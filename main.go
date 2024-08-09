/*
###############################################
Copyright (c) 2023 Baidu, Inc. All Rights Reserved
# Author        :  lichen18@baidu.com
# Organization  :  Baidu-inc
# Created Time  : 2024-08-09 12:34:46
# File Name     : ./main.go
# Description   :
###############################################
*/
package main

import (
	"llm-for-go/routes"
)

func main() {
	// Our server will live in the routes package
	//select {}
	routes.Run()
}

