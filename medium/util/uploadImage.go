package util

import (
	"context"
	"fmt"
	"io"
	"medium/config"

	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

func UploadImage(file io.Reader) string {

	cld, _ := config.SetupCloudinary()

	res, err := cld.Upload.Upload(context.Background(), file, uploader.UploadParams{
		Folder: "go_images",
	})

	if err != nil {
		fmt.Println("error adding image to cloundinary ", err)
	}

	return res.SecureURL
}
