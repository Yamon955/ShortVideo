package profile

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/Yamon955/ShortVideo/base"
	"github.com/Yamon955/ShortVideo/protocol/user/pb"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
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
	fmt.Printf("rsp%v", rsp)
	assert.NoError(t, err)
	assert.Len(t, rsp.UserInfos, 1)
	assert.Equal(t, uint64(1), rsp.UserInfos[0].MainPageInfo.ID)
	assert.Len(t, rsp.FailedUIDs, 1)
	assert.Equal(t, uint64(2), rsp.FailedUIDs[0])
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

	mockDB.On("FindUserByUsername", ctx, "newuser").Return(nil, gorm.ErrRecordNotFound)
	mockDB.On("UpdateUserInfo", ctx, uint64(1), mock.Anything).Return(nil)

	err := handler.HandleSetProfile(ctx, req, rsp)
	assert.NoError(t, err)
	assert.Empty(t, rsp.FailedTypes)
}

func TestHandleSetProfile_UsernameTooLong(t *testing.T) {
	ctx := context.Background()
	mockDB := new(MockDBClient)
	handler := &handlerImpl{db: mockDB}

	// 设置预期的 UpdateUserInfo 方法调用和返回值
	mockDB.On("UpdateUserInfo", mock.Anything, uint64(1), mock.Anything).Return(nil)

	req := &pb.SetProfileReq{
		Uid:          1,
		ProfileTypes: []pb.PROFILE_TYPES{pb.PROFILE_TYPES_USERNAME},
		Username:     "thisusernameiswaytoolongandshouldfailvalidation",
	}
	rsp := &pb.SetProfileRsp{}

	err := handler.HandleSetProfile(ctx, req, rsp)
	assert.NoError(t, err)
	assert.NotEmpty(t, rsp.FailedTypes)
	assert.Equal(t, "username too long", rsp.FailedTypes[int32(pb.PROFILE_TYPES_USERNAME)])
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

	existingUser := &base.User{
		ID:       2,
		Username: "existinguser",
	}

	mockDB.On("FindUserByUsername", ctx, "existinguser").Return(existingUser, nil)
	// 设置预期的 UpdateUserInfo 方法调用和返回值
	mockDB.On("UpdateUserInfo", mock.Anything, uint64(1), mock.Anything).Return(nil)

	err := handler.HandleSetProfile(ctx, req, rsp)
	assert.NoError(t, err)
	assert.NotEmpty(t, rsp.FailedTypes)
	assert.Equal(t, "username is used", rsp.FailedTypes[int32(pb.PROFILE_TYPES_USERNAME)])
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

	mockDB.On("FindUserByUsername", ctx, "newuser").Return(nil, gorm.ErrRecordNotFound)
	mockDB.On("UpdateUserInfo", ctx, uint64(1), mock.Anything).Return(errors.New("update failed"))

	err := handler.HandleSetProfile(ctx, req, rsp)
	assert.NoError(t, err)
	assert.NotEmpty(t, rsp.FailedTypes)
	assert.Equal(t, "update failed", rsp.FailedTypes[int32(pb.PROFILE_TYPES_AVATOR)])
}

func TestCheckUsername(t *testing.T) {
	ctx := context.Background()
	mockDB := new(MockDBClient)

	// Test case: Username too long
	failedTypes := make(map[int32]string)
	result := checkUsername(ctx, "thisusernameiswaytoolongandshouldfailvalidation", failedTypes, mockDB)
	assert.False(t, result)
	assert.Equal(t, "username too long", failedTypes[int32(pb.PROFILE_TYPES_USERNAME)])

	// Test case: Username already used
	failedTypes = make(map[int32]string)
	existingUser := &base.User{
		ID:       2,
		Username: "existinguser",
	}
	mockDB.On("FindUserByUsername", ctx, "existinguser").Return(existingUser, nil)
	result = checkUsername(ctx, "existinguser", failedTypes, mockDB)
	assert.False(t, result)
	assert.Equal(t, "username is used", failedTypes[int32(pb.PROFILE_TYPES_USERNAME)])

	// Test case: Username available
	failedTypes = make(map[int32]string)
	mockDB.On("FindUserByUsername", ctx, "newuser").Return(nil, gorm.ErrRecordNotFound)
	result = checkUsername(ctx, "newuser", failedTypes, mockDB)
	assert.True(t, result)
	assert.Empty(t, failedTypes)
}
