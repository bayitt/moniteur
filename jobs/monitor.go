package jobs

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

type Service struct {
	Name        string    `json:"name"`
	Url         string    `json:"url"`
	ServiceType string    `json:"type"`
	Active      bool      `json:"active"`
	SslExpiry   time.Time `json:"ssl_expiry"`
}

type ServiceData struct {
	Services []Service `json:"services"`
}

type PingResponse struct {
	Status string `json:"status"`
}

func Monitor() {
	directory, _ := os.Getwd()
	servicesPath := filepath.Join(directory, "services.json")
	serviceContent, _ := os.ReadFile(servicesPath)

	var serviceData *ServiceData
	json.Unmarshal(serviceContent, &serviceData)

	services := serviceData.Services

	serviceLength := len(services)
	serviceChannel := make(chan Service, serviceLength)
	var wg sync.WaitGroup

	wg.Add(serviceLength)

	for _, service := range services {
		go PingService(service, &wg, serviceChannel)
	}

	wg.Wait()

	close(serviceChannel)

	for service := range serviceChannel {
		fmt.Println(service)
	}

}

func PingService(service Service, wg *sync.WaitGroup, channel chan<- Service) {
	defer wg.Done()

	var requestUrl string

	if service.ServiceType == "frontend" {
		requestUrl = service.Url
	} else {
		requestUrl = service.Url + "/ping"
	}

	request, _ := http.NewRequest(http.MethodPost, requestUrl, nil)
	response, err := http.DefaultClient.Do(request)
	var pingResponse PingResponse

	if err != nil {
		service.Active = false
		channel <- service
		return
	}

	responseBody, _ := ioutil.ReadAll(response.Body)

	if service.ServiceType == "frontend" {
		service.Active = true
	} else {
		json.Unmarshal(responseBody, &pingResponse)

		if strings.ToLower(pingResponse.Status) == "ok" {
			service.Active = true
		} else {
			service.Active = false
		}
	}

	sslExpiry, err := CheckSslExpiry(strings.Split(service.Url, "https://")[1])

	if err != nil {
		channel <- service
		return
	}

	service.SslExpiry = sslExpiry

	channel <- service
}

func CheckSslExpiry(domain string) (time.Time, error) {
	connection, err := tls.Dial("tcp", domain+":443", nil)

	if err != nil {
		return time.Now(), errors.New("SSL Certificate could not be obtained")
	}

	err = connection.VerifyHostname(domain)

	if err != nil {
		return time.Now(), errors.New("SSL Certificate does not match domain")
	}

	return connection.ConnectionState().PeerCertificates[0].NotAfter, nil
}
