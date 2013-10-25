package api

import (
	"cf"
	"cf/net"
)

type FakeUserProvidedServiceInstanceRepo struct {
	CreateName string
	CreateParameters map[string]string
	CreateSyslogDrainUrl string

	UpdateServiceInstance cf.ServiceInstance
}

func (repo *FakeUserProvidedServiceInstanceRepo) Create(name string, params map[string]string, syslogDrainUrl string) (apiResponse net.ApiResponse) {
	repo.CreateName = name
	repo.CreateParameters = params
	repo.CreateSyslogDrainUrl = syslogDrainUrl
	return
}

func (repo *FakeUserProvidedServiceInstanceRepo) Update(serviceInstance cf.ServiceInstance) (apiResponse net.ApiResponse) {
	repo.UpdateServiceInstance = serviceInstance
	return
}
