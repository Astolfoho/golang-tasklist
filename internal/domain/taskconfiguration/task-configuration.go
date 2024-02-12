package taskconfiguration

import (
	internalerrors "task-list/internal/domain/internal-errors"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/rs/xid"
)

type TaskConfiguration struct {
	Id          string    `validate:"required" gorm:"size:50"`
	Name        string    `validate:"min=3,max=50" gorm:"size:50"`
	Description string    `validate:"min=10,max=500" gorm:"size:500"`
	DaysOfWeek  Weekdays  `validate:"required"`
	CreatedOn   time.Time `validate:"required"`
}

type Weekdays int

const (
	Sunday Weekdays = 1 << iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func NewTaskConfig(name string, description string, daysOfWeek Weekdays) (*TaskConfiguration, validator.FieldError) {
	ret := &TaskConfiguration{
		Id:          xid.New().String(),
		Name:        name,
		Description: description,
		DaysOfWeek:  daysOfWeek,
		CreatedOn:   time.Now(),
	}

	println(ret.Description)
	err := internalerrors.ValidateStruct(ret)
	if err != nil {
		return nil, err
	}

	return ret, nil

}
