package api

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type ScoreData struct {
	Username string `json:"username"`
	Lab      string `json:"lab"`
	Score    int    `json:"score"`
}

type GeneralResponse struct {
	Message string `json:"message"`
}

func (server CustomServer) PushScore(data ScoreData) (code int, err error) {
	// create json data
	jsonData, err := json.Marshal(data)
	if err != nil {
		return code, err
	}

	// create request
	req, err := http.NewRequest("POST", server.Host+"/v1/score", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+server.Token)

	// send request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return code, err
	}

	defer resp.Body.Close()

	return resp.StatusCode, nil
}
