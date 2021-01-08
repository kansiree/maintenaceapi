package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"maintenace/databaseManage"
	"maintenace/unit"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type BasicResponseStringMessage struct {
	Error   int    `json:"error"`
	Message string `json:"message"`
}

type BasicResponse struct {
	Error   int             `json:"error"`
	Message json.RawMessage `json:"message"`
}

type ImageStructure struct {
	ImageName string
	URL       string
}

func main() {

	router := gin.Default()

	router.GET("/", homePage)
	router.GET("getMasterSystem", getMasterSystem)
	router.GET("getMasterAircraft", getMasterAircraft)
	router.GET("getMasterTechnicalOrder", getMasterTechnicalOrder)
	//	router.POST("uploadImage", uploadImage)
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
	var masters unit.Masterdata
	response, err := databaseManage.SelectDataReturnJsonFormat("SELECT * FROM " + unit.MASTER_SYSTEM_TABLE)
	if err != nil {
		log.Fatalln(err)
	} else {
		js, err := unit.ConvertStringToJsonFormat(response, masters)
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

	var masters unit.Masterdata
	response, err := databaseManage.SelectDataReturnJsonFormat("SELECT * FROM " + unit.MASTER_AIRCRAFT)
	if err != nil {
		log.Fatalln(err)
	} else {
		js, err := unit.ConvertStringToJsonFormat(response, masters)
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

	var masters unit.Masterdata
	response, err := databaseManage.SelectDataReturnJsonFormat("SELECT * FROM " + unit.MASTER_TECHNICAL_ORDER)
	if err != nil {
		log.Fatalln(err)
	} else {
		js, err := unit.ConvertStringToJsonFormat(response, masters)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return
		}
		c.Header("Content-Type", "application/json; charset=utf-8")
		dataRes := &BasicResponse{0, js}
		c.JSON(http.StatusOK, dataRes)
	}

}

func getDetail(c *gin.Context) {
	var dataDetail unit.Detaildata

	response, err := databaseManage.SelectDataReturnJsonFormat("SELECT * FROM " + unit.MAINTANACE_DETAIL)
	if err != nil {
		log.Fatalln(err)
	} else {
		js, err := unit.ConvertStringToDetailJsonFormat(response, dataDetail)
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

func insertDetail(c *gin.Context) {
	var detailData unit.Detaildata
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
	stmt, err := db.Prepare("insert into " + unit.MAINTANACE_DETAIL + " values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?) ")
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

func getPort() string {
	var port = os.Getenv("PORT")
	if port == "" {
		port = "8080"
		fmt.Println("No Port In Heroku " + port)
	}
	return ":" + port
}
