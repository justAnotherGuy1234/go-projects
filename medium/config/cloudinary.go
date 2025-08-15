package config

import (
	"fmt"
	"os"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/joho/godotenv"
)

func SetupCloudinary() (*cloudinary.Cloudinary, error) {

	if err := godotenv.Load(); err != nil {
		fmt.Println("error loading env in setup cloudinary", err)
	}

	cld, _ := cloudinary.NewFromParams(os.Getenv("CLOUD_NAME"), os.Getenv("CLOUD_KEY"), os.Getenv("CLOUD_SECRET"))

	return cld, nil
}
