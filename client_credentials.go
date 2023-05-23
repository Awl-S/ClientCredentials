package main

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type OAuthServer struct {
	clients      map[string]string
	accessTokens map[string]string
}

func NewOAuthServer() *OAuthServer {
	return &OAuthServer{
		clients:      make(map[string]string),
		accessTokens: make(map[string]string),
	}
}

func (s *OAuthServer) TokenHandler(c *gin.Context) {
	clientID := c.PostForm("client_id")
	clientSecret := c.PostForm("client_secret")

	storedClientSecret, ok := s.clients[clientID]
	if !ok || storedClientSecret != clientSecret {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid_client"})
		return
	}

	// Возвращаем access token (в реальности должен быть JSON Web Token или аналогичным типом токена)
	accessToken := randSeq(20)
	s.accessTokens[accessToken] = clientID

	c.JSON(http.StatusOK, gin.H{"access_token": accessToken, "token_type": "Bearer"})
}

func (s *OAuthServer) ResourceHandler(c *gin.Context) {
	accessToken := c.GetHeader("Authorization")

	if _, ok := s.accessTokens[accessToken]; !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid_token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "secure data"})
}

func randSeq(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func main() {
	rand.Seed(time.Now().UnixNano())

	oauthServer := NewOAuthServer()
	oauthServer.clients["my_client_id"] = "my_client_secret"

	r := gin.Default()

	r.POST("/token", oauthServer.TokenHandler)
	r.GET("/resource", oauthServer.ResourceHandler)

	r.Run(":8080")
}
