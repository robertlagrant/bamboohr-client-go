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

	fmt.Println(employees[0])
}

// func TestGetAvailableFields(t *testing.T) {
// 	GetAvailableFields()
// }

func TestGetEmployee(t *testing.T) {
	GetEmployee(0) // Employee 0 is the employee associated with the API key
}
