# logo [![GoDoc](https://godoc.org/github.com/Xuyuanp/logo?status.svg)](https://godoc.org/github.com/Xuyuanp/logo) 
log for go

## Getting Started

Install `logo` package:

`go get github.com/Xuyuanp/logo`

## Usage

```go
package main

import (
    "os"

    "github.com/Xuyuanp/logo"
)

func main() {
    // use default logo
    logo.Info("logo example")

    // new logo with stdout.
    tl := logo.New(logo.LevelDebug, os.Stdout, "", logo.LdefaultFlags)
    tl.Debug("hello %s", "jack")

    // new logo with file.
    f, err := logo.OpenFile("example.log", 0644)
    if err != nil {
        // ...
        return
    }
    defer f.Close()
    fl := logo.New(logo.LevelInfo, f, "", logo.LstdFlags|logo.Llevel|logo.Lshortfile)
    fl.Warning("something wrong")

    // new logo with smtp
    sw, err := logo.NewSMTPWriter(
        "smtp.example.com:465", // smtp addr
        "username",             // username
        "password",             // password
        "TestLogo",             // email subject
        "you@email.com",        // to-list
    )
    if err != nil {
        // ...
        return
    }
    defer sw.Close()
    sl := logo.New(logo.LevelError, sw, "", logo.LstdFlags|logo.Llevel|logo.Lshortfile)
    sl.Error("uncatched exception!")

    // new logo group with all above logos
    l := logo.Group(logo.LevelDebug, tl, fl, sl)
    l.Info("group message")
}
```
