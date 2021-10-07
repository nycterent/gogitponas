package slack

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// Message stores slack message
type Message struct {
	Text        string              `json:"text"`
	Attachments []MessageAttachment `json:"attachments"`
}

// MessageAttachment slack message formatting and body
type MessageAttachment struct {
	Title     string  `json:"title"`
	TitleLink string  `json:"title_link"`
	Text      string  `json:"text"`
	ImageURL  *string `json:"image_url"`
	Color     *string `json:"color"`
}

// Slack what we need to connect to slack
type Slack struct {
	hook string
	http *http.Client
}

// New a constructor
func New(SlackHook string) *Slack {
	return &Slack{
		hook: SlackHook,
		http: &http.Client{},
	}
}

// Send used to send a slack message
func (sc Slack) Send(message Message) error {
	payloadBuf := new(bytes.Buffer)

	err := json.NewEncoder(payloadBuf).Encode(message)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := sc.http.Post(sc.hook, "application/json", payloadBuf)

	if err != nil {
		log.Println(err)
	}

	b, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	log.Printf("%s %s", b, message.Text)
	return nil
}
