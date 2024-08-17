package meta

import (
	"golang-prisma-api/src/interfaces"
	"math"
)

func Meta(page, perPage, totalData int) interfaces.MetaResponse {
	totalPages := int(math.Ceil(float64(totalData) / float64(perPage)))

	return interfaces.MetaResponse{
		Page:       page,
		PerPage:    perPage,
		TotalData:  totalData,
		TotalPages: totalPages,
	}
}
