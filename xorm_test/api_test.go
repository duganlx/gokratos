package xormtest

import (
	"fmt"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
	"xorm.io/xorm"
)

type Student struct {
	Id         int64    `json:"id" xorm:"id pk autoincr comment(主键)"`
	Name       string   `json:"name" xorm:"name varchar(20) comment(姓名)"`
	Age        int32    `json:"age" xorm:"age comment(年龄)"`
	Characters []string `json:"characters" xorm:"characters comment(个性)"`
	LuckyNums  []int32  `json:"luckyNums" xorm:"luckyNums comment(幸运数)"`

	CreatedAt time.Time  `json:"created_at" xorm:"created_at created comment(创建时间)"`
	UpdatedAt time.Time  `json:"updated_at" xorm:"updated_at updated comment(最后更新时间)"`
	DeletedAt *time.Time `json:"-" xorm:"deleted_at deleted comment(删除时间)"`
}

func (*Student) TableName() string {
	return "students"
}

var mEngine *xorm.Engine

func init() {
	var err error
	mEngine, err = xorm.NewEngine("mysql", "root:root@tcp(192.168.15.42:3306)/jhl_uc")
	if err != nil {
		panic(err)
	}

	mEngine.Sync2(
		new(Student),
	)

	mEngine.ShowSQL(true)
}

func TestDBMetas(t *testing.T) {

	dbMetas, err := mEngine.DBMetas()
	assert.Nil(t, err)

	tbName := "students"
	columns := []string{"id", "name", "characters", "luckyNums", "created_at", "updated_at", "deleted_at"}
	for _, dbMeta := range dbMetas {
		if dbMeta.Name == tbName {
			// fmt.Printf("%+v\n", dbMeta)
			assert.Equal(t, columns, dbMeta.ColumnsSeq())
			assert.Equal(t, []string{"id"}, dbMeta.PrimaryKeys)
			assert.Equal(t, "InnoDB", dbMeta.StoreEngine)
			assert.Equal(t, "id", dbMeta.AutoIncrement)
		}
	}
}

func TestInsert(t *testing.T) {

	var n int64
	var err error
	var item *Student

	item = &Student{Name: "Tom", Characters: []string{"optimistic"}, LuckyNums: []int32{1, 3}}
	n, err = mEngine.Insert(item)
	assert.Nil(t, err)
	assert.Equal(t, n, int64(1))
	// fmt.Printf("%+v\n", item)

	item = &Student{Name: "XiaoMing", Characters: []string{"optimistic"}, LuckyNums: []int32{1, 3}}
	n, err = mEngine.InsertOne(item)
	assert.Nil(t, err)
	assert.Equal(t, n, int64(1))

	// len <= 150
	items := []*Student{
		{Name: "Jim1", Characters: []string{"optimistic"}, LuckyNums: []int32{4, 6}},
		{Name: "Jim2", Characters: []string{"pessimistic"}, LuckyNums: []int32{3, 5}},
	}
	n, err = mEngine.Insert(items)
	assert.Nil(t, err)
	assert.Equal(t, n, int64(2))

}

func TestSelectOne(t *testing.T) {
	// ID(interface{}) 传入一个主键字段的值，作为查询条件
	// Alias(string) 给Table设定一个别名
	// And(string, …interface{}) 和Where函数中的条件基本相同，作为条件
	// Or(interface{}, …interface{}) 和Where函数中的条件基本相同，作为条件
	// Get() 单条查询
	item := Student{}
	var ok bool
	var err error

	// ok, err = mEngine.Alias("o").Where("o.name = ?", "Tom").Get(&item)
	// assert.Nil(t, err)
	// assert.True(t, ok)
	// fmt.Println(item)

	// ok, err = mEngine.Where("name = ?", "Tom").And("id > ?", 15).Get(&item)
	// assert.Nil(t, err)
	// assert.True(t, ok)
	// fmt.Println(item)

	// ok, err = mEngine.ID(15).Get(&item)
	// assert.Nil(t, err)
	// assert.True(t, ok)
	// fmt.Println(item)

	item.Id = 15
	ok, err = mEngine.Get(&item)
	assert.Nil(t, err)
	assert.True(t, ok)

}

func TestSelectList(t *testing.T) {
	// Asc(…string) 指定字段名正序排序，可以组合
	// Desc(…string) 指定字段名逆序排序，可以组合
	// OrderBy(string) 按照指定的顺序进行排序
	// Select(string) 指定select语句的字段部分内容
	// In(string, …interface{}) 某字段在一些值中，这里需要注意必须是[]interface{}才可以展开，由于Go语言的限制，[]int64等不可以直接展开，而是通过传递一个slice。第二个参数也可以是一个*builder.Builder 指针。
	// SQL(string, …interface{}) 执行指定的Sql语句，并把结果映射到结构体
	// Cols(…string) 只查询或更新某些指定的字段，默认是查询所有映射的字段或者根据Update的第一个参数来判断更新的字段
	// Distinct(…string) 按照参数中指定的字段归类结果
	// Limit(int, …int) 限制获取的数目，第一个参数为条数，第二个参数表示开始位置，如果不传则为0
	// Top(int) Limit(int, 0)
	// Join(string,interface{},string) 第一个参数为连接类型，当前支持INNER, LEFT OUTER, CROSS中的一个值， 第二个参数为string类型的表名，表对应的结构体指针或者为两个值的[]string，表示表名和别名， 第三个参数为连接条件
	// GroupBy(string) Groupby的参数字符串
	// Having(string) Having的参数字符串
	items := make([]*Student, 0)

	err := mEngine.Where("name = ?", "Tom").Find(&items)
	assert.Nil(t, err)

	fmt.Println(items)
}

func TestCount(t *testing.T) {
	student := new(Student)

	total, err := mEngine.Where("id > ?", 10).Count(student)
	assert.Nil(t, err)
	assert.Greater(t, total, int64(1))
}

func TestExist(t *testing.T) {

	stu := Student{Id: -1}
	var has bool
	var err error

	has, err = mEngine.Exist(&stu)
	assert.Nil(t, err)
	assert.False(t, has)

	has, err = mEngine.SQL("select * from students where id < ?", 0).Exist()
	assert.Nil(t, err)
	assert.False(t, has)
}

func TestUpdate(t *testing.T) {
	// Cols(…string) 只查询或更新某些指定的字段，默认是查询所有映射的字段或者根据Update的第一个参数来判断更新的字段
	// AllCols() 查询或更新所有字段，一般与Update配合使用，因为默认Update只更新非0，非""，非bool的字段
	// MustCols(…string) 某些字段必须更新
	// Omit(…string) 和cols相反，此函数指定排除某些指定的字段。注意：此方法和Cols方法不可同时使用。
	//
	// 传对象只会更新 非空/非0 的字段, 传Map 有就更新

	var n int64
	var err error

	n, err = mEngine.ID(15).Update(&Student{Name: "Joker"})
	assert.Nil(t, err)
	assert.Equal(t, int64(1), n)

	n, err = mEngine.Table(new(Student)).ID(15).Update(map[string]interface{}{"age": 0})
	assert.Nil(t, err)
	assert.Equal(t, int64(1), n)
}

func TestDelete(t *testing.T) {
	stu := new(Student)

	var n int64
	var err error

	// 假删
	n, err = mEngine.ID(16).Delete(stu)
	assert.Nil(t, err)
	assert.Equal(t, int64(1), n)

	// 真删
	n, err = mEngine.ID(17).Unscoped().Delete(stu)
	assert.Nil(t, err)
	assert.Equal(t, int64(0), n)
}

func TestTransaction(t *testing.T) {
	session := mEngine.NewSession()
	defer session.Close()

	var n int64
	var err error

	err = session.Begin()
	assert.Nil(t, err)

	stu1 := Student{Name: "lvx", Age: 12, Characters: []string{"audacious"}, LuckyNums: []int32{3, 5}}
	n, err = session.Insert(&stu1)
	assert.Nil(t, err)
	assert.Equal(t, n, int64(1))

	err = session.Commit()
	assert.Nil(t, err)

}
