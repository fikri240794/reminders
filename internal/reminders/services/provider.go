package services

import "github.com/fikri240794/reminders/internal/reminders/repositories"

type ReminderService struct {
	reminderRepository repositories.IReminderRepository
}

func NewReminderService(reminderRepository repositories.IReminderRepository) *ReminderService {
	return &ReminderService{
		reminderRepository: reminderRepository,
	}
}
