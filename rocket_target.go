package main

import (
	"gogitponas/gitlab"
	"gogitponas/registry"
	"gogitponas/rocketchat"
)

var _ registry.Callback = (*Rocket)(nil)

type Rocket struct {
	gmi gitlab.GitlabMergeInformation
	rc  *rocketchat.RocketChat
}

func (r Rocket) Send() {
	r.rc.Send(rocketchat.RocketChatMessage{
		Text: r.gmi.Title,
		Attachments: []rocketchat.RocketChatMessageAttachment{
			{
				Title:     r.gmi.Reference,
				TitleLink: r.gmi.MRUrl,
				Text:      r.gmi.Author,
			},
		},
	})

}

func (r *Rocket) Set(i interface{}) {
	r.gmi = i.(gitlab.GitlabMergeInformation)
}
