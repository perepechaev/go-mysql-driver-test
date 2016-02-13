package mysql_test

import (
	"database/sql"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func TestConnectTimeout(t *testing.T) {
	db, err := sql.Open("mysql", "root@tcp(localhost:3307)/mysql?readTimeout=2s")

	timestamp := time.Now()
	err = db.Ping()
	duration := time.Since(timestamp)

	if err == nil {
		t.Log("expect bad connection")
		t.FailNow()
	}

	if duration.Seconds() < 2 || duration.Seconds() > 3 {
		t.Log("expect break connection after 2 seconds")
		t.FailNow()
	}

}

func TestSleepTimeout8Seconds(t *testing.T) {
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/mysql?readTimeout=2s")

	err = db.Ping()
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	timestamp := time.Now()
	_, err = db.Exec("SELECT SLEEP(8)")
	duration := time.Since(timestamp)
	if err == nil {
		t.Log("expect read timeout error")
		t.FailNow()
	}

	if duration.Seconds() < 6 || duration.Seconds() > 6+1 {
		t.Logf("expect break connection after 2 seconds. Duration: %0.2f", duration.Seconds())
		t.FailNow()
	}

}

func TestSleepTimeout4Seconds(t *testing.T) {
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/mysql?readTimeout=2s")

	err = db.Ping()
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	timestamp := time.Now()
	_, err = db.Exec("SELECT SLEEP(4)")
	duration := time.Since(timestamp)
	if err == nil {
		t.Log("expect read timeout error")
		t.FailNow()
	}

	if duration.Seconds() < 6 || duration.Seconds() > 6+1 {
		t.Logf("expect break connection after 2 seconds. Duration: %0.2f", duration.Seconds())
		t.FailNow()
	}

}
