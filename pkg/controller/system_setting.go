package controller

import (
	"github.com/KubeOperator/KubeOperator/pkg/constant"
	"github.com/KubeOperator/KubeOperator/pkg/controller/kolog"
	"github.com/KubeOperator/KubeOperator/pkg/dto"
	"github.com/KubeOperator/KubeOperator/pkg/service"
	"github.com/go-playground/validator/v10"
	"github.com/kataras/iris/v12/context"
)

type SystemSettingController struct {
	Ctx                  context.Context
	SystemSettingService service.SystemSettingService
}

func NewSystemSettingController() *SystemSettingController {
	return &SystemSettingController{
		SystemSettingService: service.NewSystemSettingService(),
	}
}

func (s SystemSettingController) Get() (interface{}, error) {

	item, err := s.SystemSettingService.List()
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (s SystemSettingController) GetBy(name string) (interface{}, error) {
	item, err := s.SystemSettingService.ListByTab(name)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (s SystemSettingController) Post() ([]dto.SystemSetting, error) {
	var req dto.SystemSettingCreate
	err := s.Ctx.ReadJSON(&req)
	if err != nil {
		return nil, err
	}
	validate := validator.New()
	err = validate.Struct(req)
	if err != nil {
		return nil, err
	}
	result, err := s.SystemSettingService.Create(req)
	if err != nil {
		return nil, err
	}

	operator := s.Ctx.Values().GetString("operator")
	go kolog.Save(operator, constant.CREATE_EMAIL, "-")

	return result, nil
}

func (s SystemSettingController) GetIp() string {
	return s.SystemSettingService.GetLocalHostName()
}

func (s SystemSettingController) PostCheckBy(typeName string) error {
	var req dto.SystemSettingCreate
	err := s.Ctx.ReadJSON(&req)
	if err != nil {
		return err
	}
	validate := validator.New()
	err = validate.Struct(req)
	if err != nil {
		return err
	}
	err = s.SystemSettingService.CheckSettingByType(typeName, req)
	if err != nil {
		return err
	}
	return nil
}
