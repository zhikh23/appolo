package materialgrpc

import (
	"appolo-register/internal/domain"
	"appolo-register/internal/services/materials"
	"context"
	"errors"

	appolov1 "github.com/zhikh23/appolo-protos/gen/go/register"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type MaterialService interface {
	Register(
		ctx context.Context,
		m   *domain.Material,
	) (id uint64, err error)
	MaterialById(
		ctx context.Context,
		id  uint64,
	) (m *domain.Material, err error)
	MaterialsByTags(
		ctx  context.Context,
		tags []string,
	) (ms []*domain.Material, err error)
}

type serverApi struct {
	appolov1.UnimplementedRegisterServiceServer
	service MaterialService
}

func Register(grpcServer *grpc.Server, service MaterialService) {
	appolov1.RegisterRegisterServiceServer(grpcServer, &serverApi{ service: service })
}

func (s *serverApi) RegisterMaterial(ctx context.Context, req *appolov1.RegisterRequest) (*appolov1.RegisterResponse, error) {
	id, err := s.service.Register(ctx, &domain.Material{
		Name: 		 req.Name,
		Description: req.Description,
		Tags: 		 req.Tags,
		Url: 		 req.Url,
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &appolov1.RegisterResponse{
		MaterialId: id,
	}, nil
}

func (s *serverApi) GetMaterialById(ctx context.Context, req *appolov1.GetMaterialByIdRequest) (*appolov1.GetMaterialByIdResponse, error) {
	model, err := s.service.MaterialById(ctx, req.MaterialId)
	if err != nil {
		if errors.Is(err, materials.ErrMaterialNotFound) {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &appolov1.GetMaterialByIdResponse{
		Material: convertModelToDto(model),
	}, nil
}

func (s *serverApi) GetMaterialsByTags(ctx context.Context, req *appolov1.GetMaterialsByTagsRequest) (*appolov1.GetMaterialsByTagsResponse, error) {
	models, err := s.service.MaterialsByTags(ctx, req.Tags)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &appolov1.GetMaterialsByTagsResponse{
		Materials: convertModelsToDtos(models),
	}, nil
}

func convertModelToDto(model *domain.Material) *appolov1.Material {
	return &appolov1.Material {
		Id: 	     model.Id,
		Name:		 model.Name,
		Description: model.Description,
		Tags:		 model.Tags,
		Url:		 model.Url,
	}
}

func convertModelsToDtos(models []*domain.Material) []*appolov1.Material {
	dtos := make([]*appolov1.Material, len(models))
	for i, model := range models {
		dtos[i] = convertModelToDto(model)
	}
	return dtos
}
