package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/muhuchah/traceroute/trace"
	"github.com/muhuchah/traceroute/helper"
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

