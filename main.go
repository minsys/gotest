package main

import (
	"fmt"
	"reflect"
	"strings"
)

type animal interface {
	MakeSound()
	SayName()
}

type bear struct{}

func (b *bear) MakeSound() {
	fmt.Println("Groar!")
}

func (b bear) SayName() {
	fmt.Println("Bear")
}

type duck struct{}

func (d *duck) MakeSound() {
	fmt.Println("Kwak!")
}

func (d duck) SayName() {
	fmt.Println("Duck")
}

// Print prints the method set of the value x.
func print(x interface{}) {
	v := reflect.ValueOf(x)
	t := v.Type()
	fmt.Printf("type %s\n", t)

	for i := 0; i < v.NumMethod(); i++ {
		methType := v.Method(i).Type()
		fmt.Printf("func (%s) %s%s\n", t, t.Method(i).Name,
			strings.TrimPrefix(methType.String(), "func"))
	}
}

func main() {
	animals := []animal{}

	// Method Set of bear struct literal does not implement animal interface,
	// it does not include the methods defined on a pointer receiver.
	//animals = append(animals, bear{}, duck{})

	// Method Set of *bear pointer to struct DOES implement animal interface,
	// it includes the methods defined on a non-pointer receiver.
	animals = append(animals, &bear{}, &duck{})
	for _, a := range animals {
		a.MakeSound()
	}

	fmt.Println("\nPrint Method Set of bear struct literal:")
	print(animals[0])
	fmt.Println("\nPrint Method Set of *bear pointer to struct:")
	print(bear{})

	fmt.Println("\nMethod Set of bear struct literal does not contain MakeSound(), but is callable:")
	b := bear{}
	b.MakeSound()

	fmt.Println("\nMethod Set of *bear pointer does include the methods defined on non-pointer receivers:")
	bp := &bear{}
	bp.SayName()
}
