package repositories

import (
	"database/sql"
	"product-manager/common"
	"product-manager/datamodels"
	"strconv"
)

// 1. interface
// 2. real struct

// IProduct interface
type IProduct interface {
	Conn() error
	Insert(*datamodels.Product) (int64, error)
	Delete(int64) bool
	Update(*datamodels.Product) error
	SelectByKey(int64) (*datamodels.Product, error)
	SelectAll() ([]*datamodels.Product, error)
}

// ProductManager real struct of IProduct
type ProductManager struct {
	table     string
	mysqlConn *sql.DB
}

// NewProductManager get ProductManager
func NewProductManager(table string, db *sql.DB) IProduct {
	return &ProductManager{
		table:     table,
		mysqlConn: db,
	}
}

// Conn database connection
func (p *ProductManager) Conn() (err error) {
	if p.mysqlConn == nil {
		mysql, err := common.NewMysqlConn()
		if err != nil {
			return err
		}
		p.mysqlConn = mysql
	}
	if p.table == "" {
		p.table = "product"
	}
	return
}

// Insert insert a new product
func (p *ProductManager) Insert(product *datamodels.Product) (productID int64, err error) {
	if err = p.Conn(); err != nil {
		return
	}

	sql := "INSERT product SET productName=?,productNum=?,productImage=?,productUrl=?"

	stmt, err := p.mysqlConn.Prepare(sql)
	if err != nil {
		productID = 0
		return
	}

	result, err := stmt.Exec(product.ProductName, product.ProductNum, product.ProductImage, product.ProductUrl)
	if err != nil {
		productID = 0
		return
	}

	return result.LastInsertId()
}

// Delete delete a product by key
func (p *ProductManager) Delete(productID int64) bool {
	if err := p.Conn(); err != nil {
		return false
	}

	sql := "delete from product where ID=?"
	stmt, err := p.mysqlConn.Prepare(sql)
	if err != nil {
		return false
	}

	_, err = stmt.Exec(strconv.FormatInt(productID, 10))
	if err != nil {
		return false
	}

	return true
}

// Update update a product
func (p *ProductManager) Update(product *datamodels.Product) (err error) {
	if err = p.Conn(); err != nil {
		return
	}

	sql := "Update product set productName=?,productNum=?,productImage=?,productUrl=? where ID=" + strconv.FormatInt(product.ID, 10)

	stmt, err := p.mysqlConn.Prepare(sql)
	if err != nil {
		return
	}

	_, err = stmt.Exec(product.ProductName, product.ProductNum, product.ProductImage, product.ProductUrl)
	if err != nil {
		return
	}
	return
}

// SelectByKey get a product by key
func (p *ProductManager) SelectByKey(productID int64) (productResult *datamodels.Product, err error) {
	//1.判断连接是否存在
	if err = p.Conn(); err != nil {
		return &datamodels.Product{}, err
	}

	sql := "Select * from " + p.table + " where ID =" + strconv.FormatInt(productID, 10)
	row, errRow := p.mysqlConn.Query(sql)
	defer row.Close()
	if errRow != nil {
		return &datamodels.Product{}, errRow
	}

	result := common.GetResultRow(row)
	if len(result) == 0 {
		return &datamodels.Product{}, nil
	}

	productResult = &datamodels.Product{}
	common.DataToStructByTagSQL(result, productResult)

	return
}

// SelectAll get all products
func (p *ProductManager) SelectAll() (productArray []*datamodels.Product, err error) {
	//1.判断连接是否存在
	if err = p.Conn(); err != nil {
		return
	}
	sql := "Select * from " + p.table
	rows, err := p.mysqlConn.Query(sql)
	defer rows.Close()
	if err != nil {
		return
	}

	result := common.GetResultRows(rows)
	if len(result) == 0 {
		return
	}

	for _, v := range result {
		product := &datamodels.Product{}
		common.DataToStructByTagSQL(v, product)
		productArray = append(productArray, product)
	}

	return
}
