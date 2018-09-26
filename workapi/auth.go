package workapi

import (
	"fmt"
)

// GetUserInfoByCode 通过企业微信后台返回的code一次性获取用户信息
func (wx *WorkAPI) GetUserInfoByCode(code string) ([]byte, error) {
	token, err := wx.GetTokenString()
	if err != nil {
		return nil, err
	}

	url := BaseURL + fmt.Sprintf(WorkAPIType["GetUserInfoByCode"][0], token, code)

	body, err := HTTPRequest(url, WorkAPIType["GetUserInfoByCode"][1], nil)
	if err != nil {
		return nil, err
	}

	return body, nil
}

// GetUserDetail 通过企业微信后台返回的user_ticket获取用户详细信息
func (wx *WorkAPI) GetUserDetail(userTicket string) ([]byte, error) {
	token, err := wx.GetTokenString()
	if err != nil {
		return nil, err
	}

	url := BaseURL + fmt.Sprintf(WorkAPIType["GetUserDetail"][0], token)

	d := map[string]interface{}{"user_ticket": userTicket}

	data, err := DataToReader(d)
	if err != nil {
		return nil, err
	}

	body, err := HTTPRequest(url, WorkAPIType["GetUserDetail"][1], data)
	if err != nil {
		return nil, err
	}

	return body, nil
}
