package handler

import (
	"features/internal/domain/feature/service"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Feature struct {
	featureService service.Feature
}

func NewFeature(featureService service.Feature) Feature {
	return Feature{featureService: featureService}
}

func (f *Feature) FindByCustomerID(ctx echo.Context) error {
	var req IDRequest
	if err := ctx.Bind(req); err != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}
	res, err := f.featureService.FindByCustomerID(ctx.Request().Context(), req.ID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return ctx.JSON(http.StatusOK, mapFeaturesToCustomerFeatures(res))
}

func (f *Feature) GetAll(ctx echo.Context) error {
	res, err := f.featureService.GetAll(ctx.Request().Context())
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return ctx.JSON(http.StatusOK, mapModelsToResponse(res))
}

func (f *Feature) CreateFeature(ctx echo.Context) error {
	var req *featureRequestResponse
	if err := ctx.Bind(req); err != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}
	err := f.featureService.CreateFeature(ctx.Request().Context(), mapFeatureToModel(req))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return ctx.NoContent(http.StatusCreated)
}

func (f *Feature) UpdateFeature(ctx echo.Context) error {
	var req *featureRequestResponse
	if err := ctx.Bind(req); err != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}
	err := f.featureService.UpdateFeature(ctx.Request().Context(), mapFeatureToModel(req))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return ctx.NoContent(http.StatusCreated)
}

func (f *Feature) ArchiveFeature(ctx echo.Context) error {
	var req *IDRequest
	if err := ctx.Bind(req); err != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}
	err := f.featureService.ArchiveFeature(ctx.Request().Context(), req.ID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return ctx.NoContent(http.StatusCreated)
}

type featureResponse struct {
	Features []*featureRequestResponse `json:"features"`
}

type featureRequestResponse struct {
	ID            string   `json:"id"`
	DisplayName   string   `json:"display_name"`
	TechnicalName string   `json:"technical_name"`
	ExpiresOn     int64    `json:"expires_on"`
	Description   string   `json:"description"`
	Inverted      bool     `json:"inverted"`
	CustomerIDs   []string `json:"customer_ids"`
	CreatedAt     int64    `json:"created_at"`
	UpdatedAt     int64    `json:"updated_at"`
	DeletedAt     int64    `json:"deleted_at"`
}

type IDRequest struct {
	ID string `query:"id"`
}

type CustomerResponse struct {
	Features []*CustomerFeature `json:"features"`
}

type CustomerFeature struct {
	Name     string `json:"name"`
	Active   bool   `json:"active"`
	Inverted bool   `json:"inverted"`
	Expired  bool   `json:"expired"`
}
