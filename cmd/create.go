package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"

	"github.com/RevZer0/zkpaste-cli/config"
	"github.com/RevZer0/zkpaste-cli/handler"
	"github.com/RevZer0/zkpaste-cli/service"
	"github.com/RevZer0/zkpaste-cli/utils"
)

var (
	password      string
	ttl           string
	viewLimit     int
	usePipedInput bool
	plaintext     string

	createCmd = &cobra.Command{
		Use:   "create [message]",
		Short: "Create encrypted paste and get URL to share",
		Args: func(cmd *cobra.Command, args []string) error {
			fileInfo, _ := os.Stdin.Stat()
			if fileInfo.Mode()&os.ModeCharDevice == 0 {
				usePipedInput = true
				return nil
			}
			// check for input
			if err := cobra.MinimumNArgs(1)(cmd, args); err != nil {
				return err
			}
			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			ttlMap := map[string]int{
				"10m": 600,
				"30m": 1800,
				"1h":  3600,
				"1d":  86400,
				"5d":  432000,
				"1w":  604800,
			}
			ttlInt := 86400
			if _, ok := ttlMap[ttl]; ok {
				ttlInt = ttlMap[ttl]
			}
			if usePipedInput {
				pipedText := make([]string, 0)
				scanner := bufio.NewScanner(os.Stdin)
				for scanner.Scan() {
					pipedText = append(pipedText, scanner.Text())
				}
				plaintext = strings.Join(pipedText, "\n")
			} else {
				plaintext = args[0]
			}
			ciphertext, iv, key, signature := service.EncryptPaste(plaintext, password)
			pasteId := handler.CreatePasteHandler(
				utils.ArmorValue(ciphertext),
				utils.ArmorValue(iv),
				utils.ArmorValue(signature),
				ttlInt,
				viewLimit,
				len(password) > 0,
			)
			fmt.Println("Paste has been created.")
			fmt.Print("Your paste URL: ")
			fmt.Println(
				config.ZKPasteConfig.URL.Public + "/paste/" + pasteId + "#" + utils.ArmorValue(key),
			)
		},
	}
)

func init() {
	createCmd.Flags().StringVarP(&password, "password", "p", "", "Set password for the paste")
	createCmd.Flags().
		StringVarP(&ttl, "ttl", "t", "1d", "Set paste expiration time. Possible values are: 10m, 30m, 1h, 1d, 5d, 1w. Default is 1 day")
	createCmd.Flags().
		IntVarP(&viewLimit, "views", "l", 0, "Set paste views limit. Default is unlimited")
	rootCmd.AddCommand(createCmd)
}
