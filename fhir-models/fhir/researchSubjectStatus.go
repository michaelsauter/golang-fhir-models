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

// ResearchSubjectStatus is documented here http://hl7.org/fhir/ValueSet/research-subject-status
type ResearchSubjectStatus int

const (
	ResearchSubjectStatusCandidate ResearchSubjectStatus = iota
	ResearchSubjectStatusEligible
	ResearchSubjectStatusFollowUp
	ResearchSubjectStatusIneligible
	ResearchSubjectStatusNotRegistered
	ResearchSubjectStatusOffStudy
	ResearchSubjectStatusOnStudy
	ResearchSubjectStatusOnStudyIntervention
	ResearchSubjectStatusOnStudyObservation
	ResearchSubjectStatusPendingOnStudy
	ResearchSubjectStatusPotentialCandidate
	ResearchSubjectStatusScreening
	ResearchSubjectStatusWithdrawn
)

func (code ResearchSubjectStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *ResearchSubjectStatus) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "candidate":
		*code = ResearchSubjectStatusCandidate
	case "eligible":
		*code = ResearchSubjectStatusEligible
	case "follow-up":
		*code = ResearchSubjectStatusFollowUp
	case "ineligible":
		*code = ResearchSubjectStatusIneligible
	case "not-registered":
		*code = ResearchSubjectStatusNotRegistered
	case "off-study":
		*code = ResearchSubjectStatusOffStudy
	case "on-study":
		*code = ResearchSubjectStatusOnStudy
	case "on-study-intervention":
		*code = ResearchSubjectStatusOnStudyIntervention
	case "on-study-observation":
		*code = ResearchSubjectStatusOnStudyObservation
	case "pending-on-study":
		*code = ResearchSubjectStatusPendingOnStudy
	case "potential-candidate":
		*code = ResearchSubjectStatusPotentialCandidate
	case "screening":
		*code = ResearchSubjectStatusScreening
	case "withdrawn":
		*code = ResearchSubjectStatusWithdrawn
	default:
		return fmt.Errorf("unknown ResearchSubjectStatus code `%s`", s)
	}
	return nil
}
func (code ResearchSubjectStatus) String() string {
	return code.Code()
}
func (code ResearchSubjectStatus) Code() string {
	switch code {
	case ResearchSubjectStatusCandidate:
		return "candidate"
	case ResearchSubjectStatusEligible:
		return "eligible"
	case ResearchSubjectStatusFollowUp:
		return "follow-up"
	case ResearchSubjectStatusIneligible:
		return "ineligible"
	case ResearchSubjectStatusNotRegistered:
		return "not-registered"
	case ResearchSubjectStatusOffStudy:
		return "off-study"
	case ResearchSubjectStatusOnStudy:
		return "on-study"
	case ResearchSubjectStatusOnStudyIntervention:
		return "on-study-intervention"
	case ResearchSubjectStatusOnStudyObservation:
		return "on-study-observation"
	case ResearchSubjectStatusPendingOnStudy:
		return "pending-on-study"
	case ResearchSubjectStatusPotentialCandidate:
		return "potential-candidate"
	case ResearchSubjectStatusScreening:
		return "screening"
	case ResearchSubjectStatusWithdrawn:
		return "withdrawn"
	}
	return "<unknown>"
}
func (code ResearchSubjectStatus) Display() string {
	switch code {
	case ResearchSubjectStatusCandidate:
		return "Candidate"
	case ResearchSubjectStatusEligible:
		return "Eligible"
	case ResearchSubjectStatusFollowUp:
		return "Follow-up"
	case ResearchSubjectStatusIneligible:
		return "Ineligible"
	case ResearchSubjectStatusNotRegistered:
		return "Not Registered"
	case ResearchSubjectStatusOffStudy:
		return "Off-study"
	case ResearchSubjectStatusOnStudy:
		return "On-study"
	case ResearchSubjectStatusOnStudyIntervention:
		return "On-study-intervention"
	case ResearchSubjectStatusOnStudyObservation:
		return "On-study-observation"
	case ResearchSubjectStatusPendingOnStudy:
		return "Pending on-study"
	case ResearchSubjectStatusPotentialCandidate:
		return "Potential Candidate"
	case ResearchSubjectStatusScreening:
		return "Screening"
	case ResearchSubjectStatusWithdrawn:
		return "Withdrawn"
	}
	return "<unknown>"
}
func (code ResearchSubjectStatus) Definition() string {
	switch code {
	case ResearchSubjectStatusCandidate:
		return "An identified person that can be considered for inclusion in a study."
	case ResearchSubjectStatusEligible:
		return "A person that has met the eligibility criteria for inclusion in a study."
	case ResearchSubjectStatusFollowUp:
		return "A person is no longer receiving study intervention and/or being evaluated with tests and procedures according to the protocol, but they are being monitored on a protocol-prescribed schedule."
	case ResearchSubjectStatusIneligible:
		return "A person who did not meet one or more criteria required for participation in a study is considered to have failed screening or\nis ineligible for the study."
	case ResearchSubjectStatusNotRegistered:
		return "A person for whom registration was not completed."
	case ResearchSubjectStatusOffStudy:
		return "A person that has ended their participation on a study either because their treatment/observation is complete or through not\nresponding, withdrawal, non-compliance and/or adverse event."
	case ResearchSubjectStatusOnStudy:
		return "A person that is enrolled or registered on a study."
	case ResearchSubjectStatusOnStudyIntervention:
		return "The person is receiving the treatment or participating in an activity (e.g. yoga, diet, etc.) that the study is evaluating."
	case ResearchSubjectStatusOnStudyObservation:
		return "The subject is being evaluated via tests and assessments according to the study calendar, but is not receiving any intervention. Note that this state is study-dependent and might not exist in all studies.  A synonym for this is \"short-term follow-up\"."
	case ResearchSubjectStatusPendingOnStudy:
		return "A person is pre-registered for a study."
	case ResearchSubjectStatusPotentialCandidate:
		return "A person that is potentially eligible for participation in the study."
	case ResearchSubjectStatusScreening:
		return "A person who is being evaluated for eligibility for a study."
	case ResearchSubjectStatusWithdrawn:
		return "The person has withdrawn their participation in the study before registration."
	}
	return "<unknown>"
}
