package main

import "fmt"
import (
	"strings"
	"reflect"
);

type Carrier struct {
	name, scac string;
	id int;
}

func (carrier *Carrier) String() string {
	return fmt.Sprintf("carrier %s with scac [%s], its cs4 id is %d", carrier.name, carrier.scac, carrier.id);
}

func main() {
	oocl := Carrier{name:"OOCL", scac:"OOLU",id:1};
	toLower(&oocl);
	fmt.Printf("carrier with lowercase: %s\n", oocl);

	cma := new(Carrier);
	cma.name = "CMA";
	cma.scac = "CGM_CMA";
	cma.id =2;
	toLower(cma);
	fmt.Printf("carrier with lowercase: %s\n", cma);

	var apl Carrier;
	apl.name = "APLU";
	apl.scac = "APL";
	apl.id =3;
	toLowerWithoutPointer(apl);
	// need to pass pointer to String(), as the String() definition only accept the pointer
	fmt.Printf("carrier without lowercase: %s\n", &apl);

	one := &Carrier{name:"ONEY", scac:"OceanNetwork Express",id:4};
	toLowerWithoutPointer(*one);
	fmt.Printf("carrier without lowercase: %s\n", one);

	aplType := reflect.TypeOf(apl);
	field := aplType.Field(0);
	fmt.Printf("apl type is %s, name is: %v\n ", aplType, field.Name);
}

func toLower(carrier *Carrier) {
	carrier.name = strings.ToLower(carrier.name);
	carrier.scac = strings.ToLower(carrier.scac);
}

func toLowerWithoutPointer(carrier Carrier) {
	carrier.name = strings.ToLower(carrier.name);
	carrier.scac = strings.ToLower(carrier.scac);
}
