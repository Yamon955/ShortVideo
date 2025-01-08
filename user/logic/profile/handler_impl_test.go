package profile

import (
	"context"
	"errors"
	"testing"

	"github.com/Yamon955/ShortVideo/base"
	"github.com/Yamon955/ShortVideo/protocol/user/pb"
	"github.com/Yamon955/ShortVideo/user/entity/errcode"
	MySQL "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"trpc.group/trpc-go/trpc-go/errs"
)

// MockDBClient 是一个模拟的数据库客户端
type MockDBClient struct {
	mock.Mock
}

func (m *MockDBClient) FindUserByUsername(ctx context.Context, username string) (*base.User, error) {
	args := m.Called(ctx, username)
	user := args.Get(0)
	if user == nil {
		return nil, args.Error(1)
	}
	return user.(*base.User), args.Error(1)
}

func (m *MockDBClient) FindUserByUID(ctx context.Context, uid uint64) (*base.User, error) {
	args := m.Called(ctx, uid)
	user := args.Get(0)
	if user == nil {
		return nil, args.Error(1)
	}
	return user.(*base.User), args.Error(1)
}

func (m *MockDBClient) CreateUser(ctx context.Context, user *base.User) error {
	args := m.Called(ctx, user)
	return args.Error(0)
}

func (m *MockDBClient) UpdateUserInfo(ctx context.Context, uid uint64, updateInfo map[string]interface{}) error {
	args := m.Called(ctx, uid, updateInfo)
	return args.Error(0)
}

func (m *MockDBClient) GetPublishListCount(ctx context.Context, uid uint64) (int64, error) {
	args := m.Called(ctx, uid)
	return args.Get(0).(int64), args.Error(1)
}

func (m *MockDBClient) GetLikedListCount(ctx context.Context, uid uint64) (int64, error) {
	args := m.Called(ctx, uid)
	return args.Get(0).(int64), args.Error(1)
}

func (m *MockDBClient) GetCollectListCount(ctx context.Context, uid uint64) (int64, error) {
	args := m.Called(ctx, uid)
	return args.Get(0).(int64), args.Error(1)
}

func TestHandleBatchGetProfile(t *testing.T) {
	ctx := context.Background()
	mockDB := new(MockDBClient)
	handler := &handlerImpl{db: mockDB}

	req := &pb.BatchGetProfileReq{
		Uids:         []uint64{1, 2},
		ProfileTypes: []pb.PROFILE_TYPES{pb.PROFILE_TYPES_MAIN_PAGE_INFO},
	}
	rsp := &pb.BatchGetProfileRsp{}

	user := &base.User{
		ID:       1,
		Username: "testuser",
		Avatar:   "http://example.com/avatar.jpg",
		Sign:     "Hello",
		Gender:   1,
	}

	mockDB.On("FindUserByUID", ctx, uint64(1)).Return(user, nil)
	mockDB.On("FindUserByUID", ctx, uint64(2)).Return(nil, errors.New("user not found"))

	err := handler.HandleBatchGetProfile(ctx, req, rsp)
	assert.NoError(t, err)
	assert.Len(t, rsp.UserInfos, 1)
	assert.Equal(t, uint64(1), rsp.UserInfos[0].MainPageInfo.ID)
	assert.Len(t, rsp.FailedUids, 1)
}

func TestHandleSetProfile(t *testing.T) {
	ctx := context.Background()
	mockDB := new(MockDBClient)
	handler := &handlerImpl{db: mockDB}

	req := &pb.SetProfileReq{
		Uid:          1,
		ProfileTypes: []pb.PROFILE_TYPES{pb.PROFILE_TYPES_USERNAME, pb.PROFILE_TYPES_AVATOR},
		Username:     "newuser",
		Avator:       "http://example.com/newavatar.jpg",
	}
	rsp := &pb.SetProfileRsp{}

	user := &base.User{
		ID:       1,
		Username: "testuser",
		Avatar:   "http://example.com/avatar.jpg",
		Sign:     "Hello",
		Gender:   1,
	}
	mockDB.On("FindUserByUID", ctx, req.GetUid()).Return(user, nil)
	mockDB.On("UpdateUserInfo", ctx, uint64(1), mock.Anything).Return(nil)

	err := handler.HandleSetProfile(ctx, req, rsp)
	assert.NoError(t, err)
	assert.Equal(t, 0, int(errs.Code(err)))
}

func TestHandleSetProfile_UsernameTooLong(t *testing.T) {
	ctx := context.Background()
	mockDB := new(MockDBClient)
	handler := &handlerImpl{db: mockDB}

	req := &pb.SetProfileReq{
		Uid:          1,
		ProfileTypes: []pb.PROFILE_TYPES{pb.PROFILE_TYPES_USERNAME},
		Username:     "thisusernameiswaytoolongandshouldfailvalidation",
	}
	rsp := &pb.SetProfileRsp{}

	user := &base.User{
		ID:       1,
		Username: "testuser",
		Avatar:   "http://example.com/avatar.jpg",
		Sign:     "Hello",
		Gender:   1,
	}
	mockDB.On("FindUserByUID", ctx, req.GetUid()).Return(user, nil)
	// 设置预期的 UpdateUserInfo 方法调用和返回值
	mockDB.On("UpdateUserInfo", mock.Anything, uint64(1), mock.Anything).Return(nil)

	err := handler.HandleSetProfile(ctx, req, rsp)
	assert.Error(t, err)
	assert.Equal(t, errcode.ErrUsername, int(errs.Code(err)))
}

func TestHandleSetProfile_UsernameAlreadyUsed(t *testing.T) {
	ctx := context.Background()
	mockDB := new(MockDBClient)
	handler := &handlerImpl{db: mockDB}

	req := &pb.SetProfileReq{
		Uid:          1,
		ProfileTypes: []pb.PROFILE_TYPES{pb.PROFILE_TYPES_USERNAME},
		Username:     "existinguser",
	}
	rsp := &pb.SetProfileRsp{}

	user := &base.User{
		ID:       1,
		Username: "testuser",
		Avatar:   "http://example.com/avatar.jpg",
		Sign:     "Hello",
		Gender:   1,
	}
	mockDB.On("FindUserByUID", ctx, req.GetUid()).Return(user, nil)
	// 设置预期的 UpdateUserInfo 方法调用和返回值
	mockDB.On("UpdateUserInfo", mock.Anything, uint64(1), mock.Anything).Return(
		&MySQL.MySQLError{
			Number:  1062,
			Message: "Duplicate entry 'duplicate' for key 'users.uni_users_username'"},
	)

	err := handler.HandleSetProfile(ctx, req, rsp)
	assert.Error(t, err)
	assert.Equal(t, errcode.ErrUsernameIsUsed, int(errs.Code(err)))
}

func TestHandleSetProfile_UpdateUserInfoError(t *testing.T) {
	ctx := context.Background()
	mockDB := new(MockDBClient)
	handler := &handlerImpl{db: mockDB}

	req := &pb.SetProfileReq{
		Uid:          1,
		ProfileTypes: []pb.PROFILE_TYPES{pb.PROFILE_TYPES_USERNAME, pb.PROFILE_TYPES_AVATOR},
		Username:     "newuser",
		Avator:       "http://example.com/newavatar.jpg",
	}
	rsp := &pb.SetProfileRsp{}

	user := &base.User{
		ID:       1,
		Username: "testuser",
		Avatar:   "http://example.com/avatar.jpg",
		Sign:     "Hello",
		Gender:   1,
	}
	mockDB.On("FindUserByUID", ctx, req.GetUid()).Return(user, nil)
	mockDB.On("UpdateUserInfo", ctx, uint64(1), mock.Anything).Return(errors.New("update failed"))

	err := handler.HandleSetProfile(ctx, req, rsp)
	assert.Error(t, err)
	assert.Equal(t, errcode.ErrUpdateUserInfo, int(errs.Code(err)))
}
