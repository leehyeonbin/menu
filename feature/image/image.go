package image

import (
	"menu-go/feature/image/data/local/store"
	"menu-go/feature/image/data/remote/api"
)

func FetchAndSaveImage(url string, filename string) error {
	image, fetchErr := api.FetchImage(url)
	if fetchErr != nil {
		return fetchErr
	}
	saveErr := store.SaveImage(image, filename)
	if saveErr != nil {
		return saveErr
	}
	return nil
}
