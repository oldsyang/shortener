package core

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"shortener/global"
	"shortener/lib/base62"
	"shortener/lib/db"
	"shortener/models"
	"strings"
)

type EncodeForm struct {
	Url string `json:"url" binding:"required"`
}

func Encode(ctx *gin.Context) {
	var urlModel models.UrlModel
	if err := ctx.ShouldBind(&urlModel); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ctx.JSON(http.StatusOK, gin.H{
				"msg": err.Error(),
			})
		}

		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": func(fields map[string]string) map[string]string {
				resp := map[string]string{}
				for field, err := range fields {
					resp[strings.ToLower(field[strings.Index(field, ".")+1:])] = err
				}
				return resp
			}(errs.Translate(global.Trans)),
		})
		return
	}

	db := db.GetDB()
	db.Where(models.UrlModel{Url: urlModel.Url}).FirstOrCreate(&urlModel)
	if err := db.Error; err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"success": "false",
		})
	}
	code := base62.Encode(int64(urlModel.Model.ID))
	ctx.JSON(http.StatusOK, gin.H{
		"success": "true",
		"response": gin.H{
			"url": fmt.Sprintf("%s/%d/%s", global.ServerConfig.Server, 3, code),
		},
	})

}

func Redirect(ctx *gin.Context) {
	var reRq models.RedirectRequest
	if err := ctx.ShouldBindUri(&reRq); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ctx.JSON(http.StatusOK, gin.H{
				"msg": err.Error(),
			})
		}

		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": func(fields map[string]string) map[string]string {
				resp := map[string]string{}
				for field, err := range fields {
					resp[strings.ToLower(field[strings.Index(field, ".")+1:])] = err
				}
				return resp
			}(errs.Translate(global.Trans)),
		})
		return
	}

	sqlId, _ := base62.Decode(reRq.Code)
	var urlModel models.UrlModel
	db := db.GetDB()
	db.Where("id=?", sqlId).First(&urlModel)
	if err := db.Error; err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"success": "false",
		})
	}
	ctx.Redirect(http.StatusMovedPermanently, urlModel.Url)
}

func Info(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"data": gin.H{
			"key": "test",
		},
		"meta": gin.H{
			"total": 0,
		},
	})
}
