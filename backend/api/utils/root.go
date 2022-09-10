package utils

import (
	"errors"

	"github.com/gin-gonic/gin"
)

var (
	// ErrContentTypeNotAllowed is used when request contains an incorrect Content-Type.
	errContentTypeNotAllowed = errors.New("content type not allowed")

	// Negotiate is used to express which Accept and Content-Type MIME types are allowed.
	Negotiate = []string{gin.MIMEJSON, gin.MIMEXML}
)

type ParamParser interface {
	Param(string) string
}

type QueryParser interface {
	Query(string) string
	DefaultQuery(string, string) string
	QueryArray(string) []string
}

func ErrRes(g *gin.Context, err error, statusCode int) {
	g.Negotiate(statusCode, gin.Negotiate{
		Offered: Negotiate,
		Data: WrapperResponse{
			Code: statusCode,
			Msg:  err.Error(),
		},
	})
}
