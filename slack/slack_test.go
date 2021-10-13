package slack

import (
	"io"
	"testing"
)

type SlackClientTest struct {
	body []byte
	hook string
}

func (sc *SlackClientTest) Send(hook string, body io.Reader) error {
	var err error

	sc.hook = hook
	sc.body, err = io.ReadAll(body)

	return err

}

func TestSend(t *testing.T) {
	sct := &SlackClientTest{}
	var slack = New("labashook0", sct)

	slack.Send(Message{
		Text: "labas1",
		Attachments: []MessageAttachment{
			{
				Title:     "labas2",
				TitleLink: "labas3",
				Text:      "labas4",
			},
		},
	})

	if string(sct.body[:len(sct.body)-1]) != `{"text":"labas1","attachments":[{"title":"labas2","title_link":"labas3","text":"labas4","image_url":null,"color":null}]}` {
		t.Fatalf("%v", sct.body[:len(sct.body)-1])
	}

}
