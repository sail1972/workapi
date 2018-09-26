package workapi

import (
	"fmt"
)

// CreateUser 创建企业微信用户信息
func (wx *WorkAPI) CreateUser(user User) ([]byte, error) {
	token, err := wx.GetTokenString()
	if err != nil {
		return nil, err
	}

	url := BaseURL + fmt.Sprintf(WorkAPIType["UserCreate"][0], token)

	data, err := DataToReader(user)
	if err != nil {
		return nil, err
	}

	body, err := HTTPRequest(url, WorkAPIType["UserCreate"][1], data)
	if err != nil {
		return nil, err
	}

	return body, nil
}

// UpdateUser 更新企业微信用户信息
func (wx *WorkAPI) UpdateUser(user User) ([]byte, error) {
	token, err := wx.GetTokenString()
	if err != nil {
		return nil, err
	}

	url := BaseURL + fmt.Sprintf(WorkAPIType["UserUpdate"][0], token)

	data, err := DataToReader(user)
	if err != nil {
		return nil, err
	}

	body, err := HTTPRequest(url, WorkAPIType["UserUpdate"][1], data)
	if err != nil {
		return nil, err
	}

	return body, nil
}

// DeleteUser 删除企业微信用户信息
func (wx *WorkAPI) DeleteUser(userID string) ([]byte, error) {
	token, err := wx.GetTokenString()
	if err != nil {
		return nil, err
	}

	url := BaseURL + fmt.Sprintf(WorkAPIType["UserDelete"][0], token, userID)

	body, err := HTTPRequest(url, WorkAPIType["UserDelete"][1], nil)
	if err != nil {
		return nil, err
	}

	return body, nil
}

// BatchDeleteUser 删除企业微信用户信息
func (wx *WorkAPI) BatchDeleteUser(userIDs ...string) ([]byte, error) {
	token, err := wx.GetTokenString()
	if err != nil {
		return nil, err
	}

	var url = BaseURL + fmt.Sprintf(WorkAPIType["UserBatchDelete"][0], token)

	d := map[string]interface{}{"userlist": userIDs}

	data, err := DataToReader(d)
	if err != nil {
		return nil, err
	}

	body, err := HTTPRequest(url, WorkAPIType["UserBatchDelete"][1], data)
	if err != nil {
		return nil, err
	}

	return body, nil
}

// GetUser 获取企业微信用户详情
func (wx *WorkAPI) GetUser(userID string) ([]byte, error) {
	token, err := wx.GetTokenString()
	if err != nil {
		return nil, err
	}

	var url = BaseURL + fmt.Sprintf(WorkAPIType["UserGet"][0], token, userID)

	body, err := HTTPRequest(url, WorkAPIType["UserGet"][1], nil)
	if err != nil {
		return nil, err
	}

	return body, nil
}

// GetUserList 获取企业微信用户详情列表
func (wx *WorkAPI) GetUserList(deptID int, fetchChild int) ([]byte, error) {
	token, err := wx.GetTokenString()
	if err != nil {
		return nil, err
	}

	var url = BaseURL + fmt.Sprintf(WorkAPIType["UserList"][0], token, deptID, fetchChild)

	body, err := HTTPRequest(url, WorkAPIType["UserList"][1], nil)
	if err != nil {
		return nil, err
	}

	return body, nil
}

// GetUserSimpleList 获取企业微信用户详情列表
func (wx *WorkAPI) GetUserSimpleList(deptID int, fetchChild int) ([]byte, error) {
	token, err := wx.GetTokenString()
	if err != nil {
		return nil, err
	}

	var url = BaseURL + fmt.Sprintf(WorkAPIType["UserSimpleList"][0], token, deptID, fetchChild)

	body, err := HTTPRequest(url, WorkAPIType["UserSimpleList"][1], nil)
	if err != nil {
		return nil, err
	}

	return body, nil
}

// UserIDToOpenID 企业微信用户UserID转换为OpenID
func (wx *WorkAPI) UserIDToOpenID(userID string) ([]byte, error) {
	token, err := wx.GetTokenString()
	if err != nil {
		return nil, err
	}

	url := BaseURL + fmt.Sprintf(WorkAPIType["UserIDToOpenID"][0], token)

	d := map[string]interface{}{"userid": userID}

	data, err := DataToReader(d)
	if err != nil {
		return nil, err
	}

	body, err := HTTPRequest(url, WorkAPIType["UserIDToOpenID"][1], data)
	if err != nil {
		return nil, err
	}

	return body, nil
}

// OpenIDToUserID 企业微信用户UserID转换为OpenID
func (wx *WorkAPI) OpenIDToUserID(openID string) ([]byte, error) {
	token, err := wx.GetTokenString()
	if err != nil {
		return nil, err
	}

	url := BaseURL + fmt.Sprintf(WorkAPIType["OpenIDToUserID"][0], token)

	d := map[string]interface{}{"openid": openID}

	data, err := DataToReader(d)
	if err != nil {
		return nil, err
	}

	body, err := HTTPRequest(url, WorkAPIType["OpenIDToUserID"][1], data)
	if err != nil {
		return nil, err
	}

	return body, nil
}
