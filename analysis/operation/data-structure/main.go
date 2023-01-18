package main

import (
	"fmt"
	"reflect"
)

type Artist struct {
	Name, Genre string
	Songs       int
}

func newRelease(a *Artist) int {
	a.Songs++
	return a.Songs
}

// Maps is a collection of unordered pairs of key-value
var (
	elements = map[string]map[string]string{
		"H": map[string]string{
			"name":  "Hydrogen",
			"state": "gas",
		},
		"He": map[string]string{
			"name":  "Helium",
			"state": "gas",
		},
		"Li": map[string]string{
			"name":  "Lithium",
			"state": "solid",
		},
		"Be": map[string]string{
			"name":  "Beryllium",
			"state": "solid",
		},
		"B": map[string]string{
			"name":  "Boron",
			"state": "solid",
		},
		"C": map[string]string{
			"name":  "Carbon",
			"state": "solid",
		},
		"N": map[string]string{
			"name":  "Nitrogen",
			"state": "gas",
		},
		"O": map[string]string{
			"name":  "Oxygen",
			"state": "gas",
		},
		"F": map[string]string{
			"name":  "Fluorine",
			"state": "gas",
		},
		"Ne": map[string]string{
			"name":  "Neon",
			"state": "gas",
		},
	}
	arr              [5]int = [5]int{10, 20, 30}           //numeric array to define
	slc                     = []string{"fsfd", "dsfds"}    // define a slice which just extension of array but does not have strict length
	anEmptykeyValMap        = make(map[int]string)         // create an empty object with the help of make keyword
	anUnorderKeyVal         = map[interface{}]interface{}{ // to create an unordered misclleniuos object this can be done with the help interface
		"str": 3,
		23:    "jfklsdjfjls",
	}
)

func main() {

	me := &Artist{Name: "Matt", Genre: "Electro", Songs: 42}
	fmt.Printf("%s released their %dth song\n", me.Name, newRelease(me))
	fmt.Printf("%s has a total of %d songs", me.Name, me.Songs)

	if el, ok := elements["Li"]; ok {
		fmt.Println(el["name"], el["state"])
	}

	fmt.Println(len(anEmptykeyValMap))
	fmt.Println(len(anUnorderKeyVal))
	fmt.Println(len(arr))
	fmt.Println("Slice length is : ", len(slc))

	if v, ok := anUnorderKeyVal["str"]; ok {
		fmt.Println(v)
	}

	sliceExample()
}

func sliceExample() {

	var intSlice = new([50]int)[0:10] // another way of define a slice by using new keyword

	fmt.Println(reflect.ValueOf(intSlice).Kind())
	fmt.Printf("intSlice \tLen: %v \tCap: %v\n", len(intSlice), cap(intSlice))
	fmt.Println(intSlice)
}
