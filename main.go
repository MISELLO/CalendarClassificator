package main

import (
	"fmt"
	"time"
)

func main() {
	const layout = "02/01/2006" // The layout used by Go: Jan 2, 2006 at 3:04pm (MST)
	const partValue1 = "29/02/" // To know if it is a leap year
	const partValue2 = "01/01/" // First day of the year
	const init = 2000           // First year to check
	const end = 2100            // Last year to check

	var yearsByType [14][]int

	for i := init; i <= end; i++ {
		index := 0

		// Stupid, non orthodox way of calculating a leap year
		value := partValue1 + fmt.Sprint(i)
		t, err := time.Parse(layout, value)
		if err == nil {
			// Leap year
			index += 7
		}

		value = partValue2 + fmt.Sprint(i)
		t, _ = time.Parse(layout, value)
		wd := int(t.Weekday()) // 0 is Sunday and we want it to be Monday
		if wd == 0 {
			wd = 7
		}
		wd--
		index += wd
		yearsByType[index] = append(yearsByType[index], i)
	}

	print(yearsByType)
}

func print(d [14][]int) {
	for i := 0; i < len(d); i++ {
		if i < 7 {
			fmt.Printf("\n Non leap years starting with %s:", time.Weekday((i+1)%7))
		} else {
			fmt.Printf("\n Leap years starting with %s:", time.Weekday((i+1)%7))
		}

		for j := 0; j < len(d[i]); j++ {
			fmt.Printf(" %d", d[i][j])
		}
	}
	fmt.Println()
	fmt.Println()
}
