package main

import "fmt"

func main() {
	monday := WeekDay{num:0, ShortName:"MON"};
	tuesday := WeekDay{num:1, ShortName:"TUE"};
	wednesday := WeekDay{num:2, ShortName:"WED"};
	thursday := WeekDay{num:3, ShortName:"THU"};
	friday := WeekDay{num:4, ShortName:"FRI"};
	saturday := WeekDay{num:5, ShortName:"SAT"};
	sunday := WeekDay{num:6, ShortName:"SUN"};

	weekdays := &DayArray{[]WeekDay{wednesday,thursday,monday,tuesday,friday,saturday,sunday}}
	fmt.Printf("weekdays %v\n", weekdays);
	Sort(weekdays)
	fmt.Printf("sorted weekdays %v\n", weekdays);
	for _, weekday := range weekdays.days {
		fmt.Printf("%s ", weekday.ShortName)
	}
}

type Sorter interface {
	Len() int;
	Less(i, j int) bool;
	Swap(i, j int);
}

func Sort(data Sorter){
	for pass:=1; pass< data.Len();pass++ {
		for i:=0; i< data.Len()-pass; i++ {
			if data.Less(i+1, i){
				data.Swap(i+1, i);
			}
		}
	}
}

type WeekDay struct {
	num int;
	ShortName string;
	longName string;
}

type DayArray struct {
	days [] WeekDay;
}

func (t *DayArray) Len() int {
	return len(t.days);
}

func (t * DayArray) Less(i, j int) bool{
	return t.days[i].num < t.days[j].num;
}

func (t * DayArray) Swap(i,j int) {
	t.days[j], t.days[i] = t.days[i], t.days[j]
}
