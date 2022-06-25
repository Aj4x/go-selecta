package selecta

import (
	"encoding/json"
	"reflect"
	"strings"
	"testing"
)

type Colour string

const (
	Black Colour = "Black"
	Green Colour = "Green"
	Blue  Colour = "Blue"
)

type Product struct {
	ProductCode string
	Description string
	Order       int
	Colour
}

var Products = []Product{
	{
		ProductCode: "111",
		Description: "Product 1",
		Order:       1,
		Colour:      Black,
	},
	{
		ProductCode: "222",
		Description: "Product 3",
		Order:       2,
		Colour:      Blue,
	},
	{
		ProductCode: "333",
		Description: "Product 3",
		Order:       3,
		Colour:      Blue,
	},
	{
		ProductCode: "444",
		Description: "Product 4",
		Order:       4,
		Colour:      Green,
	},
	{
		ProductCode: "555",
		Description: "Product 5",
		Order:       5,
		Colour:      Green,
	},
}

var Numbers = []int{1, 1, 1, 1, 2, 2, 2, 2, 2, 2, 3, 3, 3, 3, 3, 3, 4, 4, 5, 5, 6, 6, 7, 7, 7, 10, 9, 11, 16}

type Collection []string

var Collection1 = Collection{
	"test string 1",
	"test string 2",
	"test string 3",
	"test string 4",
	"test string 5",
}

func CheckResultType(t *testing.T, expectedType, actual reflect.Type) {
	if expectedType != actual {
		t.Fatal("incorrect return type")
	}
}

func CheckResultSize(t *testing.T, expectedSize int, actual int) {
	if expectedSize != actual {
		t.Fatal("result array incorrect size")
	}
}

func TestSelect(t *testing.T) {
	result, _ := Select[[]Product, []string](Products, func(p Product) (string, error) { return p.Description, nil })
	t.Log(result)
	var expectedType []string
	CheckResultSize(t, 5, len(result))
	CheckResultType(t, reflect.TypeOf(expectedType), reflect.TypeOf(result))
}

func TestSelectWhere(t *testing.T) {
	result, _ := SelectWhere[[]string, []int](Collection1, func(it string) (bool, int, error) {
		if strings.Contains(it, "1") {
			return false, 0, nil
		}
		return true, len(it), nil
	})
	t.Log(result)
	var expectedType []int
	CheckResultSize(t, 4, len(result))
	CheckResultType(t, reflect.TypeOf(expectedType), reflect.TypeOf(result))
}

func TestWhere(t *testing.T) {
	result, _ := Where(Numbers, func(i int) (bool, error) {
		return i > 3, nil
	})
	t.Log(result)
	var expectedType []int
	CheckResultSize(t, 13, len(result))
	CheckResultType(t, reflect.TypeOf(expectedType), reflect.TypeOf(result))
}

func TestAny(t *testing.T) {
	result1, _ := Any(Products, func(p Product) (bool, error) {
		return p.Order > 1, nil
	})
	t.Log(result1)
	result2, _ := Any(Products, func(p Product) (bool, error) {
		return p.Order < 1, nil
	})
	t.Log(result2)
	if !result1 || result2 {
		t.Fatal("Unexpected Any() outcome")
	}
}

func TestAll(t *testing.T) {
	result1, _ := All(Products, func(p Product) (bool, error) {
		return p.Order > 0, nil
	})
	t.Log(result1)
	result2, _ := All(Products, func(p Product) (bool, error) {
		return p.Order < 3, nil
	})
	t.Log(result2)
	if !result1 || result2 {
		t.Fatal("Unexpected All() outcome")
	}
}

func TestIndexOf(t *testing.T) {
	result1, _ := IndexOf(Products, func(p Product) (bool, error) {
		return p.Order == 5, nil
	})
	t.Log(result1)
	result2, _ := IndexOf(Products, func(p Product) (bool, error) {
		return p.Order == 10, nil
	})
	t.Log(result2)
	if result1 != 4 || result2 != -1 {
		t.Fatal("Unexpected IndexOf() outcome")
	}
}

func TestForEach(t *testing.T) {
	j := func(p Product) (string, error) {
		bytes, jErr := json.Marshal(p)
		return string(bytes), jErr
	}
	_ = ForEach(Products, func(p Product) error {
		jStr, err := j(p)
		if err != nil {
			return err
		}
		t.Log(jStr)
		return nil
	})
}

func TestGroupToMap(t *testing.T) {
	m, _ := GroupToMap(Products, func(p Product) Colour { return p.Colour })
	b, _ := json.MarshalIndent(m, "", "\t")
	t.Log(string(b))
	if len(m) != 3 {
		t.Fatal("incorrect length of map grouped slice")
	}
}

func TestMapToSlice(t *testing.T) {
	m, _ := GroupToMap(Products, func(p Product) Colour { return p.Colour })
	s, _ := MapToSlice(m, func(c Colour, p []Product) ([]Product, error) {
		return p, nil
	})
	b, _ := json.MarshalIndent(s, "", "\t")
	t.Log(string(b))
	if len(s) != 3 || len(s[0])+len(s[1])+len(s[2]) != 5 {
		t.Fatal("map to slice conversion has incorrect no items")
	}
}

func TestGroupBy(t *testing.T) {
	m, _ := GroupBy(Products, func(p Product) Colour { return p.Colour })
	b, _ := json.MarshalIndent(m, "", "\t")
	t.Log(string(b))
	if len(m) != 3 {
		t.Fatal("incorrect length of grouped slice")
	}
}
