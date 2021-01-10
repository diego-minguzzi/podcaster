package podcaster

import "fmt"
import "time"

// A size expressed into bytes.
type ByteSize     int
type DateTime     time.Time
type KBytePerSec  int
type Url          string

const audioFileExt = ".mp3"
const defaultFilenameLen = 60
const fileNameSeparator = rune('_') 

func (b ByteSize) String() string { return fmt.Sprintf("%d Bytes",b) }
func (d *DateTime) String() string { return (time.Time(*d)).String() }
func (k KBytePerSec) String() string { return fmt.Sprintf("%d kB/s",k) }
func (u *Url) String() string { return string(*u) }

//--------------------------------------------------------------------------------------------------
/* A podcast source, e.g. the RSS feed of the podcast*/
type PodcastSource struct {
    // Human readable name of the podcast.
    PodcastName           string     
    FeedUrl               Url     
    NumEpisodesToDownload int    
}

type PodcastFeedUpdate struct {
   Podcast  PodcastSource
   Episodes []PodcastEpisodeMeta
}

type PodcastFeedUpdateChItem struct {
    update PodcastFeedUpdate
    err    error
}

type Settings struct {
    PodcastSources      []PodcastSource
    NumMaxConnections   int
    MaxBandWidth        KBytePerSec
    StoragePath         string 
}