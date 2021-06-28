package main

import (
	"context"
	"github.com/andersfylling/disgord"
	"github.com/sirupsen/logrus"
	"math/rand"
	"os"
	"strings"
)

var log = &logrus.Logger{
	Out:       os.Stderr,
	Formatter: new(logrus.TextFormatter),
	Hooks:     make(logrus.LevelHooks),
	Level:     logrus.InfoLevel,
}

const NewMessage = disgord.EvtMessageCreate

func main() {
	client := disgord.New(disgord.Config{
		BotToken: "nick-when-is-season-4",
		Logger:   log,
	})
	defer client.StayConnectedUntilInterrupted(context.Background())

	responses := []string{
		"Hi! There is no ETA (estimated time of arrival) set for season 4 yet, we will know when we have one. Till then, please stop spamming that question over and over again to keep things less nerving :)",
		"tomorrow",
		"when the pigs fly",
		"Well hello there! Season 4 is still being worked on by our amazing build team and the amazing dev! There is no release date yet for you to know of. Thank you for your patience!",
		"yes",
		"Season 4 is being worked on! Why not check out https://npbe.net in the meantime?",
		"liek if cri",
		"no",
		"please stop",
		"mom help me im scared",
		"what is season 4 and why is he so strong",
	}

	colors := []int{
		0xea907a,
		0xfbc687,
		0xaacdbe,
		0xfaf0af,
		0xa8e6cf,
		0xdcedc1,
		0xffd3b6,
	}

	client.On(disgord.EvtReady, func(s disgord.Session, evt *disgord.Ready) {
		err := s.UpdateStatus(&disgord.UpdateStatusPayload{
			Game:   &disgord.Activity{
				Name:          "season 4 when",
				Type:          2,
			},
			Status: disgord.StatusOnline,
		})

		if err != nil{
			panic(err)
		}
	})

	client.On(NewMessage, func(s disgord.Session, evt *disgord.MessageCreate) {
		text := strings.ToLower(evt.Message.Content)
		if (strings.Contains(text, "s4") || strings.Contains(text, "season 4")) && strings.Contains(text, "when") {
			_, _ = evt.Message.Reply(context.Background(), s, &disgord.Embed{
				Description: responses[rand.Intn(len(responses)-1)],
				Color:       colors[rand.Intn(len(colors)-1)],
			})
		}
	})
}
