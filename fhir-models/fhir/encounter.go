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
	"bytes"
	"encoding/json"
)

// THIS FILE IS GENERATED BY https://github.com/michaelsauter/golang-fhir-models
// PLEASE DO NOT EDIT BY HAND

// Encounter is documented here http://hl7.org/fhir/StructureDefinition/Encounter
type Encounter struct {
	ID                *string                   `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta                     `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string                   `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string                   `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative                `bson:"text,omitempty" json:"text,omitempty"`
	Extension         []Extension               `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension               `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier        []Identifier              `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status            EncounterStatus           `bson:"status" json:"status"`
	StatusHistory     []EncounterStatusHistory  `bson:"statusHistory,omitempty" json:"statusHistory,omitempty"`
	Class             Coding                    `bson:"class" json:"class"`
	ClassHistory      []EncounterClassHistory   `bson:"classHistory,omitempty" json:"classHistory,omitempty"`
	Type              []CodeableConcept         `bson:"type,omitempty" json:"type,omitempty"`
	ServiceType       *CodeableConcept          `bson:"serviceType,omitempty" json:"serviceType,omitempty"`
	Priority          *CodeableConcept          `bson:"priority,omitempty" json:"priority,omitempty"`
	Subject           *Reference                `bson:"subject,omitempty" json:"subject,omitempty"`
	EpisodeOfCare     []Reference               `bson:"episodeOfCare,omitempty" json:"episodeOfCare,omitempty"`
	BasedOn           []Reference               `bson:"basedOn,omitempty" json:"basedOn,omitempty"`
	Participant       []EncounterParticipant    `bson:"participant,omitempty" json:"participant,omitempty"`
	Appointment       []Reference               `bson:"appointment,omitempty" json:"appointment,omitempty"`
	Period            *Period                   `bson:"period,omitempty" json:"period,omitempty"`
	Length            *Duration                 `bson:"length,omitempty" json:"length,omitempty"`
	ReasonCode        []CodeableConcept         `bson:"reasonCode,omitempty" json:"reasonCode,omitempty"`
	ReasonReference   []Reference               `bson:"reasonReference,omitempty" json:"reasonReference,omitempty"`
	Diagnosis         []EncounterDiagnosis      `bson:"diagnosis,omitempty" json:"diagnosis,omitempty"`
	Account           []Reference               `bson:"account,omitempty" json:"account,omitempty"`
	Hospitalization   *EncounterHospitalization `bson:"hospitalization,omitempty" json:"hospitalization,omitempty"`
	Location          []EncounterLocation       `bson:"location,omitempty" json:"location,omitempty"`
	ServiceProvider   *Reference                `bson:"serviceProvider,omitempty" json:"serviceProvider,omitempty"`
	PartOf            *Reference                `bson:"partOf,omitempty" json:"partOf,omitempty"`
}
type EncounterStatusHistory struct {
	ID                *string         `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Status            EncounterStatus `bson:"status" json:"status"`
	Period            Period          `bson:"period" json:"period"`
}
type EncounterClassHistory struct {
	ID                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Class             Coding      `bson:"class" json:"class"`
	Period            Period      `bson:"period" json:"period"`
}
type EncounterParticipant struct {
	ID                *string           `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension       `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension       `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Type              []CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	Period            *Period           `bson:"period,omitempty" json:"period,omitempty"`
	Individual        *Reference        `bson:"individual,omitempty" json:"individual,omitempty"`
}
type EncounterDiagnosis struct {
	ID                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Condition         Reference        `bson:"condition" json:"condition"`
	Use               *CodeableConcept `bson:"use,omitempty" json:"use,omitempty"`
	Rank              *int             `bson:"rank,omitempty" json:"rank,omitempty"`
}
type EncounterHospitalization struct {
	ID                     *string           `bson:"id,omitempty" json:"id,omitempty"`
	Extension              []Extension       `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension      []Extension       `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	PreAdmissionIdentifier *Identifier       `bson:"preAdmissionIdentifier,omitempty" json:"preAdmissionIdentifier,omitempty"`
	Origin                 *Reference        `bson:"origin,omitempty" json:"origin,omitempty"`
	AdmitSource            *CodeableConcept  `bson:"admitSource,omitempty" json:"admitSource,omitempty"`
	ReAdmission            *CodeableConcept  `bson:"reAdmission,omitempty" json:"reAdmission,omitempty"`
	DietPreference         []CodeableConcept `bson:"dietPreference,omitempty" json:"dietPreference,omitempty"`
	SpecialCourtesy        []CodeableConcept `bson:"specialCourtesy,omitempty" json:"specialCourtesy,omitempty"`
	SpecialArrangement     []CodeableConcept `bson:"specialArrangement,omitempty" json:"specialArrangement,omitempty"`
	Destination            *Reference        `bson:"destination,omitempty" json:"destination,omitempty"`
	DischargeDisposition   *CodeableConcept  `bson:"dischargeDisposition,omitempty" json:"dischargeDisposition,omitempty"`
}
type EncounterLocation struct {
	ID                *string                  `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension              `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension              `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Location          Reference                `bson:"location" json:"location"`
	Status            *EncounterLocationStatus `bson:"status,omitempty" json:"status,omitempty"`
	PhysicalType      *CodeableConcept         `bson:"physicalType,omitempty" json:"physicalType,omitempty"`
	Period            *Period                  `bson:"period,omitempty" json:"period,omitempty"`
}
type OtherEncounter Encounter

// MarshalJSON marshals the given Encounter as JSON into a byte slice
func (r Encounter) MarshalJSON() ([]byte, error) {
	buffer := bytes.Buffer{}
	enc := json.NewEncoder(&buffer)
	enc.SetEscapeHTML(false)
	err := enc.Encode(struct {
		OtherEncounter
		ResourceType string `json:"resourceType"`
	}{
		OtherEncounter: OtherEncounter(r),
		ResourceType:   "Encounter",
	})
	return buffer.Bytes(), err
}

// UnmarshalEncounter unmarshals a Encounter.
func UnmarshalEncounter(b []byte) (Encounter, error) {
	var encounter Encounter
	if err := json.Unmarshal(b, &encounter); err != nil {
		return encounter, err
	}
	return encounter, nil
}
