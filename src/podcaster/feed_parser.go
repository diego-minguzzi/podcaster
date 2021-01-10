package podcaster

import "bytes"
import "encoding/xml"
import "fmt"
import "github.com/diego-minguzzi/dmlog"
import "io"
import "io/ioutil"
import "strings"
import "time"

//-------------------------------------------------------------------------------------------------
// rssType is the root data type used to unmarshal the XML.
type rssType struct {  // item
    Channel     rssChannelType    `xml:"channel"`   
}

type rssChannelType struct {  // item
    Items       []rssItemType     `xml:"item"`   
}

type rssItemType struct {
    Title         string            `xml:"title"`
    Summary       string            `xml:"itunes_summary"` // Xml element substitured by xmlBuffReplacements 
    Enclosure     rssEnclosureType  `xml:"enclosure"`    
    EpisodeNumber int               `xml:"itunes_episode"` // Xml element substitured by xmlBuffReplacements 
    Duration      string            `xml:"itunes_duration"` // Xml element substitured by xmlBuffReplacements 
    PubDate       string            `xml:"pubDate"`
}

type rssEnclosureType struct {
    Url         string  `xml:"url,attr"`
    Type        string  `xml:"type,attr"`
    Length      int     `xml:"length,attr"`    
}

//-------------------------------------------------------------------------------------------------
const audioMpegType string = `audio/mpeg`
const durationSeparator string = `:` // Separator in duration expressed as HH:MM:SS
const tagDuration string = `itunes_duration`
const tagEpisodeNumber string = `itunes_episode`
const tagSummary string = `itunes_summary`

// Describes a replacement of XML data, from old to new.
type xmlBuffReplacement struct {
    oldBuff []byte
    newBuff []byte
}

/* A table of byte slice substitutions: it was not found a way to parse the elements in the 
 * namespace itunes, like itunes:episode, etc, so the element names are substituted with
 * elements that are easier to unmarshal. */
var xmlBuffReplacements []xmlBuffReplacement

func init() {
    xmlBuffReplacements = []xmlBuffReplacement {
        xmlBuffReplacement{ []byte("<itunes:duration>"),  []byte("<"+ tagDuration +">")},
        xmlBuffReplacement{ []byte("</itunes:duration>"), []byte("</"+ tagDuration +">")},
        xmlBuffReplacement{ []byte("<itunes:episode>"),   []byte("<"+ tagEpisodeNumber +">")},
        xmlBuffReplacement{ []byte("</itunes:episode>"),  []byte("</"+ tagEpisodeNumber +">")},
        xmlBuffReplacement{ []byte("<itunes:summary>"),   []byte("<"+ tagSummary +">")}, 
        xmlBuffReplacement{ []byte("</itunes:summary>"),  []byte("</"+ tagSummary +">")}, 
	}
}

//-------------------------------------------------------------------------------------------------
func executeXmlReplacements( xmlData []byte, replacements []xmlBuffReplacement) []byte {
    var result = xmlData
    for _, replacement := range replacements {
        result = bytes.Replace( result, replacement.oldBuff, replacement.newBuff, -1)
	}
    return result
}

//-------------------------------------------------------------------------------------------------
func ParseRssFeedAllEpisodes( podSource PodcastSource,
                              reader io.Reader ) ([]PodcastEpisode,error) {
    return ParseRssFeed(podSource, 
                        reader, 
                        func( _ *PodcastSource, _ *PodcastEpisode) bool { return true })
}

//-------------------------------------------------------------------------------------------------
func ParseRssFeed(podSource PodcastSource,
                  reader io.Reader,
                  funcAcceptEpisode func(*PodcastSource, *PodcastEpisode) bool ) ([]PodcastEpisode,error) {
    defer dmlog.MethodStartEnd()()
    readData, err := ioutil.ReadAll( reader)
    if err!=nil {
      return nil,fmt.Errorf("ReadAll() failed:%s",err) 
    }

    // Replaces XML strings that cannot be parsed correctly.
    replacedData := executeXmlReplacements( readData, xmlBuffReplacements)

    var rssContent rssType
    if err := xml.Unmarshal( replacedData, &rssContent); err!=nil {
        return nil,fmt.Errorf("Unmarshal() failed:%s",err)
    }

    var podEpisodes = make( []PodcastEpisode, 0, len(rssContent.Channel.Items) )
    for _,rssItem := range rssContent.Channel.Items { 
        enclosure := &rssItem.Enclosure
        trimmedType := strings.TrimSpace(enclosure.Type)
        dmlog.Debug("rssItem.Title:",rssItem.Title)   
        dmlog.Debug("rssItem.Summary:",rssItem.Summary)   
        dmlog.Debug("rssItem.EpisodeNumber:",rssItem.EpisodeNumber)   
        dmlog.Debug("rssItem.Duration:",rssItem.Duration)   
        dmlog.Debug("rssItem.PubDate:",rssItem.PubDate)   

        if audioMpegType != trimmedType {
            dmlog.Warn("Unexpected enclosure type:",trimmedType)
            continue
		}

        pubDate,err := time.Parse( time.RFC1123Z, rssItem.PubDate) 
        if err!=nil {
            return nil,fmt.Errorf("failed parsing the pubDate element %s:%s",rssItem.PubDate,err)
        }
        dmlog.Debug("pubDate:", pubDate)   

        duration,err := parseDuration( rssItem.Duration )
        if err!=nil {
            return nil,fmt.Errorf("failed parsing the duration %s:%s",rssItem.Duration,err)
        }

        podEpisode := PodcastEpisode {Title: rssItem.Title,
                                      Summary: rssItem.Summary,
                                      AudioFileUrl: Url(enclosure.Url),
                                      EpisodeNumber: rssItem.EpisodeNumber,
                                      PublicationDate: DateTime(pubDate), 
                                      AudioFileSize: ByteSize(enclosure.Length),
                                      AudioDuration: duration,
        }
        dmlog.Debug("podEpisode.PublicationDate:", &podEpisode.PublicationDate)   
        if funcAcceptEpisode( &podSource, &podEpisode) {
            podEpisodes= append( podEpisodes, podEpisode)
		}
       
	}
    return podEpisodes,nil
}

//-------------------------------------------------------------------------------------------------
/* According to:
   https://rfwilmut.net/notes/itunes/tag/duration.html
   the content of the itunes:duration element is expressed either in seconds, 
   or as HH:MM:SS the previous recommendation by Apple)
 */
func parseDuration( duration string ) (time.Duration, error) {
  if strings.Contains( duration, durationSeparator) {
      hhMmSsFields := strings.Split( duration, durationSeparator)
      if 3!=len(hhMmSsFields){
          return time.Duration(0), fmt.Errorf("invalid duration, format should be hh:mm:ss")
	  } else {
          return time.ParseDuration( hhMmSsFields[0]+"h"+hhMmSsFields[1]+"m"+hhMmSsFields[2]+"s")
	  }
  } else {
      var secsDuration = duration+"s"
      return time.ParseDuration( secsDuration)
  }  
}