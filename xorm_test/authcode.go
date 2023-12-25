package xormtest

import (
	"fmt"
	"time"

	"xorm.io/xorm"
)

type Authcode struct {
	ID        int32      `json:"id" xorm:"id pk autoincr comment(主键)"`
	AppId     string     `json:"appid" xorm:"appid varchar(50) comment(AppId)"`
	AppSecret string     `json:"appsecret" xorm:"appsecret varchar(50) comment(AppSecret)"`
	Name      string     `json:"name" xorm:"name varchar(50) comment(名称)"`
	Expires   time.Time  `json:"expires" xorm:"expires datetime comment(失效时间)"`
	UserId    int64      `json:"userid" xorm:"userid BIGINT comment(用户id)"`
	Remark    string     `json:"remark" xorm:"remark varchar(255) comment(备注)"`
	Aucodes   []string   `json:"aucodes" xorm:"aucodes"`
	CreatedAt time.Time  `json:"created_at" xorm:"created_at created comment(创建时间)"`
	UpdatedAt time.Time  `json:"updated_at" xorm:"updated_at updated comment(最后更新时间)"`
	DeletedAt *time.Time `json:"-" xorm:"deleted_at deleted comment(删除时间)"`
}

func (*Authcode) TableName() string {
	return "auth_code"
}

func Demo1(db *xorm.Engine) {
	ins := &Authcode{
		AppId:     "t",
		AppSecret: "tt",
		Name:      "ttt",
		Expires:   time.Now(),
		UserId:    10,
		Remark:    "ss",
		Aucodes:   []string{"ll"},
	}
	_, err := db.Insert(ins)
	if err != nil {
		panic(err)
	}

	secret := "tt"

	item := Authcode{
		AppSecret: secret,
	}

	ok, err := db.Get(&item)
	if err != nil {
		panic(err)
	}
	if !ok {
		panic(fmt.Sprintf("记录{appsecret: %s} 不存在", secret))
	}

	fmt.Println(item)
}
