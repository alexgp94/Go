package Tutoriales

import (
	"fmt"
)

const (
	first = iota
	second
)
const (
	third = iota
	fourth
)

func main() {
	fmt.Println("01-HolaMundo, mi name is alex!")
	var i int
	i = 42
	fmt.Println(i)
	var f float32 = 3.14
	fmt.Println(f)
	firstname := "alex"
	fmt.Println("my name is " + (firstname))
	b := true
	fmt.Println(b)
	c := complex(3, 5)
	fmt.Println(c)
	r, im := real(c), imag(c)
	fmt.Println(r, im)
	//punteros
	ab := 10
	bc := &ab
	fmt.Println(bc, *bc)
	ab = 3000
	fmt.Println(bc, *bc)
	ab = 20
	cp := ab * 9
	fmt.Println(cp)
	ab = -8
	ff := *bc + 15
	fmt.Println(ff, ab)

	const pi = 3.14
	fmt.Println(pi)
	const cc int = 4
	fmt.Println(cc + 3)
	fmt.Println(float32(cc) + 1.6)

	//fmt.Println(first,second,third,fourth)
	//Arrays forma larga
	var arr [3]int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	fmt.Println(arr)

	// froma cota
	arr2 := [3]int{1, 2, 3}
	fmt.Println(arr2)

	//slice
	a := []int{1, 2, 3}
	fmt.Println(a)
	a = append(a, 10, 14, 18)
	fmt.Println(a)
	bd := a[1:]
	cd := a[:2]
	dd := a[2:2]
	fmt.Println(bd, cd, dd)

	//map
	m := map[int]string{1: "alex", 2: "ander"}
	fmt.Println(m)
	fmt.Println(m[2])
	m[2] = "juan"
	fmt.Println(m[1], m[2])
	delete(m, 2)
	fmt.Println(m)

	//estructuras

	type user struct {
		id        int
		Firstname string
		lastname  string
	}

	u := user{
		id:        1,
		Firstname: "alex",
		lastname:  "gacia",
	}
	fmt.Println(u)

}
