package services

import (
	"context"
	"fmt"
	"time"
	"tsdoa/internal/constants"
	"tsdoa/internal/db/helpers"
	"tsdoa/internal/models"
	"tsdoa/internal/utils"

	"github.com/google/uuid"
)

type TaskService struct {
	ctx context.Context
}

func NewTaskService() *TaskService {
	return &TaskService{}
}

func (ts *TaskService) SetContext(ctx context.Context) {
	ts.ctx = ctx
}

func (ts *TaskService) CreateTask(title string) error {
	if title == "" {
		return constants.ErrEmptyTitle
	}

	task := &models.Task{
		Id:        uuid.New().String(),
		Title:     title,
		CreatedAt: time.Now(),
	}

	// Store in database
	err := helpers.StoreTask(task)
	if err != nil {
		return err
	}

	// Create a deafault step for the task
	defaultStep := &models.Step{
		Id:        uuid.New().String(),
		Title:     "I won't allow you to create a task without a step 😈, ( But you can delete it anyways 🤪 )",
		Done:      false,
		CreatedAt: time.Now(),
	}
	err = helpers.StoreStep(task.Id, defaultStep)
	if err != nil {
		return err
	}

	// return the task response
	stepRes := utils.NewStepResponse(defaultStep)
	res := utils.NewTaskResponse(task, []*models.StepResponse{stepRes}, false)
	utils.EmitEvent(ts.ctx, constants.TaskCreated, res)
	return nil
}

func (ts *TaskService) GetTasks(offset int) (*models.TaskTreeResponse, error) {
	tasks, hasMore, err := helpers.GetPaginatedTasks(offset)
	if err != nil {
		return nil, err
	}
	var taskResponses = &models.TaskTreeResponse{
		Tasks:        make([]*models.TaskResponse, 0, len(tasks)),
		HasMoreTasks: hasMore,
	}

	for _, task := range tasks {
		// Get steps for each task
		steps, hasMoreSteps, err := helpers.GetPaginatedStepsByTaskId(task.Id, 0, constants.STEPS_FETCH_LIMIT)
		if err != nil {
			return nil, err
		}

		// Convert steps to StepResponse
		stepResponses := make([]*models.StepResponse, 0, len(steps))
		for _, step := range steps {
			stepResponses = append(stepResponses, utils.NewStepResponse(step))
		}

		taskResponse := utils.NewTaskResponse(task, stepResponses, hasMoreSteps)
		taskResponses.Tasks = append(taskResponses.Tasks, taskResponse)
	}

	return taskResponses, nil
}

func (ts *TaskService) EditTask(id, title string) error {
	if title == "" {
		return constants.ErrEmptyTitle
	}

	task, err := helpers.GetTaskById(id)
	if err != nil {
		return err
	}

	task.Title = title
	err = helpers.StoreTask(task)
	if err != nil {
		return err
	}

	res := map[string]string{
		"taskId": task.Id,
		"title":  task.Title,
	}
	utils.EmitEvent(ts.ctx, constants.TaskEdited, res)
	return nil
}

func (ts *TaskService) DeleteTask(id string) error {
	err := helpers.DeleteTaskById(id)
	if err != nil {
		return err
	}
	utils.EmitEvent(ts.ctx, constants.TaskDeleted, map[string]string{
		"taskId": id,
	})
	return nil
}

func (ts *TaskService) DeleteAllTasks() error {
	// Delete all tasks from the database
	err := helpers.WipeOutAllDBs()
	if err != nil {
		return err
	}

	utils.EmitEmptyEvent(ts.ctx, constants.DBWiped)

	return nil
}

func (ts *TaskService) AddStepToTask(taskId, title string) error {
	if title == "" {
		return constants.ErrEmptyTitle
	}

	task, err := helpers.GetTaskById(taskId)
	if err != nil {
		return err
	}

	step := &models.Step{
		Id:        uuid.New().String(),
		Title:     title,
		Done:      false,
		CreatedAt: time.Now(),
	}

	err = helpers.StoreStep(task.Id, step)
	if err != nil {
		return err
	}

	res := utils.NewStepResponse(step)
	utils.EmitEvent(ts.ctx, constants.StepAdded, map[string]any{
		"taskId": task.Id,
		"step":   res,
	})

	return nil
}

func (ts *TaskService) GetStepsByTaskId(taskId string, offset int) (*models.StepTreeResponse, error) {
	steps, hasMore, err := helpers.GetPaginatedStepsByTaskId(taskId, offset, constants.STEPS_FETCH_LIMIT)
	if err != nil {
		return nil, err
	}

	stepTreeResponse := &models.StepTreeResponse{
		TaskId:        taskId,
		Steps:         make([]*models.StepResponse, 0, len(steps)),
		MoreAvailable: hasMore,
	}

	for _, step := range steps {
		fmt.Println("Step ID:", step.Id)
		stepTreeResponse.Steps = append(stepTreeResponse.Steps, utils.NewStepResponse(step))
	}

	return stepTreeResponse, nil

}

func (ts *TaskService) EditStep(taskId, stepId, title string) error {
	if title == "" {
		return constants.ErrEmptyTitle
	}

	step, err := helpers.GetStepById(stepId)
	if err != nil {
		return err
	}

	step.Title = title
	err = helpers.StoreStep(taskId, step)
	if err != nil {
		return err
	}

	res := utils.NewStepResponse(step)
	utils.EmitEvent(ts.ctx, constants.StepEdited, map[string]any{
		"taskId": taskId,
		"step":   res,
	})

	return nil
}

func (ts *TaskService) ToggleStep(taskId, stepId string) error {
	step, err := helpers.GetStepById(stepId)
	if err != nil {
		return err
	}

	step.Done = !step.Done
	err = helpers.StoreStep(taskId, step)
	if err != nil {
		return err
	}

	res := utils.NewStepResponse(step)
	utils.EmitEvent(ts.ctx, constants.StepEdited, map[string]any{
		"taskId": taskId,
		"step":   res,
	})

	return nil
}

func (ts *TaskService) DeleteStep(taskId, stepId string) error {
	err := helpers.DeleteStepById(taskId, stepId)
	if err != nil {
		return err
	}

	utils.EmitEvent(ts.ctx, constants.StepDeleted, map[string]string{
		"taskId": taskId,
		"stepId": stepId,
	})

	return nil
}
