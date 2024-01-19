package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws/session"
	rotator "github.com/kmesiab/go-key-rotator"
	log "github.com/kmesiab/go-klogger"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/aws/aws-sdk-go/aws"

	"github.com/kmesiab/go-key-rotator-cli/args"
	"github.com/kmesiab/go-key-rotator-cli/cmd_fetch"
	"github.com/kmesiab/go-key-rotator-cli/cmd_generate"
	"github.com/kmesiab/go-key-rotator-cli/cmd_rotate"
)

var rootCmd = &cobra.Command{
	Use:   "go-rotate",
	Short: "go-rotate is a CLI tool for managing RSA key rotation",
	Long: `
go-rotate is a tool for generating, storing, and retrieving
public/private RSA key pairs using AWS Parameter store.
	`,
}

func main() {
	log.InitializeGlobalLogger(logrus.InfoLevel, &logrus.TextFormatter{
		ForceColors:      true,
		DisableQuote:     true,
		DisableTimestamp: true,
		QuoteEmptyFields: true,
	})

	printGreeting()

	// Set the default command to show help
	rootCmd.Run = runShowHelp

	var (
		err  error
		sess *session.Session
	)

	config := aws.NewConfig() //.WithRegion("us-west-2")
	if sess, err = session.NewSession(config); err != nil {
		fmt.Printf("Error creating AWS config: %s\n", err)

		return
	}

	// Add sub commands and initialize their flags
	if err := args.Init(rootCmd,
		NewGenerateCommand(sess).Run,
		NewRotateCommand(sess).Run,
		NewFetchCommand(sess).Run,
	); err != nil {
		os.Exit(1)
	}

	// Execute the command
	if err := rootCmd.Execute(); err != nil {
		log.Logf("Error executing command: %s\n", err).Error()

		os.Exit(1)
	}
}

func NewRotateCommand(sess *session.Session) cmd_rotate.RotateCommand {
	cmd := cmd_rotate.RotateCommand{
		Session: sess,
	}

	cmd.KeyRotator = rotator.NewKeyRotator(rotator.NewAWSParameterStore(sess))
	cmd.AWSSession = sess

	return cmd
}

func NewGenerateCommand(sess *session.Session) cmd_generate.GenerateCommand {
	cmd := cmd_generate.GenerateCommand{}

	cmd.KeyRotator = rotator.NewKeyRotator(rotator.NewAWSParameterStore(sess))
	cmd.AWSSession = sess

	return cmd
}

func NewFetchCommand(sess *session.Session) cmd_fetch.FetchCommand {
	cmd := cmd_fetch.FetchCommand{}

	cmd.KeyRotator = rotator.NewKeyRotator(rotator.NewAWSParameterStore(sess))
	cmd.AWSSession = sess

	return cmd
}

func printGreeting() {
	fmt.Println(`
┏┓┏┓  ┏┓┏┓╋┏┓╋┏┓
┗┫┗┛  ┛ ┗┛┗┗┻┗┗ 
 ┛
	`)
}
