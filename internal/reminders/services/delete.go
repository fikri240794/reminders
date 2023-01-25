package services

import (
	"context"

	"github.com/fikri240794/golib/errors"
	base_enums "github.com/fikri240794/golib/models/enums"
	"github.com/fikri240794/reminders/internal/reminders/models/entities"
	"github.com/google/uuid"
)

func (i *ReminderService) Delete(ctx context.Context, id string) error {
	var (
		e         *entities.ReminderEntity
		errFields []errors.ErrorField
		err       error
	)

	errFields = []errors.ErrorField{}

	if id == "" {
		errFields = append(errFields, errors.NewErrorField("id", "required field"))
	}

	_, err = uuid.Parse(id)

	if err != nil {
		errFields = append(errFields, errors.NewErrorField("id", "value is not uuid"))
	}

	if len(errFields) > 0 {
		return errors.NewError(base_enums.BAD_REQUEST, base_enums.BAD_REQUEST.ToString(), errFields...)
	}

	e, err = i.reminderRepository.FindById(ctx, id)

	if err != nil {
		return err
	}

	err = i.reminderRepository.Delete(ctx, e)

	if err != nil {
		return err
	}

	return nil
}
