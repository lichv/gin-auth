package models

import (
	"fmt"
	"gin-auth/utils"
	"github.com/jinzhu/gorm"
)

type WechatUser struct {
	Id int `json:"id" form:"id" gorm:"id"`
	Code string `json:"code" form:"code"`
	ConfigCode string `json:"config_code" form:"config_code"`
	Openid string `json:"openid" form:"openid"`
	Unionid string `json:"unionid" form:"unionid"`
	Nickname string `json:"nickname" form:"nickname"`
	Sex string `json:"sex" form:"sex"`
	Headimage string `json:"headimage" form:"headimage"`
	Country string `json:"country" form:"country"`
	Province string `json:"province" form:"province"`
	City string `json:"city" form:"city"`
	Phone string `json:"phone" form:"phone"`
	Privilege string `json:"privilege" form:"privilege"`
	CreatedOn int64 `json:"created_on" form:"created_on" gorm:"created_on"`
	ModifiedOn int64 `json:"modified_on" form:"modified_on" gorm:"modified_on"`
	DeletedOn int64 `json:"deleted_on" form:"deleted_on" gorm:"deleted_on"`
	Flag bool `json:"flag" form:"flag"`
	State bool `json:"state" form:"state"`
}

func ExistWechatUserByCode(code string) (b bool,err error) {
	var wechatUser WechatUser
	err = db.Model(&WechatUser{}).Select("code").Where("code = ? ",code).First(&wechatUser).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false,err
	}
	return true, err
}

func GetWechatUserTotal(maps interface{}) (count int,err error) {
	err = db.Model(&WechatUser{}).Where("state = ?",true).Count(&count).Error
	if err != nil {
		return 0,err
	}
	return count, err
}

func FindWechatUserByCode( code string) ( *WechatUser, error) {
	var wechatUser WechatUser
	var err error
	err = db.Model(&WechatUser{}).Where("code = ? ",code).First(&wechatUser).Error
	fmt.Println(wechatUser)
	if err != nil && err != gorm.ErrRecordNotFound {
		return &WechatUser{},err
	}
	return &wechatUser, err
}

func GetWechatUserOne( query map[string]interface{},orderBy interface{}) ( *WechatUser,error) {
	var wechatUser WechatUser
	model := db.Model(&WechatUser{})
	for key, value := range query {
		b,err := utils.In ([]string{"code", "config_code", "openid", "unionid", "nickname", "sex", "country", "province", "city", "headimage", "privilege", "flag", "state"},key)
		if  err == nil && b{
			model = model.Where(key + "= ?", value)
		}
	}
	err := model.First(&wechatUser).Error
	if err != nil && err != gorm.ErrRecordNotFound{
		return &WechatUser{},nil
	}
	return &wechatUser, nil
}

func GetWechatUserPages( query map[string]interface{},orderBy interface{},pageNum int,pageSize int) ( []*WechatUser, []error) {
	var wechatUsers []*WechatUser
	var errs []error
	model := db.Where("state=?",true)
	for key, value := range query {
		b,err := utils.In ([]string{"code", "config_code", "openid", "unionid", "nickname", "sex", "country", "province", "city", "headimage", "privilege", "flag", "state"},key)
		if  err == nil && b{
			model = model.Where(key + "= ?", value)
		}
	}
	model =model.Offset(pageNum).Limit(pageSize).Order(orderBy)
	model = model.Find(&wechatUsers)
	errs = model.GetErrors()

	return wechatUsers, errs
}

func AddWechatUser( data map[string]interface{}) error {
	phone,ok := data["phone"]
	if !ok {
		phone = ""
	}
	dataFlag,_ := data["flag"]
	flag := utils.BoolVal(dataFlag)

	dataState,_ := data["state"]
	state := utils.BoolVal(dataState)

	dataPrivilege,_ := data["privilege"]
	privilege := utils.Strval(dataPrivilege)
	if privilege == ""{
		privilege = "{}"
	}
	wechatUser := WechatUser{
		Code:data["code"].(string),
		ConfigCode:data["config_code"].(string),
		Openid:data["openid"].(string),
		Unionid:data["unionid"].(string),
		Nickname:data["nickname"].(string),
		Sex:data["sex"].(string),
		Headimage:data["headimage"].(string),
		Country:data["country"].(string),
		Province:data["province"].(string),
		City:data["city"].(string),
		Phone:utils.Strval(phone),
		Privilege:privilege,
		Flag:flag,
		State:state,
	}
	if err:= db.Create(&wechatUser).Error;err != nil{
		return err
	}
	return nil
}

func EditWechatUser( code string,data map[string]interface{}) error {
	if err:= db.Model(&WechatUser{}).Where("code=?",code).Updates(data).Error;err != nil{
		return err
	}
	return nil
}

func DeleteWechatUser(code string) error {
	if err := db.Where("code=?",code).Delete(WechatUser{}).Error;err != nil{
		return err
	}
	return nil
}

func DeleteWechatUsers(maps map[string]interface{}) error{
	model := db
	for key, value := range maps {
		b,err := utils.In ([]string{"code", "username", "password", "name", "sex", "birthday", "phone", "email", "province", "city", "county", "address", "reference", "regtime", "remark", "is_active", "is_superWechatUser", "flag", "state"},key)
		if  err == nil && b{
			model = model.Where(key + "= ?", value)
		}
	}
	if err :=model.Delete(&WechatUser{}).Error;err != nil{
		return err
	}
	return nil
}

func ClearAllWechatUser() error {
	if err := db.Unscoped().Delete(&WechatUser{}).Error; err != nil {
		return err
	}
	return nil
}
