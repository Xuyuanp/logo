/*
 * Copyright 2015 Xuyuan Pang
 * Author: Xuyuan Pang
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package logo

import (
	"crypto/tls"
	"errors"
	"io"
	"net"
	"net/smtp"
	"strings"
	"sync"
)

// SMTPWriter struct
type SMTPWriter struct {
	cli      *smtp.Client
	subject  string
	addr     string
	username string
	password string
	to       []string
	buf      []byte
	mu       sync.Mutex
}

// NewSMTPWriter create a new SMTPWriter instance.
func NewSMTPWriter(addr, username, password, subject string, to ...string) *SMTPWriter {
	sw := &SMTPWriter{
		addr:     addr,
		username: username,
		password: password,
		subject:  subject,
		to:       to,
	}
	return sw
}

// Connect connects to SMTP server, init the client and init buffer.
func (sw *SMTPWriter) Connect() error {
	// init client
	conn, err := tls.Dial("tcp", sw.addr, nil)
	if err != nil {
		return err
	}
	host, _, err := net.SplitHostPort(sw.addr)
	if err != nil {
		return err
	}
	c, err := smtp.NewClient(conn, host)
	if err != nil {
		return err
	}
	auth := smtp.PlainAuth("", sw.username, sw.password, host)
	if err := c.Auth(auth); err != nil {
		return err
	}
	if err := c.Mail(sw.username); err != nil {
		return err
	}
	sw.cli = c

	// init SMTP header
	header := map[string]string{
		"From":         sw.username,
		"To":           strings.Join(sw.to, ";"),
		"Subject":      sw.subject,
		"MIME-Version": "1.0",
		"Content-Type": `text/plain; charset="utf-8"`,
	}
	var buf []byte
	for k, v := range header {
		buf = append(buf, k...)
		buf = append(buf, ':', ' ')
		buf = append(buf, v...)
		buf = append(buf, '\r', '\n')
	}
	buf = append(buf, "\r\n\r\n"...)
	sw.buf = buf
	return nil
}

func (sw *SMTPWriter) Write(d []byte) (n int, err error) {
	sw.mu.Lock()
	if sw.cli == nil {
		sw.mu.Unlock()
		return -1, errors.New("client not init")
	}
	body := sw.buf[:]
	body = append(body, d...)

	for _, t := range sw.to {
		if err = sw.cli.Rcpt(t); err != nil {
			sw.mu.Unlock()
			return
		}
	}
	var wc io.WriteCloser
	if wc, err = sw.cli.Data(); err == nil {
		n, err = wc.Write(body)
		wc.Close()
		sw.mu.Unlock()
		return
	}
	sw.mu.Unlock()
	return
}

// Close quites SMTP client.
func (sw *SMTPWriter) Close() (err error) {
	sw.mu.Lock()
	if sw.cli != nil {
		err = sw.cli.Quit()
		sw.cli = nil
	}
	sw.mu.Unlock()
	return
}
