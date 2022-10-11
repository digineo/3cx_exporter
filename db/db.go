package db

import (
	"database/sql"
	"errors"
	"net"
	"os"

	"time"

	"github.com/digineo/3cx_exporter/models"
	_ "github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
)

type Store struct {
	db *sql.DB
}

// Check db connection
func (s Store) Check() error {
	_, err := s.db.Exec("SELECT version()")
	return err

}

// Get list of instances
func (s Store) GetInstances() (instances []models.Instance, err error) {
	//Check username and password is not empty
	rows, err := s.db.Query("SELECT instance_id, instance_code,instance_username,instance_password,instance_host, instance_port FROM instances")
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		inst := models.Instance{}
		err = rows.Scan(&inst.InstanceId, &inst.Code, &inst.Login, &inst.Password, &inst.Host, &inst.Port)
		if err != nil {
			return
		}
		instances = append(instances, inst)
	}
	return

}

// Get last status
func (s Store) GetLastStatus() (status models.InstanceState, err error) {
	row := s.db.QueryRow(`SELECT 
	                        instance_state_id,
							instance_id,
							blacklist_size,
							calls_active,
							calls_limit,
							extensions_total,
							extensions_registred,
							last_backup,
							maintence_until,
							service_status,
							service_cpu,service_memory,
							trunc_registred,
							created_at
						FROM instances_state ORDER BY instance_state_id DESC LIMIT 1`)
	err = row.Scan(&status.Id,
		&status.InstanceId,
		&status.BlacklistSize,
		&status.CallsActive,
		&status.CallsLimit,
		&status.ExtensionsTotal,
		&status.ExtensionsRegistred,
		&status.LastBackUp,
		&status.MaintenceUntil,
		&status.ServiceStatus,
		&status.ServiceCPU,
		&status.ServiceMemory,
		&status.TruncRegistred,
		&status.CreatedAt,
	)
	return

}

// New Status
func (s Store) NewStatus(status models.InstanceState) (err error) {

	stmt, err := s.db.Prepare(`INSERT INTO instances_state(
		                        instance_id,
								blacklist_size,
							    calls_active,
							    calls_limit,
							    extensions_total,
							    extensions_registred,
							    last_backup,
							    maintence_until,
							    service_status,
							    service_cpu,
								service_memory,
							    trunc_registred,
							    created_at
							)
							   VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?)
								
								`)
	if err != nil {
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec(status.InstanceId,
		status.BlacklistSize,
		status.CallsActive,
		status.CallsLimit,
		status.ExtensionsTotal,
		status.ExtensionsRegistred,
		status.LastBackUp,
		status.MaintenceUntil,
		status.ServiceStatus,
		status.ServiceCPU,
		status.ServiceMemory,
		status.TruncRegistred,
		status.CreatedAt,
	)
	return

}

// Close db client and all connection pool
func (s Store) Close() error {
	return s.db.Close()
}

func New(dbConnection string, log *zap.Logger) *Store {

	var db *sql.DB
	var err error
	db, err = sql.Open("mysql", dbConnection)

	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(1)
	db.SetConnMaxIdleTime(2 * time.Second)

	thisStore := &Store{db: db}
	err = thisStore.Check()
	if err != nil {
		switch err.(type) {
		case *net.OpError:
			unwrapped := errors.Unwrap(err)
			switch unwrapped.(type) {
			case *net.DNSError:
				log.Panic("DNS ERROR. RDS HOST DOES NOT EXISTS")

			case *os.SyscallError:
				log.Panic("RDS CONNECTION TIMEOUT. MAYBE ACCESS DENIED AT THE NETWORK LEVEL")
			default:
				log.Panic("RDS CONNECTION OPERATION ERROR", zap.Error(err))
			}
		default:
			log.Panic("RDS CONNECTION ERROR", zap.Error(err))

		}
	}
	return thisStore

}
