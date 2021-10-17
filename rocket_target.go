package main

import (
	"gogitponas/gitlab"
	"gogitponas/registry"
	"gogitponas/rocketchat"

	"log"
)

var _ registry.Callback = (*Rocket)(nil)

// Rocket sets rocketchat client object
type Rocket struct {
	rc *rocketchat.Chat
}

// Send implements message sending for the RocketChat
func (r Rocket) Send(i interface{}) {
	gmi := i.(gitlab.MergeInformation)
	err := r.rc.Send(rocketchat.ChatMessage{
		Text: gmi.Title,
		Attachments: []rocketchat.ChatMessageAttachment{
			{
				Title:     gmi.Reference,
				TitleLink: gmi.MRURL,
				Text:      gmi.Author,
			},
		},
	})

	if err != nil {
		log.Fatal(err)
	}

}
