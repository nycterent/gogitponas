package main

import (
	"gogitponas/gitlab"
	"gogitponas/registry"
	"gogitponas/slack"
	"log"
)

var _ registry.Callback = (*Slack)(nil)

// Slack is stub for slack message and client
type Slack struct {
	gmi gitlab.GitlabMergeInformation
	sc  *slack.Slack
}

// Send implements sending message to slack
func (s Slack) Send() {
	err := s.sc.Send(slack.Message{
		Text: s.gmi.Title,
		Attachments: []slack.MessageAttachment{
			{
				Title:     s.gmi.Reference,
				TitleLink: s.gmi.MRURL,
				Text:      s.gmi.Author,
			},
		},
	})

	if err != nil {
		log.Fatal(err)
	}

}

// Set casts interface to slack message
func (s *Slack) Set(i interface{}) {
	s.gmi = i.(gitlab.GitlabMergeInformation)
}
