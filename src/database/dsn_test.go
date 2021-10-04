package database

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_splitHostUrl(t *testing.T) {
	params := []struct {
		name    string
		url     string
		expHost string
		expPort string
	}{
		{"local", "localhost", "localhost", "3306"},
		{"local_port", "localhost:100", "localhost", "100"},
		{"ip", "192.168.170.2", "192.168.170.2", "3306"},
		{"ip_port", "192.168.170.2:100", "192.168.170.2", "100"},
		{"domain", "thetadev.de", "thetadev.de", "3306"},
		{"domain_port", "thetadev.de:100", "thetadev.de", "100"},
		{"double_port", "thetadev.de:50:100", "thetadev.de:50", "100"},
		{"str_port", "thetadev.de:abc", "thetadev.de:abc", "3306"},
	}

	for _, p := range params {
		t.Run(p.name, func(t *testing.T) {
			host, port := splitHostUrl(p.url, "3306")

			assert.Equal(t, p.expHost, host)
			assert.Equal(t, p.expPort, port)
		})
	}
}

func Test_getDsn(t *testing.T) {
	params := []struct {
		name   string
		conn   Connection
		expect string
	}{
		{
			"sqlite",
			Connection{Dialect: DialectSqlite, File: "test.db"},
			"test.db",
		},
		{
			"sqlite_default",
			Connection{Dialect: DialectSqlite},
			"database.db",
		},
		{
			"mysql",
			Connection{
				Dialect: DialectMySql,
				Host:    "thetadev.de",
				DbName:  "talon",
				User:    "test",
				Pass:    "1234",
			},
			`test:"1234"@tcp(thetadev.de:3306)/talon?charset=utf8&parseTime=True&loc=Local`,
		},
		{
			"mysql_port",
			Connection{
				Dialect: DialectMySql,
				Host:    "thetadev.de:100",
				DbName:  "talon",
				User:    "test",
				Pass:    "1234",
			},
			`test:"1234"@tcp(thetadev.de:100)/talon?charset=utf8&parseTime=True&loc=Local`,
		},
		{
			"mysql_default",
			Connection{Dialect: DialectMySql},
			`user:"pass"@tcp(127.0.0.1:3306)/db?charset=utf8&parseTime=True&loc=Local`,
		},
		{
			"postgres",
			Connection{
				Dialect: DialectPostgres,
				Host:    "thetadev.de",
				DbName:  "talon",
				User:    "test",
				Pass:    "1234",
			},
			`host=thetadev.de user=test password="1234" dbname=talon port=5432 sslmode=disable`,
		},
		{
			"postgres_port",
			Connection{
				Dialect: DialectPostgres,
				Host:    "thetadev.de:100",
				DbName:  "talon",
				User:    "test",
				Pass:    "1234",
			},
			`host=thetadev.de user=test password="1234" dbname=talon port=100 sslmode=disable`,
		},
		{
			"postgres_default",
			Connection{Dialect: DialectPostgres},
			`host=127.0.0.1 user=user password="pass" dbname=db port=5432 sslmode=disable`,
		},
	}

	for _, p := range params {
		t.Run(p.name, func(t *testing.T) {
			dsn := p.conn.getDsn()
			assert.Equal(t, p.expect, dsn)
		})
	}
}
