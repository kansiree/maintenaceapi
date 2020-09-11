package main

<<<<<<< HEAD
import (
	"encoding/json"
	"net/http"
)

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
=======
// Basic response struct
type BasicResponse struct {
	Error   int    `json:"error"`
	Message string `json:"message"`
>>>>>>> 1283c7b6601c5cc5d164e284d9439b1e27d93867
}
