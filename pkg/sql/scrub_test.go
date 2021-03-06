// Copyright 2017 The Cockroach Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
// implied. See the License for the specific language governing
// permissions and limitations under the License.

package sql_test

import (
	gosql "database/sql"
	"strings"
	"testing"
	"time"

	"golang.org/x/net/context"

	"github.com/cockroachdb/cockroach/pkg/base"
	"github.com/cockroachdb/cockroach/pkg/sql"
	"github.com/cockroachdb/cockroach/pkg/sql/sem/tree"
	"github.com/cockroachdb/cockroach/pkg/sql/sqlbase"
	"github.com/cockroachdb/cockroach/pkg/testutils/serverutils"
	"github.com/cockroachdb/cockroach/pkg/util/leaktest"
)

type scrubResult struct {
	errorType  string
	database   string
	table      string
	primaryKey string
	timestamp  time.Time
	repaired   bool
	details    string
}

// getResultRows will scan and unmarshal scrubResults from a Rows
// iterator.
func getResultRows(rows *gosql.Rows) (results []scrubResult, err error) {
	defer rows.Close()

	var unused *string
	for rows.Next() {
		result := scrubResult{}
		if err := rows.Scan(
			// TODO(joey): In the future, SCRUB will run as a job during execution.
			&unused, /* job_uuid */
			&result.errorType,
			&result.database,
			&result.table,
			&result.primaryKey,
			&result.timestamp,
			&result.repaired,
			&result.details,
		); err != nil {
			return nil, err
		}
		results = append(results, result)
	}

	if rows.Err() != nil {
		return nil, err
	}

	return results, nil
}

// TestScrubIndexMissingIndexEntry tests that
// `SCRUB TABLE ... INDEX ALL`` will find missing index entries. To test
// this, a row's underlying secondary index k/v is deleted using the KV
// client. This causes a missing index entry error as the row is missing
// the expected secondary index k/v.
func TestScrubIndexMissingIndexEntry(t *testing.T) {
	defer leaktest.AfterTest(t)()
	s, db, kvDB := serverutils.StartServer(t, base.TestServerArgs{})
	defer s.Stopper().Stop(context.TODO())

	// Create the table and the row entry.
	if _, err := db.Exec(`
CREATE DATABASE t;
CREATE TABLE t.test (k INT PRIMARY KEY, v INT);
CREATE INDEX secondary ON t.test (v);
INSERT INTO t.test VALUES (10, 20);
`); err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	// Construct datums for our row values (k, v).
	values := []tree.Datum{tree.NewDInt(10), tree.NewDInt(20)}
	tableDesc := sqlbase.GetTableDescriptor(kvDB, "t", "test")
	secondaryIndex := &tableDesc.Indexes[0]

	colIDtoRowIndex := make(map[sqlbase.ColumnID]int)
	colIDtoRowIndex[tableDesc.Columns[0].ID] = 0
	colIDtoRowIndex[tableDesc.Columns[1].ID] = 1

	// Construct the secondary index key that is currently in the
	// database.
	secondaryIndexKey, err := sqlbase.EncodeSecondaryIndex(
		tableDesc, secondaryIndex, colIDtoRowIndex, values)
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	// Delete the entry.
	if err := kvDB.Del(context.TODO(), secondaryIndexKey.Key); err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	// Run SCRUB and find the index errors we created.
	rows, err := db.Query(`EXPERIMENTAL SCRUB TABLE t.test WITH OPTIONS INDEX ALL`)
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	results, err := getResultRows(rows)
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	if len(results) != 1 {
		t.Fatalf("expected 1 result, got %d. got %#v", len(results), results)
	}
	if result := results[0]; result.errorType != sql.ScrubErrorMissingIndexEntry {
		t.Fatalf("expected %q error, instead got: %s",
			sql.ScrubErrorMissingIndexEntry, result.errorType)
	} else if result.database != "t" {
		t.Fatalf("expected database %q, got %q", "t", result.database)
	} else if result.table != "test" {
		t.Fatalf("expected table %q, got %q", "test", result.table)
	} else if result.primaryKey != "(10)" {
		t.Fatalf("expected primaryKey %q, got %q", "(10)", result.primaryKey)
	} else if result.repaired {
		t.Fatalf("expected repaired %v, got %v", false, result.repaired)
	} else if !strings.Contains(result.details, `"v":"20"`) {
		t.Fatalf("expected erorr details to contain `%s`, got %s", `"v":"20"`, result.details)
	}
}

// TestScrubIndexDanglingIndexReference tests that
// `SCRUB TABLE ... INDEX`` will find dangling index references, which
// are index entries that have no corresponding primary k/v. To test
// this an index entry is generated and inserted. This creates a
// dangling index error as the corresponding primary k/v is not equal.
func TestScrubIndexDanglingIndexReference(t *testing.T) {
	defer leaktest.AfterTest(t)()
	s, db, kvDB := serverutils.StartServer(t, base.TestServerArgs{})
	defer s.Stopper().Stop(context.TODO())

	// Create the table and the row entry.
	if _, err := db.Exec(`
CREATE DATABASE t;
CREATE TABLE t.test (k INT PRIMARY KEY, v INT);
CREATE INDEX secondary ON t.test (v);
`); err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	tableDesc := sqlbase.GetTableDescriptor(kvDB, "t", "test")
	secondaryIndexDesc := &tableDesc.Indexes[0]

	colIDtoRowIndex := make(map[sqlbase.ColumnID]int)
	colIDtoRowIndex[tableDesc.Columns[0].ID] = 0
	colIDtoRowIndex[tableDesc.Columns[1].ID] = 1

	// Construct datums and secondary k/v for our row values (k, v).
	values := []tree.Datum{tree.NewDInt(10), tree.NewDInt(314)}
	secondaryIndex, err := sqlbase.EncodeSecondaryIndex(
		tableDesc, secondaryIndexDesc, colIDtoRowIndex, values)
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
	// Put the new secondary k/v into the database.
	if err := kvDB.Put(context.TODO(), secondaryIndex.Key, &secondaryIndex.Value); err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	// Run SCRUB and find the index errors we created.
	rows, err := db.Query(`EXPERIMENTAL SCRUB TABLE t.test WITH OPTIONS INDEX ALL`)
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	} else if rows.Err() != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	results, err := getResultRows(rows)
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	if len(results) != 1 {
		t.Fatalf("expected 1 result, got %d. got %#v", len(results), results)
	}
	if result := results[0]; result.errorType != sql.ScrubErrorDanglingIndexReference {
		t.Fatalf("expected %q error, instead got: %s",
			sql.ScrubErrorDanglingIndexReference, result.errorType)
	} else if result.database != "t" {
		t.Fatalf("expected database %q, got %q", "t", result.database)
	} else if result.table != "test" {
		t.Fatalf("expected table %q, got %q", "test", result.table)
	} else if result.primaryKey != "(10)" {
		t.Fatalf("expected primaryKey %q, got %q", "(10)", result.primaryKey)
	} else if result.repaired {
		t.Fatalf("expected repaired %v, got %v", false, result.repaired)
	} else if !strings.Contains(result.details, `"v":"314"`) {
		t.Fatalf("expected erorr details to contain `%s`, got %s", `"v":"314"`, result.details)
	}

	// Run SCRUB DATABASE to make sure it also catches the problem.
	rows, err = db.Query(`EXPERIMENTAL SCRUB DATABASE t`)
	if err != nil {
		t.Fatalf("unexpected error: %+v", err)
	} else if rows.Err() != nil {
		t.Fatalf("unexpected error: %+v", err)
	}
	scrubDatabaseResults, err := getResultRows(rows)
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	} else if len(scrubDatabaseResults) != 1 {
		t.Fatalf("expected 1 result, got %d. got %#v", len(scrubDatabaseResults), scrubDatabaseResults)
	} else if !(scrubDatabaseResults[0].errorType == results[0].errorType &&
		scrubDatabaseResults[0].database == results[0].database &&
		scrubDatabaseResults[0].table == results[0].table &&
		scrubDatabaseResults[0].details == results[0].details) {
		t.Fatalf("expected results to be equal, SCRUB TABLE got %v. SCRUB DATABASE got %v",
			results, scrubDatabaseResults)
	}
}

// TestScrubIndexCatchesStoringMismatch tests that
// `SCRUB TABLE ... INDEX ALL`` will fail if an index entry only differs
// by its STORING values. To test this, a row's underlying secondary
// index k/v is updated using the KV client to have a different value.
func TestScrubIndexCatchesStoringMismatch(t *testing.T) {
	defer leaktest.AfterTest(t)()
	s, db, kvDB := serverutils.StartServer(t, base.TestServerArgs{})
	defer s.Stopper().Stop(context.TODO())

	// Create the table and the row entry.
	if _, err := db.Exec(`
CREATE DATABASE t;
CREATE TABLE t.test (k INT PRIMARY KEY, v INT, data INT);
CREATE INDEX secondary ON t.test (v) STORING (data);
INSERT INTO t.test VALUES (10, 20, 1337);
`); err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	tableDesc := sqlbase.GetTableDescriptor(kvDB, "t", "test")
	secondaryIndexDesc := &tableDesc.Indexes[0]

	colIDtoRowIndex := make(map[sqlbase.ColumnID]int)
	colIDtoRowIndex[tableDesc.Columns[0].ID] = 0
	colIDtoRowIndex[tableDesc.Columns[1].ID] = 1
	colIDtoRowIndex[tableDesc.Columns[2].ID] = 2

	// Generate the existing secondary index key.
	values := []tree.Datum{tree.NewDInt(10), tree.NewDInt(20), tree.NewDInt(1337)}
	secondaryIndex, err := sqlbase.EncodeSecondaryIndex(
		tableDesc, secondaryIndexDesc, colIDtoRowIndex, values)
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
	// Delete the existing secondary k/v.
	if err := kvDB.Del(context.TODO(), secondaryIndex.Key); err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	// Generate a secondary index k/v that has a different value.
	values = []tree.Datum{tree.NewDInt(10), tree.NewDInt(20), tree.NewDInt(314)}
	secondaryIndex, err = sqlbase.EncodeSecondaryIndex(
		tableDesc, secondaryIndexDesc, colIDtoRowIndex, values)
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
	// Put the incorrect secondary k/v.
	if err := kvDB.Put(context.TODO(), secondaryIndex.Key, &secondaryIndex.Value); err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	// Run SCRUB and find the index errors we created.
	rows, err := db.Query(`EXPERIMENTAL SCRUB TABLE t.test WITH OPTIONS INDEX ALL`)
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	} else if rows.Err() != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	results, err := getResultRows(rows)
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	// We will receive both a missing_index_entry and dangling_index_reference.
	if len(results) != 2 {
		t.Fatalf("expected 2 result, got %d. got %#v", len(results), results)
	}

	// Assert the missing index error is correct.
	var missingIndexError *scrubResult
	for _, result := range results {
		if result.errorType == sql.ScrubErrorMissingIndexEntry {
			missingIndexError = &result
			break
		}
	}
	if result := missingIndexError; result == nil {
		t.Fatalf("expected errors to include %q error, but got errors: %#v",
			sql.ScrubErrorMissingIndexEntry, results)
	} else if result.database != "t" {
		t.Fatalf("expected database %q, got %q", "t", result.database)
	} else if result.table != "test" {
		t.Fatalf("expected table %q, got %q", "test", result.table)
	} else if result.primaryKey != "(10)" {
		t.Fatalf("expected primaryKey %q, got %q", "(10)", result.primaryKey)
	} else if result.repaired {
		t.Fatalf("expected repaired %v, got %v", false, result.repaired)
	} else if !strings.Contains(result.details, `"data":"1337"`) {
		t.Fatalf("expected erorr details to contain `%s`, got %s", `"data":"1337"`, result.details)
	}

	// Assert the dangling index error is correct.
	var danglingIndexResult *scrubResult
	for _, result := range results {
		if result.errorType == sql.ScrubErrorDanglingIndexReference {
			danglingIndexResult = &result
			break
		}
	}
	if result := danglingIndexResult; result == nil {
		t.Fatalf("expected errors to include %q error, but got errors: %#v",
			sql.ScrubErrorDanglingIndexReference, results)
	} else if result.database != "t" {
		t.Fatalf("expected database %q, got %q", "t", result.database)
	} else if result.table != "test" {
		t.Fatalf("expected table %q, got %q", "test", result.table)
	} else if result.primaryKey != "(10)" {
		t.Fatalf("expected primaryKey %q, got %q", "(10)", result.primaryKey)
	} else if result.repaired {
		t.Fatalf("expected repaired %v, got %v", false, result.repaired)
	} else if !strings.Contains(result.details, `"data":"314"`) {
		t.Fatalf("expected erorr details to contain `%s`, got %s", `"data":"314"`, result.details)
	}
}
