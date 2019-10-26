package ldb

import (
	"fmt"
	"database/sql"
	"github.com/gchaincl/dotsql"
)

// AllMoored todo doc
func (c connector) AllMoored(idPortinformer string) []map[string]string {
	var idControlUnitData, iso3, grossTonnage, length, width, shipType sql.NullString
	var shipName, mooringTime, currentActivity, quay, shippedGoods sql.NullString
	var agency, tsETD sql.NullString

	var result []map[string]string

	dot, err := dotsql.LoadFromFile("/home/deeper-x/go/src/github.com/deeper-x/krowki/qsql/realtime.sql")

	if err != nil {
		fmt.Println(err)
	}

	rows, err := dot.Query(c.db, "all-moored", idPortinformer, idPortinformer)

	if err != nil {
		fmt.Println(err)
	}

	for rows.Next() {
		rows.Scan(
			&idControlUnitData,
			&shipName,
			&mooringTime,
			&currentActivity,
			&quay,
			&shippedGoods,
			&iso3,
			&grossTonnage,
			&length,
			&width,
			&shipType,
			&agency,
			&tsETD,
		)

		tmpDict := map[string]string{
			"id_trip":          idControlUnitData.String,
			"ship":             shipName.String,
			"mooring_time":     mooringTime.String,
			"current_activity": currentActivity.String,
			"quay":             quay.String,
			"shipped_goods":    shippedGoods.String,
			"iso3":             iso3.String,
			"gross_tonnage":    grossTonnage.String,
			"ships_length":     length.String,
			"ships_width":      width.String,
			"ship_type":        shipType.String,
			"agency":           agency.String,
			"ts_etd":           tsETD.String,
		}

		result = append(result, tmpDict)
	}
	return result
}
