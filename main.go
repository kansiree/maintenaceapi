package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"maintenaceApi/databaseManage"
	"maintenaceApi/unit"
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
	response, err := databaseManage.SelectDataReturnJsonFormat("SELECT * FROM " + unit.MASTER_TECHNICAL_ORDER)
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
