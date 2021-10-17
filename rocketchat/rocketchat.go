package rocketchat

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

// ChatMessage json struct for rocketchat message
type ChatMessage struct {
	Text        string                  `json:"text"`
	Attachments []ChatMessageAttachment `json:"attachments"`
}

// ChatMessageAttachment json struct for rocketchat message attachment
type ChatMessageAttachment struct {
	Title     string  `json:"title"`
	TitleLink string  `json:"title_link"`
	Text      string  `json:"text"`
	ImageURL  *string `json:"image_url"`
	Color     *string `json:"color"`
}

// Chat client serializer, kinda starts to define "class"
type Chat struct {
	hook   string
	client ClientInterface
}

// Send encodes the payload and invokes real client.Send
func (rc Chat) Send(message ChatMessage) error {
	payloadBuf := new(bytes.Buffer)

	err := json.NewEncoder(payloadBuf).Encode(message)

	if err != nil {
		log.Fatal(err)
	}

	return rc.client.Send(rc.hook, payloadBuf)
}

// ClientInterface represents Sending notification
type ClientInterface interface {
	Send(string, io.Reader) error
}

// Client holds http client
type Client struct {
	http *http.Client
}

// Send sends the payload
func (rc Client) Send(hook string, body io.Reader) error {

	resp, err := rc.http.Post(hook, "application/json", body)

	if err != nil {
		log.Println(err)
	}
	b, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	log.Printf("%s", b)

	return nil

}

// New constructs rocketchat.Chat
func New(RHook string, client ClientInterface) *Chat {
	return &Chat{
		hook:   RHook,
		client: client,
	}
}

// NewClient constructs rocketchat Client from  http
func NewClient() Client {
	return Client{http: &http.Client{}}
}
