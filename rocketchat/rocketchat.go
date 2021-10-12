package rocketchat

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

type RocketChatMessage struct {
	Text        string                        `json:"text"`
	Attachments []RocketChatMessageAttachment `json:"attachments"`
}

type RocketChatMessageAttachment struct {
	Title     string  `json:"title"`
	TitleLink string  `json:"title_link"`
	Text      string  `json:"text"`
	ImageUrl  *string `json:"image_url"`
	Color     *string `json:"color"`
}

type RocketChat struct {
	hook   string
	http   *http.Client
	client RocketClientInterface
}

type RocketClientInterface interface {
	Send(string, io.Reader) error
}

type RocketClient struct {
	http *http.Client
}

func (rc RocketClient) Send(hook string, body io.Reader) error {

	resp, err := rc.http.Post(hook, "application/json", body)

	if err != nil {
		log.Println(err)
	}
	b, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	log.Printf("%s", b)

	return nil

}

func New(RHook string, client RocketClientInterface) *RocketChat {
	return &RocketChat{
		hook:   RHook,
		http:   &http.Client{},
		client: client,
	}
}

func NewClient() RocketClient {
	return RocketClient{http: &http.Client{}}
}

func (rc RocketChat) Send(message RocketChatMessage) error {
	payloadBuf := new(bytes.Buffer)

	err := json.NewEncoder(payloadBuf).Encode(message)

	if err != nil {
		log.Fatal(err)
	}

	return rc.client.Send(rc.hook, payloadBuf)
}
