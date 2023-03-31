package wechat

import (
	"github.com/zm4728/go-wechat/util"
)

// MPUniformMessageTemplateGetAll 統一消息模板
const (
	MPUniformMessageTemplateSendMsg = WXAPI + "message/wxopen/template/uniform_send?access_token="
)

// // MpTemplate 统一模板信息
type MpUniformMessageTemplate struct {
	ToUser          string      `json:"touser"`          //用户openid，可以是小程序的openid，也可以是mp_template_msg.appid对应的公众号的openid
	AppId           string      `json:"appid"`           //公众号appid，要求与小程序有绑定且同主体
	TemplateId      string      `json:"template_id"`     //公众号模板id
	Url             string      `json:"url"`             //公众号模板消息所要跳转的url
	MiniProgramId   string      `json:"miniprogramId"`   //id,公众号模板消息所要跳转的小程序，小程序的必须与公众号具有绑定关系
	MiniProgramPath string      `json:"miniprogramPath"` //id,公众号模板消息所要跳转的小程序，小程序的必须与公众号具有绑定关系
	Data            interface{} `json:"data"`
}

// UniformMessage 发送模板消息，data通常是map[string]struct{value string,color string}
func (s *Server) SendUniformMessage(info MpUniformMessageTemplate) *WxErr {
	form := map[string]interface{}{
		"touser": info.ToUser,
		"mp_template_msg": map[string]interface{}{
			"appid":       info.AppId,
			"template_id": info.TemplateId,
			"url":         info.Url,
			"miniprogram": map[string]interface{}{
				"appid":    info.MiniProgramId,
				"pagepath": info.MiniProgramPath,
			},
			"data": info.Data,
		},
	}
	ret := new(WxErr)
	err := util.PostJsonPtr(MPUniformMessageTemplateSendMsg+s.GetAccessToken(), form, &ret)
	if err != nil {
		return &WxErr{ErrCode: -1, ErrMsg: err.Error()}
	}
	return ret
}
