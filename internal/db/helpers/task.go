package helpers

import (
	"encoding/json"
	"fmt"
	"tsdoa/internal/constants"
	"tsdoa/internal/db"
	"tsdoa/internal/models"
)

func StoreTask(task *models.Task) error {
	taskData, err := json.Marshal(task)
	if err != nil {
		return err
	}

	key := fmt.Sprintf("task:%s", task.Id)
	indexKey := fmt.Sprintf("it:%020d:%s", task.CreatedAt.UnixNano(), task.Id)

	err = db.DataDBManager.Set(key, taskData)
	if err != nil {
		return err
	}

	err = db.IndexDBManager.Set(indexKey, []byte(task.Id))
	if err != nil {
		return err
	}

	return nil
}

func GetTaskById(id string) (*models.Task, error) {
	key := fmt.Sprintf("task:%s", id)
	data, err := db.DataDBManager.GetByKey(key)
	if err != nil {
		return nil, err
	}

	var task models.Task
	err = json.Unmarshal(data, &task)
	if err != nil {
		return nil, err
	}

	return &task, nil
}

func GetPaginatedTasks(offset int) ([]*models.Task, bool, error) {
	keys, hasMore, err := db.IndexDBManager.GetPaginatedWithPrefix("it:", offset, constants.TASK_FETCH_LIMIT)
	if err != nil {
		return nil, false, err
	}

	var tasks []*models.Task
	for _, key := range keys {
		task, err := GetTaskById(string(key))
		if err != nil {
			return nil, false, err
		}
		tasks = append(tasks, task)
	}

	return tasks, hasMore, nil
}

func DeleteTaskById(id string) error {

	task, err := GetTaskById(id)
	if err != nil {
		return err
	}

	key := fmt.Sprintf("task:%s", id)
	if err := db.DataDBManager.DeleteByKey(key); err != nil {
		return err
	}

	indexKey := fmt.Sprintf("it:%020d:%s", task.CreatedAt.UnixNano(), id)
	if err := db.IndexDBManager.DeleteByKey(indexKey); err != nil {
		return err
	}

	if err := DeleteAllStepsByTaskId(id); err != nil {
		return err
	}

	return nil
}

func WipeOutAllDBs() error {
	err := db.DataDBManager.WipeOut()
	if err != nil {
		return err
	}

	err = db.IndexDBManager.WipeOut()
	if err != nil {
		return err
	}

	return nil
}
