// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"github.com/facebookincubator/ent/dialect/sql/schema"
	"github.com/facebookincubator/ent/schema/field"
)

var (
	// DrugsColumns holds the columns for the "drugs" table.
	DrugsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "drug_type", Type: field.TypeString},
		{Name: "strength", Type: field.TypeInt},
		{Name: "information", Type: field.TypeString},
		{Name: "form_id", Type: field.TypeInt, Nullable: true},
		{Name: "unit_id", Type: field.TypeInt, Nullable: true},
		{Name: "user_id", Type: field.TypeInt, Nullable: true},
		{Name: "volume_id", Type: field.TypeInt, Nullable: true},
	}
	// DrugsTable holds the schema information for the "drugs" table.
	DrugsTable = &schema.Table{
		Name:       "drugs",
		Columns:    DrugsColumns,
		PrimaryKey: []*schema.Column{DrugsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "drugs_forms_drugs",
				Columns: []*schema.Column{DrugsColumns[4]},

				RefColumns: []*schema.Column{FormsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "drugs_units_drugs",
				Columns: []*schema.Column{DrugsColumns[5]},

				RefColumns: []*schema.Column{UnitsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "drugs_users_drugs",
				Columns: []*schema.Column{DrugsColumns[6]},

				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "drugs_volumes_drugs",
				Columns: []*schema.Column{DrugsColumns[7]},

				RefColumns: []*schema.Column{VolumesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// FormsColumns holds the columns for the "forms" table.
	FormsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "form_type", Type: field.TypeString},
	}
	// FormsTable holds the schema information for the "forms" table.
	FormsTable = &schema.Table{
		Name:        "forms",
		Columns:     FormsColumns,
		PrimaryKey:  []*schema.Column{FormsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// UnitsColumns holds the columns for the "units" table.
	UnitsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "unit_type", Type: field.TypeString},
	}
	// UnitsTable holds the schema information for the "units" table.
	UnitsTable = &schema.Table{
		Name:        "units",
		Columns:     UnitsColumns,
		PrimaryKey:  []*schema.Column{UnitsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "username", Type: field.TypeString},
		{Name: "email", Type: field.TypeString},
		{Name: "password", Type: field.TypeString},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:        "users",
		Columns:     UsersColumns,
		PrimaryKey:  []*schema.Column{UsersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// VolumesColumns holds the columns for the "volumes" table.
	VolumesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "volume_type", Type: field.TypeString},
	}
	// VolumesTable holds the schema information for the "volumes" table.
	VolumesTable = &schema.Table{
		Name:        "volumes",
		Columns:     VolumesColumns,
		PrimaryKey:  []*schema.Column{VolumesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		DrugsTable,
		FormsTable,
		UnitsTable,
		UsersTable,
		VolumesTable,
	}
)

func init() {
	DrugsTable.ForeignKeys[0].RefTable = FormsTable
	DrugsTable.ForeignKeys[1].RefTable = UnitsTable
	DrugsTable.ForeignKeys[2].RefTable = UsersTable
	DrugsTable.ForeignKeys[3].RefTable = VolumesTable
}
