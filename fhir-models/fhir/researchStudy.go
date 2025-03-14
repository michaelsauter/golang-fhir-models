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

import "encoding/json"

// THIS FILE IS GENERATED BY https://github.com/lifebox-healthcare/golang-fhir-models
// PLEASE DO NOT EDIT BY HAND

// ResearchStudy is documented here http://hl7.org/fhir/StructureDefinition/ResearchStudy
type ResearchStudy struct {
	Id                    *string                  `bson:"id,omitempty" json:"id,omitempty"`
	Meta                  *Meta                    `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules         *string                  `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language              *string                  `bson:"language,omitempty" json:"language,omitempty"`
	Text                  *Narrative               `bson:"text,omitempty" json:"text,omitempty"`
	Extension             []Extension              `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension     []Extension              `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier            []Identifier             `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Title                 *string                  `bson:"title,omitempty" json:"title,omitempty"`
	Protocol              []Reference              `bson:"protocol,omitempty" json:"protocol,omitempty"`
	PartOf                []Reference              `bson:"partOf,omitempty" json:"partOf,omitempty"`
	Status                ResearchStudyStatus      `bson:"status" json:"status"`
	PrimaryPurposeType    *CodeableConcept         `bson:"primaryPurposeType,omitempty" json:"primaryPurposeType,omitempty"`
	Phase                 *CodeableConcept         `bson:"phase,omitempty" json:"phase,omitempty"`
	Category              []CodeableConcept        `bson:"category,omitempty" json:"category,omitempty"`
	Focus                 []CodeableConcept        `bson:"focus,omitempty" json:"focus,omitempty"`
	Condition             []CodeableConcept        `bson:"condition,omitempty" json:"condition,omitempty"`
	Contact               []ContactDetail          `bson:"contact,omitempty" json:"contact,omitempty"`
	RelatedArtifact       []RelatedArtifact        `bson:"relatedArtifact,omitempty" json:"relatedArtifact,omitempty"`
	Keyword               []CodeableConcept        `bson:"keyword,omitempty" json:"keyword,omitempty"`
	Location              []CodeableConcept        `bson:"location,omitempty" json:"location,omitempty"`
	Description           *string                  `bson:"description,omitempty" json:"description,omitempty"`
	Enrollment            []Reference              `bson:"enrollment,omitempty" json:"enrollment,omitempty"`
	Period                *Period                  `bson:"period,omitempty" json:"period,omitempty"`
	Sponsor               *Reference               `bson:"sponsor,omitempty" json:"sponsor,omitempty"`
	PrincipalInvestigator *Reference               `bson:"principalInvestigator,omitempty" json:"principalInvestigator,omitempty"`
	Site                  []Reference              `bson:"site,omitempty" json:"site,omitempty"`
	ReasonStopped         *CodeableConcept         `bson:"reasonStopped,omitempty" json:"reasonStopped,omitempty"`
	Note                  []Annotation             `bson:"note,omitempty" json:"note,omitempty"`
	Arm                   []ResearchStudyArm       `bson:"arm,omitempty" json:"arm,omitempty"`
	Objective             []ResearchStudyObjective `bson:"objective,omitempty" json:"objective,omitempty"`
}
type ResearchStudyArm struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Name              string           `bson:"name" json:"name"`
	Type              *CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	Description       *string          `bson:"description,omitempty" json:"description,omitempty"`
}
type ResearchStudyObjective struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Name              *string          `bson:"name,omitempty" json:"name,omitempty"`
	Type              *CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
}
type OtherResearchStudy ResearchStudy

// MarshalJSON marshals the given ResearchStudy as JSON into a byte slice
func (r ResearchStudy) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherResearchStudy
		ResourceType string `json:"resourceType"`
	}{
		OtherResearchStudy: OtherResearchStudy(r),
		ResourceType:       "ResearchStudy",
	})
}

// UnmarshalResearchStudy unmarshals a ResearchStudy.
func UnmarshalResearchStudy(b []byte) (ResearchStudy, error) {
	var researchStudy ResearchStudy
	if err := json.Unmarshal(b, &researchStudy); err != nil {
		return researchStudy, err
	}
	return researchStudy, nil
}
