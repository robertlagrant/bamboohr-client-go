package bamboohr_client

import "testing"

func TestListEmployee(t *testing.T) {
	ListEmployees()
}

// func TestGetAvailableFields(t *testing.T) {
// 	GetAvailableFields()
// }

func TestGetEmployee(t *testing.T) {
	GetEmployee(0)	// Employee 0 is the employee associated with the API key	
}
