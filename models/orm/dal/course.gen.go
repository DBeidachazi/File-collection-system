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

func newCourse(db *gorm.DB, opts ...gen.DOOption) course {
	_course := course{}

	_course.courseDo.UseDB(db, opts...)
	_course.courseDo.UseModel(&model.Course{})

	tableName := _course.courseDo.TableName()
	_course.ALL = field.NewAsterisk(tableName)
	_course.CourseID = field.NewInt32(tableName, "course_id")
	_course.Name = field.NewString(tableName, "name")
	_course.ClassID = field.NewInt32(tableName, "class_id")
	_course.PublisherName = field.NewString(tableName, "publisher_name")

	_course.fillFieldMap()

	return _course
}

type course struct {
	courseDo

	ALL           field.Asterisk
	CourseID      field.Int32
	Name          field.String
	ClassID       field.Int32
	PublisherName field.String

	fieldMap map[string]field.Expr
}

func (c course) Table(newTableName string) *course {
	c.courseDo.UseTable(newTableName)
	return c.updateTableName(newTableName)
}

func (c course) As(alias string) *course {
	c.courseDo.DO = *(c.courseDo.As(alias).(*gen.DO))
	return c.updateTableName(alias)
}

func (c *course) updateTableName(table string) *course {
	c.ALL = field.NewAsterisk(table)
	c.CourseID = field.NewInt32(table, "course_id")
	c.Name = field.NewString(table, "name")
	c.ClassID = field.NewInt32(table, "class_id")
	c.PublisherName = field.NewString(table, "publisher_name")

	c.fillFieldMap()

	return c
}

func (c *course) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := c.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (c *course) fillFieldMap() {
	c.fieldMap = make(map[string]field.Expr, 4)
	c.fieldMap["course_id"] = c.CourseID
	c.fieldMap["name"] = c.Name
	c.fieldMap["class_id"] = c.ClassID
	c.fieldMap["publisher_name"] = c.PublisherName
}

func (c course) clone(db *gorm.DB) course {
	c.courseDo.ReplaceConnPool(db.Statement.ConnPool)
	return c
}

func (c course) replaceDB(db *gorm.DB) course {
	c.courseDo.ReplaceDB(db)
	return c
}

type courseDo struct{ gen.DO }

func (c courseDo) Debug() *courseDo {
	return c.withDO(c.DO.Debug())
}

func (c courseDo) WithContext(ctx context.Context) *courseDo {
	return c.withDO(c.DO.WithContext(ctx))
}

func (c courseDo) ReadDB() *courseDo {
	return c.Clauses(dbresolver.Read)
}

func (c courseDo) WriteDB() *courseDo {
	return c.Clauses(dbresolver.Write)
}

func (c courseDo) Session(config *gorm.Session) *courseDo {
	return c.withDO(c.DO.Session(config))
}

func (c courseDo) Clauses(conds ...clause.Expression) *courseDo {
	return c.withDO(c.DO.Clauses(conds...))
}

func (c courseDo) Returning(value interface{}, columns ...string) *courseDo {
	return c.withDO(c.DO.Returning(value, columns...))
}

func (c courseDo) Not(conds ...gen.Condition) *courseDo {
	return c.withDO(c.DO.Not(conds...))
}

func (c courseDo) Or(conds ...gen.Condition) *courseDo {
	return c.withDO(c.DO.Or(conds...))
}

func (c courseDo) Select(conds ...field.Expr) *courseDo {
	return c.withDO(c.DO.Select(conds...))
}

func (c courseDo) Where(conds ...gen.Condition) *courseDo {
	return c.withDO(c.DO.Where(conds...))
}

func (c courseDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *courseDo {
	return c.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (c courseDo) Order(conds ...field.Expr) *courseDo {
	return c.withDO(c.DO.Order(conds...))
}

func (c courseDo) Distinct(cols ...field.Expr) *courseDo {
	return c.withDO(c.DO.Distinct(cols...))
}

func (c courseDo) Omit(cols ...field.Expr) *courseDo {
	return c.withDO(c.DO.Omit(cols...))
}

func (c courseDo) Join(table schema.Tabler, on ...field.Expr) *courseDo {
	return c.withDO(c.DO.Join(table, on...))
}

func (c courseDo) LeftJoin(table schema.Tabler, on ...field.Expr) *courseDo {
	return c.withDO(c.DO.LeftJoin(table, on...))
}

func (c courseDo) RightJoin(table schema.Tabler, on ...field.Expr) *courseDo {
	return c.withDO(c.DO.RightJoin(table, on...))
}

func (c courseDo) Group(cols ...field.Expr) *courseDo {
	return c.withDO(c.DO.Group(cols...))
}

func (c courseDo) Having(conds ...gen.Condition) *courseDo {
	return c.withDO(c.DO.Having(conds...))
}

func (c courseDo) Limit(limit int) *courseDo {
	return c.withDO(c.DO.Limit(limit))
}

func (c courseDo) Offset(offset int) *courseDo {
	return c.withDO(c.DO.Offset(offset))
}

func (c courseDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *courseDo {
	return c.withDO(c.DO.Scopes(funcs...))
}

func (c courseDo) Unscoped() *courseDo {
	return c.withDO(c.DO.Unscoped())
}

func (c courseDo) Create(values ...*model.Course) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Create(values)
}

func (c courseDo) CreateInBatches(values []*model.Course, batchSize int) error {
	return c.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (c courseDo) Save(values ...*model.Course) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Save(values)
}

func (c courseDo) First() (*model.Course, error) {
	if result, err := c.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Course), nil
	}
}

func (c courseDo) Take() (*model.Course, error) {
	if result, err := c.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Course), nil
	}
}

func (c courseDo) Last() (*model.Course, error) {
	if result, err := c.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Course), nil
	}
}

func (c courseDo) Find() ([]*model.Course, error) {
	result, err := c.DO.Find()
	return result.([]*model.Course), err
}

func (c courseDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Course, err error) {
	buf := make([]*model.Course, 0, batchSize)
	err = c.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (c courseDo) FindInBatches(result *[]*model.Course, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return c.DO.FindInBatches(result, batchSize, fc)
}

func (c courseDo) Attrs(attrs ...field.AssignExpr) *courseDo {
	return c.withDO(c.DO.Attrs(attrs...))
}

func (c courseDo) Assign(attrs ...field.AssignExpr) *courseDo {
	return c.withDO(c.DO.Assign(attrs...))
}

func (c courseDo) Joins(fields ...field.RelationField) *courseDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Joins(_f))
	}
	return &c
}

func (c courseDo) Preload(fields ...field.RelationField) *courseDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Preload(_f))
	}
	return &c
}

func (c courseDo) FirstOrInit() (*model.Course, error) {
	if result, err := c.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Course), nil
	}
}

func (c courseDo) FirstOrCreate() (*model.Course, error) {
	if result, err := c.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Course), nil
	}
}

func (c courseDo) FindByPage(offset int, limit int) (result []*model.Course, count int64, err error) {
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

func (c courseDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = c.Count()
	if err != nil {
		return
	}

	err = c.Offset(offset).Limit(limit).Scan(result)
	return
}

func (c courseDo) Scan(result interface{}) (err error) {
	return c.DO.Scan(result)
}

func (c courseDo) Delete(models ...*model.Course) (result gen.ResultInfo, err error) {
	return c.DO.Delete(models)
}

func (c *courseDo) withDO(do gen.Dao) *courseDo {
	c.DO = *do.(*gen.DO)
	return c
}
