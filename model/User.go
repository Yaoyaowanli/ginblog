package model

import (
	"encoding/base64"
	"ginblog/utils/errmsg"
	"golang.org/x/crypto/scrypt"
	"gorm.io/gorm"
	"log"
)

//User 用户模型
type User struct {
	gorm.Model    //gorm模型
	Username string `gorm:"type:varchar(32);not null" json:"username" validate:"required,min=4,max=12" label:"用户名"`   //用户名
	Password string `gorm:"type:varchar(32);not null " json:"password" validate:"required,min=6,max=20" label:"密码"`   //密码
	Role int `gorm:"type:int;DEFAULT:2" json:"role" validate:"required,gte=2" label:"角色码"`  //身份码，0：管理员，1：阅读者
}


//CheckUser   查询用户名是否已被使用
func CheckUser(userName string)int {
	//新建user模型用于接收查询结果
	var users User
	//db.Where("username=?",userName).First(&users)
	//select id from users where username = userName ;      //返回查询到符合条件的的第一条数据给users模型
	db.Select("id").Where("username=?",userName).First(&users)
	if users.ID>0{  //如果id大于0 说明有数据存在
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
func GetUsers(pageSize,pageNo int) ([]User,int64) {
	var users []User
	var total int64

	//sql limit分页查询，offset：偏移量，要跳过的记录数
	offset := (pageNo -1)*pageSize
	if pageNo == -1 && pageSize == -1 {
		offset = -1
	}
	err = db.Limit(pageSize).Offset(offset).Find(&users).Count(&total).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil,0
	}
	return users,total
}

//BeforeSave   钩子函数
func (u *User)BeforeSave(_ *gorm.DB)(err error){
	//钩子函数，会在对user结构体进行操作的时候自动调用，对前端传入的密码进行加密可以与数据库中的密码进行匹配
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
	err = db.First(&user,id).Error
	if err != nil {
		return nil
	}
	return &user
}


//CheckLogin 登陆验证
func CheckLogin(username,password string)int{
	var user User
	//查询用户
	db.Where("username =?",username).First(&user)

	//无此用户ID
	if user.ID == 0 {
		return errmsg.ERROR_USER_NOT_EXIST
	}
	//密码错误
	if ScryptPw(password) != user.Password {
		return errmsg.ERROR_PASSWORD_WRONG
	}
	//权限不够
	if user.Role != 1 {
		return errmsg.ERROR_USER_NO_RIGHT
	}

	return errmsg.SUCCESS
}