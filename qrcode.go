package wechat

import (
	"fmt"
	"github.com/zm4728/go-wechat/util"
)

// 生成公众号带参二维码

const (
	ParamsQrcode  = "qrcode/create?access_token=%s"  // 二维码api
	QR_STR_SCENE = "QR_STR_SCENE"  // 临时二维码
	QR_LIMIT_STR_SCENE = "QR_LIMIT_STR_SCENE" // 永久二维码

)

type QrcodeParamsReq struct {
	ExpireSeconds int `json:"expire_seconds"`  //该二维码有效时间，以秒为单位。 最大不超过2592000（即30天），此字段如果不填，则默认有效期为30秒。
	ActionName   string `json:"action_name"`  // 二维码类型，QR_SCENE为临时的整型参数值，QR_STR_SCENE为临时的字符串参数值，QR_LIMIT_SCENE为永久的整型参数值，QR_LIMIT_STR_SCENE为永久的字符串参数值
	ActionInfo   SenceA `json:"action_info"`   //二维码详细信息
}
type SenceA struct {
	Scene  SenceB `json:"scene"`
}
type SenceB struct {
	SceneId      string `json:"sceneId"`      // 场景值ID，临时二维码时为32位非0整型，永久二维码时最大值为100000（目前参数只支持1--100000）
	SceneStr     map[string]interface{} `json:"scene_str"`  // 场景值ID（字符串形式的ID），字符串类型，长度限制为1到64
}

type QrcodeParamsResp struct {
	Ticket string `json:"ticket"`
	ExpireSeconds int `json:"expire_seconds"`
	Url  string `json:"url"`
}

func (s *Server) CreateQrcodeParams(v QrcodeParamsReq) (b *QrcodeParamsResp, err error) {
	url := fmt.Sprintf(s.RootUrl+ParamsQrcode, s.GetAccessToken())
	m := new(QrcodeParamsResp)
	err = util.PostJsonPtr(url, v,m)
	if err != nil {
		return m,err
	}
	return m,err
}