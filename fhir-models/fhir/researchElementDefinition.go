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

// ResearchElementDefinition is documented here http://hl7.org/fhir/StructureDefinition/ResearchElementDefinition
type ResearchElementDefinition struct {
	Id                     *string                                   `bson:"id,omitempty" json:"id,omitempty"`
	Meta                   *Meta                                     `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules          *string                                   `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language               *string                                   `bson:"language,omitempty" json:"language,omitempty"`
	Text                   *Narrative                                `bson:"text,omitempty" json:"text,omitempty"`
	Extension              []Extension                               `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension      []Extension                               `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Url                    *string                                   `bson:"url,omitempty" json:"url,omitempty"`
	Identifier             []Identifier                              `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Version                *string                                   `bson:"version,omitempty" json:"version,omitempty"`
	Name                   *string                                   `bson:"name,omitempty" json:"name,omitempty"`
	Title                  *string                                   `bson:"title,omitempty" json:"title,omitempty"`
	ShortTitle             *string                                   `bson:"shortTitle,omitempty" json:"shortTitle,omitempty"`
	Subtitle               *string                                   `bson:"subtitle,omitempty" json:"subtitle,omitempty"`
	Status                 PublicationStatus                         `bson:"status" json:"status"`
	Experimental           *bool                                     `bson:"experimental,omitempty" json:"experimental,omitempty"`
	SubjectCodeableConcept *CodeableConcept                          `bson:"subjectCodeableConcept,omitempty" json:"subjectCodeableConcept,omitempty"`
	SubjectReference       *Reference                                `bson:"subjectReference,omitempty" json:"subjectReference,omitempty"`
	Date                   *string                                   `bson:"date,omitempty" json:"date,omitempty"`
	Publisher              *string                                   `bson:"publisher,omitempty" json:"publisher,omitempty"`
	Contact                []ContactDetail                           `bson:"contact,omitempty" json:"contact,omitempty"`
	Description            *string                                   `bson:"description,omitempty" json:"description,omitempty"`
	Comment                []string                                  `bson:"comment,omitempty" json:"comment,omitempty"`
	UseContext             []UsageContext                            `bson:"useContext,omitempty" json:"useContext,omitempty"`
	Jurisdiction           []CodeableConcept                         `bson:"jurisdiction,omitempty" json:"jurisdiction,omitempty"`
	Purpose                *string                                   `bson:"purpose,omitempty" json:"purpose,omitempty"`
	Usage                  *string                                   `bson:"usage,omitempty" json:"usage,omitempty"`
	Copyright              *string                                   `bson:"copyright,omitempty" json:"copyright,omitempty"`
	ApprovalDate           *string                                   `bson:"approvalDate,omitempty" json:"approvalDate,omitempty"`
	LastReviewDate         *string                                   `bson:"lastReviewDate,omitempty" json:"lastReviewDate,omitempty"`
	EffectivePeriod        *Period                                   `bson:"effectivePeriod,omitempty" json:"effectivePeriod,omitempty"`
	Topic                  []CodeableConcept                         `bson:"topic,omitempty" json:"topic,omitempty"`
	Author                 []ContactDetail                           `bson:"author,omitempty" json:"author,omitempty"`
	Editor                 []ContactDetail                           `bson:"editor,omitempty" json:"editor,omitempty"`
	Reviewer               []ContactDetail                           `bson:"reviewer,omitempty" json:"reviewer,omitempty"`
	Endorser               []ContactDetail                           `bson:"endorser,omitempty" json:"endorser,omitempty"`
	RelatedArtifact        []RelatedArtifact                         `bson:"relatedArtifact,omitempty" json:"relatedArtifact,omitempty"`
	Library                []string                                  `bson:"library,omitempty" json:"library,omitempty"`
	Type                   ResearchElementType                       `bson:"type" json:"type"`
	VariableType           *EvidenceVariableType                     `bson:"variableType,omitempty" json:"variableType,omitempty"`
	Characteristic         []ResearchElementDefinitionCharacteristic `bson:"characteristic" json:"characteristic"`
}
type ResearchElementDefinitionCharacteristic struct {
	Id                                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension                         []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension                 []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	DefinitionCodeableConcept         CodeableConcept  `bson:"definitionCodeableConcept" json:"definitionCodeableConcept"`
	DefinitionCanonical               string           `bson:"definitionCanonical" json:"definitionCanonical"`
	DefinitionExpression              Expression       `bson:"definitionExpression" json:"definitionExpression"`
	DefinitionDataRequirement         DataRequirement  `bson:"definitionDataRequirement" json:"definitionDataRequirement"`
	UsageContext                      []UsageContext   `bson:"usageContext,omitempty" json:"usageContext,omitempty"`
	Exclude                           *bool            `bson:"exclude,omitempty" json:"exclude,omitempty"`
	UnitOfMeasure                     *CodeableConcept `bson:"unitOfMeasure,omitempty" json:"unitOfMeasure,omitempty"`
	StudyEffectiveDescription         *string          `bson:"studyEffectiveDescription,omitempty" json:"studyEffectiveDescription,omitempty"`
	StudyEffectiveDateTime            *string          `bson:"studyEffectiveDateTime,omitempty" json:"studyEffectiveDateTime,omitempty"`
	StudyEffectivePeriod              *Period          `bson:"studyEffectivePeriod,omitempty" json:"studyEffectivePeriod,omitempty"`
	StudyEffectiveDuration            *Duration        `bson:"studyEffectiveDuration,omitempty" json:"studyEffectiveDuration,omitempty"`
	StudyEffectiveTiming              *Timing          `bson:"studyEffectiveTiming,omitempty" json:"studyEffectiveTiming,omitempty"`
	StudyEffectiveTimeFromStart       *Duration        `bson:"studyEffectiveTimeFromStart,omitempty" json:"studyEffectiveTimeFromStart,omitempty"`
	StudyEffectiveGroupMeasure        *GroupMeasure    `bson:"studyEffectiveGroupMeasure,omitempty" json:"studyEffectiveGroupMeasure,omitempty"`
	ParticipantEffectiveDescription   *string          `bson:"participantEffectiveDescription,omitempty" json:"participantEffectiveDescription,omitempty"`
	ParticipantEffectiveDateTime      *string          `bson:"participantEffectiveDateTime,omitempty" json:"participantEffectiveDateTime,omitempty"`
	ParticipantEffectivePeriod        *Period          `bson:"participantEffectivePeriod,omitempty" json:"participantEffectivePeriod,omitempty"`
	ParticipantEffectiveDuration      *Duration        `bson:"participantEffectiveDuration,omitempty" json:"participantEffectiveDuration,omitempty"`
	ParticipantEffectiveTiming        *Timing          `bson:"participantEffectiveTiming,omitempty" json:"participantEffectiveTiming,omitempty"`
	ParticipantEffectiveTimeFromStart *Duration        `bson:"participantEffectiveTimeFromStart,omitempty" json:"participantEffectiveTimeFromStart,omitempty"`
	ParticipantEffectiveGroupMeasure  *GroupMeasure    `bson:"participantEffectiveGroupMeasure,omitempty" json:"participantEffectiveGroupMeasure,omitempty"`
}
type OtherResearchElementDefinition ResearchElementDefinition

// MarshalJSON marshals the given ResearchElementDefinition as JSON into a byte slice
func (r ResearchElementDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherResearchElementDefinition
		ResourceType string `json:"resourceType"`
	}{
		OtherResearchElementDefinition: OtherResearchElementDefinition(r),
		ResourceType:                   "ResearchElementDefinition",
	})
}

// UnmarshalResearchElementDefinition unmarshals a ResearchElementDefinition.
func UnmarshalResearchElementDefinition(b []byte) (ResearchElementDefinition, error) {
	var researchElementDefinition ResearchElementDefinition
	if err := json.Unmarshal(b, &researchElementDefinition); err != nil {
		return researchElementDefinition, err
	}
	return researchElementDefinition, nil
}
