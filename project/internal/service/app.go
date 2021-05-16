package service

import "project/internal/dao"

type AppService struct {
	db *dao.Dao
}

func NewAppService(db *dao.Dao) *AppService {
	return &AppService{db: db}
}
