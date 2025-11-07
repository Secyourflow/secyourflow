package migrations

import (
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	"github.com/pocketbase/pocketbase/models/schema"
)

func severityDeadlineUp(db dbx.Builder) error {
	dao := daos.New(db)

	collection, err := dao.FindCollectionByNameOrId(TicketCollectionName)
	if err != nil {
		return err
	}

	// Add severity field (text field with options: low, medium, high, critical)
	collection.Schema.AddField(&schema.SchemaField{
		Name: "severity",
		Type: schema.FieldTypeSelect,
		Options: &schema.SelectOptions{
			MaxSelect: 1,
			Values:    []string{"low", "medium", "high", "critical"},
		},
	})

	// Add deadline field (date field)
	collection.Schema.AddField(&schema.SchemaField{
		Name: "deadline",
		Type: schema.FieldTypeDate,
	})

	return dao.SaveCollection(collection)
}

func severityDeadlineDown(db dbx.Builder) error {
	dao := daos.New(db)

	collection, err := dao.FindCollectionByNameOrId(TicketCollectionName)
	if err != nil {
		return err
	}

	// Remove severity field
	collection.Schema.RemoveField(collection.Schema.GetFieldByName("severity").Id)

	// Remove deadline field
	collection.Schema.RemoveField(collection.Schema.GetFieldByName("deadline").Id)

	return dao.SaveCollection(collection)
}
