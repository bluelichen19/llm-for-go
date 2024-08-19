/*
###############################################
Copyright (c) 2023 Baidu, Inc. All Rights Reserved
# Author        :  lichen18@baidu.com
# Organization  :  Baidu-inc
# Created Time  : 2024-08-09 14:54:28
# File Name     : service/llama.go
# Description   :
###############################################
*/
package service

import (
    "fmt"
    "bytes"
    "io/ioutil"
    "encoding/json"
    "net/http"
)

type LlamaService struct {
}

type LlamaBotParams struct {
    UserName string
    ChatName string
    Msg      string
}

type LLamaResp struct {
	ID string `json:"id"`
	Created int `json:"created"`
	Choices []Choices `json:"choices"`
	Model string `json:"model"`
	Object string `json:"object"`
	Usage Usage `json:"usage"`
	FinalOut string `json:"final_out"`
}
type Message struct {
	Role string `json:"role"`
	Content string `json:"content"`
}
type Choices struct {
	Index int `json:"index"`
	FinishReason string `json:"finish_reason"`
	Message Message `json:"message"`
}
type Usage struct {
	PromptTokens int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens int `json:"total_tokens"`
}

type Messages struct {
    Role                string  `json:"role,omitempty"`
    Content             string  `json:"content,omitempty"`
}

type LlamaParams struct {
    Model              string   `json:"model,omitempty"`
    Messages           []Messages `json:"messages,omitempty"`
    Temperature        float32  `json:"temperature,omitempty"`
    Top_p              float32  `json:"top_p,omitempty"`
    Repetition_Penalty float32  `json:"repetition_penalty,omitempty"`
    Stream             bool     `json:"stream,omitempty"`
}

func NewLlamaService() *LlamaService {
    return &LlamaService{}
}

func (service *LlamaService) LlamaBot(params LlamaBotParams ,msg *string) (string) {
    messageses := make([]Messages, 0)
    messageses = append(messageses, Messages{
        Role: "assistant",
        Content: "回答下面的问题:\n" + params.Msg + "\n回答上面的问题，要求回答的字数限制在200字以内",
    })
    llamaParams := LlamaParams {
        Model: "Atom-7B-Chat",
        Messages: messageses,
        Stream: false,
    }
    jsonStr, err := json.Marshal(llamaParams)
    fmt.Println(jsonStr)
    if err != nil {
        fmt.Printf("llamaParams Marshal failed %+v\n", err)
        return ""
    }
    fmt.Println(string(jsonStr))
    url:= "https://api.atomecho.cn/v1/chat/completions"
    req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(jsonStr)))
    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("Authorization", "Bearer sk-8b0636211c1f5dd408c44d94621f8293")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return ""
        // handle error
    }
    defer resp.Body.Close()

    statuscode := resp.StatusCode
    hea := resp.Header
    body, _ := ioutil.ReadAll(resp.Body)
    fmt.Println(string(body))
    res := LLamaResp{}
    json.Unmarshal([]byte(body), &res)
    answer := res.Choices[0].Message.Content
    fmt.Println(statuscode)
    fmt.Println(hea)
    return answer
}
