package repository

import (
	"client/consumer/internal/models"
	"log"
)

func (r *MangaRepository) AddMangaChapters(manga []models.MangaChapters) ([]models.MangaChapters, error) {
	result := r.DB.Create(&manga)

	if err := result.Error; err != nil {
		log.Println(err)
		return manga, err
	}
	return manga, nil
}
