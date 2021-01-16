package podcaster


type PodcastDownloader interface {
    Start()
    Terminate()
    // AddListener( listener PodcastDownloaderListener) error 
    Wait()
}

func CreatePodcastDownloader() (PodcastDownloader,error){
    return nil, nil
}


