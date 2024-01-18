package app

import (
	"github.com/aws/aws-sdk-go/aws/session"

	"github.com/kmesiab/go-key-rotator-cli/types"
)

type Command struct {
	KeyRotator types.KeyRotatorInterface
	AWSSession *session.Session
}
