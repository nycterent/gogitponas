package main

import (
	"gogitponas/gitlab"
	"gogitponas/registry"
	"gogitponas/rocketchat"

	"log"
)

var _ registry.Callback = (*Rocket)(nil)

// Rocket sets gmi for merge information
// and rocketchat client object
type Rocket struct {
	gmi gitlab.GitlabMergeInformation
	rc  *rocketchat.RocketChat
}

// Send implements message sending for the RocketChat
func (r Rocket) Send() {
	err := r.rc.Send(rocketchat.RocketChatMessage{
		Text: r.gmi.Title,
		Attachments: []rocketchat.RocketChatMessageAttachment{
			{
				Title:     r.gmi.Reference,
				TitleLink: r.gmi.MRURL,
				Text:      r.gmi.Author,
			},
		},
	})

	if err != nil {
		log.Fatal(err)
	}

}

// Set casts gitlab merge inforfation to rocket
func (r *Rocket) Set(i interface{}) {
	r.gmi = i.(gitlab.GitlabMergeInformation)
}
