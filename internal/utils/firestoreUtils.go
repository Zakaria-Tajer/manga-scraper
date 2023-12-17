package utils

type MangaChapterModel struct {
	MangaID           int      `firestore:"MangaID"`
	MangaChaptersUrls []string `firestore:"MangaChaptersUrls"`
	MangaChapterCount int      `firestore:"MangaChaptersUrls"`
}
