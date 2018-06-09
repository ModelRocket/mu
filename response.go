/*************************************************************************
 * MIT License
 * Copyright (c) 2018 Model Rocket
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

package mu

import (
	"encoding/base64"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

type (
	// Response defines a struct that conforms to the API Gateway Labmbda Response
	Response struct {
		events.APIGatewayProxyResponse
		header        http.Header
		headerWritten bool
		body          []byte
	}
)

// NewResponse returns a new response object
func NewResponse() *Response {
	return &Response{
		APIGatewayProxyResponse: events.APIGatewayProxyResponse{
			StatusCode:      200,
			Headers:         make(map[string]string),
			IsBase64Encoded: false,
		},
		header:        make(http.Header),
		body:          make([]byte, 0),
		headerWritten: false,
	}
}

// Header implements the http.ResponseWriter.Header() method
func (r *Response) Header() http.Header {
	return r.header
}

// Write implements the http.ResponseWriter.Write() method
func (r *Response) Write(body []byte) (int, error) {
	r.body = append(r.body, body...)

	if !r.headerWritten {
		r.WriteHeader(http.StatusOK)
	}

	if r.IsBase64Encoded == true {
		r.Body = base64.StdEncoding.EncodeToString(r.body)
	} else {
		r.Body = string(r.body)
	}

	return len(body), nil
}

// WriteHeader implements the http.ResponseWriter.WriteHeader() method
func (r *Response) WriteHeader(statusCode int) {
	r.StatusCode = statusCode

	ct := r.header.Get("Content-Type")
	if ct == "" {
		ct = http.DetectContentType([]byte(r.Body))
		r.header.Set("Content-Type", ct)
	}

	if r.header.Get("content-encoding") != "" || ct == "application/octet-stream" {
		r.IsBase64Encoded = true
	}

	for k, a := range r.header {
		for _, v := range a {
			r.Headers[k] = v
		}
	}

	r.headerWritten = true
}
