// Code generated by go-queryset. DO NOT EDIT.
package models

import (
	"errors"
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
)

// ===== BEGIN of all query sets

// ===== BEGIN of query set FunctionControlQuerySet

// FunctionControlQuerySet is an queryset type for FunctionControl
type FunctionControlQuerySet struct {
	db *gorm.DB
}

// NewFunctionControlQuerySet constructs new FunctionControlQuerySet
func NewFunctionControlQuerySet(db *gorm.DB) FunctionControlQuerySet {
	return FunctionControlQuerySet{
		db: db.Model(&FunctionControl{}),
	}
}

func (qs FunctionControlQuerySet) w(db *gorm.DB) FunctionControlQuerySet {
	return NewFunctionControlQuerySet(db)
}

// Create is an autogenerated method
// nolint: dupl
func (o *FunctionControl) Create(db *gorm.DB) error {
	return db.Create(o).Error
}

// Delete is an autogenerated method
// nolint: dupl
func (o *FunctionControl) Delete(db *gorm.DB) error {
	return db.Delete(o).Error
}

// All is an autogenerated method
// nolint: dupl
func (qs FunctionControlQuerySet) All(ret *[]FunctionControl) error {
	return qs.db.Find(ret).Error
}

// Count is an autogenerated method
// nolint: dupl
func (qs FunctionControlQuerySet) Count() (int, error) {
	var count int
	err := qs.db.Count(&count).Error
	return count, err
}

// CreatedAtEq is an autogenerated method
// nolint: dupl
func (qs FunctionControlQuerySet) CreatedAtEq(createdAt time.Time) FunctionControlQuerySet {
	return qs.w(qs.db.Where("created_at = ?", createdAt))
}

// CreatedAtGt is an autogenerated method
// nolint: dupl
func (qs FunctionControlQuerySet) CreatedAtGt(createdAt time.Time) FunctionControlQuerySet {
	return qs.w(qs.db.Where("created_at > ?", createdAt))
}

// CreatedAtGte is an autogenerated method
// nolint: dupl
func (qs FunctionControlQuerySet) CreatedAtGte(createdAt time.Time) FunctionControlQuerySet {
	return qs.w(qs.db.Where("created_at >= ?", createdAt))
}

// CreatedAtLt is an autogenerated method
// nolint: dupl
func (qs FunctionControlQuerySet) CreatedAtLt(createdAt time.Time) FunctionControlQuerySet {
	return qs.w(qs.db.Where("created_at < ?", createdAt))
}

// CreatedAtLte is an autogenerated method
// nolint: dupl
func (qs FunctionControlQuerySet) CreatedAtLte(createdAt time.Time) FunctionControlQuerySet {
	return qs.w(qs.db.Where("created_at <= ?", createdAt))
}

// CreatedAtNe is an autogenerated method
// nolint: dupl
func (qs FunctionControlQuerySet) CreatedAtNe(createdAt time.Time) FunctionControlQuerySet {
	return qs.w(qs.db.Where("created_at != ?", createdAt))
}

// CreatorEq is an autogenerated method
// nolint: dupl
func (qs FunctionControlQuerySet) CreatorEq(creator string) FunctionControlQuerySet {
	return qs.w(qs.db.Where("creator = ?", creator))
}

// CreatorIn is an autogenerated method
// nolint: dupl
func (qs FunctionControlQuerySet) CreatorIn(creator ...string) FunctionControlQuerySet {
	if len(creator) == 0 {
		qs.db.AddError(errors.New("must at least pass one creator in CreatorIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("creator IN (?)", creator))
}

// CreatorNe is an autogenerated method
// nolint: dupl
func (qs FunctionControlQuerySet) CreatorNe(creator string) FunctionControlQuerySet {
	return qs.w(qs.db.Where("creator != ?", creator))
}

// CreatorNotIn is an autogenerated method
// nolint: dupl
func (qs FunctionControlQuerySet) CreatorNotIn(creator ...string) FunctionControlQuerySet {
	if len(creator) == 0 {
		qs.db.AddError(errors.New("must at least pass one creator in CreatorNotIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("creator NOT IN (?)", creator))
}

// Delete is an autogenerated method
// nolint: dupl
func (qs FunctionControlQuerySet) Delete() error {
	return qs.db.Delete(FunctionControl{}).Error
}

// DeleteNum is an autogenerated method
// nolint: dupl
func (qs FunctionControlQuerySet) DeleteNum() (int64, error) {
	db := qs.db.Delete(FunctionControl{})
	return db.RowsAffected, db.Error
}

// DeleteNumUnscoped is an autogenerated method
// nolint: dupl
func (qs FunctionControlQuerySet) DeleteNumUnscoped() (int64, error) {
	db := qs.db.Unscoped().Delete(FunctionControl{})
	return db.RowsAffected, db.Error
}

// DeletedAtEq is an autogenerated method
// nolint: dupl
func (qs FunctionControlQuerySet) DeletedAtEq(deletedAt time.Time) FunctionControlQuerySet {
	return qs.w(qs.db.Where("deleted_at = ?", deletedAt))
}

// DeletedAtGt is an autogenerated method
// nolint: dupl
func (qs FunctionControlQuerySet) DeletedAtGt(deletedAt time.Time) FunctionControlQuerySet {
	return qs.w(qs.db.Where("deleted_at > ?", deletedAt))
}

// DeletedAtGte is an autogenerated method
// nolint: dupl
func (qs FunctionControlQuerySet) DeletedAtGte(deletedAt time.Time) FunctionControlQuerySet {
	return qs.w(qs.db.Where("deleted_at >= ?", deletedAt))
}

// DeletedAtIsNotNull is an autogenerated method
// nolint: dupl
func (qs FunctionControlQuerySet) DeletedAtIsNotNull() FunctionControlQuerySet {
	return qs.w(qs.db.Where("deleted_at IS NOT NULL"))
}

// DeletedAtIsNull is an autogenerated method
// nolint: dupl
func (qs FunctionControlQuerySet) DeletedAtIsNull() FunctionControlQuerySet {
	return qs.w(qs.db.Where("deleted_at IS NULL"))
}

// DeletedAtLt is an autogenerated method
// nolint: dupl
func (qs FunctionControlQuerySet) DeletedAtLt(deletedAt time.Time) FunctionControlQuerySet {
	return qs.w(qs.db.Where("deleted_at < ?", deletedAt))
}

// DeletedAtLte is an autogenerated method
// nolint: dupl
func (qs FunctionControlQuerySet) DeletedAtLte(deletedAt time.Time) FunctionControlQuerySet {
	return qs.w(qs.db.Where("deleted_at <= ?", deletedAt))
}

// DeletedAtNe is an autogenerated method
// nolint: dupl
func (qs FunctionControlQuerySet) DeletedAtNe(deletedAt time.Time) FunctionControlQuerySet {
	return qs.w(qs.db.Where("deleted_at != ?", deletedAt))
}

// DescriptionEq is an autogenerated method
// nolint: dupl
func (qs FunctionControlQuerySet) DescriptionEq(description string) FunctionControlQuerySet {
	return qs.w(qs.db.Where("description = ?", description))
}

// DescriptionIn is an autogenerated method
// nolint: dupl
func (qs FunctionControlQuerySet) DescriptionIn(description ...string) FunctionControlQuerySet {
	if len(description) == 0 {
		qs.db.AddError(errors.New("must at least pass one description in DescriptionIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("description IN (?)", description))
}

// DescriptionNe is an autogenerated method
// nolint: dupl
func (qs FunctionControlQuerySet) DescriptionNe(description string) FunctionControlQuerySet {
	return qs.w(qs.db.Where("description != ?", description))
}

// DescriptionNotIn is an autogenerated method
// nolint: dupl
func (qs FunctionControlQuerySet) DescriptionNotIn(description ...string) FunctionControlQuerySet {
	if len(description) == 0 {
		qs.db.AddError(errors.New("must at least pass one description in DescriptionNotIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("description NOT IN (?)", description))
}

// ExtraEq is an autogenerated method
// nolint: dupl
func (qs FunctionControlQuerySet) ExtraEq(extra string) FunctionControlQuerySet {
	return qs.w(qs.db.Where("extra = ?", extra))
}

// ExtraIn is an autogenerated method
// nolint: dupl
func (qs FunctionControlQuerySet) ExtraIn(extra ...string) FunctionControlQuerySet {
	if len(extra) == 0 {
		qs.db.AddError(errors.New("must at least pass one extra in ExtraIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("extra IN (?)", extra))
}

// ExtraNe is an autogenerated method
// nolint: dupl
func (qs FunctionControlQuerySet) ExtraNe(extra string) FunctionControlQuerySet {
	return qs.w(qs.db.Where("extra != ?", extra))
}

// ExtraNotIn is an autogenerated method
// nolint: dupl
func (qs FunctionControlQuerySet) ExtraNotIn(extra ...string) FunctionControlQuerySet {
	if len(extra) == 0 {
		qs.db.AddError(errors.New("must at least pass one extra in ExtraNotIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("extra NOT IN (?)", extra))
}

// FuncCodeEq is an autogenerated method
// nolint: dupl
func (qs FunctionControlQuerySet) FuncCodeEq(funcCode string) FunctionControlQuerySet {
	return qs.w(qs.db.Where("func_code = ?", funcCode))
}

// FuncCodeIn is an autogenerated method
// nolint: dupl
func (qs FunctionControlQuerySet) FuncCodeIn(funcCode ...string) FunctionControlQuerySet {
	if len(funcCode) == 0 {
		qs.db.AddError(errors.New("must at least pass one funcCode in FuncCodeIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("func_code IN (?)", funcCode))
}

// FuncCodeNe is an autogenerated method
// nolint: dupl
func (qs FunctionControlQuerySet) FuncCodeNe(funcCode string) FunctionControlQuerySet {
	return qs.w(qs.db.Where("func_code != ?", funcCode))
}

// FuncCodeNotIn is an autogenerated method
// nolint: dupl
func (qs FunctionControlQuerySet) FuncCodeNotIn(funcCode ...string) FunctionControlQuerySet {
	if len(funcCode) == 0 {
		qs.db.AddError(errors.New("must at least pass one funcCode in FuncCodeNotIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("func_code NOT IN (?)", funcCode))
}

// FuncStatusEq is an autogenerated method
// nolint: dupl
func (qs FunctionControlQuerySet) FuncStatusEq(funcStatus bool) FunctionControlQuerySet {
	return qs.w(qs.db.Where("func_status = ?", funcStatus))
}

// FuncStatusIn is an autogenerated method
// nolint: dupl
func (qs FunctionControlQuerySet) FuncStatusIn(funcStatus ...bool) FunctionControlQuerySet {
	if len(funcStatus) == 0 {
		qs.db.AddError(errors.New("must at least pass one funcStatus in FuncStatusIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("func_status IN (?)", funcStatus))
}

// FuncStatusNe is an autogenerated method
// nolint: dupl
func (qs FunctionControlQuerySet) FuncStatusNe(funcStatus bool) FunctionControlQuerySet {
	return qs.w(qs.db.Where("func_status != ?", funcStatus))
}

// FuncStatusNotIn is an autogenerated method
// nolint: dupl
func (qs FunctionControlQuerySet) FuncStatusNotIn(funcStatus ...bool) FunctionControlQuerySet {
	if len(funcStatus) == 0 {
		qs.db.AddError(errors.New("must at least pass one funcStatus in FuncStatusNotIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("func_status NOT IN (?)", funcStatus))
}

// GetDB is an autogenerated method
// nolint: dupl
func (qs FunctionControlQuerySet) GetDB() *gorm.DB {
	return qs.db
}

// GetUpdater is an autogenerated method
// nolint: dupl
func (qs FunctionControlQuerySet) GetUpdater() FunctionControlUpdater {
	return NewFunctionControlUpdater(qs.db)
}

// IDEq is an autogenerated method
// nolint: dupl
func (qs FunctionControlQuerySet) IDEq(ID uint) FunctionControlQuerySet {
	return qs.w(qs.db.Where("id = ?", ID))
}

// IDGt is an autogenerated method
// nolint: dupl
func (qs FunctionControlQuerySet) IDGt(ID uint) FunctionControlQuerySet {
	return qs.w(qs.db.Where("id > ?", ID))
}

// IDGte is an autogenerated method
// nolint: dupl
func (qs FunctionControlQuerySet) IDGte(ID uint) FunctionControlQuerySet {
	return qs.w(qs.db.Where("id >= ?", ID))
}

// IDIn is an autogenerated method
// nolint: dupl
func (qs FunctionControlQuerySet) IDIn(ID ...uint) FunctionControlQuerySet {
	if len(ID) == 0 {
		qs.db.AddError(errors.New("must at least pass one ID in IDIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("id IN (?)", ID))
}

// IDLt is an autogenerated method
// nolint: dupl
func (qs FunctionControlQuerySet) IDLt(ID uint) FunctionControlQuerySet {
	return qs.w(qs.db.Where("id < ?", ID))
}

// IDLte is an autogenerated method
// nolint: dupl
func (qs FunctionControlQuerySet) IDLte(ID uint) FunctionControlQuerySet {
	return qs.w(qs.db.Where("id <= ?", ID))
}

// IDNe is an autogenerated method
// nolint: dupl
func (qs FunctionControlQuerySet) IDNe(ID uint) FunctionControlQuerySet {
	return qs.w(qs.db.Where("id != ?", ID))
}

// IDNotIn is an autogenerated method
// nolint: dupl
func (qs FunctionControlQuerySet) IDNotIn(ID ...uint) FunctionControlQuerySet {
	if len(ID) == 0 {
		qs.db.AddError(errors.New("must at least pass one ID in IDNotIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("id NOT IN (?)", ID))
}

// Limit is an autogenerated method
// nolint: dupl
func (qs FunctionControlQuerySet) Limit(limit int) FunctionControlQuerySet {
	return qs.w(qs.db.Limit(limit))
}

// Offset is an autogenerated method
// nolint: dupl
func (qs FunctionControlQuerySet) Offset(offset int) FunctionControlQuerySet {
	return qs.w(qs.db.Offset(offset))
}

// One is used to retrieve one result. It returns gorm.ErrRecordNotFound
// if nothing was fetched
func (qs FunctionControlQuerySet) One(ret *FunctionControl) error {
	return qs.db.First(ret).Error
}

// OrderAscByCreatedAt is an autogenerated method
// nolint: dupl
func (qs FunctionControlQuerySet) OrderAscByCreatedAt() FunctionControlQuerySet {
	return qs.w(qs.db.Order("created_at ASC"))
}

// OrderAscByDeletedAt is an autogenerated method
// nolint: dupl
func (qs FunctionControlQuerySet) OrderAscByDeletedAt() FunctionControlQuerySet {
	return qs.w(qs.db.Order("deleted_at ASC"))
}

// OrderAscByID is an autogenerated method
// nolint: dupl
func (qs FunctionControlQuerySet) OrderAscByID() FunctionControlQuerySet {
	return qs.w(qs.db.Order("id ASC"))
}

// OrderAscByUpdatedAt is an autogenerated method
// nolint: dupl
func (qs FunctionControlQuerySet) OrderAscByUpdatedAt() FunctionControlQuerySet {
	return qs.w(qs.db.Order("updated_at ASC"))
}

// OrderDescByCreatedAt is an autogenerated method
// nolint: dupl
func (qs FunctionControlQuerySet) OrderDescByCreatedAt() FunctionControlQuerySet {
	return qs.w(qs.db.Order("created_at DESC"))
}

// OrderDescByDeletedAt is an autogenerated method
// nolint: dupl
func (qs FunctionControlQuerySet) OrderDescByDeletedAt() FunctionControlQuerySet {
	return qs.w(qs.db.Order("deleted_at DESC"))
}

// OrderDescByID is an autogenerated method
// nolint: dupl
func (qs FunctionControlQuerySet) OrderDescByID() FunctionControlQuerySet {
	return qs.w(qs.db.Order("id DESC"))
}

// OrderDescByUpdatedAt is an autogenerated method
// nolint: dupl
func (qs FunctionControlQuerySet) OrderDescByUpdatedAt() FunctionControlQuerySet {
	return qs.w(qs.db.Order("updated_at DESC"))
}

// UpdatedAtEq is an autogenerated method
// nolint: dupl
func (qs FunctionControlQuerySet) UpdatedAtEq(updatedAt time.Time) FunctionControlQuerySet {
	return qs.w(qs.db.Where("updated_at = ?", updatedAt))
}

// UpdatedAtGt is an autogenerated method
// nolint: dupl
func (qs FunctionControlQuerySet) UpdatedAtGt(updatedAt time.Time) FunctionControlQuerySet {
	return qs.w(qs.db.Where("updated_at > ?", updatedAt))
}

// UpdatedAtGte is an autogenerated method
// nolint: dupl
func (qs FunctionControlQuerySet) UpdatedAtGte(updatedAt time.Time) FunctionControlQuerySet {
	return qs.w(qs.db.Where("updated_at >= ?", updatedAt))
}

// UpdatedAtLt is an autogenerated method
// nolint: dupl
func (qs FunctionControlQuerySet) UpdatedAtLt(updatedAt time.Time) FunctionControlQuerySet {
	return qs.w(qs.db.Where("updated_at < ?", updatedAt))
}

// UpdatedAtLte is an autogenerated method
// nolint: dupl
func (qs FunctionControlQuerySet) UpdatedAtLte(updatedAt time.Time) FunctionControlQuerySet {
	return qs.w(qs.db.Where("updated_at <= ?", updatedAt))
}

// UpdatedAtNe is an autogenerated method
// nolint: dupl
func (qs FunctionControlQuerySet) UpdatedAtNe(updatedAt time.Time) FunctionControlQuerySet {
	return qs.w(qs.db.Where("updated_at != ?", updatedAt))
}

// UpdatorEq is an autogenerated method
// nolint: dupl
func (qs FunctionControlQuerySet) UpdatorEq(updator string) FunctionControlQuerySet {
	return qs.w(qs.db.Where("updator = ?", updator))
}

// UpdatorIn is an autogenerated method
// nolint: dupl
func (qs FunctionControlQuerySet) UpdatorIn(updator ...string) FunctionControlQuerySet {
	if len(updator) == 0 {
		qs.db.AddError(errors.New("must at least pass one updator in UpdatorIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("updator IN (?)", updator))
}

// UpdatorNe is an autogenerated method
// nolint: dupl
func (qs FunctionControlQuerySet) UpdatorNe(updator string) FunctionControlQuerySet {
	return qs.w(qs.db.Where("updator != ?", updator))
}

// UpdatorNotIn is an autogenerated method
// nolint: dupl
func (qs FunctionControlQuerySet) UpdatorNotIn(updator ...string) FunctionControlQuerySet {
	if len(updator) == 0 {
		qs.db.AddError(errors.New("must at least pass one updator in UpdatorNotIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("updator NOT IN (?)", updator))
}

// WhiteListEq is an autogenerated method
// nolint: dupl
func (qs FunctionControlQuerySet) WhiteListEq(whiteList string) FunctionControlQuerySet {
	return qs.w(qs.db.Where("white_list = ?", whiteList))
}

// WhiteListIn is an autogenerated method
// nolint: dupl
func (qs FunctionControlQuerySet) WhiteListIn(whiteList ...string) FunctionControlQuerySet {
	if len(whiteList) == 0 {
		qs.db.AddError(errors.New("must at least pass one whiteList in WhiteListIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("white_list IN (?)", whiteList))
}

// WhiteListNe is an autogenerated method
// nolint: dupl
func (qs FunctionControlQuerySet) WhiteListNe(whiteList string) FunctionControlQuerySet {
	return qs.w(qs.db.Where("white_list != ?", whiteList))
}

// WhiteListNotIn is an autogenerated method
// nolint: dupl
func (qs FunctionControlQuerySet) WhiteListNotIn(whiteList ...string) FunctionControlQuerySet {
	if len(whiteList) == 0 {
		qs.db.AddError(errors.New("must at least pass one whiteList in WhiteListNotIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("white_list NOT IN (?)", whiteList))
}

// SetCreatedAt is an autogenerated method
// nolint: dupl
func (u FunctionControlUpdater) SetCreatedAt(createdAt time.Time) FunctionControlUpdater {
	u.fields[string(FunctionControlDBSchema.CreatedAt)] = createdAt
	return u
}

// SetCreator is an autogenerated method
// nolint: dupl
func (u FunctionControlUpdater) SetCreator(creator string) FunctionControlUpdater {
	u.fields[string(FunctionControlDBSchema.Creator)] = creator
	return u
}

// SetDeletedAt is an autogenerated method
// nolint: dupl
func (u FunctionControlUpdater) SetDeletedAt(deletedAt *time.Time) FunctionControlUpdater {
	u.fields[string(FunctionControlDBSchema.DeletedAt)] = deletedAt
	return u
}

// SetDescription is an autogenerated method
// nolint: dupl
func (u FunctionControlUpdater) SetDescription(description string) FunctionControlUpdater {
	u.fields[string(FunctionControlDBSchema.Description)] = description
	return u
}

// SetExtra is an autogenerated method
// nolint: dupl
func (u FunctionControlUpdater) SetExtra(extra string) FunctionControlUpdater {
	u.fields[string(FunctionControlDBSchema.Extra)] = extra
	return u
}

// SetFuncCode is an autogenerated method
// nolint: dupl
func (u FunctionControlUpdater) SetFuncCode(funcCode string) FunctionControlUpdater {
	u.fields[string(FunctionControlDBSchema.FuncCode)] = funcCode
	return u
}

// SetFuncStatus is an autogenerated method
// nolint: dupl
func (u FunctionControlUpdater) SetFuncStatus(funcStatus bool) FunctionControlUpdater {
	u.fields[string(FunctionControlDBSchema.FuncStatus)] = funcStatus
	return u
}

// SetID is an autogenerated method
// nolint: dupl
func (u FunctionControlUpdater) SetID(ID uint) FunctionControlUpdater {
	u.fields[string(FunctionControlDBSchema.ID)] = ID
	return u
}

// SetUpdatedAt is an autogenerated method
// nolint: dupl
func (u FunctionControlUpdater) SetUpdatedAt(updatedAt time.Time) FunctionControlUpdater {
	u.fields[string(FunctionControlDBSchema.UpdatedAt)] = updatedAt
	return u
}

// SetUpdator is an autogenerated method
// nolint: dupl
func (u FunctionControlUpdater) SetUpdator(updator string) FunctionControlUpdater {
	u.fields[string(FunctionControlDBSchema.Updator)] = updator
	return u
}

// SetWhiteList is an autogenerated method
// nolint: dupl
func (u FunctionControlUpdater) SetWhiteList(whiteList string) FunctionControlUpdater {
	u.fields[string(FunctionControlDBSchema.WhiteList)] = whiteList
	return u
}

// Update is an autogenerated method
// nolint: dupl
func (u FunctionControlUpdater) Update() error {
	return u.db.Updates(u.fields).Error
}

// UpdateNum is an autogenerated method
// nolint: dupl
func (u FunctionControlUpdater) UpdateNum() (int64, error) {
	db := u.db.Updates(u.fields)
	return db.RowsAffected, db.Error
}

// ===== END of query set FunctionControlQuerySet

// ===== BEGIN of FunctionControl modifiers

// FunctionControlDBSchemaField describes database schema field. It requires for method 'Update'
type FunctionControlDBSchemaField string

// String method returns string representation of field.
// nolint: dupl
func (f FunctionControlDBSchemaField) String() string {
	return string(f)
}

// FunctionControlDBSchema stores db field names of FunctionControl
var FunctionControlDBSchema = struct {
	ID          FunctionControlDBSchemaField
	CreatedAt   FunctionControlDBSchemaField
	UpdatedAt   FunctionControlDBSchemaField
	DeletedAt   FunctionControlDBSchemaField
	Extra       FunctionControlDBSchemaField
	FuncCode    FunctionControlDBSchemaField
	Description FunctionControlDBSchemaField
	FuncStatus  FunctionControlDBSchemaField
	WhiteList   FunctionControlDBSchemaField
	Creator     FunctionControlDBSchemaField
	Updator     FunctionControlDBSchemaField
}{

	ID:          FunctionControlDBSchemaField("id"),
	CreatedAt:   FunctionControlDBSchemaField("created_at"),
	UpdatedAt:   FunctionControlDBSchemaField("updated_at"),
	DeletedAt:   FunctionControlDBSchemaField("deleted_at"),
	Extra:       FunctionControlDBSchemaField("extra"),
	FuncCode:    FunctionControlDBSchemaField("func_code"),
	Description: FunctionControlDBSchemaField("description"),
	FuncStatus:  FunctionControlDBSchemaField("func_status"),
	WhiteList:   FunctionControlDBSchemaField("white_list"),
	Creator:     FunctionControlDBSchemaField("creator"),
	Updator:     FunctionControlDBSchemaField("updator"),
}

// Update updates FunctionControl fields by primary key
// nolint: dupl
func (o *FunctionControl) Update(db *gorm.DB, fields ...FunctionControlDBSchemaField) error {
	dbNameToFieldName := map[string]interface{}{
		"id":          o.ID,
		"created_at":  o.CreatedAt,
		"updated_at":  o.UpdatedAt,
		"deleted_at":  o.DeletedAt,
		"extra":       o.Extra,
		"func_code":   o.FuncCode,
		"description": o.Description,
		"func_status": o.FuncStatus,
		"white_list":  o.WhiteList,
		"creator":     o.Creator,
		"updator":     o.Updator,
	}
	u := map[string]interface{}{}
	for _, f := range fields {
		fs := f.String()
		u[fs] = dbNameToFieldName[fs]
	}
	if err := db.Model(o).Updates(u).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return err
		}

		return fmt.Errorf("can't update FunctionControl %v fields %v: %s",
			o, fields, err)
	}

	return nil
}

// FunctionControlUpdater is an FunctionControl updates manager
type FunctionControlUpdater struct {
	fields map[string]interface{}
	db     *gorm.DB
}

// NewFunctionControlUpdater creates new FunctionControl updater
// nolint: dupl
func NewFunctionControlUpdater(db *gorm.DB) FunctionControlUpdater {
	return FunctionControlUpdater{
		fields: map[string]interface{}{},
		db:     db.Model(&FunctionControl{}),
	}
}

// ===== END of FunctionControl modifiers

// ===== END of all query sets