package model

import "time"

const (
	SEX_WOMEN  = "W" // 女性
	SEX_MAN    = "M" // 男性
	SEX_UNKNOW = "U" // 未知
)



// 用户信息
type User struct {
	Id       int64     `xorm:"pk autoincr bigint(20)" form:"id" json:"id"`  // 用户的ID
	Mobile   string    `xorm:"varchar(20)" form:"mobile" json:"mobile"`     // 用户手机号（唯一）
	Password string    `xorm:"varchar(40)" form:"password" json:"-"`        // 用户密码 = md5(plainpwd+slat)
	Avatar   string    `xorm:"varchar(150)" form:"avatar" json:"avatar"`    // 头像
	Sex      string    `xorm:"varchar(2)" form:"sex" json:"sex"`            // 性别
	Nickname string    `xorm:"varchar(20)" form:"nickname" json:"nickname"` // 昵称
	Salt     string    `xorm:"varchar(10)" form:"salt" json:"-"`            // 随机数，fmt.Sprintf("%06d", rand.Int31n(100000))
	Online   int       `xorm:"int(10)" form:"online" json:"online"`         // 是否在线
	Token    string    `xorm:"varchar(40)" form:"token" json:"token"`       // 鉴权使用，如前端用户登录，接入ws使用。/chat?id=1&token=x
	Memo     string    `xorm:"varchar(140)" form:"memo" json:"memo"`        // 用户描述
	Createat time.Time `xorm:"datetime" form:"createat" json:"createat"`    // 统计每天用户增量
}
