package main

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
)

//go:generate stringer -type intColor -linecomment
type intColor int

const (
	intColorRed   intColor = iota // red
	intColorGreen                 // green
	intColorBlue                  // blue
	intColorWhite                 // white
	intColorBlack                 // black
)

var intColorsNames = map[string]intColor{
	intColorRed.String():   intColorRed,
	intColorGreen.String(): intColorGreen,
	intColorBlue.String():  intColorBlue,
	intColorWhite.String(): intColorWhite,
	intColorBlack.String(): intColorBlack,
}

func (in *intColor) Scan(val interface{}) error {
	switch v := val.(type) {
	case string:
		s, ok := intColorsNames[string(v)]
		if !ok {
			return fmt.Errorf("unknown color: %q", s)
		}

		*in = s
		return nil
	case []byte:
		s, ok := intColorsNames[string(v)]
		if !ok {
			return fmt.Errorf("unknown color: %q", s)
		}

		*in = s
		return nil
	}

	return fmt.Errorf("unsupported type %T", val)
}

func (in intColor) Value() (driver.Value, error) {
	return in.String(), nil
}

type lightInt struct {
	id    int
	color intColor
}

func insertIntColors(db *sql.DB) error {
	colors := []intColor{
		intColorRed,
		intColorGreen,
		intColorBlue,
		intColorWhite,
		intColorBlack,
	}

	for i, c := range colors {
		id := i + 1

		_, err := db.Exec("INSERT INTO lights (id, color) VALUES ($1, $2);", id, c)
		if err != nil {
			return fmt.Errorf("could not insert light %d: %w", id, err)
		}
	}

	return nil
}

func readIntColors(db *sql.DB) ([]lightInt, error) {
	rows, err := db.Query(`SELECT id, color FROM lights ORDER BY id`)
	if err != nil {
		return nil, fmt.Errorf("could not query db: %w", err)
	}
	defer rows.Close()

	var result []lightInt
	for rows.Next() {
		var li lightInt
		if err := rows.Scan(&li.id, &li.color); err != nil {
			return nil, fmt.Errorf("could not scan light int: %w", err)
		}

		result = append(result, li)
	}

	return result, nil
}

type stringColor string

const (
	stringColorRed   stringColor = "red"
	stringColorGreen stringColor = "green"
	stringColorBlue  stringColor = "blue"
	stringColorWhite stringColor = "white"
	stringColorBlack stringColor = "black"
)

type lightString struct {
	id    int
	color stringColor
}

func insertStringColors(db *sql.DB) error {
	colors := []stringColor{
		stringColorRed,
		stringColorGreen,
		stringColorBlue,
		stringColorWhite,
		stringColorBlack,
	}

	for i, c := range colors {
		id := i + 1

		_, err := db.Exec("INSERT INTO lights (id, color) VALUES ($1, $2);", id, c)
		if err != nil {
			return fmt.Errorf("could not insert light %d: %w", id, err)
		}
	}

	return nil
}

func readStringColors(db *sql.DB) ([]lightString, error) {
	rows, err := db.Query(`SELECT id, color FROM lights ORDER BY id`)
	if err != nil {
		return nil, fmt.Errorf("could not query db: %w", err)
	}
	defer rows.Close()

	var result []lightString
	for rows.Next() {
		var ls lightString
		if err := rows.Scan(&ls.id, &ls.color); err != nil {
			return nil, fmt.Errorf("could not scan light string: %w", err)
		}

		result = append(result, ls)
	}

	return result, nil
}
