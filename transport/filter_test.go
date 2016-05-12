// Copyright (c) 2016 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package transport_test

import (
	"bytes"
	"io/ioutil"
	"testing"
	"time"

	"github.com/yarpc/yarpc-go/encoding/raw"
	"github.com/yarpc/yarpc-go/transport"
	"github.com/yarpc/yarpc-go/transport/transporttest"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"
)

func TestNopFilter(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	o := transporttest.NewMockOutbound(mockCtrl)
	wrappedO := transport.ApplyFilter(o, transport.NopFilter)

	ctx, _ := context.WithTimeout(context.Background(), time.Second)
	req := &transport.Request{
		Caller:    "somecaller",
		Service:   "someservice",
		Encoding:  raw.Encoding,
		TTL:       time.Second,
		Procedure: "hello",
		Body:      bytes.NewReader([]byte{1, 2, 3}),
	}

	res := &transport.Response{Body: ioutil.NopCloser(bytes.NewReader([]byte{4, 5, 6}))}
	o.EXPECT().Call(ctx, req).Return(res, nil)

	got, err := wrappedO.Call(ctx, req)
	if assert.NoError(t, err) {
		assert.Equal(t, res, got)
	}
}
