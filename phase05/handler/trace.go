package handler

import (
	"net/http"
	"fmt"

	"github.com/gin-gonic/gin"

	"traceroute/helper"
	"traceroute/trace"
)

func Trace(c *gin.Context) {
	host := c.Param("host")
	if host == "" {
		fmt.Println("[Trace handler] [Error]: Invalid host to trace!")
		c.IndentedJSON(http.StatusBadRequest, gin.H{"ERROR": "Invalid host"})
		return
	}

	ipAddr, err := trace.ResolveIP(host)
	if err != nil {
		fmt.Println("[Trace handler] [Error]: IP resolution failed")
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"ERROR": "Failed to resolve IP address"})
		return
	}

	fmt.Println("[Trace handler] [Info]: Performing trace")
	traceResponses := trace.PerformTrace(ipAddr)

	c.IndentedJSON(http.StatusOK, traceResponses)

	fmt.Println("[Trace handler] [Info]: Storing results")
	helper.StoreResults(host, traceResponses)
}
