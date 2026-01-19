package main

type Product struct {
	Name   string
	Brand  string
	MadeIn string
}

type Laptop struct {
	Product
	ModelName string
	OsSystem  string
	OsVersion string
}

func main() {

	tufGammingLaptop := new(Laptop)

	tufGammingLaptop.Name = "Tuf Gamming"
	tufGammingLaptop.ModelName = "F15"
	tufGammingLaptop.OsSystem = "Win 11"
	tufGammingLaptop.Brand = "Asus"
	tufGammingLaptop.MadeIn = "America"
	tufGammingLaptop.OsVersion = "26 H1"

}
