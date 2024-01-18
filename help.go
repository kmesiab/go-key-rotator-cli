package main

import (
	log "github.com/kmesiab/go-klogger"
	"github.com/spf13/cobra"
)

func runShowHelp(cmd *cobra.Command, args []string) {
	subCmd, _, err := rootCmd.Find(args)

	if err != nil || subCmd == nil {
		// Show the primary help message
		if err := cmd.Root().Help(); err != nil {
			log.Logf("Error showing help: %s", err).Error()
		}
	} else {
		// Show the sub command help message
		if err := subCmd.Root().Help(); err != nil {
			log.Logf("Error showing help: %s", err).Error()
		}
	}
}
