package ldb

import (
	"database/sql"
	"fmt"
	"log"

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

// AllAnchored todo doc
func (c connector) AllAnchored(idPortinformer string) []map[string]string {
	var idControlUnitData sql.NullString
	var shipName, anchoringTime, currentActivity, shippedGoods, tsReadiness sql.NullString
	var shipType, iso3, grossTonnage, anchoragePoint, length, width, agency, tsPlannedMooring sql.NullString
	var result []map[string]string = []map[string]string{}

	dot, err := dotsql.LoadFromFile("/home/deeper-x/go/src/github.com/deeper-x/krowki/qsql/realtime.sql")

	if err != nil {
		fmt.Println(err)
	}

	rows, err := dot.Query(c.db, "all-anchored", idPortinformer, idPortinformer)

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(
			&idControlUnitData,
			&shipName,
			&anchoringTime,
			&currentActivity,
			&anchoragePoint,
			&shipType,
			&iso3,
			&grossTonnage,
			&length,
			&width,
			&agency,
			&shippedGoods,
			&tsPlannedMooring,
			&tsReadiness,
		)

		if err != nil {
			log.Fatal(err)
		}

		idControlUnitDataStr := idControlUnitData

		tmpDict := map[string]string{
			"id_trip":            idControlUnitDataStr.String,
			"ship":               shipName.String,
			"ship_type":          shipType.String,
			"anchoring_time":     anchoringTime.String,
			"current_activity":   currentActivity.String,
			"anchorage_point":    anchoragePoint.String,
			"iso3":               iso3.String,
			"gross_tonnage":      grossTonnage.String,
			"length":             length.String,
			"width":              width.String,
			"agency":             agency.String,
			"shipped_goods":      shippedGoods.String,
			"ts_planned_mooring": tsPlannedMooring.String,
			"ts_readiness":       tsReadiness.String,
		}

		result = append(result, tmpDict)
	}

	return result
}

//ArrivalPrevisions todo doc
func (c connector) AllArrivalPrevisions(idPortinformer string) []map[string]string {
	var idControlUnitData, shipName sql.NullString
	var tsArrivalPrevision, shipType sql.NullString
	var shipFlag, shipWidth, shipLength, grossTonnage sql.NullString
	var netTonnage, draftAft, draftFwd sql.NullString
	var agency, cargoOnBoard sql.NullString
	var lastPortOfCall sql.NullString
	var destinationQuayBerth sql.NullString
	var destinationRoadstead sql.NullString

	var result []map[string]string = []map[string]string{}

	dot, err := dotsql.LoadFromFile("/home/deeper-x/go/src/github.com/deeper-x/krowki/qsql/realtime.sql")

	if err != nil {
		fmt.Println(err)
	}

	rows, err := dot.Query(c.db, "arrival-previsions", idPortinformer)

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(
			&shipName,
			&tsArrivalPrevision,
			&shipType,
			&shipFlag,
			&shipWidth,
			&shipLength,
			&grossTonnage,
			&netTonnage,
			&draftAft,
			&draftFwd,
			&agency,
			&lastPortOfCall,
			&destinationQuayBerth,
			&destinationRoadstead,
			&cargoOnBoard,
		)

		if err != nil {
			log.Fatal(err)
		}

		idControlUnitDataStr := idControlUnitData

		tmpDict := map[string]string{
			"id_trip":                idControlUnitDataStr.String,
			"ship":                   shipName.String,
			"ts_arrival_prevision":   tsArrivalPrevision.String,
			"ship_type":              shipType.String,
			"ship_flag":              shipFlag.String,
			"ship_width":             shipWidth.String,
			"ship_length":            shipLength.String,
			"gross_tonnage":          grossTonnage.String,
			"net_tonnage":            netTonnage.String,
			"draft_aft":              draftAft.String,
			"draft_fwd":              draftFwd.String,
			"agency":                 agency.String,
			"last_port_of_call":      lastPortOfCall.String,
			"destination_quay_berth": destinationQuayBerth.String,
			"destination_roadstead":  destinationRoadstead.String,
			"cargo_on_board":         cargoOnBoard.String,
		}

		result = append(result, tmpDict)
	}

	return result
}

//GetTodayArrivals todo doc
func (c connector) GetTodayArrivals(idPortinformer string) []map[string]string {
	var idTrip, shipName, shipType, tsSighting, shipFlag, shipWidth, shipLength sql.NullString
	var grossTonnage, netTonnage, draftAft, draftFwd, agency, lastPortOfCall sql.NullString
	var portDestination, destinationQuayBerth, destinationRoadstead sql.NullString

	var result []map[string]string = []map[string]string{}

	dot, err := dotsql.LoadFromFile("/home/deeper-x/go/src/github.com/deeper-x/krowki/qsql/realtime.sql")

	if err != nil {
		fmt.Println(err)
	}

	idState := 5
	rows, err := dot.Query(c.db, "arrivals", idState, idState, idState, idPortinformer, idState)

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(
			&idTrip,
			&shipName,
			&shipType,
			&tsSighting,
			&shipFlag,
			&shipWidth,
			&shipLength,
			&grossTonnage,
			&netTonnage,
			&draftAft,
			&draftFwd,
			&agency,
			&lastPortOfCall,
			&portDestination,
			&destinationQuayBerth,
			&destinationRoadstead,
		)

		if err != nil {
			log.Fatal(err)
		}

		tmpDict := map[string]string{
			"id_trip":                idTrip.String,
			"ship_name":              shipName.String,
			"ship_type":              shipType.String,
			"ts_sighting":            tsSighting.String,
			"ship_flag":              shipFlag.String,
			"ship_width":             shipWidth.String,
			"ship_length":            shipLength.String,
			"gross_tonnage":          grossTonnage.String,
			"net_tonnage":            netTonnage.String,
			"draft_aft":              draftAft.String,
			"draft_fwd":              draftFwd.String,
			"agency":                 agency.String,
			"last_port_of_call":      lastPortOfCall.String,
			"port_destination":       portDestination.String,
			"destination_quay_berth": destinationQuayBerth.String,
			"destination_roadstead":  destinationRoadstead.String,
		}

		result = append(result, tmpDict)
	}

	return result
}
