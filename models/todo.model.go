package models

type Todo struct {
	TodoID          string `json:"id" db:"id"`
	TodoDescription string `json:"description" db:"description"`
	TodoPriority   string `json:"priority" db:"priority"`
	TodoUserId 	 string `json:"userId" db:"userId"`
}