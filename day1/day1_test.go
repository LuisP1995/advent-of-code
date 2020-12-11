package main

import (
	"fmt"
	"reflect"
	"testing"
)

func testData() map[int]interface{} {
	var void interface{}
	testMap := map[int]interface{}{
		1721: void,
		979:  void,
		366:  void,
		299:  void,
		675:  void,
		1456: void,
	}
	return testMap
}

// TestLoadFile Checks if the load function is working as expected
func TestLoadFile(t *testing.T) {
	expected := testData()

	actual := make(map[int]interface{})
	LoadFile("numbersTest", actual)

	fmt.Println(expected)
	result := reflect.DeepEqual(actual, expected)
	if !result {
		t.Errorf("Values are not equal")
	} else {
		t.Logf("Values are the same, congrats")
	}
}

func TestCalc(t *testing.T) {
	expected := testData()
	expectedCode := 514579
	testCode := 0
	for key := range expected {
		result := Calc2020(key)
		if _, ok := expected[result]; ok {
			fmt.Println("Value found")
			testCode = (key * result)
			break
		}
	}

	fmt.Println(testCode)
	res := reflect.DeepEqual(expectedCode, testCode)
	if !res {
		t.Errorf("Values are not equal")
	} else {
		t.Logf("Codes are the same, congrats")
	}
}
