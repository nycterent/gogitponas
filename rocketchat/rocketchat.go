package rocketchat

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// {
// 	"text": "Example message",
// 	"attachments": [
// 	  {
// 		"title": "Rocket.Chat",
// 		"title_link": "https://rocket.chat",
// 		"text": "Rocket.Chat, the best open source chat",
// 		"image_url": "http://images/integration-attachment-example.png",
// 		"color": "#764FA5"
// 	  }
// 	]
//   }

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
	hook string
	http *http.Client
}

func New(RHook string) *RocketChat {
	return &RocketChat{
		hook: RHook,
		http: &http.Client{},
	}
}

func (rc RocketChat) Send(message RocketChatMessage) error {
	payloadBuf := new(bytes.Buffer)

	json.NewEncoder(payloadBuf).Encode(message)

	resp, err := rc.http.Post(rc.hook, "application/json", payloadBuf)

	if err != nil {
		log.Println(err)
	}
	b, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	log.Printf("%s", b)
	return nil
}