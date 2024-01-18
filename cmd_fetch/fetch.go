package cmd_fetch

import (
	klog "github.com/kmesiab/go-klogger"

	"github.com/kmesiab/go-key-rotator-cli/app"
	"github.com/kmesiab/go-key-rotator-cli/args"
	"github.com/kmesiab/go-key-rotator-cli/aws"

	"github.com/spf13/cobra"
)

type FetchCommand struct {
	app.Command
}

func (app FetchCommand) Run(cmd *cobra.Command, _ []string) {
	klog.Logf("Fetching keys! ").Info()

	if !aws.IsValidParameterStoreName(args.GetName(cmd)) {
		klog.Logf(aws.ParameterStoreNamingRequirementsString, args.GetName(cmd)).Error()

		return
	}
}
