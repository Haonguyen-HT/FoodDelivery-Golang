package uploadprovider

import (
	"FoodDelivery/common"
	"fmt"
	"github.com/gin-gonic/gin"
	"mime/multipart"
)

func SaveUploadedLocal(
	ctx *gin.Context,
	data *multipart.FileHeader,
//fileBytes []byte,
	dst string,
) (*common.Image, error) {
	fmt.Println("vao upload")
	if err := ctx.SaveUploadedFile(data, fmt.Sprintf("static/%s", dst)); err != nil {
		panic(err)
	}

	img := &common.Image{
		ID:        0,
		Url:       fmt.Sprintf("%s/%s", "http://localhost:3001/static/", dst),
		CloudName: "local",
		Width:     0,
		Height:    0,
		Extension: "",
	}

	return img, nil
}