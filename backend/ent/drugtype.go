// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/B6001186/Contagions/ent/drugtype"
	"github.com/facebookincubator/ent/dialect/sql"
)

// DrugType is the model entity for the DrugType schema.
type DrugType struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// DrugTypeName holds the value of the "DrugTypeName" field.
	DrugTypeName string `json:"DrugTypeName,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the DrugTypeQuery when eager-loading is set.
	Edges DrugTypeEdges `json:"edges"`
}

// DrugTypeEdges holds the relations/edges for other nodes in the graph.
type DrugTypeEdges struct {
	// Drug holds the value of the drug edge.
	Drug []*Drug
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// DrugOrErr returns the Drug value or an error if the edge
// was not loaded in eager-loading.
func (e DrugTypeEdges) DrugOrErr() ([]*Drug, error) {
	if e.loadedTypes[0] {
		return e.Drug, nil
	}
	return nil, &NotLoadedError{edge: "drug"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*DrugType) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // DrugTypeName
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the DrugType fields.
func (dt *DrugType) assignValues(values ...interface{}) error {
	if m, n := len(values), len(drugtype.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	dt.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field DrugTypeName", values[0])
	} else if value.Valid {
		dt.DrugTypeName = value.String
	}
	return nil
}

// QueryDrug queries the drug edge of the DrugType.
func (dt *DrugType) QueryDrug() *DrugQuery {
	return (&DrugTypeClient{config: dt.config}).QueryDrug(dt)
}

// Update returns a builder for updating this DrugType.
// Note that, you need to call DrugType.Unwrap() before calling this method, if this DrugType
// was returned from a transaction, and the transaction was committed or rolled back.
func (dt *DrugType) Update() *DrugTypeUpdateOne {
	return (&DrugTypeClient{config: dt.config}).UpdateOne(dt)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (dt *DrugType) Unwrap() *DrugType {
	tx, ok := dt.config.driver.(*txDriver)
	if !ok {
		panic("ent: DrugType is not a transactional entity")
	}
	dt.config.driver = tx.drv
	return dt
}

// String implements the fmt.Stringer.
func (dt *DrugType) String() string {
	var builder strings.Builder
	builder.WriteString("DrugType(")
	builder.WriteString(fmt.Sprintf("id=%v", dt.ID))
	builder.WriteString(", DrugTypeName=")
	builder.WriteString(dt.DrugTypeName)
	builder.WriteByte(')')
	return builder.String()
}

// DrugTypes is a parsable slice of DrugType.
type DrugTypes []*DrugType

func (dt DrugTypes) config(cfg config) {
	for _i := range dt {
		dt[_i].config = cfg
	}
}
