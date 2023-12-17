package main

import (
	"client/consumer/initializer"
	"client/consumer/internal/api"
	"client/consumer/internal/db"
	"client/consumer/internal/repository"
	"client/consumer/internal/storage"
	"github.com/gin-gonic/gin"
)

func init() {
	db.Connection()

	initializer.LoadEnvironment()
	storage.ConnectToDrive()
	//db.ConnectFirebase()
	//utils.CheckCollectionExist()
}

type BasicUrl string

const (
	Basic BasicUrl = "/api/v1/series/"
)

func setupRoutes(c *gin.Engine, api *api.ApiService) {
	c.GET(string(Basic+"/retrieveManga"), api.RetrieveData)

	c.GET(string(Basic+"/retrieveAdditionalMangaData"), api.AdditionalService.GetAdditionalData)
	c.GET(string(Basic+"/getChapters"), api.AdditionalService.GetMangaChapters)

}

func main() {

	server := gin.Default()

	mangaRepository := repository.NewUserRepository(db.GetDB())
	apiService := api.NewService(nil, mangaRepository)
	apiService.AdditionalService = &api.AdditionalMangaMethod{
		MangaRepository: mangaRepository,
	}
	setupRoutes(server, apiService)

	server.Run()

}
