package workapi

import (
	"fmt"
)

// CreateDepartment 创建企业微信部门信息
func (wx *WorkAPI) CreateDepartment(d Department) ([]byte, error) {
	token, err := wx.GetTokenString()
	if err != nil {
		return nil, err
	}

	url := BaseURL + fmt.Sprintf(WorkAPIType["DepartmentCreate"][0], token)

	data, err := DataToReader(d)
	if err != nil {
		return nil, err
	}

	body, err := HTTPRequest(url, WorkAPIType["DepartmentCreate"][1], data)
	if err != nil {
		return nil, err
	}

	return body, nil
}

// UpdateDepartment 更新企业微信部门信息
func (wx *WorkAPI) UpdateDepartment(d Department) ([]byte, error) {
	token, err := wx.GetTokenString()
	if err != nil {
		return nil, err
	}

	url := BaseURL + fmt.Sprintf(WorkAPIType["DepartmentUpdate"][0], token)

	data, err := DataToReader(d)
	if err != nil {
		return nil, err
	}

	body, err := HTTPRequest(url, WorkAPIType["DepartmentUpdate"][1], data)
	if err != nil {
		return nil, err
	}

	return body, nil
}

// DeleteDepartment 删除企业微信部门
func (wx *WorkAPI) DeleteDepartment(id int) ([]byte, error) {
	token, err := wx.GetTokenString()
	if err != nil {
		return nil, err
	}

	url := BaseURL + fmt.Sprintf(WorkAPIType["DepartmentList"][0], token, id)

	body, err := HTTPRequest(url, WorkAPIType["DepartmentList"][1], nil)
	if err != nil {
		return nil, err
	}

	return body, nil
}

// GetDepartmentList 获取企业微信部门列表
func (wx *WorkAPI) GetDepartmentList(id int) ([]byte, error) {
	token, err := wx.GetTokenString()
	if err != nil {
		return nil, err
	}

	url := BaseURL + fmt.Sprintf(WorkAPIType["DepartmentList"][0], token)
	if id != 0 {
		url += fmt.Sprintf("&id=%v", id)
	}

	body, err := HTTPRequest(url, WorkAPIType["DepartmentList"][1], nil)
	if err != nil {
		return nil, err
	}

	return body, nil
}
