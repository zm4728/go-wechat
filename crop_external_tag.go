package wechat

import (
	"encoding/json"
	"github.com/zm4728/go-wechat/util"
)

//企业微信外部联系人标签
const (
	//企业微信外部联系人标签
	CropApiExternalTagList = CorpAPI + `externalcontact/get_corp_tag_list?access_token=` //获取企业标签库
	CropApiExternalTagAdd = CorpAPI + `externalcontact/add_corp_tag?access_token=` //添加企业标签库
	CropApiExternalTagEdit = CorpAPI + `externalcontact/edit_corp_tag?access_token=` //编辑企业标签库
	CropApiExternalTagDel = CorpAPI + `externalcontact/del_corp_tag?access_token=` //删除企业标签库

	CropApiExternalUpdateTag = CorpAPI + `externalcontact/mark_tag?access_token=` //编辑客户标签
)

type (

	// 企业标签列表请求参数
	ExTagGroupReq struct {
		TagId     []string      `json:"tag_id"`
		GroupId   []string 		`json:"group_id"`
	}
	// TagList 企业标签列表
	ExTagList struct {
		WxErr
		TagGroup []ExTagGroup `json:"tag_group"`
	}
	// 企业标签
	ExTagGroup struct {
		GroupId   string    `json:"group_id"`
		GroupName 	string `json:"group_name"`
		CreateTime  int `json:"create_time"`
		Order      int `json:"order"`
		Deleted   bool `json:"deleted"`
		Tag   []ExTagResp `json:"tag"`
	}

	// Tag 标签
	ExTagResp struct {
		Id   string `json:"id"`
		Name string `json:"name"`
		CreateTime  int `json:"create_time"`
		Order      int `json:"order"`
		Deleted   bool `json:"deleted"`
	}


	ExTagDetail struct {
		WxErr
		TagGroup ExTagGroup `json:"tag_group"`
	}

	// 企业标签新增 (请求body)
	ExAddTagReq struct {
		GroupId     string 		`json:"group_id"`
		GroupName 	string `json:"group_name"`
		Order      int `json:"order"`
		Tag   []ExTagReq`json:"tag"`
	}
	ExTagReq struct {
		Name 	string `json:"name"`
		Order   int `json:"order"`
	}

	// 编辑企业标签(请求)
	ExTagEditReq struct {
		Id      string `json:"id"` //标签或标签组的id
		ExTagReq
	}

	// 编辑客户企业标签（给外部联系人打标签）
	ExtUserTagReq struct {
		Userid  string `json:"userid"` // 添加外部联系人userid  //企业员工userId
		ExternalUserid string `json:"external_userid"`  //外部联系人userid
		AddTag   []string `json:"add_tag"` //要标记的标签列表
		RemoveTag []string `json:"remove_tag"` //要移除的标签列表
	}

)

// GetTagList 获取标签列表
func (s *Server) GetExtTagList(v ExTagGroupReq) (l ExTagList, err  WxErr) {
	url := CropApiExternalTagList + s.GetUserAccessToken()
	body, errs := util.PostJson(url, v)
	if errs != nil {
		return l,WxErr{-1, errs.Error()}
	}
	errs = json.Unmarshal(body, &l)
	if errs != nil {
		return l,WxErr{-1, errs.Error()}
	}
	return
}

// 添加企业标签库
func (s *Server) AddExtTagList(v ExAddTagReq) (l ExTagDetail, err  WxErr) {
	url := CropApiExternalTagAdd + s.GetUserAccessToken()
	body, errs := util.PostJson(url, v)
	if errs != nil {
		return l,WxErr{-1, errs.Error()}
	}
	errs = json.Unmarshal(body, &l)
	if errs != nil {
		return l,WxErr{-1, errs.Error()}
	}
	return
}

// 修改企业标签库
func (s *Server) EditExtTagList(v ExTagEditReq) (err  WxErr) {
	url := CropApiExternalTagEdit + s.GetUserAccessToken()
	body, errs := util.PostJson(url, v)
	if errs != nil {
		return WxErr{-1, errs.Error()}
	}
	errs = json.Unmarshal(body, &err)
	if errs != nil {
		return WxErr{-1, errs.Error()}
	}
	return err
}

// 删除企业标签库
func (s *Server) DelExtTagList(v ExTagGroupReq) (err  WxErr) {
	// GroupId 慎用！会删除整个标签组
	url := CropApiExternalTagDel + s.GetUserAccessToken()
	body, errs := util.PostJson(url, v)
	if errs != nil {
		return WxErr{-1, errs.Error()}
	}
	errs = json.Unmarshal(body, &err)
	if errs != nil {
		return WxErr{-1, errs.Error()}
	}
	return err
}

// 给外部联系人打企业标签
func (s *Server) UpdateExtUserTag(v ExtUserTagReq) (err  WxErr) {
	url := CropApiExternalUpdateTag + s.GetUserAccessToken()
	body, errs := util.PostJson(url, v)
	if errs != nil {
		return WxErr{-1, errs.Error()}
	}
	errs = json.Unmarshal(body, &err)
	if errs != nil {
		return WxErr{-1, errs.Error()}
	}
	return err
}


