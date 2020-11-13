package main

import "fmt"

type reader interface {
	read(b []byte) (int, error)
}

// inteface file
type file struct {
	name string
}

func (file) read(b []byte) (int, error) {
	s := "<rss><channel><title>Going Go Programming</title></channel></rss>"
	copy(b, s)
	return len(s), nil
}

type pipe struct {
	name string
}

//第二个read 函数
func (pipe) read(b []byte) (int, error) {
	s := `{name: "hoanh", title: "developer"}`
	copy(b, s)

	return len(s), nil
}

func main() {
	f := file{"xml"}
	p := pipe{"json"}

	// 传入结构体
	retriever(f)

	retriever(p)
}

func retriever(r reader) error {
	// 声明一个byte数组
	d := make([]byte, 100)

	// 将read函数里面定义的字符串数据copy给d，并返回长度和错误信息
	len, err := r.read(d)

	if err != nil {
		return err
	}

	// 打印这个byte数组，打印的长度为返回的长度
	fmt.Println(string(d[:len]))
	return nil
}
