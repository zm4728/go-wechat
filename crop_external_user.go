package wechat

import (
	"fmt"
	"github.com/zm4728/go-wechat/util"
)

const (
	// CropExternal User 外部联系人
	CorpAPIExternalUserGet    = CorpAPI + "externalcontact/get?access_token=%s&external_userid=%s"
)

type (
	ExternalUserInfo struct {
		WxErr          `json:"-"`
		ExternalContact struct{
			ExternalUserid string `json:"external_userid"`
			Name  string `json:"name"`
			Type  int `json:"type"` //1表示该外部联系人是微信用户，2表示该外部联系人是企业微信用户
			Avatar string `json:"avatar"`
			Gender int `json:"gender"` //外部联系人性别 0-未知 1-男性 2-女性
			CorpName string `json:"corp_name"` //外部联系人所在企业的简称，仅当联系人类型是企业微信用户时有此字段
			CorpFullName string `json:"corp_full_name"` //外部联系人所在企业的主体名称，仅当联系人类型是企业微信用户时有此字段
			Unionid string `json:"unionid"`
		} `json:"external_contact"`   //crop externalUserInfo
		FollowUser []FollowUser `json:"follow_user"`
		NextCursor string `json:"next_cursor"` //分页的cursor，当跟进人多于500人时返回
	}

	FollowUser struct {
		Userid string `json:"userid"`
		Remark string `json:"remark"`
		Description string `json:"description"`
		Createtime  int `json:"createtime"`
		AddWay      int `json:"add_way"` //该成员添加此客户的来源
		OperUserid  string `json:"oper_userid"` // 发起添加的userid，如果成员主动添加，为成员的userid；如果是客户主动添加，则为客户的外部联系人userid；如果是内部成员共享/管理员分配，则为对应的成员/管理员userid
		State string `json:"state"` //企业自定义的state参数，用于区分客户具体是通过哪个「联系我」添加
		Tags []FollowUserTags `json:"tags"`
	}
	FollowUserTags struct {
		GroupName string `json:"group_name"`
		TagName   string `json:"tag_name"`
		TagId     string `json:"tag_id"`
		Type      int `json:"type"`
	}
)

// GetExternalUserInfo 从企业号通过userId获取外部用户信息
func (s *Server) GetExternalUserInfo(userId string) (user ExternalUserInfo, err error) {
	url := fmt.Sprintf(CorpAPIExternalUserGet,  s.GetUserAccessToken(), userId)
	if err = util.GetJson(url, &user); err != nil {
		return
	}
	err = user.Error()
	return
}

