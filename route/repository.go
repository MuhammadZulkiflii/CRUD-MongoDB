package route

import (
	"context"

	"github.com/MuhammadZulkiflii/CRUD-MongoDB/entity"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type repository interface {
	Create(ctx context.Context, nama string) (*entity.Data, error)
	Update(ctx context.Context, id primitive.ObjectID, nama string) (*entity.Data, error)
	Delete(ctx context.Context, id primitive.ObjectID) (*entity.Data, error)
	Find(ctx context.Context) ([]*entity.Data, error)
}
