package utils

import "tsdoa/internal/models"

func NewTaskResponse(task *models.Task, steps []*models.StepResponse, hasMore bool) *models.TaskResponse {
	return &models.TaskResponse{
		Id:           task.Id,
		Title:        task.Title,
		HasMoreSteps: hasMore,
		Steps:        steps,
	}
}

func NewStepResponse(step *models.Step) *models.StepResponse {
	return &models.StepResponse{
		Id:    step.Id,
		Title: step.Title,
		Done:  step.Done,
	}
}
