package workapi

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/json-iterator/go"

	cache "github.com/patrickmn/go-cache"
)

// WorkAPI 企业微信
type WorkAPI struct {
	CorpID     string
	AppSecret  string
	AppID      int
	TokenCache *cache.Cache
}

// AccessToken 企业微信应用令牌信息
type accessToken struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"` //秒 默认返回是2小时
	ExpireAt    int64  `json:"expireAt"`   //在什么时候过期
	ErrCode     int    `json:"errcode"`
	ErrMsg      string `json:"errmsg"`
}

// BaseURL 企业微信应用接口调用链接
const BaseURL = "https://qyapi.weixin.qq.com"

// WorkAPIType 微信API链接与连接模式
var WorkAPIType = map[string][]string{
	"GetAccessToken":  {"/cgi-bin/gettoken?corpid=%s&corpsecret=%s", "GET"},
	"UserCreaet":      {"/cgi-bin/user/create?access_token=%s", "POST"},
	"UserGet":         {"/cgi-bin/user/get?access_token=%s&userid=%s", "GET"},
	"UserUpdate":      {"/cgi-bin/user/update?access_token=%s", "POST"},
	"UserDelete":      {"/cgi-bin/user/delete?access_token=%s", "GET"},
	"UserBatchDelete": {"/cgi-bin/user/batchdelete?access_token=%s", "POST"},
	"UserSimpleList":  {"/cgi-bin/user/simplelist?access_token=%s&department_id=%v&fetch_child=%v", "GET"},
	"UserList":        {"/cgi-bin/user/list?access_token=%s&department_id=%v&fetch_child=%v", "GET"},
	"UserIDToOpenID":  {"/cgi-bin/user/convert_to_openid?access_token=%s", "POST"},
	"OpenIDToUserID":  {"/cgi-bin/user/convert_to_userid?access_token=%s", "POST"},
	"UserAuthSuccess": {"/cgi-bin/user/authsucc?access_token=%s", "GET"},

	"DepartmentCreaet": {"/cgi-bin/department/create?access_token=%s", "POST"},
	"DepartmentGet":    {"/cgi-bin/department/get?access_token=%s&id=%v", "POST"},
	"DepartmentUpdate": {"/cgi-bin/department/update?access_token=%s", "POST"},
	"DepartmentDelete": {"/cgi-bin/department/delete?access_token=%s&id=%v", "GET"},
	"DepartmentList":   {"/cgi-bin/department/list?access_token=%s", "GET"},

	"TagCreate":     {"/cgi-bin/tag/create?access_token=%s", "POST"},
	"TagUpdate":     {"/cgi-bin/tag/update?access_token=%s", "POST"},
	"TagDelete":     {"/cgi-bin/tag/delete?access_token=%s&tagid=%v", "GET"},
	"TagGetUser":    {"/cgi-bin/tag/get?access_token=%s&tagid=%v", "GET"},
	"TagAddUser":    {"/cgi-bin/tag/addtagusers?access_token=%s", "POST"},
	"TagDeleteUser": {"/cgi-bin/tag/deltagusers?access_token=%s", "POST"},
	"TagGetList":    {"/cgi-bin/tag/list?access_token=%s", "GET"},

	"BatchJobGetResult": {"/cgi-bin/batch/getresult?access_token=%s", "GET"},

	"BatchInvite": {"/cgi-bin/batch/invite?access_token=%s", "POST"},

	"AgentGet":     {"/cgi-bin/agent/get?access_token=%s&agentid=%v", "GET"},
	"AgentSet":     {"/cgi-bin/agent/set?access_token=%s", "POST"},
	"AgentGetList": {"/cgi-bin/agent/list?access_token=%s", "GET"},

	"MenuCreaet": {"/cgi-bin/menu/create?access_token=%s", "POST"}, // TODO
	"MenuGet":    {"/cgi-bin/menu/get?access_token=%s", "GET"},
	"MenuDelete": {"/cgi-bin/menu/delete?access_token=%s", "GET"},

	"MessageSend": {"/cgi-bin/message/send?access_token=%s", "POST"},

	"MediaGet": {"/cgi-bin/media/get?access_token=%s", "GET"},

	"GetUserInfoByCode": {"/cgi-bin/user/getuserinfo?access_token=%s&code=%s", "GET"},
	"GetUserDetail":     {"/cgi-bin/user/getuserdetail?access_token=%s", "POST"},

	"GetTicket":      {"/cgi-bin/ticket/get?access_token=%s", "GET"},
	"GetJSAPITicket": {"/cgi-bin/get_jsapi_ticket?access_token=%s", "GET"},

	"GetCheckinOption": {"/cgi-bin/checkin/getcheckinoption?access_token=%s", "POST"},
	"GetCheckinData":   {"/cgi-bin/checkin/getcheckindata?access_token=%s", "POST"},
	"GetApprovalData":  {"/cgi-bin/corp/getapprovaldata?access_token=%s", "POST"},

	"GetInvoiceInfo":           {"/cgi-bin/card/invoice/reimburse/getinvoiceinfo?access_token=%s", "POST"},
	"UpdateInvoiceStatus":      {"/cgi-bin/card/invoice/reimburse/updateinvoicestatus?access_token=%s", "POST"},
	"BatchUpdateInvoiceStatus": {"/cgi-bin/card/invoice/reimburse/updatestatusbatch?access_token=%s", "POST"},
	"BatchGetInvoiceInfo":      {"/cgi-bin/card/invoice/reimburse/getinvoiceinfobatch?access_token=%s", "POST"},
}

// NewWorkAPI 创建新的企业微信应用接口
func NewWorkAPI(corpid, appsecret string, appid int) *WorkAPI {
	if corpid == "" || appsecret == "" {
		log.Fatal("CorpID or AppSecret is null, please check the config file.")
	}
	//return &WorkAPI{CorpID: corpid, AppSecret: appsecret, AppID: appid, TokenCache: cache.New(7100, 5*time.Second)}
	return &WorkAPI{CorpID: corpid, AppSecret: appsecret, AppID: appid, TokenCache: cache.New(7100, 5*time.Second)}
}

// GetAccessToken 企业微信获取或更新访问令牌 AccessToken 接口
func (wx *WorkAPI) GetAccessToken() {
	for {
		url := BaseURL + fmt.Sprintf(WorkAPIType["GetAccessToken"][0], wx.CorpID, wx.AppSecret)
		result, err := HTTPRequest(url, WorkAPIType["GetAccessToken"][1], nil)
		if err != nil {
			log.Printf("Get access token return error: %v", err)
			return
		}

		//res, err := ioutil.ReadAll(result)
		newAccess := accessToken{}
		err = jsoniter.Unmarshal(result, &newAccess)
		if err != nil {
			log.Printf("Get access token parse JSON error: %s", err)
			return
		}

		if newAccess.ExpiresIn == 0 || newAccess.AccessToken == "" {
			log.Printf("Get access token error code: %v, error message: %s", newAccess.ErrCode, newAccess.ErrMsg)
			time.Sleep(5 * time.Minute)
		}

		wx.TokenCache.Set("token", newAccess, time.Duration(newAccess.ExpiresIn)*time.Second)
		log.Printf("Access token update sucessful: %s, expired time: %v", newAccess.AccessToken, newAccess.ExpiresIn)
		time.Sleep(time.Duration(newAccess.ExpiresIn-100) * time.Second)
	}
}

// GetTokenString 获取企业微信接口调用令牌
func (wx *WorkAPI) GetTokenString() (string, error) {

	token, found := wx.TokenCache.Get("token")
	if found {
		accessToken, ok := token.(accessToken)
		if ok {
			return accessToken.AccessToken, nil
		}
		return "", errors.New("Parse Token string failed")
	}
	return "", errors.New("Get Token string failed")
}
