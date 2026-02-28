package repository

import (
	"edu/model"

	"gorm.io/gorm"
)

// IClassRepository 班级数据访问接口
type IClassRepository interface {
	Create(c *model.Class) error
	Update(c *model.Class) error
	Delete(id uint) error
	FindByID(id uint) (*model.Class, error)
	FindByInviteCode(code string) (*model.Class, error)
	FindPage(query *model.ClassQuery, offset, limit int) ([]*model.Class, int64, error)
	FindAll(query *model.ClassQuery) ([]*model.Class, error)
	AddStudent(classId, userId uint) error
	RemoveStudent(classId, userId uint) error
	FindStudents(classId uint) ([]*model.User, error)
	FindStudentsWithStatus(classId uint) ([]*model.ClassStudentView, error)
	UpdateStudentStatus(classId, userId uint, status int) error
	// IsStudentInOtherAdministrativeClass checks whether a user is already in an administrative class
	// other than the one identified by excludeClassId.
	IsStudentInOtherAdministrativeClass(userId, excludeClassId uint) (bool, error)
	// Teacher management
	AddTeacher(classId, userId uint) error
	RemoveTeacher(classId, userId uint) error
	FindTeachers(classId uint) ([]*model.User, error)
	IsTeacherInClass(classId, userId uint) (bool, error)
}

type classRepository struct {
	db *gorm.DB
}

func NewClassRepository(db *gorm.DB) IClassRepository {
	return &classRepository{db: db}
}

func (r *classRepository) Create(c *model.Class) error {
	return r.db.Create(c).Error
}

func (r *classRepository) Update(c *model.Class) error {
	return r.db.Model(c).Updates(c).Error
}

func (r *classRepository) Delete(id uint) error {
	return r.db.Delete(&model.Class{}, id).Error
}

func (r *classRepository) FindByID(id uint) (*model.Class, error) {
	var c model.Class
	err := r.db.Where("id = ?", id).
		Preload("AdminUser").
		First(&c).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &c, err
}

func (r *classRepository) FindByInviteCode(code string) (*model.Class, error) {
	var c model.Class
	err := r.db.Where("invite_code = ?", code).
		Preload("AdminUser").
		First(&c).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &c, err
}

func (r *classRepository) FindPage(query *model.ClassQuery, offset, limit int) ([]*model.Class, int64, error) {
	var classes []*model.Class
	var total int64

	q := r.db.Model(&model.Class{}).Preload("AdminUser")

	if query.ID != 0 {
		q = q.Where("id = ?", query.ID)
	}
	if query.AdminUserId != 0 {
		q = q.Where("admin_user_id = ?", query.AdminUserId)
	}
	if query.InviteCode != "" {
		q = q.Where("invite_code = ?", query.InviteCode)
	}
	if query.ClassType != 0 {
		q = q.Where("class_type = ?", query.ClassType)
	}

	q.Count(&total)
	err := q.Order("id DESC").Offset(offset).Limit(limit).Find(&classes).Error
	return classes, total, err
}

func (r *classRepository) FindAll(query *model.ClassQuery) ([]*model.Class, error) {
	var classes []*model.Class

	q := r.db.Model(&model.Class{}).Preload("AdminUser")

	if query.ID != 0 {
		q = q.Where("id = ?", query.ID)
	}
	if query.AdminUserId != 0 {
		q = q.Where("admin_user_id = ?", query.AdminUserId)
	}
	if query.InviteCode != "" {
		q = q.Where("invite_code = ?", query.InviteCode)
	}
	if query.ClassType != 0 {
		q = q.Where("class_type = ?", query.ClassType)
	}

	err := q.Order("id DESC").Find(&classes).Error
	return classes, err
}

func (r *classRepository) AddStudent(classId, userId uint) error {
	class := model.Class{Model: model.Model{ID: classId}}
	user := model.User{Model: model.Model{ID: userId}}
	return r.db.Model(&class).Association("Students").Append(&user)
}

func (r *classRepository) RemoveStudent(classId, userId uint) error {
	class := model.Class{Model: model.Model{ID: classId}}
	user := model.User{Model: model.Model{ID: userId}}
	return r.db.Model(&class).Association("Students").Delete(&user)
}

func (r *classRepository) FindStudents(classId uint) ([]*model.User, error) {
	var class model.Class
	err := r.db.Where("id = ?", classId).Preload("Students").First(&class).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return class.Students, err
}

func (r *classRepository) FindStudentsWithStatus(classId uint) ([]*model.ClassStudentView, error) {
	var results []*model.ClassStudentView
	err := r.db.Table("users").
		Select("users.*, COALESCE(user_class_relation.status, ?) AS student_status", model.ClassStudentStatusStudying).
		Joins("JOIN user_class_relation ON user_class_relation.user_id = users.id AND user_class_relation.class_id = ?", classId).
		Where("users.deleted_at IS NULL").
		Scan(&results).Error
	return results, err
}

func (r *classRepository) UpdateStudentStatus(classId, userId uint, status int) error {
	result := r.db.Model(&model.UserClassRelation{}).
		Where("class_id = ? AND user_id = ?", classId, userId).
		Update("status", status)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (r *classRepository) IsStudentInOtherAdministrativeClass(userId, excludeClassId uint) (bool, error) {
	var count int64
	q := r.db.Model(&model.Class{}).
		Joins("JOIN user_class_relation ON user_class_relation.class_id = classes.id").
		Where("user_class_relation.user_id = ? AND classes.class_type = ?", userId, model.ClassTypeAdministrative)
	if excludeClassId != 0 {
		q = q.Where("classes.id != ?", excludeClassId)
	}
	err := q.Count(&count).Error
	return count > 0, err
}

func (r *classRepository) AddTeacher(classId, userId uint) error {
	class := model.Class{Model: model.Model{ID: classId}}
	user := model.User{Model: model.Model{ID: userId}}
	return r.db.Model(&class).Association("Teachers").Append(&user)
}

func (r *classRepository) RemoveTeacher(classId, userId uint) error {
	class := model.Class{Model: model.Model{ID: classId}}
	user := model.User{Model: model.Model{ID: userId}}
	return r.db.Model(&class).Association("Teachers").Delete(&user)
}

func (r *classRepository) FindTeachers(classId uint) ([]*model.User, error) {
	var class model.Class
	err := r.db.Where("id = ?", classId).Preload("Teachers").First(&class).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return class.Teachers, err
}

func (r *classRepository) IsTeacherInClass(classId, userId uint) (bool, error) {
	var count int64
	err := r.db.Model(&model.Class{}).
		Joins("JOIN class_teacher_relation ON class_teacher_relation.class_id = classes.id").
		Where("classes.id = ? AND class_teacher_relation.user_id = ?", classId, userId).
		Count(&count).Error
	return count > 0, err
}

// IClassJoinRequestRepository 入班申请数据访问接口
type IClassJoinRequestRepository interface {
	Create(r *model.ClassJoinRequest) error
	Update(r *model.ClassJoinRequest) error
	FindByID(id uint) (*model.ClassJoinRequest, error)
	FindByClassAndUser(classId, userId uint) (*model.ClassJoinRequest, error)
	FindPage(query *model.ClassJoinRequestQuery, offset, limit int) ([]*model.ClassJoinRequest, int64, error)
}

type classJoinRequestRepository struct {
	db *gorm.DB
}

func NewClassJoinRequestRepository(db *gorm.DB) IClassJoinRequestRepository {
	return &classJoinRequestRepository{db: db}
}

func (r *classJoinRequestRepository) Create(req *model.ClassJoinRequest) error {
	return r.db.Create(req).Error
}

func (r *classJoinRequestRepository) Update(req *model.ClassJoinRequest) error {
	return r.db.Model(req).Updates(req).Error
}

func (r *classJoinRequestRepository) FindByID(id uint) (*model.ClassJoinRequest, error) {
	var req model.ClassJoinRequest
	err := r.db.Where("id = ?", id).
		Preload("Class").
		Preload("User").
		First(&req).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &req, err
}

func (r *classJoinRequestRepository) FindByClassAndUser(classId, userId uint) (*model.ClassJoinRequest, error) {
	var req model.ClassJoinRequest
	err := r.db.Where("class_id = ? AND user_id = ?", classId, userId).
		Order("created_at DESC").
		First(&req).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &req, err
}

func (r *classJoinRequestRepository) FindPage(query *model.ClassJoinRequestQuery, offset, limit int) ([]*model.ClassJoinRequest, int64, error) {
	var requests []*model.ClassJoinRequest
	var total int64

	q := r.db.Model(&model.ClassJoinRequest{}).
		Preload("Class").
		Preload("User")

	if query.ClassId != 0 {
		q = q.Where("class_id = ?", query.ClassId)
	}
	if query.UserId != 0 {
		q = q.Where("user_id = ?", query.UserId)
	}
	if query.Status != nil {
		q = q.Where("status = ?", *query.Status)
	}

	q.Count(&total)
	err := q.Order("id DESC").Offset(offset).Limit(limit).Find(&requests).Error
	return requests, total, err
}

// IClassTeacherApplicationRepository 教师加入班级申请数据访问接口
type IClassTeacherApplicationRepository interface {
	Create(r *model.ClassTeacherApplication) error
	Update(r *model.ClassTeacherApplication) error
	FindByID(id uint) (*model.ClassTeacherApplication, error)
	FindByClassAndUser(classId, userId uint) (*model.ClassTeacherApplication, error)
	FindPage(query *model.ClassTeacherApplicationQuery, offset, limit int) ([]*model.ClassTeacherApplication, int64, error)
}

type classTeacherApplicationRepository struct {
	db *gorm.DB
}

func NewClassTeacherApplicationRepository(db *gorm.DB) IClassTeacherApplicationRepository {
	return &classTeacherApplicationRepository{db: db}
}

func (r *classTeacherApplicationRepository) Create(req *model.ClassTeacherApplication) error {
	return r.db.Create(req).Error
}

func (r *classTeacherApplicationRepository) Update(req *model.ClassTeacherApplication) error {
	return r.db.Model(req).Updates(req).Error
}

func (r *classTeacherApplicationRepository) FindByID(id uint) (*model.ClassTeacherApplication, error) {
	var req model.ClassTeacherApplication
	err := r.db.Where("id = ?", id).
		Preload("Class").
		Preload("User").
		First(&req).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &req, err
}

func (r *classTeacherApplicationRepository) FindByClassAndUser(classId, userId uint) (*model.ClassTeacherApplication, error) {
	var req model.ClassTeacherApplication
	err := r.db.Where("class_id = ? AND user_id = ?", classId, userId).
		Order("created_at DESC").
		First(&req).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &req, err
}

func (r *classTeacherApplicationRepository) FindPage(query *model.ClassTeacherApplicationQuery, offset, limit int) ([]*model.ClassTeacherApplication, int64, error) {
	var requests []*model.ClassTeacherApplication
	var total int64

	q := r.db.Model(&model.ClassTeacherApplication{}).
		Preload("Class").
		Preload("User")

	if query.ClassId != 0 {
		q = q.Where("class_id = ?", query.ClassId)
	}
	if query.UserId != 0 {
		q = q.Where("user_id = ?", query.UserId)
	}
	if query.Status != nil {
		q = q.Where("status = ?", *query.Status)
	}

	q.Count(&total)
	err := q.Order("id DESC").Offset(offset).Limit(limit).Find(&requests).Error
	return requests, total, err
}
