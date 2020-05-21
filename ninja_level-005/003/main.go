package main

import "fmt"

func main() {
	type vehicle struct {
		doors int
		color string
	}

	type truck struct {
		vehicle
		isFourWheel bool
	}

	type sedan struct {
		vehicle
		isLuxury bool
	}

	t1 := truck{
		isFourWheel: true,
		vehicle: vehicle{
			color: "red",
			doors: 4,
		},
	}

	s1 := sedan{
		isLuxury: true,
		vehicle: vehicle{
			color: "blue",
			doors: 2,
		},
	}

	fmt.Println(t1)
	fmt.Println(s1)

	fmt.Println(t1.doors)
	fmt.Println(s1.doors)

}
