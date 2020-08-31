package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type Masterdata []struct {
	CreatedDate string `json:"created_date"`
	FullName    string `json:"full_name"`
	ID          string `json:"id"`
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

type Masters []Masterdata

type Post struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

var posts []Post

func main() {
	handleRequest()
}

func homePage(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		panic(err)
	}
	id := r.FormValue("id")
	log.Println(id)
	defer r.Body.Close()
}
func handleRequest() {
	myRouter := mux.NewRouter()

	// Handle all preflight request
	myRouter.Methods("OPTIONS").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("OPTIONS")
		log.Println(w.Header().Get("Methods"))

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Access-Control-Request-Headers, Access-Control-Request-Method, Connection, Host, Origin, User-Agent, Referer, Cache-Control, X-header")
		w.WriteHeader(http.StatusNoContent)
		return
	})

	myRouter.StrictSlash(true)
	// mux := http.NewServeMux()

	myRouter.HandleFunc("/home", homePage).Methods("POST")
	myRouter.HandleFunc("/getMasterSystem", getMasterSystem).Methods("GET")
	myRouter.HandleFunc("/getMasterAircraft", getMasterAircraft).Methods("GET")
	myRouter.HandleFunc("/getMasterTechnicalOrder", getMasterTechnicalOrder).Methods("GET")
	myRouter.HandleFunc("/getDetail", getDetail).Methods("GET")

	http.ListenAndServe(getPort(), myRouter)
}

func getMasterSystem(w http.ResponseWriter, r *http.Request) {

	var masters Masterdata
	response, err := getJSON("SELECT * FROM t016ffukzsi0y5ie.master_system")
	if err != nil {
		log.Fatalln(err)
	} else {
		js, err := convertStringToJsonFormat(response, masters)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")

		w.Write(js)
	}

}

func getMasterAircraft(w http.ResponseWriter, r *http.Request) {

	var masters Masterdata
	response, err := getJSON("SELECT * FROM t016ffukzsi0y5ie.master_aircraft")
	if err != nil {
		log.Fatalln(err)
	} else {
		js, err := convertStringToJsonFormat(response, masters)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	}

}

func getMasterTechnicalOrder(w http.ResponseWriter, r *http.Request) {

	var masters Masterdata
	response, err := getJSON("SELECT * FROM t016ffukzsi0y5ie.master_technical_order")
	if err != nil {
		log.Fatalln(err)
	} else {
		js, err := convertStringToJsonFormat(response, masters)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	}

}

func insertDetail(w http.ResponseWriter, r *http.Request) {

}

func getDetail(w http.ResponseWriter, r *http.Request) {
	var dataDetail Detaildata
	response, err := getJSON("SELECT * FROM t016ffukzsi0y5ie.maintenace_detail")
	if err != nil {
		log.Fatalln(err)
	} else {
		js, err := convertStringToDetailJsonFormat(response, dataDetail)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	}
}

func getPort() string {
	var port = os.Getenv("PORT")
	if port == "" {
		port = "8080"
		fmt.Println("No Port In Heroku " + port)
	}
	return ":" + port
}

func getJSON(sqlString string) (string, error) {
	db, err := sql.Open("mysql", "sz0debklevf8wjhf:gu2af8swu50tjc3k@tcp(u3r5w4ayhxzdrw87.cbetxkdyhwsb.us-east-1.rds.amazonaws.com:3306)/t016ffukzsi0y5ie")
	if err != nil {
		log.Fatal(err)
	}
	rows, err := db.Query(sqlString)
	if err != nil {
		return "", err
	}
	defer rows.Close()
	columns, err := rows.Columns()
	if err != nil {
		return "", err
	}

	count := len(columns)
	tableData := make([]map[string]interface{}, 0)
	values := make([]interface{}, count)
	valuePtrs := make([]interface{}, count)

	for rows.Next() {
		for i := 0; i < count; i++ {
			valuePtrs[i] = &values[i]
		}
		rows.Scan(valuePtrs...)
		entry := make(map[string]interface{})
		for i, col := range columns {
			var v interface{}
			val := values[i]
			b, ok := val.([]byte)
			if ok {
				v = string(b)
			} else {
				v = val
			}
			entry[col] = v
		}
		tableData = append(tableData, entry)
	}
	jsonData, err := json.Marshal(tableData)
	if err != nil {
		return "", err
	}
	return string(jsonData), nil
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
