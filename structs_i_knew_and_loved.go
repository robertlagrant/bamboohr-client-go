package bamboohr_client

type EmployeeListResponse struct {
	Fields    []Field    `json:"fields"`
	Employees []Employee `json:"employees"`
}
type Employee struct {
	Id          string `json:"id"`
	DisplayName string `json:"displayName"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Location    string `json:"location"`
	Department  string `json:"department"`
	JobTitle    string `json:"jobTitle"`
	Supervisor  string `json:"supervisor"`
}
type Field struct {
	ID         int32         `json:"id"`
	Manageable string        `json:"manageable"`
	Multiple   string        `json:"multiple"`
	Name       string        `json:"name"`
	Options    []FieldOption `json:"options"`
	Alias      string        `json:"employmentStatusHistory"`
}

type FieldOption struct {
	ID           int32  `json:"id"`
	Archived     string `json:"archived"`
	CreatedDate  string `json:"createdDate"`
	ArchivedDate string `json:"archivedDate"`
	Name         string `json:"name"`
}
