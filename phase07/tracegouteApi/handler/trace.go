package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"traceroute/helper"
	"traceroute/trace"
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

	helper.StoreResults(host, traceResponses)
}
