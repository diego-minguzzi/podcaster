package podcaster

// Interface to a podcast storage.
type PodcastStorage interface {
    HasEpisode( podSource PodcastSource, episode PodcastEpisode) (bool,error)
    MostRecentEpisode() (PodcastEpisode, error)
}
