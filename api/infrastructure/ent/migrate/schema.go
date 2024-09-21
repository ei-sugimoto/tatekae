// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// ProjectsColumns holds the columns for the "projects" table.
	ProjectsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "created_by", Type: field.TypeInt},
	}
	// ProjectsTable holds the schema information for the "projects" table.
	ProjectsTable = &schema.Table{
		Name:       "projects",
		Columns:    ProjectsColumns,
		PrimaryKey: []*schema.Column{ProjectsColumns[0]},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "username", Type: field.TypeString, Unique: true},
		{Name: "password", Type: field.TypeString},
		{Name: "email", Type: field.TypeString, Unique: true},
		{Name: "created_at", Type: field.TypeTime},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// UserProjectsColumns holds the columns for the "user_projects" table.
	UserProjectsColumns = []*schema.Column{
		{Name: "user_id", Type: field.TypeInt},
		{Name: "project_id", Type: field.TypeInt},
	}
	// UserProjectsTable holds the schema information for the "user_projects" table.
	UserProjectsTable = &schema.Table{
		Name:       "user_projects",
		Columns:    UserProjectsColumns,
		PrimaryKey: []*schema.Column{UserProjectsColumns[0], UserProjectsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "user_projects_user_id",
				Columns:    []*schema.Column{UserProjectsColumns[0]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "user_projects_project_id",
				Columns:    []*schema.Column{UserProjectsColumns[1]},
				RefColumns: []*schema.Column{ProjectsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		ProjectsTable,
		UsersTable,
		UserProjectsTable,
	}
)

func init() {
	UserProjectsTable.ForeignKeys[0].RefTable = UsersTable
	UserProjectsTable.ForeignKeys[1].RefTable = ProjectsTable
}
