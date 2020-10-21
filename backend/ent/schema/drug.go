package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field" 
	"github.com/facebookincubator/ent/schema/edge"
)

// Drug holds the schema definition for the Drug entity.
type Drug struct {
	ent.Schema
}

// Fields of the Drug.
func (Drug) Fields() []ent.Field {
	return []ent.Field {
		field.String("DrugType").NotEmpty(),
		field.Int("Strength").Positive(),
		field.String("Information").NotEmpty(),
	}
}

// Edges of the Drug.
func (Drug) Edges() []ent.Edge {
	return []ent.Edge {
		edge.From("user", User.Type).Ref("drugs").Unique(),
		edge.From("form", Form.Type).Ref("drugs").Unique(),
		edge.From("unit", Unit.Type).Ref("drugs").Unique(),
		edge.From("volume", Volume.Type).Ref("drugs").Unique(),
	}
}


