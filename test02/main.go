package main

import (
	"github.com/gin-gonic/gin"
    "os"
)

func main() {
    r := gin.Default()

    r.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "pong",
        })
    })

    r.GET("/env", func(c *gin.Context) {
        pod_name := os.Getenv("MY_POD_NAME")
        pod_namespace := os.Getenv("MY_POD_NAMESPACE")
        c.JSON(200, gin.H{
            "pod_name": pod_name,
            "pod_namespace": pod_namespace,
        })
    })


    r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

