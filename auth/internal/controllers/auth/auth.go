package auth

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	log    *slog.Logger
	client *http.Client
}

func New(logger *slog.Logger, client *http.Client) *AuthController {
	return &AuthController{
		log:    logger,
		client: client,
	}
}

func (ac *AuthController) Login(c *gin.Context) {

}

func (ac *AuthController) Register(c *gin.Context) {

}

func (ac *AuthController) Permissions(c *gin.Context) {

}
