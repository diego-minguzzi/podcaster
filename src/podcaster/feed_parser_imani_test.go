package podcaster

import "strings"
import "testing"

func TestImaniParseRssFeedAllEpisodes( t *testing.T) {
    podSource := PodcastSource {
        PodcastName: "Imani_State_of_Mind",
        FeedUrl: "https://omny.fm/shows/imani-state-of-mind/playlists/podcast.rss",
        NumEpisodesToDownload: 10,        
	}

    rssReader := strings.NewReader( imaniRssContent)
  
    podEpisodes,err := ParseRssFeedAllEpisodes( podSource, rssReader)
    if err!=nil {
        t.Error(t.Name(),`ParseRssFeedAllEpisodes() failed:`,err)            
        return 
    }
    {
        want:=1
        got:=len(podEpisodes)
        if want!=got {
            t.Error(t.Name(),`Num episodes, want:`,want,`got:`,got)            
            return 
        }
	}

    episode := podEpisodes[0]

    t.Log("Parse result", episode.Title, episode.AudioFileSize, episode.AudioFileUrl)
    t.Log("EpisodeNumber:", episode.EpisodeNumber)
    t.Log("AudioDuration:", episode.AudioDuration)
    {
        got:= episode.AudioFileUrl
        want:=Url(`https://chtbl.com/track/288D49/traffic.omny.fm/d/clips/aaea4e69/audio.mp3?utm_source=Podcast`)        
        if want!=got { t.Error(t.Name(),`AudioFileUrl, want:`,want,`got:`,got) }
    }
    { 
        got:= episode.AudioFileSize
        want:=ByteSize(99871330)        
        if want!=got { t.Error(t.Name(),`AudioFileSize, want:`,want,`got:`,got) }
    }
    {
        got:= episode.Title
        want:=`Schizo-what??`        
        if want!=got { t.Error(t.Name(),`Title, want:`,want,`got:`,got) }
    }
    
}


const imaniRssContent string = `
<?xml version="1.0" encoding="utf-8"?>
<rss xmlns:itunes="http://www.itunes.com/dtds/podcast-1.0.dtd" xmlns:atom="http://www.w3.org/2005/Atom" xmlns:media="http://search.yahoo.com/mrss/" xmlns:psc="https://podlove.org/simple-chapters/" xmlns:omny="https://omny.fm/rss-extensions" xmlns:content="http://purl.org/rss/1.0/modules/content/" xmlns:googleplay="http://www.google.com/schemas/play-podcasts/1.0" xmlns:acast="https://schema.acast.com/1.0/" version="2.0">
  <channel>
    <language>en-US</language>
    <atom:link rel="self" type="application/rss+xml" href="https://www.omnycontent.com/d/playlist/aaea4e69-af51-495e-afc9-a9760146922b/151a9389-3003-4d91-b575-ab99015606e2/e4ead0f3-6fe6-4c73-803c-ab99015720ef/podcast.rss" />
    <atom:link rel="first" type="application/rss+xml" href="https://www.omnycontent.com/d/playlist/aaea4e69-af51-495e-afc9-a9760146922b/151a9389-3003-4d91-b575-ab99015606e2/e4ead0f3-6fe6-4c73-803c-ab99015720ef/podcast.rss?page=1" />
    <atom:link rel="last" type="application/rss+xml" href="https://www.omnycontent.com/d/playlist/aaea4e69-af51-495e-afc9-a9760146922b/151a9389-3003-4d91-b575-ab99015606e2/e4ead0f3-6fe6-4c73-803c-ab99015720ef/podcast.rss?page=1" />
    <title>Imani State of Mind</title>
    <link>https://omny.fm/shows/imani-state-of-mind/playlists/podcast</link>
    <description><![CDATA[<p>Get into an Imani State Of Mind with psychiatrist Dr. Imani Walker. Every week, she&rsquo;ll break down mental illness and mental health through pop culture and news and in the process normalize what getting your mind right really looks like.</p>]]></description>
    <itunes:type>episodic</itunes:type>
    <itunes:summary>Get into an Imani State Of Mind with psychiatrist Dr. Imani Walker. Every week, she’ll break down mental illness and mental health through pop culture and news and in the process normalize what getting your mind right really looks like.</itunes:summary>
    <itunes:owner>
      <itunes:name>Stitcher</itunes:name>
      <itunes:email>originals@stitcher.com</itunes:email>
    </itunes:owner>
    <itunes:author>Stitcher &amp; Imani Walker</itunes:author>
    <itunes:explicit>yes</itunes:explicit>
    <itunes:category text="Health &amp; Fitness">
      <itunes:category text="Mental Health" />
    </itunes:category>
    <itunes:image href="https://www.omnycontent.com/d/programs/aaea4e69-af51-495e-afc9-a9760146922b/151a9389-3003-4d91-b575-ab99015606e2/image.jpg?t=1601931831&amp;size=Large" />
    <image>
      <url>https://www.omnycontent.com/d/programs/aaea4e69-af51-495e-afc9-a9760146922b/151a9389-3003-4d91-b575-ab99015606e2/image.jpg?t=1601931831&amp;size=Large</url>
      <title>Imani State of Mind</title>
      <link>https://omny.fm/shows/imani-state-of-mind/playlists/podcast</link>
    </image>
    <item>
      <title>Schizo-what??</title>
      <itunes:title>Schizo-what??</itunes:title>
      <description>Time for a study break! Schizophrenia is rare, but it’s the mental illness that shows up the most in pop culture. Imani explains what schizophrenia is - and isn't! - and debunks the most famous misconceptions about it. And on Pop Culture Diagnosis, Imani breaks down how The CW's Crazy Ex-Girlfriend portrays borderline personality disorder. 
Keep up with Dr. Imani on Instagram @doctor.imani and on Twitter @doctor_imani. Send your mental health questions to askdrimani@gmail.com or leave a voicemail at (424) 235-0064‬.

See omnystudio.com/listener for privacy information.</description>
      <content:encoded><![CDATA[<p>Time for a study break! Schizophrenia is rare, but it&rsquo;s the mental illness that shows up the most in pop culture. Imani explains what schizophrenia is - and isn't! - and debunks the most famous misconceptions about it. And on Pop Culture Diagnosis, Imani breaks down how The CW's Crazy Ex-Girlfriend portrays borderline personality disorder. <br>Keep up with Dr. Imani on Instagram @doctor.imani and on Twitter @doctor_imani. Send your mental health questions to <a href="mailto:askdrimani@gmail.com">askdrimani@gmail.com</a> or leave a voicemail at (424) 235-0064‬.</p><p>See <a href="https://omnystudio.com/listener">omnystudio.com/listener</a> for privacy information.</p>]]></content:encoded>
      <itunes:summary>Time for a study break! Schizophrenia is rare, but it’s the mental illness that shows up the most in pop culture. Imani explains what schizophrenia is - and isn't! - and debunks the most famous misconceptions about it. And on Pop Culture Diagnosis, Imani breaks down how The CW's Crazy Ex-Girlfriend portrays borderline personality disorder. 
Keep up with Dr. Imani on Instagram @doctor.imani and on Twitter @doctor_imani. Send your mental health questions to askdrimani@gmail.com or leave a voicemail at (424) 235-0064‬.

See omnystudio.com/listener for privacy information.</itunes:summary>
      <itunes:episodeType>full</itunes:episodeType>
      <itunes:episode>21</itunes:episode>
      <itunes:author>Stitcher &amp; Imani Walker</itunes:author>
      <itunes:image href="https://www.omnycontent.com/d/programs/aaea4e69-af51-495e-afc9-a9760146922b/151a9389-3003-4d91-b575-ab99015606e2/image.jpg?t=1601931831&amp;size=Large" />
      <media:content url="https://chtbl.com/track/288D49/traffic.omny.fm/d/clips/aaea4e69-af51-495e-afc9-a9760146922b/151a9389-3003-4d91-b575-ab99015606e2/42a12219-7c35-4a7a-97bf-ac8c0173b884/audio.mp3?utm_source=Podcast&amp;in_playlist=e4ead0f3-6fe6-4c73-803c-ab99015720ef" type="audio/mpeg">
        <media:player url="https://omny.fm/shows/imani-state-of-mind/schizo-what/embed" />
      </media:content>
      <media:content url="https://www.omnycontent.com/d/programs/aaea4e69-af51-495e-afc9-a9760146922b/151a9389-3003-4d91-b575-ab99015606e2/image.jpg?t=1601931831&amp;size=Large" type="image/jpeg" />
      <guid isPermaLink="false">42a12219-7c35-4a7a-97bf-ac8c0173b884</guid>
      <omny:clipId>42a12219-7c35-4a7a-97bf-ac8c0173b884</omny:clipId>
      <pubDate>Thu, 10 Dec 2020 05:00:00 +0000</pubDate>
      <itunes:duration>4158</itunes:duration>
      <enclosure url="https://chtbl.com/track/288D49/traffic.omny.fm/d/clips/aaea4e69/audio.mp3?utm_source=Podcast" length="99871330" type="audio/mpeg" />
      <link>https://omny.fm/shows/imani-state-of-mind/schizo-what</link>
    </item>
  </channel>
</rss>
`