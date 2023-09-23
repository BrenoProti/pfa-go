package entity_test

import (
	"testing"

	"github.com/BrenoProti/pfa-go/internal/order/entity"

	"github.com/stretchr/testify/assert"
)

func TestGivenAndEmptyId_WhenCreateANewOder_ThenShouldReceivedAndError(t *testing.T) {
	order := entity.Order{}

	assert.Error(t, order.IsValid(), "ID is required")
}

func TestGivenAndEmptyPrice_WhenCreateANewOder_ThenShouldReceivedAndError(t *testing.T) {
	order := entity.Order{ID: "123"}

	assert.Error(t, order.IsValid(), "Price is required")
}

func TestGivenAndEmptyTax_WhenCreateANewOder_ThenShouldReceivedAndError(t *testing.T) {
	order := entity.Order{ID: "123", Price: 10}

	assert.Error(t, order.IsValid(), "Tax is required")
}

func TestGivenAValidParams_WhenCallNewOder_ThenShouldReceivedOrderWithAllParams(t *testing.T) {
	order, err := entity.NewOrder("123", 10, 10)
	assert.NoError(t, err)
	assert.Equal(t, "123", order.ID)
	assert.Equal(t, 10.0, order.Price)
	assert.Equal(t, 10.0, order.Tax)
}

func TestGivenAValidParams_WhenCallCalculateFinalPrice_ThenShouldCalculateFinalPrice(t *testing.T) {
	order, err := entity.NewOrder("123", 10, 10)
	assert.NoError(t, err)
	order.CalculateFinalPrice()
	assert.Equal(t, 20.0, order.FinalPrice)
}
