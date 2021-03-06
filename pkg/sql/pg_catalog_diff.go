// Copyright 2020 The Cockroach Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

// This file contains constants and types used by pg_catalog diff tool
// that are also re-used in /pkg/cmd/generate-pg-catalog

package sql

import (
	"encoding/json"
	"io"
	"os"
)

// GetPGCatalogSQL is a query uses udt_name::regtype instead of data_type column because
// data_type only says "ARRAY" but does not say which kind of array it is.
const GetPGCatalogSQL = `
	SELECT
		c.relname AS table_name,
		a.attname AS column_name,
		t.typname AS data_type,
		t.oid AS data_type_oid
	FROM pg_class c
	JOIN pg_attribute a ON a.attrelid = c.oid
	JOIN pg_type t ON t.oid = a.atttypid
	JOIN pg_namespace n ON n.oid = c.relnamespace
	WHERE n.nspname = 'pg_catalog'
	AND a.attnum > 0
	ORDER BY 1, 2;
`

// PGCatalogColumn is a structure which contains a small description about the datatype of a column, but this can also be
// used as a diff information if populating ExpectedOid. Fields are exported for Marshaling purposes.
type PGCatalogColumn struct {
	Oid              uint32  `json:"oid"`
	DataType         string  `json:"dataType"`
	ExpectedOid      *uint32 `json:"expectedOid"`
	ExpectedDataType *string `json:"expectedDataType"`
}

// PGCatalogColumns maps column names to datatype description
type PGCatalogColumns map[string]*PGCatalogColumn

// PGCatalogTables have 2 use cases:
// First: This is used to model pg_schema for postgres and cockroach db for comparison purposes by mapping tableNames
// to columns.
// Second: This is used to store and load expected diffs:
// - Using it this way, a table name pointing to a zero length PGCatalogColumns means that we expect this table to be missing
//   in cockroach db
// - If PGCatalogColumns is not empty but columnName points to null, we expect that column to be missing in that table in
//   cockroach db
// - If column Name points to a not null PGCatalogColumn, the test column describes how we expect that data type to be
//   different between cockroach db and postgres
type PGCatalogTables map[string]PGCatalogColumns

// PGCatalogFile is used to export pg_catalog from postgres and store the representation of this structure as a
// json file
type PGCatalogFile struct {
	PgVersion string          `json:"pgVersion"`
	PgCatalog PGCatalogTables `json:"pgCatalog"`
}

func (p PGCatalogTables) addColumn(tableName, columnName string, column *PGCatalogColumn) {
	columns, ok := p[tableName]

	if !ok {
		columns = make(PGCatalogColumns)
		p[tableName] = columns
	}

	columns[columnName] = column
}

// AddColumnMetadata is used to load data from postgres or cockroach pg_catalog schema
func (p PGCatalogTables) AddColumnMetadata(
	tableName string, columnName string, dataType string, dataTypeOid uint32,
) {
	p.addColumn(tableName, columnName, &PGCatalogColumn{
		dataTypeOid,
		dataType,
		nil,
		nil,
	})
}

// addDiff is for the second use case for pgTables which objective is create a datatype diff
func (p PGCatalogTables) addDiff(
	tableName string, columnName string, expected *PGCatalogColumn, actual *PGCatalogColumn,
) {
	p.addColumn(tableName, columnName, &PGCatalogColumn{
		actual.Oid,
		actual.DataType,
		&expected.Oid,
		&expected.DataType,
	})
}

// isDiffOid verifies if there is a datatype mismatch or if the diff is an expected diff
func (p PGCatalogTables) isDiffOid(
	tableName string, columnName string, expected *PGCatalogColumn, actual *PGCatalogColumn,
) bool {
	if expected.Oid == actual.Oid {
		return false
	}

	columns, ok := p[tableName]
	if !ok {
		return true
	}

	diff, ok := columns[columnName]
	if !ok {
		return true
	}

	return !(diff.Oid == actual.Oid && *diff.ExpectedOid == expected.Oid)
}

// isExpectedMissingTable is used by the diff PGCatalogTables to verify whether missing a table in cockroach is expected
// or not
func (p PGCatalogTables) isExpectedMissingTable(tableName string) bool {
	if columns, ok := p[tableName]; !ok || len(columns) > 0 {
		return false
	}

	return true
}

// isExpectedMissingColumn is similar to isExpectedMissingTable to verify column expected misses
func (p PGCatalogTables) isExpectedMissingColumn(tableName string, columnName string) bool {
	columns, ok := p[tableName]
	if !ok {
		return false
	}

	diff, ok := columns[columnName]
	if !ok {
		return false
	}

	return diff == nil
}

// addMissingTable adds a tablename when it is not found in cockroach db
func (p PGCatalogTables) addMissingTable(tableName string) {
	p[tableName] = make(PGCatalogColumns)
}

// addMissingColumn adds a column when it is not found in cockroach db
func (p PGCatalogTables) addMissingColumn(tableName string, columnName string) {
	columns, ok := p[tableName]

	if !ok {
		columns = make(PGCatalogColumns)
		p[tableName] = columns
	}

	columns[columnName] = nil
}

// rewriteDiffs creates pg_catalog_test-diffs.json
func (p PGCatalogTables) rewriteDiffs(diffFile string) error {
	f, err := os.OpenFile(diffFile, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	byteArray, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return err
	}

	if _, err = f.Write(byteArray); err != nil {
		return err
	}

	return nil
}

// Save have the purpose of storing all the data retrieved from postgres and useful information as postgres version
func (f *PGCatalogFile) Save(writer io.Writer) {
	byteArray, err := json.MarshalIndent(f, "", "  ")
	if err != nil {
		panic(err)
	}

	if _, err = writer.Write(byteArray); err != nil {
		panic(err)
	}
}
