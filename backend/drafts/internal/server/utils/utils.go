package utils

import (
	"crypto/rsa"
	"errors"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"

	"github.com/davg/drafts/internal/customerrors"
)

func GetKey(pathToKey string) *rsa.PublicKey {
	data, err := os.ReadFile(pathToKey)
	if err != nil {
		log.Fatal(err, pathToKey)
	}
	key, err := jwt.ParseRSAPublicKeyFromPEM(data)
	if err != nil {
		log.Fatal(err)
	}

	return key
}

func HandleError(ctx *gin.Context, err error) {
	if errors.Is(err, customerrors.ErrBadRequest) {
		ctx.JSON(400, gin.H{"status": "error", "message": err.Error()})
	} else if errors.Is(err, customerrors.ErrUnauthorized) {
		ctx.JSON(401, gin.H{"status": "error", "message": err.Error()})
	} else if errors.Is(err, customerrors.ErrForbidden) {
		ctx.JSON(403, gin.H{"status": "error", "message": err.Error()})
	} else if errors.Is(err, customerrors.ErrNotFound) {
		ctx.JSON(404, gin.H{"status": "error", "message": err.Error()})
	} else if errors.Is(err, customerrors.ErrConflict) {
		ctx.JSON(409, gin.H{"status": "error", "message": err.Error()})
	} else if errors.Is(err, customerrors.ErrInternal) {
		ctx.JSON(500, gin.H{"status": "error", "message": err.Error()})
	} else {
		ctx.JSON(500, gin.H{"status": "error", "message": err.Error()})
	}
}
