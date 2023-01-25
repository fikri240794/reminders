package entities

import "time"

type ReminderEntity struct {
	ID        string
	Note      string
	CreatedAt int64
	UpdatedAt int64
	IsDeleted bool
}

func NewReminderEntity() *ReminderEntity {
	return &ReminderEntity{}
}

func (e *ReminderEntity) MarkAsDeleted() *ReminderEntity {
	e.UpdatedAt = time.Now().UnixMilli()
	e.IsDeleted = true

	return e
}
