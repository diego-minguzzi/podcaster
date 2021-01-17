package podcaster

import "github.com/diego-minguzzi/dmlog"
import "log"
import "os"
import "testing"

func TestMain(m *testing.M) {
	// Releases all log resources when the program terminates.
	defer dmlog.Terminate()

	// Adds a sink that writes to the console.
	sinkId, err := dmlog.AddConsoleSink(dmlog.DebugSeverity)
	if err != nil {
		log.Panicln("AddConsoleSink() failed:", err)
	}
	// Sets a custom format for the console sink.
	if !dmlog.SetSinkOutputFormat(sinkId, dmlog.LogMessageType,
		dmlog.FilenameLineFmt, dmlog.LineEndFmt, dmlog.SeverityFmt, dmlog.TextFmt, dmlog.LineEndFmt) {
		log.Panicln("SetSinkOutputFormat() failed.")
	}

	dmlog.Debug("About to run all tests")
	exitCode := m.Run()
	dmlog.Debug("Tests terminated")
	os.Exit(exitCode)
}

func TestUpdateFeeds(t *testing.T) {
	const numConnections int = 1

	podSources := []PodcastSource{
		PodcastSource{"Freakonomics", "http://feeds.feedburner.com/freakonomicsradio", 4},
		PodcastSource{"Hensel_Minutes", "http://feeds.feedburner.com/Hanselminutes?format=xml", 4},
	}

	feedUpdater := UpdateFeeds(podSources, numConnections)
	isSuccess := true

	updateReceiver := func(chan PodcastFeedUpdateChItem) {
		isSuccess = true
	}
	go updateReceiver(feedUpdater.PodcastFeedItemsChan())
	feedUpdater.Wait()
	if !isSuccess {
	}
}
