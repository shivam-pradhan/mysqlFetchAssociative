package main

import "database/sql"

import (
	"fmt"
	"reflect"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// sample query
	query := "select * from testTable where true limit 3"
	arrayMap := fetchArrayMap(query)
	fmt.Println("Map from DB is :")
	fmt.Println(arrayMap)
}

func fetchArrayMap(query string) ([]map[string]interface{}) {
	// connection from your desired DB
	db, err := sql.Open("mysql","root:password@tcp(127.0.0.1:3306)/testDB" )

	// get the result of query
	rows, err := db.Query(query)
	tt, err := rows.ColumnTypes()
	if err != nil {
		checkErr(err)
	}
	columns, _ := rows.Columns()
	var resultRows []map[string]interface{}
	types := make([]reflect.Type, len(tt))
	for i, tp := range tt {
		st := tp.ScanType()
		if st == nil {
			checkErr(err)
			continue
		}
		types[i] = st
	}
	values := make([]interface{}, len(tt))
	for i := range values {
		values[i] = reflect.New(types[i]).Interface()
	}
	for rows.Next() {
		err = rows.Scan(values...)
		if err != nil {
			checkErr(err)
		}
		resultRow := map[string]interface{}{}
		for i, col := range columns {
			// check whether attribute datatype is struct
			if reflect.ValueOf(values[i]).Elem().Kind() == reflect.Struct {
				// 2nd field value false represent null value from DB
				if reflect.ValueOf(values[i]).Elem().Field(1).Interface() == false {
					resultRow[col] = nil
				} else {
					resultRow[col] = reflect.ValueOf(values[i]).Elem().Field(0).Interface()
				}
			} else {
				// if attribute type is slice then it will denotes a string
				if reflect.ValueOf(values[i]).Elem().Kind() == reflect.Slice {
					testString := fmt.Sprintf("%s", values[i])
					resultRow[col] = testString[1:len(testString)]
				} else {
					resultRow[col] = reflect.ValueOf(values[i]).Elem().Interface()
				}
			}
		}
		// append each rows in array
		resultRows = append(resultRows, resultRow)
	}
	db.Close()
	return resultRows
}

// checkErr ...
/**
 *
 *Check for error and panic the system start if any error while booting occured
 */
func checkErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}

/******* TEST TABLE STRUCTURE **************
	mysql> desc testTable;
+----------------+--------------+------+-----+---------+-------+
| Field          | Type         | Null | Key | Default | Extra |
+----------------+--------------+------+-----+---------+-------+
| id             | int(11)      | YES  |     | NULL    |       |
| full_names     | varchar(150) | NO   |     | NULL    |       |
| gender         | varchar(6)   | YES  |     | NULL    |       |
| date_of_birth  | date         | YES  |     | NULL    |       |
| contact_number | varchar(75)  | YES  |     | NULL    |       |
| email          | varchar(255) | YES  |     | NULL    |       |
+----------------+--------------+------+-----+---------+-------+

	*************** TABLE VALUES ******************
	mysql> select * from testTable;
+------+------------+--------+---------------+----------------+------------------+
| id   | full_names | gender | date_of_birth | contact_number | email            |
+------+------------+--------+---------------+----------------+------------------+
|    1 | testname1  | male   | 2019-01-01    | 9909123456     | example1@abc.com |
|    2 | testname2  | male   | NULL          | 9909123457     | example2@abc.com |
|    3 | testname3  | male   | 2019-01-02    | 9909123456     | NULL             |
+------+------------+--------+---------------+----------------+------------------+

	*************** ARRAY OF MAP **********************
[
	map[
		contact_number:9909123456
		date_of_birth:2019-01-01 00:00:00 +0000 UTC
		email:example1@abc.com
		full_names:testname1
		gender:male
		id:1
	]
	map[
		contact_number:9909123457
		date_of_birth:<nil>
		email:example2@abc.com
		full_names:testname2
		gender:male
		id:2
	]
	map[
		contact_number:9909123456
		date_of_birth:2019-01-02 00:00:00 +0000 UTC
		email:
		full_names:testname3
		gender:male
		id:3
	]
]


	***********************/