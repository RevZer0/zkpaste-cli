package handler

import (
	"encoding/json"
	"io"
	"net/http"
)

type PasteData struct {
	PasteId           string `json:"paste_id"`
	Ciphertext        string `json:"paste"`
	Iv                string `json:"iv"`
	PasswordProtected bool   `json:"password_protected"`
}

func GetPasteData(pasteId string) (pasteData PasteData) {
	request, _ := http.NewRequest(
		http.MethodGet,
		"http://localhost:8000/paste/"+pasteId,
		nil,
	)
	request.Header.Add("Content-type", "application/json")
	client := &http.Client{}

	resp, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	jsonResponse, _ := io.ReadAll(resp.Body)

	json.Unmarshal(jsonResponse, &pasteData)
	return
}
