package postgres

import (
	"construccion_demo/internal/models"
	"fmt"
	"gorm.io/gorm"
	"log/slog"
)

type ProductService struct {
	DB *gorm.DB
}

// GetProducts todo: make tests
func (s *ProductService) GetProducts() *[]models.Product {
	var products []models.Product
	if err := s.DB.Raw("select product_name as name, c.category_name as category, s.company_name as supplier, " +
		"quantity_per_unit as unit, unit_price as price, units_in_stock as stock from products inner join public.categories c " +
		"on c.category_id = products.category_id inner join public.suppliers s on s.supplier_id = products.supplier_id").
		Scan(&products).Error; err != nil {
		slog.Error(err.Error())
		return nil
	}
	slog.Info(fmt.Sprintf("Consulted %d products from database", len(products)))
	return &products
}

// SearchProductsByName todo: make tests
func (s *ProductService) SearchProductsByName(name string) *[]models.Product {
	var products []models.Product
	if err := s.DB.Raw("select product_name as name, c.category_name as category, s.company_name as supplier, "+
		"quantity_per_unit as unit, unit_price as price, units_in_stock as stock from products inner join public.categories c "+
		"on c.category_id = products.category_id inner join public.suppliers s on s.supplier_id = products.supplier_id "+
		"where product_name LIKE ?", "%"+name+"%").
		Scan(&products).Error; err != nil {
		slog.Error(err.Error())
		return nil
	}
	slog.Info(fmt.Sprintf("Search products with %%%s%% in database", name))
	return &products
}
