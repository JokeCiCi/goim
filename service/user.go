package service

import (
	"errors"
	"fmt"
	"github.com/JokeCiCi/goim/model"
	"github.com/JokeCiCi/goim/util"
	"math/rand"
	"time"
)

type UserService struct{}

// 用户注册
func (s UserService) Register(mobile, plainpwd, nickname, avatar, sex string) (model.User, error) {
	tmp := model.User{}
	fmt.Println(DbEngine)
	// 检查手机号是否存在
	_, err := DbEngine.Where("mobile=?", mobile).Get(&tmp)
	if err != nil {
		return tmp, err
	}

	// 存在返回提示已注册
	if tmp.Id > 0 {
		return tmp, errors.New("该手机号已注册")
	}
	rand.Seed(time.Now().UnixNano())
	// 不存在拼接数据入库
	tmp.Mobile = mobile
	tmp.Nickname = nickname
	tmp.Avatar = avatar
	tmp.Sex = sex
	tmp.Salt = fmt.Sprintf("%06d", rand.Int31n(100000))
	tmp.Password = util.MakePasswd(plainpwd, tmp.Salt)
	tmp.Createat = time.Now()

	// 入库失败返回提示信息(恶意插入特殊字符、数据库连接失败)
	_, err = DbEngine.InsertOne(tmp)
	if err != nil {
		return tmp, err
	}

	// 入库成功返回用户信息
	return tmp, nil
}
