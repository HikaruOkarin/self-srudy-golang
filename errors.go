package main

func main() {
	x, err:=Foo() 
		if err!=nil {
			return
		}
	
}
func Foo() (int,error) {
	return 666,nil
}
func Bar() (int,error) {
	return 666,nil
}
func Baz() (int,error) {
	return 666,nil
}