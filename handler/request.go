package handler

import "fmt"

func errParamIsRequired(name, typ string) error {
	return fmt.Errorf("Param: %s (type: %s) is required", name, typ)
}

// CreateOpening

type CreateOpeningRequest struct {
	Role     string  `json: "role"`
	Company  string  `json: "company"`
	Location string  `json: "location"`
	Remote   *bool   `json: "remote"`
	Link     string  `json: "link"`
	Salary   float64 `json: "salary"`
}

func (request *CreateOpeningRequest) Validate() error {
	if request.Role == "" && request.Company == "" && request.Location == "" && request.Link == "" && request.Remote == nil && request.Salary <= 0 {
		return fmt.Errorf("Params body is empty or malformed")
	}
	if request.Role == "" {
		return errParamIsRequired("role", "string")
	}
	if request.Company == "" {
		return errParamIsRequired("company", "string")
	}
	if request.Location == "" {
		return errParamIsRequired("location", "string")
	}
	if request.Link == "" {
		return errParamIsRequired("link", "string")
	}
	if request.Remote == nil {
		return errParamIsRequired("remote", "bool")
	}
	if request.Salary <= 0 {
		return errParamIsRequired("salary", "float64")
	}
	
	return nil
}

type UpdateOpeningRequest struct {
	Role     string  `json: "role"`
	Company  string  `json: "company"`
	Location string  `json: "location"`
	Remote   *bool   `json: "remote"`
	Link     string  `json: "link"`
	Salary   float64 `json: "salary"`
}

func (request *UpdateOpeningRequest) Validate() error {
	if request.Role != "" || request.Company != "" || request.Location != "" || request.Remote != nil || request.Link != "" || request.Salary <= 0 {
		return nil
	}
	return fmt.Errorf("At least one valid must be provided")
}