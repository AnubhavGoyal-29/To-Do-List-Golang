package models

import "time"

type Task struct {
	ID        uint      `gorm:"primaryKey;autoIncrement"`
	Title     string    `gorm:"size:255;not null"`
	Status    int       `gorm:"not null;default:1"`  // Default to "Pending"
	Priority  int       `gorm:"not null;default:1"`  // Default to "Low" (1)
	Type      int       `gorm:"not null;default:1"`  // Default to "General" (1)
	UserID    uint      `gorm:"not null;constraint:OnDelete:CASCADE;foreignKey:UserID;references:ID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// TaskResponse (used for API responses)
type TaskResponse struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	Status    string    `json:"status"`
	Priority  string    `json:"priority"`
	Type      string    `json:"type"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Convert Task model to TaskResponse
func (t *Task) Serialize() TaskResponse {
	return TaskResponse{
		ID:        t.ID,
		Title:     t.Title,
		Status:    getStatusName(t.Status),
		Priority:  getPriorityName(t.Priority),
		Type:      getTypeName(t.Type),
		CreatedAt: t.CreatedAt,
		UpdatedAt: t.UpdatedAt,
	}
}

// Function to map status codes to names
func getStatusName(status int) string {
	switch status {
	case 1:
		return "PENDING"
	case 2:
		return "WORKING"
	case 3:
		return "COMPLETED"
	default:
		return "UNKNOWN"
	}
}

func getPriorityName(priority int) string {
  switch priority {
  case 1:
    return "LOW"
  case 2:
    return "MEDIUM"
  case 3:
    return "HIGH"
  default:
    return "UNKNOWN"
  }
}

func getTypeName(_type int) string {
  switch _type {
  case 1:
    return "MEETING"
  case 2:
    return "SHOPPING"
  case 3:
    return "WORK"
  default:
    return "UNKNOWN"
  }
}
