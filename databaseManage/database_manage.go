package databaseManage

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"maintenaceApi/unit"
)

func ConnectDB() (*sql.DB, error) {
	var dataSource = unit.USER_NAME + ":" + unit.PASS_WORD + "@tcp(" + unit.HOST_NAME + ":" + unit.PORT + ")/" + unit.DATABAS_NAME
	return sql.Open("mysql", dataSource)
}

func QueryDB(sqlString string) (*sql.Rows, error) {
	db, err := ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	return db.Query(sqlString)
}

func SelectDataReturnJsonFormat(sqlString string) (string, error) {
	rows, err := QueryDB(sqlString)
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
	fmt.Println(tableData)
	jsonData, err := json.Marshal(tableData)
	if err != nil {
		return "", err
	}
	return string(jsonData), nil
}
