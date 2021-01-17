package podcaster

import "time"

// Describes the connection parameters.
type ConnectionParams struct {
	// The number of connections to update the feeds.
	NumFeedConnections int

	// The number of connections to download episodes.
	NumEpisodeConnections int

	DataRate KBytePerSec

	RecvTimeout time.Duration
}

type StorageParams struct {
	Url string
}

type PodcastSourcesParams struct {
	PodcastSources []PodcastSource
}
