package podcaster

import "fmt"
import "github.com/diego-minguzzi/dmlog"

type EpisodeDownloaderId int

type EpisodeProgress struct {
    DownloadRate    KBytePerSec 
    DownloadedSize  ByteSize
    EpisodeSize     ByteSize
}

type EpisodeDownloaderListener interface {

    OnEpisodeProgress( progress *PodcastProgress)

    OnEpisodeStarted( *PodcastSource, *PodcastEpisodeMeta)
    OnEpisodeTerminated( *PodcastEpisodeMeta, error)

    Terminated( error)
}

type EpisodeDownloader interface {
    Terminate()
    TerminatedCh() ChTerminated
    AddListener( listener EpisodeDownloaderListener)  
}

func CreateEpisodeDownloader( id EpisodeDownloaderId,
                              connParams *ConnectionParams,
                              podSource *PodcastSource, 
                              episodeMeta *PodcastEpisodeMeta,
                              storage PodcastStorage) (EpisodeDownloader, bool, error) {
    dmlog.Debug("Episode title:",episodeMeta.Title)
    hasEpisode,err := storage.HasEpisode( podSource, episodeMeta)
    if err!=nil {
        return nil,false,fmt.Errorf("Storage.HasEpisode() failed:%s",err)
	} 

    if hasEpisode {
        dmlog.Debug("Already have the episode")
		return nil,hasEpisode,nil
	}
               
    episodeWriter,err:= storage.CreateEpisodeWriter( podSource, episodeMeta)
	if err!=nil {
		return nil,hasEpisode,fmt.Errorf("Storage.CreateEpisodeWriter() failed:%s",err)
	}

	var result = ctxEpisodeDownloader{
		episodeWriter: episodeWriter,
		chRequests: make( chan interface{}), 
        chResponses: make( chan interface{}), 
        chTerminated: make( chan struct{}), 
	}
    
	dmlog.Debug("Terminated correctly")
    return &result,hasEpisode,nil
}

type reqEpisodeDownloaderTerminate struct {}

type reqEpisodeDownloaderAddListener struct {
  listener EpisodeDownloaderListener
}

type ctxEpisodeDownloader struct {
	episodeWriter EpisodeWriter
    chRequests    chan interface{}
	chResponses   chan interface{}
    chTerminated  ChTerminated
}

func (c *ctxEpisodeDownloader) Terminate(){
	c.chRequests <- reqEpisodeDownloaderTerminate{}
	<- c.chResponses
}

func (c *ctxEpisodeDownloader) TerminatedCh() ChTerminated {
	return c.chTerminated
}

func (c *ctxEpisodeDownloader) AddListener( listener EpisodeDownloaderListener) {
	c.chRequests <- reqEpisodeDownloaderAddListener{ listener:listener, } 
	<- c.chResponses
}

func (c *ctxEpisodeDownloader) handler(){
	isTerminated:= false
	for !isTerminated {
		select {
			case /*request :=*/ <- c.chRequests : 
		}
	}
}