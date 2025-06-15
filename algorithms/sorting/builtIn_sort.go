package sorting

import (
	"cmp"
	"fmt"
	"slices"
	"sort"
	"strings"
)

type Person struct {
	Name string
	Age  int
}

type By func(p1, p2 *Person) bool

func (b By) Sort(people []Person) {
	ps := &PersonSorter{
		people: people,
		by:     b,
	}
	sort.Sort(ps)
}

type PersonSorter struct {
	people []Person
	by     func(p1, p2 *Person) bool
}

func (ps *PersonSorter) Len() int {
	return len(ps.people)
}

func (ps *PersonSorter) Swap(i, j int) {
	ps.people[i], ps.people[j] = ps.people[j], ps.people[i]
}

func (ps *PersonSorter) Less(i, j int) bool {
	return ps.by(&ps.people[i], &ps.people[j])
}

// =========== Standard Sort ===========

func SortByName(arr []Person) {
	sort.SliceStable(arr, func(i, j int) bool {
		return arr[i].Name < arr[j].Name
	})
}

func SortByAge(arr []Person) {
	sort.SliceStable(arr, func(i, j int) bool {
		return arr[i].Age < arr[j].Age
	})
}

func SortByBoth(arr []Person) {
	sort.SliceStable(arr, func(i, j int) bool {
		if arr[i].Name == arr[j].Name {
			return arr[i].Age < arr[j].Age
		}
		return arr[i].Name < arr[j].Name
	})
}

// =========== New Standard Sort ===========

func SortPersonByAgeNew(arr []Person) {
	slices.SortFunc(arr, func(i, j Person) int {
		if n := cmp.Compare(i.Age, j.Age); n != 0 {
			return n
		}
		return strings.Compare(i.Name, j.Name)
	})
}

func SortPersonByNameNew(arr []Person) {
	slices.SortFunc(arr, func(i, j Person) int {
		if n := strings.Compare(i.Name, j.Name); n != 0 {
			return n
		}
		return cmp.Compare(i.Age, j.Age)
	})
}

func main() {
	people := []Person{
		{"Mars", 123},
		{"Iggy", 80},
		{"Clara", 35},
		{"Robin", 2},
		{"Frans", 42},
	}

	people2 := []Person{
		{"Mars", 1},
		{"Iggy", 24},
		{"Clara", 6},
		{"Robin", 2},
		{"Frans", 42},
	}

	name := func(p1, p2 *Person) bool {
		return p1.Name < p2.Name
	}

	age := func(p1, p2 *Person) bool {
		return p1.Age < p2.Age
	}

	By(name).Sort(people)
	fmt.Println("By name: ", people)

	By(age).Sort(people)
	fmt.Println("By age", people)

	SortByName(people2)
	fmt.Println(people2)

	SortByAge(people2)
	fmt.Println(people2)
}
