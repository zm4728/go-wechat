package wechat_test

import (
	"fmt"
	"github.com/zm4728/go-wechat"
	"testing"
)

func TestWX(t *testing.T) {
	cfg := &wechat.WxConfig{
		Token:          "yourToken",
		AppId:          "yourAppID",
		Secret:         "yourSecret",
		EncodingAESKey: "yourEncodingAesKey",
	}
app := wechat.New(cfg)
	//u:=app.GetJsConfig("http://ks.mofahe.net/mp/")
 ss,err:=app.GetExternalUserInfo("wmazTqCgAAftuC2x3srH2Tjrr55vHS-Q")
	fmt.Println(err)
 fmt.Println(ss.ExternalContact)
}

	//return
	//wechat.Debug = true
	//
	//cfg := &wechat.WxConfig{
	//	Token:          "yourToken",
	//	AppId:          "yourAppID",
	//	Secret:         "yourSecret",
	//	EncodingAESKey: "yourEncodingAesKey",
	//}

	//app := wechat.New(cfg)
	//
	//	app.CreateQrcodeParams(wechat.QrcodeParamsReq{
	//	ExpireSeconds: 60,
	//	ActionName:    wechat.QR_STR_SCENE,
	//	ActionInfo:  wechat.SenceA{Scene:wechat.SenceB{SceneStr: map[string]interface{}{
	//		"bb":"11",
	//	},
	//	}},},
	//	)
	//}
	//http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//	ctx := app.VerifyURL(w, r)
	//
	//	// 根据消息类型主动回复
	//	switch ctx.Msg.MsgType {
	//	case wechat.TypeText:
	//		ctx.NewText(ctx.Msg.Content).Reply() // 回复文字
	//	case wechat.TypeImage:
	//		ctx.NewImage(ctx.Msg.MediaId).Reply() // 回复图片
	//	case wechat.TypeVoice:
	//		ctx.NewVoice(ctx.Msg.MediaId).Reply() // 回复语音
	//	case wechat.TypeVideo:
	//		ctx.NewVideo(ctx.Msg.MediaId, "video title", "video description").Reply() //回复视频
	//	case wechat.TypeFile:
	//		ctx.NewFile(ctx.Msg.MediaId).Reply() // 回复文件，仅企业微信可用
	//	default:
	//		ctx.NewText("其他消息类型" + ctx.Msg.MsgType).Reply() // 回复模板消息
	//	}
	//})

	// http.ListenAndServe(":9090", nil)


// func Example_echo() {
// 	app := wechat.New("yourToken", "yourAppID", "yourSecret", "yourEncodingAesKey")
// 	e := echo.New()
// 	e.Any("/", func(c echo.Context) error {
// 		ctx := app.VerifyURL(c.Response().Writer, c.Request())

// 		// 根据消息类型主动回复
// 		switch ctx.Msg.MsgType {
// 		case wechat.TypeText:
// 			ctx.NewText(ctx.Msg.Content).Reply() // 回复文字
// 		case wechat.TypeImage:
// 			ctx.NewImage(ctx.Msg.MediaId).Reply() // 回复图片
// 		case wechat.TypeVoice:
// 			ctx.NewVoice(ctx.Msg.MediaId).Reply() // 回复语音
// 		case wechat.TypeVideo:
// 			ctx.NewVideo(ctx.Msg.MediaId, "video title", "video description").Reply() //回复视频
// 		case wechat.TypeFile:
// 			ctx.NewFile(ctx.Msg.MediaId).Reply() // 回复文件，仅企业微信可用
// 		default:
// 			ctx.NewText("其他消息类型" + ctx.Msg.MsgType).Reply() // 回复模板消息
// 		}
// 		return nil
// 	})
// 	e.Start(":9090")
// }
