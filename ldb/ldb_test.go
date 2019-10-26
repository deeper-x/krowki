package ldb

import (
	"log"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestConnect(t *testing.T) {
	_, err := Connect()

	if err != nil {
		t.Errorf("%s", err)
	}

}

func TestAllMoored(t *testing.T) {
	db, mock, err := sqlmock.New()

	if err != nil {
		log.Println(err)
	}

	defer db.Close()

	expectedRows := sqlmock.NewRows([]string{"id_ship"})

	mock.ExpectQuery("SELECT id_control_unit_data, ship_description").WithArgs("28", "28").WillReturnRows(expectedRows)

	mConn := NewConnection(db)

	mConn.AllMoored("28")

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("%v", err)
	}
}
