// Package args_test provides a comprehensive suite of tests for the args package.
// These tests ensure the correctness and reliability of the CLI command structure
// and behavior defined in the args package, which includes commands for key generation,
// rotation, and fetching. The tests are designed to cover various aspects of the command
// implementation, including:
//
//  1. Command args.Initialization: Tests the args.Init function to verify that all commands are
//     correctly added to the root command.
//
//  2. Command Mounting: Verifies that each command (generate, store, fetch) is properly
//     mounted with the correct flags and descriptions.
//
//  3. Flag Handling: Tests the parsing and handling of flags (--name, --size) for each command,
//     ensuring that they correctly accept valid inputs and handle invalid or missing inputs.
//
//  4. Command Execution: Simulates the execution of commands with various argument combinations,
//     checking both normal and error scenarios.
//
//  5. Callback Functions: Ensures that the provided callback functions (mockGenerateKeysRunFunc,
//     mockRotateKeysRunFunc, mockFetchRunFunc) are correctly invoked during command execution.
//
// 6. Help Text: Verifies that each command provides appropriate help text.
//
//  7. Direct Testing of Flag Attachment Functions: Tests args.AttachNameFlag and AttachSizeFlag functions
//     to ensure flags are correctly attached to commands.
//
// This test suite aims to maintain high quality and reliability of the CLI tool, making sure
// that each command behaves as expected under various scenarios. It includes both unit tests
// for individual components and integration tests that assess the combined behavior of these components.
//
// Note: While this test suite provides extensive coverage, it focuses on the command structure,
// flag handling, and basic execution flow. It does not cover the internal logic within the
// callback functions or interaction with external systems, which should be addressed in separate tests.
package args_test

import (
	"strconv"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"

	"github.com/kmesiab/go-key-rotator-cli/args"
)

func mockGenerateKeysRunFunc(_ *cobra.Command, _ []string) {
	// Mock implementation
}

func mockFetchRunFunc(_ *cobra.Command, _ []string) {
	// Mock implementation
}

func mockRotateKeysRunFunc(_ *cobra.Command, _ []string) {
	// Mock implementation
}

func TestGenerateCommandSizeFlag(t *testing.T) {
	cmd, err := args.MountGenerateCommand(mockGenerateKeysRunFunc)
	assert.NoError(t, err)

	// Test with a valid size
	err = cmd.ParseFlags([]string{"--size", "4096"})
	assert.NoError(t, err)
	size, err := cmd.Flags().GetInt(args.FlagStringSize)
	assert.NoError(t, err)
	assert.Equal(t, 4096, size)

	// Test with missing size flag (should use default)
	cmd, err = args.MountGenerateCommand(mockGenerateKeysRunFunc) // Create new command to reset flags
	assert.NoError(t, err)
	err = cmd.ParseFlags([]string{})
	assert.NoError(t, err)
	size, err = cmd.Flags().GetInt(args.FlagStringSize)
	assert.NoError(t, err)
	assert.Equal(t, args.DefaultKeySize, size) // args.DefaultKeySize is the expected default value
}

func TestGenerateCommandNameFlag(t *testing.T) {
	// Test with a valid name
	cmd, err := args.MountGenerateCommand(mockGenerateKeysRunFunc)
	assert.NoError(t, err)
	err = cmd.ParseFlags([]string{"--name", "testKey"})
	assert.NoError(t, err)
	name, err := cmd.Flags().GetString(args.FlagStringName)
	assert.NoError(t, err)
	assert.Equal(t, "testKey", name)

	// Test with missing name flag
	// Create a new command instance to reset flag values
	cmd, err = args.MountGenerateCommand(mockGenerateKeysRunFunc)
	assert.NoError(t, err)
	err = cmd.ParseFlags([]string{})
	assert.NoError(t, err)
	name, err = cmd.Flags().GetString(args.FlagStringName)
	assert.NoError(t, err)
	assert.Equal(t, args.DefaultName, name) // args.DefaultName should be the expected default value
}

func TestInit(t *testing.T) {
	rootCmd := &cobra.Command{Use: "root"}
	err := args.Init(rootCmd, nil, nil, nil)
	assert.NoError(t, err)
	assert.Len(t, rootCmd.Commands(), 3)

	// Check for specific commands
	foundGenerate := false
	foundRotate := false
	foundFetch := false
	for _, cmd := range rootCmd.Commands() {
		switch cmd.Use {
		case "generate":
			foundGenerate = true
		case "store":
			foundRotate = true
		case "fetch":
			foundFetch = true
		}
	}
	assert.True(t, foundGenerate, "generate command not added")
	assert.True(t, foundRotate, "store command not added")
	assert.True(t, foundFetch, "fetch command not added")
}

func mockCommandRunFunc(_ *cobra.Command, _ []string) {}

func TestInitWithError(t *testing.T) {
	rootCmd := &cobra.Command{Use: "root"}
	err := args.Init(rootCmd, mockCommandRunFunc, mockCommandRunFunc, mockCommandRunFunc)
	assert.NoError(t, err)
}

func TestMountGenerateCommand(t *testing.T) {
	cmd, err := args.MountGenerateCommand(nil)
	assert.NoError(t, err)
	assert.NotNil(t, cmd)
	assert.Equal(t, "generate", cmd.Use)
	assert.NotNil(t, cmd.Flags().Lookup(args.FlagStringName))
	assert.NotNil(t, cmd.Flags().Lookup(args.FlagStringSize))
}

func TestMountRotateCommand(t *testing.T) {
	cmd, err := args.MountRotateCommand(nil)
	assert.NoError(t, err)
	assert.NotNil(t, cmd)
	assert.Equal(t, "store", cmd.Use)
	assert.NotNil(t, cmd.Flags().Lookup(args.FlagStringName))
	assert.NotNil(t, cmd.Flags().Lookup(args.FlagStringSize))
}

func TestMountFetchCommand(t *testing.T) {
	cmd, err := args.MountFetchCommand(nil)
	assert.NoError(t, err)
	assert.NotNil(t, cmd)
	assert.Equal(t, "fetch", cmd.Use)
	assert.NotNil(t, cmd.Flags().Lookup(args.FlagStringName))
	// No size flag for fetch command
}

func TestRotateCommandNameFlag(t *testing.T) {
	cmd, err := args.MountRotateCommand(mockRotateKeysRunFunc)
	assert.NoError(t, err)

	// Test with a valid name
	err = cmd.ParseFlags([]string{"--name", "rotateKey"})
	assert.NoError(t, err)
	name, err := cmd.Flags().GetString(args.FlagStringName)
	assert.NoError(t, err)
	assert.Equal(t, "rotateKey", name)

	// Test with missing name flag
	cmd, err = args.MountRotateCommand(mockRotateKeysRunFunc) // Create new command to reset flags
	assert.NoError(t, err)
	err = cmd.ParseFlags([]string{})
	assert.NoError(t, err)
	name, err = cmd.Flags().GetString(args.FlagStringName)
	assert.NoError(t, err)
	assert.Equal(t, args.DefaultName, name) // args.DefaultName should be the expected default value
}

// TestRotateCommandSizeFlag tests the --size flag for the rotate command.
func TestRotateCommandSizeFlag(t *testing.T) {
	cmd, err := args.MountRotateCommand(mockRotateKeysRunFunc)
	assert.NoError(t, err)

	// Test with a valid size
	err = cmd.ParseFlags([]string{"--size", "4096"})
	assert.NoError(t, err)
	size, err := cmd.Flags().GetInt(args.FlagStringSize)
	assert.NoError(t, err)
	assert.Equal(t, 4096, size)

	// Test with missing size flag (should use default)
	cmd, err = args.MountRotateCommand(mockRotateKeysRunFunc) // Create new command to reset flags
	assert.NoError(t, err)
	err = cmd.ParseFlags([]string{})
	assert.NoError(t, err)
	size, err = cmd.Flags().GetInt(args.FlagStringSize)
	assert.NoError(t, err)
	assert.Equal(t, args.DefaultKeySize, size) // args.DefaultKeySize is the expected default value
}

// TestRotateCommandInvalidSizeFlag tests error handling for invalid --size flag value.
func TestRotateCommandInvalidSizeFlag(t *testing.T) {
	cmd, err := args.MountRotateCommand(mockRotateKeysRunFunc)
	assert.NoError(t, err)

	// Test with an invalid size
	err = cmd.ParseFlags([]string{"--size", "invalid"})
	assert.Error(t, err)
}

func TestFetchCommandNameFlag(t *testing.T) {
	cmd, err := args.MountFetchCommand(mockFetchRunFunc)
	assert.NoError(t, err)

	// Test with a valid name
	cmd.SetArgs([]string{"--name", "fetchKey"})
	err = cmd.Execute()
	assert.NoError(t, err)
	name, _ := cmd.Flags().GetString(args.FlagStringName)
	assert.Equal(t, "fetchKey", name)
}

func TestFetchCommandMissingRequiredFlag(t *testing.T) {
	cmd, err := args.MountFetchCommand(mockFetchRunFunc)
	assert.NoError(t, err)

	// Test with no flags
	cmd.SetArgs([]string{})
	err = cmd.Execute()
	assert.Error(t, err) // Expect an error since --name is required
}

func TestGenerateCommandCallbackFunction(t *testing.T) {
	var callbackInvoked bool
	mockRunFunc := func(cmd *cobra.Command, args []string) {
		callbackInvoked = true
	}

	cmd, err := args.MountGenerateCommand(mockRunFunc)
	assert.NoError(t, err)

	cmd.SetArgs([]string{"--name", "testKey", "--size", "2048"})
	err = cmd.Execute()
	assert.NoError(t, err)
	assert.True(t, callbackInvoked, "The callback function should be invoked")
}

// Help Text

func TestGenerateCommandHelpText(t *testing.T) {
	cmd, err := args.MountGenerateCommand(mockGenerateKeysRunFunc)
	assert.NoError(t, err)

	// Check for non-empty help text
	helpText := cmd.Short
	assert.NotEmpty(t, helpText)
}

func TestFetchCommandHelpText(t *testing.T) {
	cmd, err := args.MountFetchCommand(mockFetchRunFunc)
	assert.NoError(t, err)

	// Check for non-empty help text
	helpText := cmd.Short
	assert.NotEmpty(t, helpText)
}

func TestRotateCommandHelpText(t *testing.T) {
	cmd, err := args.MountRotateCommand(mockRotateKeysRunFunc)
	assert.NoError(t, err)

	// Check for non-empty help text
	helpText := cmd.Short
	assert.NotEmpty(t, helpText)
}

func TestGenerateCommandInvalidFlagValue(t *testing.T) {
	cmd, err := args.MountGenerateCommand(mockGenerateKeysRunFunc)
	assert.NoError(t, err)

	// Test with an invalid size value
	cmd.SetArgs([]string{"--size", "invalidSize"})
	err = cmd.Execute()
	assert.Error(t, err)
}

func TestGenerateCommandSuccess(t *testing.T) {
	cmd, err := args.MountGenerateCommand(mockGenerateKeysRunFunc)
	assert.NoError(t, err)

	cmd.SetArgs([]string{"--name", "testKey", "--size", "2048"})
	err = cmd.Execute()
	assert.NoError(t, err)
}

func TestGenerateCommandArgumentCombinations(t *testing.T) {
	// Valid arguments
	cmd, _ := args.MountGenerateCommand(mockGenerateKeysRunFunc)
	cmd.SetArgs([]string{"--name", "validName", "--size", "2048"})
	err := cmd.Execute()
	assert.NoError(t, err)

	// Invalid size
	cmd, _ = args.MountGenerateCommand(mockGenerateKeysRunFunc)
	cmd.SetArgs([]string{"--name", "validName", "--size", "invalidSize"})
	err = cmd.Execute()
	assert.Error(t, err)

	// Missing name
	cmd, _ = args.MountGenerateCommand(mockGenerateKeysRunFunc)
	cmd.SetArgs([]string{"--size", "2048"})
	err = cmd.Execute()
	assert.Error(t, err)
}

func TestRotateCommandArgumentCombinations(t *testing.T) {
	// Valid arguments
	cmd, _ := args.MountRotateCommand(mockRotateKeysRunFunc)
	cmd.SetArgs([]string{"--name", "validName", "--size", "2048"})
	err := cmd.Execute()
	assert.NoError(t, err)

	// Invalid size
	cmd, _ = args.MountRotateCommand(mockRotateKeysRunFunc)
	cmd.SetArgs([]string{"--name", "validName", "--size", "invalidSize"})
	err = cmd.Execute()
	assert.Error(t, err)

	// Missing name
	cmd, _ = args.MountRotateCommand(mockRotateKeysRunFunc)
	cmd.SetArgs([]string{"--size", "2048"})
	err = cmd.Execute()
	assert.Error(t, err)
}

func TestFetchCommandArgumentCombinations(t *testing.T) {
	// Valid name
	cmd, _ := args.MountFetchCommand(mockFetchRunFunc)
	cmd.SetArgs([]string{"--name", "validName"})
	err := cmd.Execute()
	assert.NoError(t, err)

	// Missing name
	cmd, _ = args.MountFetchCommand(mockFetchRunFunc)
	cmd.SetArgs([]string{})
	err = cmd.Execute()
	assert.Error(t, err)
}

func TestAttachNameFlag(t *testing.T) {
	cmd := &cobra.Command{Use: "testCommand"}
	err := args.AttachNameFlag(cmd)
	assert.NoError(t, err)

	flag := cmd.Flags().Lookup(args.FlagStringName)
	assert.NotNil(t, flag, "Flag should be attached to the command")
	assert.Equal(t, args.FlagStringName, flag.Name)
	assert.Equal(t, args.FlagStringNameShorthand, flag.Shorthand)
	assert.Equal(t, args.DefaultName, flag.DefValue)
}

func TestAttachSizeFlag(t *testing.T) {
	cmd := &cobra.Command{Use: "testCommand"}
	err := args.AttachSizeFlag(cmd)
	assert.NoError(t, err) // This test is limited because AttachSizeFlag doesn't return an error in your current implementation

	flag := cmd.Flags().Lookup(args.FlagStringSize)
	assert.NotNil(t, flag, "Flag should be attached to the command")
	assert.Equal(t, args.FlagStringSize, flag.Name)
	assert.Equal(t, args.FlagStringSizeShorthand, flag.Shorthand)
	assert.Equal(t, strconv.Itoa(args.DefaultKeySize), flag.DefValue)
}

func TestGetSize(t *testing.T) {
	cmd := &cobra.Command{}
	cmd.Flags().String(args.FlagStringSize, "10", "")
	size := args.GetSize(cmd)
	expectedSize := "10"
	if size != expectedSize {
		t.Errorf("Expected size %s, but got %s", expectedSize, size)
	}
}
