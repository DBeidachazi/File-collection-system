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

	"FengfengStudy/models/orm/model"
)

func newClassUser(db *gorm.DB, opts ...gen.DOOption) classUser {
	_classUser := classUser{}

	_classUser.classUserDo.UseDB(db, opts...)
	_classUser.classUserDo.UseModel(&model.ClassUser{})

	tableName := _classUser.classUserDo.TableName()
	_classUser.ALL = field.NewAsterisk(tableName)
	_classUser.ClassID = field.NewInt32(tableName, "class_id")
	_classUser.StuID = field.NewInt32(tableName, "stu_id")
	_classUser.Name = field.NewString(tableName, "name")
	_classUser.PermissionType = field.NewInt32(tableName, "permission_type")

	_classUser.fillFieldMap()

	return _classUser
}

type classUser struct {
	classUserDo

	ALL            field.Asterisk
	ClassID        field.Int32
	StuID          field.Int32
	Name           field.String
	PermissionType field.Int32 // 1:管理员 2:普通用户

	fieldMap map[string]field.Expr
}

func (c classUser) Table(newTableName string) *classUser {
	c.classUserDo.UseTable(newTableName)
	return c.updateTableName(newTableName)
}

func (c classUser) As(alias string) *classUser {
	c.classUserDo.DO = *(c.classUserDo.As(alias).(*gen.DO))
	return c.updateTableName(alias)
}

func (c *classUser) updateTableName(table string) *classUser {
	c.ALL = field.NewAsterisk(table)
	c.ClassID = field.NewInt32(table, "class_id")
	c.StuID = field.NewInt32(table, "stu_id")
	c.Name = field.NewString(table, "name")
	c.PermissionType = field.NewInt32(table, "permission_type")

	c.fillFieldMap()

	return c
}

func (c *classUser) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := c.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (c *classUser) fillFieldMap() {
	c.fieldMap = make(map[string]field.Expr, 4)
	c.fieldMap["class_id"] = c.ClassID
	c.fieldMap["stu_id"] = c.StuID
	c.fieldMap["name"] = c.Name
	c.fieldMap["permission_type"] = c.PermissionType
}

func (c classUser) clone(db *gorm.DB) classUser {
	c.classUserDo.ReplaceConnPool(db.Statement.ConnPool)
	return c
}

func (c classUser) replaceDB(db *gorm.DB) classUser {
	c.classUserDo.ReplaceDB(db)
	return c
}

type classUserDo struct{ gen.DO }

func (c classUserDo) Debug() *classUserDo {
	return c.withDO(c.DO.Debug())
}

func (c classUserDo) WithContext(ctx context.Context) *classUserDo {
	return c.withDO(c.DO.WithContext(ctx))
}

func (c classUserDo) ReadDB() *classUserDo {
	return c.Clauses(dbresolver.Read)
}

func (c classUserDo) WriteDB() *classUserDo {
	return c.Clauses(dbresolver.Write)
}

func (c classUserDo) Session(config *gorm.Session) *classUserDo {
	return c.withDO(c.DO.Session(config))
}

func (c classUserDo) Clauses(conds ...clause.Expression) *classUserDo {
	return c.withDO(c.DO.Clauses(conds...))
}

func (c classUserDo) Returning(value interface{}, columns ...string) *classUserDo {
	return c.withDO(c.DO.Returning(value, columns...))
}

func (c classUserDo) Not(conds ...gen.Condition) *classUserDo {
	return c.withDO(c.DO.Not(conds...))
}

func (c classUserDo) Or(conds ...gen.Condition) *classUserDo {
	return c.withDO(c.DO.Or(conds...))
}

func (c classUserDo) Select(conds ...field.Expr) *classUserDo {
	return c.withDO(c.DO.Select(conds...))
}

func (c classUserDo) Where(conds ...gen.Condition) *classUserDo {
	return c.withDO(c.DO.Where(conds...))
}

func (c classUserDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *classUserDo {
	return c.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (c classUserDo) Order(conds ...field.Expr) *classUserDo {
	return c.withDO(c.DO.Order(conds...))
}

func (c classUserDo) Distinct(cols ...field.Expr) *classUserDo {
	return c.withDO(c.DO.Distinct(cols...))
}

func (c classUserDo) Omit(cols ...field.Expr) *classUserDo {
	return c.withDO(c.DO.Omit(cols...))
}

func (c classUserDo) Join(table schema.Tabler, on ...field.Expr) *classUserDo {
	return c.withDO(c.DO.Join(table, on...))
}

func (c classUserDo) LeftJoin(table schema.Tabler, on ...field.Expr) *classUserDo {
	return c.withDO(c.DO.LeftJoin(table, on...))
}

func (c classUserDo) RightJoin(table schema.Tabler, on ...field.Expr) *classUserDo {
	return c.withDO(c.DO.RightJoin(table, on...))
}

func (c classUserDo) Group(cols ...field.Expr) *classUserDo {
	return c.withDO(c.DO.Group(cols...))
}

func (c classUserDo) Having(conds ...gen.Condition) *classUserDo {
	return c.withDO(c.DO.Having(conds...))
}

func (c classUserDo) Limit(limit int) *classUserDo {
	return c.withDO(c.DO.Limit(limit))
}

func (c classUserDo) Offset(offset int) *classUserDo {
	return c.withDO(c.DO.Offset(offset))
}

func (c classUserDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *classUserDo {
	return c.withDO(c.DO.Scopes(funcs...))
}

func (c classUserDo) Unscoped() *classUserDo {
	return c.withDO(c.DO.Unscoped())
}

func (c classUserDo) Create(values ...*model.ClassUser) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Create(values)
}

func (c classUserDo) CreateInBatches(values []*model.ClassUser, batchSize int) error {
	return c.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (c classUserDo) Save(values ...*model.ClassUser) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Save(values)
}

func (c classUserDo) First() (*model.ClassUser, error) {
	if result, err := c.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.ClassUser), nil
	}
}

func (c classUserDo) Take() (*model.ClassUser, error) {
	if result, err := c.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.ClassUser), nil
	}
}

func (c classUserDo) Last() (*model.ClassUser, error) {
	if result, err := c.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.ClassUser), nil
	}
}

func (c classUserDo) Find() ([]*model.ClassUser, error) {
	result, err := c.DO.Find()
	return result.([]*model.ClassUser), err
}

func (c classUserDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.ClassUser, err error) {
	buf := make([]*model.ClassUser, 0, batchSize)
	err = c.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (c classUserDo) FindInBatches(result *[]*model.ClassUser, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return c.DO.FindInBatches(result, batchSize, fc)
}

func (c classUserDo) Attrs(attrs ...field.AssignExpr) *classUserDo {
	return c.withDO(c.DO.Attrs(attrs...))
}

func (c classUserDo) Assign(attrs ...field.AssignExpr) *classUserDo {
	return c.withDO(c.DO.Assign(attrs...))
}

func (c classUserDo) Joins(fields ...field.RelationField) *classUserDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Joins(_f))
	}
	return &c
}

func (c classUserDo) Preload(fields ...field.RelationField) *classUserDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Preload(_f))
	}
	return &c
}

func (c classUserDo) FirstOrInit() (*model.ClassUser, error) {
	if result, err := c.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.ClassUser), nil
	}
}

func (c classUserDo) FirstOrCreate() (*model.ClassUser, error) {
	if result, err := c.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.ClassUser), nil
	}
}

func (c classUserDo) FindByPage(offset int, limit int) (result []*model.ClassUser, count int64, err error) {
	result, err = c.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = c.Offset(-1).Limit(-1).Count()
	return
}

func (c classUserDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = c.Count()
	if err != nil {
		return
	}

	err = c.Offset(offset).Limit(limit).Scan(result)
	return
}

func (c classUserDo) Scan(result interface{}) (err error) {
	return c.DO.Scan(result)
}

func (c classUserDo) Delete(models ...*model.ClassUser) (result gen.ResultInfo, err error) {
	return c.DO.Delete(models)
}

func (c *classUserDo) withDO(do gen.Dao) *classUserDo {
	c.DO = *do.(*gen.DO)
	return c
}
