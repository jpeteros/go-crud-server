package productapi

import (
	"config"
	"encoding/json"
	"entities"
	"models"
	"net/http"

	"gopkg.in/mgo.v2/bson"

	"github.com/gorilla/mux"
)

/*test*/
func FindAll(response http.ResponseWriter, request *http.Request) {
	db, err := config.Connect()
	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
		return
	} else {
		productModel := models.ProductModel{
			Db: db,
		}
		products, err2 := productModel.FindAll()
		if err2 != nil {
			respondWithError(response, http.StatusBadRequest, err2.Error())
			return
		} else {
			respondWithJson(response, http.StatusOK, products)
		}
	}
}

func FindAllUsers(response http.ResponseWriter, request *http.Request) {
	db, err := config.Connect()
	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
		return
	} else {
		productModel := models.ProductModel{
			Db: db,
		}
		users, err2 := productModel.FindAllUsers()
		if err2 != nil {
			respondWithError(response, http.StatusBadRequest, err2.Error())
			return
		} else {
			respondWithJson(response, http.StatusOK, users)
		}
	}
}

/*test*/
func FindProduct(response http.ResponseWriter, request *http.Request) {
	db, err := config.Connect()
	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
		return
	} else {
		productModel := models.ProductModel{
			Db: db,
		}
		vars := mux.Vars(request)
		id := vars["id"]
		product, err2 := productModel.FindProduct(id)
		if err2 != nil {
			respondWithError(response, http.StatusBadRequest, err2.Error())
			return
		} else {
			respondWithJson(response, http.StatusOK, product)
		}
	}
}

/*test*/
func CreateProduct(response http.ResponseWriter, request *http.Request) {
	db, err := config.Connect()
	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
		return
	} else {
		productModel := models.ProductModel{
			Db: db,
		}
		var product entities.Product
		product.Id = bson.NewObjectId()
		json.NewDecoder(request.Body).Decode(&product)
		err2 := productModel.CreateProduct(&product)
		if err2 != nil {
			respondWithError(response, http.StatusBadRequest, err2.Error())
		} else {
			respondWithJson(response, http.StatusOK, product)
		}
	}
}

/*test*/
func DeleteProduct(response http.ResponseWriter, request *http.Request) {
	db, err := config.Connect()
	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
		return
	} else {
		productModel := models.ProductModel{
			Db: db,
		}
		vars := mux.Vars(request)
		id := vars["id"]
		err2 := productModel.DeleteProduct(id)
		if err2 != nil {
			respondWithError(response, http.StatusBadRequest, err2.Error())
		} else {
			respondWithJson(response, http.StatusOK, entities.Product{})
		}
	}
}

/*test*/
func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

/*test*/
func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
