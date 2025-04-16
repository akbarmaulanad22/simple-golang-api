package usecase

import (
	"api/internal/entity"
	"api/internal/repository"

	"gorm.io/gorm"
)

type RuangKelasUsecase interface {
	GetAllRuangKelass() ([]entity.RuangKelas, error)
	CreateRuangKelas(ruangKelas *entity.RuangKelas) error
	UpdateRuangKelas(id uint, ruangKelas *entity.RuangKelas) error
	DeleteRuangKelas(id uint) error
}

type ruangKelasUsecase struct {
	ruangKelasRepo repository.RuangKelasRepository
}

func NewRuangKelasUsecase(db *gorm.DB) RuangKelasUsecase {
	return &ruangKelasUsecase{
		ruangKelasRepo: repository.NewRuangKelasRepository(db),
	}
}

func (u *ruangKelasUsecase) GetAllRuangKelass() ([]entity.RuangKelas, error) {
	return u.ruangKelasRepo.FindAll()
}

func (u *ruangKelasUsecase) CreateRuangKelas(ruangKelas *entity.RuangKelas) error {
	return u.ruangKelasRepo.Create(ruangKelas)
}

func (u *ruangKelasUsecase) UpdateRuangKelas(id uint, ruangKelas *entity.RuangKelas) error {
	return u.ruangKelasRepo.Update(id, ruangKelas)
}

func (u *ruangKelasUsecase) DeleteRuangKelas(id uint) error {
	return u.ruangKelasRepo.Delete(id)
}
