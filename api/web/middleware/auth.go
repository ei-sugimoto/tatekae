package middleware

import (
	"github.com/ei-sugimoto/tatekae/api/pkg"
	"github.com/gin-gonic/gin"
)

func Auth(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		c.Abort()
		return
	}

	token, err := pkg.ParseToken(tokenString)
	if err != nil {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		c.Abort()
		return
	}

	info, err := pkg.GetUserInfo(token)
	if err != nil {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		c.Abort()
		return
	}

	c.Set("id", info.ID)
	c.Set("username", info.Username)
	c.Set("email", info.Email)
	c.Next()

}
