package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/missions", getMissions)

	router.Run("localhost:8080")
}

type Status string
const (
    Created Status = "created"
    Done Status = "done"
    Expired Status =  "expired"
)

type Mission struct {
	Uuid         string `json:"uuid"`
	Title        string `json:"title"`
	Description  string `json:"description"`
	Rewards      string `json:"rewards"`
	Status       Status `json:"status"`
	DateCreation string `json:"dateCreation"`
}


var missions = []Mission{
	{Uuid: "1", Title: "lave la vaisselle feignasse", Description: "description", Rewards: "rewards", Status: Done, DateCreation: "2022/10/22"},
	{Uuid: "2", Title: "lave la vaisselle feignasse 2", Description: "description", Rewards: "rewards", Status: "status", DateCreation: "2022/10/22"},
}

func getMissions(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, missions)
}
