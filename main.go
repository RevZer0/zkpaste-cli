package main

import (
	"fmt"

	"github.com/RevZer0/zkpaste-cli/handler"
	"github.com/RevZer0/zkpaste-cli/service"
	"github.com/RevZer0/zkpaste-cli/utils"
)

func main() {
	test_encrypt()
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
