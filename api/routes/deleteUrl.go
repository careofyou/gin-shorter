package routes

import (
	"net/http"

	"github.com/careofyou/url-short/api/database"
	"github.com/gin-gonic/gin"
)


func DeleteURL(c *gin.Context) {
    shortID := c.Param("shortID")

    r := database.CreateClient(0)
    defer r.Close()

    err := r.Del(database.Ctx, shortID).Err()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": "Delete shorten link error",
        })
        return
    }
    c.JSON(http.StatusOK, gin.H{
        "message": "Succesfully deleted",
    })
}
