// Copyright (c) 2021 Satvik Reddy
package http

import (
	"net/http"

	"github.com/SatvikR/liveassist/amnis/domain"
	"github.com/SatvikR/liveassist/omnis"
	"github.com/gin-gonic/gin"
)

type createBody struct {
	Name     string   `json:"name"`
	Keywords []string `json:"keywords"`
}

func create(c *gin.Context) {
	var reqBody createBody
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": omnis.ErrInvalidBody.Error(),
		})
		return
	}

	_uid, exists := c.Get("uid")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "could not retrieve user id from token",
		})
		return
	}

	uid := int(_uid.(int64))

	id, err := domain.Create(reqBody.Name, uid, removeDuplicateKeywords(reqBody.Keywords))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "unable to create channel",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id": id,
	})
}

func removeDuplicateKeywords(keywords []string) []string {
	output := make([]string, 0)

	seen := make(map[string]bool)
	for _, word := range keywords {
		if _, ok := seen[word]; !ok {
			output = append(output, word)
			seen[word] = true
		}
	}
	return output
}

func delete(c *gin.Context) {
	id := c.Param("id")
	err := domain.Delete(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "could not delete channel",
		})
		return
	}
	c.Status(http.StatusOK)
}

func channel(c *gin.Context) {
	id := c.Param("id")
	channel, err := domain.GetChannel(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "could not find channel",
		})
		return
	}

	c.JSON(http.StatusOK, channel)
}

func channels(c *gin.Context) {
	channels, err := domain.GetChannels()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "could not fetch channels",
		})
		return
	}
	c.JSON(http.StatusOK, channels)
}
