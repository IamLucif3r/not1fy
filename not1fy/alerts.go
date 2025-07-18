package not1fy

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

var webhookURL string

func SetWebhookURL(url string) {
	webhookURL = url
}

type Message struct {
	Content string `json:"content"`
}

type JSONPayload map[string]interface{}

type DiscordMessage struct {
	Content         string           `json:"content,omitempty"`
	Embeds          []Embed          `json:"embeds,omitempty"`
	AllowedMentions *AllowedMentions `json:"allowed_mentions,omitempty"`
}

type Embed struct {
	Title       string  `json:"title,omitempty"`
	Description string  `json:"description,omitempty"`
	Color       int     `json:"color,omitempty"`
	Fields      []Field `json:"fields,omitempty"`
}

type Field struct {
	Name   string `json:"name"`
	Value  string `json:"value"`
	Inline bool   `json:"inline,omitempty"`
}

type AllowedMentions struct {
	Parse []string `json:"parse"`
}

func SendAlert(messageContent string) error {
	if webhookURL == "" {
		return fmt.Errorf("webhook URL not set")
	}

	msg := Message{
		Content: messageContent,
	}
	return SendJSONPayload(msg)
}

func SendCustomMessage(discordMsg DiscordMessage) error {
	if webhookURL == "" {
		return fmt.Errorf("webhook URL not set")
	}

	return SendJSONPayload(discordMsg)
}

func SendJSONPayload(payload any) error {
	if webhookURL == "" {
		return fmt.Errorf("webhook URL not set")
	}

	msgJSON, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", webhookURL, bytes.NewBuffer(msgJSON))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}
