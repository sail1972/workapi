package workapi

// ErrResult 企业微信调用成功与否的返回信息
type ErrResult struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

// Department 企业微信部门信息模型
type Department struct {
	ID       int    `json:"id,omitempty"`
	Name     string `json:"name"`
	ParentID int32  `json:"parentid"`
	Order    int32  `json:"order,omitempty"`
}

// Departments 企业微信部门调用返回信息
type Departments struct {
	*ErrResult
	DepartmentList []Department `json:"department"`
}

// Users 企业微信用户调用返回信息
type Users struct {
	*ErrResult
	Userlist []User `json:"userlist"`
}

// User 企业微信用户信息模型
type User struct {
	UserID      string `json:"userid"`
	Name        string `json:"name,omitempty"`
	Department  []int  `json:"department,omitempty"`
	Mobile      string `json:"mobile,omitempty"`
	Email       string `json:"email,omitempty"`
	Status      int    `json:"status,omitempty"`
	Avatar      string `json:"avatar,omitempty"`
	Telephone   string `json:"telephone,omitempty"`
	EnglishName string `json:"english_name,omitempty"`
}

// Tag 企业微信标签信息模型
type Tag struct {
	TagName string `json:"tagname"`
	TagID   int    `json:"tagid"`
}

// Tags 企业微信标签调用返回信息
type Tags struct {
	*ErrResult
	TagList []Tag `json:"taglist"`
}

// AllowUserInfos 允许使用该应用的用户组，归属与Agent结构
type AllowUserInfos struct {
	User []User `json:"user"`
}

// AllowTagID 允许使用该应用的标签组，归属于Agent结构
type AllowTagID struct {
	Tagid []int `json:"tagid"`
}

// AllowPartyID 允许使用该应用的部门组，归属于Agent结构
type AllowPartyID struct {
	PartyID []int `json:"partyid"`
}

// Agent 应用详细信息
type Agent struct {
	AgentID        int             `json:"agentid"`
	Name           string          `json:"name"`
	SquareLogoURL  string          `json:"square_logo_url,omitempty"`
	Description    string          `json:"desription,omitempty"`
	AllowUserInfos *AllowUserInfos `json:"allow_userinfos"`
	AllowParty     *AllowPartyID   `json:"allow_partys"`
	AllowTags      *AllowTagID     `json:"allow_tags"`
	Close          int             `json:"close,omitempty"`
	RedirecDomain  string          `json:"redirect_domain,omitempty"`
	ReportLocation int             `json:"report_location,omitempty"`
	IsReportEnter  int             `json:"isreportenter,omitempty"`
	HomeURL        string          `json:"home_url,omitempty"`
}
