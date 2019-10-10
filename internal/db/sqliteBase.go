package db

import (
	"database/sql"
	"kboard/utils"
	"log"
	"strings"
	"fmt"
	"errors"
)

type SQLiteBase struct {
	Table string
}

func (mdl *SQLiteBase) getTable() string {
	return mdl.Table
}

func (mdl *SQLiteBase) Insert(data map[string]interface{}) int64 {
	if len(data) <= 0 {
		log.Println("insert data empty")
		return 0
	}

	keys, values := mdl.GetKeysAndValues(data)
	sqlString := fmt.Sprintf("insert into %s (%s) VALUES (%s)", mdl.getTable(), strings.Join(keys, ","), utils.Implode(values, ","))

	result, err := DbConn.Exec(sqlString)
	if err != nil {
		log.Println(err)
		return 0
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		log.Println(err)
		return 0
	}

	return lastId
}

func (mdl *SQLiteBase) Update(cond map[string]interface{}, data map[string]interface{}) int64 {
	queryString := fmt.Sprintf("update %s set %s where %s", mdl.getTable(), mdl.BuildValue(data), mdl.BuildCond(cond))
	result, err := DbConn.Exec(queryString)
	if err != nil {
		return 0
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return 0
	}

	return rows
}

func (mdl *SQLiteBase) Select(fields string, cond map[string]interface{}, order string) (data []map[string]interface{}, err error) {
	if fields == "" {
		err = errors.New("fields is empty")
		return
	}

	queryString := fmt.Sprintf("select %s from %s where %s", fields, mdl.getTable(), mdl.BuildCond(cond))
	if order != "" {
		queryString += " order by " + order
	}
	rows, err := DbConn.Query(queryString)
	if err != nil {
		return
	}

	data = ParseRows(rows)

	return
}

func (mdl *SQLiteBase) Delete(cond map[string]interface{}) int64 {
	queryString := fmt.Sprintf("delete from %s where %s", mdl.getTable(), mdl.BuildCond(cond))
	result, err := DbConn.Exec(queryString)
	if err != nil {
		return 0
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return 0
	}

	return rows
}

func ParseRows(rows *sql.Rows) []map[string]interface{} {
	columns, _ := rows.Columns()
	columnLen := len(columns)
	scans := make([]interface{}, columnLen)
	values := []map[string]interface{}{}

	//为每一列初始化一个指针
	for index, _ := range scans {
		var a interface{}
		scans[index] = &a
	}

	for rows.Next() {
		_ = rows.Scan(scans...)

		item := make(map[string]interface{})
		for k, v := range scans {
			item[columns[k]] = *v.(*interface{}) //取实际类型
		}
		values = append(values, item)
	}
	_ = rows.Close()

	return values
}

func (mdl *SQLiteBase) BuildCond(cond map[string]interface{}) string {
	condSlice := []string{}
	if len(cond) <= 0 {
		return "1"
	}

	for k, v := range cond {
		switch v.(type) {
		case int8:
			condSlice = append(condSlice, fmt.Sprintf("`%s` = %d", k, v.(int8)))
		case int16:
			condSlice = append(condSlice, fmt.Sprintf("`%s` = %d", k, v.(int16)))
		case int:
			condSlice = append(condSlice, fmt.Sprintf("`%s` = %d", k, v.(int)))
		case int32:
			condSlice = append(condSlice, fmt.Sprintf("`%s` = %d", k, v.(int32)))
		case int64:
			condSlice = append(condSlice, fmt.Sprintf("`%s` = %d", k, v.(int64)))
		default:
			condSlice = append(condSlice, fmt.Sprintf("`%s` = '%s'", k, utils.AddSlashes(v.(string))))
		}
	}

	if len(condSlice) > 0 {
		return strings.Join(condSlice, " AND ")
	} else {
		return "1"
	}
}

func (mdl *SQLiteBase) BuildValue(values map[string]interface{}) string {
	valueSlice := []string{}
	if len(values) <= 0 {
		return ""
	}
	for k, v := range values {
		switch v.(type) {
		case int8:
			valueSlice = append(valueSlice, fmt.Sprintf("`%s` = %d", k, v.(int8)))
		case int16:
			valueSlice = append(valueSlice, fmt.Sprintf("`%s` = %d", k, v.(int16)))
		case int:
			valueSlice = append(valueSlice, fmt.Sprintf("`%s` = %d", k, v.(int)))
		case int32:
			valueSlice = append(valueSlice, fmt.Sprintf("`%s` = %d", k, v.(int32)))
		case int64:
			valueSlice = append(valueSlice, fmt.Sprintf("`%s` = %d", k, v.(int64)))
		default:
			valueSlice = append(valueSlice, fmt.Sprintf("`%s` = '%s'", k, utils.AddSlashes(v.(string))))
		}
	}

	return strings.Join(valueSlice, ",")
}

func (mdl *SQLiteBase) GetKeysAndValues(data map[string]interface{}) (keys []string, values []interface{}) {
	if len(data) <= 0 {
		return
	}

	for k, v := range data {
		keys = append(keys, k)
		switch v.(type) {
		case int8:
			values = append(values, v.(int8))
		case int16:
			values = append(values, v.(int16))
		case int:
			values = append(values, v.(int))
		case int32:
			values = append(values, v.(int32))
		case int64:
			values = append(values, v.(int64))
		default:
			values = append(values, "'"+ utils.AddSlashes(v.(string)) + "'")
		}
	}

	return
}

