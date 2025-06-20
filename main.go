package main

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/RevZer0/zkpaste-cli/handler"
	"github.com/RevZer0/zkpaste-cli/service"
	"github.com/RevZer0/zkpaste-cli/utils"
)

func main() {
	test_decrypt()
}

func test_decrypt() {
	pasteUrl := "http://localhost:3000/paste/ecf3b22d-d6cc-4178-8787-23aac5ba2943#HTCe3hASUsmhzkHffaZiwa1zV+FWpcy0VaP5WgyJmVk="
	parsed, _ := url.Parse(pasteUrl)
	splitPath := strings.Split(parsed.Path, "/")

	pasteId := splitPath[len(splitPath)-1]
	key := parsed.Fragment

	pasteData := handler.GetPasteData(pasteId)
	plaintext := service.DecryptPaste(
		utils.DearmorValue(pasteData.Ciphertext),
		utils.DearmorValue(pasteData.Iv),
		utils.DearmorValue(key),
		"12345",
	)
	fmt.Println(plaintext)
}

func test_encrypt() {
	message := "This is password protected paste"
	password := "12345"

	ciphertext, iv, key, signature := service.EncryptPaste(message, password)

	pasteId := handler.CreatePasteHandler(
		utils.ArmorValue(ciphertext),
		utils.ArmorValue(iv),
		utils.ArmorValue(signature),
	)
	fmt.Println("http://localhost:3000/paste/" + pasteId + "#" + utils.ArmorValue(key))
}
