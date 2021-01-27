package podcaster

import "github.com/diego-minguzzi/dmlog"
import "fmt"
import "os"
import "path"

type FsStorageParams struct {
	RootDirpath string
}

type fsStorage struct {
	rootDirpath string
}

// Creates a podcast storage based on the filesystem.
func CreateFsStorage(params *FsStorageParams) (PodcastStorage, error) {
	rootPathInfo, err := os.Stat(params.RootDirpath)
	if err != nil {
		return nil, fmt.Errorf("Stat() failed:%s", err)
	}
	if !rootPathInfo.IsDir() {
		return nil, fmt.Errorf("param RootDirpath=\"%s\" is not a directory: failed.", params.RootDirpath)
	}
	var fsObj = fsStorage{rootDirpath: params.RootDirpath}
	return &fsObj, nil
}

//-------------------------------------------------------------------------------------------------
func (f *fsStorage) HasEpisode(podSource *PodcastSource,
	episodeMeta *PodcastEpisodeMeta) (bool, error) {

	audioFilepath, err := f.getEpisodeAudioFilepath(podSource, episodeMeta)
	dmlog.Debug("audioFilepath:", audioFilepath)
	if err != nil {
		return false, fmt.Errorf("getEpisodeAudioFilepath() failed:%s", err)
	}
	fileInfo, err := os.Stat(audioFilepath)
	if (err != nil) || fileInfo.IsDir() {
		return false, nil
	}

	return true, nil
}

//-------------------------------------------------------------------------------------------------
func (f *fsStorage) CreateEpisodeWriter(podSource *PodcastSource,
	episodeMeta *PodcastEpisodeMeta) (EpisodeWriter, error) {
	audioFilepath, err := f.getEpisodeAudioFilepath(podSource, episodeMeta)
	if err != nil {
		return nil, fmt.Errorf("getEpisodeAudioFilepath() failed. Got:%s", err)
	}
	dmlog.Debug("audioFilepath:", audioFilepath, "Audio file size:", episodeMeta.AudioFileSize)

	writer, err := CreateFileEpisodeWriter(audioFilepath, episodeMeta.AudioFileSize)
	if err != nil {
		return nil, fmt.Errorf("CreateFileEpisodeWriter() failed. Got:%s", err)
	}

	return writer, nil
}

// Composes the full path to the audio file of a given episode.
func (f *fsStorage) getEpisodeAudioFilepath(podSource *PodcastSource,
	episodeMeta *PodcastEpisodeMeta) (string, error) {

	podcastName := NameToSymbol(podSource.PodcastName)

	fileBaseName, err := GetAudioFileBaseName(podSource, episodeMeta)
	if err != nil {
		return "", fmt.Errorf("GetAudioFileBaseName() failed:%s", err)
	}

	return path.Join(f.rootDirpath, podcastName, fileBaseName+audioFileExt), nil
}
