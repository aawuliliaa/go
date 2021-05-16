package service

import (
	"context"
	pb "project/api"
	"project/internal/dao"
)

type GroupService struct {
	*pb.UnimplementedGroupServiceServer
	dao *dao.Dao
}

func NewGroupService(db *dao.Dao) *GroupService {
	return &GroupService{dao: db}
}

func (groupService *GroupService) GetGroupInfo(ctx context.Context, in *pb.GetGroupRequest) (*pb.GetGroupResponse, error) {
	name := in.Name
	name = groupService.dao.GetGroupInfo(name)
	return &pb.GetGroupResponse{
		Name: name,
	}, nil
}
