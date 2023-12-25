package xormtest

import (
	"fmt"
	"testing"
	"time"

	"xorm.io/xorm"
)

type Student struct {
	Id         int64    `json:"id" xorm:"id pk autoincr comment(主键)"`
	Name       string   `json:"name" xorm:"name varchar(20) comment(姓名)"`
	Characters []string `json:"characters" xorm:"characters comment(个性)"`
	LuckyNums  []int32  `json:"luckyNums" xorm:"luckyNums comment(幸运数)"`

	CreatedAt time.Time  `json:"created_at" xorm:"created_at created comment(创建时间)"`
	UpdatedAt time.Time  `json:"updated_at" xorm:"updated_at updated comment(最后更新时间)"`
	DeletedAt *time.Time `json:"-" xorm:"deleted_at deleted comment(删除时间)"`
}

var mEngine *xorm.Engine

func init() {
	var err error
	mEngine, err = xorm.NewEngine("mysql", "root:root@tcp(192.168.15.42:3306)/jhl_uc")
	panic(err)
}

func TestXxxx(t *testing.T) {
	fmt.Print("--2")
}
