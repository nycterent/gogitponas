package slack

import (
	"bytes"
	"encoding/json"
	"io"
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
	hook   string
	http   *http.Client
	client ClientInterface
}

// New a constructor
func New(SlackHook string, client ClientInterface) *Slack {
	return &Slack{
		hook:   SlackHook,
		http:   &http.Client{},
		client: client,
	}
}

// Send implements the interface to send the notification
func (sc Client) Send(hook string, body io.Reader) error {

	resp, err := sc.http.Post(hook, "application/json", body)

	if err != nil {
		log.Println(err)
	}

	b, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	log.Printf("%s", b)

	return nil

}

// ClientInterface mainly needed for testing
type ClientInterface interface {
	Send(string, io.Reader) error
}

// Client holds http client
type Client struct {
	http *http.Client
}

// NewClient constructs slack Client
func NewClient() Client {
	return Client{http: &http.Client{}}
}

// Send used to send a slack message
func (sc Slack) Send(message Message) error {
	payloadBuf := new(bytes.Buffer)

	err := json.NewEncoder(payloadBuf).Encode(message)

	if err != nil {
		log.Fatal(err)
	}

	return sc.client.Send(sc.hook, payloadBuf)
}
