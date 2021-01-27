package podcaster

import "encoding/xml"
import "fmt"
import "io"
import "io/ioutil"
import "os"
import "strings"
import "time"

// Describes the connection parameters.
type ConnectionParams struct {
	// The number of connections to update the feeds.
	NumFeedConnections int

	// The number of connections to download the episodes.
	NumEpisodeConnections int

	DownloadDataRate KBytePerSec

	RecvTimeout time.Duration
}

type StorageParams struct {
	Url string
}

type PodcasterParams struct {
	PodcastSources []PodcastSource
	Storage        StorageParams
	Connection     ConnectionParams
}

func ReadParams(reader io.Reader) (*PodcasterParams, error) {

	readData, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, fmt.Errorf("ReadAll() failed:%s", err)
	}

	var params PodcasterParams

	if err := xml.Unmarshal(readData, &params); err != nil {
		return nil, fmt.Errorf("Unmarshal() failed:%s", err)
	}
	return &params, nil
}

func WriteParams(params *PodcasterParams, writer io.Writer) error {
	serializedData, err := xml.MarshalIndent(*params, "", "  ")
	if err != nil {
		return fmt.Errorf("Marshal() failed:%s", err)
	}

	_, err = writer.Write(serializedData)
	if err != nil {
		return fmt.Errorf("Write() failed:%s", err)
	}

	return nil
}

func WriteParamsToFile(params *PodcasterParams, filepath string) error {
	file, err := os.Create(filepath)
	if err != nil {
		return fmt.Errorf("File creation failed:%s", err)
	}
	defer file.Close()

	err = WriteParams(params, file)
	if err != nil {
		return fmt.Errorf("WriteParams() failed:%s", err)
	}

	return nil
}

func WriteParamsToBuilder(params *PodcasterParams, b *strings.Builder) error {
	err := WriteParams(params, b)
	if err != nil {
		return fmt.Errorf("WriteParams() failed:%s", err)
	}

	return nil
}

func EvaluateParams(params *PodcasterParams) error {
	return nil
}
