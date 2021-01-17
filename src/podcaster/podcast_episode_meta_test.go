package podcaster

import "testing"
import "time"

func TestGetAudioFileBaseName(t *testing.T) {
	type testTableItem struct {
		gotPodSource   PodcastSource
		gotEpisodeMeta PodcastEpisodeMeta
		want           string
	}

	podSource := PodcastSource{
		PodcastName:           "Hensel_Minutes",
		FeedUrl:               Url("http://feeds.feedburner.com/Hanselminutes?format=xml"),
		NumEpisodesToDownload: 10,
	}

	duration, err := parseDuration("00:38:27")
	if err != nil {
		t.Error(t.Name(), `parseDuration() failed:`, err)
		return
	}

	episodeMeta := PodcastEpisodeMeta{
		Title:           "Shipping the Sandman Doppler with Palo Alto Innovation's Alex Tramiel",
		Summary:         "",
		AudioFileUrl:    Url(""),
		AudioFileSize:   ByteSize(36916853),
		PublicationDate: DateTime(time.Date(2020, 05, 23, 10, 12, 30, 40, time.FixedZone("UTC", 0))),
		EpisodeNumber:   45,
		AudioDuration:   duration,
	}

	testTable := []testTableItem{testTableItem{podSource, episodeMeta, "hensel_minutes_20200523_0045"}}
	for indx, testItem := range testTable {
		got, err := GetAudioFileBaseName(&testItem.gotPodSource, &testItem.gotEpisodeMeta)
		if err != nil {
			t.Error(t.Name(), `GetAudioFileBaseName() failed on item:`, indx, err)
		}
		if got != testItem.want {
			t.Error(t.Name(), `GetAudioFileBaseName() failed on item:`, indx, `got:`, got, `want:`, testItem.want)
		}
		t.Log(`item`, indx, `got:`, got, `want:`, testItem.want)
	}
}
