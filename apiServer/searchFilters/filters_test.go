package searchFilters

import (
	"testing"
	"fmt"
)

func Test1(t *testing.T) {
	Fs := Filters{}
	F1 := ageFilter{minAge: 30, maxAge: 30}
	F2 := onlineFilter{uidSlice: append([]int{}, 1, 2, 3)}
	// F2 := onlineFilter{uidSlice: append([]int{}, 1)}
	// F2 := onlineFilter{uidSlice: []int{}}
	F3 := ratingFilter{minRating: 0, maxRating: 5}
	F4 := interestsFilter{interests: append([]string{}, "fun", "football")}

	Fs.filters = append(Fs.filters, F1, F2, F3, F4)
	println(Fs.PrepareQuery(""))
	fmt.Println("online filter type ", F2.getFilterType())
	fmt.Println("age filter type ", F1.getFilterType())
	fmt.Println("rating filter type ", F3.getFilterType())
}