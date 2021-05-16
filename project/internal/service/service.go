package service

import "github.com/google/wire"

var Provider = wire.NewSet(NewGroupService, NewNamespaceService, NewService)

type Service struct {
	GroupService     *GroupService
	NamespaceService *NamespaceService
}
var AppService *Service

func NewService(groupService *GroupService, namespaceService *NamespaceService)error  {
	AppService= &Service{
		GroupService:     groupService,
		NamespaceService: namespaceService,
	}
	return nil
}
