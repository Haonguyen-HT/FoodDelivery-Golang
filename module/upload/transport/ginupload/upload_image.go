package ginupload

import (
	"FoodDelivery/common"
	"FoodDelivery/components/appcontext"
	"FoodDelivery/module/upload/uploadbusiness"
	"github.com/gin-gonic/gin"
)

func UploadImage(ctx appcontext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		fileHeader, err := c.FormFile("file")

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		folder := c.DefaultPostForm("folder", "img")
		file, err := fileHeader.Open()

		if err != nil {
			panic(common.ErrInternal(err))
		}

		defer file.Close()
		dataBytes := make([]byte, fileHeader.Size)
		if _, err := file.Read(dataBytes); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		biz := uploadbusiness.NewUploadBiz(ctx.UploadProvider())
		img, err := biz.UploadLocal(c, fileHeader, folder, fileHeader.Filename)

		if err != nil {
			panic(err)
		}
		c.JSON(200, common.SimpleSuccessResponse(img))
	}
}