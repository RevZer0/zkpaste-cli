package cmd

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/spf13/cobra"

	"github.com/RevZer0/zkpaste-cli/handler"
	"github.com/RevZer0/zkpaste-cli/service"
	"github.com/RevZer0/zkpaste-cli/utils"
)

var (
	deletePassword string

	deleteCmd = &cobra.Command{
		Use:   "delete [pasteUrl]",
		Short: "Delete paste specified by URL",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			parsed, _ := url.Parse(args[0])
			splitPath := strings.Split(parsed.Path, "/")
			encryptionKey := parsed.Fragment
			if len(splitPath) == 0 || len(encryptionKey) == 0 {
				fmt.Println("Error: pasteUrl must be a valid zkpaste.com paste URL")
				return
			}
			pasteId := splitPath[len(splitPath)-1]
			pasteData, err := handler.GetPasteData(pasteId)
			if err != nil {
				fmt.Println("Error: " + err.Error())
				return
			}
			if pasteData.PasswordProtected && len(deletePassword) == 0 {
				fmt.Println(
					"Error: This paste is protected with the password. Use --password flag to set the password",
				)
				return
			}
			plaintext, err := service.DecryptPaste(
				utils.DearmorValue(pasteData.Ciphertext),
				utils.DearmorValue(pasteData.Iv),
				utils.DearmorValue(encryptionKey),
				deletePassword,
			)
			if err != nil {
				fmt.Println("Error: Failed to decrypt paste. Seems like URL or password is invalid")
				return
			}
			handler.DeletePaste(
				pasteId,
				utils.ArmorValue(
					service.ProofOfKnowlege(
						utils.DearmorValue(encryptionKey),
						plaintext,
						deletePassword,
					),
				),
			)
			fmt.Println("Paste has been deleted.")
		},
	}
)

func init() {
	deleteCmd.Flags().StringVarP(&deletePassword, "password", "p", "", "Password to decrypt paste")

	rootCmd.AddCommand(deleteCmd)
}

