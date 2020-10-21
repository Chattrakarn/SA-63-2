package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field" 
	"github.com/facebookincubator/ent/schema/edge"
)
// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field {
		field.String("Username").NotEmpty(), 
		field.String("Email").NotEmpty(), 
		field.String("Password").NotEmpty(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge {
		edge.To("drugs" ,Drug.Type).StorageKey(edge.Column("user_id")),
	}
}


