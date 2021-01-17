package podcaster

import "io/ioutil"
import "os"
import "testing"
import "time"

type myEpisodeDownloaderListener struct {
	t                              *testing.T
	wasOnEpisodeStartedExecuted    bool
	wasOnEpisodeProgressExecuted   bool
	wasOnEpisodeTerminatedExecuted bool
	errOnTermination               error
}

func NewMyEpisodeDownloaderListener(t *testing.T) *myEpisodeDownloaderListener {
	obj := myEpisodeDownloaderListener{
		t: t,
		wasOnEpisodeStartedExecuted:    false,
		wasOnEpisodeProgressExecuted:   false,
		wasOnEpisodeTerminatedExecuted: false,
	}
	return &obj
}

func (m *myEpisodeDownloaderListener) OnEpisodeStarted(podSource *PodcastSource, episodeMeta *PodcastEpisodeMeta) {
	m.t.Log("OnEpisodeStarted():", podSource.PodcastName, episodeMeta.Title)
	m.wasOnEpisodeStartedExecuted = true
}

func (m *myEpisodeDownloaderListener) OnEpisodeProgress(progress *EpisodeProgress) {
	m.t.Log("OnEpisodeProgress():", progress.ToPercentage(), "% ", progress.DownloadRate)
	m.wasOnEpisodeProgressExecuted = true
}

func (m *myEpisodeDownloaderListener) OnEpisodeTerminated(_ *PodcastEpisodeMeta, err error) {
	m.t.Log("OnEpisodeTerminated()")
	m.errOnTermination = err
	m.wasOnEpisodeTerminatedExecuted = true
}

//-------------------------------------------------------------------------------------------------
func TestCreateEpisodeDownloaderWithListener(t *testing.T) {
	t.Log("TestCreateEpisodeDownloaderWithListener() started.")

	tempDirName, err := ioutil.TempDir("", "podcaster_episode_download_")
	if err != nil {
		t.Error(t.Name(), `TempDir() failed:`, err)
		return
	}
	t.Log("tempDirName:", tempDirName)
	defer os.RemoveAll(tempDirName)

	var epDownloaderId EpisodeDownloaderId = 1

	timeout, err := time.ParseDuration("1000ms")
	if err != nil {
		t.Error(t.Name(), `ParseDuration() failed:`, err)
		return
	}
	var connParams = ConnectionParams{
		NumFeedConnections:    1,
		NumEpisodeConnections: 1,
		DataRate:              KBytePerSec(100),
		RecvTimeout:           timeout,
	}

	var podcastSource = PodcastSource{
		PodcastName:           "Under_The_Influence",
		FeedUrl:               "https://www.cbc.ca/podcasting/includes/undertheinfluence.xml",
		NumEpisodesToDownload: 1,
	}

	audioDuration, err := time.ParseDuration("2m24s")
	if err != nil {
		t.Error(t.Name(), `ParseDuration() failed:`, err)
		return
	}
	var episodeMeta = PodcastEpisodeMeta{
		Title:           "Under the Influence is back January 7th",
		Summary:         "We’ve got a fun 2021 season planned for you. Here’s a sneak peek at what's in store...",
		AudioFileUrl:    "https://cbc.mc.tritondigital.com/CBC_UNDERTHEINFLUENCE_P/media/undertheinfluence-qRhSXiLt-20210106.mp3",
		AudioFileSize:   ByteSize(1369011),
		PublicationDate: DateTime(time.Now()),
		EpisodeNumber:   -1,
		AudioDuration:   audioDuration,
	}

	var storageParams = FsStorageParams{RootDirpath: tempDirName}

	storage, err := CreateFsStorage(&storageParams)
	if err != nil {
		t.Error(t.Name(), `CreateFsStorage() failed:`, err)
		return
	}

	myListener := NewMyEpisodeDownloaderListener(t)

	downloader, hasEpisode, err := CreateEpisodeDownloader(epDownloaderId, &connParams, &podcastSource, &episodeMeta, storage, myListener)
	if err != nil {
		t.Error(t.Name(), `CreateEpisodeDownloader() failed:`, err)
		return
	}
	if hasEpisode {
		t.Error(t.Name(), `hasEpisode got:`, hasEpisode)
	}
	if downloader == nil {
		t.Error(t.Name(), `Invalid episode downloader`)
	}

	<-downloader.TerminatedCh()
	if !myListener.wasOnEpisodeStartedExecuted {
		t.Error(t.Name(), `OnEpisodeStarted() not executed`)
	}

	if !myListener.wasOnEpisodeProgressExecuted {
		t.Error(t.Name(), `OnEpisodeProgress() not executed`)
	}

	if !myListener.wasOnEpisodeTerminatedExecuted {
		t.Error(t.Name(), `OnEpisodeTerminated() not executed`)
	}

	if myListener.errOnTermination != nil {
		t.Error(t.Name(), `Error on termination:`, myListener.errOnTermination)
	}
}
