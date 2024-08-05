package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"traceroute/helper"
	"traceroute/trace"
)

func Trace(c *gin.Context) {
	logger := helper.InitLogger()

	host := c.Param("host")
	if host == "" {
		logger.Warn("Invalid host")
		c.IndentedJSON(http.StatusBadRequest, gin.H{"ERROR": "Invalid host"})
		return
	}

	ipAddr, err := trace.ResolveIP(host)
	if err != nil {
		logger.Error("IP resolution failed", "Host", host)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"ERROR": "Failed to resolve IP address"})
		return
	}

	logger.Info("performing Trace")
	traceResponses := trace.PerformTrace(ipAddr)

	c.IndentedJSON(http.StatusOK, traceResponses)

	helper.StoreResults(host, traceResponses)
}
