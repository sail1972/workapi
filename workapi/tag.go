package workapi

import (
	"fmt"
)

// CreateTag 创建企业微信标签
func (wx *WorkAPI) CreateTag(tag Tag) ([]byte, error) {
	token, err := wx.GetTokenString()
	if err != nil {
		return nil, err
	}

	url := BaseURL + fmt.Sprintf(WorkAPIType["TagCreate"][0], token)

	data, err := DataToReader(tag)
	if err != nil {
		return nil, err
	}

	body, err := HTTPRequest(url, WorkAPIType["TagCreate"][1], data)
	if err != nil {
		return nil, err
	}
	return body, nil
}

// UpdateTag 更新企业微信标签名称
func (wx *WorkAPI) UpdateTag(tag Tag) ([]byte, error) {
	token, err := wx.GetTokenString()
	if err != nil {
		return nil, err
	}

	url := BaseURL + fmt.Sprintf(WorkAPIType["TagUpdate"][0], token)

	data, err := DataToReader(tag)
	if err != nil {
		return nil, err
	}

	body, err := HTTPRequest(url, WorkAPIType["TagUpdate"][1], data)
	if err != nil {
		return nil, err
	}
	return body, nil
}

// DeleteTag 删除企业微信标签
func (wx *WorkAPI) DeleteTag(tagID int) ([]byte, error) {
	token, err := wx.GetTokenString()
	if err != nil {
		return nil, err
	}

	url := BaseURL + fmt.Sprintf(WorkAPIType["TagDelete"][0], token, tagID)

	body, err := HTTPRequest(url, WorkAPIType["TagDelete"][1], nil)
	if err != nil {
		return nil, err
	}
	return body, nil
}

// GetTagList 获取企业微信标签列表
func (wx *WorkAPI) GetTagList() ([]byte, error) {
	token, err := wx.GetTokenString()
	if err != nil {
		return nil, err
	}

	url := BaseURL + fmt.Sprintf(WorkAPIType["TagGetList"][0], token)

	body, err := HTTPRequest(url, WorkAPIType["TagGetList"][1], nil)
	if err != nil {
		return nil, err
	}

	return body, nil
}

// AddTagUser 添加企业微信标签用户或部门
func (wx *WorkAPI) AddTagUser(tagID int, userID []string, partyID []int) ([]byte, error) {
	token, err := wx.GetTokenString()
	if err != nil {
		return nil, err
	}

	url := BaseURL + fmt.Sprintf(WorkAPIType["TagAddUser"][0], token)

	d := map[string]interface{}{"tagid": tagID, "userlist": userID, "partylist": partyID}

	data, err := DataToReader(d)
	if err != nil {
		return nil, err
	}

	body, err := HTTPRequest(url, WorkAPIType["TagAddUser"][1], data)
	if err != nil {
		return nil, err
	}

	return body, nil
}

// DeleteTagUser 添加企业微信标签用户或部门
func (wx *WorkAPI) DeleteTagUser(tagID int, userID []string, partyID []int) ([]byte, error) {
	token, err := wx.GetTokenString()
	if err != nil {
		return nil, err
	}

	url := BaseURL + fmt.Sprintf(WorkAPIType["TagDeleteUser"][0], token)

	d := map[string]interface{}{"tagid": tagID, "userlist": userID, "partylist": partyID}

	data, err := DataToReader(d)
	if err != nil {
		return nil, err
	}

	body, err := HTTPRequest(url, WorkAPIType["TagDeleteUser"][1], data)
	if err != nil {
		return nil, err
	}

	return body, nil
}

// GetTagUser 获取企业微信标签用户
func (wx *WorkAPI) GetTagUser(tagID int) ([]byte, error) {
	token, err := wx.GetTokenString()
	if err != nil {
		return nil, err
	}

	url := BaseURL + fmt.Sprintf(WorkAPIType["TagGetUser"][0], token, tagID)

	body, err := HTTPRequest(url, WorkAPIType["TagGetUser"][1], nil)
	if err != nil {
		return nil, err
	}

	return body, nil
}
