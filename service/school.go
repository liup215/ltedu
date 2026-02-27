package service

import (
"crypto/rand"
"encoding/hex"
"edu/repository"
"edu/model"
"errors"
)

var SchoolSvr = &SchoolService{baseService: newBaseService()}

type SchoolService struct {
baseService
}

// 年级管理
func (svr *SchoolService) SelectGradeList(q model.GradeQuery) ([]*model.Grade, int64, error) {
page := q.Page.CheckPage()
list, total, err := repository.GradeRepo.FindPage(&q, (page.PageIndex-1)*page.PageSize, page.PageSize)
return list, total, err
}

func (svr *SchoolService) SelectGradeById(id uint) (*model.Grade, error) {
if id == 0 {
return nil, errors.New("无效的ID")
}
return repository.GradeRepo.FindByID(id)
}

func (svr *SchoolService) SelectGradeAll(q model.GradeQuery) ([]*model.Grade, error) {
return repository.GradeRepo.FindAll(&q)
}

func (svr *SchoolService) CreateGrade(or model.GradeCreateEditRequest) (*model.Grade, error) {
if or.ID != 0 {
or.ID = 0
}
o := model.Grade{
Name:               or.Name,
GradeLeadTeacherId: or.GradeLeadTeacherId,
}
err := repository.GradeRepo.Create(&o)
return &o, err
}

func (svr *SchoolService) EditGrade(or model.GradeCreateEditRequest) error {
if or.ID == 0 {
return errors.New("无效的ID")
}
g := model.Grade{
Model:              model.Model{ID: or.ID},
Name:               or.Name,
GradeLeadTeacherId: or.GradeLeadTeacherId,
}
err := repository.GradeRepo.Update(&g)
return err
}

func (svr *SchoolService) DeleteGrade(id uint) error {
if id == 0 {
return errors.New("无效的ID")
}
return repository.GradeRepo.Delete(id)
}

/* 班级类型相关方法补充，保证 controller 编译通过 */
func (svr *SchoolService) SelectClassTypeList(q interface{}) ([]interface{}, int64, error) {
	return nil, 0, nil
}
func (svr *SchoolService) SelectClassTypeById(id uint) (interface{}, error) {
	return nil, nil
}
func (svr *SchoolService) SelectClassTypeAll(q interface{}) ([]interface{}, error) {
	return nil, nil
}
func (svr *SchoolService) CreateClassType(o interface{}) error {
	return nil
}
func (svr *SchoolService) EditClassType(o interface{}) error {
	return nil
}
func (svr *SchoolService) DeleteClassType(id uint) error {
return nil
}

// generateInviteCode 生成唯一邀请码
func generateInviteCode() (string, error) {
	b := make([]byte, 8)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}

// 班级管理

func (svr *SchoolService) SelectClassList(q model.ClassQuery) ([]*model.Class, int64, error) {
	page := q.Page.CheckPage()
	return repository.ClassRepo.FindPage(&q, (page.PageIndex-1)*page.PageSize, page.PageSize)
}

func (svr *SchoolService) SelectClassById(id uint) (*model.Class, error) {
	if id == 0 {
		return nil, errors.New("无效的ID")
	}
	return repository.ClassRepo.FindByID(id)
}

func (svr *SchoolService) SelectClassAll(q model.ClassQuery) ([]*model.Class, error) {
	return repository.ClassRepo.FindAll(&q)
}

// CreateClass 创建班级（仅教师可创建，创建者自动成为管理员）
func (svr *SchoolService) CreateClass(req model.ClassCreateEditRequest, creatorId uint) (*model.Class, error) {
	if req.Name == "" {
		return nil, errors.New("班级名称不能为空")
	}
	user, err := repository.UserRepo.FindByID(creatorId)
	if err != nil || user == nil {
		return nil, errors.New("用户不存在")
	}
	if !user.IsTeacher {
		return nil, errors.New("只有教师身份可以创建班级")
	}
	classType := req.ClassType
	if classType != model.ClassTypeTeaching && classType != model.ClassTypeAdministrative {
		classType = model.ClassTypeTeaching
	}
	code, err := generateInviteCode()
	if err != nil {
		return nil, errors.New("邀请码生成失败")
	}
	class := &model.Class{
		Name:        req.Name,
		ClassType:   classType,
		InviteCode:  code,
		AdminUserId: creatorId,
	}
	if err := repository.ClassRepo.Create(class); err != nil {
		return nil, err
	}
	return class, nil
}

func (svr *SchoolService) EditClass(req model.ClassCreateEditRequest) error {
	if req.ID == 0 {
		return errors.New("无效的ID")
	}
	class := &model.Class{
		Model: model.Model{ID: req.ID},
		Name:  req.Name,
	}
	return repository.ClassRepo.Update(class)
}

func (svr *SchoolService) DeleteClass(id uint) error {
	if id == 0 {
		return errors.New("无效的ID")
	}
	return repository.ClassRepo.Delete(id)
}

func (svr *SchoolService) GetStudentsByClassId(classId uint) ([]*model.User, error) {
	return repository.ClassRepo.FindStudents(classId)
}

func (svr *SchoolService) DeleteStudentFromClass(classId, userId uint) error {
	return repository.ClassRepo.RemoveStudent(classId, userId)
}

// AddStudentDirectly 超级管理员直接添加学生到班级（绕过邀请码流程）
func (svr *SchoolService) AddStudentDirectly(classId, userId, adminId uint) error {
	admin, err := repository.UserRepo.FindByID(adminId)
	if err != nil || admin == nil || !admin.IsAdmin {
		return errors.New("只有超级管理员可以直接添加学生")
	}
	if classId == 0 || userId == 0 {
		return errors.New("班级ID和用户ID不能为空")
	}
	class, err := repository.ClassRepo.FindByID(classId)
	if err != nil || class == nil {
		return errors.New("班级不存在")
	}
	user, err := repository.UserRepo.FindByID(userId)
	if err != nil || user == nil {
		return errors.New("用户不存在")
	}
	if err := svr.checkAdministrativeClassConstraint(class, userId); err != nil {
		return err
	}
	return repository.ClassRepo.AddStudent(classId, userId)
}

// ApplyToJoinClass 学生使用邀请码申请加入班级
func (svr *SchoolService) ApplyToJoinClass(inviteCode string, userId uint, message string) (*model.ClassJoinRequest, error) {
	class, err := repository.ClassRepo.FindByInviteCode(inviteCode)
	if err != nil {
		return nil, err
	}
	if class == nil {
		return nil, errors.New("邀请码无效")
	}
	// Check if already a member
	students, _ := repository.ClassRepo.FindStudents(class.ID)
	for _, s := range students {
		if s.ID == userId {
			return nil, errors.New("您已经是该班级成员")
		}
	}
	// Check for existing pending request
	existing, err := repository.ClassJoinRequestRepo.FindByClassAndUser(class.ID, userId)
	if err != nil {
		return nil, err
	}
	if existing != nil && existing.Status == model.ClassJoinStatusPending {
		return nil, errors.New("您已经提交过申请，请等待审核")
	}
	req := &model.ClassJoinRequest{
		ClassId: class.ID,
		UserId:  userId,
		Status:  model.ClassJoinStatusPending,
		Message: message,
	}
	if err := repository.ClassJoinRequestRepo.Create(req); err != nil {
		return nil, err
	}
	return req, nil
}

// ListClassJoinRequests 列出班级的加入申请（管理员）
func (svr *SchoolService) ListClassJoinRequests(q model.ClassJoinRequestQuery) ([]*model.ClassJoinRequest, int64, error) {
	page := q.Page.CheckPage()
	return repository.ClassJoinRequestRepo.FindPage(&q, (page.PageIndex-1)*page.PageSize, page.PageSize)
}

// checkAdministrativeClassConstraint ensures a user does not join a second administrative class
func (svr *SchoolService) checkAdministrativeClassConstraint(class *model.Class, userId uint) error {
	if class.ClassType != model.ClassTypeAdministrative {
		return nil
	}
	already, err := repository.ClassRepo.IsStudentInOtherAdministrativeClass(userId, class.ID)
	if err != nil {
		return err
	}
	if already {
		return errors.New("每个用户只能属于一个行政班")
	}
	return nil
}

// checkClassAdminPermission verifies that the user is either the class admin or a system admin
func (svr *SchoolService) checkClassAdminPermission(class *model.Class, userId uint) error {
	if class.AdminUserId == userId {
		return nil
	}
	user, err := repository.UserRepo.FindByID(userId)
	if err != nil || user == nil || !user.IsAdmin {
		return errors.New("只有班级管理员或系统管理员可以审核申请")
	}
	return nil
}

// ApproveJoinRequest 审核通过加入申请
func (svr *SchoolService) ApproveJoinRequest(requestId, adminUserId uint) error {
	req, err := repository.ClassJoinRequestRepo.FindByID(requestId)
	if err != nil {
		return err
	}
	if req == nil {
		return errors.New("申请不存在")
	}
	if req.Status != model.ClassJoinStatusPending {
		return errors.New("该申请已处理")
	}
	class, err := repository.ClassRepo.FindByID(req.ClassId)
	if err != nil || class == nil {
		return errors.New("班级不存在")
	}
	if err := svr.checkClassAdminPermission(class, adminUserId); err != nil {
		return err
	}
	if err := svr.checkAdministrativeClassConstraint(class, req.UserId); err != nil {
		return err
	}
	req.Status = model.ClassJoinStatusApproved
	if err := repository.ClassJoinRequestRepo.Update(req); err != nil {
		return err
	}
	return repository.ClassRepo.AddStudent(req.ClassId, req.UserId)
}

// RejectJoinRequest 拒绝加入申请
func (svr *SchoolService) RejectJoinRequest(requestId, adminUserId uint) error {
	req, err := repository.ClassJoinRequestRepo.FindByID(requestId)
	if err != nil {
		return err
	}
	if req == nil {
		return errors.New("申请不存在")
	}
	if req.Status != model.ClassJoinStatusPending {
		return errors.New("该申请已处理")
	}
	class, err := repository.ClassRepo.FindByID(req.ClassId)
	if err != nil || class == nil {
		return errors.New("班级不存在")
	}
	if err := svr.checkClassAdminPermission(class, adminUserId); err != nil {
		return err
	}
	req.Status = model.ClassJoinStatusRejected
	return repository.ClassJoinRequestRepo.Update(req)
}

