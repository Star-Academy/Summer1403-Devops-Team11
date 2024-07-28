package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"your_module/trace"
	"your_module/helper"
)

func Trace(c *gin.Context) {
	host := c.Param("host")

	ipAddr, err := trace.ResolveIP(host)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"ERROR": "Failed to resolve IP address"})
		return
	}

	traceResponses := trace.PerformTrace(ipAddr)

	c.IndentedJSON(http.StatusOK, traceResponses)

	helper.CacheTraceResults(host, traceResponses)
}

