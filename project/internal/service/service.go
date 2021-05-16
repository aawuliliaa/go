package service

import "github.com/google/wire"

var Provider = wire.NewSet(NewGroupService, NewNamespaceService, NewService)

type Service struct {
	GroupService     *GroupService
	NamespaceService *NamespaceService
}


func NewService(groupService *GroupService, namespaceService *NamespaceService)*Service  {
	appService:= &Service{
		GroupService:     groupService,
		NamespaceService: namespaceService,
	}
	return appService
}
