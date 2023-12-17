package api

import (
	"client/consumer/internal/models"
	"client/consumer/internal/repository"
	"client/consumer/internal/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AdditionalServiceMangaMethods interface {
	GetAdditionalData(c *gin.Context)
	GetMangaChapters(c *gin.Context)
}

type AdditionalMangaMethod struct {
	MangaRepository *repository.MangaRepository
	MangaServices   *services.MangaServiceMethod
}

func (s *AdditionalMangaMethod) GetAdditionalData(c *gin.Context) {
	// Fetch data from the database
	dataFetched, err := s.MangaRepository.FetchBySizeRecords()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Iterate over the fetched data
	for _, manga := range dataFetched {
		// Retrieve additional data for each URL
		updatedManga, err := s.MangaServices.RetrieveAdditionalData([]models.Manga{manga})

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		// Save the updated data
		_, err = s.MangaRepository.SaveAdditionalDataManga(updatedManga)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Data processing and saving completed successfully",
	})
}
func (s *AdditionalMangaMethod) GetMangaChapters(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"message": "Data processing and saving completed successfully",
	})
}
