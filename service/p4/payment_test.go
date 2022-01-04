package service_test

import (
	"context"
	"errors"
	"testing"

	"github.com/libsv/go-bc/spv"
	"github.com/libsv/go-p4"
	p4mocks "github.com/libsv/go-p4/mocks"
	"github.com/libsv/p4-server/log"
	"github.com/libsv/p4-server/service"
	"github.com/stretchr/testify/assert"
)

func TestPayment_Create(t *testing.T) {
	tests := map[string]struct {
		paymentCreateFn func(context.Context, p4.PaymentCreateArgs, p4.Payment) (*p4.PaymentACK, error)
		args            p4.PaymentCreateArgs
		req             p4.Payment
		expErr          error
	}{
		"successful payment create": {
			paymentCreateFn: func(context.Context, p4.PaymentCreateArgs, p4.Payment) (*p4.PaymentACK, error) {
				return &p4.PaymentACK{}, nil
			},
			req: p4.Payment{
				SPVEnvelope: &spv.Envelope{
					RawTx: "01000000000000000000",
					TxID:  "d21633ba23f70118185227be58a63527675641ad37967e2aa461559f577aec43",
				},
				MerchantData: p4.Merchant{
					ExtendedData: map[string]interface{}{"paymentReference": "omgwow"},
				},
			},
			args: p4.PaymentCreateArgs{
				PaymentID: "abc123",
			},
		},
		"invalid args errors": {
			paymentCreateFn: func(context.Context, p4.PaymentCreateArgs, p4.Payment) (*p4.PaymentACK, error) {
				return &p4.PaymentACK{}, nil
			},
			args: p4.PaymentCreateArgs{},
			req: p4.Payment{
				SPVEnvelope: &spv.Envelope{
					RawTx: "01000000000000000000",
					TxID:  "d21633ba23f70118185227be58a63527675641ad37967e2aa461559f577aec43",
				},
				MerchantData: p4.Merchant{
					ExtendedData: map[string]interface{}{"paymentReference": "omgwow"},
				},
			},
			expErr: errors.New("[paymentID: value cannot be empty]"),
		},
		"missing raw tx errors": {
			paymentCreateFn: func(context.Context, p4.PaymentCreateArgs, p4.Payment) (*p4.PaymentACK, error) {
				return &p4.PaymentACK{}, nil
			},
			args: p4.PaymentCreateArgs{
				PaymentID: "abc123",
			},
			req: p4.Payment{
				MerchantData: p4.Merchant{
					ExtendedData: map[string]interface{}{"paymentReference": "omgwow"},
				},
			},
			expErr: errors.New("[spvEnvelope/rawTx: either an SPVEnvelope or a rawTX are required]"),
		},
		"error on payment create is handled": {
			paymentCreateFn: func(context.Context, p4.PaymentCreateArgs, p4.Payment) (*p4.PaymentACK, error) {
				return nil, errors.New("lol oh boi")
			},
			args: p4.PaymentCreateArgs{
				PaymentID: "abc123",
			},
			req: p4.Payment{
				SPVEnvelope: &spv.Envelope{
					RawTx: "01000000000000000000",
					TxID:  "d21633ba23f70118185227be58a63527675641ad37967e2aa461559f577aec43",
				},
				MerchantData: p4.Merchant{
					ExtendedData: map[string]interface{}{"paymentReference": "omgwow"},
				},
			},
			expErr: errors.New("lol oh boi"),
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			svc := service.NewPayment(
				log.Noop{},
				&p4mocks.PaymentWriterMock{
					PaymentCreateFunc: test.paymentCreateFn,
				})

			_, err := svc.PaymentCreate(context.TODO(), test.args, test.req)
			if test.expErr != nil {
				assert.Error(t, err)
				assert.EqualError(t, err, test.expErr.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
