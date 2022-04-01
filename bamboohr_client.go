package bamboohr_client

import (
	"encoding/json"
	"fmt"
	"strings"

	_ "github.com/joho/godotenv/autoload"
)

var (
	urlBase            = "https://api.bamboohr.com/api/gateway.php/"
	employeeFieldNames = []string{"id", "employmentHistoryStatus", "jobTitle", "location", "department", "flsaCode", "division", "payChangeReason", "payRate", "paySchedule", "nationality"}
)

func EmployeeDirectory(config Config) ([]Employee, error) {
	listUrl := urlBase + config.Tenant + "/v1/employees/directory"
	body, err := CallJsonApi(listUrl, config.ApiKey, "GET")
	if err != nil {
		return nil, fmt.Errorf("Could not retrieve employees. Reason: %s", err.Error())
	}

	var response EmployeeListResponse
	json.Unmarshal([]byte(body), &response)

	return response.Employees, nil
}

func ListMyEmployees(config Config) ([]Employee, error) {
	listUrl := urlBase + config.Tenant + "/v1/reports/custom?format=JSON&onlyCurrent=false"
	payload := strings.NewReader("{\"filters\":{\"lastChanged\":{\"includeNull\":\"yes\"}},\"fields\":[\"supervisorId\",\"supervisorEId\",\"firstName\",\"lastName\",\"displayName\",\"payRate\",\"employmentHistoryStatus\",\"jobTitle\",\"location\",\"department\",\"payChangeReason\",\"paySchedule\",\"nationality\",\"employeeNumber\",\"birthday\",\"hireDate\",\"status\",\"terminationDate\",\"gender\",\"originalHireDate\",\"division\",\"createdByUserId\"]}")
	body, err := CallJsonApiWithPayload(listUrl, config.ApiKey, "POST", payload)
	if err != nil {
		return nil, fmt.Errorf("Could not retrieve my employees. Reason: %s", err.Error())
	}

	var response EmployeeListResponse
	json.Unmarshal([]byte(body), &response)

	employees := response.Employees

	for i, _ := range employees[:] {
		salary, currency, _ := parsePayRate(employees[i].PayRate)
		employees[i].PayRateParsedSalary = salary
		employees[i].PayRateParsedCurrency = currency
	}

	return employees, nil
}

func ListEmployees(config Config) ([]Employee, error) {
	myEmployees, err := ListMyEmployees(config)
	if err != nil {
		return nil, fmt.Errorf("Could not retrieve my employees. Reason: %s", err.Error())
	}
	myEmployeeIds := []string{}
	for _, employee := range myEmployees {
		myEmployeeIds = append(myEmployeeIds, employee.ID)
	}

	employeeDirectory, err := EmployeeDirectory(config)
	if err != nil {
		return nil, fmt.Errorf("Could not retrieve employee directory. Reason: %s", err.Error())
	}

	allEmployees := myEmployees

	for _, employee := range employeeDirectory {
		if !isElementExist(myEmployeeIds, employee.ID) {
			allEmployees = append(allEmployees, employee)
		}
	}

	return allEmployees, nil
}

func GetEmployee(config Config, id int) (*Employee, error) {
	getUrl := urlBase + config.Tenant + "/v1/employees/" + fmt.Sprint(id) + "?fields=" + strings.Join(employeeFieldNames, ",")
	body, err := CallJsonApi(getUrl, config.ApiKey, "GET")
	if err != nil {
		return nil, fmt.Errorf("Could not retrieve this employee. Reason: %s", err.Error())
	}

	var employee Employee
	json.Unmarshal([]byte(body), &employee)

	return &employee, nil
}

// func GetMyReportIds() (interface{}, error) {
// 	// me, err := GetEmployee(0)
// 	// if err != nil {
// 	// 	return nil, fmt.Errorf("Could not retrieve your details. Reason: %s", err.Error())
// 	// }

// 	// lineManagers := []string{fmt.Sprintf("%s %s", me.FirstName, me.LastName)}

// }

// func GetAvailableFields() ([]string, error) {
// 	urlPrefix := urlBase + tenantName
// 	getUrl := urlPrefix + "/v1/meta/lists/"
// 	fields, err := CallJsonApiList(getUrl, apiKey, "GET")
// 	if err != nil {
// 		return nil, fmt.Errorf("Could not retrieve fields. Reason: %s", err.Error())
// 	}
// 	var fieldNames []string
// 	for _, field := range fields {
// 		if str, ok := field["alias"].(string); ok {
// 			fieldNames = append(fieldNames, str)
// 		} else {
// 			return nil, fmt.Errorf("Could not retrieve fields. Reason: %s", err.Error())
// 		}
// 	}
// 	return fieldNames, nil
// }

func isElementExist(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}
