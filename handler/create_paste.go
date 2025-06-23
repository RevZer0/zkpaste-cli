package handler

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/RevZer0/zkpaste-cli/config"
)

type CreatePasteMetadata struct {
	Ttl               int  `json:"ttl"`
	PasswordProtected bool `json:"password_protected"`
	ViewLimit         int  `json:"opens_count"`
}

type CreatePastePayload struct {
	Ciphertext string              `json:"ciphertext"`
	Iv         string              `json:"iv"`
	Signature  string              `json:"signature"`
	Metadata   CreatePasteMetadata `json:"metadata"`
}

type CreatePayloadResponse struct {
	PasteId string `json:"paste_id"`
}

func CreatePasteHandler(
	ciphertext, iv, signature string,
	ttl, viewLimit int,
	passwordProtected bool,
) string {
	payload := CreatePastePayload{
		Ciphertext: ciphertext,
		Iv:         iv,
		Signature:  signature,
		Metadata: CreatePasteMetadata{
			Ttl:               ttl,
			PasswordProtected: passwordProtected,
			ViewLimit:         viewLimit,
		},
	}
	jsonPayload, _ := json.Marshal(payload)

	request, _ := http.NewRequest(
		http.MethodPost,
		config.ZKPasteConfig.URL.CoreApi+"/paste",
		bytes.NewBuffer(jsonPayload),
	)
	request.Header.Add("Content-type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	jsonResponse, _ := io.ReadAll(resp.Body)

	var response CreatePayloadResponse
	json.Unmarshal(jsonResponse, &response)

	return response.PasteId
}
