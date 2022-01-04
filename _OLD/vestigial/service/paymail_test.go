package service

import (
	"errors"
	"testing"

	"github.com/nch-bowstave/paymail"
	"github.com/stretchr/testify/assert"
)

func TestAccountArgs(t *testing.T) {
	tests := map[string]struct {
		args paymail.AccountArgs
		err  error
	}{
		"only domain should return an error": {
			args: paymail.AccountArgs{
				Email: "missing.com",
			},
			err: errors.New("invalid email address"),
		},
		"name with no domain should return an error": {
			args: paymail.AccountArgs{
				Email: "bob@",
			},
			err: errors.New("invalid email address"),
		},
		"valid email should an address": {
			args: paymail.AccountArgs{
				Email: "bob@domain.com",
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			err := test.args.Validate()
			if test.err != nil {
				assert.EqualError(t, test.err, err.Error())
			}
		})
	}
}
