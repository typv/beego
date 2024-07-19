package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type CreateDepartmentTable_20240719_071649 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CreateDepartmentTable_20240719_071649{}
	m.Created = "20240719_071649"

	migration.Register("CreateDepartmentTable_20240719_071649", m)
}

// Run the migrations
func (m *CreateDepartmentTable_20240719_071649) Up() {
	m.SQL(`
        CREATE TABLE departments (
                       id SERIAL PRIMARY KEY,
                       name VARCHAR(255) NOT NULL
		);
    `)
}

// Reverse the migrations
func (m *CreateDepartmentTable_20240719_071649) Down() {
	m.SQL(`DROP TABLE IF EXISTS departments;`)
}
