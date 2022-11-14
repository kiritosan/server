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

func newSubjectTopicComment(db *gorm.DB, opts ...gen.DOOption) subjectTopicComment {
	_subjectTopicComment := subjectTopicComment{}

	_subjectTopicComment.subjectTopicCommentDo.UseDB(db, opts...)
	_subjectTopicComment.subjectTopicCommentDo.UseModel(&dao.SubjectTopicComment{})

	tableName := _subjectTopicComment.subjectTopicCommentDo.TableName()
	_subjectTopicComment.ALL = field.NewAsterisk(tableName)
	_subjectTopicComment.ID = field.NewUint32(tableName, "sbj_pst_id")
	_subjectTopicComment.TopicID = field.NewUint32(tableName, "sbj_pst_mid")
	_subjectTopicComment.UID = field.NewUint32(tableName, "sbj_pst_uid")
	_subjectTopicComment.Related = field.NewUint32(tableName, "sbj_pst_related")
	_subjectTopicComment.Content = field.NewString(tableName, "sbj_pst_content")
	_subjectTopicComment.State = field.NewUint8(tableName, "sbj_pst_state")
	_subjectTopicComment.CreatedTime = field.NewUint32(tableName, "sbj_pst_dateline")

	_subjectTopicComment.fillFieldMap()

	return _subjectTopicComment
}

type subjectTopicComment struct {
	subjectTopicCommentDo subjectTopicCommentDo

	ALL         field.Asterisk
	ID          field.Uint32
	TopicID     field.Uint32
	UID         field.Uint32
	Related     field.Uint32
	Content     field.String
	State       field.Uint8
	CreatedTime field.Uint32

	fieldMap map[string]field.Expr
}

func (s subjectTopicComment) Table(newTableName string) *subjectTopicComment {
	s.subjectTopicCommentDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s subjectTopicComment) As(alias string) *subjectTopicComment {
	s.subjectTopicCommentDo.DO = *(s.subjectTopicCommentDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *subjectTopicComment) updateTableName(table string) *subjectTopicComment {
	s.ALL = field.NewAsterisk(table)
	s.ID = field.NewUint32(table, "sbj_pst_id")
	s.TopicID = field.NewUint32(table, "sbj_pst_mid")
	s.UID = field.NewUint32(table, "sbj_pst_uid")
	s.Related = field.NewUint32(table, "sbj_pst_related")
	s.Content = field.NewString(table, "sbj_pst_content")
	s.State = field.NewUint8(table, "sbj_pst_state")
	s.CreatedTime = field.NewUint32(table, "sbj_pst_dateline")

	s.fillFieldMap()

	return s
}

func (s *subjectTopicComment) WithContext(ctx context.Context) *subjectTopicCommentDo {
	return s.subjectTopicCommentDo.WithContext(ctx)
}

func (s subjectTopicComment) TableName() string { return s.subjectTopicCommentDo.TableName() }

func (s subjectTopicComment) Alias() string { return s.subjectTopicCommentDo.Alias() }

func (s *subjectTopicComment) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *subjectTopicComment) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 7)
	s.fieldMap["sbj_pst_id"] = s.ID
	s.fieldMap["sbj_pst_mid"] = s.TopicID
	s.fieldMap["sbj_pst_uid"] = s.UID
	s.fieldMap["sbj_pst_related"] = s.Related
	s.fieldMap["sbj_pst_content"] = s.Content
	s.fieldMap["sbj_pst_state"] = s.State
	s.fieldMap["sbj_pst_dateline"] = s.CreatedTime
}

func (s subjectTopicComment) clone(db *gorm.DB) subjectTopicComment {
	s.subjectTopicCommentDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s subjectTopicComment) replaceDB(db *gorm.DB) subjectTopicComment {
	s.subjectTopicCommentDo.ReplaceDB(db)
	return s
}

type subjectTopicCommentDo struct{ gen.DO }

func (s subjectTopicCommentDo) Debug() *subjectTopicCommentDo {
	return s.withDO(s.DO.Debug())
}

func (s subjectTopicCommentDo) WithContext(ctx context.Context) *subjectTopicCommentDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s subjectTopicCommentDo) ReadDB() *subjectTopicCommentDo {
	return s.Clauses(dbresolver.Read)
}

func (s subjectTopicCommentDo) WriteDB() *subjectTopicCommentDo {
	return s.Clauses(dbresolver.Write)
}

func (s subjectTopicCommentDo) Session(config *gorm.Session) *subjectTopicCommentDo {
	return s.withDO(s.DO.Session(config))
}

func (s subjectTopicCommentDo) Clauses(conds ...clause.Expression) *subjectTopicCommentDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s subjectTopicCommentDo) Returning(value interface{}, columns ...string) *subjectTopicCommentDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s subjectTopicCommentDo) Not(conds ...gen.Condition) *subjectTopicCommentDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s subjectTopicCommentDo) Or(conds ...gen.Condition) *subjectTopicCommentDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s subjectTopicCommentDo) Select(conds ...field.Expr) *subjectTopicCommentDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s subjectTopicCommentDo) Where(conds ...gen.Condition) *subjectTopicCommentDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s subjectTopicCommentDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *subjectTopicCommentDo {
	return s.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (s subjectTopicCommentDo) Order(conds ...field.Expr) *subjectTopicCommentDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s subjectTopicCommentDo) Distinct(cols ...field.Expr) *subjectTopicCommentDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s subjectTopicCommentDo) Omit(cols ...field.Expr) *subjectTopicCommentDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s subjectTopicCommentDo) Join(table schema.Tabler, on ...field.Expr) *subjectTopicCommentDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s subjectTopicCommentDo) LeftJoin(table schema.Tabler, on ...field.Expr) *subjectTopicCommentDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s subjectTopicCommentDo) RightJoin(table schema.Tabler, on ...field.Expr) *subjectTopicCommentDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s subjectTopicCommentDo) Group(cols ...field.Expr) *subjectTopicCommentDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s subjectTopicCommentDo) Having(conds ...gen.Condition) *subjectTopicCommentDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s subjectTopicCommentDo) Limit(limit int) *subjectTopicCommentDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s subjectTopicCommentDo) Offset(offset int) *subjectTopicCommentDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s subjectTopicCommentDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *subjectTopicCommentDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s subjectTopicCommentDo) Unscoped() *subjectTopicCommentDo {
	return s.withDO(s.DO.Unscoped())
}

func (s subjectTopicCommentDo) Create(values ...*dao.SubjectTopicComment) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s subjectTopicCommentDo) CreateInBatches(values []*dao.SubjectTopicComment, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s subjectTopicCommentDo) Save(values ...*dao.SubjectTopicComment) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s subjectTopicCommentDo) First() (*dao.SubjectTopicComment, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*dao.SubjectTopicComment), nil
	}
}

func (s subjectTopicCommentDo) Take() (*dao.SubjectTopicComment, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*dao.SubjectTopicComment), nil
	}
}

func (s subjectTopicCommentDo) Last() (*dao.SubjectTopicComment, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*dao.SubjectTopicComment), nil
	}
}

func (s subjectTopicCommentDo) Find() ([]*dao.SubjectTopicComment, error) {
	result, err := s.DO.Find()
	return result.([]*dao.SubjectTopicComment), err
}

func (s subjectTopicCommentDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*dao.SubjectTopicComment, err error) {
	buf := make([]*dao.SubjectTopicComment, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s subjectTopicCommentDo) FindInBatches(result *[]*dao.SubjectTopicComment, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s subjectTopicCommentDo) Attrs(attrs ...field.AssignExpr) *subjectTopicCommentDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s subjectTopicCommentDo) Assign(attrs ...field.AssignExpr) *subjectTopicCommentDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s subjectTopicCommentDo) Joins(fields ...field.RelationField) *subjectTopicCommentDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s subjectTopicCommentDo) Preload(fields ...field.RelationField) *subjectTopicCommentDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s subjectTopicCommentDo) FirstOrInit() (*dao.SubjectTopicComment, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*dao.SubjectTopicComment), nil
	}
}

func (s subjectTopicCommentDo) FirstOrCreate() (*dao.SubjectTopicComment, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*dao.SubjectTopicComment), nil
	}
}

func (s subjectTopicCommentDo) FindByPage(offset int, limit int) (result []*dao.SubjectTopicComment, count int64, err error) {
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

func (s subjectTopicCommentDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s subjectTopicCommentDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s subjectTopicCommentDo) Delete(models ...*dao.SubjectTopicComment) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *subjectTopicCommentDo) withDO(do gen.Dao) *subjectTopicCommentDo {
	s.DO = *do.(*gen.DO)
	return s
}
