package main

import (
	// "moniteur/jobs"

	"crypto/tls"
	"fmt"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	// bot.Init()
	// jobs.Monitor()

	connection, err := tls.Dial("tcp", "admin.usedo.me:443", nil)

	if err != nil {
		panic("Application does not have SSL certificate " + err.Error())
	}

	err = connection.VerifyHostname("admin.usedo.me")

	if err != nil {
		panic("Hostname does not match SSL Certificate: " + err.Error())
	}

	issuer := connection.ConnectionState().PeerCertificates[0].Issuer
	expiry := connection.ConnectionState().PeerCertificates[0].NotAfter

	fmt.Printf("Issuer: %s\nExpiry:%v", issuer, expiry.Format(time.RFC850))
}
