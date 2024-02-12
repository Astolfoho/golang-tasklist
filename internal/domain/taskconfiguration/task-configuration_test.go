package taskconfiguration

import (
	"testing"
	"time"

	"github.com/jaswdr/faker/v2"
	"github.com/stretchr/testify/assert"
)

const (
	name        = "task_1"
	description = "task_1 description"
	daysOfWeek  = Monday | Thursday
)

var fake = faker.New()

func Test_TaskConfig_Create(t *testing.T) {
	assert := assert.New(t)

	var tc, _ = NewTaskConfig(name, description, daysOfWeek)

	assert.Equal(tc.Name, name)
	assert.Equal(tc.Description, description)
	assert.Equal(tc.DaysOfWeek, daysOfWeek)
}

func Test_TaskConfig_IdIsNotNill(t *testing.T) {
	assert := assert.New(t)

	var tc, _ = NewTaskConfig(name, description, daysOfWeek)

	assert.NotNil(tc.Id)
}

func Test_TaskConfig_CreatedOnMustBeNow(t *testing.T) {
	assert := assert.New(t)

	createdOn := time.Now().Add(-time.Minute)

	var tc, _ = NewTaskConfig(name, description, daysOfWeek)
	assert.Greater(tc.CreatedOn, createdOn)
}

func Test_TaskConfig_MustValidateNameMin(t *testing.T) {
	assert := assert.New(t)
	_, err := NewTaskConfig("", description, daysOfWeek)
	assert.Equal(err.StructField(), "Name")
	assert.Equal(err.Tag(), "min")
}

func Test_TaskConfig_MustValidateNameMax(t *testing.T) {
	assert := assert.New(t)
	_, err := NewTaskConfig(fake.Lorem().Text(200), description, daysOfWeek)
	assert.Equal(err.StructField(), "Name")
	assert.Equal(err.Tag(), "max")
}

func Test_TaskConfig_MustValidateDescriptionMin(t *testing.T) {
	assert := assert.New(t)
	_, err := NewTaskConfig(name, "", daysOfWeek)
	assert.Equal(err.StructField(), "Description")
	assert.Equal(err.Tag(), "min")
}

func Test_TaskConfig_MustValidateDescriptionMax(t *testing.T) {
	assert := assert.New(t)
	_, err := NewTaskConfig(name, fake.Lorem().Text(1000), daysOfWeek)
	assert.Equal(err.StructField(), "Description")
	assert.Equal(err.Tag(), "max")
}

func Test_TaskConfig_MustValidateDaysOfWeek(t *testing.T) {
	assert := assert.New(t)
	_, err := NewTaskConfig(name, description, 0)
	assert.Equal(err.StructField(), "DaysOfWeek")
	assert.Equal(err.Tag(), "required")
}
