package podcaster

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
func CreateFsStorage( params *FsStorageParams) (PodcastStorage, error) {
    rootPathInfo, err := os.Stat( params.RootDirpath)
    if err!=nil {
        return nil, fmt.Errorf("Stat() failed:%s",err)
    }
    if ! rootPathInfo.IsDir() {
        return nil, fmt.Errorf("param RootDirpath=\"%s\" is not a directory: failed.",params.RootDirpath)
    }
    var fsObj = fsStorage { rootDirpath: params.RootDirpath, }
    return &fsObj,nil
}

func (f *fsStorage) HasEpisode( podSource *PodcastSource, 
                                episodeMeta *PodcastEpisodeMeta) (bool,error){

    audioFilepath,err := f.getEpisodeAudioFilepath( podSource, episodeMeta)
    if err!=nil {
        return false, fmt.Errorf("getEpisodeAudioFilepath() failed:%s",err)
    }
    fileInfo,err := os.Stat( audioFilepath)
    if (err!=nil) || fileInfo.IsDir() {
        return false,nil
	}

    return true,nil
}

// Composes the full path to the audio file of a given episode.
func (f *fsStorage) getEpisodeAudioFilepath(podSource *PodcastSource, 
                                            episodeMeta *PodcastEpisodeMeta) (string,error) {

    fileBaseName,err := GetAudioFileBaseName( podSource, episodeMeta) 
    if err!=nil {
        return "", fmt.Errorf("GetAudioFileBaseName() failed:%s",err)
    }
    
    return path.Join( f.rootDirpath, fileBaseName+ audioFileExt ), nil
}
