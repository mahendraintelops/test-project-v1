/*
 * Swagger Petstore
 *
 * This is a sample server Petstore server.  You can find out more about Swagger at [http://swagger.io](http://swagger.io) or on [irc.freenode.net, #swagger](http://swagger.io/irc/).  For this sample, you can use the api key `special-key` to test the authorization filters.
 *
 * API version: 1.0.6
 * Contact: apiteam@swagger.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"github.com/gin-gonic/gin"
)

type StoreAPI struct {
	// Delete /v2/store/order/:orderId
	// Delete purchase order by ID
	DeleteOrder gin.HandlerFunc
	// Get /v2/store/inventory
	// Returns pet inventories by status
	GetInventory gin.HandlerFunc
	// Get /v2/store/order/:orderId
	// Find purchase order by ID
	GetOrderById gin.HandlerFunc
	// Post /v2/store/order
	// Place an order for a pet
	PlaceOrder gin.HandlerFunc
}
