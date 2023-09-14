package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_If_It_Gets_An_Error_If_ID_Is_Black(t *testing.T) {
	order := &Order{}
	assert.Error(t, order.Validate(), "id is required")
}

func Test_If_It_Gets_An_Error_If_Price_Is_Black(t *testing.T) {
	order := &Order{
		ID: "1",
	}
	assert.Error(t, order.Validate(), "price must be greater than 0")
}

func Test_If_It_Gets_An_Error_If_Tax_Is_Black(t *testing.T) {
	order := &Order{
		ID:    "1",
		Price: 10,
	}
	assert.Error(t, order.Validate(), "tax must be greater than 0")
}

func Test_Final_Price(t *testing.T) {
	order := &Order{
		ID:    "1",
		Price: 10,
		Tax:   5,
	}
	assert.NoError(t, order.Validate())
	assert.Equal(t, "1", order.ID)
	assert.Equal(t, float64(10), order.Price)
	assert.Equal(t, float64(5), order.Tax)
	assert.NoError(t, order.CalculateFinalPrice())
	assert.Equal(t, float64(15), order.FinalPrice)
}
