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

func newTimeLine(db *gorm.DB, opts ...gen.DOOption) timeLine {
	_timeLine := timeLine{}

	_timeLine.timeLineDo.UseDB(db, opts...)
	_timeLine.timeLineDo.UseModel(&dao.TimeLine{})

	tableName := _timeLine.timeLineDo.TableName()
	_timeLine.ALL = field.NewAsterisk(tableName)
	_timeLine.ID = field.NewField(tableName, "tml_id")
	_timeLine.UID = field.NewField(tableName, "tml_uid")
	_timeLine.Cat = field.NewUint16(tableName, "tml_cat")
	_timeLine.Type = field.NewUint16(tableName, "tml_type")
	_timeLine.Related = field.NewString(tableName, "tml_related")
	_timeLine.Memo = field.NewBytes(tableName, "tml_memo")
	_timeLine.Img = field.NewBytes(tableName, "tml_img")
	_timeLine.Batch = field.NewUint8(tableName, "tml_batch")
	_timeLine.Source = field.NewUint8(tableName, "tml_source")
	_timeLine.Replies = field.NewUint32(tableName, "tml_replies")
	_timeLine.Dateline = field.NewUint32(tableName, "tml_dateline")
	_timeLine.Status = field.NewUint8(tableName, "tml_status")

	_timeLine.fillFieldMap()

	return _timeLine
}

type timeLine struct {
	timeLineDo timeLineDo

	ALL      field.Asterisk
	ID       field.Field
	UID      field.Field
	Cat      field.Uint16
	Type     field.Uint16
	Related  field.String
	Memo     field.Bytes
	Img      field.Bytes
	Batch    field.Uint8
	Source   field.Uint8  // 更新来源
	Replies  field.Uint32 // 回复数
	Dateline field.Uint32
	Status   field.Uint8

	fieldMap map[string]field.Expr
}

func (t timeLine) Table(newTableName string) *timeLine {
	t.timeLineDo.UseTable(newTableName)
	return t.updateTableName(newTableName)
}

func (t timeLine) As(alias string) *timeLine {
	t.timeLineDo.DO = *(t.timeLineDo.As(alias).(*gen.DO))
	return t.updateTableName(alias)
}

func (t *timeLine) updateTableName(table string) *timeLine {
	t.ALL = field.NewAsterisk(table)
	t.ID = field.NewField(table, "tml_id")
	t.UID = field.NewField(table, "tml_uid")
	t.Cat = field.NewUint16(table, "tml_cat")
	t.Type = field.NewUint16(table, "tml_type")
	t.Related = field.NewString(table, "tml_related")
	t.Memo = field.NewBytes(table, "tml_memo")
	t.Img = field.NewBytes(table, "tml_img")
	t.Batch = field.NewUint8(table, "tml_batch")
	t.Source = field.NewUint8(table, "tml_source")
	t.Replies = field.NewUint32(table, "tml_replies")
	t.Dateline = field.NewUint32(table, "tml_dateline")
	t.Status = field.NewUint8(table, "tml_status")

	t.fillFieldMap()

	return t
}

func (t *timeLine) WithContext(ctx context.Context) *timeLineDo { return t.timeLineDo.WithContext(ctx) }

func (t timeLine) TableName() string { return t.timeLineDo.TableName() }

func (t timeLine) Alias() string { return t.timeLineDo.Alias() }

func (t *timeLine) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := t.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (t *timeLine) fillFieldMap() {
	t.fieldMap = make(map[string]field.Expr, 12)
	t.fieldMap["tml_id"] = t.ID
	t.fieldMap["tml_uid"] = t.UID
	t.fieldMap["tml_cat"] = t.Cat
	t.fieldMap["tml_type"] = t.Type
	t.fieldMap["tml_related"] = t.Related
	t.fieldMap["tml_memo"] = t.Memo
	t.fieldMap["tml_img"] = t.Img
	t.fieldMap["tml_batch"] = t.Batch
	t.fieldMap["tml_source"] = t.Source
	t.fieldMap["tml_replies"] = t.Replies
	t.fieldMap["tml_dateline"] = t.Dateline
	t.fieldMap["tml_status"] = t.Status
}

func (t timeLine) clone(db *gorm.DB) timeLine {
	t.timeLineDo.ReplaceConnPool(db.Statement.ConnPool)
	return t
}

func (t timeLine) replaceDB(db *gorm.DB) timeLine {
	t.timeLineDo.ReplaceDB(db)
	return t
}

type timeLineDo struct{ gen.DO }

func (t timeLineDo) Debug() *timeLineDo {
	return t.withDO(t.DO.Debug())
}

func (t timeLineDo) WithContext(ctx context.Context) *timeLineDo {
	return t.withDO(t.DO.WithContext(ctx))
}

func (t timeLineDo) ReadDB() *timeLineDo {
	return t.Clauses(dbresolver.Read)
}

func (t timeLineDo) WriteDB() *timeLineDo {
	return t.Clauses(dbresolver.Write)
}

func (t timeLineDo) Session(config *gorm.Session) *timeLineDo {
	return t.withDO(t.DO.Session(config))
}

func (t timeLineDo) Clauses(conds ...clause.Expression) *timeLineDo {
	return t.withDO(t.DO.Clauses(conds...))
}

func (t timeLineDo) Returning(value interface{}, columns ...string) *timeLineDo {
	return t.withDO(t.DO.Returning(value, columns...))
}

func (t timeLineDo) Not(conds ...gen.Condition) *timeLineDo {
	return t.withDO(t.DO.Not(conds...))
}

func (t timeLineDo) Or(conds ...gen.Condition) *timeLineDo {
	return t.withDO(t.DO.Or(conds...))
}

func (t timeLineDo) Select(conds ...field.Expr) *timeLineDo {
	return t.withDO(t.DO.Select(conds...))
}

func (t timeLineDo) Where(conds ...gen.Condition) *timeLineDo {
	return t.withDO(t.DO.Where(conds...))
}

func (t timeLineDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *timeLineDo {
	return t.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (t timeLineDo) Order(conds ...field.Expr) *timeLineDo {
	return t.withDO(t.DO.Order(conds...))
}

func (t timeLineDo) Distinct(cols ...field.Expr) *timeLineDo {
	return t.withDO(t.DO.Distinct(cols...))
}

func (t timeLineDo) Omit(cols ...field.Expr) *timeLineDo {
	return t.withDO(t.DO.Omit(cols...))
}

func (t timeLineDo) Join(table schema.Tabler, on ...field.Expr) *timeLineDo {
	return t.withDO(t.DO.Join(table, on...))
}

func (t timeLineDo) LeftJoin(table schema.Tabler, on ...field.Expr) *timeLineDo {
	return t.withDO(t.DO.LeftJoin(table, on...))
}

func (t timeLineDo) RightJoin(table schema.Tabler, on ...field.Expr) *timeLineDo {
	return t.withDO(t.DO.RightJoin(table, on...))
}

func (t timeLineDo) Group(cols ...field.Expr) *timeLineDo {
	return t.withDO(t.DO.Group(cols...))
}

func (t timeLineDo) Having(conds ...gen.Condition) *timeLineDo {
	return t.withDO(t.DO.Having(conds...))
}

func (t timeLineDo) Limit(limit int) *timeLineDo {
	return t.withDO(t.DO.Limit(limit))
}

func (t timeLineDo) Offset(offset int) *timeLineDo {
	return t.withDO(t.DO.Offset(offset))
}

func (t timeLineDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *timeLineDo {
	return t.withDO(t.DO.Scopes(funcs...))
}

func (t timeLineDo) Unscoped() *timeLineDo {
	return t.withDO(t.DO.Unscoped())
}

func (t timeLineDo) Create(values ...*dao.TimeLine) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Create(values)
}

func (t timeLineDo) CreateInBatches(values []*dao.TimeLine, batchSize int) error {
	return t.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (t timeLineDo) Save(values ...*dao.TimeLine) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Save(values)
}

func (t timeLineDo) First() (*dao.TimeLine, error) {
	if result, err := t.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*dao.TimeLine), nil
	}
}

func (t timeLineDo) Take() (*dao.TimeLine, error) {
	if result, err := t.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*dao.TimeLine), nil
	}
}

func (t timeLineDo) Last() (*dao.TimeLine, error) {
	if result, err := t.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*dao.TimeLine), nil
	}
}

func (t timeLineDo) Find() ([]*dao.TimeLine, error) {
	result, err := t.DO.Find()
	return result.([]*dao.TimeLine), err
}

func (t timeLineDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*dao.TimeLine, err error) {
	buf := make([]*dao.TimeLine, 0, batchSize)
	err = t.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (t timeLineDo) FindInBatches(result *[]*dao.TimeLine, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return t.DO.FindInBatches(result, batchSize, fc)
}

func (t timeLineDo) Attrs(attrs ...field.AssignExpr) *timeLineDo {
	return t.withDO(t.DO.Attrs(attrs...))
}

func (t timeLineDo) Assign(attrs ...field.AssignExpr) *timeLineDo {
	return t.withDO(t.DO.Assign(attrs...))
}

func (t timeLineDo) Joins(fields ...field.RelationField) *timeLineDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Joins(_f))
	}
	return &t
}

func (t timeLineDo) Preload(fields ...field.RelationField) *timeLineDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Preload(_f))
	}
	return &t
}

func (t timeLineDo) FirstOrInit() (*dao.TimeLine, error) {
	if result, err := t.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*dao.TimeLine), nil
	}
}

func (t timeLineDo) FirstOrCreate() (*dao.TimeLine, error) {
	if result, err := t.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*dao.TimeLine), nil
	}
}

func (t timeLineDo) FindByPage(offset int, limit int) (result []*dao.TimeLine, count int64, err error) {
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

func (t timeLineDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = t.Count()
	if err != nil {
		return
	}

	err = t.Offset(offset).Limit(limit).Scan(result)
	return
}

func (t timeLineDo) Scan(result interface{}) (err error) {
	return t.DO.Scan(result)
}

func (t timeLineDo) Delete(models ...*dao.TimeLine) (result gen.ResultInfo, err error) {
	return t.DO.Delete(models)
}

func (t *timeLineDo) withDO(do gen.Dao) *timeLineDo {
	t.DO = *do.(*gen.DO)
	return t
}
