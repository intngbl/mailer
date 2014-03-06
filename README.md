# github.com/intngbl/mailer

A dead-simple mail queue based on redis lists.

## Install

```
go get github.com/intngbl/mailer
```

## Code example

```go
m, err := mailer.NewMailer()

if err != nil {
  log.Fatalf("Error: %s\n", err.Error())
}

err = m.Send(mailer.Message{
  Subject: "Hi!",
  To:      []string{"test@example.com"},
  Content: []byte("It has been a while!"),
})
```

## Configuring mailer

Configuration settings should be part of the environment:

```
SMTP_HOST="smtp.example.org" ./my-go-program
```

This is a list of the enviroment variables the `mailer` package expects:

* `SMTP_HOST` IP or FQDN of the SMTP server. Defaults to `127.0.0.1`.
* `SMTP_PORT` Defaults to 25
* `SMTP_USERNAME` Username, required if the SMTP server requires authentication.
* `SMTP_PASSWORD` Password, required if the SMTP server requires authentication.
* `MAIL_QUEUE` Queue name. Defaults to `go-mailer`.

## License

> Copyright (c) 2014, Intangible Investments S.A.P.I. de C.V.
>
> Permission is hereby granted, free of charge, to any person obtaining a copy
> of this software and associated documentation files (the "Software"), to deal
> in the Software without restriction, including without limitation the rights
> to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
> copies of the Software, and to permit persons to whom the Software is
> furnished to do so, subject to the following conditions:
>
> The above copyright notice and this permission notice shall be included in
> all copies or substantial portions of the Software.
>
> THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
> IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
> FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
> AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
> LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
> OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
> THE SOFTWARE.
