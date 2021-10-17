package main

import (
	"gogitponas/gitlab"
	"gogitponas/registry"
	"gogitponas/slack"
	"log"
)

var _ registry.Callback = (*Slack)(nil)

// Slack is stub for slack client
type Slack struct {
	sc *slack.Slack
}

// Send implements sending message to slack
func (s Slack) Send(i interface{}) {
	gmi := i.(gitlab.MergeInformation)
	err := s.sc.Send(slack.Message{
		Text: gmi.Title,
		Attachments: []slack.MessageAttachment{
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
