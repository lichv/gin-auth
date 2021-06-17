package models

import (
	"github.com/jinzhu/gorm"
	lichv "github.com/lichv/go"
)

type Sysparam struct {
	Id int `json:"id" form:"id" gorm:"id"`
	Code string `json:"code" form:"code" gorm:"code"`
	Name string `json:"name" form:"name" gorm:"name"`
	Label string `json:"label" form:"label" gorm:"label"`
	Value string `json:"value" form:"value" gorm:"value"`
	Datatype string `json:"datatype" form:"datatype" gorm:"datatype"`
	Options string `json:"options" form:"options" gorm:"options"`
	CreatedOn int64 `json:"created_on" form:"created_on" gorm:"created_on"`
	ModifiedOn int64 `json:"modified_on" form:"modified_on" gorm:"modified_on"`
	DeletedOn int64 `json:"deleted_on" form:"deleted_on" gorm:"deleted_on"`
	Flag bool `json:"flag" form:"flag" gorm:"flag"`
	State bool `json:"state" form:"state" gorm:"state"`
}

func ExistSysparamByCode(code string) (b bool,err error) {
	var sysparam Sysparam
	err = db.Model(&Sysparam{}).Select("code").Where("code = ? ",code).First(&sysparam).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false,err
	}
	return true, err
}

func GetSysparamTotal(maps interface{}) (count int,err error) {
	err = db.Model(&Sysparam{}).Where("state = ?",1).Count(&count).Error
	if err != nil {
		return 0,err
	}
	return count, err
}

func FindSysparamByCode( code string) ( sysparam *Sysparam, err error) {
	err = db.Model(&Sysparam{}).Where("code = ? ",code).First(&sysparam).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return &Sysparam{},err
	}
	return
}

func GetSysparamOne( query map[string]interface{},orderBy interface{}) ( *Sysparam,error) {
	var sysparam Sysparam
	model := db.Model(&Sysparam{})
	for key, value := range query {
		b,err := lichv.In ([]string{"code", "name", "label", "value", "datatype", "options", "flag", "state"},key)
		if  err == nil && b{
			model = model.Where(key + "= ?", value)
		}
	}
	err := model.First(&sysparam).Error
	if err != nil && err != gorm.ErrRecordNotFound{
		return &Sysparam{},nil
	}
	return &sysparam, nil
}

func GetSysparamPages( query map[string]interface{},orderBy interface{},pageNum int,pageSize int) ( []*Sysparam, []error) {
	var sysparams []*Sysparam
	var errs []error
	model := db.Model(&Sysparam{})
	for key, value := range query {
		b,err := lichv.In ([]string{"code", "name", "label", "value", "datatype", "options", "flag", "state"},key)
		if  err == nil && b{
			model = model.Where(key + "= ?", value)
		}
	}
	model =model.Offset(pageNum).Limit(pageSize).Order(orderBy).Find(&sysparams)
	errs = model.GetErrors()

	return sysparams, errs
}

func GetAllSysparamCode( query map[string]interface{},orderBy interface{},limit int) ( []string, []error) {
	var sysparams []Sysparam
	var errs []error
	var result []string

	model := db.Table("demo_whitelist_user").Select("code")
	for key, value := range query {
		b,err := lichv.In ([]string{"code", "name", "label", "value", "datatype", "options", "flag", "state"},key)
		if  err == nil && b{
			model = model.Where(key + "= ?", value)
		}
	}
	if limit > 0 {
		model =model.Limit(limit)
	}
	model =model.Order(orderBy).Find(&sysparams)
	errs = model.GetErrors()

	for _, v := range sysparams {
		result = append(result, v.Code)
	}

	return result, errs
}

func GetSysparams( query map[string]interface{},orderBy interface{},limit int) ( []*Sysparam, []error) {
	var Sysparams []*Sysparam
	var errs []error
	model := db.Model(&Sysparam{})
	for key, value := range query {
		b,err := lichv.In ([]string{"code", "name", "label", "value", "datatype", "options", "flag", "state"},key)
		if  err == nil && b{
			model = model.Where(key + "= ?", value)
		}
	}
	if limit > 0 {
		model =model.Limit(limit)
	}
	errs = model.Order(orderBy).Find(&Sysparams).GetErrors()

	return Sysparams, errs
}

func AddSysparam( data map[string]interface{}) error {
	sysparam := Sysparam{
		Code:data["code"].(string),
		Name:data["name"].(string),
		Label:data["label"].(string),
		Value:data["value"].(string),
		Datatype:data["datatype"].(string),
		Options:data["options"].(string),
		Flag:data["flag"].(bool),
		State:data["state"].(bool),

	}
	if err:= db.Create(&sysparam).Error;err != nil{
		return err
	}
	return nil
}

func EditSysparam( code string,data map[string]interface{}) error {
	if err:= db.Model(&Sysparam{}).Where("code=?",code).Updates(data).Error;err != nil{
		return err
	}
	return nil
}

func DeleteSysparamByCode(code string) error {
	if err := db.Where("code=?",code).Delete(Sysparam{}).Error;err != nil{
		return err
	}
	return nil
}

func DeleteSysparams(maps map[string]interface{}) error{
	model := db
	for key, value := range maps {
		b,err := lichv.In ([]string{"code", "name", "label", "value", "datatype", "options", "flag", "state"},key)
		if  err == nil && b{
			model = model.Where(key + "= ?", value)
		}
	}
	if err :=model.Delete(&Sysparam{}).Error;err != nil{
		return err
	}
	return nil
}

func ClearAllSysparam() error {
	if err := db.Unscoped().Delete(&Sysparam{}).Error; err != nil {
		return err
	}
	return nil
}
