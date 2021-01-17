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

func (e *EpisodeProgress) ToPercentage() float64 {
	if e.EpisodeSize == 0 {
		return 0.0
	}
	return float64((100.0 * e.DownloadedSize) / e.EpisodeSize)
}

type EpisodeDownloaderListener interface {
	OnEpisodeStarted(*PodcastSource, *PodcastEpisodeMeta)
	OnEpisodeProgress(*EpisodeProgress)
	OnEpisodeTerminated(*PodcastEpisodeMeta, error)
}

type EpisodeDownloader interface {
	// Forces the termination of the download
	Terminate() error

	AddListener(listener EpisodeDownloaderListener) error

	// When the returned channel is closed, the episode dowloader is finished.
	TerminatedCh() ChTerminated
}

const durEpisodeDownloaderCycle = 1 * time.Second

// Describes the terminate request.
type reqEpisodeDownloaderTerminate struct{}

// Describes the add listener request.
type reqEpisodeDownloaderAddListener struct {
	listener EpisodeDownloaderListener
}

type replyEpisodeDownloader struct{ err error }

func CreateEpisodeDownloader(id EpisodeDownloaderId,
	connParams *ConnectionParams,
	podSource *PodcastSource,
	episodeMeta *PodcastEpisodeMeta,
	storage PodcastStorage,
	listeners ...EpisodeDownloaderListener) (EpisodeDownloader, bool, error) {

	defer dmlog.MethodStartEnd()()
	hasEpisode, err := storage.HasEpisode(podSource, episodeMeta)
	if err != nil {
		return nil, false, fmt.Errorf("Storage.HasEpisode() failed:%s", err)
	}

	if hasEpisode {
		dmlog.Debug("Already have the episode with title:", episodeMeta.Title)
		return nil, hasEpisode, nil
	}

	episodeWriter, err := storage.CreateEpisodeWriter(podSource, episodeMeta)
	if err != nil {
		return nil, hasEpisode, fmt.Errorf("Storage.CreateEpisodeWriter() failed:%s", err)
	}

	var ctx = ctxEpisodeDownloader{
		connParams:      connParams,
		episodeWriter:   episodeWriter,
		chRequests:      make(chan interface{}),
		chReplies:       make(chan interface{}),
		chTerminated:    make(chan struct{}),
		podcastSource:   podSource,
		episodeMeta:     episodeMeta,
		listeners:       make([]EpisodeDownloaderListener, 0, len(listeners)),
		buffer:          make([]byte, int(connParams.DataRate.toBytePerSec())),
		tmDownloadStart: time.Now(),
	}

	for _, listener := range listeners {
		ctx.listeners = append(ctx.listeners, listener)
	}

	ctx.episodeProgress.EpisodeSize = episodeMeta.AudioFileSize
	go episodeDownloaderHandler(&ctx)

	return &ctx, hasEpisode, nil
}

type ctxEpisodeDownloader struct {
	connParams      *ConnectionParams
	episodeWriter   EpisodeWriter
	chRequests      chan interface{} // Channel where requests are received.
	chReplies       chan interface{} // The channel where replies are sent.
	chTerminated    ChTerminated     // This channel is closed upon termination.
	podcastSource   *PodcastSource
	episodeMeta     *PodcastEpisodeMeta
	listeners       []EpisodeDownloaderListener
	buffer          []byte
	episodeProgress EpisodeProgress
	tmDownloadStart time.Time // The time when the download started.
	isTerminate     bool
}

//-------------------------------------------------------------------------------------------------
func (c *ctxEpisodeDownloader) Terminate() error {
	defer dmlog.MethodExecuted()
	c.chRequests <- reqEpisodeDownloaderTerminate{}
	reply := <-c.chReplies
	switch reply := reply.(type) {
	case replyEpisodeDownloader:
		return reply.err
	}

	panic("Unexpected case detected.")
}

//-------------------------------------------------------------------------------------------------
func (c *ctxEpisodeDownloader) AddListener(listener EpisodeDownloaderListener) error {
	defer dmlog.MethodExecuted()
	c.chRequests <- reqEpisodeDownloaderAddListener{listener: listener}
	reply := <-c.chReplies
	switch reply := reply.(type) {
	case replyEpisodeDownloader:
		return reply.err
	}

	panic("Unexpected case detected.")
}

//-------------------------------------------------------------------------------------------------
func (c *ctxEpisodeDownloader) TerminatedCh() ChTerminated {
	defer dmlog.MethodExecuted()
	return c.chTerminated
}

//-------------------------------------------------------------------------------------------------
func episodeDownloaderHandler(ctx *ctxEpisodeDownloader) {
	defer dmlog.MethodStartEnd()()
	defer close(ctx.chTerminated)

	episodeDownloaderEpisodeStarted(ctx)

	dmlog.Debug("Downloading episode from URL:", ctx.episodeMeta.AudioFileUrl)
	resp, err := http.Get(string(ctx.episodeMeta.AudioFileUrl))
	if err != nil {
		dmlog.Error("http.Get() failed:", err)
		ctx.episodeWriter.CloseAndDiscard()
		episodeDownloaderEpisodeTerminated(ctx, err)
		return
	}

	var audioFileReader = io_utils.NewTimeoutReader(resp.Body, ctx.connParams.RecvTimeout)
	defer audioFileReader.Close()

	var tmNextCycle = time.Now().Add(durEpisodeDownloaderCycle)
	var durToNextDownload time.Duration
	ctx.isTerminate = false

	for !ctx.isTerminate {
		var tmStartCycle = time.Now()
		if tmNextCycle.Before(tmStartCycle) {
			durToNextDownload = 0 * time.Second
		} else {
			durToNextDownload = tmNextCycle.Sub(tmStartCycle)
		}

		select {
		case request := <-ctx.chRequests:
			{
				ctx.chReplies <- handleRequest(ctx, request)
			}
		case <-time.After(durToNextDownload):
			{
				tmNextCycle = tmNextCycle.Add(durEpisodeDownloaderCycle)
				isCompleted, err := downloadNextChunk(ctx, audioFileReader)
				episodeDownloaderEpisodeProgress(ctx, &ctx.episodeProgress)
				if err != nil {
					err = fmt.Errorf("chunk download failed:%s", err)
					ctx.episodeWriter.CloseAndDiscard()
				}
				if isCompleted {
					ctx.episodeWriter.Close()
				}
				if isCompleted || (err != nil) {
					episodeDownloaderEpisodeTerminated(ctx, err)
					ctx.isTerminate = true
				}
			}
		}
	}
}

//-------------------------------------------------------------------------------------------------
func handleRequest(ctx *ctxEpisodeDownloader, request interface{}) replyEpisodeDownloader {

	switch request := request.(type) {
	case reqEpisodeDownloaderTerminate:
		{
			dmlog.Debug("Terminate request received.")
			ctx.isTerminate = true
			ctx.episodeWriter.CloseAndDiscard()
			return replyEpisodeDownloader{err: nil}
		}

	case reqEpisodeDownloaderAddListener:
		{
			dmlog.Debug("Add listener request received.")
			if request.listener == nil {
			} else {
				ctx.listeners = append(ctx.listeners, request.listener)
			}
			return replyEpisodeDownloader{err: nil}
		}
	}
	return replyEpisodeDownloader{fmt.Errorf("Unknown request type.")}
}

//-------------------------------------------------------------------------------------------------
func downloadNextChunk(ctx *ctxEpisodeDownloader, audioFileReader io.ReadCloser) (bool, error) {
	chunkSizeTarget := int(ctx.connParams.DataRate.toBytePerSec())
	var err error
	var n int
	// Tries to accumulate enough size to reach the target of the period.
	for accumSize := 0; (accumSize < chunkSizeTarget) && (err == nil); accumSize += n {
		n, err = audioFileReader.Read(ctx.buffer)
		ctx.episodeProgress.DownloadedSize += ByteSize(n)

		if (err != nil) && (err == io.EOF) {
			ctx.episodeWriter.Write(ctx.buffer)
			dmlog.Debug("End of file detected.")
			return true, nil
		} else if err != nil {
			dmlog.Error("Error:", err)
			return true, fmt.Errorf("io.Read() failed:%s", err)
		}

		ctx.episodeWriter.Write(ctx.buffer)
	}

	downloadSecs := time.Now().Sub(ctx.tmDownloadStart).Seconds()
	if downloadSecs < epsilon {
		dmlog.Error("Invalid downloadSecs:", downloadSecs)
		ctx.episodeProgress.DownloadRate = KBytePerSec(0)
	} else {
		downloadRateKBps := float64(ctx.episodeProgress.DownloadedSize) / (1024.0 * downloadSecs)
		ctx.episodeProgress.DownloadRate = KBytePerSec(downloadRateKBps)
	}

	return false, nil
}

// Notifies all listeners of the state of an episode download.
func episodeDownloaderEpisodeProgress(ctx *ctxEpisodeDownloader, episodeProgress *EpisodeProgress) {
	for _, listener := range ctx.listeners {
		listener.OnEpisodeProgress(episodeProgress)
	}
}

// Notifies all listeners that an episode downloading started.
func episodeDownloaderEpisodeStarted(ctx *ctxEpisodeDownloader) {
	for _, listener := range ctx.listeners {
		listener.OnEpisodeStarted(ctx.podcastSource, ctx.episodeMeta)
	}
}

// Notifies all listeners that an episode downloading is terminated.
func episodeDownloaderEpisodeTerminated(ctx *ctxEpisodeDownloader, err error) {
	for _, listener := range ctx.listeners {
		listener.OnEpisodeTerminated(ctx.episodeMeta, err)
	}
}
