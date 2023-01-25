package repositories

import (
	"context"

	"github.com/fikri240794/golib/errors"
	base_enums "github.com/fikri240794/golib/models/enums"
	"github.com/fikri240794/reminders/internal/reminders/models/dtos"
	"github.com/fikri240794/reminders/internal/reminders/models/entities"
)

func (r *ReminderRepository) findAllFromPersistence(ctx context.Context, dto *dtos.ReminderPaginationParamsDTO) ([]*entities.ReminderEntity, error) {
	var (
		entityDatas []*entities.ReminderEntity
		err         error
	)

	entityDatas, err = r.reminderPersistence.FindAllReminder(ctx, dto)

	if err != nil {
		return nil, err
	}

	return entityDatas, nil
}

func (r *ReminderRepository) FindAll(ctx context.Context, dto *dtos.ReminderPaginationParamsDTO) ([]*entities.ReminderEntity, error) {
	var (
		entityDatas []*entities.ReminderEntity
		err         error
	)

	if dto == nil {
		return nil, errors.NewError(base_enums.BAD_REQUEST, "reminder pagination params dto is null")
	}

	entityDatas, err = r.findAllFromPersistence(ctx, dto)

	if err != nil {
		return nil, err
	}

	return entityDatas, nil
}
