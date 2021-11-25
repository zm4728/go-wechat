package wechat

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/zm4728/go-wechat/util"
)

const (
	addNewsURL          = "https://api.weixin.qq.com/cgi-bin/material/add_news"
	uploadimgURL        = "https://api.weixin.qq.com/cgi-bin/media/uploadimg"
	updateNewsURL       = "https://api.weixin.qq.com/cgi-bin/material/update_news"
	addMaterialURL      = "https://api.weixin.qq.com/cgi-bin/material/add_material"
	delMaterialURL      = "https://api.weixin.qq.com/cgi-bin/material/del_material"
	getMaterialURL      = "https://api.weixin.qq.com/cgi-bin/material/get_material"
	getMaterialCountURL = "https://api.weixin.qq.com/cgi-bin/material/get_materialcount"
	batchGetMaterialURL = "https://api.weixin.qq.com/cgi-bin/material/batchget_material"
)

const (
	// PermanentMaterialTypeImage 永久素材图片类型（image）
	PermanentMaterialTypeImage string = "image"
	// PermanentMaterialTypeVideo 永久素材视频类型（video）
	PermanentMaterialTypeVideo string = "video"
	// PermanentMaterialTypeVoice 永久素材语音类型 （voice）
	PermanentMaterialTypeVoice string = "voice"
	// PermanentMaterialTypeNews 永久素材图文类型（news）
	PermanentMaterialTypeNews string = "news"
)


// Article 永久图文素材
type MaterialArticle struct {
	Title            string `json:"title"`
	ThumbMediaID     string `json:"thumb_media_id"`
	ThumbURL         string `json:"thumb_url"`
	Author           string `json:"author"`
	Digest           string `json:"digest"`
	ShowCoverPic     int    `json:"show_cover_pic"`
	Content          string `json:"content"`
	ContentSourceURL string `json:"content_source_url"`
	URL              string `json:"url"`
	DownURL          string `json:"down_url"`
	NeedOpenComment  int `json:"need_open_comment"`
	OnlyFansCanComment int `json:"only_fans_can_comment"`
}

// GetNews 获取/下载永久素材
func (s *Server) GetNews(id string) ([]*MaterialArticle, error) {
	uri := fmt.Sprintf("%s?access_token=%s", getMaterialURL,  s.GetAccessToken())

	var req struct {
		MediaID string `json:"media_id"`
	}
	req.MediaID = id
	responseBytes, err := util.PostJson(uri, req)
	if err != nil {
		return nil, err
	}

	var res struct {
		NewsItem []*MaterialArticle `json:"news_item"`
	}
	err = json.Unmarshal(responseBytes, &res)
	if err != nil {
		return nil, err
	}

	return res.NewsItem, nil
}

// reqArticles 永久性图文素材请求信息
type reqArticles struct {
	Articles []*MaterialArticle `json:"articles"`
}

// resArticles 永久性图文素材返回结果
type resArticles struct {
	WxErr

	MediaID string `json:"media_id"`
}

// AddNews 新增永久图文素材
func (s *Server) AddNews(articles []*MaterialArticle) (mediaID string, err error) {
	req := &reqArticles{articles}

	uri := fmt.Sprintf("%s?access_token=%s", addNewsURL, s.GetAccessToken())
	responseBytes, err := util.PostJson(uri, req)
	if err != nil {
		return
	}
	var res resArticles
	err = json.Unmarshal(responseBytes, &res)
	if err != nil {
		return
	}
	if res.ErrCode != 0 {
		return "", fmt.Errorf("errcode=%d,errmsg=%s", res.ErrCode, res.ErrMsg)
	}
	mediaID = res.MediaID
	return
}

// reqUpdateArticle 更新永久性图文素材请求信息
type reqUpdateArticle struct {
	MediaID  string   `json:"media_id"`
	Index    int64    `json:"index"`
	Articles *MaterialArticle `json:"articles"`
}

// UpdateNews 更新永久图文素材
func (s *Server) UpdateNews(article *MaterialArticle, mediaID string, index int64) (err error) {
	req := &reqUpdateArticle{mediaID, index, article}

	uri := fmt.Sprintf("%s?access_token=%s", updateNewsURL, s.GetAccessToken())
	var response []byte
	response, err = util.PostJson(uri, req)
	if err != nil {
		return
	}
	return errors.New(string(response))
}

// resAddMaterial 永久性素材上传返回的结果
type resAddMaterial struct {
	WxErr

	MediaID string `json:"media_id"`
	URL     string `json:"url"`
}

// AddMaterial 上传永久性素材（处理视频需要单独上传）

func (s *Server) AddMaterial(mediaType string, filename string) (mediaID string, url string, err error) {
	if mediaType == PermanentMaterialTypeVideo {
		err = errors.New("永久视频素材上传使用 AddVideo 方法")
		return
	}

	if err != nil {
		return
	}

	uri := fmt.Sprintf("%s?access_token=%s&type=%s", addMaterialURL, s.GetAccessToken(), mediaType)
	var response []byte
	response, err = util.PostFile("media", filename, uri)
	if err != nil {
		return
	}
	var resMaterial resAddMaterial
	err = json.Unmarshal(response, &resMaterial)
	if err != nil {
		return
	}
	if resMaterial.ErrCode != 0 {
		err = fmt.Errorf("AddMaterial error : errcode=%v , errmsg=%v", resMaterial.ErrCode, resMaterial.ErrMsg)
		return
	}
	mediaID = resMaterial.MediaID
	url = resMaterial.URL
	return
}

/* 上传图文消息内的图片获取URL
本接口所上传的图片不占用公众号的素材库中图片数量的100000个的限制。图片仅支持jpg/png格式，大小必须在1MB以下。
**/
func (s *Server) Uploadimg(filename string) (url string, err error) {

	uri := fmt.Sprintf("%s?access_token=%s", uploadimgURL, s.GetAccessToken())
	var response []byte
	response, err = util.PostFile("media", filename, uri)
	if err != nil {
		return
	}
	var resMaterial resAddMaterial
	err = json.Unmarshal(response, &resMaterial)
	if err != nil {
		return
	}
	if resMaterial.ErrCode != 0 {
		err = fmt.Errorf("AddMaterial error : errcode=%v , errmsg=%v", resMaterial.ErrCode, resMaterial.ErrMsg)
		return
	}
	url = resMaterial.URL
	return
}

type reqVideo struct {
	Title        string `json:"title"`
	Introduction string `json:"introduction"`
}

// AddVideo 永久视频素材文件上传
func (s *Server) AddVideo(filename, title, introduction string) (mediaID string, url string, err error) {


	uri := fmt.Sprintf("%s?access_token=%s&type=video", addMaterialURL, s.GetAccessToken())

	videoDesc := &reqVideo{
		Title:        title,
		Introduction: introduction,
	}
	var fieldValue []byte
	fieldValue, err = json.Marshal(videoDesc)
	if err != nil {
		return
	}

	fields := []util.MultipartFormField{
		{
			IsFile:    true,
			Fieldname: "media",
			Filename:  filename,
		},
		{
			IsFile:    false,
			Fieldname: "description",
			Value:     fieldValue,
		},
	}

	var response []byte
	response, err = util.PostMultipartForm(fields, uri)
	if err != nil {
		return
	}

	var resMaterial resAddMaterial
	err = json.Unmarshal(response, &resMaterial)
	if err != nil {
		return
	}
	if resMaterial.ErrCode != 0 {
		err = fmt.Errorf("AddMaterial error : errcode=%v , errmsg=%v", resMaterial.ErrCode, resMaterial.ErrMsg)
		return
	}
	mediaID = resMaterial.MediaID
	url = resMaterial.URL
	return
}

type reqDeleteMaterial struct {
	MediaID string `json:"media_id"`
}

// DeleteMaterial 删除永久素材
func (s *Server) DeleteMaterial(mediaID string) error {

	uri := fmt.Sprintf("%s?access_token=%s", delMaterialURL, s.GetAccessToken())
	_, err := util.PostJson(uri, reqDeleteMaterial{mediaID})
	if err != nil {
		return err
	}
	return err
}

// ArticleList 永久素材列表
type ArticleList struct {
	WxErr
	TotalCount int64             `json:"total_count"`
	ItemCount  int64             `json:"item_count"`
	Item       []ArticleListItem `json:"item"`
}

// ArticleListItem 用于ArticleList的item节点
type ArticleListItem struct {
	MediaID    string             `json:"media_id"`
	Content    ArticleListContent `json:"content"`
	Name       string             `json:"name"`
	URL        string             `json:"url"`
	UpdateTime int64              `json:"update_time"`
}

// ArticleListContent 用于ArticleListItem的content节点
type ArticleListContent struct {
	NewsItem   []MaterialArticle `json:"news_item"`
	UpdateTime int64     `json:"update_time"`
	CreateTime int64     `json:"create_time"`
}

// reqBatchGetMaterial BatchGetMaterial请求参数
type reqBatchGetMaterial struct {
	Type   string                `json:"type"`
	Count  int64                 `json:"count"`
	Offset int64                 `json:"offset"`
}

// BatchGetMaterial 批量获取永久素材
//reference:https://developers.weixin.qq.com/doc/offiaccount/Asset_Management/Get_materials_list.html
func (s *Server) BatchGetMaterial(permanentMaterialType string, offset, count int64) (list ArticleList, err error) {

	uri := fmt.Sprintf("%s?access_token=%s", batchGetMaterialURL, s.GetAccessToken())

	req := reqBatchGetMaterial{
		Type:   permanentMaterialType,
		Offset: offset,
		Count:  count,
	}

	var response []byte
	response, err = util.PostJson(uri, req)
	if err != nil {
		return
	}
	if err = json.Unmarshal(response, &list); err != nil {
		return
	}
	return
}

// ResMaterialCount 素材总数
type ResMaterialCount struct {
	WxErr
	VoiceCount int64 `json:"voice_count"` // 语音总数量
	VideoCount int64 `json:"video_count"` // 视频总数量
	ImageCount int64 `json:"image_count"` // 图片总数量
	NewsCount  int64 `json:"news_count"`  // 图文总数量
}

// GetMaterialCount 获取素材总数.
func (s *Server) GetMaterialCount() (res ResMaterialCount, err error) {


	uri := fmt.Sprintf("%s?access_token=%s", getMaterialCountURL, s.GetAccessToken())
	var response []byte
	response, err = util.GetBody(uri)
	if err != nil {
		return
	}
	if err = json.Unmarshal(response, &res); err != nil {
		return
	}
	return
}