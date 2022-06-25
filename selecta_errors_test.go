package selecta

import (
	"fmt"
	"strings"
	"testing"
)

func TestSelectError(t *testing.T) {
	errString := "expected Select error"
	result, err := Select[[]Product, []string](Products, func(p Product) (string, error) { return "", fmt.Errorf(errString) })
	if err == nil {
		t.Log(result)
		t.Fatal("Didn't received expected error")
	}
	t.Log(err.Error())
	if err.Error() != errString {
		t.Fatal("Incorrect SelectWhere error")
	}
}

func TestSelectWhereError(t *testing.T) {
	errString := "expected SelectWhere error"
	result, err := SelectWhere[[]string, []int](Collection1, func(it string) (bool, int, error) {
		if strings.Contains(it, "1") {
			return false, 0, nil
		}
		if strings.Contains(it, "test") {
			return false, 0, fmt.Errorf(errString)
		}
		return true, len(it), nil
	})
	if err == nil {
		t.Log(result)
		t.Fatal("Didn't received expected error")
	}
	t.Log(err.Error())
	if err.Error() != errString {
		t.Fatal("Incorrect SelectWhere error")
	}
}

func TestWhereError(t *testing.T) {
	errString := "expected Where error"
	result, err := Where(Numbers, func(i int) (bool, error) {
		return false, fmt.Errorf(errString)
	})
	if err == nil {
		t.Log(result)
		t.Fatal("Didn't received expected error")
	}
	t.Log(err.Error())
	if err.Error() != errString {
		t.Fatal("Incorrect Where error")
	}
}

func TestAnyError(t *testing.T) {
	errString := "expected Any error"
	result, err := Any(Products, func(p Product) (bool, error) {
		return false, fmt.Errorf(errString)
	})
	if err == nil {
		t.Log(result)
		t.Fatal("Didn't received expected error")
	}
	t.Log(err.Error())
	if result {
		t.Fatal("Unexpected Any() outcome")
	}
	if err.Error() != errString {
		t.Fatal("Incorrect Any error")
	}
}

func TestAllError(t *testing.T) {
	errString := "expected All error"
	result, err := All(Products, func(p Product) (bool, error) {
		return false, fmt.Errorf(errString)
	})
	if err == nil {
		t.Log(result)
		t.Fatal("Didn't received expected error")
	}
	t.Log(err.Error())
	if result {
		t.Fatal("Unexpected All() outcome")
	}
	if err.Error() != errString {
		t.Fatal("Incorrect All error")
	}
}

func TestIndexOfError(t *testing.T) {
	errString := "expected IndexOf error"
	result, err := IndexOf(Products, func(p Product) (bool, error) {
		return false, fmt.Errorf(errString)
	})
	if err == nil {
		t.Log(result)
		t.Fatal("Didn't received expected error")
	}
	t.Log(err.Error())
	if result != -1 {
		t.Fatal("Unexpected IndexOf() outcome")
	}
	if err.Error() != errString {
		t.Fatal("Incorrect IndexOf error")
	}
}

func TestForEachError(t *testing.T) {
	errString := "expected ForEach error"
	err := ForEach(Products, func(p Product) error {
		return fmt.Errorf(errString)
	})
	if err == nil {
		t.Fatal("Didn't received expected error")
	}
	if err.Error() != errString {
		t.Fatal("Incorrect ForEach error")
	}
}

func TestMapToSliceError(t *testing.T) {
	errString := "expected MapToSlice error"
	m, _ := GroupToMap(Products, func(p Product) Colour { return p.Colour })
	result, err := MapToSlice(m, func(c Colour, p []Product) ([]Product, error) {
		return nil, fmt.Errorf(errString)
	})
	if err == nil {
		t.Log(result)
		t.Fatal("Didn't received expected error")
	}
	t.Log(err.Error())
	if err.Error() != errString {
		t.Fatal("Incorrect MapToSlice error")
	}
}
