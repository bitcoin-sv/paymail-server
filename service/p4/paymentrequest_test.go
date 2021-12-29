package service_test

// import (
// 	"context"
// 	"errors"
// 	"testing"
// 	"time"

// 	"github.com/libsv/go-p4"
// 	p4mocks "github.com/libsv/go-p4/mocks"
// 	"github.com/nch-bowstave/paymail/config"
// 	"github.com/nch-bowstave/paymail/service"
// 	"github.com/stretchr/testify/assert"
// )

// func TestPaymentRequest_PaymentRequest(t *testing.T) {
// 	created := time.Now()
// 	expired := created.Add(time.Hour * 24)
// 	tests := map[string]struct {
// 		destinationsFunc func(context.Context, p4.PaymentRequestArgs) (*p4.Destinations, error)
// 		ownerFunc        func(context.Context) (*p4.Merchant, error)
// 		config           *config.Server
// 		args             p4.PaymentRequestArgs
// 		expResp          *p4.PaymentRequest
// 		expErr           error
// 	}{
// 		"successful request": {
// 			args: p4.PaymentRequestArgs{
// 				PaymentID: "abc123",
// 			},
// 			config: &config.Server{
// 				FQDN: "iamsotest",
// 			},
// 			destinationsFunc: func(context.Context, p4.PaymentRequestArgs) (*p4.Destinations, error) {
// 				return &p4.Destinations{
// 					SPVRequired: false,
// 					CreatedAt:   created,
// 					ExpiresAt:   expired,
// 					Outputs: []p4.Output{{
// 						Amount: 500,
// 						Script: "abc123",
// 					}},
// 				}, nil
// 			},
// 			ownerFunc: func(context.Context) (*p4.Merchant, error) {
// 				return &p4.Merchant{
// 					ExtendedData: map[string]interface{}{},
// 				}, nil
// 			},
// 			expResp: &p4.PaymentRequest{
// 				SPVRequired:         false,
// 				CreationTimestamp:   created,
// 				ExpirationTimestamp: expired,
// 				Destinations: p4.PaymentDestinations{
// 					Outputs: []p4.Output{{
// 						Amount: 500,
// 						Script: "abc123",
// 					}},
// 				},
// 				PaymentURL: "http://iamsotest/api/v1/payment/abc123",
// 				Memo:       "invoice abc123",
// 				MerchantData: &p4.Merchant{
// 					ExtendedData: map[string]interface{}{"paymentReference": "abc123"},
// 				},
// 			},
// 		},
// 		"successful request with nil extended data from owner reader": {
// 			args: p4.PaymentRequestArgs{
// 				PaymentID: "abc123",
// 			},
// 			config: &config.Server{
// 				FQDN: "iamsotest",
// 			},
// 			destinationsFunc: func(context.Context, p4.PaymentRequestArgs) (*p4.Destinations, error) {
// 				return &p4.Destinations{
// 					SPVRequired: false,
// 					CreatedAt:   created,
// 					ExpiresAt:   expired,
// 					Outputs: []p4.Output{{
// 						Amount: 500,
// 						Script: "abc123",
// 					}},
// 				}, nil
// 			},
// 			ownerFunc: func(context.Context) (*p4.Merchant, error) {
// 				return &p4.Merchant{
// 					ExtendedData: nil,
// 				}, nil
// 			},
// 			expResp: &p4.PaymentRequest{
// 				SPVRequired:         false,
// 				CreationTimestamp:   created,
// 				ExpirationTimestamp: expired,
// 				Destinations: p4.PaymentDestinations{
// 					Outputs: []p4.Output{{
// 						Amount: 500,
// 						Script: "abc123",
// 					}},
// 				},
// 				PaymentURL: "http://iamsotest/api/v1/payment/abc123",
// 				Memo:       "invoice abc123",
// 				MerchantData: &p4.Merchant{
// 					ExtendedData: map[string]interface{}{"paymentReference": "abc123"},
// 				},
// 			},
// 		},
// 		"config fqdn is reflected in payment url": {
// 			args: p4.PaymentRequestArgs{
// 				PaymentID: "abc123",
// 			},
// 			config: &config.Server{
// 				FQDN: "iamsodifferent",
// 			},
// 			destinationsFunc: func(context.Context, p4.PaymentRequestArgs) (*p4.Destinations, error) {
// 				return &p4.Destinations{
// 					SPVRequired: false,
// 					CreatedAt:   created,
// 					ExpiresAt:   expired,
// 					Outputs: []p4.Output{{
// 						Amount: 500,
// 						Script: "abc123",
// 					}},
// 				}, nil
// 			},
// 			ownerFunc: func(context.Context) (*p4.Merchant, error) {
// 				return &p4.Merchant{
// 					ExtendedData: map[string]interface{}{},
// 				}, nil
// 			},
// 			expResp: &p4.PaymentRequest{
// 				SPVRequired:         false,
// 				CreationTimestamp:   created,
// 				ExpirationTimestamp: expired,
// 				Destinations: p4.PaymentDestinations{
// 					Outputs: []p4.Output{{
// 						Amount: 500,
// 						Script: "abc123",
// 					}},
// 				},
// 				PaymentURL: "http://iamsodifferent/api/v1/payment/abc123",
// 				Memo:       "invoice abc123",
// 				MerchantData: &p4.Merchant{
// 					ExtendedData: map[string]interface{}{"paymentReference": "abc123"},
// 				},
// 			},
// 		},
// 		"paymentID is reflected in the payment url, the memo, and the extended data": {
// 			args: p4.PaymentRequestArgs{
// 				PaymentID: "456def",
// 			},
// 			config: &config.Server{
// 				FQDN: "iamsotest",
// 			},
// 			destinationsFunc: func(context.Context, p4.PaymentRequestArgs) (*p4.Destinations, error) {
// 				return &p4.Destinations{
// 					SPVRequired: false,
// 					CreatedAt:   created,
// 					ExpiresAt:   expired,
// 					Outputs: []p4.Output{{
// 						Amount: 500,
// 						Script: "abc123",
// 					}},
// 				}, nil
// 			},
// 			ownerFunc: func(context.Context) (*p4.Merchant, error) {
// 				return &p4.Merchant{
// 					ExtendedData: map[string]interface{}{},
// 				}, nil
// 			},
// 			expResp: &p4.PaymentRequest{
// 				SPVRequired:         false,
// 				CreationTimestamp:   created,
// 				ExpirationTimestamp: expired,
// 				Destinations: p4.PaymentDestinations{
// 					Outputs: []p4.Output{{
// 						Amount: 500,
// 						Script: "abc123",
// 					}},
// 				},
// 				PaymentURL: "http://iamsotest/api/v1/payment/456def",
// 				Memo:       "invoice 456def",
// 				MerchantData: &p4.Merchant{
// 					ExtendedData: map[string]interface{}{"paymentReference": "456def"},
// 				},
// 			},
// 		},
// 		"invalid args rejected": {
// 			config: &config.Server{
// 				FQDN: "iamsotest",
// 			},
// 			destinationsFunc: func(context.Context, p4.PaymentRequestArgs) (*p4.Destinations, error) {
// 				return &p4.Destinations{
// 					SPVRequired: false,
// 					CreatedAt:   created,
// 					ExpiresAt:   expired,
// 					Outputs: []p4.Output{{
// 						Amount: 500,
// 						Script: "abc123",
// 					}},
// 				}, nil
// 			},
// 			ownerFunc: func(context.Context) (*p4.Merchant, error) {
// 				return &p4.Merchant{
// 					ExtendedData: map[string]interface{}{},
// 				}, nil
// 			},
// 			expErr: errors.New("[paymentID: value cannot be empty]"),
// 		},
// 		"destination reader error handled and reported": {
// 			args: p4.PaymentRequestArgs{
// 				PaymentID: "abc123",
// 			},
// 			config: &config.Server{
// 				FQDN: "iamsotest",
// 			},
// 			destinationsFunc: func(context.Context, p4.PaymentRequestArgs) (*p4.Destinations, error) {
// 				return nil, errors.New("oh boi")
// 			},
// 			ownerFunc: func(context.Context) (*p4.Merchant, error) {
// 				return &p4.Merchant{
// 					ExtendedData: map[string]interface{}{},
// 				}, nil
// 			},
// 			expErr: errors.New("failed to get destinations for paymentID abc123: oh boi"),
// 		},
// 		"owner reader error handled and reported": {
// 			args: p4.PaymentRequestArgs{
// 				PaymentID: "abc123",
// 			},
// 			config: &config.Server{
// 				FQDN: "iamsotest",
// 			},
// 			destinationsFunc: func(context.Context, p4.PaymentRequestArgs) (*p4.Destinations, error) {
// 				return &p4.Destinations{
// 					SPVRequired: false,
// 					CreatedAt:   created,
// 					ExpiresAt:   expired,
// 					Outputs: []p4.Output{{
// 						Amount: 500,
// 						Script: "abc123",
// 					}},
// 				}, nil
// 			},
// 			ownerFunc: func(context.Context) (*p4.Merchant, error) {
// 				return nil, errors.New("yikes")
// 			},
// 			expErr: errors.New("failed to read merchant data when constructing payment request: yikes"),
// 		},
// 	}

// 	for name, test := range tests {
// 		t.Run(name, func(t *testing.T) {
// 			svc := service.NewPaymentRequest(test.config, &p4mocks.DestinationReaderMock{
// 				DestinationsFunc: test.destinationsFunc,
// 			}, &p4mocks.MerchantReaderMock{
// 				OwnerFunc: test.ownerFunc,
// 			})

// 			resp, err := svc.PaymentRequest(context.TODO(), test.args)
// 			if test.expErr != nil {
// 				assert.Error(t, err)
// 				assert.EqualError(t, err, test.expErr.Error())
// 				return
// 			}

// 			assert.NoError(t, err)
// 			assert.NotNil(t, resp)
// 			assert.Equal(t, *test.expResp, *resp)
// 		})
// 	}
// }
