package proxy_database

import (
	"encoding/json"
	"os"
)

type Database struct {
	File string
}

func NewDatabase(file string) *Database {
	return &Database{File: file}
}

func (db *Database) Set(input Input) {
	data := Data{}.FromFile(db.File)

	data[input.Id] = input.Content

	data.ToFile(db.File)
}

func (db *Database) Get(id string) string {
	data := Data{}.FromFile(db.File)

	value, ok := data[id]
	if !ok {
		return "not found"
	}

	return value
}

// data structure
type Data map[string]string

func (d Data) FromFile(file string) Data {
	dbBytes, _ := os.ReadFile(file)

	var data Data
	json.Unmarshal(dbBytes, &data)

	if data == nil {
		data = make(Data)
	}

	return data
}

func (d Data) ToFile(file string) {
	updatedDbBytes, _ := json.Marshal(d)
	os.WriteFile(file, updatedDbBytes, os.ModePerm)
}
