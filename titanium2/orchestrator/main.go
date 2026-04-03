package main
import (
    "github.com/gin-gonic/gin"
    "net/http"
)
func main() {
    r := gin.Default()
    r.GET("/api/v1/status", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{"mesh": "TITANIUM_ACTIVE", "drive": "E:"})
    })
    r.Run(":8080")
}
