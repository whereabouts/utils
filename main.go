package main

import (
	"fmt"
	"github.com/whereabouts/utils/code"
	"github.com/whereabouts/utils/jwt"
	"github.com/whereabouts/utils/mapper"
	"github.com/whereabouts/utils/slice"
	"github.com/whereabouts/utils/timer"
	"time"
)

type Stu struct {
	Name string `json:"name" bson:"_name"`
	Age  int    `json:"age" bson:"_age"`
}

func main() {
	// example of slice
	fmt.Println("#######################################################")
	s := []int{1, 1, 2, 3, 4, 5, 6, 7, 8, 8, 0}
	fmt.Println("default s:", s)
	s = slice.DeleteInt(s, 8)
	fmt.Println("s delete one 8:", s)
	s = slice.DeleteInt(s, 1, true)
	fmt.Println("s delete all 1:", s)
	stus := []Stu{{"a", 1}, {"b", 2}, {"c", 3}}
	fmt.Println("struct slice stus:", stus)
	stus = slice.Delete(stus, Stu{"a", 1}).([]Stu)
	fmt.Println("stus delete one Stu{a 1}:", stus)
	// example of jwt
	fmt.Println("#######################################################")
	token := jwt.NewToken().SetOwner("korbin").String()
	fmt.Println(token)
	fmt.Println(jwt.Check(token))
	fmt.Println(jwt.Check(token + "1"))
	fmt.Println(jwt.IsExpire(token))
	fmt.Printf("%+v\n", jwt.GetPayload(token))
	token = jwt.Refresh(token)
	fmt.Println(token)
	fmt.Printf("%+v\n", jwt.GetPayload(token))
	// example of time
	fmt.Println("#######################################################")
	now := time.Now()
	fmt.Println(timer.Format(now))
	fmt.Println(timer.Format2Milli(now))
	t, err := timer.Parse("2021-12-12 12:12:12")
	if err != nil {
		return
	}
	fmt.Println(t.Unix() == 1639282332)
	fmt.Println(timer.Format(timer.Unix2Time(1639282332)))
	// example of map
	fmt.Println("#######################################################")
	fmt.Println(mapper.Json2Map(`{"name": "korbin", "age": 22}`))
	fmt.Println(mapper.Map2Json(map[string]string{"name": "korbin", "age": "22"}))
	stu := Stu{}
	err = mapper.Map2Struct(map[string]interface{}{"name": "korbin", "age": 22}, &stu)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(stu)
	fmt.Println(mapper.Struct2Map(stu))
	fmt.Println(timer.IsUnixTime(1614359148))
	fmt.Println(timer.IsUnixTime("1614359148"))
	fmt.Println(timer.IsStrTime("2021-02-27 02:35:17"))
	fmt.Println(timer.IsTime(now))
	// example of code
	fmt.Println("#######################################################")
	fmt.Println(code.String())
	fmt.Println(code.StrBlend(4))
	fmt.Println(code.VerifyStr("rfbd67", "RFbD67"))
	fmt.Println(code.VerifyStr("rfbd67", "RFbD67", true))
	fmt.Println(code.VerifyStr("RFbD67", "RFbD67", true))
	fmt.Println(code.ZN_CN())
	fmt.Println(code.NewZN_CN(4))
	fmt.Println(code.VerifyZN_CN("生却块正步成", "生却块正步成"))
	fmt.Println(code.VerifyZN_CN("生却块正步啊", "生却块正步成"))
	fmt.Println(code.Custom("123abcDEF却块正步"))
	fmt.Println(code.NewCustom(4, "123abcDEF却块正步"))
	fmt.Println(code.VerifyCustom("块d步a1d", "块d步A1D"))
	fmt.Println(code.VerifyCustom("块d步a1d", "块d步A1D", true))
	fmt.Println(code.VerifyCustom("块d步A1D", "块d步A1D", true))
}
