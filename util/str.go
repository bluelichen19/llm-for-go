/*
###############################################
Copyright (c) 2023 Baidu, Inc. All Rights Reserved
# Author        :  lichen18@baidu.com
# Organization  :  Baidu-inc
# Created Time  : 2024-08-19 22:28:54
# File Name     : util/str.go
# Description   :
###############################################
*/
package util

import (
        "crypto/md5"
        "encoding/hex"
        "errors"
        "regexp"

        //"fmt"
        "strconv"
        "strings"
)

// EscapeUnicode 字符转码成unicode编码
func EscapeUnicode(text string) string {
        unicodeText := strconv.QuoteToASCII(text)
        // 去掉返回内容两端多余的双引号
        return unicodeText[1 : len(unicodeText)-1]
}

// UnescapeUnicode 将unicode编码转换成中文
func UnescapeUnicode(uContent string) (string, error) {
        // 转码前需要先增加上双引号，Quote增加双引号会将\u转义成\\u，同时会处理一些非打印字符
        content := strings.Replace(strconv.Quote(uContent), `\\u`, `\u`, -1)
        text, err := strconv.Unquote(content)
        if err != nil {
                return "", err
        }
        return text, nil
}

func GetMD5Hash(text string) string {
        hasher := md5.New()
        hasher.Write([]byte(text))
        return hex.EncodeToString(hasher.Sum(nil))
}

func ReverseStringSlice(s []string) {
        for i := 0; i < len(s)/2; i++ {
                j := len(s) - i - 1
                s[i], s[j] = s[j], s[i]
        }
}

func ReplaceStringByRegex(str, rule, replace string) (string, error) {
        reg, err := regexp.Compile(rule)
        if reg == nil || err != nil {
                return str, errors.New("正则MustCompile错误:" + err.Error())
        }
        return reg.ReplaceAllString(str, replace), nil
}

