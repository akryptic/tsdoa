package helpers

import (
	"encoding/json"
	"fmt"
	"tsdoa/internal/constants"
	"tsdoa/internal/db"
	"tsdoa/internal/models"
)

func StoreStep(taskId string, step *models.Step) error {
	stepData, err := json.Marshal(step)
	if err != nil {
		return err
	}

	stepKey := fmt.Sprintf(constants.StepKeyFmt, step.Id)
	indexKey := fmt.Sprintf(constants.IndexKeyFmt, taskId, step.CreatedAt.UnixNano(), step.Id)

	// Store step data in main data DB
	if err := db.DataDBManager.Set(stepKey, stepData); err != nil {
		return err
	}

	// Store relationship in index DB
	if err := db.IndexDBManager.Set(indexKey, []byte(step.Id)); err != nil {
		return err
	}

	return nil
}

func GetStepById(id string) (*models.Step, error) {
	stepKey := fmt.Sprintf(constants.StepKeyFmt, id)
	data, err := db.DataDBManager.GetByKey(stepKey)
	if err != nil {
		return nil, err
	}

	var step models.Step
	err = json.Unmarshal(data, &step)
	if err != nil {
		return nil, err
	}

	return &step, nil
}

func GetPaginatedStepsByTaskId(taskId string, offset int, limit int) ([]*models.Step, bool, error) {
	indexKey := fmt.Sprintf(constants.IndexPrefixFmt, taskId)
	stepsData, hasMore, err := db.IndexDBManager.GetPaginatedWithPrefix(indexKey, offset, limit)
	if err != nil {
		return nil, false, err
	}

	var steps []*models.Step
	for _, data := range stepsData {
		stepId := string(data)
		step, err := GetStepById(stepId)
		if err != nil {
			return nil, false, err
		}
		steps = append(steps, step)
	}

	return steps, hasMore, nil
}

func DeleteStepById(taskId string, id string) error {
	stepKey := fmt.Sprintf(constants.StepKeyFmt, id)

	step, err := GetStepById(id)
	if err != nil {
		return err
	}

	indexKey := fmt.Sprintf(constants.IndexKeyFmt, taskId, step.CreatedAt.UnixNano(), id)

	// Delete step data from main data DB
	if err := db.DataDBManager.DeleteByKey(stepKey); err != nil {
		return err
	}

	// Delete relationship from index DB
	if err := db.IndexDBManager.DeleteByKey(indexKey); err != nil {
		return err
	}

	return nil
}

func DeleteAllStepsByTaskId(taskId string) error {
	steps, _, err := GetPaginatedStepsByTaskId(taskId, 0, -1)
	if err != nil {
		return err
	}

	var failed []string

	for _, step := range steps {
		if err := DeleteStepById(taskId, step.Id); err != nil {
			failed = append(failed, step.Id)
		}
	}

	if len(failed) > 0 {
		return fmt.Errorf("failed to delete steps: %v", failed)
	}

	return nil
}
