// Automatically generated by pegomock. DO NOT EDIT!
// Source: github.com/runatlantis/atlantis/server/events (interfaces: AtlantisWorkspaceLocker)

package mocks

import (
	"reflect"

	pegomock "github.com/petergtz/pegomock"
)

type MockAtlantisWorkspaceLocker struct {
	fail func(message string, callerSkip ...int)
}

func NewMockAtlantisWorkspaceLocker() *MockAtlantisWorkspaceLocker {
	return &MockAtlantisWorkspaceLocker{fail: pegomock.GlobalFailHandler}
}

func (mock *MockAtlantisWorkspaceLocker) TryLock(repoFullName string, workspace string, pullNum int) bool {
	params := []pegomock.Param{repoFullName, workspace, pullNum}
	result := pegomock.GetGenericMockFrom(mock).Invoke("TryLock", params, []reflect.Type{reflect.TypeOf((*bool)(nil)).Elem()})
	var ret0 bool
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(bool)
		}
	}
	return ret0
}

func (mock *MockAtlantisWorkspaceLocker) Unlock(repoFullName string, workspace string, pullNum int) {
	params := []pegomock.Param{repoFullName, workspace, pullNum}
	pegomock.GetGenericMockFrom(mock).Invoke("Unlock", params, []reflect.Type{})
}

func (mock *MockAtlantisWorkspaceLocker) VerifyWasCalledOnce() *VerifierAtlantisWorkspaceLocker {
	return &VerifierAtlantisWorkspaceLocker{mock, pegomock.Times(1), nil}
}

func (mock *MockAtlantisWorkspaceLocker) VerifyWasCalled(invocationCountMatcher pegomock.Matcher) *VerifierAtlantisWorkspaceLocker {
	return &VerifierAtlantisWorkspaceLocker{mock, invocationCountMatcher, nil}
}

func (mock *MockAtlantisWorkspaceLocker) VerifyWasCalledInOrder(invocationCountMatcher pegomock.Matcher, inOrderContext *pegomock.InOrderContext) *VerifierAtlantisWorkspaceLocker {
	return &VerifierAtlantisWorkspaceLocker{mock, invocationCountMatcher, inOrderContext}
}

type VerifierAtlantisWorkspaceLocker struct {
	mock                   *MockAtlantisWorkspaceLocker
	invocationCountMatcher pegomock.Matcher
	inOrderContext         *pegomock.InOrderContext
}

func (verifier *VerifierAtlantisWorkspaceLocker) TryLock(repoFullName string, workspace string, pullNum int) *AtlantisWorkspaceLocker_TryLock_OngoingVerification {
	params := []pegomock.Param{repoFullName, workspace, pullNum}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "TryLock", params)
	return &AtlantisWorkspaceLocker_TryLock_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type AtlantisWorkspaceLocker_TryLock_OngoingVerification struct {
	mock              *MockAtlantisWorkspaceLocker
	methodInvocations []pegomock.MethodInvocation
}

func (c *AtlantisWorkspaceLocker_TryLock_OngoingVerification) GetCapturedArguments() (string, string, int) {
	repoFullName, workspace, pullNum := c.GetAllCapturedArguments()
	return repoFullName[len(repoFullName)-1], workspace[len(workspace)-1], pullNum[len(pullNum)-1]
}

func (c *AtlantisWorkspaceLocker_TryLock_OngoingVerification) GetAllCapturedArguments() (_param0 []string, _param1 []string, _param2 []int) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]string, len(params[0]))
		for u, param := range params[0] {
			_param0[u] = param.(string)
		}
		_param1 = make([]string, len(params[1]))
		for u, param := range params[1] {
			_param1[u] = param.(string)
		}
		_param2 = make([]int, len(params[2]))
		for u, param := range params[2] {
			_param2[u] = param.(int)
		}
	}
	return
}

func (verifier *VerifierAtlantisWorkspaceLocker) Unlock(repoFullName string, workspace string, pullNum int) *AtlantisWorkspaceLocker_Unlock_OngoingVerification {
	params := []pegomock.Param{repoFullName, workspace, pullNum}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "Unlock", params)
	return &AtlantisWorkspaceLocker_Unlock_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type AtlantisWorkspaceLocker_Unlock_OngoingVerification struct {
	mock              *MockAtlantisWorkspaceLocker
	methodInvocations []pegomock.MethodInvocation
}

func (c *AtlantisWorkspaceLocker_Unlock_OngoingVerification) GetCapturedArguments() (string, string, int) {
	repoFullName, workspace, pullNum := c.GetAllCapturedArguments()
	return repoFullName[len(repoFullName)-1], workspace[len(workspace)-1], pullNum[len(pullNum)-1]
}

func (c *AtlantisWorkspaceLocker_Unlock_OngoingVerification) GetAllCapturedArguments() (_param0 []string, _param1 []string, _param2 []int) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]string, len(params[0]))
		for u, param := range params[0] {
			_param0[u] = param.(string)
		}
		_param1 = make([]string, len(params[1]))
		for u, param := range params[1] {
			_param1[u] = param.(string)
		}
		_param2 = make([]int, len(params[2]))
		for u, param := range params[2] {
			_param2[u] = param.(int)
		}
	}
	return
}