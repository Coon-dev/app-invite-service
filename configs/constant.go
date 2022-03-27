package configs

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"time"
)

var (
	Timeout       time.Duration
	HTTPTransport http.Transport
	HTTPCilent    *http.Client
	Clog          *log.Logger
	TimeZone      *time.Location
)

func InitialConfig() {
	Clog = log.New(new(logWriter), "", log.LstdFlags|log.Lshortfile)
	Timeout = 5 * time.Second

	HTTPTransport = http.Transport{
		TLSClientConfig:       &tls.Config{InsecureSkipVerify: true},
		MaxIdleConns:          500,
		MaxIdleConnsPerHost:   500,
		IdleConnTimeout:       Timeout,
		TLSHandshakeTimeout:   Timeout,
		ResponseHeaderTimeout: Timeout,
		ExpectContinueTimeout: Timeout,
	}
	HTTPCilent = &http.Client{
		Timeout:   Timeout,
		Transport: &HTTPTransport,
	}

	tz, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		Clog.Fatalln("parse location error:", err)
	}
	TimeZone = tz
}

type logWriter struct {
}

func (writer logWriter) Write(bytes []byte) (int, error) {
	return fmt.Print("[LOG-DEBUG] " + string(bytes))
}
