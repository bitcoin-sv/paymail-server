package payd_test

import (
	"context"
	"encoding/json"
	"errors"
	"testing"
	"time"

	"github.com/libsv/go-bc/spv"
	"github.com/libsv/go-bt/v2"
	"github.com/libsv/go-p4"
	"github.com/libsv/p4-server/data/payd"
	"github.com/libsv/p4-server/data/payd/models"
	"github.com/libsv/p4-server/mocks"
	"github.com/nch-bowstave/paymail/config"
	"github.com/stretchr/testify/assert"
)

func TestPayd_PaymentCreate(t *testing.T) {
	tests := map[string]struct {
		doFunc func(context.Context, string, string, int, interface{}, interface{}) error
		args   p4.PaymentCreateArgs
		req    p4.Payment
		cfg    *config.PayD
		expURL string
		expReq models.PayDPaymentRequest
		expErr error
	}{
		"successful payment created": {
			doFunc: func(context.Context, string, string, int, interface{}, interface{}) error {
				return nil
			},
			args: p4.PaymentCreateArgs{
				PaymentID: "qwe123",
			},
			req: p4.Payment{
				RawTX:       func() *string { s := "rawrawraw"; return &s }(),
				SPVEnvelope: &spv.Envelope{},
				ProofCallbacks: map[string]p4.ProofCallback{
					"abc.com": {Token: "mYtOkEn"},
				},
			},
			expReq: models.PayDPaymentRequest{
				RawTX:       func() *string { s := "rawrawraw"; return &s }(),
				SPVEnvelope: &spv.Envelope{},
				ProofCallbacks: map[string]p4.ProofCallback{
					"abc.com": {Token: "mYtOkEn"},
				},
			},
			cfg: &config.PayD{
				Host: "paydhost",
				Port: ":8080",
			},
			expURL: "http://paydhost:8080/api/v1/payments/qwe123",
		},
		"successful https payment created": {
			doFunc: func(context.Context, string, string, int, interface{}, interface{}) error {
				return nil
			},
			args: p4.PaymentCreateArgs{
				PaymentID: "qwe123",
			},
			req: p4.Payment{
				RawTX:       func() *string { s := "rawrawraw"; return &s }(),
				SPVEnvelope: &spv.Envelope{},
				ProofCallbacks: map[string]p4.ProofCallback{
					"abc.com": {Token: "mYtOkEn"},
				},
			},
			expReq: models.PayDPaymentRequest{
				RawTX:       func() *string { s := "rawrawraw"; return &s }(),
				SPVEnvelope: &spv.Envelope{},
				ProofCallbacks: map[string]p4.ProofCallback{
					"abc.com": {Token: "mYtOkEn"},
				},
			},
			cfg: &config.PayD{
				Host:   "securepaydhost",
				Port:   ":8081",
				Secure: true,
			},
			expURL: "https://securepaydhost:8081/api/v1/payments/qwe123",
		},
		"error is handled and returned": {
			doFunc: func(context.Context, string, string, int, interface{}, interface{}) error {
				return errors.New("i tried so hard")
			},
			args: p4.PaymentCreateArgs{
				PaymentID: "qwe123",
			},
			req: p4.Payment{
				RawTX:       func() *string { s := "rawrawraw"; return &s }(),
				SPVEnvelope: &spv.Envelope{},
				ProofCallbacks: map[string]p4.ProofCallback{
					"abc.com": {Token: "mYtOkEn"},
				},
			},
			expReq: models.PayDPaymentRequest{
				RawTX:       func() *string { s := "rawrawraw"; return &s }(),
				SPVEnvelope: &spv.Envelope{},
				ProofCallbacks: map[string]p4.ProofCallback{
					"abc.com": {Token: "mYtOkEn"},
				},
			},
			cfg: &config.PayD{
				Host:   "securepaydhost",
				Port:   ":8081",
				Secure: true,
			},
			expURL: "https://securepaydhost:8081/api/v1/payments/qwe123",
			expErr: errors.New("i tried so hard"),
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			pd := payd.NewPayD(test.cfg, &mocks.HTTPClientMock{
				DoFunc: func(ctx context.Context, method string, url string, statusCode int, req, out interface{}) error {
					assert.Equal(t, test.expURL, url)
					assert.Equal(t, test.expReq, req)
					return test.doFunc(ctx, method, url, statusCode, req, out)
				},
			})
			_, err := pd.PaymentCreate(context.Background(), test.args, test.req)
			if test.expErr != nil {
				assert.Error(t, err)
				assert.EqualError(t, err, test.expErr.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestPayd_Owner(t *testing.T) {
	tests := map[string]struct {
		doFunc func(context.Context, string, string, int, interface{}, interface{}) error
		cfg    *config.PayD
		expURL string
		expErr error
	}{
		"successful owner request": {
			doFunc: func(ctx context.Context, method string, url string, statusCode int, req, out interface{}) error {
				return nil
			},
			cfg: &config.PayD{
				Host: "paydowner",
				Port: ":1122",
			},
			expURL: "http://paydowner:1122/api/v1/owner",
		},
		"successful https owner request": {
			doFunc: func(ctx context.Context, method string, url string, statusCode int, req, out interface{}) error {
				return nil
			},
			cfg: &config.PayD{
				Host:   "securepaydowner",
				Port:   ":2122",
				Secure: true,
			},
			expURL: "https://securepaydowner:2122/api/v1/owner",
		},
		"error is reported": {
			doFunc: func(ctx context.Context, method string, url string, statusCode int, req, out interface{}) error {
				return errors.New("oh no")
			},
			cfg: &config.PayD{
				Host:   "securepaydowner",
				Port:   ":2122",
				Secure: true,
			},
			expURL: "https://securepaydowner:2122/api/v1/owner",
			expErr: errors.New("oh no"),
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			pd := payd.NewPayD(test.cfg, &mocks.HTTPClientMock{
				DoFunc: func(ctx context.Context, method string, url string, statusCode int, req, out interface{}) error {
					assert.Equal(t, test.expURL, url)
					return test.doFunc(ctx, method, url, statusCode, req, out)
				},
			})
			_, err := pd.Owner(context.Background())
			if test.expErr != nil {
				assert.Error(t, err)
				assert.EqualError(t, err, test.expErr.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestPayd_Destinations(t *testing.T) {
	tests := map[string]struct {
		doFunc   func(ctx context.Context, method string, url string, statusCode int, req, out interface{}) error
		args     p4.PaymentRequestArgs
		cfg      *config.PayD
		expURL   string
		expDests *p4.Destinations
		expErr   error
	}{
		"successful destination request": {
			doFunc: func(ctx context.Context, method string, url string, statusCode int, req, out interface{}) error {
				return json.Unmarshal([]byte(`{
					"spvRequired": true,
					"network": "mainnet",
					"outputs": [{
						"script": "abc123",
						"satoshis": 100
					}, {
						"script": "def456",
						"satoshis": 400
					}],
					"fees": {
						"data": {
							"miningFee": {
								"satoshis": 5,
								"bytes": 10
							},
							"relayFee": {
								"satoshis": 5,
								"bytes": 10
							}
						},
						"standard": {
							"miningFee": {
								"satoshis": 5,
								"bytes": 10
							},
							"relayFee": {
								"satoshis": 5,
								"bytes": 10
							}
						}
					},
					"createdAt": "2021-10-15T08:33:51.51229Z",
					"expiresAt": "2021-10-16T08:33:51.51229Z"
				}`), &out)
			},
			args: p4.PaymentRequestArgs{
				PaymentID: "qwe123",
			},
			cfg: &config.PayD{
				Host: "payddest",
				Port: ":445",
			},
			expURL: "http://payddest:445/api/v1/destinations/qwe123",
			expDests: &p4.Destinations{
				SPVRequired: true,
				Network:     "mainnet",
				Outputs: []p4.Output{{
					Script: "abc123",
					Amount: 100,
				}, {
					Script: "def456",
					Amount: 400,
				}},
				Fees:      bt.NewFeeQuote(),
				CreatedAt: func() time.Time { t, _ := time.Parse(time.RFC3339, "2021-10-15T08:33:51.51229Z"); return t }(),
				ExpiresAt: func() time.Time { t, _ := time.Parse(time.RFC3339, "2021-10-16T08:33:51.51229Z"); return t }(),
			},
		},
		"successful https destination request": {
			doFunc: func(ctx context.Context, method string, url string, statusCode int, req, out interface{}) error {
				return json.Unmarshal([]byte(`{
					"spvRequired": true,
					"network": "mainnet",
					"outputs": [{
						"script": "abc123",
						"satoshis": 100
					}, {
						"script": "def456",
						"satoshis": 400
					}],
					"fees": {
						"data": {
							"miningFee": {
								"satoshis": 5,
								"bytes": 10
							},
							"relayFee": {
								"satoshis": 5,
								"bytes": 10
							}
						},
						"standard": {
							"miningFee": {
								"satoshis": 5,
								"bytes": 10
							},
							"relayFee": {
								"satoshis": 5,
								"bytes": 10
							}
						}
					},
					"createdAt": "2021-10-15T08:33:51.51229Z",
					"expiresAt": "2021-10-16T08:33:51.51229Z"
				}`), &out)
			},
			args: p4.PaymentRequestArgs{
				PaymentID: "bwe123",
			},
			cfg: &config.PayD{
				Host:   "securepayddest",
				Port:   ":4445",
				Secure: true,
			},
			expURL: "https://securepayddest:4445/api/v1/destinations/bwe123",
			expDests: &p4.Destinations{
				SPVRequired: true,
				Network:     "mainnet",
				Outputs: []p4.Output{{
					Script: "abc123",
					Amount: 100,
				}, {
					Script: "def456",
					Amount: 400,
				}},
				Fees:      bt.NewFeeQuote(),
				CreatedAt: func() time.Time { t, _ := time.Parse(time.RFC3339, "2021-10-15T08:33:51.51229Z"); return t }(),
				ExpiresAt: func() time.Time { t, _ := time.Parse(time.RFC3339, "2021-10-16T08:33:51.51229Z"); return t }(),
			},
		},
		"error is handled": {
			doFunc: func(ctx context.Context, method string, url string, statusCode int, req, out interface{}) error {
				return errors.New("yikes")
			},
			args: p4.PaymentRequestArgs{
				PaymentID: "bwe123",
			},
			cfg: &config.PayD{
				Host:   "securepayddest",
				Port:   ":4445",
				Secure: true,
			},
			expURL: "https://securepayddest:4445/api/v1/destinations/bwe123",
			expErr: errors.New("yikes"),
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			pd := payd.NewPayD(test.cfg, &mocks.HTTPClientMock{
				DoFunc: func(ctx context.Context, method string, url string, statusCode int, req, out interface{}) error {
					assert.Equal(t, test.expURL, url)
					return test.doFunc(ctx, method, url, statusCode, req, out)
				},
			})
			dests, err := pd.Destinations(context.Background(), test.args)
			if test.expErr != nil {
				assert.Error(t, err)
				assert.EqualError(t, err, test.expErr.Error())
			} else {
				assert.NoError(t, err)
			}

			if test.expDests != nil {
				assert.NotNil(t, dests)
				assert.Equal(t, test.expDests.CreatedAt.String(), dests.CreatedAt.String())
				assert.Equal(t, test.expDests.ExpiresAt.String(), dests.ExpiresAt.String())

				ts := time.Now()
				dests.Fees.UpdateExpiry(ts)
				test.expDests.Fees.UpdateExpiry(ts)

				assert.Equal(t, *test.expDests, *dests)
			} else {
				assert.Nil(t, dests)
			}
		})
	}
}
