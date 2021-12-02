package wechat

import "github.com/zm4728/go-wechat/util"

//统一服务消息
const (
	WxUniformMessage = WXAPI + "message/subscribe/send?access_token=" // 订阅消息
)

// Template 模板信息
type WxMessageTemplate struct { //mp_template_msg
	Touser          string `json:"touser"` //用户openid
	TemplateId      string `json:"template_id"` // 模板id
	Page 			string `json:"page"` // 跳转页面 index?foo=bar
	Data            map[string]interface{} `json:"data"` // 模板消息内容
}


//   _d := map[string]interface{}{
//   	"thing1":map[string]interface{}{
//		"value":"说法水电费，水电费水电费，电饭锅电饭锅",
//	},
//   	"character_string2":map[string]interface{}{
//		"value":"c655983323546",
//	},
//	 "phrase3":map[string]interface{}{
//		"value":"待提交资料",
//	 },
//	 "thing4":map[string]interface{}{
//		"value":"请提交下一步资料",
//	 },
//   }
//	err := SendWxTemplate("","","/pages/index/index?id=1","developer",_d)
// SendTemplate 发送模板消息，data通常是map[string]struct{value string,color string}
func (s *Server) SendWxTemplate(to, id, pagepath,state string, data interface{}) *WxErr {
	form := map[string]interface{}{
		"touser":      to,
		"template_id": id,
		"page":			pagepath,
		"miniprogram_state":state,//跳转小程序类型 developer为开发版；trial为体验版；formal为正式版；默认为正式版
		"data":        data,
	}
	ret := new(WxErr)
	err := util.PostJsonPtr(WxUniformMessage+s.GetAccessToken(), form, &ret)
	if err != nil {
		return &WxErr{ErrCode: -1, ErrMsg: err.Error()}
	}

	return ret
}
