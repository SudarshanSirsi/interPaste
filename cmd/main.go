package main

import (
	"interPaste/routes"

	"github.com/gin-gonic/gin"
)

func main() {
  r := gin.Default();
  routes.RegisterRoutes(r)
  r.Run(":3001")
}
