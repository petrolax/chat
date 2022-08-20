package helpers

import (
	"github.com/petrolax/chat/pkg/plants/dao"
	"github.com/petrolax/chat/pkg/plants/dto"
)

func ConvertSlicesDAOToDTO(src []*dao.Plant) []*dto.Plant {
	dst := make([]*dto.Plant, len(src))
	for i := range src {
		dst[i] = &dto.Plant{
			ID:   src[i].ID,
			Name: src[i].Name,
		}
	}
	return dst
}
