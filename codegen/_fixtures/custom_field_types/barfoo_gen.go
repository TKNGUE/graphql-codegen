// This code is genereated by graphql-codegen
// DO NOT EDIT!

package custom_field_types

import (
	"encoding/json"

	graphql "github.com/neelance/graphql-go"
)

// BarFoo
type BarFoo struct {
	StartCursor *graphql.ID `json:"startCursor"`
	EndCursor   *graphql.ID `json:"endCursor"`
	HasNextPage bool        `json:"hasNextPage"`
}

// BarFooResolver resolver for BarFoo
type BarFooResolver struct {
	BarFoo
}

func (r *BarFooResolver) StartCursor() *graphql.ID {
	return r.BarFoo.StartCursor
}

func (r *BarFooResolver) EndCursor() *graphql.ID {
	return r.BarFoo.EndCursor
}

func (r *BarFooResolver) HasNextPage() bool {
	return r.BarFoo.HasNextPage
}

func (r *BarFooResolver) MarshalJSON() ([]byte, error) {
	return json.Marshal(&r.BarFoo)
}

func (r *BarFooResolver) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.BarFoo)
}