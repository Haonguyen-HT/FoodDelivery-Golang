package uploadbusiness

import (
	"FoodDelivery/common"
	"FoodDelivery/components/uploadprovider"
	"bytes"
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slices"
	"image"
	"io"
	"mime/multipart"
	"path/filepath"
	"strings"
	"time"
)

type uploadBiz struct {
	provider uploadprovider.UploadProvider
}

func NewUploadBiz(provider uploadprovider.UploadProvider) *uploadBiz {
	return &uploadBiz{
		provider: provider,
	}
}

func (biz *uploadBiz) Upload(
	ctx context.Context,
	data []byte, folder, fileName string,
) (*common.Image, error) {
	fileBytes := bytes.NewBuffer(data)
	w, h, err := getImageDimension(fileBytes)

	if strings.TrimSpace(folder) == "" {
		folder = "img"
	}

	fileExt := filepath.Ext(fileName)
	fileName = fmt.Sprintf("%d%s", time.Now().Nanosecond(), fileExt)

	img, err := biz.provider.SaveFileUploaded(ctx, data, fmt.Sprintf("%s/%s", folder, fileName))
	if err != nil {
		return nil, err
	}

	img.Width = w
	img.Height = h
	img.Extension = fileExt

	return img, nil

}
func (biz *uploadBiz) UploadLocal(
	c *gin.Context,
	file *multipart.FileHeader,
	folder, fileName string,
) (*common.Image, error) {

	if strings.TrimSpace(folder) == "" {
		folder = "img"
	}

	fileExt := filepath.Ext(fileName)
	fileName = fmt.Sprintf("%d%s", time.Now().Nanosecond(), fileExt)

	fileExtCheck := []string{".jpg", ".png", ".jpeg", ".gif"}
	if slices.Contains(fileExtCheck, fileExt) {
		if err := c.SaveUploadedFile(file, fmt.Sprintf("static/%s", fileName)); err != nil {
			panic(err)
		}

		return &common.Image{
			ID:        0,
			Url:       "http://localhost:3001/static/" + fileName,
			Width:     0,
			Height:    0,
			CloudName: "",
			Extension: "",
		}, nil
	}

	return nil, errors.New("unsupported file format")

	//img, err := biz.provider.SaveUploadedLocal(c, file)
	//if err != nil {
	//	return nil, err
	//}

}

func getImageDimension(reader io.Reader) (int, int, error) {
	img, _, err := image.DecodeConfig(reader)

	if err != nil {
		return 0, 0, err
	}

	return img.Width, img.Height, nil
}