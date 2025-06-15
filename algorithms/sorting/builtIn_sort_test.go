package sorting

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_BuiltinSort(t *testing.T) {
	people := []Person{
		{"Mars", 123},
		{"Iggy", 80},
		{"Clara", 35},
		{"Robin", 2},
		{"Frans", 42},
	}

	expectedSortedByAge1 := []Person{
		{"Robin", 2},
		{"Clara", 35},
		{"Frans", 42},
		{"Iggy", 80},
		{"Mars", 123},
	}

	expectedSortedByName1 := []Person{
		{"Clara", 35},
		{"Frans", 42},
		{"Iggy", 80},
		{"Mars", 123},
		{"Robin", 2},
	}

	people2 := []Person{
		{"Mars", 1},
		{"Iggy", 24},
		{"Clara", 6},
		{"Robin", 2},
		{"Frans", 42},
	}

	expectedSortedByAge2 := []Person{
		{"Mars", 1},
		{"Robin", 2},
		{"Clara", 6},
		{"Iggy", 24},
		{"Frans", 42},
	}

	expectedSortedByName2 := []Person{
		{"Clara", 6},
		{"Frans", 42},
		{"Iggy", 24},
		{"Mars", 1},
		{"Robin", 2},
	}

	t.Run("Custom", func(t *testing.T) {
		name := func(p1, p2 *Person) bool {
			return p1.Name < p2.Name
		}

		age := func(p1, p2 *Person) bool {
			return p1.Age < p2.Age
		}

		By(name).Sort(people)
		assert.Equal(t, expectedSortedByName1, people)

		By(age).Sort(people)
		assert.Equal(t, expectedSortedByAge1, people)
	})

	t.Run("Builtin standard", func(t *testing.T) {
		SortByName(people2)
		assert.Equal(t, expectedSortedByName2, people2)

		SortByAge(people2)
		assert.Equal(t, expectedSortedByAge2, people2)
	})

	t.Run("Custom standard", func(t *testing.T) {
		p := []Person{
			{"Mars", 123},
			{"Iggy", 80},
			{"Clara", 35},
		}

		expectedSortedByAge3 := []Person{
			{"Clara", 35},
			{"Iggy", 80},
			{"Mars", 123},
		}
		expectedSortedByName3 := []Person{
			{"Clara", 35},
			{"Iggy", 80},
			{"Mars", 123},
		}

		SortPersonByAgeNew(p)
		assert.Equal(t, expectedSortedByAge3, p)

		SortPersonByNameNew(p)
		assert.Equal(t, expectedSortedByName3, p)
	})
}

func Benchmark_BuiltinSort(b *testing.B) {
	people := []Person{
		{"Mars", 123},
		{"Iggy", 80},
		{"Clara", 35},
		{"Robin", 2},
		{"Frans", 42},
	}

	b.Run("SortPersonByAgeNew", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			SortPersonByAgeNew(people)
		}
	})

	b.Run("SortPersonByNameNew", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			SortPersonByNameNew(people)
		}
	})

	b.Run("SortByAge", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			SortByAge(people)
		}
	})

	b.Run("SortByName", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			SortByName(people)
		}
	})

	b.Run("Custom - Name", func(b *testing.B) {
		name := func(p1, p2 *Person) bool {
			return p1.Name < p2.Name
		}

		for i := 0; i < b.N; i++ {
			By(name).Sort(people)
		}
	})

	b.Run("Custom - Age", func(b *testing.B) {
		age := func(p1, p2 *Person) bool {
			return p1.Age < p2.Age
		}

		for i := 0; i < b.N; i++ {
			By(age).Sort(people)
		}
	})

}
