package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"maintenaceApi/databaseManage"
	"maintenaceApi/unit"
	"net/http"
	"os"

	"cloud.google.com/go/firestore"
	"cloud.google.com/go/storage"
	firebase "firebase.google.com/go"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"
)

type BasicResponseStringMessage struct {
	Error   int    `json:"error"`
	Message string `json:"message"`
}

type BasicResponse struct {
	Error   int             `json:"error"`
	Message json.RawMessage `json:"message"`
}

type Masterdata []struct {
	CreatedDate string `json:"created_date"`
	FullName    string `json:"full_name"`
	ID          string `json:"id"`
}

type ImageStructure struct {
	ImageName string
	URL       string
}

type Detaildata []struct {
	ID                  string `json:"id"`
	CreatedDate         string `json:"created_date"`
	AircraftType        string `json:"aircraft_type"`
	AircraftSN          string `json:"aircraft_sn"`
	System              string `json:"system"`
	PrimaryPilot        string `json:"primary_pilot"`
	SecondaryPilot      string `json:"secondary_pilot"`
	Recorder            string `json:"recorder"`
	Trouble             string `json:"trouble"`
	TechnicalOrder      string `json:"technical_order"`
	TroubleShooting     string `json:"trouble_shooting"`
	Replace             string `json:"replace"`
	Name                string `json:"name"`
	PartNumber          string `json:"part_number"`
	SerailNumberRemove  string `json:"serail_number_remove"`
	SerailNumberinstall string `json:"serail_number_install"`
	Remark              string `json:"remark"`
}

func main() {

	router := gin.Default()

	router.GET("/", homePage)
	router.GET("getMasterSystem", getMasterSystem)
	router.GET("getMasterAircraft", getMasterAircraft)
	router.GET("getMasterTechnicalOrder", getMasterTechnicalOrder)
	router.POST("uploadImage", uploadImage)
	log.Fatal(
		// start on port
		http.ListenAndServe(getPort(), router),
	)

	router.Run()
}

func getIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello world!")
}

func homePage(c *gin.Context) {
	c.Header("Content-Type", "application/json; charset=utf-8")

	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})

}

func getMasterSystem(c *gin.Context) {
	var masters Masterdata
	response, err := databaseManage.SelectDataReturnJsonFormat("SELECT * FROM " + unit.MASTER_SYSTEM_TABLE)
	if err != nil {
		log.Fatalln(err)
	} else {
		js, err := convertStringToJsonFormat(response, masters)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return
		}
		c.Header("Content-Type", "application/json; charset=utf-8")
		dataRes := &BasicResponse{0, js}
		c.JSON(http.StatusOK, dataRes)

	}

}

func getMasterAircraft(c *gin.Context) {

	var masters Masterdata
	response, err := databaseManage.SelectDataReturnJsonFormat("SELECT * FROM t016ffukzsi0y5ie.master_aircraft")
	if err != nil {
		log.Fatalln(err)
	} else {
		js, err := convertStringToJsonFormat(response, masters)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return
		}
		c.Header("Content-Type", "application/json; charset=utf-8")
		dataRes := &BasicResponse{0, js}
		c.JSON(http.StatusOK, dataRes)
	}

}

func getMasterTechnicalOrder(c *gin.Context) {

	var masters Masterdata
	response, err := databaseManage.SelectDataReturnJsonFormat("SELECT * FROM t016ffukzsi0y5ie.master_technical_order")
	if err != nil {
		log.Fatalln(err)
	} else {
		js, err := convertStringToJsonFormat(response, masters)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return
		}
		c.Header("Content-Type", "application/json; charset=utf-8")
		dataRes := &BasicResponse{0, js}
		c.JSON(http.StatusOK, dataRes)
	}

}

func insertDetail(c *gin.Context) {
	var detailData Detaildata
	body, err := ioutil.ReadAll(io.LimitReader(c.Request.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := c.Request.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &detailData); err != nil {

		c.Status(http.StatusInternalServerError)
		c.JSON(http.StatusInternalServerError, json.NewEncoder(c.Writer).Encode(err))
		panic(err)

	}
	db, err := databaseManage.ConnectDB()
	stmt, err := db.Prepare("insert into maintanace_detail values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?) ")
	if err != nil {
		panic(err)
	}
	_, err = stmt.Exec("", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "")
	if err != nil {
		panic(err.Error())
	}
	c.Header("Content-Type", "application/json; charset=utf-8")

	c.JSON(0, json.NewEncoder(c.Writer).Encode(BasicResponseStringMessage{
		0,
		"success",
	}))

}

func getDetail(c *gin.Context) {
	var dataDetail Detaildata

	response, err := databaseManage.SelectDataReturnJsonFormat("SELECT * FROM t016ffukzsi0y5ie.maintenace_detail")
	if err != nil {
		log.Fatalln(err)
	} else {
		js, err := convertStringToDetailJsonFormat(response, dataDetail)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			c.JSON(http.StatusInternalServerError, json.NewEncoder(c.Writer).Encode(err))
			panic(err)
		}
		{

			c.Header("Content-Type", "application/json; charset=utf-8")
			unit.SetResponseSuccess(c, js)
			dataRes := &BasicResponse{0, js}
			c.JSON(http.StatusOK, dataRes)

		}
	}
}

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

func getPort() string {
	var port = os.Getenv("PORT")
	if port == "" {
		port = "8080"
		fmt.Println("No Port In Heroku " + port)
	}
	return ":" + port
}

func convertStringToJsonFormat(message string, format Masterdata) ([]byte, error) {
	errpare := json.Unmarshal([]byte(message), &format)
	if errpare != nil {
		return nil, errpare
	} else {
		json, err := json.Marshal(format)
		if err != nil {
			return nil, err
		}
		return json, nil
	}
}

func convertStringToDetailJsonFormat(message string, format Detaildata) ([]byte, error) {
	errpare := json.Unmarshal([]byte(message), &format)
	if errpare != nil {
		return nil, errpare
	} else {
		json, err := json.Marshal(format)
		if err != nil {
			return nil, err
		}
		return json, nil
	}
}
