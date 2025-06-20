package handler

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/RevZer0/zkpaste-cli/config"
)

type ViewPastePayload struct {
	Signature string `json:"signature"`
}

func ViewPaste(pasteId string, signature string) {
	payload, _ := json.Marshal(DeletePastePayload{
		Signature: signature,
	})
	request, _ := http.NewRequest(
		http.MethodPut,
		config.ZKPasteConfig.URL.CoreApi+"/paste/"+pasteId+"/view",
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
