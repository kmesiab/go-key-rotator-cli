package cmd_generate

import (
	klog "github.com/kmesiab/go-klogger"
	"github.com/spf13/cobra"

	"github.com/kmesiab/go-key-rotator-cli/app"
	"github.com/kmesiab/go-key-rotator-cli/args"
	"github.com/kmesiab/go-key-rotator-cli/aws"
)

type GenerateCommand struct {
	app.Command
}

func (app GenerateCommand) Run(cmd *cobra.Command, _ []string) {
	klog.Logf("Rotating new keys! ").Info()

	if !aws.IsValidParameterStoreName(args.GetName(cmd)) {
		klog.Logf(aws.ParameterStoreNamingRequirementsString, args.GetName(cmd)).Error()

		return
	}

	klog.Logf("Generating new keys! ").Info()
}
