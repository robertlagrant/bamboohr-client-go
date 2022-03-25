package bamboohr_client

import (
	"fmt"
	"testing"
)

func TestEmployeeDirectory(t *testing.T) {
	employees, err := EmployeeDirectory()
	if err != nil {
		t.Log("Couldn't list employees", err)
		t.Fail()
	}

	fmt.Printf("%#v\n", employees[0])
}

func TestListMyEmployees(t *testing.T) {
	employees, err := ListMyEmployees()
	if err != nil {
		t.Log("Couldn't list my employees", err)
		t.Fail()
	}

	fmt.Printf("%#v\n", employees[0])
}

func TestListAllEmployees(t *testing.T) {
	employees, err := ListEmployees()
	if err != nil {
		t.Log("Couldn't list all employees", err)
		t.Fail()
	}

	fmt.Printf("%#v\n", employees)
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
