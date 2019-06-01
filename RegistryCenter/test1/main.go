/**
* Created by GoLand
* User: dollarkiller
* Date: 19-6-1
* Time: 下午8:49
* */
package main

import "fmt"

type Options struct {
	StrOption1 string
	StrOption2 string
	StrOption3 string
	IntOption1 int
	IntOption2 int
	IntOption3 int
}

type OptionFunc func(opts *Options)

func InitOption(opfs ...OptionFunc) {
	option := &Options{}
	for _,opf := range opfs {
		opf(option)
	}
	fmt.Println(option)
}

func WithStingOption1(str string) OptionFunc {
	return func(opts *Options) {
		opts.StrOption1 = str
	}
}

func WithStringOption2(str string) OptionFunc {
	return func(opts *Options) {
		opts.StrOption2 = str
	}
}

func main() {
	InitOption(
		WithStingOption1("hello"),
		WithStringOption2("asdasdasd"),
		)
}
