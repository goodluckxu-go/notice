package notice

import (
	"github.com/gofrs/uuid"
	pb "github.com/goodluckxu-go/notice/code"
)

func getUUID() string {
	v4, _ := uuid.NewV4()
	return v4.String()
}

func inArray[T comparable](val T, list []T) bool {
	for _, v := range list {
		if v == val {
			return true
		}
	}
	return false
}

func toMetadata(v any) *pb.Metadata {
	switch val := v.(type) {
	case int:
		return &pb.Metadata{
			Value: &pb.Metadata_Int{Int: int64(val)},
		}
	case int8:
		return &pb.Metadata{
			Value: &pb.Metadata_Int{Int: int64(val)},
		}
	case int16:
		return &pb.Metadata{
			Value: &pb.Metadata_Int{Int: int64(val)},
		}
	case int32:
		return &pb.Metadata{
			Value: &pb.Metadata_Int{Int: int64(val)},
		}
	case int64:
		return &pb.Metadata{
			Value: &pb.Metadata_Int{Int: val},
		}
	case uint:
		return &pb.Metadata{
			Value: &pb.Metadata_Uint{Uint: uint64(val)},
		}
	case uint8:
		return &pb.Metadata{
			Value: &pb.Metadata_Uint{Uint: uint64(val)},
		}
	case uint16:
		return &pb.Metadata{
			Value: &pb.Metadata_Uint{Uint: uint64(val)},
		}
	case uint32:
		return &pb.Metadata{
			Value: &pb.Metadata_Uint{Uint: uint64(val)},
		}
	case uint64:
		return &pb.Metadata{
			Value: &pb.Metadata_Uint{Uint: val},
		}
	case float32:
		return &pb.Metadata{
			Value: &pb.Metadata_Float{Float: float64(val)},
		}
	case float64:
		return &pb.Metadata{
			Value: &pb.Metadata_Float{Float: val},
		}
	case string:
		return &pb.Metadata{
			Value: &pb.Metadata_String_{String_: val},
		}
	case bool:
		return &pb.Metadata{
			Value: &pb.Metadata_Bool{Bool: val},
		}
	}
	return nil
}
