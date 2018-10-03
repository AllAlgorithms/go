package opeval

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)

type testpair struct {
	eq     string
	expRes float64
}

func TestEvaluate(t *testing.T) {
	t.Run("Simple expressions", func(t *testing.T) {
		var tests = []testpair{
			{
				"1 + 2",
				float64(1) + float64(2),
			},
			{
				"3 * 4",
				float64(3) * float64(4),
			},
			{
				"5 - 6",
				float64(5) - float64(6),
			},
			{
				"7 / 8",
				float64(7) / float64(8),
			},
		}

		for _, test := range tests {
			t.Run(test.eq, func(t *testing.T) {
				res, err := Evaluate(test.eq)
				if err != nil {
					t.Errorf("No error expected, got %+v", err)
				}

				if res != test.expRes {
					t.Errorf("Expected %f, got %f", test.expRes, res)
				}
			})
		}
	})

	t.Run("float64 expressions", func(t *testing.T) {
		var tests = []testpair{
			{
				"1.1*0.8",
				float64(1.1) * float64(0.8),
			},
			{
				"1/3 - 0.333333333333",
				float64(1)/float64(3) - float64(0.333333333333),
			},
			{
				"2. + 8.",
				float64(2.) + float64(8.),
			},
		}

		for _, test := range tests {
			t.Run(test.eq, func(t *testing.T) {
				res, err := Evaluate(test.eq)
				if err != nil {
					t.Errorf("No error expected, got %+v", err)
				}

				if res != test.expRes {
					t.Errorf("Expected %f, got %f", test.expRes, res)
				}
			})
		}
	})

	t.Run("Complex expressions", func(t *testing.T) {
		var tests = []testpair{
			{
				"(10.8 * (1.2 - 0.87) / 14) * 10",
				(float64(10.8) * (float64(1.2) - float64(0.87)) / float64(14)) * float64(10),
			},
			{
				"((((1500)+127)*900)-1)",
				((((float64(1500)) + float64(127)) * float64(900)) - float64(1)),
			},
			{
				"(1500+(127*(900-(1))))",
				(float64(1500) + (float64(127) * (float64(900) - (float64(1))))),
			},
		}

		for _, test := range tests {
			t.Run(test.eq, func(t *testing.T) {
				res, err := Evaluate(test.eq)
				if err != nil {
					t.Errorf("No error expected, got %+v", err)
				}

				if res != test.expRes {
					t.Errorf("Expected %f, got %f", test.expRes, res)
				}
			})
		}
	})

	t.Run("Incomplete expressions", func(t *testing.T) {
		var tests = []string{
			"1 + 2 +",
			"13*-",
			"/ 19 -",
		}

		for _, test := range tests {
			t.Run(test, func(t *testing.T) {
				_, err := Evaluate(test)
				if err == nil {
					t.Error("Error expected")
				}

				expected := errors.New("opeval - invalid expression")
				if !reflect.DeepEqual(expected, err) {
					t.Errorf("Expected %+v, got %+v", expected, err)
				}
			})
		}
	})

	t.Run("Bad float64 expressions", func(t *testing.T) {
		var tests = []string{
			"1..3 + 1",
			"8-3..1",
			"1..4 * 5....5",
		}

		for _, test := range tests {
			t.Run(test, func(t *testing.T) {
				_, err := Evaluate(test)
				if err == nil {
					t.Error("Error expected")
				}

				expected := errors.New("opeval - invalid expression")
				if !reflect.DeepEqual(expected, err) {
					t.Errorf("Expected %+v, got %+v", expected, err)
				}
			})
		}
	})

	t.Run("Invalid expressions", func(t *testing.T) {
		var tests = [][2]string{
			[2]string{
				"a+b",
				"a",
			},
			[2]string{
				"10x - 1",
				"x",
			},
			[2]string{
				"not a valid expression",
				"n",
			},
		}

		for _, test := range tests {
			t.Run(test[0], func(t *testing.T) {
				_, err := Evaluate(test[0])
				if err == nil {
					t.Error("Error expected")
				}

				expected := fmt.Errorf("opeval - invalid character '%s'", test[1])
				if !reflect.DeepEqual(expected, err) {
					t.Errorf("Expected %+v, got %+v", expected, err)
				}
			})
		}
	})

	t.Run("Bad parenthesis configuration", func(t *testing.T) {
		var tests = []string{
			"(1+2",
			"(1+2))",
			")(1+2)(",
		}

		for _, test := range tests {
			t.Run(test, func(t *testing.T) {
				_, err := Evaluate(test)
				if err == nil {
					t.Error("Error expected")
				}

				expected := errors.New("opeval - invalid parenthesis configuration")
				if !reflect.DeepEqual(expected, err) {
					t.Errorf("Expected %+v, got %+v", expected, err)
				}
			})
		}
	})

	t.Run("Division by zero", func(t *testing.T) {
		expression := "1 / 0"
		_, err := Evaluate(expression)
		if err == nil {
			t.Error("Error expected")
		}

		expected := errors.New("opeval - division by 0")
		if !reflect.DeepEqual(expected, err) {
			t.Errorf("Expected %+v, got %+v", expected, err)
		}
	})
}
