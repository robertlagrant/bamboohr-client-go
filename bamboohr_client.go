package bamboohr_client

import "fmt"
import "os"
import "strings"
import _ "github.com/joho/godotenv/autoload"

var (
	urlBase            = "https://api.bamboohr.com/api/gateway.php/"
	apiKey             = os.Getenv("BAMBOOHR_API_KEY")
	tenantName         = os.Getenv("BAMBOOHR_TENANT")
	employeeFieldNames = []string{"id", "employmentHistoryStatus", "jobTitle", "location", "department", "flsaCode", "division", "payChangeReason", "payRate", "paySchedule", "nationality"}
)

func ListEmployees() {
	urlPrefix := urlBase + tenantName
	listUrl := urlPrefix + "/v1/employees/directory"
	body, err := CallJsonApiMap(listUrl, apiKey, "GET")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(body["fields"])
	}
}

func GetEmployee(id int) {
	urlPrefix := urlBase + tenantName
	getUrl := urlPrefix + "/v1/employees/" + fmt.Sprint(id) + "?fields=" + strings.Join(employeeFieldNames, ",")
	fmt.Println(getUrl)
	body, err := CallJsonApiMap(getUrl, apiKey, "GET")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(body)
	}
}

func GetAvailableFields() ([]string, error) {
	urlPrefix := urlBase + tenantName
	getUrl := urlPrefix + "/v1/meta/lists/"
	fields, err := CallJsonApiList(getUrl, apiKey, "GET")
	if err != nil {
		return nil, fmt.Errorf("Could not retrieve fields. Reason: %s", err.Error())
	}
	var fieldNames []string
	for _, field := range fields {
		if str, ok := field["alias"].(string); ok {
			// fmt.Println(str)
			fieldNames = append(fieldNames, str)
		} else {
			return nil, fmt.Errorf("Could not retrieve fields. Reason: %s", err.Error())
		}
	}
	return fieldNames, nil
}
