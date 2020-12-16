package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

type Device struct {
	Name string `json:"name"`
	Host string `json:"host"`
	Info map[string]interface{} `json:"info"`
	Time time.Time `json:"time"`
}

func main() {
	// In memory store of registered devices, key -
	devices := make(map[string]Device, 0)

	bindAddress := "0.0.0.0:8080"
	flag.StringVar(&bindAddress, "bind", bindAddress, "example: 0.0.0.0:8080")
	flag.Parse()
	r := gin.Default()

	r.GET("/device", func(c *gin.Context) {
		c.JSON(200, devices)
	})

	r.POST("/device", func(c *gin.Context) {
		device := Device{}
		err := c.BindJSON(&device)
		if err != nil {
			_ = c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("binding failed: %v", err))
		}
		device.Time = time.Now()
		device.Host = strings.Split(c.Request.RemoteAddr, ":")[0]
		devices[device.Host] = device
		c.Status(201)
	})

	// listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	if err := r.Run(bindAddress); err != nil {
		panic(err)
	}
}
