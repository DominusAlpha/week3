/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/spf13/cobra"
	telebot "gopkg.in/telebot.v4"
)

var teleToken = os.Getenv("TELE_TOKEN")

// chahlikBotCmd represents the chahlikBot command
var chahlikBotCmd = &cobra.Command{
	Use:     "chahlikBot",
	Aliases: []string{"start"},
	Short:   "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("chahlikBot %s started", AppVersion)
		chahlikBot, err := telebot.NewBot(telebot.Settings{
			URL:    "",
			Token:  teleToken,
			Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
		})

		if err != nil {
			log.Fatalf("PLease check the TELE_TOKEN env variable. %s", err)
			return
		}

		chahlikBot.Handle(telebot.OnText, func(ctx telebot.Context) error {
			log.Print(ctx.Message().Payload, ctx.Text())
			payload := ctx.Message().Text
			switch payload {
			case "hello":
				err = ctx.Send(fmt.Sprintf("Hello I am chahlikBot %s", AppVersion))

			}
			return err
		})

		chahlikBot.Start()
	},
}

func init() {
	rootCmd.AddCommand(chahlikBotCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// chahlikBotCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// chahlikBotCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
