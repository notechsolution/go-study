package main

import (
	"fmt"
)

func main() {
	carrierMap := map[string]string {"OOLU": "OOCL", "KKLU": "K Line", "CMDU": "CMA_CGM"};
	carrierMap["COSU"] = "COSCON";
	fmt.Printf("carrierMap:%s \n", carrierMap);

	carrier, isExist := carrierMap["OOLU"];
	fmt.Printf("is OOLU exist %t, its name is %s\n", isExist, carrier);

	for key := range carrierMap {
		fmt.Printf("key: %s\n",key);
	}

	capitals := map[string] string {"France":"Paris", "Italy":"Rome", "Japan":"Tokyo" }
	for key := range capitals {
		fmt.Println("Map item: Capital of", key, "is", capitals[key]);
	}

	//test map copy, reference assignment only
	copyMap := carrierMap;
	fmt.Printf("copyMap initial values: %s\n", copyMap);
	copyMap["OOLU"] = "Orient Oversea Container Liner";
	fmt.Printf("copyMap now is: %s\n", copyMap)
	fmt.Printf("carrierMap now is: %s\n", carrierMap)

}

