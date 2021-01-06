package main

import (
	"context"
	"io"
	"log"
	"net/http"

	"cloud.google.com/go/firestore"
	"cloud.google.com/go/storage"
	firebase "firebase.google.com/go"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"
)

func uploadImage(c *gin.Context) {
	config := &firebase.Config{
		StorageBucket: "maintenance-7f16b.appspot.com",
	}

	file, handler, err := c.Request.FormFile("image")

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	defer file.Close()
	imagePath := handler.Filename
	// fmt.Println("imagePath: " + imagePah)
	opt := option.WithCredentialsFile("maintenance-7f16b-key.json")

	ctx := context.Background()
	app, err := firebase.NewApp(ctx, config, opt)
	if err != nil {
		log.Fatalln(err)
		c.JSON(http.StatusInternalServerError, err.Error())

	}
	client, err := app.Storage(ctx)
	//client1, err := firestore.NewClient(ctx, "34322657306")

	if err != nil {
		log.Fatalln(err)
		c.JSON(http.StatusInternalServerError, err.Error())

	}

	bucket, err := client.DefaultBucket()
	if err != nil {
		log.Fatalln(err)
		c.JSON(http.StatusInternalServerError, err.Error())

	}
	writer := bucket.Object(imagePath).NewWriter(ctx)
	writer.ObjectAttrs.CacheControl = "no-cache"
	writer.ObjectAttrs.ACL = []storage.ACLRule{
		{
			Entity: storage.AllUsers,
			Role:   storage.RoleReader,
		},
	}
	//createImageUrl(imagePath, config.StorageBucket, ctx, client1)
	if _, err = io.Copy(writer, file); err != nil {
		log.Fatalln(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	defer file.Close()

	if err := writer.Close(); err != nil {
		log.Fatalln(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.Header("Content-Type", "application/json; charset=utf-8")

	c.JSON(http.StatusCreated, "Create image success.")

}

func createImageUrl(imagePath string, bucket string, ctx context.Context, client *firestore.Client) error {
	imageStructure := ImageStructure{
		ImageName: imagePath,
		URL:       "https://storage.cloud.google.com/" + bucket + "/" + imagePath,
	}

	_, _, err := client.Collection("image").Add(ctx, imageStructure)
	if err != nil {
		return err
	}

	return nil
}
