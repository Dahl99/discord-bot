package discordbot

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
	"sync"

	"github.com/bwmarrin/discordgo"
	"github.com/jonas747/dca"
)

type VoiceInstance struct {
	voice      *discordgo.VoiceConnection
	session    *discordgo.Session
	encoder    *dca.EncodeSession
	stream     *dca.StreamingSession
	run        *exec.Cmd
	queueMutex sync.Mutex
	audioMutex sync.Mutex
	nowPlaying Song
	queue      []Song
	recv       []int16
	guildID    string
	channelID  string
	speaking   bool
	pause      bool
	stop       bool
	skip       bool
	radioFlag  bool
}

type Song struct {
	ChannelID string
	User      string
	ID        string
	VidID     string
	Title     string
	Duration  string
	VideoURL  string
}

var (
	voiceInstances = map[string]*VoiceInstance{}
	mutex sync.Mutex
)

func music(cmd []string, v *VoiceInstance, s *discordgo.Session, m *discordgo.MessageCreate) {

	if len(cmd) < 1 {
		return
	}

	fmt.Println(cmd)

	switch(cmd[0]) {
	case "join":
		joinVoice(v, s, m)
	case "play":
		playMusic(cmd[1:], s, m)
	case "skip":
		skipMusic(s, m)
	default:
		return
	}
}

func joinVoice(v *VoiceInstance, s *discordgo.Session, m *discordgo.MessageCreate) {
	voiceChannelID := searchVoiceChannel(m.Author.ID)
	if voiceChannelID == "" {
		log.Println("Voice channel id not found")
		s.ChannelMessageSend(m.ChannelID, "You need to join a voice channel first!")
		return
	}

	if v != nil {
		log.Println("Voice instance already created")
	} else {
		guildID := searchGuild(m.ChannelID)
		mutex.Lock()
		v = new(VoiceInstance)
		voiceInstances[guildID] = v
		v.guildID = guildID
		v.session = s
		mutex.Unlock()
	}

	var err error
	v.voice, err = dg.ChannelVoiceJoin(v.guildID, voiceChannelID, false, false)
	if err != nil {
		v.Stop()
		log.Println("Error when joining voice channel")
		return
	}

	v.voice.Speaking(false)
	log.Println("New voice instance created")
	s.ChannelMessageSend(m.ChannelID, "Voice channel joined!")
}

func playMusic(n []string, s *discordgo.Session, m *discordgo.MessageCreate) {

	name := replaceSpace(n)
	url := youtubeEndpoint + conf.Ytkey + "&q=" + name

	result := strings.Split(ytSearch(url), "|")

	if result[0] == ytSearchFailed || result[0] == decodingFailed {
		s.ChannelMessageSend(m.ChannelID, result[0])
	}

	// vid, err := v2.GetVideoInfo(ytVideoUrl + result[0])
	// if err != nil {
	// 	return
	// }

	// client := youtube.Client{}
	// video, err := client.GetVideo(result[0])
	// if err != nil {
	// 	return
	// }

	// go func () {
	// 	songSignal <- video
	// }()


	s.ChannelMessageSend(m.ChannelID, result[0] + "\t" + result[1])
}

func skipMusic(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSend(m.ChannelID, "Skip music!")
}

// searchVoiceChannel search the voice channel id into from guild.
func searchVoiceChannel(user string) (voiceChannelID string) {
	for _, g := range dg.State.Guilds {
		for _, v := range g.VoiceStates {
			if v.UserID == user {
				return v.ChannelID
			}
		}
	}
	return ""
}

// Stop stop the audio
func (v *VoiceInstance) Stop() {
	v.stop = true
	if v.encoder != nil {
		v.encoder.Cleanup()
	}
}
