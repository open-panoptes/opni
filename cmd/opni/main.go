package main

import (
	"github.com/gin-gonic/gin"
	"github.com/open-panoptes/opni/pkg/opni"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	opni.Execute()
}
