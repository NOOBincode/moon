package builder

import (
	"context"

	"github.com/aide-family/moon/api"
	adminapi "github.com/aide-family/moon/api/admin"
	strategyapi "github.com/aide-family/moon/api/admin/strategy"
	"github.com/aide-family/moon/api/merr"
	"github.com/aide-family/moon/cmd/server/palace/internal/biz/bo"
	"github.com/aide-family/moon/pkg/helper/middleware"
	"github.com/aide-family/moon/pkg/palace/imodel"
	"github.com/aide-family/moon/pkg/palace/model"
	"github.com/aide-family/moon/pkg/palace/model/bizmodel"
	"github.com/aide-family/moon/pkg/util/types"
	"github.com/aide-family/moon/pkg/vobj"
)

var _ IStrategyModuleBuilder = (*strategyModuleBuilder)(nil)

type (
	strategyModuleBuilder struct {
		ctx context.Context
	}

	IStrategyModuleBuilder interface {
		WithCreateStrategyGroupRequest(*strategyapi.CreateStrategyGroupRequest) ICreateStrategyGroupRequestBuilder
		WithDeleteStrategyGroupRequest(*strategyapi.DeleteStrategyGroupRequest) IDeleteStrategyGroupRequestBuilder
		WithListStrategyGroupRequest(*strategyapi.ListStrategyGroupRequest) IListStrategyGroupRequestBuilder
		WithUpdateStrategyGroupRequest(*strategyapi.UpdateStrategyGroupRequest) IUpdateStrategyGroupRequestBuilder
		WithUpdateStrategyGroupStatusRequest(*strategyapi.UpdateStrategyGroupStatusRequest) IUpdateStrategyGroupStatusRequestBuilder
		DoStrategyGroupBuilder() IDoStrategyGroupBuilder

		WithCreateStrategyRequest(*strategyapi.CreateStrategyRequest) ICreateStrategyRequestBuilder
		WithUpdateStrategyRequest(*strategyapi.UpdateStrategyRequest) IUpdateStrategyRequestBuilder
		WithListStrategyRequest(*strategyapi.ListStrategyRequest) IListStrategyRequestBuilder
		WithUpdateStrategyStatusRequest(*strategyapi.UpdateStrategyStatusRequest) IUpdateStrategyStatusRequestBuilder
		DoStrategyBuilder() IDoStrategyBuilder
		DoStrategyLevelBuilder() IDoStrategyLevelBuilder

		WithCreateTemplateStrategyRequest(*strategyapi.CreateTemplateStrategyRequest) ICreateTemplateStrategyRequestBuilder
		WithUpdateTemplateStrategyRequest(*strategyapi.UpdateTemplateStrategyRequest) IUpdateTemplateStrategyRequestBuilder
		WithListTemplateStrategyRequest(*strategyapi.ListTemplateStrategyRequest) IListTemplateStrategyRequestBuilder
		WithUpdateTemplateStrategyStatusRequest(*strategyapi.UpdateTemplateStrategyStatusRequest) IUpdateTemplateStrategyStatusRequestBuilder
		DoTemplateStrategyBuilder() IDoTemplateStrategyBuilder

		APIMutationStrategyLevelTemplateItems() IMutationStrategyLevelTemplateBuilder
		APIMutationStrategyLevelItems() IMutationStrategyLevelBuilder
		DoStrategyLevelTemplateBuilder() IDoStrategyLevelTemplateBuilder

		BoStrategyBuilder() IBoStrategyBuilder
	}

	ICreateStrategyGroupRequestBuilder interface {
		ToBo() *bo.CreateStrategyGroupParams
	}

	createStrategyGroupRequestBuilder struct {
		ctx context.Context
		*strategyapi.CreateStrategyGroupRequest
	}

	IDeleteStrategyGroupRequestBuilder interface {
		ToBo() *bo.DelStrategyGroupParams
	}

	deleteStrategyGroupRequestBuilder struct {
		ctx context.Context
		*strategyapi.DeleteStrategyGroupRequest
	}

	IListStrategyGroupRequestBuilder interface {
		ToBo() *bo.QueryStrategyGroupListParams
	}

	listStrategyGroupRequestBuilder struct {
		ctx context.Context
		*strategyapi.ListStrategyGroupRequest
	}

	IUpdateStrategyGroupRequestBuilder interface {
		ToBo() *bo.UpdateStrategyGroupParams
	}

	updateStrategyGroupRequestBuilder struct {
		ctx context.Context
		*strategyapi.UpdateStrategyGroupRequest
	}

	IUpdateStrategyGroupStatusRequestBuilder interface {
		ToBo() *bo.UpdateStrategyGroupStatusParams
	}

	updateStrategyGroupStatusRequestBuilder struct {
		ctx context.Context
		*strategyapi.UpdateStrategyGroupStatusRequest
	}

	ICreateStrategyRequestBuilder interface {
		ToBo() *bo.CreateStrategyParams
	}

	createStrategyRequestBuilder struct {
		ctx context.Context
		*strategyapi.CreateStrategyRequest
	}

	IUpdateStrategyRequestBuilder interface {
		ToBo() *bo.UpdateStrategyParams
	}

	updateStrategyRequestBuilder struct {
		ctx context.Context
		*strategyapi.UpdateStrategyRequest
	}

	IListStrategyRequestBuilder interface {
		ToBo() *bo.QueryStrategyListParams
	}

	listStrategyRequestBuilder struct {
		ctx context.Context
		*strategyapi.ListStrategyRequest
	}

	IUpdateStrategyStatusRequestBuilder interface {
		ToBo() *bo.UpdateStrategyStatusParams
	}

	updateStrategyStatusRequestBuilder struct {
		ctx context.Context
		*strategyapi.UpdateStrategyStatusRequest
	}

	ICreateTemplateStrategyRequestBuilder interface {
		ToBo() *bo.CreateTemplateStrategyParams
	}

	createTemplateStrategyRequestBuilder struct {
		ctx context.Context
		*strategyapi.CreateTemplateStrategyRequest
	}

	IUpdateTemplateStrategyRequestBuilder interface {
		ToBo() *bo.UpdateTemplateStrategyParams
	}

	updateTemplateStrategyRequestBuilder struct {
		ctx context.Context
		*strategyapi.UpdateTemplateStrategyRequest
	}

	IListTemplateStrategyRequestBuilder interface {
		ToBo() *bo.QueryTemplateStrategyListParams
	}

	listTemplateStrategyRequestBuilder struct {
		ctx context.Context
		*strategyapi.ListTemplateStrategyRequest
	}

	IUpdateTemplateStrategyStatusRequestBuilder interface {
		ToBo() *bo.UpdateTemplateStrategyStatusParams
	}

	updateTemplateStrategyStatusRequestBuilder struct {
		ctx context.Context
		*strategyapi.UpdateTemplateStrategyStatusRequest
	}

	IDoStrategyGroupBuilder interface {
		WithStrategyCountMap(*bo.StrategyCountMap) IDoStrategyGroupBuilder
		ToAPI(*bizmodel.StrategyGroup) *adminapi.StrategyGroupItem
		ToAPIs([]*bizmodel.StrategyGroup) []*adminapi.StrategyGroupItem
		ToSelect(*bizmodel.StrategyGroup) *adminapi.SelectItem
		ToSelects([]*bizmodel.StrategyGroup) []*adminapi.SelectItem
	}

	doStrategyGroupBuilder struct {
		ctx              context.Context
		strategyCountMap *bo.StrategyCountMap
	}

	IDoStrategyBuilder interface {
		ToAPI(*bizmodel.Strategy) *adminapi.StrategyItem
		ToAPIs([]*bizmodel.Strategy) []*adminapi.StrategyItem
		ToSelect(*bizmodel.Strategy) *adminapi.SelectItem
		ToSelects([]*bizmodel.Strategy) []*adminapi.SelectItem
		ToBos(*bizmodel.Strategy) []*bo.Strategy
	}

	doStrategyBuilder struct {
		ctx context.Context
	}

	IDoTemplateStrategyBuilder interface {
		ToAPI(*model.StrategyTemplate) *adminapi.StrategyTemplateItem
		ToAPIs([]*model.StrategyTemplate) []*adminapi.StrategyTemplateItem
		ToSelect(*model.StrategyTemplate) *adminapi.SelectItem
		ToSelects([]*model.StrategyTemplate) []*adminapi.SelectItem
	}

	doTemplateStrategyBuilder struct {
		ctx context.Context
	}

	IDoStrategyLevelTemplateBuilder interface {
		ToAPI(*model.StrategyLevelTemplate) *adminapi.StrategyLevelTemplateItem
		ToAPIs([]*model.StrategyLevelTemplate) []*adminapi.StrategyLevelTemplateItem
	}

	doStrategyLevelTemplateBuilder struct {
		ctx context.Context
	}

	IMutationStrategyLevelTemplateBuilder interface {
		WithStrategyTemplateID(uint32) IMutationStrategyLevelTemplateBuilder
		ToBo(*strategyapi.MutationStrategyLevelTemplateItem) *bo.CreateStrategyLevelTemplate
		ToBos([]*strategyapi.MutationStrategyLevelTemplateItem) []*bo.CreateStrategyLevelTemplate
	}

	mutationStrategyLevelTemplateBuilder struct {
		ctx                context.Context
		StrategyTemplateID uint32
	}

	IMutationStrategyLevelBuilder interface {
		WithStrategyID(uint32) IMutationStrategyLevelBuilder
		ToBo(*strategyapi.CreateStrategyLevelRequest) *bo.CreateStrategyLevel
		ToBos([]*strategyapi.CreateStrategyLevelRequest) []*bo.CreateStrategyLevel
	}

	mutationStrategyLevelBuilder struct {
		ctx        context.Context
		StrategyID uint32
	}

	IDoStrategyLevelBuilder interface {
		ToAPI(*bizmodel.StrategyLevel) *adminapi.StrategyLevelItem
		ToAPIs([]*bizmodel.StrategyLevel) []*adminapi.StrategyLevelItem
	}

	doStrategyLevelBuilder struct {
		ctx context.Context
	}

	IBoStrategyBuilder interface {
		ToAPI(*bo.Strategy) *api.Strategy
		ToAPIs([]*bo.Strategy) []*api.Strategy
	}

	boStrategyBuilder struct {
		ctx context.Context
	}
)

func (d *doStrategyBuilder) ToBos(strategy *bizmodel.Strategy) []*bo.Strategy {
	if types.IsNil(strategy) || types.IsNil(d) {
		return nil
	}

	return types.SliceTo(strategy.Levels, func(level *bizmodel.StrategyLevel) *bo.Strategy {
		return &bo.Strategy{
			ID:                         strategy.ID,
			LevelID:                    level.ID,
			Alert:                      strategy.Name,
			Expr:                       strategy.Expr,
			For:                        level.Duration,
			Count:                      level.Count,
			SustainType:                level.SustainType,
			MultiDatasourceSustainType: 0, // TODO 多数据源控制
			Labels:                     strategy.Labels,
			Annotations:                strategy.Annotations,
			Interval:                   level.Interval,
			Datasource:                 NewParamsBuild().WithContext(d.ctx).DatasourceModuleBuilder().DoDatasourceBuilder().ToBos(strategy.Datasource),
			Status:                     strategy.Status,
			Step:                       strategy.Step,
			Condition:                  level.Condition,
			Threshold:                  level.Threshold,
		}
	})
}

func (b *boStrategyBuilder) ToAPI(strategyItem *bo.Strategy) *api.Strategy {
	if types.IsNil(strategyItem) || types.IsNil(b) {
		return nil
	}

	return &api.Strategy{
		Alert:                      strategyItem.Alert,
		Expr:                       strategyItem.Expr,
		For:                        strategyItem.For.GetDuration(),
		Count:                      strategyItem.Count,
		SustainType:                api.SustainType(strategyItem.SustainType),
		MultiDatasourceSustainType: api.MultiDatasourceSustainType(strategyItem.MultiDatasourceSustainType),
		Labels:                     strategyItem.Labels.Map(),
		Annotations:                strategyItem.Annotations,
		Interval:                   strategyItem.Interval.GetDuration(),
		Datasource:                 NewParamsBuild().WithContext(b.ctx).DatasourceModuleBuilder().BoDatasourceBuilder().ToAPIs(strategyItem.Datasource),
		Id:                         strategyItem.ID,
		Status:                     api.Status(strategyItem.Status),
		Step:                       strategyItem.Step,
		Condition:                  api.Condition(strategyItem.Condition),
		Threshold:                  strategyItem.Threshold,
		LevelID:                    strategyItem.LevelID,
	}
}

func (b *boStrategyBuilder) ToAPIs(strategies []*bo.Strategy) []*api.Strategy {
	if types.IsNil(strategies) || types.IsNil(b) {
		return nil
	}

	return types.SliceTo(strategies, func(item *bo.Strategy) *api.Strategy {
		return b.ToAPI(item)
	})
}

func (s *strategyModuleBuilder) BoStrategyBuilder() IBoStrategyBuilder {
	return &boStrategyBuilder{ctx: s.ctx}
}

func (m *mutationStrategyLevelBuilder) WithStrategyID(u uint32) IMutationStrategyLevelBuilder {
	if !types.IsNil(m) {
		m.StrategyID = u
	}

	return m
}

func (m *mutationStrategyLevelBuilder) ToBo(request *strategyapi.CreateStrategyLevelRequest) *bo.CreateStrategyLevel {
	if types.IsNil(m) || types.IsNil(request) {
		return nil
	}

	return &bo.CreateStrategyLevel{
		StrategyTemplateID: m.StrategyID,
		Duration:           types.NewDuration(request.GetDuration()),
		Count:              request.Count,
		SustainType:        vobj.Sustain(request.SustainType),
		Interval:           types.NewDuration(request.GetInterval()),
		Condition:          vobj.Condition(request.Condition),
		Threshold:          request.Threshold,
		LevelID:            request.LevelId,
		Status:             vobj.Status(request.Status),
		AlarmPageIds:       request.GetAlarmPageIds(),
		AlarmGroupIds:      request.GetAlarmGroupIds(),
		StrategyID:         m.StrategyID,
		LabelNotices:       NewParamsBuild().WithContext(m.ctx).AlarmNoticeGroupModuleBuilder().APICreateStrategyLabelNoticeRequest().ToBos(request.GetLabelNotices()),
	}
}

func (m *mutationStrategyLevelBuilder) ToBos(requests []*strategyapi.CreateStrategyLevelRequest) []*bo.CreateStrategyLevel {
	if types.IsNil(m) || types.IsNil(requests) {
		return nil
	}

	return types.SliceTo(requests, func(request *strategyapi.CreateStrategyLevelRequest) *bo.CreateStrategyLevel {
		return m.ToBo(request)
	})
}

func (s *strategyModuleBuilder) APIMutationStrategyLevelItems() IMutationStrategyLevelBuilder {
	return &mutationStrategyLevelBuilder{ctx: s.ctx}
}

func (d *doStrategyLevelBuilder) ToAPI(level *bizmodel.StrategyLevel) *adminapi.StrategyLevelItem {
	if types.IsNil(d) || types.IsNil(level) {
		return nil
	}

	return &adminapi.StrategyLevelItem{
		Duration:     level.Duration.GetDuration(),
		Count:        level.Count,
		SustainType:  api.SustainType(level.SustainType),
		Interval:     level.Interval.GetDuration(),
		Status:       api.Status(level.Status),
		Id:           level.ID,
		LevelId:      level.LevelID,
		Level:        NewParamsBuild().WithContext(d.ctx).DictModuleBuilder().DoDictBuilder().ToSelect(level.Level),
		AlarmPages:   NewParamsBuild().WithContext(d.ctx).DictModuleBuilder().DoDictBuilder().ToSelects(types.SliceTo(level.AlarmPage, func(item *bizmodel.SysDict) imodel.IDict { return item })),
		Threshold:    level.Threshold,
		StrategyId:   level.StrategyID,
		Condition:    api.Condition(level.Condition),
		AlarmGroups:  NewParamsBuild().WithContext(d.ctx).AlarmNoticeGroupModuleBuilder().DoAlarmNoticeGroupItemBuilder().ToAPIs(level.AlarmGroups),
		LabelNotices: NewParamsBuild().WithContext(d.ctx).AlarmNoticeGroupModuleBuilder().DoLabelNoticeBuilder().ToAPIs(level.LabelNotices),
	}
}

func (d *doStrategyLevelBuilder) ToAPIs(levels []*bizmodel.StrategyLevel) []*adminapi.StrategyLevelItem {
	if types.IsNil(d) || types.IsNil(levels) {
		return nil
	}

	return types.SliceTo(levels, func(level *bizmodel.StrategyLevel) *adminapi.StrategyLevelItem {
		return d.ToAPI(level)
	})
}

func (s *strategyModuleBuilder) DoStrategyLevelBuilder() IDoStrategyLevelBuilder {
	return &doStrategyLevelBuilder{ctx: s.ctx}
}

func (m *mutationStrategyLevelTemplateBuilder) ToBos(items []*strategyapi.MutationStrategyLevelTemplateItem) []*bo.CreateStrategyLevelTemplate {
	if types.IsNil(m) || types.IsNil(items) {
		return nil
	}

	return types.SliceTo(items, func(item *strategyapi.MutationStrategyLevelTemplateItem) *bo.CreateStrategyLevelTemplate {
		return m.ToBo(item)
	})
}

func (m *mutationStrategyLevelTemplateBuilder) WithStrategyTemplateID(u uint32) IMutationStrategyLevelTemplateBuilder {
	if !types.IsNil(m) {
		m.StrategyTemplateID = u
	}
	return m
}

func (m *mutationStrategyLevelTemplateBuilder) ToBo(item *strategyapi.MutationStrategyLevelTemplateItem) *bo.CreateStrategyLevelTemplate {
	if types.IsNil(m) || types.IsNil(item) {
		return nil
	}

	return &bo.CreateStrategyLevelTemplate{
		StrategyTemplateID: m.StrategyTemplateID,
		Duration:           types.NewDuration(item.GetDuration()),
		Count:              item.GetCount(),
		SustainType:        vobj.Sustain(item.GetSustainType()),
		Condition:          vobj.Condition(item.Condition),
		Threshold:          item.Threshold,
		LevelID:            item.GetLevelId(),
		Status:             vobj.StatusEnable,
	}
}

func (s *strategyModuleBuilder) APIMutationStrategyLevelTemplateItems() IMutationStrategyLevelTemplateBuilder {
	return &mutationStrategyLevelTemplateBuilder{ctx: s.ctx}
}

func (d *doStrategyLevelTemplateBuilder) ToAPI(template *model.StrategyLevelTemplate) *adminapi.StrategyLevelTemplateItem {
	if types.IsNil(d) || types.IsNil(template) {
		return nil
	}

	return &adminapi.StrategyLevelTemplateItem{
		Id:          template.ID,
		Duration:    template.Duration.GetDuration(),
		Count:       template.Count,
		SustainType: api.SustainType(template.SustainType),
		Status:      api.Status(template.Status),
		LevelId:     template.LevelID,
		Level:       NewParamsBuild().WithContext(d.ctx).DictModuleBuilder().DoDictBuilder().ToSelect(template.Level),
		Threshold:   template.Threshold,
		Condition:   api.Condition(template.Condition),
		StrategyId:  template.StrategyTemplateID,
	}
}

func (d *doStrategyLevelTemplateBuilder) ToAPIs(templates []*model.StrategyLevelTemplate) []*adminapi.StrategyLevelTemplateItem {
	if types.IsNil(d) || types.IsNil(templates) {
		return nil
	}

	return types.SliceTo(templates, func(item *model.StrategyLevelTemplate) *adminapi.StrategyLevelTemplateItem {
		return d.ToAPI(item)
	})
}

func (s *strategyModuleBuilder) DoStrategyLevelTemplateBuilder() IDoStrategyLevelTemplateBuilder {
	return &doStrategyLevelTemplateBuilder{ctx: s.ctx}
}

func (d *doTemplateStrategyBuilder) ToAPI(template *model.StrategyTemplate) *adminapi.StrategyTemplateItem {
	if types.IsNil(d) || types.IsNil(template) {
		return nil
	}

	return &adminapi.StrategyTemplateItem{
		Id:          template.ID,
		Alert:       template.Alert,
		Expr:        template.Expr,
		Levels:      NewParamsBuild().WithContext(d.ctx).StrategyModuleBuilder().DoStrategyLevelTemplateBuilder().ToAPIs(template.StrategyLevelTemplates),
		Labels:      template.Labels.Map(),
		Annotations: template.Annotations,
		Status:      api.Status(template.Status),
		CreatedAt:   template.CreatedAt.String(),
		UpdatedAt:   template.UpdatedAt.String(),
		Remark:      template.Remark,
		Creator:     nil, // TODO 创建人绑定
		Categories:  NewParamsBuild().WithContext(d.ctx).DictModuleBuilder().DoDictBuilder().ToSelects(types.SliceTo(template.Categories, func(item *model.SysDict) imodel.IDict { return item })),
	}
}

func (d *doTemplateStrategyBuilder) ToAPIs(templates []*model.StrategyTemplate) []*adminapi.StrategyTemplateItem {
	if types.IsNil(d) || types.IsNil(templates) {
		return nil
	}

	return types.SliceTo(templates, func(item *model.StrategyTemplate) *adminapi.StrategyTemplateItem {
		return d.ToAPI(item)
	})
}

func (d *doTemplateStrategyBuilder) ToSelect(template *model.StrategyTemplate) *adminapi.SelectItem {
	if types.IsNil(d) || types.IsNil(template) {
		return nil
	}

	return &adminapi.SelectItem{
		Value:    template.ID,
		Label:    template.Alert,
		Children: nil,
		Disabled: template.DeletedAt > 0 || !template.Status.IsEnable(),
		Extend:   &adminapi.SelectExtend{Remark: template.Remark},
	}
}

func (d *doTemplateStrategyBuilder) ToSelects(templates []*model.StrategyTemplate) []*adminapi.SelectItem {
	if types.IsNil(d) || types.IsNil(templates) {
		return nil
	}

	return types.SliceTo(templates, func(item *model.StrategyTemplate) *adminapi.SelectItem {
		return d.ToSelect(item)
	})
}

func (u *updateTemplateStrategyStatusRequestBuilder) ToBo() *bo.UpdateTemplateStrategyStatusParams {
	if types.IsNil(u) || types.IsNil(u.UpdateTemplateStrategyStatusRequest) {
		return nil
	}

	return &bo.UpdateTemplateStrategyStatusParams{
		IDs:    u.GetIds(),
		Status: vobj.Status(u.GetStatus()),
	}
}

func (l *listTemplateStrategyRequestBuilder) ToBo() *bo.QueryTemplateStrategyListParams {
	if types.IsNil(l) || types.IsNil(l.ListTemplateStrategyRequest) {
		return nil
	}

	return &bo.QueryTemplateStrategyListParams{
		Keyword: l.GetKeyword(),
		Page:    types.NewPagination(l.GetPagination()),
		Alert:   l.GetAlert(),
		Status:  vobj.Status(l.GetStatus()),
	}
}

func (u *updateTemplateStrategyRequestBuilder) ToBo() *bo.UpdateTemplateStrategyParams {
	if types.IsNil(u) || types.IsNil(u.UpdateTemplateStrategyRequest) {
		return nil
	}

	return &bo.UpdateTemplateStrategyParams{
		ID: u.GetId(),
		Data: &bo.CreateTemplateStrategyParams{
			Alert:                  u.GetAlert(),
			Expr:                   u.GetExpr(),
			Remark:                 u.GetRemark(),
			Labels:                 vobj.NewLabels(u.GetLabels()),
			Annotations:            u.GetAnnotations(),
			StrategyLevelTemplates: NewParamsBuild().WithContext(u.ctx).StrategyModuleBuilder().APIMutationStrategyLevelTemplateItems().WithStrategyTemplateID(u.GetId()).ToBos(u.GetLevels()),
			CategoriesIDs:          u.GetCategoriesIds(),
		},
	}
}

func (c *createTemplateStrategyRequestBuilder) ToBo() *bo.CreateTemplateStrategyParams {
	if types.IsNil(c) || types.IsNil(c.CreateTemplateStrategyRequest) {
		return nil
	}

	return &bo.CreateTemplateStrategyParams{
		Alert:                  c.GetAlert(),
		Expr:                   c.GetExpr(),
		Remark:                 c.GetRemark(),
		Labels:                 vobj.NewLabels(c.GetLabels()),
		Annotations:            c.GetAnnotations(),
		StrategyLevelTemplates: NewParamsBuild().WithContext(c.ctx).StrategyModuleBuilder().APIMutationStrategyLevelTemplateItems().ToBos(c.GetLevels()),
		CategoriesIDs:          c.GetCategoriesIds(),
	}
}

func (d *doStrategyBuilder) ToAPI(strategy *bizmodel.Strategy) *adminapi.StrategyItem {
	if types.IsNil(d) || types.IsNil(strategy) {
		return nil
	}

	return &adminapi.StrategyItem{
		Name:              strategy.Name,
		Expr:              strategy.Expr,
		Levels:            NewParamsBuild().WithContext(d.ctx).StrategyModuleBuilder().DoStrategyLevelBuilder().ToAPIs(strategy.Levels),
		Labels:            strategy.Labels.Map(),
		Annotations:       strategy.Annotations,
		Datasource:        NewParamsBuild().WithContext(d.ctx).DatasourceModuleBuilder().DoDatasourceBuilder().ToAPIs(strategy.Datasource),
		Id:                strategy.ID,
		Status:            api.Status(strategy.Status),
		CreatedAt:         strategy.CreatedAt.String(),
		UpdatedAt:         strategy.UpdatedAt.String(),
		Remark:            strategy.Remark,
		GroupId:           strategy.GroupID,
		Group:             NewParamsBuild().WithContext(d.ctx).StrategyModuleBuilder().DoStrategyGroupBuilder().ToAPI(strategy.Group),
		TemplateId:        strategy.TemplateID,
		TemplateSource:    api.TemplateSourceType(strategy.TemplateSource),
		Categories:        NewParamsBuild().WithContext(d.ctx).DictModuleBuilder().DoDictBuilder().ToAPIs(types.SliceTo(strategy.Categories, func(item *bizmodel.SysDict) imodel.IDict { return item })),
		AlarmNoticeGroups: NewParamsBuild().WithContext(d.ctx).AlarmNoticeGroupModuleBuilder().DoAlarmNoticeGroupItemBuilder().ToAPIs(strategy.AlarmNoticeGroups),
	}
}

func (d *doStrategyBuilder) ToAPIs(strategies []*bizmodel.Strategy) []*adminapi.StrategyItem {
	if types.IsNil(d) || types.IsNil(strategies) {
		return nil
	}

	return types.SliceTo(strategies, func(item *bizmodel.Strategy) *adminapi.StrategyItem {
		return d.ToAPI(item)
	})
}

func (d *doStrategyBuilder) ToSelect(strategy *bizmodel.Strategy) *adminapi.SelectItem {
	if types.IsNil(d) || types.IsNil(strategy) {
		return nil
	}

	return &adminapi.SelectItem{
		Value:    strategy.ID,
		Label:    strategy.Name,
		Children: nil,
		Disabled: strategy.DeletedAt > 0 || !strategy.Status.IsEnable(),
		Extend:   &adminapi.SelectExtend{Remark: strategy.Remark},
	}
}

func (d *doStrategyBuilder) ToSelects(strategies []*bizmodel.Strategy) []*adminapi.SelectItem {
	if types.IsNil(d) || types.IsNil(strategies) {
		return nil
	}

	return types.SliceTo(strategies, func(item *bizmodel.Strategy) *adminapi.SelectItem {
		return d.ToSelect(item)
	})
}

func (u *updateStrategyStatusRequestBuilder) ToBo() *bo.UpdateStrategyStatusParams {
	if types.IsNil(u) || types.IsNil(u.UpdateStrategyStatusRequest) {
		return nil
	}

	return &bo.UpdateStrategyStatusParams{
		Ids:    u.GetIds(),
		Status: vobj.Status(u.GetStatus()),
	}
}

func (l *listStrategyRequestBuilder) ToBo() *bo.QueryStrategyListParams {
	if types.IsNil(l) || types.IsNil(l.ListStrategyRequest) {
		return nil
	}

	return &bo.QueryStrategyListParams{
		Keyword: l.GetKeyword(),
		Page:    types.NewPagination(l.GetPagination()),
		Status:  vobj.Status(l.GetStatus()),
	}
}

func (u *updateStrategyRequestBuilder) ToBo() *bo.UpdateStrategyParams {
	if types.IsNil(u) || types.IsNil(u.UpdateStrategyRequest) {
		return nil
	}

	return &bo.UpdateStrategyParams{
		ID:          u.GetId(),
		UpdateParam: NewParamsBuild().WithContext(u.ctx).StrategyModuleBuilder().WithCreateStrategyRequest(u.GetData()).ToBo(),
	}
}

func (c *createStrategyRequestBuilder) ToBo() *bo.CreateStrategyParams {
	if types.IsNil(c) || types.IsNil(c.CreateStrategyRequest) {
		return nil
	}

	claims, ok := middleware.ParseJwtClaims(c.ctx)
	if !ok {
		panic(merr.ErrorI18nUnLoginErr(c.ctx))
	}

	return &bo.CreateStrategyParams{
		GroupID:        c.GetGroupId(),
		TemplateID:     c.GetTemplateId(),
		Remark:         c.GetRemark(),
		Status:         vobj.Status(c.GetStatus()),
		Step:           c.GetStep(),
		DatasourceIDs:  c.GetDatasourceIds(),
		TemplateSource: vobj.StrategyTemplateSource(c.GetSourceType()),
		Name:           c.GetName(),
		TeamID:         claims.TeamID,
		Levels:         NewParamsBuild().WithContext(c.ctx).StrategyModuleBuilder().APIMutationStrategyLevelItems().ToBos(c.GetStrategyLevel()),
		Labels:         vobj.NewLabels(c.GetLabels()),
		Annotations:    c.GetAnnotations(),
		Expr:           c.GetExpr(),
		CategoriesIds:  c.GetCategoriesIds(),
		AlarmGroupIds:  c.GetAlarmGroupIds(),
	}
}

func (d *doStrategyGroupBuilder) WithStrategyCountMap(countMap *bo.StrategyCountMap) IDoStrategyGroupBuilder {
	if !types.IsNil(d) {
		d.strategyCountMap = countMap
	}
	return d
}

func (d *doStrategyGroupBuilder) ToAPI(group *bizmodel.StrategyGroup) *adminapi.StrategyGroupItem {
	if types.IsNil(d) || types.IsNil(group) {
		return nil
	}
	strategyCount := d.strategyCountMap
	return &adminapi.StrategyGroupItem{
		Id:                  group.ID,
		Name:                group.Name,
		Status:              api.Status(group.Status),
		CreatedAt:           group.CreatedAt.String(),
		UpdatedAt:           group.UpdatedAt.String(),
		Remark:              group.Remark,
		Creator:             "", // TODO 获取系统成员信息
		CreatorId:           group.CreatorID,
		Strategies:          NewParamsBuild().WithContext(d.ctx).StrategyModuleBuilder().DoStrategyBuilder().ToAPIs(group.Strategies),
		StrategyCount:       strategyCount.GetStrategyCountMap(group.ID),
		EnableStrategyCount: strategyCount.GetStrategyEnableMap(group.ID),
		Categories:          NewParamsBuild().WithContext(d.ctx).DictModuleBuilder().DoDictBuilder().ToAPIs(types.SliceTo(group.Categories, func(item *bizmodel.SysDict) imodel.IDict { return item })),
	}
}

func (d *doStrategyGroupBuilder) ToAPIs(groups []*bizmodel.StrategyGroup) []*adminapi.StrategyGroupItem {
	if types.IsNil(d) || types.IsNil(groups) {
		return nil
	}

	return types.SliceTo(groups, func(item *bizmodel.StrategyGroup) *adminapi.StrategyGroupItem {
		return d.ToAPI(item)
	})
}

func (d *doStrategyGroupBuilder) ToSelect(group *bizmodel.StrategyGroup) *adminapi.SelectItem {
	if types.IsNil(d) || types.IsNil(group) {
		return nil
	}

	return &adminapi.SelectItem{
		Value:    group.ID,
		Label:    group.Name,
		Children: nil,
		Disabled: group.DeletedAt > 0 || !group.Status.IsEnable(),
		Extend:   &adminapi.SelectExtend{Remark: group.Remark},
	}
}

func (d *doStrategyGroupBuilder) ToSelects(groups []*bizmodel.StrategyGroup) []*adminapi.SelectItem {
	if types.IsNil(d) || types.IsNil(groups) {
		return nil
	}

	return types.SliceTo(groups, func(item *bizmodel.StrategyGroup) *adminapi.SelectItem {
		return d.ToSelect(item)
	})
}

func (u *updateStrategyGroupStatusRequestBuilder) ToBo() *bo.UpdateStrategyGroupStatusParams {
	if types.IsNil(u) || types.IsNil(u.UpdateStrategyGroupStatusRequest) {
		return nil
	}

	return &bo.UpdateStrategyGroupStatusParams{
		IDs:    u.GetIds(),
		Status: vobj.Status(u.GetStatus()),
	}
}

func (u *updateStrategyGroupRequestBuilder) ToBo() *bo.UpdateStrategyGroupParams {
	if types.IsNil(u) || types.IsNil(u.UpdateStrategyGroupRequest) {
		return nil
	}

	return &bo.UpdateStrategyGroupParams{
		ID:          u.GetId(),
		UpdateParam: NewParamsBuild().WithContext(u.ctx).StrategyModuleBuilder().WithCreateStrategyGroupRequest(u.GetUpdate()).ToBo(),
	}
}

func (l *listStrategyGroupRequestBuilder) ToBo() *bo.QueryStrategyGroupListParams {
	if types.IsNil(l) || types.IsNil(l.ListStrategyGroupRequest) {
		return nil
	}

	return &bo.QueryStrategyGroupListParams{
		Keyword:       l.GetKeyword(),
		Page:          types.NewPagination(l.GetPagination()),
		Status:        vobj.Status(l.GetStatus()),
		CategoriesIds: l.GetCategoriesIds(),
	}
}

func (d *deleteStrategyGroupRequestBuilder) ToBo() *bo.DelStrategyGroupParams {
	if types.IsNil(d) || types.IsNil(d.DeleteStrategyGroupRequest) {
		return nil
	}

	return &bo.DelStrategyGroupParams{
		ID: d.GetId(),
	}
}

func (c *createStrategyGroupRequestBuilder) ToBo() *bo.CreateStrategyGroupParams {
	if types.IsNil(c) || types.IsNil(c.CreateStrategyGroupRequest) {
		return nil
	}

	return &bo.CreateStrategyGroupParams{
		Name:          c.GetName(),
		Remark:        c.GetRemark(),
		Status:        vobj.Status(c.GetStatus()),
		CategoriesIds: c.GetCategoriesIds(),
	}
}

func (s *strategyModuleBuilder) WithCreateStrategyGroupRequest(request *strategyapi.CreateStrategyGroupRequest) ICreateStrategyGroupRequestBuilder {
	return &createStrategyGroupRequestBuilder{ctx: s.ctx, CreateStrategyGroupRequest: request}
}

func (s *strategyModuleBuilder) WithDeleteStrategyGroupRequest(request *strategyapi.DeleteStrategyGroupRequest) IDeleteStrategyGroupRequestBuilder {
	return &deleteStrategyGroupRequestBuilder{ctx: s.ctx, DeleteStrategyGroupRequest: request}
}

func (s *strategyModuleBuilder) WithListStrategyGroupRequest(request *strategyapi.ListStrategyGroupRequest) IListStrategyGroupRequestBuilder {
	return &listStrategyGroupRequestBuilder{ctx: s.ctx, ListStrategyGroupRequest: request}
}

func (s *strategyModuleBuilder) WithUpdateStrategyGroupRequest(request *strategyapi.UpdateStrategyGroupRequest) IUpdateStrategyGroupRequestBuilder {
	return &updateStrategyGroupRequestBuilder{ctx: s.ctx, UpdateStrategyGroupRequest: request}
}

func (s *strategyModuleBuilder) WithUpdateStrategyGroupStatusRequest(request *strategyapi.UpdateStrategyGroupStatusRequest) IUpdateStrategyGroupStatusRequestBuilder {
	return &updateStrategyGroupStatusRequestBuilder{ctx: s.ctx, UpdateStrategyGroupStatusRequest: request}
}

func (s *strategyModuleBuilder) DoStrategyGroupBuilder() IDoStrategyGroupBuilder {
	return &doStrategyGroupBuilder{ctx: s.ctx}
}

func (s *strategyModuleBuilder) WithCreateStrategyRequest(request *strategyapi.CreateStrategyRequest) ICreateStrategyRequestBuilder {
	return &createStrategyRequestBuilder{ctx: s.ctx, CreateStrategyRequest: request}
}

func (s *strategyModuleBuilder) WithUpdateStrategyRequest(request *strategyapi.UpdateStrategyRequest) IUpdateStrategyRequestBuilder {
	return &updateStrategyRequestBuilder{ctx: s.ctx, UpdateStrategyRequest: request}
}

func (s *strategyModuleBuilder) WithListStrategyRequest(request *strategyapi.ListStrategyRequest) IListStrategyRequestBuilder {
	return &listStrategyRequestBuilder{ctx: s.ctx, ListStrategyRequest: request}
}

func (s *strategyModuleBuilder) WithUpdateStrategyStatusRequest(request *strategyapi.UpdateStrategyStatusRequest) IUpdateStrategyStatusRequestBuilder {
	return &updateStrategyStatusRequestBuilder{ctx: s.ctx, UpdateStrategyStatusRequest: request}
}

func (s *strategyModuleBuilder) DoStrategyBuilder() IDoStrategyBuilder {
	return &doStrategyBuilder{ctx: s.ctx}
}

func (s *strategyModuleBuilder) WithCreateTemplateStrategyRequest(request *strategyapi.CreateTemplateStrategyRequest) ICreateTemplateStrategyRequestBuilder {
	return &createTemplateStrategyRequestBuilder{ctx: s.ctx, CreateTemplateStrategyRequest: request}
}

func (s *strategyModuleBuilder) WithUpdateTemplateStrategyRequest(request *strategyapi.UpdateTemplateStrategyRequest) IUpdateTemplateStrategyRequestBuilder {
	return &updateTemplateStrategyRequestBuilder{ctx: s.ctx, UpdateTemplateStrategyRequest: request}
}

func (s *strategyModuleBuilder) WithListTemplateStrategyRequest(request *strategyapi.ListTemplateStrategyRequest) IListTemplateStrategyRequestBuilder {
	return &listTemplateStrategyRequestBuilder{ctx: s.ctx, ListTemplateStrategyRequest: request}
}

func (s *strategyModuleBuilder) WithUpdateTemplateStrategyStatusRequest(request *strategyapi.UpdateTemplateStrategyStatusRequest) IUpdateTemplateStrategyStatusRequestBuilder {
	return &updateTemplateStrategyStatusRequestBuilder{ctx: s.ctx, UpdateTemplateStrategyStatusRequest: request}
}

func (s *strategyModuleBuilder) DoTemplateStrategyBuilder() IDoTemplateStrategyBuilder {
	return &doTemplateStrategyBuilder{ctx: s.ctx}
}