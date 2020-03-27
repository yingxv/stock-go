package lab

import (
	"log"
	"reflect"
	"strings"
)

// Test 实验
type Test struct {
	ID            *string    `bson:"some,omitempty"` //用户id
	Name          *string    //活动名
	Describe      *string    //描述
	Barbolas      *[]*string //图片描述
	PublisherRole string     //发布接龙角色ID
	HeadImg       *string    //头图
	PrivacyType   *string    //传播隐私类型
	BizStartDate  string     //活动开始时间 2019-02-02 18:00:00
	BizEndData    string     //活动结束时间
	SignState     string     //签到状态

	PrivacyVisable *string //参与者可见状态
	CanComment     *string //是否可以留言
	BizType        *string //业务类型
	BizStatus      *string //业务状态
	TestArr        []interface{}
	TestMap        map[interface{}]interface{}
	TestNumber     int32
}

// TestFn 实验
func TestFn() {
	s := "String"
	sa := []*string{&s}

	t := &Test{
		Name:          &s,
		Describe:      &s,
		Barbolas:      &sa,
		PublisherRole: "5e73018c324cecdcfda7bcac",
		HeadImg:       &s,
		PrivacyType:   &s,
		BizStartDate:  "String",
		// BizEndData:     "String",
		SignState:      "关闭",
		PrivacyVisable: &s,
		CanComment:     &s,
		BizType:        &s,
		// BizStatus:      &s,
		TestArr: []interface{}{1, "a"},
		TestMap: map[interface{}]interface{}{"a": "a", 0: 1},
	}

	log.Println(t)
	m := make(map[string]interface{})

	// i := 1
	// slice1 := []int{0, 1, 3}
	Struct2Map(t, &m)

	log.Println(m)

}

// Struct2Map Struct转成Interface
func Struct2Map(s interface{}, m *map[string]interface{}) {
	structType := reflect.TypeOf(s)
	var elem reflect.Value
	if structType.Kind() == reflect.Ptr {
		elem = reflect.ValueOf(s).Elem()
	} else {
		elem = reflect.ValueOf(s)
	}

	relKind := elem.Kind() // 泛型
	if relKind != reflect.Struct {
		panic("错误")
	}

	relType := elem.Type() // 真实类型

	for i := 0; i < relType.NumField(); i++ {
		canEmpty := false
		tagsName := relType.Field(i).Name
		curElem := elem.Field(i)
		if tags, ok := relType.Field(i).Tag.Lookup("bson"); ok {
			tagsArr := strings.Split(tags, ",")
			for _, v := range tagsArr {
				if v == "omitempty" {
					canEmpty = true
				} else {
					tagsName = v
				}
			}
		}

		if curElem.Kind() == reflect.Ptr {
			if curElem.IsNil() {
				if canEmpty {
					continue
				}
				(*m)[tagsName] = nil
			} else {
				(*m)[tagsName] = curElem.Elem().Interface()
			}
		} else {
			(*m)[tagsName] = curElem.Interface()
		}
	}
}
