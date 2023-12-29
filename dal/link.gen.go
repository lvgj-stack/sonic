// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dal

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/go-sonic/sonic/model/entity"
)

func newLink(db *gorm.DB, opts ...gen.DOOption) link {
	_link := link{}

	_link.linkDo.UseDB(db, opts...)
	_link.linkDo.UseModel(&entity.Link{})

	tableName := _link.linkDo.TableName()
	_link.ALL = field.NewAsterisk(tableName)
	_link.ID = field.NewInt32(tableName, "id")
	_link.CreateTime = field.NewTime(tableName, "create_time")
	_link.UpdateTime = field.NewTime(tableName, "update_time")
	_link.Description = field.NewString(tableName, "description")
	_link.Logo = field.NewString(tableName, "logo")
	_link.Name = field.NewString(tableName, "name")
	_link.Priority = field.NewInt32(tableName, "priority")
	_link.Team = field.NewString(tableName, "team")
	_link.URL = field.NewString(tableName, "url")

	_link.fillFieldMap()

	return _link
}

type link struct {
	linkDo linkDo

	ALL         field.Asterisk
	ID          field.Int32
	CreateTime  field.Time
	UpdateTime  field.Time
	Description field.String
	Logo        field.String
	Name        field.String
	Priority    field.Int32
	Team        field.String
	URL         field.String

	fieldMap map[string]field.Expr
}

func (l link) Table(newTableName string) *link {
	l.linkDo.UseTable(newTableName)
	return l.updateTableName(newTableName)
}

func (l link) As(alias string) *link {
	l.linkDo.DO = *(l.linkDo.As(alias).(*gen.DO))
	return l.updateTableName(alias)
}

func (l *link) updateTableName(table string) *link {
	l.ALL = field.NewAsterisk(table)
	l.ID = field.NewInt32(table, "id")
	l.CreateTime = field.NewTime(table, "create_time")
	l.UpdateTime = field.NewTime(table, "update_time")
	l.Description = field.NewString(table, "description")
	l.Logo = field.NewString(table, "logo")
	l.Name = field.NewString(table, "name")
	l.Priority = field.NewInt32(table, "priority")
	l.Team = field.NewString(table, "team")
	l.URL = field.NewString(table, "url")

	l.fillFieldMap()

	return l
}

func (l *link) WithContext(ctx context.Context) *linkDo { return l.linkDo.WithContext(ctx) }

func (l link) TableName() string { return l.linkDo.TableName() }

func (l link) Alias() string { return l.linkDo.Alias() }

func (l link) Columns(cols ...field.Expr) gen.Columns { return l.linkDo.Columns(cols...) }

func (l *link) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := l.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (l *link) fillFieldMap() {
	l.fieldMap = make(map[string]field.Expr, 9)
	l.fieldMap["id"] = l.ID
	l.fieldMap["create_time"] = l.CreateTime
	l.fieldMap["update_time"] = l.UpdateTime
	l.fieldMap["description"] = l.Description
	l.fieldMap["logo"] = l.Logo
	l.fieldMap["name"] = l.Name
	l.fieldMap["priority"] = l.Priority
	l.fieldMap["team"] = l.Team
	l.fieldMap["url"] = l.URL
}

func (l link) clone(db *gorm.DB) link {
	l.linkDo.ReplaceConnPool(db.Statement.ConnPool)
	return l
}

func (l link) replaceDB(db *gorm.DB) link {
	l.linkDo.ReplaceDB(db)
	return l
}

type linkDo struct{ gen.DO }

func (l linkDo) Debug() *linkDo {
	return l.withDO(l.DO.Debug())
}

func (l linkDo) WithContext(ctx context.Context) *linkDo {
	return l.withDO(l.DO.WithContext(ctx))
}

func (l linkDo) ReadDB() *linkDo {
	return l.Clauses(dbresolver.Read)
}

func (l linkDo) WriteDB() *linkDo {
	return l.Clauses(dbresolver.Write)
}

func (l linkDo) Session(config *gorm.Session) *linkDo {
	return l.withDO(l.DO.Session(config))
}

func (l linkDo) Clauses(conds ...clause.Expression) *linkDo {
	return l.withDO(l.DO.Clauses(conds...))
}

func (l linkDo) Returning(value interface{}, columns ...string) *linkDo {
	return l.withDO(l.DO.Returning(value, columns...))
}

func (l linkDo) Not(conds ...gen.Condition) *linkDo {
	return l.withDO(l.DO.Not(conds...))
}

func (l linkDo) Or(conds ...gen.Condition) *linkDo {
	return l.withDO(l.DO.Or(conds...))
}

func (l linkDo) Select(conds ...field.Expr) *linkDo {
	return l.withDO(l.DO.Select(conds...))
}

func (l linkDo) Where(conds ...gen.Condition) *linkDo {
	return l.withDO(l.DO.Where(conds...))
}

func (l linkDo) Order(conds ...field.Expr) *linkDo {
	return l.withDO(l.DO.Order(conds...))
}

func (l linkDo) Distinct(cols ...field.Expr) *linkDo {
	return l.withDO(l.DO.Distinct(cols...))
}

func (l linkDo) Omit(cols ...field.Expr) *linkDo {
	return l.withDO(l.DO.Omit(cols...))
}

func (l linkDo) Join(table schema.Tabler, on ...field.Expr) *linkDo {
	return l.withDO(l.DO.Join(table, on...))
}

func (l linkDo) LeftJoin(table schema.Tabler, on ...field.Expr) *linkDo {
	return l.withDO(l.DO.LeftJoin(table, on...))
}

func (l linkDo) RightJoin(table schema.Tabler, on ...field.Expr) *linkDo {
	return l.withDO(l.DO.RightJoin(table, on...))
}

func (l linkDo) Group(cols ...field.Expr) *linkDo {
	return l.withDO(l.DO.Group(cols...))
}

func (l linkDo) Having(conds ...gen.Condition) *linkDo {
	return l.withDO(l.DO.Having(conds...))
}

func (l linkDo) Limit(limit int) *linkDo {
	return l.withDO(l.DO.Limit(limit))
}

func (l linkDo) Offset(offset int) *linkDo {
	return l.withDO(l.DO.Offset(offset))
}

func (l linkDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *linkDo {
	return l.withDO(l.DO.Scopes(funcs...))
}

func (l linkDo) Unscoped() *linkDo {
	return l.withDO(l.DO.Unscoped())
}

func (l linkDo) Create(values ...*entity.Link) error {
	if len(values) == 0 {
		return nil
	}
	return l.DO.Create(values)
}

func (l linkDo) CreateInBatches(values []*entity.Link, batchSize int) error {
	return l.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (l linkDo) Save(values ...*entity.Link) error {
	if len(values) == 0 {
		return nil
	}
	return l.DO.Save(values)
}

func (l linkDo) First() (*entity.Link, error) {
	if result, err := l.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*entity.Link), nil
	}
}

func (l linkDo) Take() (*entity.Link, error) {
	if result, err := l.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*entity.Link), nil
	}
}

func (l linkDo) Last() (*entity.Link, error) {
	if result, err := l.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*entity.Link), nil
	}
}

func (l linkDo) Find() ([]*entity.Link, error) {
	result, err := l.DO.Find()
	return result.([]*entity.Link), err
}

func (l linkDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*entity.Link, err error) {
	buf := make([]*entity.Link, 0, batchSize)
	err = l.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (l linkDo) FindInBatches(result *[]*entity.Link, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return l.DO.FindInBatches(result, batchSize, fc)
}

func (l linkDo) Attrs(attrs ...field.AssignExpr) *linkDo {
	return l.withDO(l.DO.Attrs(attrs...))
}

func (l linkDo) Assign(attrs ...field.AssignExpr) *linkDo {
	return l.withDO(l.DO.Assign(attrs...))
}

func (l linkDo) Joins(fields ...field.RelationField) *linkDo {
	for _, _f := range fields {
		l = *l.withDO(l.DO.Joins(_f))
	}
	return &l
}

func (l linkDo) Preload(fields ...field.RelationField) *linkDo {
	for _, _f := range fields {
		l = *l.withDO(l.DO.Preload(_f))
	}
	return &l
}

func (l linkDo) FirstOrInit() (*entity.Link, error) {
	if result, err := l.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*entity.Link), nil
	}
}

func (l linkDo) FirstOrCreate() (*entity.Link, error) {
	if result, err := l.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*entity.Link), nil
	}
}

func (l linkDo) FindByPage(offset int, limit int) (result []*entity.Link, count int64, err error) {
	result, err = l.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = l.Offset(-1).Limit(-1).Count()
	return
}

func (l linkDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = l.Count()
	if err != nil {
		return
	}

	err = l.Offset(offset).Limit(limit).Scan(result)
	return
}

func (l linkDo) Scan(result interface{}) (err error) {
	return l.DO.Scan(result)
}

func (l linkDo) Delete(models ...*entity.Link) (result gen.ResultInfo, err error) {
	return l.DO.Delete(models)
}

func (l *linkDo) withDO(do gen.Dao) *linkDo {
	l.DO = *do.(*gen.DO)
	return l
}
