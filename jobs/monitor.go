package jobs

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Service struct {
	Name         string `json:"name"`
	Url          string `json:"url"`
	Service_type string `json:"type"`
}

type ServiceData struct {
	Services []Service `json:"services"`
}

func Monitor() {
	directory, _ := os.Getwd()
	servicesPath := filepath.Join(directory, "services.json")
	serviceContent, _ := os.ReadFile(servicesPath)

	var serviceData *ServiceData
	json.Unmarshal(serviceContent, &serviceData)

	fmt.Println(serviceData)
	services := serviceData.Services

	fmt.Println(services[0].Name)
}
