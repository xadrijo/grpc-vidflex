package product

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestProductService(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	t.Run("test get product by id", func(t *testing.T) {
		productStorageMock := NewMockStorage(mockCtrl)
		id := int32(1)
		productStorageMock.
			EXPECT().
			Get(id).
			Return(
				Product{
					ID:     1,
					Label:  "product-1",
					Type:   "Type-1",
					Url:    "",
					Weight: 3.35,
				},
				nil,
			)

		productService := New(productStorageMock)
		p, err := productService.GetProductByID(
			context.Background(),
			id,
		)

		assert.NoError(t, err)
		assert.Equal(t, int32(1), p.ID)
	})

	t.Run("test save product", func(t *testing.T) {
		productStorageMock := NewMockStorage(mockCtrl)
		product := Product{
			ID:     1,
			Label:  "product-1",
			Type:   "Type-1",
			Url:    "",
			Weight: 3.35,
		}

		productStorageMock.
			EXPECT().
			Save(product).
			Return(
			product,
				nil,
			)

		productService := New(productStorageMock)
		p, err := productService.InsertProduct(
			context.Background(),
			product,
		)

		assert.NoError(t, err)
		assert.EqualValues(t, product, p)
	})

	// TODO: Achieve 100% of coverage
}
