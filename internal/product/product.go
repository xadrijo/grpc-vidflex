//go:generate mockgen -destination=product_mocks_test.go -package=product github.com/xadrijo/grpc-vidflex/internal/product Storage

package product

import "context"

// Product contain the definition of a product
type Product struct {
	ID     int32
	Label  string
	Type   string
	Url    string
	Weight float64
}

// Storage defines the interface we expect in our db.
type Storage interface {
	Get(id int32) (Product, error)
	Save(product Product) (Product, error)
}
// Service is responsible for get product
type Service struct {
	Storage Storage
}

// New returns a new instance of our product service.
func New(storage Storage) Service {
	return Service {
		Storage: storage,
	}
}

// GetProductByID retrieves a rocket based on the ID from db.
func (s Service) GetProductByID(ctx context.Context, id int32) (Product, error) {
	product, err := s.Storage.Get(id)
	if err != nil {
		return Product{}, err
	}
	return product, nil
}

func (s Service) InsertProduct(ctx context.Context, p Product) (Product, error) {
	product, err := s.Storage.Save(p)
	if err != nil {
		return Product{}, err
	}
	return product, nil
}
