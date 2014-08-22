/*
	Extremely simple redis-based e-mail queue.
	---

	Copyright (c) 2014, Intangible Investments S.A.P.I. de C.V.

	Permission is hereby granted, free of charge, to any person obtaining a copy
	of this software and associated documentation files (the "Software"), to deal
	in the Software without restriction, including without limitation the rights
	to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
	copies of the Software, and to permit persons to whom the Software is
	furnished to do so, subject to the following conditions:

	The above copyright notice and this permission notice shall be included in
	all copies or substantial portions of the Software.

	THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
	IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
	FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
	AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
	LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
	OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
	THE SOFTWARE.
*/
package mailer

import (
	"net/http"
	"testing"
	"time"
)

func TestSendMail(t *testing.T) {
	var err error
	var m *Mailer
	var headers http.Header

	m, err = NewMailer()

	if err != nil {
		t.Fatalf("Error: %s\n", err.Error())
	}

	err = m.Send(Message{
		Subject: "Normal message",
		To:      []string{"jose.carlos+spam@intangible.mx"},
		Content: []byte("E-mail content."),
	})

	if err != nil {
		t.Fatalf("Error: %s\n", err.Error())
	}

	err = m.Enqueue(Message{
		Subject: "Queued message",
		To:      []string{"jose.carlos+spam@intangible.mx"},
		Content: []byte("E-mail content."),
	})

	if err != nil {
		t.Fatalf("Error: %s\n", err.Error())
	}

	headers = make(http.Header)
	headers.Add(`MIME-version`, `1.0`)
	headers.Add(`Content-type`, `text/html; charset="UTF-8"`)

	err = m.Send(Message{
		Subject: "HTML message",
		To:      []string{"jose.carlos+spam@intangible.mx"},
		Content: []byte("<html><body><h1>Rich content</h1></body></html>"),
		Headers: headers,
	})

	if err != nil {
		t.Fatalf("Error: %s\n", err.Error())
	}

	// Waiting for queue to drain.
	time.Sleep(time.Second * 5)
}
