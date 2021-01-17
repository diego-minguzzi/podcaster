package podcaster

import "io"

//-------------------------------------------------------------------------------------------------
// Interface to a type that can write a podcast episode to storage.
type EpisodeWriter interface {
	io.WriteCloser

	// Closes the writer and discard what has written so far.
	CloseAndDiscard() error
}
