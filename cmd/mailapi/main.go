package main

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"os"

	"os/signal"
	"syscall"

	"github.com/l-vitaly/jsonrpc"
	"github.com/l-vitaly/jsonrpc/json2"
	"github.com/l-vitaly/mailapi"
	"gopkg.in/gomail.v2"
)

func main() {
	username := os.Getenv("MAIL_USERNAME")
	password := os.Getenv("MAIL_PASSWD")
	sendTo := os.Getenv("MAIL_TO")
	sendFrom := os.Getenv("MAIL_FROM")
	subject := os.Getenv("MAIL_SUBJECT")

	if username == "" || password == "" || sendTo == "" {
		fmt.Println("MAIL_TO, MAIL_USERNAME and MAIL_PASSWD env must be required")
		os.Exit(1)
	}

	d := gomail.NewDialer("smtp.gmail.com", 465, username, password)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	jrpc := jsonrpc.NewServer()
	jrpc.RegisterCodec(json2.NewCodec(), "application/json")
	jrpc.RegisterService(mailapi.NewService(d, sendTo, sendFrom, subject), "Service")

	errCh := make(chan error)

	// Interrupt handler.
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errCh <- fmt.Errorf("%s", <-c)
	}()

	go func() {
		errCh <- http.ListenAndServe(":9000", jrpc)
	}()

	fmt.Println("exit", <-errCh)
}
