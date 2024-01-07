package migrations

import (
	"encoding/json"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/models"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		jsonData := `{
			"id": "6barvao062uij9a",
			"created": "2024-01-07 15:07:38.050Z",
			"updated": "2024-01-07 15:07:38.050Z",
			"name": "items",
			"type": "base",
			"system": false,
			"schema": [
				{
					"system": false,
					"id": "tenmxxsj",
					"name": "title",
					"type": "text",
					"required": false,
					"presentable": false,
					"unique": false,
					"options": {
						"min": null,
						"max": null,
						"pattern": ""
					}
				},
				{
					"system": false,
					"id": "db7avq1x",
					"name": "description",
					"type": "editor",
					"required": false,
					"presentable": false,
					"unique": false,
					"options": {
						"convertUrls": false
					}
				},
				{
					"system": false,
					"id": "tebm9baz",
					"name": "due_on",
					"type": "date",
					"required": false,
					"presentable": false,
					"unique": false,
					"options": {
						"min": "",
						"max": ""
					}
				},
				{
					"system": false,
					"id": "trf1tavp",
					"name": "completed_on",
					"type": "date",
					"required": false,
					"presentable": false,
					"unique": false,
					"options": {
						"min": "",
						"max": ""
					}
				}
			],
			"indexes": [
				"CREATE INDEX ` + "`" + `idx_RhfMH3V` + "`" + ` ON ` + "`" + `items` + "`" + ` (` + "`" + `due_on` + "`" + `)",
				"CREATE INDEX ` + "`" + `idx_nSJNk4W` + "`" + ` ON ` + "`" + `items` + "`" + ` (` + "`" + `completed_on` + "`" + `)"
			],
			"listRule": null,
			"viewRule": null,
			"createRule": null,
			"updateRule": null,
			"deleteRule": null,
			"options": {}
		}`

		collection := &models.Collection{}
		if err := json.Unmarshal([]byte(jsonData), &collection); err != nil {
			return err
		}

		return daos.New(db).SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("6barvao062uij9a")
		if err != nil {
			return err
		}

		return dao.DeleteCollection(collection)
	})
}
