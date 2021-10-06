package main

import (
	"fmt"
	"gogitponas/registry"
)

var _ registry.Callback = (*Slack)(nil)

// Slack is stub for slack message and client
type Slack struct {
	name string
}

// Send implements sending message to slack
func (s Slack) Send() {
	fmt.Printf("jee, slack, %v\n", s.name)
}

// Set casts interface to slack message
func (s *Slack) Set(i interface{}) {
	s.name = i.(string)
}
