package services

import (
	"fmt"
	"gin-auth/app/models"
	"gin-auth/utils/jwt"
	lichv "github.com/lichv/go"
)

type User struct {
	Id int `json:"id" form:"id" gorm:"id"`
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
	CreatedOn int64 `json:"created_on" form:"created_on" gorm:"created_on"`
	ModifiedOn int64 `json:"modified_on" form:"modified_on" gorm:"modified_on"`
	DeletedOn int64 `json:"deleted_on" form:"deleted_on" gorm:"deleted_on"`
	FLag bool `json:"flag" form:"flag" gorm:"flag"`
	State bool `json:"state" form:"state" gorm:"state"`
}

func ExistUserByCode(code string) (b bool,err error) {
	b,err = models.ExistUserByCode(code)
	return b, err
}

func GetUserTotal(maps interface{}) (count int,err error) {
	count,err = models.GetUserTotal(map[string]interface{}{})
	return count, err
}
func GetUserOne( query map[string]interface{},orderBy interface{}) (user *User, err error) {
	var nu *models.User
	nu,err = models.GetUserOne(query,orderBy)
	return TransferUserModel(nu),nil
}

func GetUserPages( query map[string]interface{},orderBy interface{},pageNum int,pageSize int) (users []*User, total int, errs []error) {
	count,err := models.GetUserTotal(query)
	fmt.Println(count)
	if err != nil {
		return nil,0,errs
	}
	us,errs := models.GetUserPages(query,orderBy,pageNum,pageSize)
	users = TransferUsers(us)
	return users,total,nil
}

func AddUser( data map[string]interface{}) (err error ){
	err = models.AddUser(data)
	return err
}

func EditUser( code string,data map[string]interface{}) (err error) {
	err = models.EditUser(code,data)
	return err
}

func DeleteUser(maps map[string]interface{}) (err error) {
	err = models.DeleteUsers(maps)
	return nil
}

func ClearAllUser() (err error) {
	err = models.ClearAllUser()
	return err
}

func TransferUserModel(u *models.User)(user *User){
	user =  &User{
		Id:u.Id,
		Code:u.Code,
		Username:u.Username,
		Name:u.Name,
		Sex:u.Sex,
		Birthday:u.Birthday,
		Phone:u.Phone,
		Email:u.Email,
		Province:u.Province,
		City:u.City,
		County:u.County,
		Address:u.Address,
		Reference:u.Reference,
		Regtime:u.Regtime,
		Remark:u.Remark,
		IsActive:u.IsActive,
		IsSuperUser:u.IsSuperUser,
		CreatedOn: u.CreatedOn,
		ModifiedOn: u.ModifiedOn,
		DeletedOn: u.DeletedOn,
		FLag:u.FLag,
		State:u.State,
	}
	return
}
func TransferUsers(us []*models.User) (users []*User) {
	for _,value := range us {
		user := TransferUserModel(value)
		users = append(users, user)
	}
	return users
}

func GenerateToken(code, username string) (string, error) {
	token_validity_time, _ := FindSysparamValueByCode("token_validity_time")
	return jwt.GenerateToken(code,username,lichv.IntVal(token_validity_time))
}

func Auth(username, password string) (string,error){
	user,err := models.Check(username,password)
	if err != nil {
		return "",err
	}
	token_validity_time, _ := FindSysparamValueByCode("token_validity_time")
	token,err := jwt.GenerateToken(user.Code,user.Username,lichv.IntVal(token_validity_time))
	if err != nil {
		return "",err
	}
	return token,nil
}