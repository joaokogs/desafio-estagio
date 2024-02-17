package handler

import "fmt"

func errParamIsRequired(name, typ string) error {
	return fmt.Errorf("parms : %s (type: %s) is required", name, typ)
}

// Create Mission

type CreateMissionRequest struct {
	Name        string  `json: "name"`
	Description string  `json: "description"`
	Difficulty  string  `json: "difficulty"`
	Reward      float32 `json: "reward"`
}

func (r *CreateMissionRequest) Validate() error {

	if r.Name == "" && r.Description == "" && r.Difficulty == "" && r.Reward <= 0 {
		return fmt.Errorf("request body is empty or malformed")
	}

	if r.Name == "" {
		return errParamIsRequired("name", "string")
	}
	if r.Description == "" {
		return errParamIsRequired("description", "string")
	}
	if r.Difficulty == "" {
		return errParamIsRequired("difficulty", "string")
	}
	if r.Reward <= 0 {
		return errParamIsRequired("reward", "float32")
	}
	return nil
}
