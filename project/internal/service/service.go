package service

import "github.com/google/wire"

var Provider = wire.NewSet(NewAppService,NewNamespaceService,NewService)

type Service struct {
	AppService       *AppService
	NamespaceService *NamespaceService
}

func NewService(appService *AppService, namespaceService *NamespaceService) *Service {
	return &Service{
		AppService:       appService,
		NamespaceService: namespaceService,
	}
}
