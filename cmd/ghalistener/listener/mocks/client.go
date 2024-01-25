// Code generated by mockery v2.36.1. DO NOT EDIT.

package mocks

import (
	context "context"

	actions "github.com/actions/actions-runner-controller/github/actions"

	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// Client is an autogenerated mock type for the Client type
type Client struct {
	mock.Mock
}

// AcquireJobs provides a mock function with given fields: ctx, runnerScaleSetId, messageQueueAccessToken, requestIds
func (_m *Client) AcquireJobs(ctx context.Context, runnerScaleSetId int, messageQueueAccessToken string, requestIds []int64) ([]int64, error) {
	ret := _m.Called(ctx, runnerScaleSetId, messageQueueAccessToken, requestIds)

	var r0 []int64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int, string, []int64) ([]int64, error)); ok {
		return rf(ctx, runnerScaleSetId, messageQueueAccessToken, requestIds)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int, string, []int64) []int64); ok {
		r0 = rf(ctx, runnerScaleSetId, messageQueueAccessToken, requestIds)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]int64)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int, string, []int64) error); ok {
		r1 = rf(ctx, runnerScaleSetId, messageQueueAccessToken, requestIds)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateMessageSession provides a mock function with given fields: ctx, runnerScaleSetId, owner
func (_m *Client) CreateMessageSession(ctx context.Context, runnerScaleSetId int, owner string) (*actions.RunnerScaleSetSession, error) {
	ret := _m.Called(ctx, runnerScaleSetId, owner)

	var r0 *actions.RunnerScaleSetSession
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int, string) (*actions.RunnerScaleSetSession, error)); ok {
		return rf(ctx, runnerScaleSetId, owner)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int, string) *actions.RunnerScaleSetSession); ok {
		r0 = rf(ctx, runnerScaleSetId, owner)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*actions.RunnerScaleSetSession)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int, string) error); ok {
		r1 = rf(ctx, runnerScaleSetId, owner)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteMessage provides a mock function with given fields: ctx, messageQueueUrl, messageQueueAccessToken, messageId
func (_m *Client) DeleteMessage(ctx context.Context, messageQueueUrl string, messageQueueAccessToken string, messageId int64) error {
	ret := _m.Called(ctx, messageQueueUrl, messageQueueAccessToken, messageId)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, int64) error); ok {
		r0 = rf(ctx, messageQueueUrl, messageQueueAccessToken, messageId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteMessageSession provides a mock function with given fields: ctx, runnerScaleSetId, sessionId
func (_m *Client) DeleteMessageSession(ctx context.Context, runnerScaleSetId int, sessionId *uuid.UUID) error {
	ret := _m.Called(ctx, runnerScaleSetId, sessionId)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int, *uuid.UUID) error); ok {
		r0 = rf(ctx, runnerScaleSetId, sessionId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAcquirableJobs provides a mock function with given fields: ctx, runnerScaleSetId
func (_m *Client) GetAcquirableJobs(ctx context.Context, runnerScaleSetId int) (*actions.AcquirableJobList, error) {
	ret := _m.Called(ctx, runnerScaleSetId)

	var r0 *actions.AcquirableJobList
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int) (*actions.AcquirableJobList, error)); ok {
		return rf(ctx, runnerScaleSetId)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int) *actions.AcquirableJobList); ok {
		r0 = rf(ctx, runnerScaleSetId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*actions.AcquirableJobList)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, runnerScaleSetId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetMessage provides a mock function with given fields: ctx, messageQueueUrl, messageQueueAccessToken, lastMessageId
func (_m *Client) GetMessage(ctx context.Context, messageQueueUrl string, messageQueueAccessToken string, lastMessageId int64) (*actions.RunnerScaleSetMessage, error) {
	ret := _m.Called(ctx, messageQueueUrl, messageQueueAccessToken, lastMessageId)

	var r0 *actions.RunnerScaleSetMessage
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, int64) (*actions.RunnerScaleSetMessage, error)); ok {
		return rf(ctx, messageQueueUrl, messageQueueAccessToken, lastMessageId)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string, int64) *actions.RunnerScaleSetMessage); ok {
		r0 = rf(ctx, messageQueueUrl, messageQueueAccessToken, lastMessageId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*actions.RunnerScaleSetMessage)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string, int64) error); ok {
		r1 = rf(ctx, messageQueueUrl, messageQueueAccessToken, lastMessageId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RefreshMessageSession provides a mock function with given fields: ctx, runnerScaleSetId, sessionId
func (_m *Client) RefreshMessageSession(ctx context.Context, runnerScaleSetId int, sessionId *uuid.UUID) (*actions.RunnerScaleSetSession, error) {
	ret := _m.Called(ctx, runnerScaleSetId, sessionId)

	var r0 *actions.RunnerScaleSetSession
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int, *uuid.UUID) (*actions.RunnerScaleSetSession, error)); ok {
		return rf(ctx, runnerScaleSetId, sessionId)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int, *uuid.UUID) *actions.RunnerScaleSetSession); ok {
		r0 = rf(ctx, runnerScaleSetId, sessionId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*actions.RunnerScaleSetSession)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int, *uuid.UUID) error); ok {
		r1 = rf(ctx, runnerScaleSetId, sessionId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewClient creates a new instance of Client. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *Client {
	mock := &Client{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
