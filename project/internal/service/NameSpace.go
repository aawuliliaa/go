package service

import "project/internal/dao"

type NamespaceService struct {
	db *dao.Dao
}

func NewNamespaceService(db *dao.Dao) *NamespaceService {
	return &NamespaceService{db: db}
}
