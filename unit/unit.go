package unit

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
)

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
}

func SetResponseSuccess(c *gin.Context, response json.RawMessage) {
	c.Header("Content-Type", "application/json; charset=utf-8")

	c.JSON(0, json.NewEncoder(c.Writer).Encode(BasicResponse{
		0,
		response,
	}))
}

func ConvertStringToJsonFormat(message string, format Masterdata) ([]byte, error) {
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

func ConvertStringToDetailJsonFormat(message string, format Detaildata) ([]byte, error) {
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
