package service

import "shipping/model"

type PathRouter interface {

	GetRoute(src string, dst string) (model.Route, error )
}