package main

import (
	"fmt"
	"reflect"
	"strings"
)

type animal interface {
	MakeSound()
}

type bear struct{}

func (b *bear) MakeSound() {
	fmt.Println("Groar!")
}

type duck struct{}

func (b *duck) MakeSound() {
	fmt.Println("Kwak!")
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
	//animals = append(animals, bear{}, duck{})   // Method Set of Bear struct literal does not implement animal interface
	animals = append(animals, &bear{}, &duck{}) // Method Set of *Bear pointer to struct DOES implement animal interface
	for _, a := range animals {
		a.MakeSound()
	}

	fmt.Println("\nPrint Method Set of Bear struct literal:")
	print(animals[0])
	fmt.Println("\nPrint Method Set of *Bear pointer to struct:")
	print(bear{})

	fmt.Println("\nMethod Set of Bear struct literal does not contain MakeSound(), but is callable:")
	b := bear{}
	b.MakeSound()
}
