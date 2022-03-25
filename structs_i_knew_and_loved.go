package bamboohr_client

type EmployeeListResponse struct {
	Fields    []Field    `json:"fields"`
	Employees []Employee `json:"employees"`
}
type Employee struct {
	ID               string `json:"id"`
	Department       string `json:"department"`
	DisplayName      string `json:"displayName"`
	Division         string `json:"division"`
	EmployeeNumber   string `json:"employeeNumber"`
	FirstName        string `json:"firstName"`
	HireDate         string `json:"hireDate"`
	JobTitle         string `json:"jobTitle"`
	LastName         string `json:"lastName"`
	Location         string `json:"location"`
	OriginalHireDate string `json:"originalHireDate"`
	PayRate          string `json:"payRate"`
	Status           string `json:"status"`
	Supervisor       string `json:"supervisor"`
	SupervisorID     string `json:"supervisorId"`
	SupervisorEID    string `json:"supervisorEId"`
	TerminationDate  string `json:"terminationDate"`
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
