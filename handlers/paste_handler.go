package handlers

import (
	"interPaste/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UploadPaste(c *gin.Context){
  text := c.PostForm("text")
  cid, err := services.UploadToIPFS(text)
  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error":err.Error()})
  }

  c.JSON(http.StatusOK, gin.H{"cid":cid})
}

func FetchPaste(c *gin.Context){
  cid := c.Param("cid")

  text, err := services.FetchFromIPFS(cid)
  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
  }

  c.JSON(http.StatusOK, gin.H{"text":text})
}
