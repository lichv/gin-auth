package services

import (
	"gin-auth/app/models"
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
	b,err = models.ExistSysparamByCode(code)
	return b, err
}

func GetSysparamTotal(maps interface{}) (count int,err error) {
	count,err = models.GetSysparamTotal(map[string]interface{}{})
	return count, err
}
func GetSysparamOne( query map[string]interface{},orderBy interface{}) (*Sysparam, error) {
	var nu *models.Sysparam
	nu,err := models.GetSysparamOne(query,orderBy)
	if err != nil {
		return nil,err
	}
	return TransferSysparamModel(nu),nil
}

func FindSysparamValueByCode( code string) (string, error) {
	var nu *models.Sysparam
	nu,err := models.GetSysparamOne(map[string]interface{}{"code":code},"code asc")
	if err != nil {
		return "",err
	}
	return nu.Value,nil
}

func GetSysparamPages( query map[string]interface{},orderBy interface{},pageNum int,pageSize int) (sysparams []*Sysparam, total int, errs []error) {
	total,err := models.GetSysparamTotal(query)
	if err != nil {
		return nil,0,errs
	}
	us,errs := models.GetSysparamPages(query,orderBy,pageNum,pageSize)
	sysparams = TransferSysparams(us)
	return sysparams,total,nil
}
func GetAllSysparamCode( query map[string]interface{},orderBy interface{},limit int)([]string,[]error){
	codes, errors := models.GetAllSysparamCode(query, orderBy, limit)
	return codes,errors
}
func GetSysparams( query map[string]interface{},orderBy interface{},limit int) ([]*Sysparam,[]error) {
	users, errors := models.GetSysparams(query, orderBy, limit)
	sysparams := TransferSysparams(users)
	return sysparams,errors
}

func AddSysparam( data map[string]interface{}) (err error ){
	err = models.AddSysparam(data)
	return err
}

func EditSysparam( code string,data map[string]interface{}) (err error) {
	err = models.EditSysparam(code,data)
	return err
}

func DeleteSysparam(maps map[string]interface{}) (err error) {
	err = models.DeleteSysparams(maps)
	return nil
}

func ClearAllSysparam() (err error) {
	err = models.ClearAllSysparam()
	return err
}

func TransferSysparamModel(u *models.Sysparam)(sysparam *Sysparam){
	sysparam =  &Sysparam{
		Id:u.Id,
		Code:u.Code,
		Name:u.Name,
		Label:u.Label,
		Value:u.Value,
		Datatype:u.Datatype,
		Options:u.Options,
		CreatedOn: u.CreatedOn,
		ModifiedOn: u.ModifiedOn,
		DeletedOn: u.DeletedOn,
		Flag:u.Flag,
		State:u.State,
	}
	return
}
func TransferSysparams(us []*models.Sysparam) (sysparams []*Sysparam) {
	for _,value := range us {
		sysparam := TransferSysparamModel(value)
		sysparams = append(sysparams, sysparam)
	}
	return sysparams
}
