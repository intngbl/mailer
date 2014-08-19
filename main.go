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
	"bytes"
	"fmt"
	"log"
	"net/http"
	"net/smtp"
	"os"
	"strconv"
	"time"
	"upper.io/queue"
)

const (
	defaultHost  = `127.0.0.1`
	defaultPort  = 25
	defaultFrom  = `someone@example.org`
	envHost      = `SMTP_HOST`
	envPort      = `SMTP_PORT`
	envUsername  = `SMTP_USERNAME`
	envPassword  = `SMTP_PASSWORD`
	envQueueName = `MAIL_QUEUE`
)

var sleepTime = time.Millisecond * 500

type Mailer struct {
	Host     string
	Port     int
	Username string
	Password string
	From     string
	queue    *queue.Queue
}

type Message struct {
	To      []string
	Subject string
	Content []byte
	Headers http.Header
}

// Creates a new mailer and a message queue.
func NewMailer() (self *Mailer, err error) {
	self = &Mailer{}

	self.Host = os.Getenv(envHost)
	self.Port, _ = strconv.Atoi(os.Getenv(envPort))
	self.Username = os.Getenv(envUsername)
	self.Password = os.Getenv(envPassword)

	if self.Host == "" {
		self.Host = defaultHost
	}
	if self.Port == 0 {
		self.Port = defaultPort
	}

	queueName := os.Getenv(envQueueName)

	if queueName == "" {
		queueName = "go-mailer"
	}

	self.queue, err = queue.NewQueue(queueName)

	if err != nil {
		return nil, err
	}

	go self.dequeue()

	return self, nil
}

func (self *Mailer) dequeue() {
	// TODO: use concurrency to send more messages.
	var message Message
	var err error
	for {
		err = self.queue.Shift(&message)

		if err == nil {
			err = self.Send(message)
			if err != nil {
				log.Printf("Error sending mail: %s\n", err.Error())
			}
		}

		time.Sleep(sleepTime)
	}
}

// Adds a mail to the sending queue.
func (self *Mailer) Enqueue(message Message) (err error) {
	err = self.queue.Push(&message)
	return err
}

// Attempts to send and e-mail.
func (self *Mailer) Send(message Message) (err error) {
	auth := smtp.PlainAuth(
		"",
		self.Username,
		self.Password,
		self.Host,
	)

	var from string
	from = self.From
	if from == "" {
		from = defaultFrom
	}

	var content []byte

	// TODO: Add support for headers besides Subject, maybe we could use
	// http.Header?
	content = []byte(fmt.Sprintf("Subject: %s\n%s\r\n\r\n", message.Subject, message.getHeaders()))
	content = append(content, message.Content...)

	err = smtp.SendMail(
		fmt.Sprintf("%s:%d", self.Host, self.Port),
		auth,
		from,
		message.To,
		content,
	)

	return err
}

func (self *Message) getHeaders() []byte {
	var buf bytes.Buffer
	self.Headers.Write(&buf)

	return buf.Bytes()
}
