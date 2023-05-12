package repository

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/MuhammadZulkiflii/CRUD-MongoDB/entity"
	"github.com/MuhammadZulkiflii/CRUD-MongoDB/errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository struct {
	db *mongo.Collection
}

func NewRepository() (*mongo.Client, *Repository) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("DB_URI")))
	if err != nil {
		log.Fatalln(err)
	}
	return client, &Repository{
		db: client.Database(os.Getenv("DB_NAME")).Collection(os.Getenv("COL_NAME")),
	}
}

func toError(err error) error {
	switch {
	case err == nil:
		return nil
	case errors.IsError(err):
		return err
	case mongo.IsDuplicateKeyError(err):
		return errors.Conflict(err)
	case err == mongo.ErrNoDocuments:
		return errors.NotFound(err)
	default:
		return errors.Internal(err)
	}
}

func (repo *Repository) Find(ctx context.Context) ([]*entity.Data, error) {
	cur, err := repo.db.Find(ctx, bson.M{})
	if err != nil {
		return nil, toError(err)
	}
	items := make([]*entity.Data, 0)
	if err := cur.All(ctx, &items); err != nil {
		return nil, toError(err)
	}
	return items, nil
}

func (repo *Repository) Delete(ctx context.Context, id primitive.ObjectID) (*entity.Data, error) {
	var result entity.Data
	if err := repo.db.FindOneAndDelete(
		ctx,
		bson.M{"_id": bson.M{"$eq": id}},
	).Decode(&result); err != nil {
		return nil, toError(err)
	}
	return &result, nil
}

func (repo *Repository) Create(ctx context.Context, nama string) (*entity.Data, error) {
	result := entity.Create(nama)
	if _, err := repo.db.InsertOne(ctx, result); err != nil {
		return nil, toError(err)
	}
	return result, nil
}

func (repo *Repository) Update(ctx context.Context, id primitive.ObjectID, nama string) (*entity.Data, error) {
	var data entity.Data
	err := repo.db.FindOneAndUpdate(ctx, bson.M{"_id": bson.M{"$eq": id}}, bson.M{"$set": bson.M{"nama": nama}}).Decode(&data)
	if err != nil {
		return nil, toError(err)
	}
	return &data, nil
}
