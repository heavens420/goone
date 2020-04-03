package main

import "fmt"

func main()  {

	var s = Student{
		Person: Person{"jj",2},
		school: "beijing",
	}
	//子类调用父类方法
	s.eat()
	//子类调用自己的方法
	s.say()
	//子类重写父类方法
	s.eat()
}

//父类
type Person struct {
	name string
	age int
}


//子类
type Student struct {
	Person
	school string
}

func (p *Person)eat()  {
	fmt.Println("父类的方法")
}

func (s *Student)say()  {
	fmt.Println("子类调用自己的方法")
}

func (s *Student)eat()  {
	fmt.Println("重写父类方法")
}