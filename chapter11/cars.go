package main

import "fmt"

type Any interface {}

type Car struct {
	Model string;
	Manufacturer string;
	BuildYear int;
}

type Cars []*Car;

func main() {
	// make some cars:
	ford := &Car{"Fiesta","Ford", 2008}
	bmw  := &Car{"XL 450", "BMW", 2011}
	merc := &Car{"D600", "Mercedes", 2009}
	bmw2 := &Car{"X 800", "BMW", 2008}
	// query:
	allCars := Cars([]*Car{ford, bmw, merc, bmw2})
	allNewBMWs := allCars.FindAll(func(car *Car) bool {
		return (car.Manufacturer == "BMW") && (car.BuildYear > 2010)
	})
	fmt.Println("AllCars: ", allCars)
	fmt.Println("New BMWs: ", allNewBMWs)
}

func (cars Cars) process (f func (car *Car)){
	for _,c := range cars {
		f(c);
	}
}

func(cs Cars) FindAll (f func(car * Car) bool) Cars{
	cars := make([]*Car, 0);
	cs.process(func(car *Car){
		if f(car){
			cars = append(cars, car);
		}
	})
	return cars;
}

func(cs Cars) String() string{
	result := "";
	for _, car := range cs {
		result += fmt.Sprintf("Model:%s, Manufacturer:%s, BuildYear: %d \n", car.Model, car.Manufacturer, car.BuildYear);
	}
	return result;
}




