package goradbot

import (
	"io"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/jonas747/dca"
)

var waitTime = 250 * time.Millisecond
var vc *discordgo.VoiceConnection
var isPlaying = false

// Change these accordingly
func playurl(s *discordgo.Session, serverID, channelID, url string) (err error) {

	//Setting a new DCA session
	options := dca.StdEncodeOptions
	options.RawOutput = true
	options.Bitrate = 128
	options.Application = "lowdelay"

	// Join the provided voice channel.
	vc, err := s.ChannelVoiceJoin(serverID, channelID, false, true)
	if err != nil {
		return err
	}

	// Sleep for a specified amount of time before playing the sound
	time.Sleep(waitTime)

	// Start speaking.
	vc.Speaking(true)

	//Connect to the server (a.k.a: do the magic trick)
	encodeSession, err := dca.EncodeFile(url, options)
	if err != nil {
		return
	}

	//And now, stream!
	done := make(chan error)
	dca.NewStream(encodeSession, vc, done)
	err = <-done
	if err != nil && err != io.EOF {
		return
	}

	return nil
}

func stop(s *discordgo.Session, serverID, channelID string) (err error) {
	// Join the provided voice channel.
	vc, err := s.ChannelVoiceJoin(serverID, channelID, false, false)

	if err != nil {
		return
	}

	// Stop speaking
	vc.Speaking(false)

	// Sleep for a specificed amount of time before ending.
	time.Sleep(waitTime)

	// Disconnect from the provided voice channel.
	vc.Disconnect()

	return
}
