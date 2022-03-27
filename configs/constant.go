package configs

import (
	"crypto/tls"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var (
	Timeout       time.Duration
	HTTPTransport http.Transport
	HTTPCilent    *http.Client
	Clog          *log.Logger
	TimeZone      *time.Location
	DB            *sql.DB
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

	DB, err = sql.Open("mysql", "sql6481885:sql6481885@tcp(sql6.freemysqlhosting.net:3306)/sql6481885")
	if err != nil {
		Clog.Fatalln("Open SQL Failed:", err)
	}

	DB.SetMaxOpenConns(50)
	DB.SetMaxIdleConns(50)
	DB.SetConnMaxIdleTime(Timeout)
	DB.SetConnMaxLifetime(Timeout)

	err = DB.Ping()
	if err != nil {
		Clog.Fatalln("Ping SQL Failed:", err)
	}
	defer DB.Close()
}

type logWriter struct {
}

func (writer logWriter) Write(bytes []byte) (int, error) {
	return fmt.Print("[LOG-DEBUG] " + string(bytes))
}
