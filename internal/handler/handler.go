package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/aisha-deriv/migrate_affiliates_go_service/internal/domain/model"
)

// In-memory data store for demonstration
var items = []model.Item{
	{ID: 1, Name: "Item One", Price: 10.99},
	{ID: 2, Name: "Item Two", Price: 20.99},
}

// GetItems handles GET requests to retrieve all items
func GetItems(w http.ResponseWriter, r *http.Request) {
	response := Response[[]model.Item]{
		Data:       items,
		Success:    true,
		Message:    "Items retrieved successfully",
		StatusCode: http.StatusOK,
	}
	WriteJSON(w, response)
}

// GetItem handles GET requests to retrieve a single item by ID
func GetItem(w http.ResponseWriter, r *http.Request) {
	// Extract ID from query parameters
	ids, ok := r.URL.Query()["id"]
	if !ok || len(ids[0]) < 1 {
		response := Response[struct{}]{
			Success:    false,
			Message:    "Missing 'id' query parameter",
			StatusCode: http.StatusBadRequest,
		}
		WriteJSON(w, response)
		return
	}

	id, err := strconv.Atoi(ids[0])
	if err != nil {
		response := Response[struct{}]{
			Success:    false,
			Message:    "Invalid 'id' parameter",
			StatusCode: http.StatusBadRequest,
		}
		WriteJSON(w, response)
		return
	}

	// Find the item
	for _, item := range items {
		if item.ID == id {
			response := Response[model.Item]{
				Data:       item,
				Success:    true,
				Message:    "Item retrieved successfully",
				StatusCode: http.StatusOK,
			}
			WriteJSON(w, response)
			return
		}
	}

	// Item not found
	response := Response[struct{}]{
		Success:    false,
		Message:    "Item not found",
		StatusCode: http.StatusNotFound,
	}
	WriteJSON(w, response)
}

// CreateItem handles POST requests to create a new item
func CreateItem(w http.ResponseWriter, r *http.Request) {
	var newItem model.Item
	err := json.NewDecoder(r.Body).Decode(&newItem)
	if err != nil {
		response := Response[struct{}]{
			Success:    false,
			Message:    "Invalid request payload",
			StatusCode: http.StatusBadRequest,
		}
		WriteJSON(w, response)
		return
	}

	// Simple ID assignment for demonstration
	newItem.ID = len(items) + 1
	items = append(items, newItem)

	response := Response[model.Item]{
		Data:       newItem,
		Success:    true,
		Message:    "Item created successfully",
		StatusCode: http.StatusCreated,
	}
	WriteJSON(w, response)
}

// UpdateItem handles PUT requests to update an existing item
func UpdateItem(w http.ResponseWriter, r *http.Request) {
	var updatedItem model.Item
	err := json.NewDecoder(r.Body).Decode(&updatedItem)
	if err != nil {
		response := Response[struct{}]{
			Success:    false,
			Message:    "Invalid request payload",
			StatusCode: http.StatusBadRequest,
		}
		WriteJSON(w, response)
		return
	}

	// Find and update the item
	for i, item := range items {
		if item.ID == updatedItem.ID {
			items[i] = updatedItem
			response := Response[model.Item]{
				Data:       updatedItem,
				Success:    true,
				Message:    "Item updated successfully",
				StatusCode: http.StatusOK,
			}
			WriteJSON(w, response)
			return
		}
	}

	// Item not found
	response := Response[struct{}]{
		Success:    false,
		Message:    "Item not found",
		StatusCode: http.StatusNotFound,
	}
	WriteJSON(w, response)
}

// DeleteItem handles DELETE requests to remove an item by ID
func DeleteItem(w http.ResponseWriter, r *http.Request) {
	// Extract ID from query parameters
	ids, ok := r.URL.Query()["id"]
	if !ok || len(ids[0]) < 1 {
		response := Response[struct{}]{
			Success:    false,
			Message:    "Missing 'id' query parameter",
			StatusCode: http.StatusBadRequest,
		}
		WriteJSON(w, response)
		return
	}

	id, err := strconv.Atoi(ids[0])
	if err != nil {
		response := Response[struct{}]{
			Success:    false,
			Message:    "Invalid 'id' parameter",
			StatusCode: http.StatusBadRequest,
		}
		WriteJSON(w, response)
		return
	}

	// Find and delete the item
	for i, item := range items {
		if item.ID == id {
			items = append(items[:i], items[i+1:]...)
			response := Response[struct{}]{
				Success:    true,
				Message:    "Item deleted successfully",
				StatusCode: http.StatusOK,
			}
			WriteJSON(w, response)
			return
		}
	}

	// Item not found
	response := Response[struct{}]{
		Success:    false,
		Message:    "Item not found",
		StatusCode: http.StatusNotFound,
	}
	WriteJSON(w, response)
}
