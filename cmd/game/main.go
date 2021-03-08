package main

import "github.com/ymohl-cl/dice/pkg/loto"

var (
	files = []string{
		"../../resources/loto_201902.csv",
		"../../resources/loto_201911.csv",
		"../../resources/loto2017.csv",
		"../../resources/nouveau_loto.csv",
	}
	//"../../resources/loto.csv",
)

func main() {
	var l loto.Loto
	var err error

	if l, err = loto.New(files); err != nil {
		panic(err)
	}
	l.Print()
	return
}

/*
	nbFace = 6
	if d, err = dice.New(nbFace); err != nil {
		panic(err)
	}
	for i := 0; i < 1000; i++ {
		_ = d.Throw()
		// fmt.Printf("throw number %d with result %d\n", i, v)
	}

	fmt.Println("verification")
	for i := 0; i < int(nbFace); i++ {
		v := d.Result(int32(i) + 1)
		fmt.Printf("%d face was returned %d times\n", i+1, v)
	}



*/
