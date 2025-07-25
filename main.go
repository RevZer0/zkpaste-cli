package main

import (
	"net/url"
	"strings"

	"github.com/RevZer0/zkpaste-cli/cmd"
	"github.com/RevZer0/zkpaste-cli/handler"
	"github.com/RevZer0/zkpaste-cli/service"
	"github.com/RevZer0/zkpaste-cli/utils"
)

func main() {
	cmd.Execute()
}

func test_delete_paste() {
	pasteUrl := "http://localhost:3000/paste/1d6b1057-520f-4e53-a147-d0d82739158c#3b73BnQCtZqBTGNB6nU1k9NnQcIp8YwXN8avaq7/hlc="

	parsed, _ := url.Parse(pasteUrl)
	splitPath := strings.Split(parsed.Path, "/")

	pasteId := splitPath[len(splitPath)-1]
	key := parsed.Fragment

	pasteData, _ := handler.GetPasteData(pasteId)
	plaintext, _ := service.DecryptPaste(
		utils.DearmorValue(pasteData.Ciphertext),
		utils.DearmorValue(pasteData.Iv),
		utils.DearmorValue(key),
		"12345",
	)
	signature := service.ProofOfKnowlege(utils.DearmorValue(key), plaintext, "12345")
	handler.DeletePaste(pasteData.PasteId, utils.ArmorValue(signature))
}
