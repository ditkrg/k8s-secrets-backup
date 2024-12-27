package options

import (
	"errors"

	"github.com/rs/zerolog/log"
)

type Options struct {
	S3          S3          `env:",prefix=S3__"`
	Secret      K8sSecret   `env:",prefix=SECRET__"`
	ClusterInfo ClusterInfo `env:",prefix=CLUSTER__"`

	AgeRecipientPublicKey string `env:"AGE_RECIPIENT_PUBLIC_KEY"`
}

func (opts *Options) Validate() error {
	log.Info().Msg("Validating options")

	// #############################
	// validate the S3 options
	// #############################
	if err := opts.S3.Validate(); err != nil {
		return err
	}

	// #############################
	// validate the Secret options
	// #############################
	if err := opts.Secret.Validate(); err != nil {
		return err
	}

	// #############################
	// validate the ClusterInfo options
	// #############################
	if err := opts.ClusterInfo.Validate(); err != nil {
		return err
	}

	// #############################
	// validate the age recipient public key
	// #############################
	if opts.AgeRecipientPublicKey == "" {
		return errors.New("AGE_RECIPIENT_PUBLIC_KEY is required")
	}

	log.Info().Msg("Options are valid")

	return nil
}
