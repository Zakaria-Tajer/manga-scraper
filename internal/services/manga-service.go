package services

import (
	"client/consumer/internal/models"
	"errors"
	"fmt"
	"github.com/gocolly/colly"
	"log"
	"strings"
)

type ServiceMethods interface {
	GetFirstPageMangaData() []models.Manga
	GetAllManga() []models.Manga
	RetrieveAdditionalData(mangaList []models.Manga) ([]models.Manga, error)
	GetChapters(manga []models.Manga) ([]models.Manga, error)
}

type MangaServiceMethod struct {
	ServiceMethod ServiceMethods
}

type DataToReturn struct {
	Data     []models.Manga
	Chapters int
}

func (s *MangaServiceMethod) GetFirstPageMangaData() []models.Manga {

	collector := Connection()

	var structMangaSlice []models.Manga

	collector.OnHTML(".page-item-detail", func(element *colly.HTMLElement) {
		mangaSlice := models.Manga{
			Title:  element.ChildText("div.post-title h3 a"),
			Url:    element.ChildAttr("div.item-thumb a", "href"),
			Image:  element.ChildAttr("div.item-thumb img", "src"),
			Rating: element.ChildText("div.meta-item span"),
		}

		structMangaSlice = append(structMangaSlice, mangaSlice)

	})

	collector.OnRequest(func(r *colly.Request) {
		fmt.Println("Loading", r.URL.String())
	})

	err := collector.Visit("https://azoramoon.com/")
	if err != nil {
		return nil
	}

	return structMangaSlice
}

func (s *MangaServiceMethod) GetAllManga() []models.Manga {
	collector := Connection()

	var structMangaSlice []models.Manga

	url := Link + "/page/"
	pagesNumber := 28

	for i := 1; i <= pagesNumber; i++ {
		pageURL := fmt.Sprintf("%s%d", url, i)

		// Use a closure to capture the current value of pageURL
		func(url string) {
			collector.OnHTML(".page-item-detail", func(element *colly.HTMLElement) {
				mangaSlice := models.Manga{
					Title:  element.ChildText("div.post-title h3 a"),
					Url:    element.ChildAttr("div.item-thumb a", "href"),
					Image:  element.ChildAttr("div.item-thumb img", "src"),
					Rating: element.ChildText("div.meta-item span"),
				}

				structMangaSlice = append(structMangaSlice, mangaSlice)
			})

			collector.OnRequest(func(r *colly.Request) {
				fmt.Println("Loading", r.URL.String())
			})

			// Visit the URL
			err := collector.Visit(url)
			if err != nil {
				log.Println(err)
			}
		}(pageURL)
	}

	// Wait for all requests to finish before returning
	collector.Wait()

	return structMangaSlice
}
func (s *MangaServiceMethod) RetrieveAdditionalData(mangaList []models.Manga) ([]models.Manga, error) {
	if len(mangaList) == 0 {
		return nil, errors.New("manga list is empty")
	}

	collector := Connection()

	for i := 0; i < len(mangaList); i++ {
		url := mangaList[i].Url

		// Create a closure to capture the current manga index
		func(index int) {
			collector.OnHTML(".post-content", func(element *colly.HTMLElement) {
				mangaDetails := models.MangaDetails{
					Description: element.ChildText("div.manga-summary p"),
				}

				element.ForEach("div.genres-content a", func(_ int, genreElement *colly.HTMLElement) {
					mangaDetails.Genre += genreElement.Text + ","
				})

				// Remove the trailing comma, if any
				mangaDetails.Genre = strings.TrimSuffix(mangaDetails.Genre, ",")

				// Append MangaDetails to the corresponding manga instance
				mangaList[index].MangaDetails = append(mangaList[index].MangaDetails, mangaDetails)
			})
		}(i)

		collector.OnRequest(func(r *colly.Request) {
			fmt.Println("Loading", r.URL.String(), mangaList[i].Url)
		})

		err := collector.Visit(url)

		if err != nil {
			log.Println(err)
		}
	}

	// Wait for the collector only if there are URLs to visit
	if len(mangaList) > 0 {
		collector.Wait()
	}

	return mangaList, nil
}
func (s *MangaServiceMethod) GetChapters(mangaList []models.Manga) ([]models.Manga, error) {
	if len(mangaList) == 0 {
		return nil, errors.New("manga list is empty")
	}

	collector := Connection()

	// Create a function to handle the OnHTML event
	handleHTML := func(m *models.Manga) func(element *colly.HTMLElement) {
		return func(element *colly.HTMLElement) {
			var ChapterUrls []string
			element.ForEach("ul.version-chap li", func(_ int, genreElement *colly.HTMLElement) {
				urls := genreElement.ChildAttr("div.wp-manga a", "href")
				ChapterUrls = append(ChapterUrls, urls)
			})

			mangaChapters := models.MangaChapters{
				MangaChapterCount: len(ChapterUrls),
				MangaChaptersUrls: ChapterUrls,
			}

			m.MangaChapters = append(m.MangaChapters, mangaChapters)

		}
	}

	// Register the OnHTML event outside the loop
	collector.OnHTML(".version-chap", func(element *colly.HTMLElement) {})
	collector.OnRequest(func(r *colly.Request) {
		fmt.Println("Loading", r.URL.String())
	})

	for i := 0; i < len(mangaList); i++ {
		url := mangaList[i].Url

		// Use the closure to capture the current manga item
		collector.OnHTML(".version-chap", handleHTML(&mangaList[i]))

		err := collector.Visit(url)

		if err != nil {
			log.Println(err)
		}

		// Clear ChapterUrls for the next iteration
		// This step is not needed since we are using a closure
	}

	// Wait for the collector only if there are URLs to visit
	if len(mangaList) > 0 {
		collector.Wait()
	}

	return mangaList, nil
}
