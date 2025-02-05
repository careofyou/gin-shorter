package routes

import (
	"encoding/json"
	"net/http"

	"github.com/careofyou/url-short/api/database"
	"github.com/gin-gonic/gin"
)

type TagRequest struct {
    ShortID string `json:"shortID"`
    Tag string `json:"tag"`
}

func AddTag(c *gin.Context) {
    var tagRequest TagRequest
    if err := c.ShouldBindJSON(&tagRequest); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "Invalid request body",
        })
        return
    }

    shortId := tagRequest.ShortID
    tag := tagRequest.Tag

    r := database.CreateClient(0)
    defer r.Close()

    val, err := r.Get(database.Ctx, shortId).Result()
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{
            "error": "No data for the providen ShortID",
        })
        return
    }

    var data map[string]interface{} 
    if err := json.Unmarshal([]byte(val), &data); err != nil {
        // if data isnt a JSON obj, assume it as a plain string
        data = make(map[string]interface{})
        data["data"] = val
    }

    // checl if "tags" field already exists and its a slice of strings
    var tags []string
    if existingTags, ok := data["tags"].([]interface{}); ok{
        for _, t := range existingTags{
            if strTag, ok := t.(string); ok{
                tags = append(tags, strTag)
            }
        }
    }
    
    // check for the duplicated tags
    for _, existingTags := range tags {
        if existingTags == tag {
            c.JSON(http.StatusBadRequest, gin.H{
                "error": "This tag already exists",
            })
            return 
        }
    }

    // Add the new tag to the slice of tags
    tags = append(tags, tag)
    data["tags"] = tags


    // Marshaling the updated data back to JSOn
    updatedData, err := json.Marshal(data)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": "Failed to Marshal updated data",
        })
        return
    }

    err = r.Set(database.Ctx, shortId, updatedData, 0).Err()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": "Failed to update the database",
        })
        return
    }

    // Response with the updated data
    c.JSON(http.StatusOK, data)
}

