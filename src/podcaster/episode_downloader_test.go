package podcaster

import "io/ioutil"
import "testing"
import "time"

func TestCreateEpisodeDownloader(t *testing.T) {
	t.Log("TestCreateEpisodeDownloader() started.")

	tempDirName, err := ioutil.TempDir("", "podcaster_episode_download_")
	if err != nil {
		t.Error(t.Name(), `TempDir() failed:`, err)
		return
	}
	t.Log("tempDirName:", tempDirName)
	//defer os.RemoveAll( tempDirName)

	var epDownloaderId EpisodeDownloaderId = 1

	timeout, err := time.ParseDuration("500ms")
	if err != nil {
		t.Error(t.Name(), `ParseDuration() failed:`, err)
		return
	}
	var connParams = ConnectionParams{
		NumFeedConnections:    1,
		NumEpisodeConnections: 1,
		DataRate:              KBytePerSec(200),
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
		AudioFileSize:   ByteSize(1843200),
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

	downloader, hasEpisode, err := CreateEpisodeDownloader(epDownloaderId, &connParams, &podcastSource, &episodeMeta, storage)
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
	t.Log("TestCreateEpisodeDownloader() terminated.")
}
