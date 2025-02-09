package routes

import (
  "interPaste/handlers"
  "github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
  r.POST("/upload", handlers.UploadPaste);
  r.GET("/fetch/:cid", handlers.FetchPaste)
}
