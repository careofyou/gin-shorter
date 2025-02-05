package routes

import (
	"net/http"

	"github.com/careofyou/url-short/api/database"
	"github.com/gin-gonic/gin"
)

func GetShortByID(c *gin.Context) {
    shortID := c.Param("shortID")

    r := database.CreateClient(0)
    defer r.Close()

    val, err := r.Get(database.Ctx, shortID).Result()
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{
            "error": "Data doesnt exist for providen ShortID",
        })
        return 
    }

    c.JSON(http.StatusOK, gin.H{"data": val})
}
