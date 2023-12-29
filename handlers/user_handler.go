package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	user := gin.H{
		"id":       1,
		"name":     "Sasha",
		"lastname": "Naumov",
		"imageurl": "https://sun9-6.userapi.com/impg/imCF4uDEPQe0-Qk7gH1vFdHEDb8j8wsXvYPMsw/5o4uUIRpQgE.jpg?size=648x807&quality=95&sign=610781ddb5e0a086304f6eb837a21382&c_uniq_tag=LL1kIZYnBcmrP8n-1AutAe-ANwiwlVW7B2LdWE45Dys&type=album",
	}
	c.JSON(http.StatusOK, user)
}
