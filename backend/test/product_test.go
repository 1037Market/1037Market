package test

import (
	"1037Market/dao"
	"1037Market/ds"
	"strconv"
	"testing"
)

func makeProduct() ds.ProductPublished {
	uris := make([]string, 0)
	uris = append(uris, "1111")
	uris = append(uris, "longonglonglonglonglong")

	categories := make([]string, 0)
	categories = append(categories, "category")

	product := ds.ProductPublished{
		Title:      "long long long title",
		Content:    "long long long content",
		Categories: categories,
		ImageURIs:  uris,
		Price:      123.456,
	}
	return product
}

func TestPublishProduct(t *testing.T) {

	product := makeProduct()

	_, err := dao.PublishProduct("t1", product)
	if err != nil {
		t.Error(err)
	}
}

func TestUpdateProduct(t *testing.T) {
	product := makeProduct()
	productId, err := dao.PublishProduct("t1", product)
	if err != nil {
		t.Error(err)
	}

	uri := make([]string, 0)
	uri = append(uri, "uri")

	update := ds.ProductUpdated{
		ProductId:  productId,
		Title:      "11223344",
		Content:    "112",
		Categories: uri,
		ImageURIs:  uri,
		Price:      -111,
		IsSoldOut:  true,
	}

	err = dao.UpdateProduct("t1", update)
	if err != nil {
		t.Error(err)
	}

	got, err := dao.GetProductById("t1", strconv.Itoa(productId))
	if err != nil {
		t.Error(err)
	}

	if got.IsSoldOut != true || got.Price != -111 || len(got.ImageURIs) != 1 || got.Title != "11223344" {
		t.Error("update error, expect: ", update, "got", got)
	}
}

func TestUpdateProductInvalidCategory(t *testing.T) {
	product := makeProduct()

	productId, err := dao.PublishProduct("t1", product)
	if err != nil {
		t.Error(err)
	}

	uri := make([]string, 0)
	uri = append(uri, "uri")

	category := make([]string, 0)
	category = append(category, "not in categoryies")

	update := ds.ProductUpdated{
		ProductId:  productId,
		Title:      "11223344",
		Content:    "112",
		Categories: category,
		ImageURIs:  uri,
		Price:      -111,
		IsSoldOut:  true,
	}

	err = dao.UpdateProduct("t1", update)
	if err == nil {
		t.Error("expect to fail to update, but succeeded")
	}

}

func TestUpdateProductInvalidStudentId(t *testing.T) {
	product := makeProduct()

	productId, err := dao.PublishProduct("t1", product)
	if err != nil {
		t.Error(err)
	}

	uri := make([]string, 0)
	uri = append(uri, "uri")

	update := ds.ProductUpdated{
		ProductId:  productId,
		Title:      "11223344",
		Content:    "112",
		Categories: uri,
		ImageURIs:  uri,
		Price:      -111,
		IsSoldOut:  true,
	}

	err = dao.UpdateProduct("not exist", update)
	if err == nil {
		t.Error("expect to fail to update, but succeeded")
	}

}
