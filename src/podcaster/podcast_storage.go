package podcaster

// Interface to a podcast storage.
type PodcastStorage interface {
	// Whether a given episode exists in the storage.
	HasEpisode(podSource *PodcastSource,
		episode *PodcastEpisodeMeta) (bool, error)
	/* Returns an episode writer in order to save the given episode, related to the given
	   podcast source */
	CreateEpisodeWriter(podSource *PodcastSource, episodeMeta *PodcastEpisodeMeta) (EpisodeWriter, error)
}
