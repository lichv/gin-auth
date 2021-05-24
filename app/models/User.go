package models

import (
	"errors"
	"gin-auth/utils"
	"gin-auth/utils/setting"
	"github.com/jinzhu/gorm"
)

type User struct {
	Code string `json:"code" form:"code" gorm:"code"`
	Username string `json:"username" form:"username" gorm:"username"`
	Password string `json:"password" form:"password" gorm:"password"`
	Name string `json:"name" form:"name" gorm:"name"`
	Sex string `json:"sex" form:"sex" gorm:"sex"`
	Birthday string `json:"birthday" form:"birthday" gorm:"birthday"`
	Phone string `json:"phone" form:"phone" gorm:"phone"`
	Email string `json:"email" form:"email" gorm:"email"`
	Province string `json:"province" form:"province" gorm:"province"`
	City string `json:"city" form:"city" gorm:"city"`
	County string `json:"county" form:"county" gorm:"county"`
	Address string `json:"address" form:"address" gorm:"address"`
	Reference string `json:"reference" form:"reference" gorm:"reference"`
	Regtime int64 `json:"regtime" form:"regtime" gorm:"regtime"`
	Remark string `json:"remark" form:"remark" gorm:"remark"`
	IsActive bool `json:"is_active" form:"is_active" gorm:"is_active"`
	IsSuperUser bool `json:"is_super_user" form:"is_super_user" gorm:"is_super_user"`
	FLag bool `json:"flag" form:"flag" gorm:"flag"`
	State bool `json:"state" form:"state" gorm:"state"`
}

func ExistUserByCode(code string) (b bool,err error) {
	var user User
	err = db.Model(&User{}).Select("code").Where("code = ? ",code).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false,err
	}
	return true, err
}

func GetUserTotal(maps interface{}) (count int,err error) {
	err = db.Model(&User{}).Where("state = ?",true).Count(&count).Error
	if err != nil {
		return 0,err
	}
	return count, err
}

func FindUserByCode( code string) ( user *User, err error) {
	err = db.Model(&User{}).Where("code = ? ",code).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return &User{},err
	}
	return user, err
}

func GetUserOne( query map[string]interface{},orderBy interface{}) ( *User,error) {
	var user User
	model := db.Model(&User{})
	for key, value := range query {
		b,err := utils.In ([]string{"code", "username", "password", "name", "sex", "birthday", "phone", "email", "province", "city", "county", "address", "reference", "regtime", "remark", "is_active", "is_superuser", "flag", "state"},key)
		if  err != nil && b{
			model = model.Where(key + "= ?", value)
		}
	}
	err := model.First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound{
		return &User{},nil
	}
	return &user, nil
}

func GetUserPages( query map[string]interface{},orderBy interface{},pageNum int,pageSize int) ( []*User, []error) {
	var users []*User
	var errs []error
	model := db.Where("state=?",true)
	for key, value := range query {
		b,err := utils.In ([]string{"code", "username", "password", "name", "sex", "birthday", "phone", "email", "province", "city", "county", "address", "reference", "regtime", "remark", "is_active", "is_superuser", "flag", "state"},key)
		if  err != nil && b{
			model = model.Where(key + "= ?", value)
		}
	}
	model =model.Offset(pageNum).Limit(pageSize).Order(orderBy)
	model = model.Find(&users)
	errs = model.GetErrors()
	//err = model.Offset(pageNum).Limit(pageSize).Order(orderBy).Find(&users).Error

	return users, errs
}

func AddUser( data map[string]interface{}) error {
	user := User{
		Code:data["code"].(string),
		Username:data["username"].(string),
		Password:data["password"].(string),
		Name:data["name"].(string),
		Sex:data["sex"].(string),
		Birthday:data["birthday"].(string),
		Phone:data["phone"].(string),
		Email:data["email"].(string),
		Province:data["province"].(string),
		City:data["city"].(string),
		County:data["country"].(string),
		Address:data["address"].(string),
		Reference:data["reference"].(string),
		Regtime:data["regtime"].(int64),
		Remark:data["remark"].(string),
		IsActive:data["is_active"].(bool),
		IsSuperUser:data["is_super_user"].(bool),
		FLag:data["flag"].(bool),
		State:data["state"].(bool),
	}
	if err:= db.Create(&user).Error;err != nil{
		return err
	}
	return nil
}

func EditUser( code string,data map[string]interface{}) error {
	if err:= db.Model(&User{}).Where("code=?",code).Updates(data).Error;err != nil{
		return err
	}
	return nil
}

func DeleteUser(code string) error {
	if err := db.Where("code=?",code).Delete(User{}).Error;err != nil{
		return err
	}
	return nil
}

func DeleteUsers(maps map[string]interface{}) error{
	model := db
	for key, value := range maps {
		b,err := utils.In ([]string{"code", "username", "password", "name", "sex", "birthday", "phone", "email", "province", "city", "county", "address", "reference", "regtime", "remark", "is_active", "is_superuser", "flag", "state"},key)
		if  err != nil && b{
			model = model.Where(key + "= ?", value)
		}
	}
	if err :=model.Delete(&User{}).Error;err != nil{
		return err
	}
	return nil
}

func ClearAllUser() error {
	if err := db.Unscoped().Delete(&User{}).Error; err != nil {
		return err
	}
	return nil
}

func Check(username, password string) (*User, error) {
	newpassword := utils.EncodeMD5(setting.AppSetting.SecretSalt+password)
	user,_:=GetUserOne(map[string]interface{}{"username":username},"code asc")

	if user.Code == "" {
		return &User{},errors.New("用户不存在或密码错误")
	}
	if user.Password != newpassword {
		return &User{},errors.New("用户密码错误或账户不存在")
	}
	return user,nil
}