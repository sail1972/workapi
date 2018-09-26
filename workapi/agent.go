package workapi

import (
	"fmt"
)

// GetAgent 通过agentid获取企业微信应用信息
func (wx *WorkAPI) GetAgent(agentID int) ([]byte, error) {
	token, err := wx.GetTokenString()
	if err != nil {
		return nil, err
	}

	url := BaseURL + fmt.Sprintf(WorkAPIType["AgentGet"][0], token, agentID)

	body, err := HTTPRequest(url, WorkAPIType["AgentGet"][1], nil)
	if err != nil {
		return nil, err
	}

	return body, nil
}

// GetAgentList 通过agentid获取企业微信应用信息
func (wx *WorkAPI) GetAgentList() ([]byte, error) {
	token, err := wx.GetTokenString()
	if err != nil {
		return nil, err
	}

	url := BaseURL + fmt.Sprintf(WorkAPIType["AgentGetList"][0], token)

	body, err := HTTPRequest(url, WorkAPIType["AgentGetList"][1], nil)
	if err != nil {
		return nil, err
	}

	return body, nil
}

// SetAgent 设置企业微信应用信息
func (wx *WorkAPI) SetAgent(agent Agent) ([]byte, error) {
	token, err := wx.GetTokenString()
	if err != nil {
		return nil, err
	}

	url := BaseURL + fmt.Sprintf(WorkAPIType["AgentSet"][0], token)

	data, err := DataToReader(agent)
	if err != nil {
		return nil, err
	}

	body, err := HTTPRequest(url, WorkAPIType["AgentSet"][1], data)
	if err != nil {
		return nil, err
	}

	return body, nil
}
