package service

import (
	"fmt"
	"hospitie/constant"
	"hospitie/core"
	"hospitie/model"

	"gorm.io/gorm"
)

type HospitieService struct {
	DB *gorm.DB
}

func NewHospitieService() *HospitieService {
	db := constant.DBInstance()
	return &HospitieService{
		DB: db,
	}
}

func (s *HospitieService) CreateHospitie(h model.Hospitie) (model.Hospitie, error) {
	res := s.DB.Create(&h)
	if res.Error != nil {
		fmt.Println(res.Error)
		// s.DB.Rollback()
	}
	return h, res.Error
}

func (s *HospitieService) GetAllHospities() []model.Hospitie {
	hs := []model.Hospitie{}
	res := s.DB.Preload("Cars").Find(&hs)
	if res.Error != nil {
		core.ErrorHandler(res.Error)
	}
	return hs
}

func (s *HospitieService) FilterHospitie(filter map[string]interface{}) []model.Hospitie {
	hs := []model.Hospitie{}
	res := s.DB.Where(filter).Find(&hs)
	if res.Error != nil {
		core.ErrorHandler(res.Error)
	}
	return hs
}

func (s *HospitieService) UpdateHospitie(h_id string, v map[string]interface{}) model.Hospitie {
	h := model.Hospitie{}
	res := s.DB.Model(s.DB.Find(&h, h_id)).Updates(v)
	if res.Error != nil {
		core.ErrorHandler(res.Error)
	}
	return h
}

func (s *HospitieService) DeleteHospitie(h_id string) model.Hospitie {
	h := model.Hospitie{}
	res := s.DB.Delete(&h, h_id)
	if res.Error != nil {
		core.ErrorHandler(res.Error)
	}
	return h
}
