package web

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/nch-bowstave/paymail"
	"github.com/nch-bowstave/paymail/mocks"
	"github.com/nch-bowstave/paymail/service"
	"github.com/stretchr/testify/assert"
)

func TestVerify(t *testing.T) {
	e := echo.New()
	tests := map[string]struct {
		mockAccFunc func(ctx context.Context, handle paymail.Handle) (*paymail.PublicAccount, error)
		handle      string
		pubkey      string
		exp         *paymail.Verification
		code        int
	}{
		"should return 404 for non-existant paymail": {
			mockAccFunc: func(ctx context.Context, hanlde paymail.Handle) (*paymail.PublicAccount, error) {
				return nil, nil
			},
			code: http.StatusNotFound,
		},
		"should return 200 for paymail address but false for pubkey": {
			mockAccFunc: func(ctx context.Context, handle paymail.Handle) (*paymail.PublicAccount, error) {
				return &paymail.PublicAccount{
					Alias:     "bob",
					Handle:    "bob@somedomain.com",
					PublicKey: "abc123",
				}, nil
			},
			code:   http.StatusOK,
			handle: "bob@somedomain.com",
			pubkey: "abc1233",
			exp: &paymail.Verification{
				Handle: "bob@somedomain.com",
				Pubkey: "abc1233",
				Match:  false,
			},
		},
		"should return 200 for paymail address and true for pubkey": {
			mockAccFunc: func(ctx context.Context, handle paymail.Handle) (*paymail.PublicAccount, error) {
				return &paymail.PublicAccount{
					Alias:     "bob",
					Handle:    "bob@somedomain.com",
					PublicKey: "abc123",
				}, nil
			},
			code:   http.StatusOK,
			handle: "bob@somedomain.com",
			pubkey: "abc123",
			exp: &paymail.Verification{
				Handle: "bob@somedomain.com",
				Pubkey: "abc123",
				Match:  true,
			},
		},
	}

	for name, test := range tests {

		svc := service.NewPaymailService(&mocks.AccountReaderWriterMock{
			AccountFunc: test.mockAccFunc,
			CreateFunc:  func(ctx context.Context, account paymail.Account) error { return nil },
		}, "somedomain.com")
		a := NewBsvAlias(svc)
		a.RegisterRoutes(e.Group(""))

		t.Run(name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, routeVerify, nil)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			ctx := e.NewContext(req, rec)

			ctx.SetParamNames("handle", "pubkey")
			ctx.SetParamValues(test.handle, test.pubkey)

			if assert.NoError(t, a.Verify(ctx)) {
				assert.Equal(t, test.code, rec.Code)

				var verified *paymail.Verification
				json.Unmarshal(rec.Body.Bytes(), &verified)

				assert.Equal(t, test.code, rec.Code)
				assert.Equal(t, test.exp, verified)
			}
		})
	}
}

func TestPKI(t *testing.T) {
	e := echo.New()
	tests := map[string]struct {
		mockAccFunc func(ctx context.Context, handle paymail.Handle) (*paymail.PublicAccount, error)
		handle      string
		exp         *paymail.PKI
		code        int
	}{
		"should return 404 for missing handle": {
			mockAccFunc: func(ctx context.Context, handle paymail.Handle) (*paymail.PublicAccount, error) {
				return nil, nil
			},
			code:   http.StatusNotFound,
			handle: "bob@somedomain.com",
		},
		"should return 200 and pki for valid handle": {
			mockAccFunc: func(ctx context.Context, handle paymail.Handle) (*paymail.PublicAccount, error) {
				return &paymail.PublicAccount{
					Alias:     "bob",
					Handle:    "bob@somedomain.com",
					PublicKey: "abc123",
				}, nil
			},
			code:   http.StatusOK,
			handle: "bob@somedomain.com",
			exp: &paymail.PKI{
				BsvAlias:  "1.0",
				Handle:    "bob@somedomain.com",
				PublicKey: "abc123",
			},
		},
	}

	for name, test := range tests {
		svc := service.NewPaymailService(&mocks.AccountReaderWriterMock{
			AccountFunc: test.mockAccFunc,
			CreateFunc:  func(ctx context.Context, account paymail.Account) error { return nil },
		}, "somedomain.com")
		a := NewBsvAlias(svc)
		a.RegisterRoutes(e.Group(""))

		t.Run(name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, routePki, nil)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			ctx := e.NewContext(req, rec)

			ctx.SetParamNames("handle")
			ctx.SetParamValues(test.handle)

			if assert.NoError(t, a.PKI(ctx)) {
				assert.Equal(t, test.code, rec.Code)

				var pki *paymail.PKI
				json.Unmarshal(rec.Body.Bytes(), &pki)

				assert.Equal(t, test.code, rec.Code)
				assert.Equal(t, test.exp, pki)
			}
		})
	}
}

func TestPublicProfile(t *testing.T) {
	e := echo.New()
	tests := map[string]struct {
		mockAccFunc func(ctx context.Context, handle paymail.Handle) (*paymail.PublicAccount, error)
		handle      string
		exp         *paymail.PublicProfile
		code        int
	}{
		"should return 404 for missing handle": {
			mockAccFunc: func(ctx context.Context, handle paymail.Handle) (*paymail.PublicAccount, error) {
				return nil, nil
			},
			code:   http.StatusNotFound,
			handle: "bob@somedomain.com",
		},
		"should return 200 and profile for valid handle": {
			mockAccFunc: func(ctx context.Context, handle paymail.Handle) (*paymail.PublicAccount, error) {
				return &paymail.PublicAccount{
					AvatarUrl: "https://somedomain.com/avatar",
					Name:      "Bob Bobson",
				}, nil
			},
			code:   http.StatusOK,
			handle: "bob@somedomain.com",
			exp: &paymail.PublicProfile{
				AvatarUrl: "https://somedomain.com/avatar",
				Name:      "Bob Bobson",
			},
		},
	}

	for name, test := range tests {
		svc := service.NewPaymailService(&mocks.AccountReaderWriterMock{
			AccountFunc: test.mockAccFunc,
			CreateFunc:  func(ctx context.Context, account paymail.Account) error { return nil },
		}, "somedomain.com")
		a := NewBsvAlias(svc)
		a.RegisterRoutes(e.Group(""))

		t.Run(name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, routePublicProfile, nil)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			ctx := e.NewContext(req, rec)

			ctx.SetParamNames("handle")
			ctx.SetParamValues(test.handle)

			if assert.NoError(t, a.PublicProfile(ctx)) {
				assert.Equal(t, test.code, rec.Code)

				var profile *paymail.PublicProfile
				json.Unmarshal(rec.Body.Bytes(), &profile)

				assert.Equal(t, test.code, rec.Code)
				assert.Equal(t, test.exp, profile)
			}
		})
	}
}
