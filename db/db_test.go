package db

import (
	"database/sql"
	"os"
	"testing"
	"time"

	"github.com/digineo/3cx_exporter/models"
)

// Get db connection string. Create tables (db must be clean)
func testDbPrepere(testDb string) (*sql.DB, error) {
	db, err := sql.Open("mysql", testDb)
	if err != nil {
		return db, err
	}
	tx, err := db.Begin()
	if err != nil {
		return db, err
	}
	defer tx.Rollback()
	_, err = tx.Exec("CREATE TABLE IF NOT EXISTS instances(instance_id SERIAL PRIMARY KEY, user_id INTEGER, instance_code VARCHAR(255),instance_username VARCHAR(255),instance_password VARCHAR(255), instance_host VARCHAR(255), instance_port VARCHAR(255), update_time DATETIME, created_at DATETIME) ")
	if err != nil {
		return db, err
	}
	_, err = tx.Exec("CREATE TABLE IF NOT EXISTS instances_state(instance_state_id SERIAL PRIMARY KEY, instance_id INTEGER, blacklist_size INTEGER,calls_active INTEGER,calls_limit INTEGER, extensions_total INTEGER, extensions_registred INTEGER, last_backup DATETIME, maintence_until DATETIME, service_status VARCHAR(255),service_cpu VARCHAR(255), service_memory VARCHAR(255), trunc_registred VARCHAR(255), created_at DATETIME) ")
	if err != nil {
		return db, err
	}
	stmt, err := tx.Prepare("INSERT INTO instances(user_id,instance_code,instance_username,instance_password,instance_host,instance_port,update_time,created_at ) VALUES(?,?,?,?,?,?,?,?)")
	if err != nil {
		return db, err
	}
	defer stmt.Close()
	_, err = stmt.Exec(1, "TEST", "TEST_USER", "TEST_PASS", "TEST_HOST", "TEST_PORT", time.Now(), time.Now())
	if err != nil {
		return db, err
	}
	err = tx.Commit()
	return db, err

}
func testDbClose(db *sql.DB) {
	db.Exec("DROP TABLE instances")
	db.Exec("DROP TABLE instances_state")
	db.Close()

}

// Test get instance, creating new instance state and get it
// This is integration test. To run it set TEST_DB environvent variable thats equal test db connection string
// Format TEST_DB connection string is username:password@protocol(address)/dbname?parseTime=true
// Use ?parseTime=true flag
func TestDB(t *testing.T) {
	connectionString := os.Getenv("TEST_DB")
	if connectionString == "" {
		t.Skip()
	}
	db, err := testDbPrepere(connectionString)
	if err != nil {
		t.Error(err)
		return
	}
	//defer testDbClose(db)
	store := Store{db: db}
	instance, err := store.GetInstances()
	if err != nil {
		t.Error(err)
		return
	}
	if len(instance) != 1 {
		t.Errorf("No instances present")
		return
	}

	if instance[0].Code != "TEST" || instance[0].Host != "TEST_HOST" || instance[0].Login != "TEST_USER" || instance[0].Password != "TEST_PASS" {
		t.Errorf("Bad instance attribute")
		return
	}
	status := models.InstanceState{
		Id:                  1,
		InstanceId:          instance[0].InstanceId,
		BlacklistSize:       10,
		CallsActive:         5,
		CallsLimit:          10,
		ExtensionsTotal:     100,
		ExtensionsRegistred: 10,
		ServiceStatus:       "OK",
		ServiceCPU:          "10",
		ServiceMemory:       "8",
		TruncRegistred:      "1",
		CreatedAt:           time.Now(),
	}
	err = store.NewStatus(status)
	if err != nil {
		t.Error(err)
		return
	}
	_, err = store.GetLastStatus()
	if err != nil {
		t.Error(err)
		return
	}

}
