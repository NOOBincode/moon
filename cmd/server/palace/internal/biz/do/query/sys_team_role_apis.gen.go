// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/aide-cloud/moon/cmd/server/palace/internal/biz/do/model"
)

func newSysTeamRoleAPI(db *gorm.DB, opts ...gen.DOOption) sysTeamRoleAPI {
	_sysTeamRoleAPI := sysTeamRoleAPI{}

	_sysTeamRoleAPI.sysTeamRoleAPIDo.UseDB(db, opts...)
	_sysTeamRoleAPI.sysTeamRoleAPIDo.UseModel(&model.SysTeamRoleAPI{})

	tableName := _sysTeamRoleAPI.sysTeamRoleAPIDo.TableName()
	_sysTeamRoleAPI.ALL = field.NewAsterisk(tableName)
	_sysTeamRoleAPI.SysAPIID = field.NewUint32(tableName, "sys_api_id")
	_sysTeamRoleAPI.SysTeamRoleID = field.NewUint32(tableName, "sys_team_role_id")
	_sysTeamRoleAPI.ID = field.NewUint32(tableName, "id")

	_sysTeamRoleAPI.fillFieldMap()

	return _sysTeamRoleAPI
}

type sysTeamRoleAPI struct {
	sysTeamRoleAPIDo

	ALL           field.Asterisk
	SysAPIID      field.Uint32
	SysTeamRoleID field.Uint32
	ID            field.Uint32

	fieldMap map[string]field.Expr
}

func (s sysTeamRoleAPI) Table(newTableName string) *sysTeamRoleAPI {
	s.sysTeamRoleAPIDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s sysTeamRoleAPI) As(alias string) *sysTeamRoleAPI {
	s.sysTeamRoleAPIDo.DO = *(s.sysTeamRoleAPIDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *sysTeamRoleAPI) updateTableName(table string) *sysTeamRoleAPI {
	s.ALL = field.NewAsterisk(table)
	s.SysAPIID = field.NewUint32(table, "sys_api_id")
	s.SysTeamRoleID = field.NewUint32(table, "sys_team_role_id")
	s.ID = field.NewUint32(table, "id")

	s.fillFieldMap()

	return s
}

func (s *sysTeamRoleAPI) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *sysTeamRoleAPI) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 3)
	s.fieldMap["sys_api_id"] = s.SysAPIID
	s.fieldMap["sys_team_role_id"] = s.SysTeamRoleID
	s.fieldMap["id"] = s.ID
}

func (s sysTeamRoleAPI) clone(db *gorm.DB) sysTeamRoleAPI {
	s.sysTeamRoleAPIDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s sysTeamRoleAPI) replaceDB(db *gorm.DB) sysTeamRoleAPI {
	s.sysTeamRoleAPIDo.ReplaceDB(db)
	return s
}

type sysTeamRoleAPIDo struct{ gen.DO }

type ISysTeamRoleAPIDo interface {
	gen.SubQuery
	Debug() ISysTeamRoleAPIDo
	WithContext(ctx context.Context) ISysTeamRoleAPIDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ISysTeamRoleAPIDo
	WriteDB() ISysTeamRoleAPIDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ISysTeamRoleAPIDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ISysTeamRoleAPIDo
	Not(conds ...gen.Condition) ISysTeamRoleAPIDo
	Or(conds ...gen.Condition) ISysTeamRoleAPIDo
	Select(conds ...field.Expr) ISysTeamRoleAPIDo
	Where(conds ...gen.Condition) ISysTeamRoleAPIDo
	Order(conds ...field.Expr) ISysTeamRoleAPIDo
	Distinct(cols ...field.Expr) ISysTeamRoleAPIDo
	Omit(cols ...field.Expr) ISysTeamRoleAPIDo
	Join(table schema.Tabler, on ...field.Expr) ISysTeamRoleAPIDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ISysTeamRoleAPIDo
	RightJoin(table schema.Tabler, on ...field.Expr) ISysTeamRoleAPIDo
	Group(cols ...field.Expr) ISysTeamRoleAPIDo
	Having(conds ...gen.Condition) ISysTeamRoleAPIDo
	Limit(limit int) ISysTeamRoleAPIDo
	Offset(offset int) ISysTeamRoleAPIDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ISysTeamRoleAPIDo
	Unscoped() ISysTeamRoleAPIDo
	Create(values ...*model.SysTeamRoleAPI) error
	CreateInBatches(values []*model.SysTeamRoleAPI, batchSize int) error
	Save(values ...*model.SysTeamRoleAPI) error
	First() (*model.SysTeamRoleAPI, error)
	Take() (*model.SysTeamRoleAPI, error)
	Last() (*model.SysTeamRoleAPI, error)
	Find() ([]*model.SysTeamRoleAPI, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SysTeamRoleAPI, err error)
	FindInBatches(result *[]*model.SysTeamRoleAPI, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.SysTeamRoleAPI) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ISysTeamRoleAPIDo
	Assign(attrs ...field.AssignExpr) ISysTeamRoleAPIDo
	Joins(fields ...field.RelationField) ISysTeamRoleAPIDo
	Preload(fields ...field.RelationField) ISysTeamRoleAPIDo
	FirstOrInit() (*model.SysTeamRoleAPI, error)
	FirstOrCreate() (*model.SysTeamRoleAPI, error)
	FindByPage(offset int, limit int) (result []*model.SysTeamRoleAPI, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ISysTeamRoleAPIDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (s sysTeamRoleAPIDo) Debug() ISysTeamRoleAPIDo {
	return s.withDO(s.DO.Debug())
}

func (s sysTeamRoleAPIDo) WithContext(ctx context.Context) ISysTeamRoleAPIDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s sysTeamRoleAPIDo) ReadDB() ISysTeamRoleAPIDo {
	return s.Clauses(dbresolver.Read)
}

func (s sysTeamRoleAPIDo) WriteDB() ISysTeamRoleAPIDo {
	return s.Clauses(dbresolver.Write)
}

func (s sysTeamRoleAPIDo) Session(config *gorm.Session) ISysTeamRoleAPIDo {
	return s.withDO(s.DO.Session(config))
}

func (s sysTeamRoleAPIDo) Clauses(conds ...clause.Expression) ISysTeamRoleAPIDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s sysTeamRoleAPIDo) Returning(value interface{}, columns ...string) ISysTeamRoleAPIDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s sysTeamRoleAPIDo) Not(conds ...gen.Condition) ISysTeamRoleAPIDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s sysTeamRoleAPIDo) Or(conds ...gen.Condition) ISysTeamRoleAPIDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s sysTeamRoleAPIDo) Select(conds ...field.Expr) ISysTeamRoleAPIDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s sysTeamRoleAPIDo) Where(conds ...gen.Condition) ISysTeamRoleAPIDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s sysTeamRoleAPIDo) Order(conds ...field.Expr) ISysTeamRoleAPIDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s sysTeamRoleAPIDo) Distinct(cols ...field.Expr) ISysTeamRoleAPIDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s sysTeamRoleAPIDo) Omit(cols ...field.Expr) ISysTeamRoleAPIDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s sysTeamRoleAPIDo) Join(table schema.Tabler, on ...field.Expr) ISysTeamRoleAPIDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s sysTeamRoleAPIDo) LeftJoin(table schema.Tabler, on ...field.Expr) ISysTeamRoleAPIDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s sysTeamRoleAPIDo) RightJoin(table schema.Tabler, on ...field.Expr) ISysTeamRoleAPIDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s sysTeamRoleAPIDo) Group(cols ...field.Expr) ISysTeamRoleAPIDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s sysTeamRoleAPIDo) Having(conds ...gen.Condition) ISysTeamRoleAPIDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s sysTeamRoleAPIDo) Limit(limit int) ISysTeamRoleAPIDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s sysTeamRoleAPIDo) Offset(offset int) ISysTeamRoleAPIDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s sysTeamRoleAPIDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ISysTeamRoleAPIDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s sysTeamRoleAPIDo) Unscoped() ISysTeamRoleAPIDo {
	return s.withDO(s.DO.Unscoped())
}

func (s sysTeamRoleAPIDo) Create(values ...*model.SysTeamRoleAPI) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s sysTeamRoleAPIDo) CreateInBatches(values []*model.SysTeamRoleAPI, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s sysTeamRoleAPIDo) Save(values ...*model.SysTeamRoleAPI) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s sysTeamRoleAPIDo) First() (*model.SysTeamRoleAPI, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysTeamRoleAPI), nil
	}
}

func (s sysTeamRoleAPIDo) Take() (*model.SysTeamRoleAPI, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysTeamRoleAPI), nil
	}
}

func (s sysTeamRoleAPIDo) Last() (*model.SysTeamRoleAPI, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysTeamRoleAPI), nil
	}
}

func (s sysTeamRoleAPIDo) Find() ([]*model.SysTeamRoleAPI, error) {
	result, err := s.DO.Find()
	return result.([]*model.SysTeamRoleAPI), err
}

func (s sysTeamRoleAPIDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SysTeamRoleAPI, err error) {
	buf := make([]*model.SysTeamRoleAPI, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s sysTeamRoleAPIDo) FindInBatches(result *[]*model.SysTeamRoleAPI, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s sysTeamRoleAPIDo) Attrs(attrs ...field.AssignExpr) ISysTeamRoleAPIDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s sysTeamRoleAPIDo) Assign(attrs ...field.AssignExpr) ISysTeamRoleAPIDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s sysTeamRoleAPIDo) Joins(fields ...field.RelationField) ISysTeamRoleAPIDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s sysTeamRoleAPIDo) Preload(fields ...field.RelationField) ISysTeamRoleAPIDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s sysTeamRoleAPIDo) FirstOrInit() (*model.SysTeamRoleAPI, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysTeamRoleAPI), nil
	}
}

func (s sysTeamRoleAPIDo) FirstOrCreate() (*model.SysTeamRoleAPI, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysTeamRoleAPI), nil
	}
}

func (s sysTeamRoleAPIDo) FindByPage(offset int, limit int) (result []*model.SysTeamRoleAPI, count int64, err error) {
	result, err = s.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = s.Offset(-1).Limit(-1).Count()
	return
}

func (s sysTeamRoleAPIDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s sysTeamRoleAPIDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s sysTeamRoleAPIDo) Delete(models ...*model.SysTeamRoleAPI) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *sysTeamRoleAPIDo) withDO(do gen.Dao) *sysTeamRoleAPIDo {
	s.DO = *do.(*gen.DO)
	return s
}
