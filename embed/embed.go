package embed

import (
	"fmt"
	"io/ioutil"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/ghodss/yaml"
)

var commandLocation = "commands"

//GetEmbedMessage returns the MessageEmbed object
func GetEmbedMessage(command string) (m *discordgo.MessageEmbed, e error) {
	//Read the embed file
	m, e = readEmbedFile(command)
	if e != nil {
		return
	}

	//Finally, set the timestamp
	m.Timestamp = time.Now().Format(time.RFC3339)

	return
}

func readEmbedFile(command string) (y *discordgo.MessageEmbed, e error) {
	//Read the file
	location := fmt.Sprintf("%s/%s.yaml", commandLocation, command)

	dat, e := ioutil.ReadFile(location)
	if e != nil {
		fmt.Printf("[%s] There was an error reading the embed command: %v\n", location, e)
		return
	}

	//Return the Yamelette, Kudasai!
	e = yaml.Unmarshal([]byte(dat), &y)
	if e != nil {
		fmt.Printf("[%s] There was an error parsing the embed command: %v\n", location, e)
	}

	return
}
