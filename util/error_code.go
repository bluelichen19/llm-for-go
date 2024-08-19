/*
###############################################
Copyright (c) 2023 Baidu, Inc. All Rights Reserved
# Author        :  lichen18@baidu.com
# Organization  :  Baidu-inc
# Created Time  : 2024-08-16 18:08:05
# File Name     : util/error_code.go
# Description   :
###############################################
*/
package util

const (
    APP_TOOLS_OK int = iota
    H_SetRedis_Failed
    H_GetRedis_Failed
)

var ErrCode = map[int]string {
    0: "OK",
    1: "h_set_redis_failed",
    2: "h_met_redis_failed",
}


