package constants

import "errors"

var (
	ErrEmptyTitle = errors.New("empty-title")
	ErrDBNotReady = errors.New("db-not-ready")
)

const (
	TASK_FETCH_LIMIT  = 10
	STEPS_FETCH_LIMIT = 7
)

const (
	StepKeyFmt     = "step:%s"
	IndexKeyFmt    = "task:%s:step:%020d_%s"
	IndexPrefixFmt = "task:%s:step:"
)

const (
	TaskKeyFmt         = "task:%s"
	IndexTaskKeyFmt    = "it:%020d:%s"
	IndexTaskPrefixFmt = "it:"
)

const (
	TaskCreated = "task-created"
	TaskEdited  = "task-edited"
	TaskDeleted = "task-deleted"

	StepAdded   = "step-added"
	StepEdited  = "step-edited"
	StepDeleted = "step-deleted"

	DBWiped = "db-wiped"
)
