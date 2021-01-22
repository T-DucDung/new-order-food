package models

import (
	"log"
)

func GetDataByQuery(query string) ([]map[string][]byte, error) {
	data := []map[string][]byte{}

	rows, err := db.Query(query)
	defer rows.Close()
	if err != nil {
		log.Println(err)
	}

	colNames, err := rows.Columns()
	if err != nil {
		log.Println(err)
	}

	cols := make([]interface{}, len(colNames))
	colPtrs := make([]interface{}, len(colNames))
	for i := 0; i < len(colNames); i++ {
		colPtrs[i] = &cols[i]
	}

	for rows.Next() {
		myMap := map[string][]byte{}

		err = rows.Scan(colPtrs...)
		if err != nil {
			log.Println(err)
		}
		for i, col := range cols {
			if col == nil {
				col = []byte("0")
			}
			myMap[colNames[i]] = col.([]byte)
		}
		data = append(data, myMap)
	}
	return data, nil
}

func ExecNonQuery(query string) error {
	_, err := db.Exec(query)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func InsertIntoDB(data map[string]string, nameTable string) error {
	query := "insert into " + nameTable
	fields := "("
	values := "values("
	total := len(data)
	for key, value := range data {
		fields += key
		values += value
		total--
		if total != 0 {
			fields += ","
			values += ","
		}
	}
	fields += ") "
	values += ")"
	query = query + fields + values
	log.Println("insert into db : ", query)
	//return ExecNonQuery(query)
	return nil
}
