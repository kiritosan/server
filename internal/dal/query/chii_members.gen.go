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

func newMember(db *gorm.DB, opts ...gen.DOOption) member {
	_member := member{}

	_member.memberDo.UseDB(db, opts...)
	_member.memberDo.UseModel(&dao.Member{})

	tableName := _member.memberDo.TableName()
	_member.ALL = field.NewAsterisk(tableName)
	_member.ID = field.NewField(tableName, "uid")
	_member.Username = field.NewString(tableName, "username")
	_member.Nickname = field.NewString(tableName, "nickname")
	_member.Avatar = field.NewString(tableName, "avatar")
	_member.Groupid = field.NewUint8(tableName, "groupid")
	_member.Regdate = field.NewInt64(tableName, "regdate")
	_member.Lastvisit = field.NewUint32(tableName, "lastvisit")
	_member.Lastactivity = field.NewUint32(tableName, "lastactivity")
	_member.Lastpost = field.NewUint32(tableName, "lastpost")
	_member.Dateformat = field.NewString(tableName, "dateformat")
	_member.Timeformat = field.NewBool(tableName, "timeformat")
	_member.Timeoffset = field.NewString(tableName, "timeoffset")
	_member.Newpm = field.NewBool(tableName, "newpm")
	_member.NewNotify = field.NewUint16(tableName, "new_notify")
	_member.Sign = field.NewString(tableName, "sign")
	_member.PasswordCrypt = field.NewBytes(tableName, "password_crypt")
	_member.Email = field.NewString(tableName, "email")
	_member.Fields = memberHasOneFields{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Fields", "dao.MemberField"),
	}

	_member.fillFieldMap()

	return _member
}

type member struct {
	memberDo memberDo

	ALL           field.Asterisk
	ID            field.Field
	Username      field.String
	Nickname      field.String
	Avatar        field.String
	Groupid       field.Uint8
	Regdate       field.Int64
	Lastvisit     field.Uint32
	Lastactivity  field.Uint32
	Lastpost      field.Uint32
	Dateformat    field.String
	Timeformat    field.Bool
	Timeoffset    field.String
	Newpm         field.Bool
	NewNotify     field.Uint16 // 新提醒
	Sign          field.String
	PasswordCrypt field.Bytes
	Email         field.String
	Fields        memberHasOneFields

	fieldMap map[string]field.Expr
}

func (m member) Table(newTableName string) *member {
	m.memberDo.UseTable(newTableName)
	return m.updateTableName(newTableName)
}

func (m member) As(alias string) *member {
	m.memberDo.DO = *(m.memberDo.As(alias).(*gen.DO))
	return m.updateTableName(alias)
}

func (m *member) updateTableName(table string) *member {
	m.ALL = field.NewAsterisk(table)
	m.ID = field.NewField(table, "uid")
	m.Username = field.NewString(table, "username")
	m.Nickname = field.NewString(table, "nickname")
	m.Avatar = field.NewString(table, "avatar")
	m.Groupid = field.NewUint8(table, "groupid")
	m.Regdate = field.NewInt64(table, "regdate")
	m.Lastvisit = field.NewUint32(table, "lastvisit")
	m.Lastactivity = field.NewUint32(table, "lastactivity")
	m.Lastpost = field.NewUint32(table, "lastpost")
	m.Dateformat = field.NewString(table, "dateformat")
	m.Timeformat = field.NewBool(table, "timeformat")
	m.Timeoffset = field.NewString(table, "timeoffset")
	m.Newpm = field.NewBool(table, "newpm")
	m.NewNotify = field.NewUint16(table, "new_notify")
	m.Sign = field.NewString(table, "sign")
	m.PasswordCrypt = field.NewBytes(table, "password_crypt")
	m.Email = field.NewString(table, "email")

	m.fillFieldMap()

	return m
}

func (m *member) WithContext(ctx context.Context) *memberDo { return m.memberDo.WithContext(ctx) }

func (m member) TableName() string { return m.memberDo.TableName() }

func (m member) Alias() string { return m.memberDo.Alias() }

func (m *member) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := m.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (m *member) fillFieldMap() {
	m.fieldMap = make(map[string]field.Expr, 18)
	m.fieldMap["uid"] = m.ID
	m.fieldMap["username"] = m.Username
	m.fieldMap["nickname"] = m.Nickname
	m.fieldMap["avatar"] = m.Avatar
	m.fieldMap["groupid"] = m.Groupid
	m.fieldMap["regdate"] = m.Regdate
	m.fieldMap["lastvisit"] = m.Lastvisit
	m.fieldMap["lastactivity"] = m.Lastactivity
	m.fieldMap["lastpost"] = m.Lastpost
	m.fieldMap["dateformat"] = m.Dateformat
	m.fieldMap["timeformat"] = m.Timeformat
	m.fieldMap["timeoffset"] = m.Timeoffset
	m.fieldMap["newpm"] = m.Newpm
	m.fieldMap["new_notify"] = m.NewNotify
	m.fieldMap["sign"] = m.Sign
	m.fieldMap["password_crypt"] = m.PasswordCrypt
	m.fieldMap["email"] = m.Email

}

func (m member) clone(db *gorm.DB) member {
	m.memberDo.ReplaceConnPool(db.Statement.ConnPool)
	return m
}

func (m member) replaceDB(db *gorm.DB) member {
	m.memberDo.ReplaceDB(db)
	return m
}

type memberHasOneFields struct {
	db *gorm.DB

	field.RelationField
}

func (a memberHasOneFields) Where(conds ...field.Expr) *memberHasOneFields {
	if len(conds) == 0 {
		return &a
	}

	exprs := make([]clause.Expression, 0, len(conds))
	for _, cond := range conds {
		exprs = append(exprs, cond.BeCond().(clause.Expression))
	}
	a.db = a.db.Clauses(clause.Where{Exprs: exprs})
	return &a
}

func (a memberHasOneFields) WithContext(ctx context.Context) *memberHasOneFields {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a memberHasOneFields) Model(m *dao.Member) *memberHasOneFieldsTx {
	return &memberHasOneFieldsTx{a.db.Model(m).Association(a.Name())}
}

type memberHasOneFieldsTx struct{ tx *gorm.Association }

func (a memberHasOneFieldsTx) Find() (result *dao.MemberField, err error) {
	return result, a.tx.Find(&result)
}

func (a memberHasOneFieldsTx) Append(values ...*dao.MemberField) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a memberHasOneFieldsTx) Replace(values ...*dao.MemberField) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a memberHasOneFieldsTx) Delete(values ...*dao.MemberField) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a memberHasOneFieldsTx) Clear() error {
	return a.tx.Clear()
}

func (a memberHasOneFieldsTx) Count() int64 {
	return a.tx.Count()
}

type memberDo struct{ gen.DO }

func (m memberDo) Debug() *memberDo {
	return m.withDO(m.DO.Debug())
}

func (m memberDo) WithContext(ctx context.Context) *memberDo {
	return m.withDO(m.DO.WithContext(ctx))
}

func (m memberDo) ReadDB() *memberDo {
	return m.Clauses(dbresolver.Read)
}

func (m memberDo) WriteDB() *memberDo {
	return m.Clauses(dbresolver.Write)
}

func (m memberDo) Session(config *gorm.Session) *memberDo {
	return m.withDO(m.DO.Session(config))
}

func (m memberDo) Clauses(conds ...clause.Expression) *memberDo {
	return m.withDO(m.DO.Clauses(conds...))
}

func (m memberDo) Returning(value interface{}, columns ...string) *memberDo {
	return m.withDO(m.DO.Returning(value, columns...))
}

func (m memberDo) Not(conds ...gen.Condition) *memberDo {
	return m.withDO(m.DO.Not(conds...))
}

func (m memberDo) Or(conds ...gen.Condition) *memberDo {
	return m.withDO(m.DO.Or(conds...))
}

func (m memberDo) Select(conds ...field.Expr) *memberDo {
	return m.withDO(m.DO.Select(conds...))
}

func (m memberDo) Where(conds ...gen.Condition) *memberDo {
	return m.withDO(m.DO.Where(conds...))
}

func (m memberDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *memberDo {
	return m.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (m memberDo) Order(conds ...field.Expr) *memberDo {
	return m.withDO(m.DO.Order(conds...))
}

func (m memberDo) Distinct(cols ...field.Expr) *memberDo {
	return m.withDO(m.DO.Distinct(cols...))
}

func (m memberDo) Omit(cols ...field.Expr) *memberDo {
	return m.withDO(m.DO.Omit(cols...))
}

func (m memberDo) Join(table schema.Tabler, on ...field.Expr) *memberDo {
	return m.withDO(m.DO.Join(table, on...))
}

func (m memberDo) LeftJoin(table schema.Tabler, on ...field.Expr) *memberDo {
	return m.withDO(m.DO.LeftJoin(table, on...))
}

func (m memberDo) RightJoin(table schema.Tabler, on ...field.Expr) *memberDo {
	return m.withDO(m.DO.RightJoin(table, on...))
}

func (m memberDo) Group(cols ...field.Expr) *memberDo {
	return m.withDO(m.DO.Group(cols...))
}

func (m memberDo) Having(conds ...gen.Condition) *memberDo {
	return m.withDO(m.DO.Having(conds...))
}

func (m memberDo) Limit(limit int) *memberDo {
	return m.withDO(m.DO.Limit(limit))
}

func (m memberDo) Offset(offset int) *memberDo {
	return m.withDO(m.DO.Offset(offset))
}

func (m memberDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *memberDo {
	return m.withDO(m.DO.Scopes(funcs...))
}

func (m memberDo) Unscoped() *memberDo {
	return m.withDO(m.DO.Unscoped())
}

func (m memberDo) Create(values ...*dao.Member) error {
	if len(values) == 0 {
		return nil
	}
	return m.DO.Create(values)
}

func (m memberDo) CreateInBatches(values []*dao.Member, batchSize int) error {
	return m.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (m memberDo) Save(values ...*dao.Member) error {
	if len(values) == 0 {
		return nil
	}
	return m.DO.Save(values)
}

func (m memberDo) First() (*dao.Member, error) {
	if result, err := m.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*dao.Member), nil
	}
}

func (m memberDo) Take() (*dao.Member, error) {
	if result, err := m.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*dao.Member), nil
	}
}

func (m memberDo) Last() (*dao.Member, error) {
	if result, err := m.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*dao.Member), nil
	}
}

func (m memberDo) Find() ([]*dao.Member, error) {
	result, err := m.DO.Find()
	return result.([]*dao.Member), err
}

func (m memberDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*dao.Member, err error) {
	buf := make([]*dao.Member, 0, batchSize)
	err = m.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (m memberDo) FindInBatches(result *[]*dao.Member, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return m.DO.FindInBatches(result, batchSize, fc)
}

func (m memberDo) Attrs(attrs ...field.AssignExpr) *memberDo {
	return m.withDO(m.DO.Attrs(attrs...))
}

func (m memberDo) Assign(attrs ...field.AssignExpr) *memberDo {
	return m.withDO(m.DO.Assign(attrs...))
}

func (m memberDo) Joins(fields ...field.RelationField) *memberDo {
	for _, _f := range fields {
		m = *m.withDO(m.DO.Joins(_f))
	}
	return &m
}

func (m memberDo) Preload(fields ...field.RelationField) *memberDo {
	for _, _f := range fields {
		m = *m.withDO(m.DO.Preload(_f))
	}
	return &m
}

func (m memberDo) FirstOrInit() (*dao.Member, error) {
	if result, err := m.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*dao.Member), nil
	}
}

func (m memberDo) FirstOrCreate() (*dao.Member, error) {
	if result, err := m.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*dao.Member), nil
	}
}

func (m memberDo) FindByPage(offset int, limit int) (result []*dao.Member, count int64, err error) {
	result, err = m.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = m.Offset(-1).Limit(-1).Count()
	return
}

func (m memberDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = m.Count()
	if err != nil {
		return
	}

	err = m.Offset(offset).Limit(limit).Scan(result)
	return
}

func (m memberDo) Scan(result interface{}) (err error) {
	return m.DO.Scan(result)
}

func (m memberDo) Delete(models ...*dao.Member) (result gen.ResultInfo, err error) {
	return m.DO.Delete(models)
}

func (m *memberDo) withDO(do gen.Dao) *memberDo {
	m.DO = *do.(*gen.DO)
	return m
}
