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

	"github.com/bangumi/server/internal/dal/dao"
)

func newAccessToken(db *gorm.DB, opts ...gen.DOOption) accessToken {
	_accessToken := accessToken{}

	_accessToken.accessTokenDo.UseDB(db, opts...)
	_accessToken.accessTokenDo.UseModel(&dao.AccessToken{})

	tableName := _accessToken.accessTokenDo.TableName()
	_accessToken.ALL = field.NewAsterisk(tableName)
	_accessToken.ID = field.NewUint32(tableName, "id")
	_accessToken.Type = field.NewUint8(tableName, "type")
	_accessToken.AccessToken = field.NewString(tableName, "access_token")
	_accessToken.ClientID = field.NewString(tableName, "client_id")
	_accessToken.UserID = field.NewString(tableName, "user_id")
	_accessToken.ExpiredAt = field.NewTime(tableName, "expires")
	_accessToken.Scope = field.NewString(tableName, "scope")
	_accessToken.Info = field.NewBytes(tableName, "info")

	_accessToken.fillFieldMap()

	return _accessToken
}

type accessToken struct {
	accessTokenDo accessTokenDo

	ALL         field.Asterisk
	ID          field.Uint32
	Type        field.Uint8
	AccessToken field.String
	ClientID    field.String
	UserID      field.String
	ExpiredAt   field.Time
	Scope       field.String
	Info        field.Bytes

	fieldMap map[string]field.Expr
}

func (a accessToken) Table(newTableName string) *accessToken {
	a.accessTokenDo.UseTable(newTableName)
	return a.updateTableName(newTableName)
}

func (a accessToken) As(alias string) *accessToken {
	a.accessTokenDo.DO = *(a.accessTokenDo.As(alias).(*gen.DO))
	return a.updateTableName(alias)
}

func (a *accessToken) updateTableName(table string) *accessToken {
	a.ALL = field.NewAsterisk(table)
	a.ID = field.NewUint32(table, "id")
	a.Type = field.NewUint8(table, "type")
	a.AccessToken = field.NewString(table, "access_token")
	a.ClientID = field.NewString(table, "client_id")
	a.UserID = field.NewString(table, "user_id")
	a.ExpiredAt = field.NewTime(table, "expires")
	a.Scope = field.NewString(table, "scope")
	a.Info = field.NewBytes(table, "info")

	a.fillFieldMap()

	return a
}

func (a *accessToken) WithContext(ctx context.Context) *accessTokenDo {
	return a.accessTokenDo.WithContext(ctx)
}

func (a accessToken) TableName() string { return a.accessTokenDo.TableName() }

func (a accessToken) Alias() string { return a.accessTokenDo.Alias() }

func (a *accessToken) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := a.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (a *accessToken) fillFieldMap() {
	a.fieldMap = make(map[string]field.Expr, 8)
	a.fieldMap["id"] = a.ID
	a.fieldMap["type"] = a.Type
	a.fieldMap["access_token"] = a.AccessToken
	a.fieldMap["client_id"] = a.ClientID
	a.fieldMap["user_id"] = a.UserID
	a.fieldMap["expires"] = a.ExpiredAt
	a.fieldMap["scope"] = a.Scope
	a.fieldMap["info"] = a.Info
}

func (a accessToken) clone(db *gorm.DB) accessToken {
	a.accessTokenDo.ReplaceConnPool(db.Statement.ConnPool)
	return a
}

func (a accessToken) replaceDB(db *gorm.DB) accessToken {
	a.accessTokenDo.ReplaceDB(db)
	return a
}

type accessTokenDo struct{ gen.DO }

func (a accessTokenDo) Debug() *accessTokenDo {
	return a.withDO(a.DO.Debug())
}

func (a accessTokenDo) WithContext(ctx context.Context) *accessTokenDo {
	return a.withDO(a.DO.WithContext(ctx))
}

func (a accessTokenDo) ReadDB() *accessTokenDo {
	return a.Clauses(dbresolver.Read)
}

func (a accessTokenDo) WriteDB() *accessTokenDo {
	return a.Clauses(dbresolver.Write)
}

func (a accessTokenDo) Session(config *gorm.Session) *accessTokenDo {
	return a.withDO(a.DO.Session(config))
}

func (a accessTokenDo) Clauses(conds ...clause.Expression) *accessTokenDo {
	return a.withDO(a.DO.Clauses(conds...))
}

func (a accessTokenDo) Returning(value interface{}, columns ...string) *accessTokenDo {
	return a.withDO(a.DO.Returning(value, columns...))
}

func (a accessTokenDo) Not(conds ...gen.Condition) *accessTokenDo {
	return a.withDO(a.DO.Not(conds...))
}

func (a accessTokenDo) Or(conds ...gen.Condition) *accessTokenDo {
	return a.withDO(a.DO.Or(conds...))
}

func (a accessTokenDo) Select(conds ...field.Expr) *accessTokenDo {
	return a.withDO(a.DO.Select(conds...))
}

func (a accessTokenDo) Where(conds ...gen.Condition) *accessTokenDo {
	return a.withDO(a.DO.Where(conds...))
}

func (a accessTokenDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *accessTokenDo {
	return a.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (a accessTokenDo) Order(conds ...field.Expr) *accessTokenDo {
	return a.withDO(a.DO.Order(conds...))
}

func (a accessTokenDo) Distinct(cols ...field.Expr) *accessTokenDo {
	return a.withDO(a.DO.Distinct(cols...))
}

func (a accessTokenDo) Omit(cols ...field.Expr) *accessTokenDo {
	return a.withDO(a.DO.Omit(cols...))
}

func (a accessTokenDo) Join(table schema.Tabler, on ...field.Expr) *accessTokenDo {
	return a.withDO(a.DO.Join(table, on...))
}

func (a accessTokenDo) LeftJoin(table schema.Tabler, on ...field.Expr) *accessTokenDo {
	return a.withDO(a.DO.LeftJoin(table, on...))
}

func (a accessTokenDo) RightJoin(table schema.Tabler, on ...field.Expr) *accessTokenDo {
	return a.withDO(a.DO.RightJoin(table, on...))
}

func (a accessTokenDo) Group(cols ...field.Expr) *accessTokenDo {
	return a.withDO(a.DO.Group(cols...))
}

func (a accessTokenDo) Having(conds ...gen.Condition) *accessTokenDo {
	return a.withDO(a.DO.Having(conds...))
}

func (a accessTokenDo) Limit(limit int) *accessTokenDo {
	return a.withDO(a.DO.Limit(limit))
}

func (a accessTokenDo) Offset(offset int) *accessTokenDo {
	return a.withDO(a.DO.Offset(offset))
}

func (a accessTokenDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *accessTokenDo {
	return a.withDO(a.DO.Scopes(funcs...))
}

func (a accessTokenDo) Unscoped() *accessTokenDo {
	return a.withDO(a.DO.Unscoped())
}

func (a accessTokenDo) Create(values ...*dao.AccessToken) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Create(values)
}

func (a accessTokenDo) CreateInBatches(values []*dao.AccessToken, batchSize int) error {
	return a.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (a accessTokenDo) Save(values ...*dao.AccessToken) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Save(values)
}

func (a accessTokenDo) First() (*dao.AccessToken, error) {
	if result, err := a.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*dao.AccessToken), nil
	}
}

func (a accessTokenDo) Take() (*dao.AccessToken, error) {
	if result, err := a.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*dao.AccessToken), nil
	}
}

func (a accessTokenDo) Last() (*dao.AccessToken, error) {
	if result, err := a.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*dao.AccessToken), nil
	}
}

func (a accessTokenDo) Find() ([]*dao.AccessToken, error) {
	result, err := a.DO.Find()
	return result.([]*dao.AccessToken), err
}

func (a accessTokenDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*dao.AccessToken, err error) {
	buf := make([]*dao.AccessToken, 0, batchSize)
	err = a.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (a accessTokenDo) FindInBatches(result *[]*dao.AccessToken, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return a.DO.FindInBatches(result, batchSize, fc)
}

func (a accessTokenDo) Attrs(attrs ...field.AssignExpr) *accessTokenDo {
	return a.withDO(a.DO.Attrs(attrs...))
}

func (a accessTokenDo) Assign(attrs ...field.AssignExpr) *accessTokenDo {
	return a.withDO(a.DO.Assign(attrs...))
}

func (a accessTokenDo) Joins(fields ...field.RelationField) *accessTokenDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Joins(_f))
	}
	return &a
}

func (a accessTokenDo) Preload(fields ...field.RelationField) *accessTokenDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Preload(_f))
	}
	return &a
}

func (a accessTokenDo) FirstOrInit() (*dao.AccessToken, error) {
	if result, err := a.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*dao.AccessToken), nil
	}
}

func (a accessTokenDo) FirstOrCreate() (*dao.AccessToken, error) {
	if result, err := a.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*dao.AccessToken), nil
	}
}

func (a accessTokenDo) FindByPage(offset int, limit int) (result []*dao.AccessToken, count int64, err error) {
	result, err = a.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = a.Offset(-1).Limit(-1).Count()
	return
}

func (a accessTokenDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = a.Count()
	if err != nil {
		return
	}

	err = a.Offset(offset).Limit(limit).Scan(result)
	return
}

func (a accessTokenDo) Scan(result interface{}) (err error) {
	return a.DO.Scan(result)
}

func (a accessTokenDo) Delete(models ...*dao.AccessToken) (result gen.ResultInfo, err error) {
	return a.DO.Delete(models)
}

func (a *accessTokenDo) withDO(do gen.Dao) *accessTokenDo {
	a.DO = *do.(*gen.DO)
	return a
}
