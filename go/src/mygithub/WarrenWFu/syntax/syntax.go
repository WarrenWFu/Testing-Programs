package main

/*
//使用os.Args就可以获取到命令行参数，Args[0]为args.exe。
func main() {
		if len(os.Args) > 0 {
			v := os.Args[0]
			fmt.Printf(v)
		}
}
*/

/*
//使用func foo() (error){...} 来匹配type hook func() error函数类型
type hook func() error

func foo() error {
	log.Println("test")

	return nil
}

func main() {
	var bar hook = foo
	bar()
}
*/

/*
//如果map找不到则ok为false
func main() {
	m := make(map[string]int, 10)

	m["1"] = 1
	m["2"] = 2
	m["3"] = 3
	m["4"] = 4

	if _, ok := m["5"]; !ok {
		log.Printf("not found")
	}

	for k, v := range m {
		log.Printf("key[%s]value[%d]", k, v)
	}
}
*/

/*
//接口可以内嵌到struct中

//Fooer test
type Fooer interface {
	Foo()
}

//Bar test
type Bar struct {
	x int
	Fooer
}

func main() {
	bar := Bar{x: 10}
	var barB Fooer

	barB = bar

	log.Println(barB)
}
*/

/*
//超过slice范围的取值抛出panic
func main() {
	s := []string{"1"}
	fmt.Println("len(nil) = ", len(s[1]))
}
*/

/*
//使用strconv库的函数来进行字符串转换操作，执行时直接go run syntax.go 100
func main() {
	if v, err := strconv.Atoi(os.Args[1]); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(v)
	}
}
*/

/*
//编译错误，字符串初始化后不能修改
func main() {
	s := "hello"
	s[1] = 'X'

	fmt.Println(s)
}
*/

/*
//因为本源码编码是utf-8存储的时候是按照utf-8存储的，每个汉子三个字节
//使用range遍历的话，是转换成rune，代表的是unicode不是utf-8
func main() {
	s := "hello,世界"

	//	for i := 0; i < len(s); i++ {
	//		fmt.Println(i, s[i])
	//	}

	for i, ch := range s {
		fmt.Println(i, ch)
	}
}
*/
