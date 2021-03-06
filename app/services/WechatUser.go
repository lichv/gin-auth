package services

import (
	"fmt"
	"gin-auth/app/models"
	"gin-auth/utils"
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
	b,err = models.ExistWechatUserByCode(code)
	return b, err
}
func GetWechatUserTotal(maps interface{}) (count int,err error) {
	count,err = models.GetWechatUserTotal(map[string]interface{}{})
	return count, err
}
func GetWechatUserOne( query map[string]interface{},orderBy interface{}) (wechatUser *WechatUser, err error) {
	var nu *models.WechatUser
	nu,err = models.GetWechatUserOne(query,orderBy)
	return TransferWechatUserModel(nu),nil
}
func GetWechatUserPages( query map[string]interface{},orderBy interface{},pageNum int,pageSize int) (wechatUser []*WechatUser, total int, errs []error) {
	count,err := models.GetWechatUserTotal(query)
	fmt.Println(count)
	if err != nil {
		return nil,0,errs
	}
	us,errs := models.GetWechatUserPages(query,orderBy,pageNum,pageSize)
	wechatUser = TransferWechatUsers(us)
	return wechatUser,total,nil
}
func AddWechatUser( data map[string]interface{}) (wu *WechatUser,err error ){
	one, err := models.GetWechatUserOne(data, "code desc")
	fmt.Println("AddWechatUser")
	fmt.Println(one)
	fmt.Println(err)
	if err != nil || one.Code == ""{
		code :="wx_"+utils.GeneTimeUUID()
		data["code"] = code

		err := models.AddWechatUser(data)
		if err != nil {
			return &WechatUser{},nil
		}
		fmt.Println(code)
		user, err := models.FindWechatUserByCode(code)
		if err != nil{
			return &WechatUser{},nil
		}
		return TransferWechatUserModel(user),nil
	}

	return TransferWechatUserModel(one),nil
}
func EditWechatUser( code string,data map[string]interface{}) (err error) {
	err = models.EditWechatUser(code,data)
	return err
}
func DeleteWechatUser(maps map[string]interface{}) (err error) {
	err = models.DeleteWechatUsers(maps)
	return nil
}
func ClearAllWechatUser() (err error) {
	err = models.ClearAllWechatUser()
	return err
}

func TransferWechatUserModel(u *models.WechatUser)(wechatUser *WechatUser){
	wechatUser =  &WechatUser{
		Id:u.Id,
		Code:u.Code,
		ConfigCode:u.ConfigCode,
		Openid:u.Openid,
		Unionid:u.Unionid,
		Nickname:u.Nickname,
		Sex:u.Sex,
		Country:u.Country,
		Province:u.Province,
		City:u.City,
		Headimage:u.Headimage,
		Phone:u.Phone,
		Privilege: u.Privilege,
		CreatedOn: u.CreatedOn,
		ModifiedOn: u.ModifiedOn,
		DeletedOn: u.DeletedOn,
		Flag:u.Flag,
		State:u.State,
	}
	return
}
func TransferWechatUsers(us []*models.WechatUser) (wechatUsers []*WechatUser) {
	for _,value := range us {
		WechatUser := TransferWechatUserModel(value)
		wechatUsers = append(wechatUsers, WechatUser)
	}
	return wechatUsers
}
