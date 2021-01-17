package podcaster

import "fmt"
import "github.com/diego-minguzzi/dmlog"
import "os"

/* Creates a File Episode Writer that writes filepath a size of fileSize bytes.*/
func CreateFileEpisodeWriter(filepath string, fileSize ByteSize) (EpisodeWriter, error) {
	if fileSize < 0 {
		return nil, fmt.Errorf("invalid fileSize. Got:%d", fileSize)
	}
	newFile, err := os.Create(filepath)
	if err != nil {
		return nil, fmt.Errorf("os.Create() failed on %s:%s", filepath, err)
	}

	result := fileEpisodeWriter{
		episodeData: make([]byte, 0, fileSize),
		file:        newFile,
	}

	return &result, nil
}

//-------------------------------------------------------------------------------------------------
type fileEpisodeWriter struct {
	episodeData []byte
	file        *os.File
}

//-------------------------------------------------------------------------------------------------
/* Implements the WriteCloser interface.
   Panic in case Write() is called after Close() or CloseAndDiscard() */
func (f *fileEpisodeWriter) Write(p []byte) (n int, err error) {
	if nil == f.file {
		panic("Write() called on invalid writer.")
	}
	f.episodeData = append(f.episodeData, p...)
	return len(p), nil
}

//-------------------------------------------------------------------------------------------------
/* Implements the WriteCloser interface */
func (f *fileEpisodeWriter) Close() error {
	_, err := f.file.Write(f.episodeData)
	if err != nil {
		return fmt.Errorf("os.Write() failed:%s", err)
	}
	return f.innerClose()
}

//-------------------------------------------------------------------------------------------------
func (f *fileEpisodeWriter) CloseAndDiscard() error {
	return f.innerClose()
}

//-------------------------------------------------------------------------------------------------
func (f *fileEpisodeWriter) innerClose() error {
	if nil == f.file {
		dmlog.Warn("Writer already closed.")
		return nil
	}
	err := f.file.Close()
	f.file = nil
	if err != nil {
		return fmt.Errorf("os.Close() failed:%s", err)
	}
	return nil
}
