package main

import (
	"fmt"

	"github.com/golangg/src/channel_package"
	"github.com/golangg/src/struct_package"
)

func modifyArrayWithValue(a []int) {
	a[0] = 2;
}

func modifyPointerArray(a *[]int) {
	(*a)[0] = 2
}

func main() {
	// Difference between Array and Slice is
	// Array is copy value and type fixed
	// Slice is reference
	array_name := []int{1, 2, 3}

	array_name = append(array_name, 4, 5, 6)
	fmt.Println(array_name)

	// Delete i
	i := 1
	array_name = append(array_name[:i], array_name[i+1:]...)
	fmt.Println("After delete:", array_name)

	// -------- Struct ---------
	type Person struct {
		name   string
		points []int
	}

	per1 := Person{"Huy dep trai", []int{1, 3, 3}}
	fmt.Println(per1)
	per2 := per1
	per2.name = "Cuoi di haha"
	fmt.Println("Person 2:", per2)
	fmt.Println("Person 1:", per1)

	// ----------- Map -----------
	var mapInterface = make(map[interface{}]interface{})
	mapInterface["good_boy"] = "Huy"
	mapInterface["bad_boy"] = "Yuh"

	mapInterface1 := mapInterface
	mapInterface1["good_boy"] = "Faker"

	fmt.Println("mapInterface:", mapInterface)
	fmt.Println("mapInterface1:", mapInterface1)

	// Deep copy map
	deepObject := make(map[interface{}]interface{})
	for k, v := range mapInterface {
		deepObject[k] = v
	}

	deepObject["good_boy"] = "Who?"
	fmt.Println("mapInterface:", mapInterface)
	fmt.Println("deepObject:", deepObject)

	// Delete key
	delete(deepObject, "bad_boy")

	// Check key
	_, ok := deepObject["good_boy"]
	fmt.Println("Check key:", ok)

	// ------- Pointer ----------
	// Passing a slice to array
	arrs := []int{1, 2, 3}
	modifyArrayWithValue(arrs)
	fmt.Println("Modify array with value:", arrs)
	// Can't passing array
	arrayNormal := [3]int{1, 2, 3}
	modifyArrayWithValue(arrayNormal[:])
	fmt.Println("arrayNormal:", arrayNormal)
	// Pass pointer
	pointerArray := []int{1, 2, 3}
	modifyPointerArray(&pointerArray)
	fmt.Println("pointerArray:", pointerArray)

	// -------- Struct ----------- export
	newVar := struct_package.React{ Height: 4, Width: 2 }
	struct_package.TestingFunc(newVar)
	fmt.Println("====Area====", newVar.Area())

	// Channel
	fmt.Println("---------------------------------------")
	channel_package.ChannelSample()
}
