package models

import "time"

type Task struct {
	Id    string
	Title string

	CreatedAt time.Time
}

type Step struct {
	Id    string
	Title string
	Done  bool

	CreatedAt time.Time
}

type StepResponse struct {
	Id    string `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

type TaskResponse struct {
	Id           string          `json:"id"`
	Title        string          `json:"title"`
	Steps        []*StepResponse `json:"steps"`
	HasMoreSteps bool            `json:"has_more_steps"`
}

type TaskTreeResponse struct {
	Tasks        []*TaskResponse `json:"tasks"`
	HasMoreTasks bool            `json:"has_more_tasks"`
}

type StepTreeResponse struct {
	TaskId        string          `json:"task_id"`
	Steps         []*StepResponse `json:"steps"`
	MoreAvailable bool            `json:"more_available"`
}
