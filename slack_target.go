package main

import (
	"fmt"
	"gogitponas/registry"
)

var _ registry.Callback = (*Slack)(nil)

type Slack struct {
	name string
}

func (s Slack) Send() {
	fmt.Printf("jee, slack, %v\n", s.name)
}

func (s *Slack) Set(i interface{}) {
	s.name = i.(string)
}
