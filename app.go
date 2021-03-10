package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

const (
	addNumbersUsageMsg      = "USAGE: http://localhost/add/[int]/[int]"
	subtractNumbersUsageMsg = "USAGE: http://localhost/subtract/[int]/[int]"
)

func StartApp() {
	setRoutes()
	if err := router.Run(":1060"); err != nil {
		panic(err)
	}
}

func AddNumbers(c *gin.Context) {
	x, err := strconv.Atoi(c.Param("strX"))
	if err != nil {
		c.String(http.StatusBadRequest, addNumbersUsageMsg)
		return
	}
	y, err := strconv.Atoi(c.Param("strY"))
	if err != nil {
		c.String(http.StatusBadRequest, addNumbersUsageMsg)
		return
	}

	c.String(http.StatusOK, strconv.Itoa(x+y))
}

func SubtractNumbers(c *gin.Context) {
	x, err := strconv.Atoi(c.Param("strX"))
	if err != nil {
		c.String(http.StatusBadRequest, subtractNumbersUsageMsg)
		return
	}
	y, err := strconv.Atoi(c.Param("strY"))
	if err != nil {
		c.String(http.StatusBadRequest, subtractNumbersUsageMsg)
		return
	}

	c.String(http.StatusOK, strconv.Itoa(x-y))
}
