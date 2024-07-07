package struct_package

import ("fmt")

type React struct {
	Height int;
	Width int;
}

func (r *React) Area() int {
	return r.Width * r.Height
}

func TestingFunc(variable React) {
	fmt.Println("=============", variable)
}
