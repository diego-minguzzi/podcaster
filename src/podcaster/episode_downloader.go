package podcaster

import "fmt"
import "github.com/diego-minguzzi/dmlog"
import "io"
import "io_utils"
import "net/http"
import "time"

type EpisodeDownloaderId int

type EpisodeProgress struct {
	DownloadRate   KBytePerSec
	DownloadedSize ByteSize
	EpisodeSize    ByteSize
}

type EpisodeDownloaderListener interface {
	OnEpisodeProgress(progress *PodcastProgress)

	OnEpisodeStarted(*PodcastSource, *PodcastEpisodeMeta)
	OnEpisodeTerminated(*PodcastEpisodeMeta, error)

	Terminated(error)
}

type EpisodeDownloader interface {
	Terminate()
	TerminatedCh() ChTerminated
	AddListener(listener EpisodeDownloaderListener)
}

const durEpisodeDownloaderCycle = 1 * time.Second

func CreateEpisodeDownloader(id EpisodeDownloaderId,
	connParams *ConnectionParams,
	podSource *PodcastSource,
	episodeMeta *PodcastEpisodeMeta,
	storage PodcastStorage) (EpisodeDownloader, bool, error) {
	dmlog.Debug("Episode title:", episodeMeta.Title)
	hasEpisode, err := storage.HasEpisode(podSource, episodeMeta)
	if err != nil {
		return nil, false, fmt.Errorf("Storage.HasEpisode() failed:%s", err)
	}

	if hasEpisode {
		dmlog.Debug("Already have the episode")
		return nil, hasEpisode, nil
	}

	episodeWriter, err := storage.CreateEpisodeWriter(podSource, episodeMeta)
	if err != nil {
		return nil, hasEpisode, fmt.Errorf("Storage.CreateEpisodeWriter() failed:%s", err)
	}

	var ctx = ctxEpisodeDownloader{
		connParams:    connParams,
		episodeWriter: episodeWriter,
		chRequests:    make(chan interface{}),
		chResponses:   make(chan interface{}),
		chTerminated:  make(chan struct{}),
		episodeMeta:   episodeMeta,
		listeners:     make([]EpisodeDownloaderListener, 0),
		buffer:        make([]byte, 1024*int(connParams.DataRate)),
	}

	ctx.episodeProgress.EpisodeSize = episodeMeta.AudioFileSize
	dmlog.Debug("ctx.episodeProgress.EpisodeSize:", int(ctx.episodeProgress.EpisodeSize))
	go episodeDownloaderHandler(&ctx)

	dmlog.Debug("Terminated correctly")
	return &ctx, hasEpisode, nil
}

type reqEpisodeDownloaderTerminate struct{}

type reqEpisodeDownloaderAddListener struct {
	listener EpisodeDownloaderListener
}

type ctxEpisodeDownloader struct {
	connParams      *ConnectionParams
	episodeWriter   EpisodeWriter
	chRequests      chan interface{}
	chResponses     chan interface{}
	chTerminated    ChTerminated
	episodeMeta     *PodcastEpisodeMeta
	listeners       []EpisodeDownloaderListener
	buffer          []byte
	episodeProgress EpisodeProgress
}

func (c *ctxEpisodeDownloader) Terminate() {
	c.chRequests <- reqEpisodeDownloaderTerminate{}
	<-c.chResponses
}

func (c *ctxEpisodeDownloader) TerminatedCh() ChTerminated {
	return c.chTerminated
}

func (c *ctxEpisodeDownloader) AddListener(listener EpisodeDownloaderListener) {
	c.chRequests <- reqEpisodeDownloaderAddListener{listener: listener}
	<-c.chResponses
}

func episodeDownloaderHandler(ctx *ctxEpisodeDownloader) {
	defer dmlog.MethodStartEnd()()
	isTerminated := false

	defer close(ctx.chTerminated)

	dmlog.Debug("Trying to download:", ctx.episodeMeta.AudioFileUrl)
	resp, err := http.Get(string(ctx.episodeMeta.AudioFileUrl))
	if err != nil {
		dmlog.Error("http.Get() failed:", err)
		ctx.episodeWriter.CloseAndDiscard()
		episodeDownloaderNotifyTerminated(ctx, err)
		return
	}

	var audioFileReader = io_utils.NewTimeoutReader(resp.Body, ctx.connParams.RecvTimeout)
	defer audioFileReader.Close()

	var tmNextCycle = time.Now().Add(durEpisodeDownloaderCycle)
	var durToNextDownload time.Duration
	for !isTerminated {
		dmlog.Debug("Cycle started")
		var tmStartCycle = time.Now()
		if tmNextCycle.Before(tmStartCycle) {
			durToNextDownload = 0 * time.Second
		} else {
			durToNextDownload = tmNextCycle.Sub(tmStartCycle)
		}

		select {
		// case  <-ctx.chRequests:
		case <-time.After(durToNextDownload):
			{
				dmlog.Debug("Download timer expired.")
				tmNextCycle = tmNextCycle.Add(durEpisodeDownloaderCycle)
				isCompleted, err := downloadNextChunk(ctx, audioFileReader)
				if err != nil {
					ctx.episodeWriter.CloseAndDiscard()
					episodeDownloaderNotifyTerminated(ctx, err)
					isTerminated = true
				}
				if isCompleted {
					ctx.episodeWriter.Close()
					episodeDownloaderNotifyTerminated(ctx, nil)
					isTerminated = true
				}
			}
		}
	}
}

//-------------------------------------------------------------------------------------------------
func downloadNextChunk(ctx *ctxEpisodeDownloader, audioFileReader io.ReadCloser) (bool, error) {
	n, err := audioFileReader.Read(ctx.buffer)
	ctx.episodeProgress.DownloadedSize += ByteSize(n)
	dmlog.Debug("ctx.episodeProgress.DownloadedSize:", int(ctx.episodeProgress.DownloadedSize))

	if (err != nil) && (err == io.EOF) {
		ctx.episodeWriter.Write(ctx.buffer)
		dmlog.Debug("End of file detected.")
		return true, nil
	} else if err != nil {
		dmlog.Error("Error:", err)
		return true, err
	}

	ctx.episodeWriter.Write(ctx.buffer)

	return false, nil
}

func episodeDownloaderNotifyTerminated(ctx *ctxEpisodeDownloader, err error) {
	for _, listener := range ctx.listeners {
		listener.Terminated(err)
	}
}
