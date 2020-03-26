package lab

import (
	"log"
	"reflect"
)

// Test 实验
type Test struct {
	ID             *string    `bson:"some,omitempty"` //用户id
	Name           *string    //活动名
	Describe       *string    //描述
	Barbolas       *[]*string //图片描述
	PublisherRole  string     //发布接龙角色ID
	HeadImg        *string    //头图
	PrivacyType    *string    //传播隐私类型
	BizStartDate   string     //活动开始时间 2019-02-02 18:00:00
	BizEndData     string     //活动结束时间
	SignState      string     //签到状态
	PrivacyVisable *string    //参与者可见状态
	CanComment     *string    //是否可以留言
	BizType        *string    //业务类型
	BizStatus      *string    //业务状态
	TestArr        []interface{}
	TestMap        map[interface{}]interface{}
}

// TestFn 实验
func TestFn() {
	s := "String"
	sa := []*string{&s}

	t := &Test{
		Name:           &s,
		Describe:       &s,
		Barbolas:       &sa,
		PublisherRole:  "5e73018c324cecdcfda7bcac",
		HeadImg:        &s,
		PrivacyType:    &s,
		BizStartDate:   "String",
		BizEndData:     "String",
		SignState:      "关闭",
		PrivacyVisable: &s,
		CanComment:     &s,
		BizType:        &s,
		BizStatus:      &s,
		TestArr:        []interface{}{1, "a"},
		TestMap:        map[interface{}]interface{}{"a": "a", 0: 1},
	}

	log.Println(t)
	m := make(map[string]interface{})

	elem := reflect.ValueOf(t).Elem()
	// elemT := reflect.TypeOf(t).Elem()
	relType := elem.Type()
	// fmt.Println(telem.Field(i).Tag)
	for i := 0; i < relType.NumField(); i++ {

		log.Println(relType.Field(i).Name, elem.Field(i).Kind(), relType.Field(i).Tag.Get("bson"))

		m[relType.Field(i).Name] = elem.Field(i).Interface()
	}
	log.Println(m)

}