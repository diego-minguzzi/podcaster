package podcaster

// Interface to a podcast storage.
type PodcastStorage interface {
    // Whether a given episode exists in the storage.
    HasEpisode( podSource *PodcastSource, 
                episode *PodcastEpisodeMeta) (bool,error)
    //MostRecentEpisodeMeta() (PodcastEpisodeMeta, error)
}
