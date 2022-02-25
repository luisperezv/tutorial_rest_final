package repository

import (
	"rest.com/pkg/model"
)

func FetchGophers() ([]model.Gopher, error) {
	listGophers := []model.Gopher{
		{
			ID:    "1",
			Name:  "nombre1",
			Image: "image1",
			Age:   1,
		},
		{
			ID:    "2",
			Name:  "nombre2",
			Image: "image2",
			Age:   2,
		},
		{
			ID:    "3",
			Name:  "nombre3",
			Image: "image3",
			Age:   3,
		},
	}
	return listGophers, nil
}

func FetchGopherByID(id string) (model.Gopher, error) {
	someGopher := model.Gopher{
		ID:    "1",
		Name:  "nombre1",
		Image: "image1",
		Age:   1,
	}
	return someGopher, nil
}
