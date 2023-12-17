package main

import (
	"encoding/json"
	"gofr.dev/pkg/gofr"
)

type InventoryItem struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
}

func main() {
	app := gofr.New()
	
	type RequestPayload struct {
		Name     string `json:"name"`
		Quantity int    `json:"quantity"`
	}

	app.POST("/addInventory", func(ctx *gofr.Context) (interface{}, error) {
		var requestPayload RequestPayload
		if err := json.NewDecoder(ctx.Request().Body).Decode(&requestPayload); err != nil {
			return nil, err
		}

		// Insert data into the 'inventory' table
		result, err := ctx.DB().ExecContext(ctx, "INSERT INTO inventory (name, quantity) VALUES (?, ?)", requestPayload.Name, requestPayload.Quantity)
		if err != nil {
			return nil, err
		}

		// Get the ID of the newly inserted item
		newID, err := result.LastInsertId()
		if err != nil {
			return nil, err
		}

		return InventoryItem{ID: int(newID), Name: requestPayload.Name, Quantity: requestPayload.Quantity}, nil
	})

	// Read all inventory items
	app.GET("/allInventory", func(ctx *gofr.Context) (interface{}, error) {
		var inventoryItems []InventoryItem

		rows, err := ctx.DB().QueryContext(ctx, "SELECT * FROM inventory")
		if err != nil {
			return nil, err
		}
		defer rows.Close()

		for rows.Next() {
			var item InventoryItem
			if err := rows.Scan(&item.ID, &item.Name, &item.Quantity); err != nil {
				return nil, err
			}
			inventoryItems = append(inventoryItems, item)
		}

		return inventoryItems, nil
	})


	app.Start()
}
