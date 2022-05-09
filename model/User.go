package model

import (
	"encoding/base64"
	"ginblog/utils/errmsg"
	"golang.org/x/crypto/scrypt"
	"gorm.io/gorm"
	"log"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(32);not null" json:"username"`   //用户名
	Password string `gorm:"type:varchar(32);not null " json:"password"`   //密码
	Role int `gorm:"type:int " json:"role"`  //身份码，0：管理员，1：阅读者
}


//CheckUser   查询用户是否存在
func CheckUser(userName string)int {
	var users User
	//db.Where("username=?",userName).First(&users)
	//select id from users where username = userName ;
	db.Select("id").Where("username=?",userName).First(&users)
	if users.ID>0{
		return errmsg.ERROR_USERNAME_USED
	}
	return errmsg.SUCCESS
}

//CreateUser 新增用户，写入数据库
func CreateUser(data *User) int {
	//加密密码使用钩子函数BeforeSave执行
	//Create insert the value into database
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// GetUsers 查询用户列表,返回user切片，需要分页显示，不然如果用户列表庞大，一次性查询性能低下
func GetUsers(pageSize,pageNo int) []User {
	var users []User
	//sql limit分页查询
	offset := (pageNo -1)*pageSize
	if pageNo == -1 && pageSize == -1 {
		offset = -1
	}
	err = db.Limit(pageSize).Offset(offset).Find(&users).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return users
}

//BeforeSave   钩子函数
func (u *User)BeforeSave(_ *gorm.DB)(err error){
	u.Password=ScryptPw(u.Password)
	return nil
}

//ScryptPw  密码加密
func ScryptPw(password string)string{
	const KeyLen = 10
	salt := make([]byte,8)
	salt = []byte{12,32,4,6,66,22,222,11}
	HashPw,err := scrypt.Key([]byte(password),salt,16384,8,1,KeyLen)
	if err != nil {
		log.Fatal(err)
	}

	retPw := base64.StdEncoding.EncodeToString(HashPw)
	return retPw
}

//EditUser  编辑用户信息（除密码，密码单独做）
func EditUser (id int,data *User)int{
	maps := make(map[string]interface{})
	maps["username"]=data.Username
	maps["role"]=data.Role
	err = db.Model(&User{}).Where("id=?",id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

//DeleteUser  删除用户
func DeleteUser(id int) int {
	var user User
	err = db.Where("id=?",id).Delete(&user).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

//FindByID 查询单个用户
func FindByID(id int) *User{
	var user User
	err = db.Last(&user,id).Error
	if err != nil {
		return nil
	}
	return &user
}