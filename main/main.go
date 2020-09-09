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

// 1) Struct for a Route
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type BasicResponse struct {
	Error   int             `json:"error"`
	Message json.RawMessage `json:"message"`
}

type Routes []Route

func getIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello world!")
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
func handleRequest() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)

	var routes = Routes{
		Route{
			"getIndex",
			"GET",
			"/",
			getIndex,
		},
		Route{
			"home",
			"GET",
			"/",
			homePage,
		},
		Route{
			"getMasterSystem",
			"GET",
			"/getMasterSystem",
			getMasterSystem,
		},
		Route{
			"getMasterAircraft",
			"GET",
			"/",
			getMasterAircraft,
		},
		Route{
			"getMasterTechnicalOrder",
			"GET",
			"/",
			getMasterTechnicalOrder,
		},
		Route{
			"getDetail",
			"GET",
			"/",
			getDetail,
		},
	}

	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}
	return router
}

func main() {
	router := handleRequest()

	log.Fatal(
		// start on port
		http.ListenAndServe(getPort(), router),
	)
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
		println(response)
		w.Header().Set("Content-type", "application/json; charset=UTF-8;")
		json.NewEncoder(w).Encode(BasicResponse{
			0,
			(js),
		})
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
