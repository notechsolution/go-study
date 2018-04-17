package main

import (
	"fmt"
	"sort"
)

func main() {
	carrierMap := map[string]string{"OOLU": "OOCL", "KKLU": "K Line", "CMDU": "CMA_CGM"};
	carrierMap["COSU"] = "COSCON";
	sortKey1(carrierMap);
	sortWithSlice(carrierMap);

}

func sortKey1(carrierMap map[string]string)  {
	keys := make([]string, len(carrierMap));
	i := 0;
	for key := range carrierMap {
		keys[i] = key;
		i++;
	}
	sort.Strings(keys);
	fmt.Printf("sortKey1 - sorted key:%s\n", keys);
}

func sortWithSlice(carrierMap map[string]string)  {
	var keys []string;
	for key := range carrierMap {
		keys = append(keys, key);
	}
	sort.Strings(keys);
	fmt.Printf("sortKey2 - sorted key:%s\n", keys);
}

