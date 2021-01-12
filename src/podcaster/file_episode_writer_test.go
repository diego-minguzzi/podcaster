package podcaster

import "io/ioutil"
import "os"
import "testing"

//--------------------------------------------------------------------------------------------------
func TestCreateFileEpisodeWriter( t *testing.T) {
    tmpFile, err := ioutil.TempFile("","tmp_test_podcaster_file_")
    if err!=nil {
        t.Error(t.Name(),`TempFile() failed:`,err)            
        return
    }    
    tmpFilename := tmpFile.Name()        
    tmpFile.Close()    
    os.Remove(tmpFilename)

    wantFileSize:= ByteSize(256)
    wantEpisodeData := make( []byte,wantFileSize)
    
    for indx:=0; indx<int(wantFileSize); indx++ {
        wantEpisodeData[indx]= byte(indx)
	}
    epWriter,err:= CreateFileEpisodeWriter( tmpFilename, wantFileSize) 
    defer os.Remove(tmpFilename)
    if err!=nil {
        t.Error(t.Name(),`CreateFileEpisodeWriter() failed:`,err)            
        return
    }

    _,err= epWriter.Write( wantEpisodeData)
    if err!=nil {
        t.Error(t.Name(),`Write() failed:`,err)            
        return
    }

    err= epWriter.Close()
    if err!=nil {
        t.Error(t.Name(),`Close() failed:`,err)            
        return
    }

    writtenFileInfo,err:= os.Stat( tmpFilename)
    if err!=nil {
        t.Error(t.Name(),`Stat() failed:`,err)            
        return
    }
    if writtenFileInfo.Size() != int64(wantFileSize) {
        t.Error(t.Name(),`written file size got:`,writtenFileInfo.Size(),`want:`,wantFileSize)            
    }
}
