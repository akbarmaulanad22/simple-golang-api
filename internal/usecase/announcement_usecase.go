package usecase

import (
	"api/internal/entity"
	"api/internal/repository"
	"encoding/json"

	"gorm.io/gorm"
)

type AnnouncementUsecase interface {
	GetAllAnnouncements() ([]entity.Announcement, error)
	CreateAnnouncement(announcement *entity.Announcement) error
	UpdateAnnouncement(id uint, announcement *entity.Announcement) error
	DeleteAnnouncement(id uint) error
	FindByIdAnnouncement(id uint) (entity.Announcement, error)
}

type announcementUsecase struct {
	announcementRepo repository.AnnouncementRepository
	logRepo repository.LogRepository
}

func NewAnnouncementUsecase(db *gorm.DB) AnnouncementUsecase {
	return &announcementUsecase{
		announcementRepo: repository.NewAnnouncementRepository(db),
		logRepo: repository.NewLogRepository(db),
	}
}

func (u *announcementUsecase) FindByIdAnnouncement(id uint) (entity.Announcement, error) {
	
	return u.announcementRepo.FindById(id)
	
}

func (u *announcementUsecase) GetAllAnnouncements() ([]entity.Announcement, error) {
	return u.announcementRepo.FindAll()
}

func (u *announcementUsecase) CreateAnnouncement(announcement *entity.Announcement) error {
	// currentAnnouncement, err := osAnnouncement.Current()
	// if err != nil {
	// 	return err
	// }
	
	// idInt, err := strconv.Atoi(currentAnnouncement.Uid)
	// if err != nil {
	// 	return err
	// }

	// id := new(uint)
    // *id = uint(idInt)
	
	errCreate := u.announcementRepo.Create(announcement)
	if errCreate != nil {
		return errCreate
	}

	jsonBytes, err := json.Marshal(announcement)
	if err != nil {
		return err
	}
	
	errCreateLog := u.logRepo.Create(&entity.Log{
		// AnnouncementID: id,
		Action: "CREATE",
		EntityType: "ANNOUNCEMENT",
		EntityID: announcement.ID,
		Details: string(jsonBytes),
	})

	if errCreateLog != nil{
		return errCreateLog
	}

	return nil
}

func (u *announcementUsecase) UpdateAnnouncement(id uint, announcement *entity.Announcement) error {
	
	errUpdate := u.announcementRepo.Update(id, announcement)
	if errUpdate != nil {
		return errUpdate
	}

	announcement.ID = id

	jsonBytes, err := json.Marshal(announcement)
	if err != nil {
		return err
	}
	
	errCreateLog := u.logRepo.Create(&entity.Log{
		// AnnouncementID: id,
		Action: "UPDATE",
		EntityType: "ANNOUNCEMENT",
		EntityID: id,
		Details: string(jsonBytes),
	})

	if errCreateLog != nil{
		return errCreateLog
	}
	
	return nil
}

func (u *announcementUsecase) DeleteAnnouncement(id uint) error {
	errCreateLog := u.logRepo.Create(&entity.Log{
		// AnnouncementID: id,
		Action: "DELETE",
		EntityType: "ANNOUNCEMENT",
		EntityID: id,
	})

	if errCreateLog != nil{
		return errCreateLog
	}

	return u.announcementRepo.Delete(id)

}
