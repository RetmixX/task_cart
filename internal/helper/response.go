package helper

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strings"
	"task_cart/internal/model/consts"
)

//Statuses:
//[200, 201, 404, 500, 422]

func OkResponse(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}

func CreatedResponse(c *gin.Context, data interface{}) {
	c.JSON(http.StatusCreated, gin.H{
		"data": data,
	})
}

func NotFoundResponse(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"error": consts.NotFoundErr,
	})
}

func ValidationErr(c *gin.Context, err error) {
	c.JSON(http.StatusUnprocessableEntity, gin.H{
		"error":  consts.InvalidJsonErr,
		"errors": parseError(err.(validator.ValidationErrors)),
	})
}

func UrlParamErr(c *gin.Context) {
	c.JSON(http.StatusBadRequest, gin.H{
		"error": consts.InvalidURLErr,
	})
}

func ServerErr(c *gin.Context) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"error": consts.ServerErr,
	})
}

func parseError(err validator.ValidationErrors) map[string]string {
	errorMsg := make(map[string]string, len(err))
	for _, v := range err {
		tag := v.Tag()
		switch tag {
		case "required":
			errorMsg[strings.ToLower(v.Field())] = fmt.Sprintf("The field is %s", tag)
		case "email":
			errorMsg[strings.ToLower(v.Field())] = fmt.Sprintf("The field is not %s", tag)
		case "min":
			errorMsg[strings.ToLower(v.Field())] = fmt.Sprintf("The field is smaller %s", v.Param())
		case "max":
			errorMsg[strings.ToLower(v.Field())] = fmt.Sprintf("The field is bigger %s", v.Param())
		}
	}

	return errorMsg
}
