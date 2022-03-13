package bamboohr_client

import "fmt"
import "os"
import _ "github.com/joho/godotenv/autoload"


var (
	urlBase = "https://api.bamboohr.com/api/gateway.php/"
	apiKey = os.Getenv("BAMBOOHR_API_KEY")
	tenantName = os.Getenv("BAMBOOHR_TENANT")
)

func ListEmployee() {
	urlPrefix :=urlBase + tenantName
	listUrl := urlPrefix + "/v1/employees/directory"
	body, err := CallBamboohrJson(listUrl, apiKey, "GET")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(body["employees"])
	}
}


// type Employee struct {
// 	Id string `json:"id"`
// 	DisplayName string `json:"displayName"`
// 	FirstName string `json:"firstName"`
// 	LastName string `json:"lastName"`
// }

// type Field struct {
//     Id string `json:"id"`
// 	Type string `json:"type"`
//     Name string `json:"name"`
// }

// type BambooHrResponse struct {
// 	FieldList []Field `json"fields"`
// 	EmployeeList []Employee `json"employees"`
// }
