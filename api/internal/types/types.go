// Code generated by goctl. DO NOT EDIT.
package types

type AddUserReq struct {
	Name     string `json:"name"`
	NickName string `json:"nickName"`
	Password string `json:"password,optional"`
	Email    string `json:"email"`
	RoleId   int64  `json:"roleId"`
	Status   int64  `json:"status,default=1"`
}

type AddUserResp struct {
	Code    int64           `json:"code"`
	Message string          `json:"message"`
	Data    ReceiptUserData `json:"data"`
}

type ListMenuTree struct {
	Id       int64  `json:"id"`
	Path     string `json:"path"`
	Name     string `json:"name"`
	ParentId int64  `json:"parentId"`
	Icon     string `json:"icon"`
}

type ListMenuTreeVue struct {
	Id           int64        `json:"id"`
	ParentId     int64        `json:"parentId"`
	Title        string       `json:"title"`
	Path         string       `json:"path"`
	Name         string       `json:"name"`
	Icon         string       `json:"icon"`
	VueRedirent  string       `json:"vueRedirent"`
	VueComponent string       `json:"vueComponent"`
	Meta         MenuTreeMeta `json:"meta"`
}

type MenuTreeMeta struct {
	Title string `json:"title"`
	Icon  string `json:"icon"`
}

type ReceiptUserData struct {
	Id int64 `json:"id"`
}

type UserInfoData struct {
	Avatar      string             `json:"avatar"`
	Name        string             `json:"name"`
	MenuTree    []*ListMenuTree    `json:"menuTree"`
	MenuTreeVue []*ListMenuTreeVue `json:"menuTreeVue"`
	ResetPwd    bool               `json:"resetPwd,default=false"`
}

type UserInfoResp struct {
	Code    int64        `json:"code"`
	Message string       `json:"message"`
	Data    UserInfoData `json:"data"`
}
