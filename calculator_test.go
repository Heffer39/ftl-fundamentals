package calculator_test

import (
	"calculator"
	"fmt"
	"math/rand"
	"testing"
	"time"
)

type testCase struct {
	a, b        float64
	want        float64
	name        string
	errExpected bool
}

func TestAdd(t *testing.T) {
	t.Parallel()
	testCases := []testCase{
		{a: 2, b: 2, want: 4, name: "two plus two"},
		{a: 1, b: 1, want: 2, name: "one plus one"},
		{a: 5, b: 0, want: 5, name: "five plus zero"},
	}

	for _, tc := range testCases {
		got := calculator.Add(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("%s want %f, got %f", tc.name, tc.want, got)
		}
	}
}

func TestSubtract(t *testing.T) {
	t.Parallel()
	testCases := []testCase{
		{a: 2, b: 2, want: 0, name: "two minus two"},
		{a: 7, b: 1, want: 6, name: "seven minus one"},
		{a: 5, b: 0, want: 5, name: "five minus zero"},
	}

	for _, tc := range testCases {
		got := calculator.Subtract(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("%s want %f, got %f", tc.name, tc.want, got)
		}
	}
}

func TestMultiply(t *testing.T) {
	t.Parallel()
	testCases := []testCase{
		{a: 2, b: 3, want: 6, name: "two times three"},
		{a: 7, b: 1, want: 7, name: "seven times one"},
		{a: 5, b: 0, want: 0, name: "five times zero"},
	}

	for _, tc := range testCases {
		got := calculator.Multiply(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("%s want %f, got %f", tc.name, tc.want, got)
		}
	}
}

func TestDivide(t *testing.T) {
	t.Parallel()
	testCases := []testCase{
		{a: 6, b: 2, want: 3, name: "six divided by two", errExpected: false},
		{a: 10, b: 2, want: 5, name: "ten divided by two", errExpected: false},
		{a: 12, b: 3, want: 4, name: "twelve divided by three", errExpected: false},
		{a: 3, b: 2, want: 1.5, name: "three divided by two", errExpected: false},
		{a: 3, b: 0, want: 0, name: "zero division", errExpected: true},
	}

	for _, tc := range testCases {
		got, err := calculator.Divide(tc.a, tc.b)
		if (err != nil) != tc.errExpected {
			t.Errorf("%s - error: %s", tc.name, err)
		} else if tc.want != got {
			t.Errorf("%s want %f, got %f", tc.name, tc.want, got)
		}
	}
}

func TestAddRandom(t *testing.T) {
	t.Parallel()
	var testCases []testCase
	rand.Seed(time.Now().UTC().UnixNano())

	for i := 0; i < 100; i++ {
		a := rand.Float64() * 100
		b := rand.Float64() * 100
		want := a + b
		name := fmt.Sprintf("%f plus %f", a, b)
		testCases = append(testCases,
			testCase{a: a, b: b, want: want, name: name})
	}

	for _, tc := range testCases {
		got := calculator.Add(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("%s want %f, got %f", tc.name, tc.want, got)
		} /*else {
			fmt.Printf("Pass! %f %f %f\n", tc.a, tc.b, tc.want)
		}*/
	}
}
