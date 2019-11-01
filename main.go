package main

import (
	"flag"
	"fmt"
)

//rebase
//定义一个类型，用于增加该类型方法
type sliceValue []string

//flag为slice的默认值default is me,和return返回值没有关系
func (s *sliceValue) String() string {
	*s = sliceValue([]string{})
	return "It's none of my business"
}

func main() {
	//	var languages []string
	var languages sliceValue
	flag.Var(&languages, "slice", "I like programming `languages`")
	flag.Parse()

	//打印结果slice接收到的值
	fmt.Println(languages)
}
