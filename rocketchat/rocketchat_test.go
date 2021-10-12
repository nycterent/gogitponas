package rocketchat

import (
	"io"
	"testing"
)

type RocketClientTest struct {
	body []byte
	hook string
}

func (rc *RocketClientTest) Send(hook string, body io.Reader) error {
	var err error

	rc.hook = hook
	rc.body, err = io.ReadAll(body)

	return err

}

func TestSend(t *testing.T) {
	rct := &RocketClientTest{}
	var rocket = New("labashook", rct)

	rocket.Send(RocketChatMessage{
		Text: "labas1",
		Attachments: []RocketChatMessageAttachment{
			{
				Title:     "labas2",
				TitleLink: "labas3",
				Text:      "labas4",
			},
		},
	})

	//	if bytes.Compare(rct.body, []byte(`{"text":"labas1","attachments":[{"title":"labas2","title_link":"labas3","text":"labas4","image_url":null,"color":null}]}`)) != 0 {

	if string(rct.body[:len(rct.body)-1]) != `{"text":"labas1","attachments":[{"title":"labas2","title_link":"labas3","text":"labas4","image_url":null,"color":null}]}` {
		t.Fatalf("%v", rct.body[:len(rct.body)-1])
	}

}
