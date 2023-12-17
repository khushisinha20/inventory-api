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

type RequestPayload struct {
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
}

func main() {
	app := gofr.New()

	app.POST("/addInventory", addInventoryHandler)

	app.GET("/allInventory", allInventoryHandler)

	app.GET("/getInventory/{id}", getInventoryHandler)

	app.PUT("/updateInventory/{id}", updateInventoryHandler)

	app.DELETE("/deleteInventory/{id}", deleteInventoryHandler)

	app.Start()
}

func addInventoryHandler(ctx *gofr.Context) (interface{}, error) {
	var requestPayload RequestPayload
	if err := json.NewDecoder(ctx.Request().Body).Decode(&requestPayload); err != nil {
		return nil, err
	}

	result, err := ctx.DB().ExecContext(ctx, "INSERT INTO inventory (name, quantity) VALUES (?, ?)", requestPayload.Name, requestPayload.Quantity)
	if err != nil {
		return nil, err
	}

	newID, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	return InventoryItem{ID: int(newID), Name: requestPayload.Name, Quantity: requestPayload.Quantity}, nil
}

func allInventoryHandler(ctx *gofr.Context) (interface{}, error) {
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
}

func getInventoryHandler(ctx *gofr.Context) (interface{}, error) {
	id := ctx.PathParam("id")

	row := ctx.DB().QueryRowContext(ctx, "SELECT * FROM inventory WHERE id=?", id)

	var item InventoryItem
	if err := row.Scan(&item.ID, &item.Name, &item.Quantity); err != nil {
		return nil, err
	}

	return item, nil
}

func updateInventoryHandler(ctx *gofr.Context) (interface{}, error) {
	var requestPayload RequestPayload
	if err := json.NewDecoder(ctx.Request().Body).Decode(&requestPayload); err != nil {
		return nil, err
	}

	itemID := ctx.Param("id")

	_, err := ctx.DB().ExecContext(ctx, "UPDATE inventory SET name = ?, quantity = ? WHERE id = ?", requestPayload.Name, requestPayload.Quantity, itemID)
	if err != nil {
		return nil, err
	}

	return "Inventory item updated successfully", nil
}

func deleteInventoryHandler(ctx *gofr.Context) (interface{}, error) {
	id := ctx.PathParam("id")

	_, err := ctx.DB().ExecContext(ctx, "DELETE FROM inventory WHERE id=?", id)
	if err != nil {
		return nil, err
	}

	return "Inventory item deleted successfully", nil
}
