package storage

import (
	"context"
	"fmt"
	"google.golang.org/api/drive/v2"
	"google.golang.org/api/option"
	"net/http"
	"os"
)

type Folder struct {
	Name     string
	ParentID string
}

type File struct {
	Name     string
	MimeType string
	ParentID string
	Path     string
}

func uploadFile(client *http.Client, file *File) (*drive.File, error) {
	srv, err := drive.NewService(context.Background(), option.WithHTTPClient(client))
	if err != nil {
		return nil, fmt.Errorf("unable to retrieve Drive client: %v", err)
	}

	fileContent, err := os.Open(file.Path)
	if err != nil {
		return nil, fmt.Errorf("unable to open file: %v", err)
	}
	defer func(fileContent *os.File) {
		err := fileContent.Close()
		if err != nil {

		}
	}(fileContent)

	driveFile := &drive.File{
		OriginalFilename: file.Name,
		MimeType:         file.MimeType,
		Parents:          []*drive.ParentReference{{Id: file.ParentID}}, // Set the parent folder ID
	}

	createdFile, err := srv.Files.Insert(driveFile).Media(fileContent).Do()
	if err != nil {
		return nil, fmt.Errorf("unable to create file: %v", err)
	}

	return createdFile, nil
}

func CreateFolder(client *http.Client, folder *Folder) (*drive.File, error) {
	srv, err := drive.NewService(context.Background(), option.WithHTTPClient(client))
	if err != nil {
		return nil, fmt.Errorf("unable to retrieve Drive client: %v", err)
	}

	driveFolder := &drive.File{
		OriginalFilename: folder.Name,
		Parents:          []*drive.ParentReference{{Id: folder.ParentID}},
		MimeType:         "application/vnd.google-apps.folder",
	}

	createdFolder, err := srv.Files.Insert(driveFolder).Do()
	if err != nil {
		return nil, fmt.Errorf("unable to create folder: %v", err)
	}

	return createdFolder, nil
}
