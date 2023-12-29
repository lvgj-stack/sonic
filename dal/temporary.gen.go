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

func newTemporary(db *gorm.DB, opts ...gen.DOOption) temporary {
	_temporary := temporary{}

	_temporary.temporaryDo.UseDB(db, opts...)
	_temporary.temporaryDo.UseModel(&entity.Temporary{})

	tableName := _temporary.temporaryDo.TableName()
	_temporary.ALL = field.NewAsterisk(tableName)
	_temporary.Key = field.NewString(tableName, "key")
	_temporary.Lang = field.NewString(tableName, "lang")
	_temporary.Content = field.NewString(tableName, "content")
	_temporary.Password = field.NewString(tableName, "password")
	_temporary.ClientIP = field.NewString(tableName, "client_ip")
	_temporary.Username = field.NewString(tableName, "username")
	_temporary.CreatedAt = field.NewTime(tableName, "created_at")
	_temporary.ExpireSecond = field.NewInt64(tableName, "expire_second")
	_temporary.ExpireCount = field.NewInt64(tableName, "expire_count")

	_temporary.fillFieldMap()

	return _temporary
}

type temporary struct {
	temporaryDo temporaryDo

	ALL          field.Asterisk
	Key          field.String
	Lang         field.String
	Content      field.String
	Password     field.String
	ClientIP     field.String
	Username     field.String
	CreatedAt    field.Time
	ExpireSecond field.Int64
	ExpireCount  field.Int64

	fieldMap map[string]field.Expr
}

func (t temporary) Table(newTableName string) *temporary {
	t.temporaryDo.UseTable(newTableName)
	return t.updateTableName(newTableName)
}

func (t temporary) As(alias string) *temporary {
	t.temporaryDo.DO = *(t.temporaryDo.As(alias).(*gen.DO))
	return t.updateTableName(alias)
}

func (t *temporary) updateTableName(table string) *temporary {
	t.ALL = field.NewAsterisk(table)
	t.Key = field.NewString(table, "key")
	t.Lang = field.NewString(table, "lang")
	t.Content = field.NewString(table, "content")
	t.Password = field.NewString(table, "password")
	t.ClientIP = field.NewString(table, "client_ip")
	t.Username = field.NewString(table, "username")
	t.CreatedAt = field.NewTime(table, "created_at")
	t.ExpireSecond = field.NewInt64(table, "expire_second")
	t.ExpireCount = field.NewInt64(table, "expire_count")

	t.fillFieldMap()

	return t
}

func (t *temporary) WithContext(ctx context.Context) *temporaryDo {
	return t.temporaryDo.WithContext(ctx)
}

func (t temporary) TableName() string { return t.temporaryDo.TableName() }

func (t temporary) Alias() string { return t.temporaryDo.Alias() }

func (t temporary) Columns(cols ...field.Expr) gen.Columns { return t.temporaryDo.Columns(cols...) }

func (t *temporary) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := t.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (t *temporary) fillFieldMap() {
	t.fieldMap = make(map[string]field.Expr, 9)
	t.fieldMap["key"] = t.Key
	t.fieldMap["lang"] = t.Lang
	t.fieldMap["content"] = t.Content
	t.fieldMap["password"] = t.Password
	t.fieldMap["client_ip"] = t.ClientIP
	t.fieldMap["username"] = t.Username
	t.fieldMap["created_at"] = t.CreatedAt
	t.fieldMap["expire_second"] = t.ExpireSecond
	t.fieldMap["expire_count"] = t.ExpireCount
}

func (t temporary) clone(db *gorm.DB) temporary {
	t.temporaryDo.ReplaceConnPool(db.Statement.ConnPool)
	return t
}

func (t temporary) replaceDB(db *gorm.DB) temporary {
	t.temporaryDo.ReplaceDB(db)
	return t
}

type temporaryDo struct{ gen.DO }

func (t temporaryDo) Debug() *temporaryDo {
	return t.withDO(t.DO.Debug())
}

func (t temporaryDo) WithContext(ctx context.Context) *temporaryDo {
	return t.withDO(t.DO.WithContext(ctx))
}

func (t temporaryDo) ReadDB() *temporaryDo {
	return t.Clauses(dbresolver.Read)
}

func (t temporaryDo) WriteDB() *temporaryDo {
	return t.Clauses(dbresolver.Write)
}

func (t temporaryDo) Session(config *gorm.Session) *temporaryDo {
	return t.withDO(t.DO.Session(config))
}

func (t temporaryDo) Clauses(conds ...clause.Expression) *temporaryDo {
	return t.withDO(t.DO.Clauses(conds...))
}

func (t temporaryDo) Returning(value interface{}, columns ...string) *temporaryDo {
	return t.withDO(t.DO.Returning(value, columns...))
}

func (t temporaryDo) Not(conds ...gen.Condition) *temporaryDo {
	return t.withDO(t.DO.Not(conds...))
}

func (t temporaryDo) Or(conds ...gen.Condition) *temporaryDo {
	return t.withDO(t.DO.Or(conds...))
}

func (t temporaryDo) Select(conds ...field.Expr) *temporaryDo {
	return t.withDO(t.DO.Select(conds...))
}

func (t temporaryDo) Where(conds ...gen.Condition) *temporaryDo {
	return t.withDO(t.DO.Where(conds...))
}

func (t temporaryDo) Order(conds ...field.Expr) *temporaryDo {
	return t.withDO(t.DO.Order(conds...))
}

func (t temporaryDo) Distinct(cols ...field.Expr) *temporaryDo {
	return t.withDO(t.DO.Distinct(cols...))
}

func (t temporaryDo) Omit(cols ...field.Expr) *temporaryDo {
	return t.withDO(t.DO.Omit(cols...))
}

func (t temporaryDo) Join(table schema.Tabler, on ...field.Expr) *temporaryDo {
	return t.withDO(t.DO.Join(table, on...))
}

func (t temporaryDo) LeftJoin(table schema.Tabler, on ...field.Expr) *temporaryDo {
	return t.withDO(t.DO.LeftJoin(table, on...))
}

func (t temporaryDo) RightJoin(table schema.Tabler, on ...field.Expr) *temporaryDo {
	return t.withDO(t.DO.RightJoin(table, on...))
}

func (t temporaryDo) Group(cols ...field.Expr) *temporaryDo {
	return t.withDO(t.DO.Group(cols...))
}

func (t temporaryDo) Having(conds ...gen.Condition) *temporaryDo {
	return t.withDO(t.DO.Having(conds...))
}

func (t temporaryDo) Limit(limit int) *temporaryDo {
	return t.withDO(t.DO.Limit(limit))
}

func (t temporaryDo) Offset(offset int) *temporaryDo {
	return t.withDO(t.DO.Offset(offset))
}

func (t temporaryDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *temporaryDo {
	return t.withDO(t.DO.Scopes(funcs...))
}

func (t temporaryDo) Unscoped() *temporaryDo {
	return t.withDO(t.DO.Unscoped())
}

func (t temporaryDo) Create(values ...*entity.Temporary) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Create(values)
}

func (t temporaryDo) CreateInBatches(values []*entity.Temporary, batchSize int) error {
	return t.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (t temporaryDo) Save(values ...*entity.Temporary) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Save(values)
}

func (t temporaryDo) First() (*entity.Temporary, error) {
	if result, err := t.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*entity.Temporary), nil
	}
}

func (t temporaryDo) Take() (*entity.Temporary, error) {
	if result, err := t.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*entity.Temporary), nil
	}
}

func (t temporaryDo) Last() (*entity.Temporary, error) {
	if result, err := t.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*entity.Temporary), nil
	}
}

func (t temporaryDo) Find() ([]*entity.Temporary, error) {
	result, err := t.DO.Find()
	return result.([]*entity.Temporary), err
}

func (t temporaryDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*entity.Temporary, err error) {
	buf := make([]*entity.Temporary, 0, batchSize)
	err = t.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (t temporaryDo) FindInBatches(result *[]*entity.Temporary, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return t.DO.FindInBatches(result, batchSize, fc)
}

func (t temporaryDo) Attrs(attrs ...field.AssignExpr) *temporaryDo {
	return t.withDO(t.DO.Attrs(attrs...))
}

func (t temporaryDo) Assign(attrs ...field.AssignExpr) *temporaryDo {
	return t.withDO(t.DO.Assign(attrs...))
}

func (t temporaryDo) Joins(fields ...field.RelationField) *temporaryDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Joins(_f))
	}
	return &t
}

func (t temporaryDo) Preload(fields ...field.RelationField) *temporaryDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Preload(_f))
	}
	return &t
}

func (t temporaryDo) FirstOrInit() (*entity.Temporary, error) {
	if result, err := t.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*entity.Temporary), nil
	}
}

func (t temporaryDo) FirstOrCreate() (*entity.Temporary, error) {
	if result, err := t.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*entity.Temporary), nil
	}
}

func (t temporaryDo) FindByPage(offset int, limit int) (result []*entity.Temporary, count int64, err error) {
	result, err = t.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = t.Offset(-1).Limit(-1).Count()
	return
}

func (t temporaryDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = t.Count()
	if err != nil {
		return
	}

	err = t.Offset(offset).Limit(limit).Scan(result)
	return
}

func (t temporaryDo) Scan(result interface{}) (err error) {
	return t.DO.Scan(result)
}

func (t temporaryDo) Delete(models ...*entity.Temporary) (result gen.ResultInfo, err error) {
	return t.DO.Delete(models)
}

func (t *temporaryDo) withDO(do gen.Dao) *temporaryDo {
	t.DO = *do.(*gen.DO)
	return t
}
