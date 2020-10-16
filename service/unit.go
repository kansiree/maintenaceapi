package unit

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
)

type BasicResponse struct {
	Error   int             `json:"error"`
	Message json.RawMessage `json:"message"`
}

func setResSucceed(c *gin.Context, response json.RawMessage) {
	c.Header("Content-Type", "application/json; charset=utf-8")

	c.JSON(0, json.NewEncoder(c.Writer).Encode(BasicResponse{
		0,
		response,
	}))
}
