package args

import (
	"github.com/spf13/cobra"
)

// Constants.  Real life constant values
const (
	// arg: --name

	DefaultName             = ""
	FlagStringName          = "name"
	FlagStringNameShorthand = "n"

	// arg: --size

	DefaultKeySize          = 2048
	FlagStringSize          = "size"
	FlagStringSizeShorthand = "s"
)

type CommandRunFunc func(cmd *cobra.Command, args []string)

func Init(rootCmd *cobra.Command, runGenerateKeys, RunRotateKeys, runFetch CommandRunFunc) error {
	var (
		err         error
		getCmd      *cobra.Command
		rotateCmd   *cobra.Command
		generateCmd *cobra.Command
	)

	if generateCmd, err = MountGenerateCommand(runGenerateKeys); err != nil {
		return err
	}

	if rotateCmd, err = MountRotateCommand(RunRotateKeys); err != nil {
		return err
	}

	if getCmd, err = MountFetchCommand(runFetch); err != nil {
		return err
	}

	rootCmd.AddCommand(generateCmd)
	rootCmd.AddCommand(rotateCmd)
	rootCmd.AddCommand(getCmd)

	return nil
}

func MountGenerateCommand(runGenerateKeys CommandRunFunc) (*cobra.Command, error) {
	generateCmd := &cobra.Command{
		Use:   "generate",
		Short: "Generates a new public/private key pair, but does not store it",
		Run: func(cmd *cobra.Command, args []string) {
			runGenerateKeys(cmd, args)
		},
	}

	// --name flag
	if err := AttachNameFlag(generateCmd); err != nil {
		return nil, err
	}

	// --size flag
	if err := AttachSizeFlag(generateCmd); err != nil {
		return nil, err
	}

	return generateCmd, nil
}

func MountRotateCommand(RunRotateKeys CommandRunFunc) (*cobra.Command, error) {
	rotateCommand := &cobra.Command{
		Use:   "store",
		Short: "Generates and stores a public/private key pair",
		Run: func(cmd *cobra.Command, args []string) {
			RunRotateKeys(cmd, args)
		},
	}

	// --name flag
	if err := AttachNameFlag(rotateCommand); err != nil {
		return nil, err
	}

	// --size flag
	if err := AttachSizeFlag(rotateCommand); err != nil {
		return nil, err
	}

	return rotateCommand, nil
}

func MountFetchCommand(runFetch CommandRunFunc) (*cobra.Command, error) {
	getCommand := &cobra.Command{
		Use:   "fetch",
		Short: "Downloads your public/private key pair",
		Run: func(cmd *cobra.Command, args []string) {
			runFetch(cmd, args)
		},
	}

	// --name flag
	if err := AttachNameFlag(getCommand); err != nil {
		return nil, err
	}

	return getCommand, nil
}

func AttachNameFlag(cmd *cobra.Command) error {
	// --name flag
	cmd.Flags().StringP(
		FlagStringName, FlagStringNameShorthand, "",
		"Specify the name prefix for your keys. Keys will have "+
			"_public.pem and _private.pem appended to the name")

	if err := cmd.MarkFlagRequired("name"); err != nil {
		return err
	}

	return nil
}

func AttachSizeFlag(cmd *cobra.Command) error {
	var size int

	// --size flag
	cmd.Flags().IntVarP(&size, FlagStringSize, FlagStringSizeShorthand, DefaultKeySize,
		"Specify the size of your keys in bits.  Default is 2048")

	return nil
}

func GetName(cmd *cobra.Command) string {
	return cmd.Flag(FlagStringName).Value.String()
}
