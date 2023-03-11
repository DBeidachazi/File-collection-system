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

func newRole(db *gorm.DB, opts ...gen.DOOption) role {
	_role := role{}

	_role.roleDo.UseDB(db, opts...)
	_role.roleDo.UseModel(&model.Role{})

	tableName := _role.roleDo.TableName()
	_role.ALL = field.NewAsterisk(tableName)
	_role.StuID = field.NewInt32(tableName, "stu_id")
	_role.RoleID = field.NewInt32(tableName, "role_id")
	_role.RoleName = field.NewString(tableName, "role_name")
	_role.ClassID = field.NewInt32(tableName, "class_id")
	_role.Status = field.NewInt32(tableName, "status")

	_role.fillFieldMap()

	return _role
}

type role struct {
	roleDo

	ALL      field.Asterisk
	StuID    field.Int32
	RoleID   field.Int32
	RoleName field.String
	ClassID  field.Int32
	Status   field.Int32 // 1:已提交 2:未提交 3:发布者(不显示)

	fieldMap map[string]field.Expr
}

func (r role) Table(newTableName string) *role {
	r.roleDo.UseTable(newTableName)
	return r.updateTableName(newTableName)
}

func (r role) As(alias string) *role {
	r.roleDo.DO = *(r.roleDo.As(alias).(*gen.DO))
	return r.updateTableName(alias)
}

func (r *role) updateTableName(table string) *role {
	r.ALL = field.NewAsterisk(table)
	r.StuID = field.NewInt32(table, "stu_id")
	r.RoleID = field.NewInt32(table, "role_id")
	r.RoleName = field.NewString(table, "role_name")
	r.ClassID = field.NewInt32(table, "class_id")
	r.Status = field.NewInt32(table, "status")

	r.fillFieldMap()

	return r
}

func (r *role) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := r.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (r *role) fillFieldMap() {
	r.fieldMap = make(map[string]field.Expr, 5)
	r.fieldMap["stu_id"] = r.StuID
	r.fieldMap["role_id"] = r.RoleID
	r.fieldMap["role_name"] = r.RoleName
	r.fieldMap["class_id"] = r.ClassID
	r.fieldMap["status"] = r.Status
}

func (r role) clone(db *gorm.DB) role {
	r.roleDo.ReplaceConnPool(db.Statement.ConnPool)
	return r
}

func (r role) replaceDB(db *gorm.DB) role {
	r.roleDo.ReplaceDB(db)
	return r
}

type roleDo struct{ gen.DO }

func (r roleDo) Debug() *roleDo {
	return r.withDO(r.DO.Debug())
}

func (r roleDo) WithContext(ctx context.Context) *roleDo {
	return r.withDO(r.DO.WithContext(ctx))
}

func (r roleDo) ReadDB() *roleDo {
	return r.Clauses(dbresolver.Read)
}

func (r roleDo) WriteDB() *roleDo {
	return r.Clauses(dbresolver.Write)
}

func (r roleDo) Session(config *gorm.Session) *roleDo {
	return r.withDO(r.DO.Session(config))
}

func (r roleDo) Clauses(conds ...clause.Expression) *roleDo {
	return r.withDO(r.DO.Clauses(conds...))
}

func (r roleDo) Returning(value interface{}, columns ...string) *roleDo {
	return r.withDO(r.DO.Returning(value, columns...))
}

func (r roleDo) Not(conds ...gen.Condition) *roleDo {
	return r.withDO(r.DO.Not(conds...))
}

func (r roleDo) Or(conds ...gen.Condition) *roleDo {
	return r.withDO(r.DO.Or(conds...))
}

func (r roleDo) Select(conds ...field.Expr) *roleDo {
	return r.withDO(r.DO.Select(conds...))
}

func (r roleDo) Where(conds ...gen.Condition) *roleDo {
	return r.withDO(r.DO.Where(conds...))
}

func (r roleDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *roleDo {
	return r.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (r roleDo) Order(conds ...field.Expr) *roleDo {
	return r.withDO(r.DO.Order(conds...))
}

func (r roleDo) Distinct(cols ...field.Expr) *roleDo {
	return r.withDO(r.DO.Distinct(cols...))
}

func (r roleDo) Omit(cols ...field.Expr) *roleDo {
	return r.withDO(r.DO.Omit(cols...))
}

func (r roleDo) Join(table schema.Tabler, on ...field.Expr) *roleDo {
	return r.withDO(r.DO.Join(table, on...))
}

func (r roleDo) LeftJoin(table schema.Tabler, on ...field.Expr) *roleDo {
	return r.withDO(r.DO.LeftJoin(table, on...))
}

func (r roleDo) RightJoin(table schema.Tabler, on ...field.Expr) *roleDo {
	return r.withDO(r.DO.RightJoin(table, on...))
}

func (r roleDo) Group(cols ...field.Expr) *roleDo {
	return r.withDO(r.DO.Group(cols...))
}

func (r roleDo) Having(conds ...gen.Condition) *roleDo {
	return r.withDO(r.DO.Having(conds...))
}

func (r roleDo) Limit(limit int) *roleDo {
	return r.withDO(r.DO.Limit(limit))
}

func (r roleDo) Offset(offset int) *roleDo {
	return r.withDO(r.DO.Offset(offset))
}

func (r roleDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *roleDo {
	return r.withDO(r.DO.Scopes(funcs...))
}

func (r roleDo) Unscoped() *roleDo {
	return r.withDO(r.DO.Unscoped())
}

func (r roleDo) Create(values ...*model.Role) error {
	if len(values) == 0 {
		return nil
	}
	return r.DO.Create(values)
}

func (r roleDo) CreateInBatches(values []*model.Role, batchSize int) error {
	return r.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (r roleDo) Save(values ...*model.Role) error {
	if len(values) == 0 {
		return nil
	}
	return r.DO.Save(values)
}

func (r roleDo) First() (*model.Role, error) {
	if result, err := r.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Role), nil
	}
}

func (r roleDo) Take() (*model.Role, error) {
	if result, err := r.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Role), nil
	}
}

func (r roleDo) Last() (*model.Role, error) {
	if result, err := r.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Role), nil
	}
}

func (r roleDo) Find() ([]*model.Role, error) {
	result, err := r.DO.Find()
	return result.([]*model.Role), err
}

func (r roleDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Role, err error) {
	buf := make([]*model.Role, 0, batchSize)
	err = r.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (r roleDo) FindInBatches(result *[]*model.Role, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return r.DO.FindInBatches(result, batchSize, fc)
}

func (r roleDo) Attrs(attrs ...field.AssignExpr) *roleDo {
	return r.withDO(r.DO.Attrs(attrs...))
}

func (r roleDo) Assign(attrs ...field.AssignExpr) *roleDo {
	return r.withDO(r.DO.Assign(attrs...))
}

func (r roleDo) Joins(fields ...field.RelationField) *roleDo {
	for _, _f := range fields {
		r = *r.withDO(r.DO.Joins(_f))
	}
	return &r
}

func (r roleDo) Preload(fields ...field.RelationField) *roleDo {
	for _, _f := range fields {
		r = *r.withDO(r.DO.Preload(_f))
	}
	return &r
}

func (r roleDo) FirstOrInit() (*model.Role, error) {
	if result, err := r.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Role), nil
	}
}

func (r roleDo) FirstOrCreate() (*model.Role, error) {
	if result, err := r.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Role), nil
	}
}

func (r roleDo) FindByPage(offset int, limit int) (result []*model.Role, count int64, err error) {
	result, err = r.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = r.Offset(-1).Limit(-1).Count()
	return
}

func (r roleDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = r.Count()
	if err != nil {
		return
	}

	err = r.Offset(offset).Limit(limit).Scan(result)
	return
}

func (r roleDo) Scan(result interface{}) (err error) {
	return r.DO.Scan(result)
}

func (r roleDo) Delete(models ...*model.Role) (result gen.ResultInfo, err error) {
	return r.DO.Delete(models)
}

func (r *roleDo) withDO(do gen.Dao) *roleDo {
	r.DO = *do.(*gen.DO)
	return r
}
