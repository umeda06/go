package main

import (
	"bytes"
	"encoding/json"
	_ "fmt"
	"net/http"
	_ "net/url"
)

type Push struct {
	To []string `json:"to"`
	Messages []Message `json:"messages"`
}

type Message struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

func main() {
	to := []string{"U68a1ff1883b23c5b65c6c7115e88b514"}
	messages := []Message{Message{Type: "text", Text: "じろりんちょ"}}
	push := Push{To: to, Messages: messages}
	b, _ := json.Marshal(&push)
	
	req, _ := http.NewRequest("POST", "https://api.line.me/v2/bot/message/multicast", bytes.NewBuffer(b))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer du4lrrAEzOclxVvdh9aCR7tyqCJmWnByE0BuKPH4n2LZPHRa0BvR4KxBccZqSye/EyWYQLeO9wAcgjalueHdFovYj1vqP4AKOW9ykTWIWisXWoQ5qtIKEXtlnCGsfp8nIFbXwJcROjeMJ9U4/e11zgdB04t89/1O/w1cDnyilFU=")

	client := &http.Client{}
	resp, _ := client.Do(req)
	
	defer resp.Body.Close()
}
