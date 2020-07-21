package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type errorResponse struct {
	Error string `json:"error"`
}

type addrResponse struct {
	Addr    string    `json:"addr"`
	Records []*record `json:"records"`
}

func lookupAddr(c *gin.Context) {
	obj, ok := c.Get("db")
	if !ok {
		panic("Fail to get database instance in lookupAddr")
	}
	db := obj.(database)

	addr := c.Param("addr")
	records, err := db.getName(addr)
	if err != nil {
		c.JSON(500, errorResponse{Error: err.Error()})
		return
	}

	c.JSON(200, addrResponse{
		Addr:    addr,
		Records: records,
	})
	return
}

func runServer(bindAddr string, bindPort int, db database) error {
	r := gin.Default()

	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})
	// API route group
	r.GET("/addr/:addr", lookupAddr)

	// Start server
	if err := r.Run(fmt.Sprintf("%s:%d", bindAddr, bindPort)); err != nil {
		return err
	}

	return nil
}
