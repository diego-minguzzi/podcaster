package podcaster

import "io/ioutil"
import "testing"
import "time"

func TestCreateEpisodeDownloader( t *testing.T) {
	t.Log("TestCreateEpisodeDownloader() started.")

    tempDirName, err := ioutil.TempDir("", "podcaster_episode_download_")
    if err != nil {
        t.Error(t.Name(),`TempDir() failed:`,err)            
        return
    }
	t.Log("tempDirName:",tempDirName)
    //defer os.RemoveAll( tempDirName)

	var epDownloaderId EpisodeDownloaderId= 1

    timeout,err := time.ParseDuration("500ms")
	if err!=nil {
		t.Error(t.Name(),`ParseDuration() failed:`,err)         
		return
	}
	var connParams= ConnectionParams {
		NumFeedConnections: 1,
		NumEpisodeConnections: 1,
		DataRate: KBytePerSec(100),	
		RecvTimeout: timeout,
	}

	var podcastSource= PodcastSource {
		PodcastName:"Short_and_Curly",
		FeedUrl:"https://www.abc.net.au/radio/programs/shortandcurly/feed/7388142/podcast.xml",
		NumEpisodesToDownload: 1,
	}

	audioDuration,err := time.ParseDuration("2m24s")
	if err!=nil {
		t.Error(t.Name(),`ParseDuration() failed:`,err)         
		return
	}
    var episodeMeta= PodcastEpisodeMeta{
		Title: "A little gift for you",
    	Summary: "Molly and Carl got you something for Christmas. They made it themselves, on their phones. What could it be?",
    	AudioFileUrl: "https://abcmedia.akamaized.net/radio/podcast/short_and_curly/scy-2020-12-25-s13-bonus.mp3",
    	AudioFileSize: ByteSize(58388337),
    	PublicationDate: DateTime(time.Now()),
    	EpisodeNumber: -1,
    	AudioDuration: audioDuration,
	}

	var storageParams= FsStorageParams { RootDirpath: tempDirName,}

	storage,err:= CreateFsStorage( &storageParams) 
	if err!=nil {
		t.Error(t.Name(),`CreateFsStorage() failed:`,err)         
		return
	}

	downloader,hasEpisode,err:= CreateEpisodeDownloader( epDownloaderId,&connParams,&podcastSource,&episodeMeta,storage)
	if err!=nil {
		t.Error(t.Name(),`CreateEpisodeDownloader() failed:`,err)         
		return
	}
	if hasEpisode {
		t.Error(t.Name(),`hasEpisode got:`,hasEpisode)         
	}
	if downloader==nil {
		t.Error(t.Name(),`Invalid episode downloader`)         
	}
	t.Log("TestCreateEpisodeDownloader() terminated.")
}
