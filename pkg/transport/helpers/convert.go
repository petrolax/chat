package helpers

import (
	"github.com/petrolax/chat/pkg/api/private"
	"github.com/petrolax/chat/pkg/api/public"
	"github.com/petrolax/chat/pkg/plants/dto"
)

func ConvertDTOToPulbicPlant(src *dto.Plant) *public.Plant {
	return &public.Plant{
		Id:   uint32(src.ID),
		Name: src.Name,
	}
}

func ConvertSlicesDTOToPublicPlant(src []*dto.Plant) []*public.Plant {
	dst := make([]*public.Plant, len(src))
	for i := range src {
		dst[i] = &public.Plant{
			Id:   uint32(src[i].ID),
			Name: src[i].Name,
		}
	}
	return dst
}

func ConvertSlicesDTOToPrivatePlant(src []*dto.Plant) []*private.Plant {
	dst := make([]*private.Plant, len(src))
	for i := range src {
		dst[i] = &private.Plant{
			Id:   uint32(src[i].ID),
			Name: src[i].Name,
		}
	}
	return dst
}
