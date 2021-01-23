package ctrl

import (
	"fmt"
	"github.com/JokeCiCi/goim/model"
	"github.com/JokeCiCi/goim/service"
	"github.com/JokeCiCi/goim/util"
	"math/rand"
	"net/http"
	"time"
)

func Register(w http.ResponseWriter, r *http.Request) {
	// 解析请求参数
	err := r.ParseForm()
	if err != nil {
		util.Resp(w, -1, err.Error(), nil)
		return
	}
	rand.Seed(time.Now().UnixNano())

	mobile := r.Form.Get("mobile")
	plainpwd := r.Form.Get("passwd")
	nickname := fmt.Sprintf("user%05d", rand.Int31n(100000))
	avatar := ""
	sex := model.SEX_UNKNOW

	user, err := service.UserService{}.Register(mobile, plainpwd, nickname, avatar, sex)
	if err != nil {
		util.Resp(w, -1, err.Error(), nil)
	} else {
		util.Resp(w, 0, "", user)
	}
}

func Login(w http.ResponseWriter, r *http.Request) {
	// 解析请求参数
	err := r.ParseForm()
	if err != nil {
		util.Resp(w, -1, err.Error(), nil)
		return
	}
	mobile := r.Form.Get("mobile")
	passwd := r.Form.Get("passwd")

	// 逻辑判断
	loginok := false
	if mobile == "13600000000" && passwd == "1234" {
		loginok = true
	}

	// 返回json响应
	if loginok {
		data := make(map[string]interface{})
		data["id"] = "1"
		data["token"] = "test"
		util.Resp(w, -1, "", data)
		return
	} else {
		util.Resp(w, -1, "用户名或密码错误", nil)
		return
	}
}
