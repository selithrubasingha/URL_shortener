package handler

import (
	"net/http"

	"github.com/eddywm/go-shortner/shortener"
	"github.com/eddywm/go-shortner/store"
	"github.com/gin-gonic/gin"
	"github.com/selithrubasingha/url-shortener/store"
)

//request model definition
type urlCreateRequest struct {
	LongUrl string `json:"long_url" binding:"required"`
	UserId  string `json:"user_id" binding:"required"`
}

func CreateShortUrl(c *gin.Context) {
	
    var creationRequest urlCreateRequest
	if err := c.ShouldBindJSON(&creationRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	shortUrl, err := shortener.CreateShortUrl(creationRequest.LongUrl, creationRequest.UserId)
	
	store.SaveUrlMapping(shortUrl,creationRequest.LongUrl , creationRequest.UserId)

	host := "http://localhost:9808/"
	c.JSON(200, gin.H{
		"message":   "short url created successfully",
		"short_url": host +shortUrl,
	})

}

func HandleShortUrlRedirect(c *gin.Context) {

	/// Implementation to be added

}