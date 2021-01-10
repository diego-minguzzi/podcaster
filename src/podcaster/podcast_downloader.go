package podcaster

type PodcastDownloaderId int

type PodcastDownloaderListener interface {
    OnPodcastStarted( PodcastDownloaderId) 
    OnPodcastTerminated( PodcastDownloaderId, error)

    OnEpisodeStarted( PodcastDownloaderId, episode PodcastEpisodeMeta)
    OnEpisodeTerminated( PodcastDownloaderId, error)
}

type PodcastDownloader interface {
    Start()
    Terminate()
    AddListener( listener PodcastDownloaderListener) error 
}

func CreatePodcastDownloader( id PodcastDownloaderId,
                              podcast PodcastSource, 
                              episodes []PodcastEpisodeMeta, 
                              writerCreator EpisodeWriterCreator) (PodcastDownloader, error) {
    return nil,nil
}
