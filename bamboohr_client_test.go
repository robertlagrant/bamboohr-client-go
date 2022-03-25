package bamboohr_client

import (
	"fmt"
	"testing"
)

func TestListEmployee(t *testing.T) {
	employees, err := ListEmployees()
	if err != nil {
		t.Log("Couldn't list employees", err)
		t.Fail()
	}

	fmt.Printf("%#v\n", employees[0])
}

// func TestGetAvailableFields(t *testing.T) {
// 	GetAvailableFields()
// }

func TestGetEmployee(t *testing.T) {
	me, err := GetEmployee(0) // Employee 0 is the employee associated with the API key
	if err != nil {
		t.Log("Couldn't get employee", err)
		t.Fail()
	}

	fmt.Printf("%#v\n", me)
}
