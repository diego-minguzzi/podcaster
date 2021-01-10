package podcaster

import "strings"
import "testing"
import "time"

func TestFreakonomicsParseRssFeedAllEpisodes( t *testing.T) {
    podSource := PodcastSource {
        PodcastName: "FreakonomicsRadio",
        FeedUrl: "http://feeds.feedburner.com/freakonomicsradio",
        NumEpisodesToDownload: 10,
        OverwriteFilename: false, 
	}

    rssReader := strings.NewReader( freakRssUpdate)
  
    podEpisodes,err := ParseRssFeedAllEpisodes( podSource, rssReader)
    if err!=nil {
        t.Error(t.Name(),`ParseRssFeedAllEpisodes() failed:`,err)            
        return 
    }

	var indxEpisode = 0
    var gotEpisode PodcastEpisode
    if indxEpisode>=len(podEpisodes) {
        t.Error(t.Name(),`indxEpisode:`,indxEpisode,` num episodes:`,len(podEpisodes))            
	} else {
        duration, err:= time.ParseDuration("2596s")
        if err!=nil {
            t.Error(t.Name(),"ParseDuration failed",err)
            return
		}
		gotEpisode= podEpisodes[indxEpisode]
	    wantEpisode := PodcastEpisode {
            Title: "A Sneak Peek at Biden’s Top Economist",
            Summary: `The incoming president argues that the economy and the environment are deeply connected. This is reflected in his choice for National Economic Council director — Brian Deese, a climate-policy wonk and veteran of the no-drama-Obama era. But don’t mistake Deese’s lack of drama for a lack of intensity.`,
            AudioFileUrl: `https://chtbl.com/track/288D49/traffic.omny.fm/d/clips/aaea4e69-af51-495e-afc9-a9760146922b/14a43378-edb2-49be-8511-ab0d000a7030/32f622dd-6044-4f9a-a6d6-ac8d0008b84e/audio.mp3?utm_source=Podcast&in_playlist=d1b9612f-bb1b-4b85-9c0c-ab0d004ab37a`,
            AudioFileSize: ByteSize(62328096),
            PublicationDate: DateTime( time.Date(2020, time.December, 12, 4, 0, 0, 0, time.UTC) ),
            EpisodeNumber: 443,
            AudioDuration: duration,
        }
        if gotEpisode.IsEqualForTestTo( &wantEpisode) { 
            t.Error(t.Name(),"Failed in index:",indxEpisode,
                    "\ngot:",gotEpisode,"\nwant:",wantEpisode)
        }
    } 

    indxEpisode++

    if indxEpisode>=len(podEpisodes) {
        t.Error(t.Name(),`indxEpisode:`,indxEpisode,` num episodes:`,len(podEpisodes))            
	} else {
        duration, err:= time.ParseDuration("3422s")
        if err!=nil {
            t.Error(t.Name(),"ParseDuration failed",err)
            return
		}
		gotEpisode= podEpisodes[indxEpisode]
	    wantEpisode := PodcastEpisode {
            Title: "PLAYBACK (2015): Could the Next Brooklyn Be ... Las Vegas?!",
            Summary: `Tony Hsieh, the longtime C.E.O. of Zappos, was an iconoclast and a dreamer. Five years ago, we sat down with him around a desert campfire to talk about those dreams. Hsieh died recently from injuries sustained in a house fire; he was 46.`,
            AudioFileUrl: `https://chtbl.com/track/288D49/traffic.omny.fm/d/clips/aaea4e69-af51-495e-afc9-a9760146922b/14a43378-edb2-49be-8511-ab0d000a7030/4df1ab4b-de1f-4b50-aced-ac880003b5a5/audio.mp3?utm_source=Podcast&;in_playlist=d1b9612f-bb1b-4b85-9c0c-ab0d004ab37a`,
            AudioFileSize: ByteSize(62328096),
            PublicationDate: DateTime( time.Date(2020, time.December, 06, 10, 0, 0, 0, time.UTC) ),
            EpisodeNumber: 0,
            AudioDuration: duration,
        }
        t.Log(t.Name(),"got Title:",gotEpisode.Title)
        t.Log(t.Name(),"got PublicationDate:",&gotEpisode.PublicationDate)
        t.Log(t.Name(),"got AudioDuration:",gotEpisode.AudioDuration)
        if gotEpisode.IsEqualForTestTo( &wantEpisode) { 
            t.Error(t.Name(),"Failed in index:",indxEpisode,
                    "\ngot:",gotEpisode,"\nwant:",wantEpisode)
        }
    } 

    indxEpisode++

    if indxEpisode>=len(podEpisodes) {
        t.Error(t.Name(),`indxEpisode:`,indxEpisode,` num episodes:`,len(podEpisodes))            
	} else {
        duration, err:= time.ParseDuration("2557s")
        if err!=nil {
            t.Error(t.Name(),"ParseDuration failed",err)
            return
		}
		gotEpisode= podEpisodes[indxEpisode]
	    wantEpisode := PodcastEpisode {
            Title: "How to Make Meetings Less Terrible (Ep. 389 Rebroadcast)",
            Summary: `In the U.S. alone, we hold 55 million meetings a day. Most of them are woefully unproductive, and tyrannize our offices. The revolution begins now — with better agendas, smaller invite lists, and an embrace of healthy conflict.`,
            AudioFileUrl: `https://chtbl.com/track/288D49/traffic.omny.fm/d/clips/aaea4e69-af51-495e-afc9-a9760146922b/14a43378-edb2-49be-8511-ab0d000a7030/461751b7-0335-4534-ad2f-abc80154e54c/audio.mp3?utm_source=Podcast&in_playlist=d1b9612f-bb1b-4b85-9c0c-ab0d004ab37a`,
            AudioFileSize: ByteSize(40935499),
            PublicationDate: DateTime( time.Date(2020, time.May, 28, 3, 0, 0, 0, time.UTC) ),
            EpisodeNumber: 0,
            AudioDuration: duration,
        }
        t.Log(t.Name(),"got Title:",gotEpisode.Title)
        t.Log(t.Name(),"got PublicationDate:",&gotEpisode.PublicationDate)
        t.Log(t.Name(),"got AudioDuration:",gotEpisode.AudioDuration)
        if gotEpisode.IsEqualForTestTo( &wantEpisode) { 
            t.Error(t.Name(),"Failed in index:",indxEpisode,
                    "\ngot:",gotEpisode,"\nwant:",wantEpisode)
        }
    } 

}


const freakRssUpdate string = `
<?xml version="1.0" encoding="utf-8"?>
<rss xmlns:itunes="http://www.itunes.com/dtds/podcast-1.0.dtd" xmlns:atom="http://www.w3.org/2005/Atom" xmlns:media="http://search.yahoo.com/mrss/" xmlns:psc="https://podlove.org/simple-chapters/" xmlns:omny="https://omny.fm/rss-extensions" xmlns:content="http://purl.org/rss/1.0/modules/content/" xmlns:googleplay="http://www.google.com/schemas/play-podcasts/1.0" xmlns:acast="https://schema.acast.com/1.0/" version="2.0">
  <channel>
    <language>en-US</language>
    <atom:link rel="self" type="application/rss+xml" href="https://www.omnycontent.com/d/playlist/aaea4e69-af51-495e-afc9-a9760146922b/14a43378-edb2-49be-8511-ab0d000a7030/d1b9612f-bb1b-4b85-9c0c-ab0d004ab37a/podcast.rss" />
    <atom:link rel="first" type="application/rss+xml" href="https://www.omnycontent.com/d/playlist/aaea4e69-af51-495e-afc9-a9760146922b/14a43378-edb2-49be-8511-ab0d000a7030/d1b9612f-bb1b-4b85-9c0c-ab0d004ab37a/podcast.rss?page=1" />
    <atom:link rel="last" type="application/rss+xml" href="https://www.omnycontent.com/d/playlist/aaea4e69-af51-495e-afc9-a9760146922b/14a43378-edb2-49be-8511-ab0d000a7030/d1b9612f-bb1b-4b85-9c0c-ab0d004ab37a/podcast.rss?page=1" />
    <title>Freakonomics Radio</title>
    <link>http://freakonomics.com/</link>
    <description><![CDATA[<p>Discover the hidden side of everything with Stephen J. Dubner, co-author of the&nbsp;<em>Freakonomics</em>&nbsp;books. Each week,&nbsp;<em>Freakonomics Radio</em>&nbsp;tells you things you always thought you knew (but didn&rsquo;t) and things you never thought you wanted to know (but do) &mdash;&nbsp;from the economics of sleep to how to become great at just about anything. Dubner speaks with Nobel laureates and provocateurs, intellectuals and entrepreneurs, and various other underachievers. The entire archive, going back to 2010, is available on the <a href="https://www.stitcher.com/download">Stitcher podcast app</a> and at&nbsp;<a href="http://freakonomics.com/" data-saferedirecturl="https://www.google.com/url?q=http://freakonomics.com&amp;source=gmail&amp;ust=1590509555837000&amp;usg=AFQjCNFHpG0N-_WcZE2ZMsOAPtNFtbJ6bw">freakonomics.com</a>.</p>]]></description>
    <itunes:type>episodic</itunes:type>
    <itunes:summary>Discover the hidden side of everything with Stephen J. Dubner, co-author of the Freakonomics books. Each week, Freakonomics Radio tells you things you always thought you knew (but didn’t) and things you never thought you wanted to know (but do) — from the economics of sleep to how to become great at just about anything. Dubner speaks with Nobel laureates and provocateurs, intellectuals and entrepreneurs, and various other underachievers. The entire archive, going back to 2010, is available on the Stitcher podcast app and at freakonomics.com.</itunes:summary>
    <itunes:owner>
      <itunes:name>Freakonomics Radio + Stitcher</itunes:name>
      <itunes:email>editor@freakonomics.com</itunes:email>
    </itunes:owner>
    <itunes:author>Freakonomics Radio + Stitcher</itunes:author>
    <copyright>2020 ​Dubner Productions and Stitcher</copyright>
    <itunes:explicit>clean</itunes:explicit>
    <itunes:category text="Society &amp; Culture">
      <itunes:category text="Documentary" />
    </itunes:category>
    <itunes:image href="https://www.omnycontent.com/d/playlist/aaea4e69-af51-495e-afc9-a9760146922b/14a43378-edb2-49be-8511-ab0d000a7030/d1b9612f-bb1b-4b85-9c0c-ab0d004ab37a/image.jpg?t=1589407970&amp;size=Large" />
    <image>
      <url>https://www.omnycontent.com/d/playlist/aaea4e69-af51-495e-afc9-a9760146922b/14a43378-edb2-49be-8511-ab0d000a7030/d1b9612f-bb1b-4b85-9c0c-ab0d004ab37a/image.jpg?t=1589407970&amp;size=Large</url>
      <title>Freakonomics Radio</title>
      <link>http://freakonomics.com/</link>
    </image>
    <item>
      <title>443. A Sneak Peek at Biden’s Top Economist</title>
      <itunes:title>A Sneak Peek at Biden’s Top Economist</itunes:title>
      <description>The incoming president argues that the economy and the environment are deeply connected. This is reflected in his choice for National Economic Council director — Brian Deese, a climate-policy wonk and veteran of the no-drama-Obama era. But don’t mistake Deese’s lack of drama for a lack of intensity.</description>
      <content:encoded><![CDATA[<p>The incoming president argues that the economy and the environment are deeply connected. This is reflected in his choice for National Economic Council director &mdash; Brian Deese, a climate-policy wonk and veteran of the no-drama-Obama era. But don&rsquo;t mistake Deese&rsquo;s lack of drama for a lack of intensity.</p>]]></content:encoded>
      <itunes:summary>The incoming president argues that the economy and the environment are deeply connected. This is reflected in his choice for National Economic Council director — Brian Deese, a climate-policy wonk and veteran of the no-drama-Obama era. But don’t mistake Deese’s lack of drama for a lack of intensity.</itunes:summary>
      <itunes:episodeType>full</itunes:episodeType>
      <itunes:episode>443</itunes:episode>
      <itunes:author>Freakonomics Radio + Stitcher</itunes:author>
      <itunes:image href="https://www.omnycontent.com/d/playlist/aaea4e69-af51-495e-afc9-a9760146922b/14a43378-edb2-49be-8511-ab0d000a7030/d1b9612f-bb1b-4b85-9c0c-ab0d004ab37a/image.jpg?t=1589407970&amp;size=Large" />
      <media:content url="https://chtbl.com/track/288D49/traffic.omny.fm/d/clips/aaea4e69-af51-495e-afc9-a9760146922b/14a43378-edb2-49be-8511-ab0d000a7030/32f622dd-6044-4f9a-a6d6-ac8d0008b84e/audio.mp3?utm_source=Podcast&amp;in_playlist=d1b9612f-bb1b-4b85-9c0c-ab0d004ab37a" type="audio/mpeg">
        <media:player url="https://omny.fm/shows/freakonomics-radio/a-sneak-peek-at-biden-s-top-economist/embed" />
      </media:content>
      <media:content url="https://www.omnycontent.com/d/playlist/aaea4e69-af51-495e-afc9-a9760146922b/14a43378-edb2-49be-8511-ab0d000a7030/d1b9612f-bb1b-4b85-9c0c-ab0d004ab37a/image.jpg?t=1589407970&amp;size=Large" type="image/jpeg" />
      <guid isPermaLink="false">32f622dd-6044-4f9a-a6d6-ac8d0008b84e</guid>
      <omny:clipId>32f622dd-6044-4f9a-a6d6-ac8d0008b84e</omny:clipId>
      <pubDate>Thu, 10 Dec 2020 04:00:00 +0000</pubDate>
      <itunes:duration>2596</itunes:duration>
      <enclosure url="https://chtbl.com/track/288D49/traffic.omny.fm/d/clips/aaea4e69-af51-495e-afc9-a9760146922b/14a43378-edb2-49be-8511-ab0d000a7030/32f622dd-6044-4f9a-a6d6-ac8d0008b84e/audio.mp3?utm_source=Podcast&amp;in_playlist=d1b9612f-bb1b-4b85-9c0c-ab0d004ab37a" length="62328096" type="audio/mpeg" />
      <link>https://omny.fm/shows/freakonomics-radio/a-sneak-peek-at-biden-s-top-economist</link>
    </item>
    <item>
      <title>PLAYBACK (2015): Could the Next Brooklyn Be ... Las Vegas?!</title>
      <itunes:title>PLAYBACK (2015): Could the Next Brooklyn Be ... Las Vegas?!</itunes:title>
      <description>Tony Hsieh, the longtime C.E.O. of Zappos, was an iconoclast and a dreamer. Five years ago, we sat down with him around a desert campfire to talk about those dreams. Hsieh died recently from injuries sustained in a house fire; he was 46.</description>
      <content:encoded><![CDATA[<p>Tony Hsieh, the longtime C.E.O. of Zappos, was an iconoclast and a dreamer. Five years ago, we sat down with him around a desert campfire to talk about those dreams. Hsieh died recently from injuries sustained in a house fire; he was 46.</p>]]></content:encoded>
      <itunes:summary>Tony Hsieh, the longtime C.E.O. of Zappos, was an iconoclast and a dreamer. Five years ago, we sat down with him around a desert campfire to talk about those dreams. Hsieh died recently from injuries sustained in a house fire; he was 46.</itunes:summary>
      <itunes:episodeType>full</itunes:episodeType>
      <itunes:author>Freakonomics Radio + Stitcher</itunes:author>
      <itunes:image href="https://www.omnycontent.com/d/playlist/aaea4e69-af51-495e-afc9-a9760146922b/14a43378-edb2-49be-8511-ab0d000a7030/d1b9612f-bb1b-4b85-9c0c-ab0d004ab37a/image.jpg?t=1589407970&amp;size=Large" />
      <media:content url="https://chtbl.com/track/288D49/traffic.omny.fm/d/clips/aaea4e69-af51-495e-afc9-a9760146922b/14a43378-edb2-49be-8511-ab0d000a7030/4df1ab4b-de1f-4b50-aced-ac880003b5a5/audio.mp3?utm_source=Podcast&amp;in_playlist=d1b9612f-bb1b-4b85-9c0c-ab0d004ab37a" type="audio/mpeg">
        <media:player url="https://omny.fm/shows/freakonomics-radio/playback-2015-could-the-next-brooklyn-be-las-vegas/embed" />
      </media:content>
      <media:content url="https://www.omnycontent.com/d/playlist/aaea4e69-af51-495e-afc9-a9760146922b/14a43378-edb2-49be-8511-ab0d000a7030/d1b9612f-bb1b-4b85-9c0c-ab0d004ab37a/image.jpg?t=1589407970&amp;size=Large" type="image/jpeg" />
      <guid isPermaLink="false">4df1ab4b-de1f-4b50-aced-ac880003b5a5</guid>
      <omny:clipId>4df1ab4b-de1f-4b50-aced-ac880003b5a5</omny:clipId>
      <pubDate>Sun, 06 Dec 2020 10:00:00 +0000</pubDate>
      <itunes:duration>3422</itunes:duration>
      <enclosure url="https://chtbl.com/track/288D49/traffic.omny.fm/d/clips/aaea4e69-af51-495e-afc9-a9760146922b/14a43378-edb2-49be-8511-ab0d000a7030/4df1ab4b-de1f-4b50-aced-ac880003b5a5/audio.mp3?utm_source=Podcast&amp;in_playlist=d1b9612f-bb1b-4b85-9c0c-ab0d004ab37a" length="82151920" type="audio/mpeg" />
      <link>https://omny.fm/shows/freakonomics-radio/playback-2015-could-the-next-brooklyn-be-las-vegas</link>
    </item>
    <item>
      <title>How to Make Meetings Less Terrible (Ep. 389 Rebroadcast)</title>
      <itunes:title>How to Make Meetings Less Terrible (Ep. 389 Rebroadcast)</itunes:title>
      <description>In the U.S. alone, we hold 55 million meetings a day. Most of them are woefully unproductive, and tyrannize our offices. The revolution begins now — with better agendas, smaller invite lists, and an embrace of healthy conflict.</description>
      <content:encoded><![CDATA[<p>In the U.S. alone, we hold 55 million meetings a day. Most of them are woefully unproductive, and tyrannize our offices. The revolution begins now &mdash; with better agendas, smaller invite lists, and an embrace of healthy conflict.</p>]]></content:encoded>
      <itunes:summary>In the U.S. alone, we hold 55 million meetings a day. Most of them are woefully unproductive, and tyrannize our offices. The revolution begins now — with better agendas, smaller invite lists, and an embrace of healthy conflict.</itunes:summary>
      <itunes:episodeType>full</itunes:episodeType>
      <itunes:author>Freakonomics Radio + Stitcher</itunes:author>
      <itunes:image href="https://www.omnycontent.com/d/playlist/aaea4e69-af51-495e-afc9-a9760146922b/14a43378-edb2-49be-8511-ab0d000a7030/d1b9612f-bb1b-4b85-9c0c-ab0d004ab37a/image.jpg?t=1589407970&amp;size=Large" />
      <media:content url="https://chtbl.com/track/288D49/traffic.omny.fm/d/clips/aaea4e69-af51-495e-afc9-a9760146922b/14a43378-edb2-49be-8511-ab0d000a7030/461751b7-0335-4534-ad2f-abc80154e54c/audio.mp3?utm_source=Podcast&amp;in_playlist=d1b9612f-bb1b-4b85-9c0c-ab0d004ab37a" type="audio/mpeg">
        <media:player url="https://omny.fm/shows/freakonomics-radio/how-to-make-meetings-less-terrible-ep-389-rebroadc/embed" />
      </media:content>
      <media:content url="https://www.omnycontent.com/d/playlist/aaea4e69-af51-495e-afc9-a9760146922b/14a43378-edb2-49be-8511-ab0d000a7030/d1b9612f-bb1b-4b85-9c0c-ab0d004ab37a/image.jpg?t=1589407970&amp;size=Large" type="image/jpeg" />
      <guid isPermaLink="false">461751b7-0335-4534-ad2f-abc80154e54c</guid>
      <omny:clipId>461751b7-0335-4534-ad2f-abc80154e54c</omny:clipId>
      <pubDate>Thu, 28 May 2020 03:00:00 +0000</pubDate>
      <itunes:duration>2557</itunes:duration>
      <enclosure url="https://chtbl.com/track/288D49/traffic.omny.fm/d/clips/aaea4e69-af51-495e-afc9-a9760146922b/14a43378-edb2-49be-8511-ab0d000a7030/461751b7-0335-4534-ad2f-abc80154e54c/audio.mp3?utm_source=Podcast&amp;in_playlist=d1b9612f-bb1b-4b85-9c0c-ab0d004ab37a" length="40935499" type="audio/mpeg" />
      <link>https://omny.fm/shows/freakonomics-radio/how-to-make-meetings-less-terrible-ep-389-rebroadc</link>
      <omny:stitcherId>69981293</omny:stitcherId>
    </item>
  </channel>
</rss>
`