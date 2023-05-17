package utils

import (
	"math"

	"gorm.io/gorm"
)

func Paginate(value any, pagination *GormPaginationData, db *gorm.DB) func(db *gorm.DB) *gorm.DB {
	var totalRows int64

	db.Model(value).Count(&totalRows)

	pagination.TotalRows = totalRows

	totalPages := int(math.Ceil(float64(totalRows) / float64(pagination.GetLimit())))

	pagination.TotalPages = totalPages

	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(pagination.GetOffset()).Limit(pagination.GetLimit()).Order(pagination.GetSort())
	}
}

func Contains[T comparable](ss []T, v T) bool {
	for _, s := range ss {
		if s == v {
			return true
		}
	}

	return false
}
