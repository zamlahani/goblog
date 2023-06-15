package controllers

import (
	"context"
	// "fmt"
	"goblog/configs"
	"goblog/models"
	"goblog/responses"
	"net/http"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var postCollection *mongo.Collection = configs.GetCollection(configs.DB, "posts")
var validate = validator.New()

func CreatePost(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var post models.Post
	defer cancel()

	//validate the request body
	if err := c.BodyParser(&post); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.PostResponse{Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	//use the validator library to validate required fields
	if validationErr := validate.Struct(&post); validationErr != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.PostResponse{Message: "error", Data: &fiber.Map{"data": validationErr.Error()}})
	}

	newPost := models.Post{
		Title: post.Title,
		Body:  post.Body,
	}

	result, err := postCollection.InsertOne(ctx, newPost)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.PostResponse{Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	return c.Status(http.StatusCreated).JSON(responses.PostResponse{Message: "success", Data: &fiber.Map{"data": result}})
}

func GetAllPosts(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var posts []models.Post
	defer cancel()

	results, err := postCollection.Find(ctx, bson.M{})

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.PostResponse{Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	//reading from the db in an optimal way
	defer results.Close(ctx)
	for results.Next(ctx) {
		var singlePost models.Post
		err = results.Decode(&singlePost)
		if err != nil {
			var res responses.PostResponse
			res.Message = "error"
			res.Data = &fiber.Map{"data": err.Error()}
			return c.Status(http.StatusInternalServerError).JSON(res)
		}

		posts = append(posts, singlePost)
	}
	var res responses.PostResponse
	res.Message = "success"
	res.Data = &fiber.Map{"data": posts}

	return c.Status(http.StatusOK).JSON(res)
}

func Fib(c *fiber.Ctx) error {
	// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// defer cancel()
	n := c.Query("n")
	nNum, err := strconv.ParseUint(n, 10, 64)
	// fmt.Println(nNum, n)
	if n == "" || err != nil {
		// fmt.Println(err)
		return c.Status(http.StatusBadRequest).JSON(
			responses.PostResponse{
				Message: "failed",
				Data: &fiber.Map{
					"message": "wrong input n",
					"n":       n,
				},
			},
		)
	}
	return c.Status(http.StatusOK).JSON(
		responses.PostResponse{Message: "success", Data: &fiber.Map{"result": fibo(nNum)}})

}

func fibo(n uint64) uint64 {
	if n == 0 || n == 1 {
		return 1
	} else {
		return (fibo(n-2) + fibo(n-1))
	}
}
