package podcaster

import "fmt"
import "strings"
import "time"
import "unicode"

// Metadata about a podcast episode.
type PodcastEpisodeMeta struct {
	Title           string
	Summary         string
	AudioFileUrl    Url
	AudioFileSize   ByteSize
	PublicationDate DateTime
	EpisodeNumber   int
	AudioDuration   time.Duration
}

// Whether the given episode metadata has a valid episode number.
func (p *PodcastEpisodeMeta) HasEpisodeNumber() bool {
	return p.EpisodeNumber > 0
}

// Whether two ep.metadata are equal for test purpose.
func (p *PodcastEpisodeMeta) IsEqualForTestTo(that *PodcastEpisodeMeta) bool {
	return (p.Title == that.Title) &&
		(p.Summary == that.Summary) &&
		(p.AudioFileUrl == that.AudioFileUrl) &&
		(p.AudioFileSize == that.AudioFileSize) &&
		(p.PublicationDate == that.PublicationDate) &&
		(p.EpisodeNumber == that.EpisodeNumber)
}

func getNameWithoutSpaces(name string, substSpace rune) string {

	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return substSpace
		} else {
			return r
		}
	}, name)
}

// Given a name, it returns the corresponding symbol, (no space characters), lowercase
func NameToSymbol(name string) string {
	return getNameWithoutSpaces(strings.ToLower(strings.TrimSpace(name)),
		fileNameSeparator)
}

/* Composes the audio file base name (no extension), as: PodcastTitle_yyyymmdd_nnnn  */
func GetAudioFileBaseName(podSource *PodcastSource,
	episodeMeta *PodcastEpisodeMeta) (string, error) {
	var fileName strings.Builder
	fileName.Grow(defaultFilenameLen)

	podcastName := NameToSymbol(podSource.PodcastName)
	var err error = nil
	_, err = fileName.WriteString(podcastName)
	if err != nil {
		return "", fmt.Errorf("WriteString() failed:%s", err)
	}

	_, err = fileName.WriteRune(fileNameSeparator)
	if err != nil {
		return "", fmt.Errorf("WriteRune() failed:%s", err)
	}

	pubDate := (*time.Time)(&episodeMeta.PublicationDate)
	timestamp := fmt.Sprintf("%04d%02d%02d", pubDate.Year(),
		pubDate.Month(),
		pubDate.Day())
	_, err = fileName.WriteString(timestamp)
	if err != nil {
		return "", fmt.Errorf("WriteString() failed:%s", err)
	}

	_, err = fileName.WriteRune(fileNameSeparator)
	if err != nil {
		return "", fmt.Errorf("WriteRune() failed:%s", err)
	}

	episodeNumber := 0
	if episodeMeta.HasEpisodeNumber() {
		episodeNumber = episodeMeta.EpisodeNumber
	}
	_, err = fileName.WriteString(fmt.Sprintf("%04d", episodeNumber))
	if err != nil {
		return "", fmt.Errorf("WriteString() failed:%s", err)
	}

	return fileName.String(), nil
}
