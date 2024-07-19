package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type CreateUserTable_20240719_070516 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CreateUserTable_20240719_070516{}
	m.Created = "20240719_070516"

	migration.Register("CreateUserTable_20240719_070516", m)
}

// Run the migrations
func (m *CreateUserTable_20240719_070516) Up() {
	m.SQL(`
        CREATE TABLE users (
            id SERIAL PRIMARY KEY,
            department_id INTEGER NULL,
            name VARCHAR(255) NOT NULL,
            email VARCHAR(255) UNIQUE NOT NULL,
            password VARCHAR(255) NULL,
            deleted_at TIMESTAMP WITH TIME ZONE NULL
        );
    `)
}

// Reverse the migrations
func (m *CreateUserTable_20240719_070516) Down() {
	m.SQL(`DROP TABLE IF EXISTS users;`)
}
