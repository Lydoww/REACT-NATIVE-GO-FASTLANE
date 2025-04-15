package models

import (
	"context"
	"time"
)


type Ticket struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	EventID   uint      `json:"eventId"`
	Event     Event     `json:"event" gorm:"foreignKey:EventID;references:ID;constraint:OnDelete:CASCADE,OnUpdate:CASCADE"`
	UserID    uint      `json:"userId"`
	User      User      `json:"user" gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE,OnUpdate:CASCADE"`
	Entered   bool      `json:"entered" gorm:"default:false"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}


type TicketRepository interface {
	GetMany(ctx context.Context, userId uint) ([]*Ticket, error)
	GetOne(ctx context.Context, userId uint, ticketId uint) (*Ticket, error)
	CreateOne(ctx context.Context, userId uint, ticket *Ticket) (*Ticket, error)
	UpdateOne(ctx context.Context,userId uint, ticketId uint, updateData map[string]interface{}) (*Ticket, error)
}

type ValidateTicket struct {
	TicketId uint `json:"ticketId"`
	OwnerId uint `json:"ownerId"`
}