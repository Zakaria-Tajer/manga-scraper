package repository

import (
	"client/consumer/internal/models"
	"log"

	"gorm.io/gorm"
)

type MangaRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *MangaRepository {
	return &MangaRepository{
		DB: db,
	}
}

func (r *MangaRepository) CreateManga(manga []models.Manga) ([]models.Manga, error) {
	result := r.DB.Create(&manga)

	if err := result.Error; err != nil {
		log.Println(err)
		return manga, err
	}
	return manga, nil
}

func (r *MangaRepository) FindById(manga []models.Manga) ([]models.Manga, error) {
	result := r.DB.Create(&manga)

	if err := result.Error; err != nil {
		log.Println(err)
		return manga, err
	}
	return manga, nil
}



func (r *MangaRepository) FindAllByUrl(url string) ([]models.Manga, error) {
	var mangas []models.Manga

	result := r.DB.Where("url = ?", url).Find(&mangas)

	if err := result.Error; err != nil {
		log.Println(err)
		return nil, err
	}

	return mangas, nil
}

func (r *MangaRepository) SaveAdditionalDataManga(manga []models.Manga) ([]models.Manga, error) {
	result := r.DB.Save(&manga)

	if err := result.Error; err != nil {
		log.Println(err)
		return manga, err
	}
	return manga, nil
}

func (r *MangaRepository) GetAllMangaData() ([]models.Manga, error) {
	var data []models.Manga

	if err := r.DB.Preload("MangaDetails").Preload("MangaChapters").Find(&data).Error; err != nil {
		return nil, err
	}

	return data, nil

}

func (r *MangaRepository) FetchBySizeRecords() ([]models.Manga, error) {

	var data []models.Manga

	result := r.DB.Limit(10).Find(&data)

	if err := result.Error; err != nil {
		log.Println(err)
		return nil, err
	}

	if result.RowsAffected == 0 {
		log.Printf("No more records with pageSize: \n")
		return data, nil
	}

	return data, nil

}
