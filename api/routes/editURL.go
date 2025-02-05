package routes

import (
	"net/http"
	"time"

	"github.com/careofyou/url-short/api/database"
	"github.com/careofyou/url-short/api/models"
	"github.com/gin-gonic/gin"
)

func EditURL(c *gin.Context) {
    shortID := c.Param("shortID")
    var body models.Request

    if err := c.ShouldBind(&body); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "Parsing JSON error",
        })
    }

    r := database.CreateClient(0)
    defer r.Close()

    // checking weather shortID is exists in DB -- or not
    val, err := r.Get(database.Ctx, shortID).Result()
    if err != nil || val == "" {
        c.JSON(http.StatusNotFound, gin.H{
            "error": "shortID doesnt exist",
        })
    }

    // update the content of the URL and expitry time with the shortID
    err = r.Set(database.Ctx, shortID, body.URL, body.Expiry*3600*time.Second).Err()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": "Unable to updated the shorten link",
        })
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message": "Succesfully updated",
    })
}

