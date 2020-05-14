package goradbot

import (
	"fmt"
	"strings"
	"net/url"
	
	"github.com/PAWA-cl/goradbot/embed"
	"github.com/bwmarrin/discordgo"
)

func checkValidCommand(s *discordgo.Session, m *discordgo.MessageCreate) (e error) {
	// But first: Validations
	
	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}
	
	// Find the channel that the message came from.
	c, e := s.State.Channel(m.ChannelID)
	
	if e != nil {
		fmt.Println("Error verifying channel:", e)
		return
	}
	
	//Ignore the message if it comes from a DM
	if c.Type == discordgo.ChannelTypeDM {
		return
	}
	
	// Find the guild for that channel.
	g, e := s.State.Guild(c.GuildID)
	
	if e != nil {
		fmt.Println("Error verifying guild/server:", e)
		return
	}
	
	// It's Showtime!
	var rm RegexMap
	
	//Rare case where only the command is written.
	if strings.ToLower(m.Content) == strings.ToLower(prefix) {
		//Show available commands
		msg := fmt.Sprintf("Hello my little friend! Do you need some help? Try with **%s help**", prefix)
		s.ChannelMessageSend(m.ChannelID, msg)
	}
	
	//All commands are in form of (prefix) (command) (arguments)
	if rm.Matches(fmt.Sprintf(`^%s( ?(?P<Command>\w+))( (?P<Argument>.*))?$`, prefix), m.Content) {
		
		//Obtains the commands. Easy, right?
		command := rm.MappedResult["Command"]
		
		switch command {
		case "help":
			//Get the embed
			msg, e := embed.GetEmbedMessage("help")
			
			if e != nil {
				fmt.Println("Error sending the help message:", e)
				return e
			}
			
			//Replace fields with the command
			for _, f := range msg.Fields {
				f.Name = strings.ReplaceAll(f.Name, "%prefix", prefix)
				f.Value = strings.ReplaceAll(f.Value, "%prefix", prefix)
			}
			
			//Show available commands
			s.ChannelMessageSendEmbed(m.ChannelID, msg)
		case "play":
			// Look for the message sender in that guild's current voice states.
			check := false
			for _, vs := range g.VoiceStates {
				if vs.UserID == m.Author.ID {
					//Ok, you are connected.
					check = true
					
					//Check if the argument is a valid url
					uri := rm.MappedResult["Argument"]
					_, err := url.ParseRequestURI(uri)
					if len(uri)== 0 || err != nil {
						msg := "I'm sorry Dave, your URL is malformed. "
						msg += "Please input a valid URL and try again."
						s.ChannelMessageSend(m.ChannelID, msg)
						return
					}
					
					//Check if is a valid URL format
					if(!strings.HasSuffix(uri, ".mp3")){
						msg := "Sorry, only mp3 for now."
						s.ChannelMessageSend(m.ChannelID, msg)
						return
					}
					
					//C'mon DJ, play that song.
					/*e = playurl(s, g.ID, vs.ChannelID, uri)
					if e != nil {
						fmt.Println("Error playing sound:", e)
					}*/

					go playurl(s, g.ID, vs.ChannelID, uri)
					
					msg, e := embed.GetEmbedMessage("play")
					
					if e != nil {
						fmt.Println("Error sending the play message:", e)
						return e
					}
					
					//Replace fields with the server status
					msg.Description = strings.ReplaceAll(msg.Description, "%server", uri)

					//Show Play
					s.ChannelMessageSendEmbed(m.ChannelID, msg)
				}
			}
			
			//If the user is not connected, notify it.
			if !check {
				msg := "I'm sorry, but i cannot join automatically without you. "
				msg += "Try connecting to a voice channel first."
				s.ChannelMessageSend(m.ChannelID, msg)
			}
		case "stop":
			//Stop Playing
			for _, vs := range g.VoiceStates {
				if vs.UserID == m.Author.ID {
					stop(s, g.ID, vs.ChannelID)
				}
			}
			//Get the embed
			msg, e := embed.GetEmbedMessage("stop")
			
			if e != nil {
				fmt.Println("Error sending the stop message:", e)
				return e
			}
			
			//Show stop meme
			s.ChannelMessageSendEmbed(m.ChannelID, msg)
		}
	}
	
	return
}
