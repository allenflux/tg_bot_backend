package platform

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"io"
	"net/http"
	"tg_bot_backend/internal/consts"
)

type TokenData struct {
	Token string `json:"token"`
}

type APIResponse struct {
	UUID    string    `json:"uuid"`
	Success bool      `json:"success"`
	Code    int       `json:"code"`
	Msg     string    `json:"msg"`
	Data    TokenData `json:"data"`
}

func GetPlatformToken(ctx context.Context, domain, apiUser, secretKey string) (string, error) {
	url := domain + consts.PlatformPathLogin

	payload := map[string]string{
		"account":  apiUser,
		"password": secretKey,
	}

	bytesData, err := json.Marshal(payload)
	if err != nil {
		g.Log().Errorf(ctx, "Failed to marshal login payload: %v", err)
		return "", err
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(bytesData))
	if err != nil {
		g.Log().Errorf(ctx, "HTTP POST request failed: %v", err)
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		g.Log().Errorf(ctx, "Failed to read response body: %v", err)
		return "", err
	}

	if len(body) == 0 {
		err = errors.New("platform API response is empty, please verify your credentials; " + url)
		g.Log().Errorf(ctx, err.Error())
		return "", err
	}

	var platformResp APIResponse
	if err := json.Unmarshal(body, &platformResp); err != nil {
		g.Log().Errorf(ctx, "Failed to unmarshal platform response: %v", err)
		return "", err
	}

	if !platformResp.Success || platformResp.Code != 1000 {
		err := errors.New("platform login failed: " + platformResp.Msg)
		g.Log().Errorf(ctx, err.Error())
		return "", err
	}

	return platformResp.Data.Token, nil
}

type VerifyCustomerResponse struct {
	UUID    string `json:"uuid"`
	Success bool   `json:"success"`
	Code    int    `json:"code"`
	Msg     string `json:"msg"`
	Data    struct {
		Warning              bool    `json:"warning"`
		TotalBalance         float64 `json:"total_balance"`
		TotalRechargeBalance float64 `json:"total_recharge_balance"`
		Count                int     `json:"count"`
		Customers            []any   `json:"customers"` // 你可以根据实际结构定义更具体的类型
	} `json:"data"`
}

func VerifyCustomerId(ctx context.Context, domain, token string, id int) (bool, error) {
	url := domain + consts.PlatformVerifyCustomerId
	payload := map[string]int{
		"id":        id,
		"available": 1,
	}
	bytesData, err := json.Marshal(payload)
	if err != nil {
		g.Log().Errorf(ctx, "Failed to marshal Verify Customer payload: %v", err)
		return false, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(bytesData))
	if err != nil {
		g.Log().Errorf(ctx, " Request creation failed: %s", err.Error())
		return false, err
	}

	// 设置请求头
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		g.Log().Errorf(ctx, " Request failed: %s", err.Error())
		return false, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		g.Log().Errorf(ctx, " Read response failed : %s ", err.Error())
		return false, err
	}

	if len(body) == 0 {
		err = errors.New("platform API response is empty; " + url)
		g.Log().Errorf(ctx, err.Error())
		return false, err
	}
	g.Log().Infof(ctx, "Verify customer id: %d", id)
	g.Log().Infof(ctx, "Verify customer resp : %s", string(body))

	var result VerifyCustomerResponse
	if err = json.Unmarshal(body, &result); err != nil {
		return false, err
	}

	if result.Data.Count == 0 {
		return false, errors.New("platform API response Data is 0; " + url)
	}

	return true, nil
}

type BusinessVerifyResponse struct {
	UUID    string      `json:"uuid"`
	Success bool        `json:"success"`
	Code    int         `json:"code"`
	Msg     string      `json:"msg"`
	Data    interface{} `json:"data"` // 或定义具体结构体也可以
}

func VerifyBusinessId(ctx context.Context, domain, token, id string) (bool, error) {
	url := domain + consts.PlatformVerifyBusinessId + id

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		g.Log().Errorf(ctx, "Request creation failed: %s", err.Error())
		return false, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		g.Log().Errorf(ctx, "Request failed: %s", err.Error())
		return false, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		g.Log().Errorf(ctx, "Read response failed: %s", err.Error())
		return false, err
	}

	if len(body) == 0 {
		err = errors.New("platform API response is empty: " + url)
		g.Log().Errorf(ctx, err.Error())
		return false, err
	}

	g.Log().Infof(ctx, "Verify business id: %s", id)
	g.Log().Infof(ctx, "Verify business resp: %s", string(body))

	var result BusinessVerifyResponse
	if err = json.Unmarshal(body, &result); err != nil {
		g.Log().Errorf(ctx, "Unmarshal failed: %s", err.Error())
		return false, err
	}

	if result.Success && result.Code == 1000 {
		return true, nil
	}

	// 失败时返回错误信息
	return false, errors.New(fmt.Sprintf("verify business failed: code=%d, msg=%s", result.Code, result.Msg))
}

type CustomerFindAPIResponse struct {
	UUID    string `json:"uuid"`
	Success bool   `json:"success"`
	Code    int    `json:"code"`
	Msg     string `json:"msg"`
	Data    struct {
		ID            int     `json:"id"`
		CustomerID    int     `json:"customer_id"`
		CustomerName  string  `json:"customer_name"`
		OrderID       string  `json:"order_id"`
		CustomerOrder string  `json:"customer_order"`
		BusinessType  int     `json:"business_type"`
		Amount        float64 `json:"amount"`
		Balance       float64 `json:"balance"`
		TotalBalance  float64 `json:"total_balance"`
		Bank          string  `json:"bank"`
		Remark        string  `json:"remark"`
		Created       string  `json:"created"`
		Updated       string  `json:"updated"`
	} `json:"data"`
}

func AddCustomerFind(ctx context.Context, domain, token string, customerId, amount int) (bool, error) {
	url := domain + consts.PlatformAddCustomerFind
	payload := map[string]interface{}{
		"customer_id": customerId,
		"bank":        "-1",
		"amount":      amount,
		"remark":      "call api by tg bot",
		"type":        1,
	}
	bytesData, err := json.Marshal(payload)
	if err != nil {
		g.Log().Errorf(ctx, "Failed to marshal AddCustomerFind payload: %v", err)
		return false, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(bytesData))
	if err != nil {
		g.Log().Errorf(ctx, " Request creation failed: %s", err.Error())
		return false, err
	}

	// 设置请求头
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		g.Log().Errorf(ctx, " Request failed: %s", err.Error())
		return false, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		g.Log().Errorf(ctx, " Read response failed : %s ", err.Error())
		return false, err
	}

	if len(body) == 0 {
		err = errors.New("platform API response is empty; " + url)
		g.Log().Errorf(ctx, err.Error())
		return false, err
	}
	g.Log().Infof(ctx, "Verify CustomerFind resp : %s", string(body))

	var result CustomerFindAPIResponse
	if err = json.Unmarshal(body, &result); err != nil {
		return false, fmt.Errorf("failed to parse API response: %w", err)
	}

	if !result.Success {
		return false, fmt.Errorf("API call failed: code=%d, msg=%s, url=%s", result.Code, result.Msg, url)
	}

	if result.Data.ID == 0 {
		return false, fmt.Errorf("API returned success, but data is empty or invalid; url=%s", url)
	}

	return true, nil
}
