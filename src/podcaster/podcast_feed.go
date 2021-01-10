package podcaster

import "fmt"
import "github.com/diego-minguzzi/dmlog"
import "log"
import "net/http"
import "sync"

type PodcastFeedUpdater interface {
    PodcastFeedItemsChan() chan PodcastFeedUpdateChItem
    Wait()
    Terminate()
}

const defaultNumEpisodesCap int = 100

//-------------------------------------------------------------------------------------------------
type concretePodcastFeed struct {
    numMaxConnections       int
    chPodcastSources        chan PodcastSource
    chPodcastFeedItems      chan PodcastFeedUpdateChItem 
    chTerminate             chan struct{}
    
    waitUpdaters            sync.WaitGroup
}

//-------------------------------------------------------------------------------------------------
/* Executes an update from the given podcast sources, using at most numMaxConnections connections  
    Returns a feed updater interface.
    numMaxConnections must be greater than zero, otherwise panic. */
func UpdateFeeds( podSources []PodcastSource, numMaxConnections int ) PodcastFeedUpdater {
    defer dmlog.MethodStartEnd()()
    if numMaxConnections<=0 {
      log.Panic("invalid numMaxConnections:got",numMaxConnections)
	}

    var ctx concretePodcastFeed
    ctx.numMaxConnections= numMaxConnections
    ctx.chPodcastSources= make( chan PodcastSource, len(podSources))
    ctx.chPodcastFeedItems= make( chan PodcastFeedUpdateChItem, len(podSources))
    ctx.chTerminate = make( chan struct{})
    
    for _, podcastSource := range podSources {
        ctx.chPodcastSources <- podcastSource
	}

    for indx:=0; indx<numMaxConnections; indx++ {
        ctx.waitUpdaters.Add(1)
        go podcastFeedUpdater( &ctx)
	}
    
    return &ctx
}

//-------------------------------------------------------------------------------------------------
func podcastFeedUpdater( ctx *concretePodcastFeed) {
    defer dmlog.MethodStartEnd()()
    defer ctx.waitUpdaters.Done()


    for isTerminated:=false; !isTerminated;  {
        select {
            case podcastSource := <- ctx.chPodcastSources : {
                dmlog.Debug("Updating feed from Url:", podcastSource.FeedUrl.String())
                feedUpdate := PodcastFeedUpdateChItem {
                    err: nil, 
                    update: PodcastFeedUpdate {
                        Podcast: podcastSource,
                        Episodes: make( []PodcastEpisodeMeta, 0, defaultNumEpisodesCap),
                    },
                }
                resp, err := http.Get(podcastSource.FeedUrl.String())
                dmlog.Debug("Got response from ", podcastSource.FeedUrl.String())  
                if err!=nil {
                    dmlog.Error("http.Get():", err)
                    feedUpdate.err = fmt.Errorf("error while getting from %s:%s",podcastSource.FeedUrl.String(),err)
                    ctx.chPodcastFeedItems <- feedUpdate 
			    }        
			    defer resp.Body.Close()  
                // Parse the body

                ctx.chPodcastFeedItems <- feedUpdate             
			}
            case <- ctx.chTerminate : {
                dmlog.Debug("Got terminate from chTerminate")
                isTerminated= true
			}

            default : {
                dmlog.Debug("Nothing to do: terminating")
                isTerminated= true 
			}
        }
	}
}

//-------------------------------------------------------------------------------------------------
func (c *concretePodcastFeed) PodcastFeedItemsChan() chan PodcastFeedUpdateChItem {
    return c.chPodcastFeedItems
}

//-------------------------------------------------------------------------------------------------
func (c *concretePodcastFeed) Wait() {
    c.waitUpdaters.Wait()    
}

//-------------------------------------------------------------------------------------------------
func (c *concretePodcastFeed) Terminate() {
    close( c.chTerminate)    
}


