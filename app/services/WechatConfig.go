package services

import (
	"fmt"
	"gin-auth/app/models"
)

type WechatConfig struct {
	Id int `json:"id" form:"id" gorm:"id"`
	Code string `json:"code" form:"code"`
	Type string `json:"type" form:"type"`
	Appid string `json:"appid" form:"appid"`
	Appsecret string `json:"appsecret" form:"appsecret"`
	Scope string `json:"scope" form:"scope"`
	AuthRedirectUrl string `json:"auth_redirect_url" form:"auth_redirect_url"`
	NoticeUrl string `json:"notice_url" form:"notice_url"`
	Group string `json:"group" form:"group"`
	Company string `json:"company" form:"company"`
	CreatedOn int64 `json:"created_on" form:"created_on" gorm:"created_on"`
	ModifiedOn int64 `json:"modified_on" form:"modified_on" gorm:"modified_on"`
	DeletedOn int64 `json:"deleted_on" form:"deleted_on" gorm:"deleted_on"`
	Flag bool `json:"flag" form:"flag"`
	State bool `json:"state" form:"state"`
}

func ExistWechatConfigByCode(code string) (b bool,err error) {
	b,err = models.ExistWechatConfigByCode(code)
	return b, err
}

func GetWechatConfigTotal(maps interface{}) (count int,err error) {
	count,err = models.GetWechatConfigTotal(map[string]interface{}{})
	return count, err
}
func GetWechatConfigOne( query map[string]interface{},orderBy interface{}) (wechatConfig *WechatConfig, err error) {
	wc, err := models.GetWechatConfigOne(query, orderBy)
	return TransferWechatConfigModel(wc),nil
}

func GetWechatConfigPages( query map[string]interface{},orderBy interface{},pageNum int,pageSize int) (wechatConfigs []*WechatConfig, total int, errs []error) {
	count,err := models.GetWechatConfigTotal(query)
	fmt.Println(count)
	if err != nil {
		return nil,0,errs
	}
	us,errs := models.GetWechatConfigPages(query,orderBy,pageNum,pageSize)
	wechatConfigs = TransferWechatConfigs(us)
	return wechatConfigs,total,nil
}

func AddWechatConfig( data map[string]interface{}) (err error ){
	err = models.AddWechatConfig(data)
	return err
}

func EditWechatConfig( code string,data map[string]interface{}) (err error) {
	err = models.EditWechatConfig(code,data)
	return err
}

func DeleteWechatConfig(maps map[string]interface{}) (err error) {
	err = models.DeleteWechatConfigs(maps)
	return nil
}

func ClearAllWechatConfig() (err error) {
	err = models.ClearAllWechatConfig()
	return err
}

func TransferWechatConfigModel(u *models.WechatConfig)(wechatConfigs *WechatConfig){
	wechatConfigs =  &WechatConfig{
		Id:u.Id,
		Code:u.Code,
		Type:u.Type,
		Appid:u.Appid,
		Appsecret:u.Appsecret,
		Scope:u.Scope,
		AuthRedirectUrl:u.AuthRedirectUrl,
		NoticeUrl:u.NoticeUrl,
		Group:u.Group,
		Company:u.Company,
		CreatedOn: u.CreatedOn,
		ModifiedOn: u.ModifiedOn,
		DeletedOn: u.DeletedOn,
		Flag:u.Flag,
		State:u.State,
	}
	return
}
func TransferWechatConfigs(us []*models.WechatConfig) (wechatConfigs []*WechatConfig) {
	for _,value := range us {
		WechatConfig := TransferWechatConfigModel(value)
		wechatConfigs = append(wechatConfigs, WechatConfig)
	}
	return wechatConfigs
}