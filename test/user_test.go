package test

import (
	"fmt"
	"log"
	"rbac/app/models"
	Mysql "rbac/databases"
	"rbac/util"
	"testing"
)

func Test(t *testing.T) {
	//	t.Run("测试密码比对",testPassword)
	// 数据库迁移
	//t.Run("数据库迁移",testAutoMigrate)
	//t.Run("测试gorm多对多", testManyToMany)
	t.Run("测试用户的权限", testGetUserPermissions)
	//t.Run("测试用户的角色", testGetUserRoles)
	//t.Run("测试用户角色的所有权限",testGetUserRolePermissions)
}

func testPassword(t *testing.T) {
	inputPwd := "1234sss"
	user, err := models.GetUserById(43)
	if err != nil {
		log.Fatalln(err)
		return
	}
	isOK, err := user.CheckPassword(inputPwd)
	if err != nil {
		log.Fatalln(err)
		return
	}
	if !isOK {
		log.Fatalln(util.PasswordCheckError.Error())
		return
	}
	log.Println("测试通过")
}

type User struct {
	models.Base
	Name      string     `json:"name"`
	Languages []Language `gorm:"many2many:user_languages;"`
}

func (user *User) TableName() string {
	return "admin_users"
}

type Language struct {
	models.Base
	Name string
}

func testManyToMany(t *testing.T) {
	//var log models.Log
	//res := Mysql.DB.Find(&log,1)
	//fmt.Println(res);
	//var user models.User
	//Mysql.DB.Find(&user,1)
	//Mysql.DB.Model(&user).Related(&roles)
	//fmt.Println(user)
	//fmt.Println(roles)
	//var langs []Language
	//Mysql.DB.Find(&langs)
	//u1 := &User{
	//	Name: "user1",
	//
	//	Languages: langs,
	//}
	//Mysql.DB.Create(u1)
	var roles []models.Role
	var user models.User
	fmt.Println("请输入测试的用户id:")
	res := Mysql.DB.Find(&user, 1).RecordNotFound()
	if res {
		log.Fatalln("用户数据未找到！")
	}
	//var languages []Language
	res = Mysql.DB.Model(&user).Related(&roles, "Roles").RecordNotFound()
	if res {
		log.Fatalln("用户角色数据未找到！")
	}
	log.Println(roles)

}

func testAutoMigrate(t *testing.T) {
	Mysql.DB.AutoMigrate(&User{}, &Language{})
}

func testGetUserPermissions(t *testing.T) {
	user, e := models.GetUserById(53)
	if e != nil {
		t.Fatal(e)
	}

	permissions := user.GetAllPermissions()
	//if err != nil {
	//	log.Fatalln(err)
	//}
	t.Log(permissions)
}

func testGetUserRoles(t *testing.T) {
	var user models.User
	res := Mysql.DB.Find(&user, 53).RecordNotFound()
	if res {
		t.Error("用户数据未找到！")
	}
	roles, err := user.GetRoles()
	if err != nil {
		t.Error(err)
	}
	t.Log(roles[0].Name)
}

func testGetUserRolePermissions(t *testing.T) {
	//var user models.User
	//res := Mysql.DB.Find(&user, 53).RecordNotFound()
	//if res {
	//	t.Error("用户数据未找到！")
	//}
	//permissions, err := user.GetRolePermissions()
	//if err != nil {
	//	t.Error(err)
	//}
	//t.Log(permissions)
}
