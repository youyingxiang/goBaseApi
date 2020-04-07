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
	t.Run("测试gorm多对多", testManyToMany)
}

func testPassword(t *testing.T) {
	inputPwd := "1234sss"
	userModel := &models.User{}
	user, err := userModel.GetUserById(43)
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

func testManyToMany(t *testing.T) {
	var user models.User
	Mysql.DB.Find(&user,1)
	var roles []models.Role
	Mysql.DB.Model(&user).Related(&roles)
	fmt.Println(user)
	fmt.Println(roles)

}
