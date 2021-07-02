package image

import "gocatering/model"

type Repository interface {
	CreateImage(img *model.Image) error
}
