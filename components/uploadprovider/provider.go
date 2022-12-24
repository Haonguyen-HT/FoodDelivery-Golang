package uploadprovider

import (
	"FoodDelivery/common"
	"context"
	"github.com/gin-gonic/gin"
	"mime/multipart"
)

type UploadProvider interface {
	SaveFileUploaded(ctx context.Context, data []byte, dst string) (*common.Image, error)
	SaveUploadedLocal(
		ctx *gin.Context,
		data *multipart.FileHeader,
	//fileBytes []byte,
		dst string,
	) (*common.Image, error)
}