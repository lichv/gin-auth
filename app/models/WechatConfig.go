package models

import (
	"gin-auth/utils"
	"github.com/jinzhu/gorm"
)

type WechatConfig struct {
	Code string `json:"code" form:"code"`
	Type string `json:"type" form:"type"`
	Appid string `json:"appid" form:"appid"`
	Appsecret string `json:"appsecret" form:"appsecret"`
	Scope string `json:"scope" form:"scope"`
	AuthRedirectUrl string `json:"auth_redirect_url" form:"auth_redirect_url"`
	NoticeUrl string `json:"notice_url" form:"notice_url"`
	Group string `json:"group" form:"group"`
	Company string `json:"company" form:"company"`
	Flag bool `json:"flag" form:"flag"`
	State bool `json:"state" form:"state"`
}

func ExistWechatConfigByCode(code string) (b bool,err error) {
	var wechatConfig WechatConfig
	err = db.Model(&WechatConfig{}).Select("code").Where("code = ? ",code).First(&wechatConfig).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false,err
	}
	return true, err
}

func GetWechatConfigTotal(maps interface{}) (count int,err error) {
	err = db.Model(&WechatConfig{}).Where("state = ?",true).Count(&count).Error
	if err != nil {
		return 0,err
	}
	return count, err
}

func FindWechatConfigByCode( code string) ( wechatConfig *WechatConfig, err error) {
	err = db.Model(&WechatConfig{}).Where("code = ? ",code).First(&wechatConfig).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return &WechatConfig{},err
	}
	return wechatConfig, err
}

func GetWechatConfigOne( query map[string]interface{},orderBy interface{}) ( *WechatConfig,error) {
	var wechatConfig WechatConfig
	model := db.Model(&WechatConfig{})
	for key, value := range query {
		b,err := utils.In ([]string{"code", "type", "appid", "appsecret", "scope", "auth_redirect_url", "notice_url", "group", "company", "flag", "state"},key)
		if  err == nil && b{
			model = model.Where(key + "= ?", value)
		}
	}
	err := model.First(&wechatConfig).Error
	if err != nil && err != gorm.ErrRecordNotFound{
		return &WechatConfig{},nil
	}
	return &wechatConfig, nil
}

func GetWechatConfigPages( query map[string]interface{},orderBy interface{},pageNum int,pageSize int) ( []*WechatConfig, []error) {
	var WechatConfigs []*WechatConfig
	var errs []error
	model := db.Where("state=?",true)
	for key, value := range query {
		b,err := utils.In ([]string{"code", "WechatConfigname", "password", "name", "sex", "birthday", "phone", "email", "province", "city", "county", "address", "reference", "regtime", "remark", "is_active", "is_superWechatConfig", "flag", "state"},key)
		if  err == nil && b{
			model = model.Where(key + "= ?", value)
		}
	}
	model =model.Offset(pageNum).Limit(pageSize).Order(orderBy)
	model = model.Find(&WechatConfigs)
	errs = model.GetErrors()
	//err = model.Offset(pageNum).Limit(pageSize).Order(orderBy).Find(&WechatConfigs).Error

	return WechatConfigs, errs
}

func AddWechatConfig( data map[string]interface{}) error {
	dataFlag,_ := data["flag"]
	flag := utils.BoolVal(dataFlag)

	dataState,_ := data["state"]
	state := utils.BoolVal(dataState)
	WechatConfig := WechatConfig{
		Code:data["code"].(string),
		Type:data["type"].(string),
		Appid:data["appid"].(string),
		Appsecret:data["appsecret"].(string),
		Scope:data["scope"].(string),
		AuthRedirectUrl:data["auth_redirect_url"].(string),
		NoticeUrl:data["notice_url"].(string),
		Group:data["group"].(string),
		Company:data["company"].(string),
		Flag:flag,
		State:state,
	}
	if err:= db.Create(&WechatConfig).Error;err != nil{
		return err
	}
	return nil
}

func EditWechatConfig( code string,data map[string]interface{}) error {
	if err:= db.Model(&WechatConfig{}).Where("code=?",code).Updates(data).Error;err != nil{
		return err
	}
	return nil
}

func DeleteWechatConfig(code string) error {
	if err := db.Where("code=?",code).Delete(WechatConfig{}).Error;err != nil{
		return err
	}
	return nil
}

func DeleteWechatConfigs(maps map[string]interface{}) error{
	model := db
	for key, value := range maps {
		b,err := utils.In ([]string{"code", "WechatConfigname", "password", "name", "sex", "birthday", "phone", "email", "province", "city", "county", "address", "reference", "regtime", "remark", "is_active", "is_superWechatConfig", "flag", "state"},key)
		if  err == nil && b{
			model = model.Where(key + "= ?", value)
		}
	}
	if err :=model.Delete(&WechatConfig{}).Error;err != nil{
		return err
	}
	return nil
}

func ClearAllWechatConfig() error {
	if err := db.Unscoped().Delete(&WechatConfig{}).Error; err != nil {
		return err
	}
	return nil
}
