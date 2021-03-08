package main

import "fmt"

type F interface {
	f()
}

type S1 struct{}

func (s S1) f() {

}

type S2 struct{}

func (s *S2) f() {

}

type errNotFound struct {
	file string
}

func (e errNotFound) Error() string {
	return fmt.Sprintf("file %q not found", e.file)
}

func open(file string) error {
	return errNotFound{file: file}
}

func use(){
	if err :=open("testfile.txt"); err!=nil {
		if _,ok := err.(errNotFound); ok{
			fmt.Errorf()
		}else{
			panic(unknown error)
		}
	}
}
func main() {
	s1Val := S1{}
	s1Ptr := &S1{}
	s2Val := S2{}
	s2Ptr := &S2{}

	var i F

	i = s1Val
	i = s1Ptr
	i = s2Ptr
	i = s2Val

}
