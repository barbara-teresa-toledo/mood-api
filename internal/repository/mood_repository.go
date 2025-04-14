package repository

import (
	"database/sql"
	"mood-api/internal/models"
	"time"
)

type MoodRepository struct {
	DB *sql.DB
}

func (r *MoodRepository) SaveMoodEntry(entry *models.MoodEntry) error {
	_, err := r.DB.Exec("INSERT INTO mood_entries (user_id, level, timestamp) VALUES ($1, $2, $3)", entry.UserID, entry.Level, entry.Timestamp)
	return err
}

func (r *MoodRepository) GetMoodEntriesByUserAndWeek(userID string) ([]models.MoodEntry, error) {
	weekAgo := time.Now().AddDate(0, 0, -7)
	rows, err := r.DB.Query("SELECT id, user_id, level, timestamp FROM mood_entries WHERE user_id = $1 AND timestamp >= $2", userID, weekAgo)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var entries []models.MoodEntry
	for rows.Next() {
		var entry models.MoodEntry
		if err := rows.Scan(&entry.ID, &entry.UserID, &entry.Level, &entry.Timestamp); err != nil {
			return nil, err
		}
		entries = append(entries, entry)
	}
	return entries, nil
}
