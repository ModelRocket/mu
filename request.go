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
	"bytes"
	"encoding/base64"
	"net/http"
	"net/url"

	"github.com/aws/aws-lambda-go/events"
)

// NewRequest returns a new *http.Request for the APIGatewayProxyRequest event
func NewRequest(event events.APIGatewayProxyRequest) (request *http.Request, err error) {
	// remove the resource root
	path, _ := url.Parse(event.Path)
	q := path.Query()
	for k, v := range event.QueryStringParameters {
		q.Set(k, v)
	}
	path.RawQuery = q.Encode()

	body := []byte(event.Body)

	if event.IsBase64Encoded {
		body, _ = base64.StdEncoding.DecodeString(event.Body)
	}

	request, err = http.NewRequest(event.HTTPMethod, path.String(), bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	request.RemoteAddr = event.RequestContext.Identity.SourceIP

	// Add the headers to the request
	for k, v := range event.Headers {
		request.Header.Set(k, v)
	}

	return request, nil
}
