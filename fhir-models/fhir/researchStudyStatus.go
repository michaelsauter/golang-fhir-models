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

// ResearchStudyStatus is documented here http://hl7.org/fhir/ValueSet/research-study-status
type ResearchStudyStatus int

const (
	ResearchStudyStatusActive ResearchStudyStatus = iota
	ResearchStudyStatusAdministrativelyCompleted
	ResearchStudyStatusApproved
	ResearchStudyStatusClosedToAccrual
	ResearchStudyStatusClosedToAccrualAndIntervention
	ResearchStudyStatusCompleted
	ResearchStudyStatusDisapproved
	ResearchStudyStatusInReview
	ResearchStudyStatusTemporarilyClosedToAccrual
	ResearchStudyStatusTemporarilyClosedToAccrualAndIntervention
	ResearchStudyStatusWithdrawn
)

func (code ResearchStudyStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *ResearchStudyStatus) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "active":
		*code = ResearchStudyStatusActive
	case "administratively-completed":
		*code = ResearchStudyStatusAdministrativelyCompleted
	case "approved":
		*code = ResearchStudyStatusApproved
	case "closed-to-accrual":
		*code = ResearchStudyStatusClosedToAccrual
	case "closed-to-accrual-and-intervention":
		*code = ResearchStudyStatusClosedToAccrualAndIntervention
	case "completed":
		*code = ResearchStudyStatusCompleted
	case "disapproved":
		*code = ResearchStudyStatusDisapproved
	case "in-review":
		*code = ResearchStudyStatusInReview
	case "temporarily-closed-to-accrual":
		*code = ResearchStudyStatusTemporarilyClosedToAccrual
	case "temporarily-closed-to-accrual-and-intervention":
		*code = ResearchStudyStatusTemporarilyClosedToAccrualAndIntervention
	case "withdrawn":
		*code = ResearchStudyStatusWithdrawn
	default:
		return fmt.Errorf("unknown ResearchStudyStatus code `%s`", s)
	}
	return nil
}
func (code ResearchStudyStatus) String() string {
	return code.Code()
}
func (code ResearchStudyStatus) Code() string {
	switch code {
	case ResearchStudyStatusActive:
		return "active"
	case ResearchStudyStatusAdministrativelyCompleted:
		return "administratively-completed"
	case ResearchStudyStatusApproved:
		return "approved"
	case ResearchStudyStatusClosedToAccrual:
		return "closed-to-accrual"
	case ResearchStudyStatusClosedToAccrualAndIntervention:
		return "closed-to-accrual-and-intervention"
	case ResearchStudyStatusCompleted:
		return "completed"
	case ResearchStudyStatusDisapproved:
		return "disapproved"
	case ResearchStudyStatusInReview:
		return "in-review"
	case ResearchStudyStatusTemporarilyClosedToAccrual:
		return "temporarily-closed-to-accrual"
	case ResearchStudyStatusTemporarilyClosedToAccrualAndIntervention:
		return "temporarily-closed-to-accrual-and-intervention"
	case ResearchStudyStatusWithdrawn:
		return "withdrawn"
	}
	return "<unknown>"
}
func (code ResearchStudyStatus) Display() string {
	switch code {
	case ResearchStudyStatusActive:
		return "Active"
	case ResearchStudyStatusAdministrativelyCompleted:
		return "Administratively Completed"
	case ResearchStudyStatusApproved:
		return "Approved"
	case ResearchStudyStatusClosedToAccrual:
		return "Closed to Accrual"
	case ResearchStudyStatusClosedToAccrualAndIntervention:
		return "Closed to Accrual and Intervention"
	case ResearchStudyStatusCompleted:
		return "Completed"
	case ResearchStudyStatusDisapproved:
		return "Disapproved"
	case ResearchStudyStatusInReview:
		return "In Review"
	case ResearchStudyStatusTemporarilyClosedToAccrual:
		return "Temporarily Closed to Accrual"
	case ResearchStudyStatusTemporarilyClosedToAccrualAndIntervention:
		return "Temporarily Closed to Accrual and Intervention"
	case ResearchStudyStatusWithdrawn:
		return "Withdrawn"
	}
	return "<unknown>"
}
func (code ResearchStudyStatus) Definition() string {
	switch code {
	case ResearchStudyStatusActive:
		return "Study is opened for accrual."
	case ResearchStudyStatusAdministrativelyCompleted:
		return "Study is completed prematurely and will not resume; patients are no longer examined nor treated."
	case ResearchStudyStatusApproved:
		return "Protocol is approved by the review board."
	case ResearchStudyStatusClosedToAccrual:
		return "Study is closed for accrual; patients can be examined and treated."
	case ResearchStudyStatusClosedToAccrualAndIntervention:
		return "Study is closed to accrual and intervention, i.e. the study is closed to enrollment, all study subjects have completed treatment or intervention but are still being followed according to the primary objective of the study."
	case ResearchStudyStatusCompleted:
		return "Study is closed to accrual and intervention, i.e. the study is closed to enrollment, all study subjects have completed treatment\nor intervention but are still being followed according to the primary objective of the study."
	case ResearchStudyStatusDisapproved:
		return "Protocol was disapproved by the review board."
	case ResearchStudyStatusInReview:
		return "Protocol is submitted to the review board for approval."
	case ResearchStudyStatusTemporarilyClosedToAccrual:
		return "Study is temporarily closed for accrual; can be potentially resumed in the future; patients can be examined and treated."
	case ResearchStudyStatusTemporarilyClosedToAccrualAndIntervention:
		return "Study is temporarily closed for accrual and intervention and potentially can be resumed in the future."
	case ResearchStudyStatusWithdrawn:
		return "Protocol was withdrawn by the lead organization."
	}
	return "<unknown>"
}
