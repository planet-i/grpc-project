package models

type TokenValue struct {
	Token     string   `json:"token" `    // token
	UserID    string   `json:"userId" `   // 用户ID
	UserName  string   `json:"userName" ` // 用户名
	NickName  string   `json:"nickName" ` // 昵称
	Auth      []string `json:"auth" `     // 权限码
	LoginTime string   `json:"loginTime" `
	RootIDs   string   `json:"root_ids" `
	UserToken string   `json:"user_token" `
}

type UserIDValue struct {
	UsefulToken []string `json:"usefulToken" `
	UserToken   string   `json:"user_token" `
}
