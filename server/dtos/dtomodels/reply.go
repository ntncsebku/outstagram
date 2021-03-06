package dtomodels

import "time"

type Reply struct {
	ID            uint      `json:"id"`
	Content       *string   `json:"content"`
	OwnerID       uint      `json:"ownerID"`
	OwnerFullname string    `json:"ownerFullname"`
	OwnerUsername string    `json:"ownerUsername"`
	CreatedAt     time.Time `json:"createdAt"`
	Reactors      []string  `json:"reactors"`
	ReactCount    int       `json:"reactCount"`
	ReactableID   uint      `json:"reactableID"`
	Reacted       bool      `json:"reacted"`
}
