/**
 * @Author: DollarKiller
 * @Description:
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 21:53 2019-09-23
 */
package main

import "log"

type App struct {
	StrOption1 string
	StrOption2 string
}

type config func(*App)

func init_app(configs ...config) App {
	a := App{}
	for _,k := range configs {
		k(&a)
	}

	return a
}

func setStr1(str string) config {
	return func(ops *App) {
		ops.StrOption1 = str
	}
}

func setStr2(str string) config {
	return func(app *App) {
		app.StrOption2 = str
	}
}

func main() {
	app := init_app(setStr1("你好"), setStr2("golang"))

	log.Println(app)
}
