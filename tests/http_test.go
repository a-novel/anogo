package tests

import (
	"github.com/a-novel/anogo"
	"github.com/gin-gonic/gin"
	"testing"
)

type NumberInst struct {
	Value int `json:"value"`
}

func TestPostJSON(t *testing.T) {
	serv, url := anogo.StartGinTestServer([]func(*gin.Engine){
		func(engine *gin.Engine) {
			engine.POST("/add", func(c *gin.Context) {
				inst := &NumberInst{}

				if err := c.Bind(inst); err != nil {
					c.AbortWithStatusJSON(400, gin.H{"error": "non valid payload"})
					return
				}

				c.JSON(200, gin.H{"value": inst.Value + 1})
			})
		},
	})

	defer serv.Close()

	if err := anogo.TestPostJSON(
		url("/add"),
		map[string]interface{}{"value": 4},
		200,
		map[string]interface{}{"value": float64(5)},
	); err != nil {
		t.Error(err.Error())
	}

	if err := anogo.TestPostJSON(
		url("/add"),
		map[string]interface{}{"value": "hello world"},
		400,
		map[string]interface{}{"error": "non valid payload"},
	); err != nil {
		t.Error(err.Error())
	}
}
