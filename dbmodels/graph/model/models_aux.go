package model

// TableName overrides the table name used by User to `profiles`
func (ProductAdmin) TableName() string {
	return "products"
}
