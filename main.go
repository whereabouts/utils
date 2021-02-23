package main

import (
	"fmt"
	"github.com/whereabouts/utils/jwt"
	"github.com/whereabouts/utils/slice"
)

type Stu struct {
	Name string
	Age  int
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
}
