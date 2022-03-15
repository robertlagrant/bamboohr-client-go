package bamboohr_client

type Employee struct {
	Id string `json:"id"`
	DisplayName string `json:"displayName"`
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
}

type BambooHrResponse struct {
	FieldList []Field `json"fields"`
	EmployeeList []Employee `json"employees"`
}

type Field struct {
	ID int32 `json:"id"`
	Manageable string `json:"manageable"`
	Multiple string `json:"multiple"`
	Name string `json:"name"`
	Options []FieldOption `json:"options"`
	Alias string `json:"employmentStatusHistory"`
}

type FieldOption struct {
	ID int32 `json:"id"`
    Archived string `json:"archived"`
    CreatedDate string `json:"createdDate"`
    ArchivedDate string `json:"archivedDate"`
	Name string `json:"name"`
}
