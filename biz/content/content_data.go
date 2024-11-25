package content

import (
	"context"
	"time"

	"github.com/dgdts/UniversalServer/pkg/mongo"
)

type User struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Avatar    string    `json:"avatar"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	IsDeleted bool      `json:"is_deleted"`
}

type ContentData struct {
	ID        string    `json:"id" bson:"_id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	IsDeleted bool      `json:"is_deleted"`
}

func InsertContent(ctx context.Context, content *ContentData) error {
	r := mongo.Inserter(ContentCollection(ctx)).Insert(ctx, content)
	return r.Error()
}

func ContentByID(ctx context.Context, contentID string) (*ContentData, error) {
	content := &ContentData{}
	err := mongo.Finder(ContentCollection(ctx)).FindOne(ctx, "_id", contentID).Read(content)
	return content, err
}