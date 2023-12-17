package models

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Manga struct {
	gorm.Model

	ID            int             `gorm:"primarykey" json:"id"`
	Title         string          `json:"title"`
	Url           string          `json:"url"`
	Image         string          `json:"image"`
	Rating        string          `json:"rating"`
	MangaDetails  []MangaDetails  `gorm:"foreignKey:MangaID"`
	MangaChapters []MangaChapters `gorm:"foreignKey:MangaID"`
}

type MangaDetails struct {
	gorm.Model

	ID          int    `gorm:"primarykey" json:"id"`
	MangaID     int    // Foreign key
	Description string `json:"description"`
	Genre       string `json:"genre"`
}

type MangaChapters struct {
	gorm.Model

	ID                int            `gorm:"primarykey" json:"id"`
	MangaID           int            // Foreign key
	MangaChaptersUrls pq.StringArray `gorm:"type:text[]" json:"chapter_urls"`

	MangaChapterCount int `json:"manga_chapter_count"`
}

type MangaChaptersHolder struct {
	MangaChaptersUrls pq.StringArray `json:"chapter_urls"`

	MangaChapterCount      int    `json:"manga_chapter_count"`
	MangaChapterUrlVisited string `json:"manga_chapter_url_visited"`
}
