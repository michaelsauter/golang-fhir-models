// Copyright 2019 The Samply Development Community
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package fhir

import (
	"encoding/json"
	"fmt"
	"strings"
)

// THIS FILE IS GENERATED BY https://github.com/lifebox-healthcare/golang-fhir-models
// PLEASE DO NOT EDIT BY HAND

// TaskStatus is documented here http://hl7.org/fhir/ValueSet/task-status
type TaskStatus int

const (
	TaskStatusDraft TaskStatus = iota
	TaskStatusRequested
	TaskStatusReceived
	TaskStatusAccepted
	TaskStatusRejected
	TaskStatusReady
	TaskStatusCancelled
	TaskStatusInProgress
	TaskStatusOnHold
	TaskStatusFailed
	TaskStatusCompleted
	TaskStatusEnteredInError
)

func (code TaskStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *TaskStatus) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "draft":
		*code = TaskStatusDraft
	case "requested":
		*code = TaskStatusRequested
	case "received":
		*code = TaskStatusReceived
	case "accepted":
		*code = TaskStatusAccepted
	case "rejected":
		*code = TaskStatusRejected
	case "ready":
		*code = TaskStatusReady
	case "cancelled":
		*code = TaskStatusCancelled
	case "in-progress":
		*code = TaskStatusInProgress
	case "on-hold":
		*code = TaskStatusOnHold
	case "failed":
		*code = TaskStatusFailed
	case "completed":
		*code = TaskStatusCompleted
	case "entered-in-error":
		*code = TaskStatusEnteredInError
	default:
		return fmt.Errorf("unknown TaskStatus code `%s`", s)
	}
	return nil
}
func (code TaskStatus) String() string {
	return code.Code()
}
func (code TaskStatus) Code() string {
	switch code {
	case TaskStatusDraft:
		return "draft"
	case TaskStatusRequested:
		return "requested"
	case TaskStatusReceived:
		return "received"
	case TaskStatusAccepted:
		return "accepted"
	case TaskStatusRejected:
		return "rejected"
	case TaskStatusReady:
		return "ready"
	case TaskStatusCancelled:
		return "cancelled"
	case TaskStatusInProgress:
		return "in-progress"
	case TaskStatusOnHold:
		return "on-hold"
	case TaskStatusFailed:
		return "failed"
	case TaskStatusCompleted:
		return "completed"
	case TaskStatusEnteredInError:
		return "entered-in-error"
	}
	return "<unknown>"
}
func (code TaskStatus) Display() string {
	switch code {
	case TaskStatusDraft:
		return "Draft"
	case TaskStatusRequested:
		return "Requested"
	case TaskStatusReceived:
		return "Received"
	case TaskStatusAccepted:
		return "Accepted"
	case TaskStatusRejected:
		return "Rejected"
	case TaskStatusReady:
		return "Ready"
	case TaskStatusCancelled:
		return "Cancelled"
	case TaskStatusInProgress:
		return "In Progress"
	case TaskStatusOnHold:
		return "On Hold"
	case TaskStatusFailed:
		return "Failed"
	case TaskStatusCompleted:
		return "Completed"
	case TaskStatusEnteredInError:
		return "Entered in Error"
	}
	return "<unknown>"
}
func (code TaskStatus) Definition() string {
	switch code {
	case TaskStatusDraft:
		return "The task is not yet ready to be acted upon."
	case TaskStatusRequested:
		return "The task is ready to be acted upon and action is sought."
	case TaskStatusReceived:
		return "A potential performer has claimed ownership of the task and is evaluating whether to perform it."
	case TaskStatusAccepted:
		return "The potential performer has agreed to execute the task but has not yet started work."
	case TaskStatusRejected:
		return "The potential performer who claimed ownership of the task has decided not to execute it prior to performing any action."
	case TaskStatusReady:
		return "The task is ready to be performed, but no action has yet been taken.  Used in place of requested/received/accepted/rejected when request assignment and acceptance is a given."
	case TaskStatusCancelled:
		return "The task was not completed."
	case TaskStatusInProgress:
		return "The task has been started but is not yet complete."
	case TaskStatusOnHold:
		return "The task has been started but work has been paused."
	case TaskStatusFailed:
		return "The task was attempted but could not be completed due to some error."
	case TaskStatusCompleted:
		return "The task has been completed."
	case TaskStatusEnteredInError:
		return "The task should never have existed and is retained only because of the possibility it may have used."
	}
	return "<unknown>"
}
