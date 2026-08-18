package handler


import (
	"github.com/eddywm/go-shortner/shortener"
	"github.com/eddywm/go-shortner/store"
	"github.com/gin-gonic/gin"
	"net/http"
)

//request model definition
type urlCreateRequest struct {
	LongUrl string `json:"long_url" binding:"required"`
	UserId  string `json:"user_id" binding:"required"`
}

func CreateShortUrl(c *gin.Context) {
	
    /// Implementation to be added

}

func HandleShortUrlRedirect(c *gin.Context) {

	/// Implementation to be added

}