// Code generated by sqlc. DO NOT EDIT.

package db

import (
	"fmt"
	"time"
)

type DevLevel string

const (
	DevLevelNew      DevLevel = "new"
	DevLevelJunior   DevLevel = "junior"
	DevLevelMidlevel DevLevel = "midlevel"
	DevLevelSenior   DevLevel = "senior"
)

func (e *DevLevel) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = DevLevel(s)
	case string:
		*e = DevLevel(s)
	default:
		return fmt.Errorf("unsupported scan type for DevLevel: %T", src)
	}
	return nil
}

type Account struct {
	ID        int64     `json:"id"`
	FullName  string    `json:"full_name"`
	Email     string    `json:"email"`
	OptedIn   bool      `json:"opted_in"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Category struct {
	ID        int64     `json:"id"`
	Category  string    `json:"category"`
	CreatedAt time.Time `json:"created_at"`
}

type Question struct {
	ID             int64     `json:"id"`
	ChallengeName  string    `json:"challenge_name"`
	Description    string    `json:"description"`
	Example        string    `json:"example"`
	Difficulty     int32     `json:"difficulty"`
	Complexity     string    `json:"complexity"`
	CompletionTime int32     `json:"completion_time"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type QuestionCategory struct {
	ID         int64 `json:"id"`
	QuestionID int64 `json:"question_id"`
	CategoryID int64 `json:"category_id"`
}

type User struct {
	ID        int64     `json:"id"`
	AccountID int64     `json:"account_id"`
	Username  string    `json:"username"`
	Grade     DevLevel  `json:"grade"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserQuestionScore struct {
	ID         int64     `json:"id"`
	UserID     int64     `json:"user_id"`
	QuestionID int64     `json:"question_id"`
	Score      int32     `json:"score"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
