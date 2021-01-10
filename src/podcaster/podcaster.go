package podcaster

import "fmt"
import "time"

// A size expressed into bytes.
type ByteSize     int
type DateTime     time.Time
type KBytePerSec  int
type Url          string

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
    OverwriteFilename     bool // Whether the    
}

type PodcastEpisode struct {
    Title                 string
    Summary               string 
    AudioFileUrl          Url     
    AudioFileSize         ByteSize
    PublicationDate       DateTime
    EpisodeNumber         int
    AudioDuration         time.Duration
}

func (p *PodcastEpisode) HasEpisodeNumber() bool { 
    return p.EpisodeNumber > 0
}

func (p *PodcastEpisode) IsEqualForTestTo( that *PodcastEpisode ) bool { 
    return (p.Title == that.Title) && 
            (p.Summary == that.Summary) &&
            (p.AudioFileUrl == that.AudioFileUrl) &&
            (p.AudioFileSize == that.AudioFileSize) &&
            (p.PublicationDate == that.PublicationDate) &&
            (p.EpisodeNumber == that.EpisodeNumber)
}

type PodcastFeedUpdate struct {
   Podcast  PodcastSource
   Episodes []PodcastEpisode
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