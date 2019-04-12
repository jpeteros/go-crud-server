package models

import (
	"entities"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type ProductModel struct {
	Db *mgo.Database
}

/*test*/
func (productModel ProductModel) FindAll() ([]entities.Product, error) {
	var products []entities.Product
	err := productModel.Db.C("product").Find(bson.M{}).All(&products)
	if err != nil {
		return nil, err
	} else {
		return products, nil
	}
}

/*test*/
func (productModel ProductModel) FindAllUsers() ([]entities.User, error) {
	var users []entities.User
	err := productModel.Db.C("user").Find(bson.M{}).All(&users)
	if err != nil {
		return nil, err
	} else {
		return users, nil
	}
}

/*test*/
func (productModel ProductModel) FindProduct(id string) (entities.Product, error) {
	var product entities.Product
	err := productModel.Db.C("product").FindId(bson.ObjectIdHex(id)).One(&product)
	return product, err
}

/*test*/
func (productModel ProductModel) CreateProduct(product *entities.Product) error {
	return productModel.Db.C("product").Insert(&product)
}

/*test*/
func (productModel ProductModel) DeleteProduct(id string) error {
	return productModel.Db.C("product").RemoveId(bson.ObjectIdHex(id))
}

/*test*/
func (productModel ProductModel) UpdateProduct(product *entities.Product) error {
	return productModel.Db.C("product").UpdateId(product.Id, product)
}
