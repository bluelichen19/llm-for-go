/*
###############################################
Copyright (c) 2023 Baidu, Inc. All Rights Reserved
# Author        :  lichen18@baidu.com
# Organization  :  Baidu-inc
# Created Time  : 2024-08-09 15:19:34
# File Name     : middleware/debug.go
# Description   :
###############################################
*/
package middleware

import (
    "bytes"
	"fmt"
	"io/ioutil"
	"log"

    "github.com/gin-gonic/gin"
)

func CreateOrUpdateInfluencer() gin.HandlerFunc {
	return func(c *gin.Context) {
		body, _ := ioutil.ReadAll(c.Request.Body)
		log.Printf("%+v\n", string(body))
		for k, v := range c.Request.Header {
			fmt.Println(k, v)
		}
		c.Request.Body = ioutil.NopCloser(bytes.NewReader(body))
		c.Next()
	}
}
