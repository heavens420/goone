package main

import "fmt"

func main(){
	mystruct()
	mystruct2()
}

func mystruct()  {
	type Person struct {
		name string
		age int
		phone string
		addr string

	}

	p1 := Person{"zhgnsan",2,"111","beijing"}
	fmt.Println(p1)

	p2 := new(Person)
	fmt.Println(&p2)

	p2.name = "hh"
	p2.age = 9
	p2.phone = "222"
	p2.addr = "jii"

	fmt.Println(p2)

	//匿名结构体
	 var st1 = struct {
		 name string
		 age int
	 }{
	 	//莫忘逗号
	 	name:"nii",
	 	age: 32,
	 }

	 fmt.Println(st1)

	 //匿名字段 默认把变量类型当作属性名 所以变量类型不能重复
	type PP struct {
		string
		int
	}

	var st2 = PP{"hhh",233}
	fmt.Println(st2)

	//匿名函数 后面有括弧
	func(){
		fmt.Println("匿名函数")
	}()

}


//结构体嵌套
func mystruct2()  {
	type Book struct {
		name string
		price float32
	}

	type Student struct {
		name string
		age int
		book *Book  //赋值指针
	}

	var bk = Book{name:"六大名著",price:34.561}
	var st  = new(Student)
	st.name = "jj"
	st.age = 11
	st.book = &bk //取地址

	fmt.Printf("姓名：%s,年龄:%d,看的书：%s,书价:%.2f",st.name,st.age,st.book.name, st.book.price)




}