package common

import "time"

func NewMate() *Mate {
	return &Mate{
		CreatedAt: time.Now().Unix(),
	}
}

type Mate struct {
	Id        int   `json:"id" gorm:"column:id"`
	CreatedAt int64 `json:"created_at" gorm:"created_at"`
	UpdatedAt int64 `json:"updated_at" gorm:"updated_at"`
}
