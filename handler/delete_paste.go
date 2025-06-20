package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type DeletePastePayload struct {
	Signature string `json:"signature"`
}

func DeletePaste(pasteId string, signature string) {
	payload, _ := json.Marshal(DeletePastePayload{
		Signature: signature,
	})
	request, _ := http.NewRequest(
		http.MethodPost,
		"http://localhost:8000/paste/"+pasteId+"/delete",
		bytes.NewBuffer(payload),
	)
	request.Header.Add("Content-type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
}
