package service

import (
	"errors"
	"mood-api/internal/models"
	"mood-api/internal/repository"
	"time"
)

type MoodService struct {
	MoodRepo *repository.MoodRepository
}

func (s *MoodService) AddMoodEntry(entry *models.MoodEntry) error {
	entry.Timestamp = time.Now()
	return s.MoodRepo.SaveMoodEntry(entry)
}

func (s *MoodService) GenerateReport(userID string) (*models.MoodReport, error) {
	entries, err := s.MoodRepo.GetMoodEntriesByUserAndWeek(userID)
	if err != nil {
		return nil, err
	}
	if len(entries) == 0 {
		return nil, errors.New("sem entradas de humor esta semana")
	}
	var sum int
	for _, e := range entries {
		sum += e.Level
	}
	avg := float64(sum) / float64(len(entries))

	return &models.MoodReport{
		UserID:     userID,
		Average:    avg,
		From:       time.Now().AddDate(0, 0, -7),
		To:         time.Now(),
		EntryCount: len(entries),
	}, nil
}
