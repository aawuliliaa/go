package service

import "project/internal/dao"

type NamespaceService struct {
	dao *dao.Dao
}

func NewNamespaceService(db *dao.Dao) *NamespaceService {
	return &NamespaceService{dao: db}
}
