package repositories

import (
	"context"
	"time"

	"github.com/Lydoww/react-native-go-fastlane/models"
)

type EventRepository struct {
	db any 
}

func (r *EventRepository) GetMany(ctx context.Context) ([]*models.Event, error) {
	events := []*models.Event{}

	events = append(events, &models.Event{
		ID:"0250505262612512",
		Name: "favorite band",
		Location: "favorite place",
		Date: time.Now(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
})
return events, nil
}

func (r *EventRepository) GetOne(ctx context.Context, eventId string) (*models.Event, error) {
	return nil, nil
}

func (r *EventRepository) CreateOne(ctx context.Context, event models.Event) (*models.Event, error) {
	return nil, nil
}


func NewEventRepository(db any) models.EventRepository {
	return &EventRepository{
		db: db,
		
	}
}