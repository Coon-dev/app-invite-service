package configs

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	Timeout       time.Duration
	HTTPTransport http.Transport
	HTTPCilent    *http.Client
	Clog          *log.Logger
	TimeZone      *time.Location
	MongoClient   *mongo.Client
)

const AuthKey string = "Basic a2mJIp6IOyZihYvw60WSwzprkB8AHGyOxtvmh0k1U4Lr0upv1LVpi4y"

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

	// err = DB.Ping()
	// if err != nil {
	// 	Clog.Fatalln("Ping SQL Failed:", err)
	// }
	option := options.Client()
	option.SetMaxPoolSize(50)
	option.SetMaxConnIdleTime(Timeout)
	option.SetMaxConnecting(50)

	MongoClient, err = mongo.NewClient(option.ApplyURI("mongodb+srv://pulseid:pulseid123@cluster0.ncnqr.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}

	err = MongoClient.Connect(context.Background())
	if err != nil {
		log.Fatal(err)
	}
}

type logWriter struct {
}

func (writer logWriter) Write(bytes []byte) (int, error) {
	return fmt.Print("[LOG-DEBUG] " + string(bytes))
}
