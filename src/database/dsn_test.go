package database

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_splitHostUrl(t *testing.T) {
	tests := []struct {
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

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			host, port := splitHostUrl(tt.url, "3306")

			assert.Equal(t, tt.expHost, host)
			assert.Equal(t, tt.expPort, port)
		})
	}
}

func Test_prepare(t *testing.T) {
	tests := []struct {
		name   string
		conn   Connection
		expect Connection
		err    error
	}{
		{
			name:   "sqlite",
			conn:   Connection{Dialect: DialectSqlite, File: "test.db"},
			expect: Connection{Dialect: DialectSqlite, File: "test.db", DbName: "test.db"},
		},
		{
			name: "sqlite_default",
			conn: Connection{},
			expect: Connection{
				Dialect: DialectSqlite,
				File:    "database.db",
				DbName:  "database.db",
			},
		},
		{
			name: "mysql",
			conn: Connection{
				Dialect: DialectMySql,
				Host:    "thetadev.de",
				DbName:  "talon",
				User:    "test",
				Pass:    "1234",
			},
			expect: Connection{
				Dialect: DialectMySql,
				Host:    "thetadev.de",
				DbName:  "talon",
				User:    "test",
				Pass:    "1234",
			},
		},
		{
			name: "mysql_default",
			conn: Connection{
				Dialect: DialectMySql,
				DbName:  "talon",
				User:    "test",
				Pass:    "1234",
			},
			expect: Connection{
				Dialect: DialectMySql,
				Host:    "127.0.0.1",
				DbName:  "talon",
				User:    "test",
				Pass:    "1234",
			},
		},
		{
			name: "mysql_no_db",
			conn: Connection{
				Dialect: DialectMySql,
				User:    "test",
				Pass:    "1234",
			},
			err: ErrEmptyDbName,
		},
		{
			name: "mysql_no_user",
			conn: Connection{
				Dialect: DialectMySql,
				DbName:  "talon",
				Pass:    "1234",
			},
			err: ErrEmptyDbUsername,
		},
		{
			name: "mysql_no_pw",
			conn: Connection{
				Dialect: DialectMySql,
				DbName:  "talon",
				User:    "test",
			},
			err: ErrEmptyDbPassword,
		},
		{
			name: "postgres",
			conn: Connection{
				Dialect: DialectPostgres,
				Host:    "thetadev.de",
				DbName:  "talon",
				User:    "test",
				Pass:    "1234",
			},
			expect: Connection{
				Dialect: DialectPostgres,
				Host:    "thetadev.de",
				DbName:  "talon",
				User:    "test",
				Pass:    "1234",
			},
		},
		{
			name: "postgres_default",
			conn: Connection{
				Dialect: DialectPostgres,
				DbName:  "talon",
				User:    "test",
				Pass:    "1234",
			},
			expect: Connection{
				Dialect: DialectPostgres,
				Host:    "127.0.0.1",
				DbName:  "talon",
				User:    "test",
				Pass:    "1234",
			},
		},
		{
			name: "invalid",
			conn: Connection{
				Dialect: "XYZ",
			},
			err: ErrUnknownSqlDialect,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &tt.conn
			err := c.prepare()

			if tt.err == nil {
				assert.Nil(t, err)
				assert.Equal(t, *c, tt.expect)
			} else {
				assert.ErrorIs(t, err, tt.err)
			}
		})
	}
}

func Test_getDsn(t *testing.T) {
	tests := []struct {
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
			"mysql",
			Connection{
				Dialect: DialectMySql,
				Host:    "thetadev.de",
				DbName:  "talon",
				User:    "test",
				Pass:    "1234",
			},
			`test:1234@tcp(thetadev.de:3306)/talon?charset=utf8&parseTime=True&loc=Local`,
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
			`test:1234@tcp(thetadev.de:100)/talon?charset=utf8&parseTime=True&loc=Local`,
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
			`host=thetadev.de user=test password=1234 dbname=talon port=5432 sslmode=disable`,
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
			`host=thetadev.de user=test password=1234 dbname=talon port=100 sslmode=disable`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dsn := tt.conn.getDsn()
			assert.Equal(t, tt.expect, dsn)
		})
	}
}
