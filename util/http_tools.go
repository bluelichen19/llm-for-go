package util

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type RalLLMParams struct {
	Msg string `json:"msg"`
	ChatName string `json:"chat_name"`
}

type RalLLMResp struct {
	Msg string `json:"msg"`
}
// offline
//const LLLMUrl string = "http://127.0.0.1:9866/v2/apptools/llm/qianfan"
//const LLLMUrlTest string = "http://127.0.0.1:9866/v2/apptools/llm/qianfantest"

// online
const LLLMUrl string = "http://lmqfeanp.wx-agent.hjdi17hc.p3648brp.com:9866/v2/apptools/llm/qianfan"
const LLLMUrlTest string = "http://lmqfeanp.wx-agent.hjdi17hc.p3648brp.com:9866/v2/apptools/llm/qianfantest"


func RalLLM(params RalLLMParams) (string, error){
	input := RalLLMParams{
		ChatName: params.ChatName,
		Msg: params.Msg,
	}
	output := RalLLMResp{}
	jsStr, _ := json.Marshal(input)
	payload := strings.NewReader(string(jsStr))
	fmt.Println(string(jsStr))
	client := &http.Client {
		Timeout: 15 * time.Second,
	}
	req, err := http.NewRequest("POST", LLLMUrl, payload)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	req.Header.Add("Content-Type", "application/json")
	//req.Header.Add("Authorization", "dee842e459ef91f78546e8222e867b3dab25eaf612edc962722320a09c9787d2")
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	fmt.Println("模型回复：",string(body))
	err = json.Unmarshal(body, &output)
	//js, _ := json.Marshal(output.Msg)
	return output.Msg, nil
}

func RalLLMTest(params RalLLMParams) (string, error){
	input := RalLLMParams{
		ChatName: params.ChatName,
		Msg: params.Msg,
	}
	output := RalLLMResp{}
	jsStr, _ := json.Marshal(input)
	payload := strings.NewReader(string(jsStr))
	fmt.Println(string(jsStr))
	client := &http.Client {
		//Timeout: 15 * time.Second,
	}
	fmt.Println("call agent:", LLLMUrlTest)
	req, err := http.NewRequest("POST", LLLMUrlTest, payload)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	req.Header.Add("Content-Type", "application/json")
	//req.Header.Add("Authorization", "dee842e459ef91f78546e8222e867b3dab25eaf612edc962722320a09c9787d2")
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	fmt.Println("模型回复：",string(body))
	err = json.Unmarshal(body, &output)
	//js, _ := json.Marshal(output.Msg)
	return output.Msg, nil
}
