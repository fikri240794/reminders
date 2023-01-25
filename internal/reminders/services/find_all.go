package services

import (
	"context"

	"github.com/fikri240794/reminders/internal/reminders/models/dtos"
	"github.com/fikri240794/reminders/internal/reminders/models/entities"
	"golang.org/x/sync/errgroup"
)

func (i *ReminderService) FindAll(ctx context.Context, dto *dtos.ReminderPaginationParamsDTO) (*dtos.ReminderPaginationResultDTO, error) {
	var (
		entityDatas []*entities.ReminderEntity
		count       int64
		errGroup    *errgroup.Group
		errGroupCtx context.Context
		result      *dtos.ReminderPaginationResultDTO
		err         error
	)

	err = dto.Validate()

	if err != nil {
		return nil, err
	}

	errGroup, errGroupCtx = errgroup.WithContext(ctx)
	errGroup.Go(func() error {
		var errRoutine error

		count, errRoutine = i.reminderRepository.CountAll(errGroupCtx, dto)

		if errRoutine != nil {
			return errRoutine
		}

		return nil
	})
	errGroup.Go(func() error {
		var errRoutine error

		entityDatas, errRoutine = i.reminderRepository.FindAll(errGroupCtx, dto)

		if errRoutine != nil {
			return errRoutine
		}

		return nil
	})
	err = errGroup.Wait()

	if err != nil {
		return nil, err
	}

	result = dtos.NewReminderPaginationResultDTO().
		SetCollection(entityDatas).
		SetCount(count)

	return result, nil
}
