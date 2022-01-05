package handlers

import (
	"context"
	"time"

	"example.com/tomdoestech/internal/db"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	ID        primitive.ObjectID `json:"_id" bson:"_id" validate:"required"`
	CreatedAt time.Time          `json:"createdAt" bson:"created_at" validate:"required"`
	UpdatedAt time.Time          `json:"updatedAt" bson:"updated_at" validate:"required"`
	Title     string             `json:"title" bson:"title" validate:"required,min=12"`
}

type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}

func ValidateProductStruct(p Product) []*ErrorResponse {
	var errors []*ErrorResponse
	validate := validator.New()
	err := validate.Struct(p)

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)

		}
	}
	return errors
}

func CreateProduct(c *fiber.Ctx) error {
	product := Product{
		ID:        primitive.NewObjectID(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := c.BodyParser(&product); err != nil {
		return err
	}

	// Could also be broken down to be
	// err := c.BodyParser(&product)

	// if err != nil {
	// 	return err
	// }

	errors := ValidateProductStruct(product)

	if errors != nil {
		return c.JSON(errors)
	}

	client, err := db.GetMongoClient()

	if err != nil {
		return err
	}

	collection := client.Database(db.Database).Collection(string(db.ProductsCollection))

	_, err = collection.InsertOne(context.TODO(), product)

	if err != nil {
		return err
	}

	return c.JSON(product)

}

func GetAllProducts(c *fiber.Ctx) error {
	client, err := db.GetMongoClient()

	var products []*Product

	if err != nil {
		return err
	}

	collection := client.Database(db.Database).Collection(string(db.ProductsCollection))

	cur, err := collection.Find(context.TODO(), bson.D{
		primitive.E{},
	})

	if err != nil {
		return err
	}

	for cur.Next(context.TODO()) {
		var p Product
		err := cur.Decode(&p)

		if err != nil {
			return err
		}

		products = append(products, &p)

	}

	return c.JSON(products)

}
