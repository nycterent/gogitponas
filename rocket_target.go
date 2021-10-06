package main

import (
	"gogitponas/gitlab"
	"gogitponas/registry"
	"gogitponas/rocketchat"

	"log"
)

var _ registry.Callback = (*Rocket)(nil)

type Rocket struct {
	gmi gitlab.GitlabMergeInformation
	rc  *rocketchat.RocketChat
}

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

func (r *Rocket) Set(i interface{}) {
	r.gmi = i.(gitlab.GitlabMergeInformation)
}
