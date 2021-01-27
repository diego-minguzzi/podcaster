package podcaster

import "strings"
import "testing"
import "time"

func createTestParams() *PodcasterParams {
	var connParams = ConnectionParams{
		NumFeedConnections:    1,
		NumEpisodeConnections: 3,
		DownloadDataRate:      KBytePerSec(1024),
		RecvTimeout:           5 * time.Second,
	}

	var storageParams = StorageParams{
		Url: "filesystem:///home/minguzzi/tmp",
	}

	var podcastSource = PodcastSource{
		PodcastName:           "Freakonomics_Radio",
		FeedUrl:               "http://feeds.feedburner.com/freakonomicsradio",
		NumEpisodesToDownload: 5,
	}

	var podcastSource2 = PodcastSource{
		PodcastName:           "Imani_State_of_Mind",
		FeedUrl:               "https://omny.fm/shows/imani-state-of-mind/playlists/podcast.rss",
		NumEpisodesToDownload: 10,
	}

	var podcastParams = PodcasterParams{
		PodcastSources: []PodcastSource{podcastSource, podcastSource2},
		Storage:        storageParams,
		Connection:     connParams,
	}

	return &podcastParams
}

//-------------------------------------------------------------------------------------------------
func TestWriteParamsToBuilder(t *testing.T) {
	podcastParams := createTestParams()
	if podcastParams == nil {
		t.Error(t.Name(), `createTestParams() failed.`)
		return
	}
	var serializedParams strings.Builder

	err := WriteParamsToBuilder(podcastParams, &serializedParams)
	if err != nil {
		t.Error(t.Name(), `WriteParamsToBuilder() failed:`, err)
		return
	}

	t.Log("serializedParams:\n", serializedParams.String())
}

//-------------------------------------------------------------------------------------------------
func TestReadWriteParams(t *testing.T) {
	podcastParams := createTestParams()
	if podcastParams == nil {
		t.Error(t.Name(), `createTestParams() failed.`)
		return
	}
	var serializedParams strings.Builder

	err := WriteParamsToBuilder(podcastParams, &serializedParams)
	if err != nil {
		t.Error(t.Name(), `WriteParamsToBuilder() failed:`, err)
		return
	}

	t.Log("serializedParams:\n", serializedParams.String())

	reader := strings.NewReader(serializedParams.String())
	if reader == nil {
		t.Error(t.Name(), `Invalid reader.`)
		return
	}

	gotParams,err := ReadParams( reader)
	if err != nil {
		t.Error(t.Name(), `ReadParams() failed:`, err)
		return
	}

	gotNumSources := len(gotParams.PodcastSources)
	wantNumSources :=len(podcastParams.PodcastSources)

	if gotNumSources!=gotNumSources {
		t.Error(t.Name(), `gotNumSources:`, gotNumSources,`wantNumSources:`,wantNumSources)
	}

	t.Log(`gotNumSources:`, gotNumSources,`wantNumSources:`,wantNumSources)

	var gotSerializedParams strings.Builder

	err = WriteParamsToBuilder(gotParams, &gotSerializedParams)
	if err != nil {
		t.Error(t.Name(), `WriteParamsToBuilder() failed:`, err)
		return
	}
	t.Log("gotSerializedParams:\n", gotSerializedParams.String())
}
