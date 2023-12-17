package api

import (
	"client/consumer/internal/repository"
	"client/consumer/internal/services"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ApiMethods interface {
	AddManga(c *gin.Context)
	RetrieveData(c *gin.Context)
}

type ApiService struct {
	ApiMethod         ApiMethods
	AdditionalService AdditionalServiceMangaMethods
	MangaRepository   *repository.MangaRepository
	Services          *services.MangaServiceMethod
}

func NewService(apiMethod ApiMethods, mangaRepository *repository.MangaRepository) *ApiService {
	return &ApiService{
		ApiMethod:       apiMethod,
		MangaRepository: mangaRepository,
	}
}

func (s *ApiService) AddManga(c *gin.Context) {
	mangaData := s.Services.GetAllManga()

	batchSize := 100

	numBatches := (len(mangaData) + batchSize - 1) / batchSize

	for i := 0; i < numBatches; i++ {
		start := i * batchSize
		end := (i + 1) * batchSize
		if end > len(mangaData) {
			end = len(mangaData)
		}

		// Extract a batch of manga data
		batch := mangaData[start:end]

		// Attempt to create the batch of manga in the database
		created, err := s.MangaRepository.CreateManga(batch)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		fmt.Println("Batch inserted successfully:", created)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Manga data successfully saved to the database.",
	})
}

func (s *ApiService) RetrieveData(c *gin.Context) {

	result, err := s.MangaRepository.GetAllMangaData()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		}) // Process the results if needed (created contains the results of the batch insertion)

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": result,
	})
}
