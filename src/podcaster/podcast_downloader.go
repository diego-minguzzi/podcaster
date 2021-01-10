package podcaster

type PodcastDownloaderId int

type PodcastDownloaderListener interface {
    OnPodcastStarted( PodcastDownloaderId) 
    OnPodcastTerminated( PodcastDownloaderId, error)

    OnEpisodeStarted( PodcastDownloaderId, episode PodcastEpisode)
    OnEpisodeTerminated( PodcastDownloaderId, error)
}

type PodcastDownloader interface {
    Start()
    Terminate()
    AddListener( listener PodcastDownloaderListener) error 
}

func CreatePodcastDownloader( id PodcastDownloaderId,
                              podcast PodcastSource, 
                              episodes []PodcastEpisode, 
                              writerCreator EpisodeWriterCreator) (PodcastDownloader, error) {
    return nil,nil
}
