package route

import (
	"net/http"
	"os"

	"github.com/MuhammadZulkiflii/CRUD-MongoDB/entity"
	"github.com/MuhammadZulkiflii/CRUD-MongoDB/errors"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func NewRoute(repo repository) *http.Server {
	route := NewRoutes()
	url := route.Group("/test/v1")
	url.GET("", find(repo))
	url.POST("", create(repo))
	url.PUT("/:id", update(repo))
	url.DELETE("/:id", delete(repo))
	srv := &http.Server{
		Addr:    ":" + os.Getenv("PORT"),
		Handler: route,
	}
	return srv
}

func paramId(ctx *gin.Context) (primitive.ObjectID, error) {
	id, err := primitive.ObjectIDFromHex(ctx.Param("id"))
	if err != nil {
		return id, err
	}
	return id, nil
}

func find(repo repository) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		result, err := repo.Find(ctx)
		Response(ctx, result, err)
	}
}

func create(repo repository) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var data entity.Data
		if err := ctx.ShouldBind(&data); err != nil {
			Response(ctx, nil, errors.BadRequest(err))
			return
		}
		result, err := repo.Create(ctx, data.Nama)
		Response(ctx, result, err)
	}
}
func update(repo repository) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := paramId(ctx)
		if err != nil {
			Response(ctx, nil, errors.BadRequest(err))
			return
		}
		var data entity.Data
		if err := ctx.ShouldBind(&data); err != nil {
			Response(ctx, nil, errors.BadRequest(err))
			return
		}
		result, err := repo.Update(ctx, id, data.Nama)
		Response(ctx, result, err)
	}
}

func delete(repo repository) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := paramId(ctx)
		if err != nil {
			Response(ctx, nil, errors.BadRequest(err))
			return
		}
		result, err := repo.Delete(ctx, id)
		Response(ctx, result, err)
	}
}
