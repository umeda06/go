package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	body := []byte(`{
		"events": [
			{
				"replyToken": "0f3779fba3b349968c5d07db31eab56f",
				"type": "message",
				"timestamp": 1462629479859,
				"source": {
					"type": "user",
					"userId": "U4af4980629..."
				},
				"message": {
					"id": "325708",
					"type": "text",
					"text": "Hello, world"
				}
			},
			{
				"replyToken": "8cf9239d56244f4197887e939187e19e",
				"type": "follow",
				"timestamp": 1462629479859,
				"source": {
					"type": "user",
					"userId": "U4af4980629..."
				}
			}
		]
	}`)

	var events interface{}
	json.Unmarshal(body, &events)
	
	replyToken := events.(map[string]interface{})["events"].([]interface{})[0].(map[string]interface{})["replyToken"].(string)
	userId := events.(map[string]interface{})["events"].([]interface{})[0].(map[string]interface{})["source"].(map[string]interface{})["userId"].(string)
	fmt.Println(replyToken)
	fmt.Println(userId)
}
