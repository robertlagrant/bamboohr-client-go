package bamboohr_client

import (
	"fmt"
	"os"
	"testing"
)

func TestEmployeeDirectory(t *testing.T) {
	employees, err := EmployeeDirectory(makeConfig())
	if err != nil {
		t.Log("Couldn't list employees", err)
		t.Fail()
	}

	fmt.Printf("%#v\n", employees[0])
}

func TestListMyEmployees(t *testing.T) {
	employees, err := ListMyEmployees(makeConfig())
	if err != nil {
		t.Log("Couldn't list my employees", err)
		t.Fail()
	}

	fmt.Printf("%#v\n", employees[0])
}

func TestListAllEmployees(t *testing.T) {
	employees, err := ListEmployees(makeConfig())
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
	me, err := GetEmployee(makeConfig(), 0) // Employee 0 is the employee associated with the API key
	if err != nil {
		t.Log("Couldn't get employee", err)
		t.Fail()
	}

	fmt.Printf("%#v\n", me)
}

func makeConfig() Config {
	apiKey := os.Getenv("BAMBOOHR_API_KEY")
	tenantName := os.Getenv("BAMBOOHR_TENANT")
	includeSalary := true

	return Config{apiKey, tenantName, includeSalary}
}
