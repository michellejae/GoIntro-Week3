package main

import (
	"fmt"
	"log"
	"math"
	"sort"
)

func main() {
	loc := Location{
		Lat: 33.3,
		Lng: 22.2,
	}

	fmt.Println("loc", loc)           // just give us values
	fmt.Printf("loc (+): %+v\n", loc) // gives us keys (%+v) dont' forget plus sign
	fmt.Printf("loc (#): %#v\n", loc) // kind of gives us type so if a number is actual a string it will be in ""

	loc2 := Location{12.3, 4.56} // can omit fields. if you leave one out, it will just be 0
	fmt.Printf("loc2 (#): %#v\n", loc2)

	loc3 := &Location{12.3, 4.56} // Pointer to Location
	fmt.Printf("loc3: lat=%f, lng=%f\n", loc3.Lat, loc3.Lng)

	c1 := Car{
		Plate:    "95",
		Location: Location{1.1, 2.3},
	}

	fmt.Printf("c1: %#v\n", c1)
	fmt.Printf("c1 lat: %f\n", c1.Lat) // can access the field (key) directly from c1

	c1.wrongMove(10, 20)
	fmt.Printf("c1 (move): %#v\n", c1) // this will still print 1.1 and 2.3 as lat and long cause below we passed by value

	c1.Move(10, 20)
	fmt.Printf("c1 (move): %#v\n", c1) // this will work cause below we passed by vlaue

	c2, err := NewCar("007", 45, 56)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("c2: %#v\n", c2) // type is a pointer since we pass in the address to the pointer in below NewCar function

	// do not HAVE to use NewCar func
	c3 := Car{"", Location{-200, 300}}
	fmt.Printf("c3: %#v\n", c3)

	cars := []Car{
		Car{"1", Location{1, 1}},
		Car{"2", Location{2, 2}},
		Car{"3", Location{3, 3}},
		Car{"4", Location{4, 4}},
	}

	customer := Location{2.2, 2.2}
	cs := Cars{
		cars: cars,
		loc:  customer,
	}

	sort.Sort(&cs)
	fmt.Println("cars", cars)

	var n Number = 17
	fmt.Println(n)
}

type Number int

// implement fmt.Stringer interface
func (n Number) String() string {
	return fmt.Sprintf("<Numer: %d>", n)
}

// Accept interfacces, return concrete types
// Use Small Interfaces

type Cars struct {
	cars []Car
	loc  Location
}

func (c *Cars) Len() int      { return len(c.cars) }
func (c *Cars) Swap(i, j int) { c.cars[i], c.cars[j] = c.cars[j], c.cars[i] }
func (c *Cars) Less(i, j int) bool {
	return c.loc.Distance(c.cars[i].Location) < c.loc.Distance(c.cars[j].Location)
}

func NewCar(plate string, lat, lng float64) (*Car, error) {
	if plate == "" {
		return nil, fmt.Errorf("empty plate")
	}
	// TODO: Check validity of lat, lng (like between 180 - 180)
	c := Car{
		Plate:    plate,
		Location: Location{Lat: lat, Lng: lng},
	}

	// have to pass & for address cause we pointer above in  function
	return &c, nil

}

// REMEMBER to pass by pointer rather than value
// function defition or method
// have to rememer to put recevir then Type
func (c *Car) Move(lat, lng float64) {
	c.Lat = lat
	c.Lng = lng
}

// PASS BY VALUE
// function defition or method
// have to rememer to put recevir then Type
func (c Car) wrongMove(lat, lng float64) {
	c.Lat = lat
	c.Lng = lng
}

type Car struct {
	Plate string
	//Location Location
	Location // car embeds Location type
}

// method on Location
func (loc Location) Distance(other Location) float64 {
	dLat := loc.Lat - other.Lat
	dLng := loc.Lng - other.Lng

	return math.Sqrt(dLat*dLat + dLng*dLng)
}

type Location struct {
	Lat float64
	Lng float64
}
