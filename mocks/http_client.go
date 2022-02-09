// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package mocks

import (
	"context"
	"github.com/nch-bowstave/paymail/data"
	"net/http"
	"sync"
)

// Ensure, that HTTPClientMock does implement data.HTTPClient.
// If this is not the case, regenerate this file with moq.
var _ data.HTTPClient = &HTTPClientMock{}

// HTTPClientMock is a mock implementation of data.HTTPClient.
//
// 	func TestSomethingThatUsesHTTPClient(t *testing.T) {
//
// 		// make and configure a mocked data.HTTPClient
// 		mockedHTTPClient := &HTTPClientMock{
// 			DoFunc: func(ctx context.Context, method string, endpoint string, expStatus int, hdrs http.Header, req interface{}, out interface{}) error {
// 				panic("mock out the Do method")
// 			},
// 		}
//
// 		// use mockedHTTPClient in code that requires data.HTTPClient
// 		// and then make assertions.
//
// 	}
type HTTPClientMock struct {
	// DoFunc mocks the Do method.
	DoFunc func(ctx context.Context, method string, endpoint string, expStatus int, hdrs http.Header, req interface{}, out interface{}) error

	// calls tracks calls to the methods.
	calls struct {
		// Do holds details about calls to the Do method.
		Do []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Method is the method argument value.
			Method string
			// Endpoint is the endpoint argument value.
			Endpoint string
			// ExpStatus is the expStatus argument value.
			ExpStatus int
			// Hdrs is the hdrs argument value.
			Hdrs http.Header
			// Req is the req argument value.
			Req interface{}
			// Out is the out argument value.
			Out interface{}
		}
	}
	lockDo sync.RWMutex
}

// Do calls DoFunc.
func (mock *HTTPClientMock) Do(ctx context.Context, method string, endpoint string, expStatus int, hdrs http.Header, req interface{}, out interface{}) error {
	if mock.DoFunc == nil {
		panic("HTTPClientMock.DoFunc: method is nil but HTTPClient.Do was just called")
	}
	callInfo := struct {
		Ctx       context.Context
		Method    string
		Endpoint  string
		ExpStatus int
		Hdrs      http.Header
		Req       interface{}
		Out       interface{}
	}{
		Ctx:       ctx,
		Method:    method,
		Endpoint:  endpoint,
		ExpStatus: expStatus,
		Hdrs:      hdrs,
		Req:       req,
		Out:       out,
	}
	mock.lockDo.Lock()
	mock.calls.Do = append(mock.calls.Do, callInfo)
	mock.lockDo.Unlock()
	return mock.DoFunc(ctx, method, endpoint, expStatus, hdrs, req, out)
}

// DoCalls gets all the calls that were made to Do.
// Check the length with:
//     len(mockedHTTPClient.DoCalls())
func (mock *HTTPClientMock) DoCalls() []struct {
	Ctx       context.Context
	Method    string
	Endpoint  string
	ExpStatus int
	Hdrs      http.Header
	Req       interface{}
	Out       interface{}
} {
	var calls []struct {
		Ctx       context.Context
		Method    string
		Endpoint  string
		ExpStatus int
		Hdrs      http.Header
		Req       interface{}
		Out       interface{}
	}
	mock.lockDo.RLock()
	calls = mock.calls.Do
	mock.lockDo.RUnlock()
	return calls
}
