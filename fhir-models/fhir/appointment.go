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

// Appointment is documented here http://hl7.org/fhir/StructureDefinition/Appointment
type Appointment struct {
	Id                    *string                  `bson:"id,omitempty" json:"id,omitempty"`
	Meta                  *Meta                    `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules         *string                  `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language              *string                  `bson:"language,omitempty" json:"language,omitempty"`
	Text                  *Narrative               `bson:"text,omitempty" json:"text,omitempty"`
	Extension             []Extension              `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension     []Extension              `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier            []Identifier             `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status                AppointmentStatus        `bson:"status" json:"status"`
	CancelationReason     *CodeableConcept         `bson:"cancelationReason,omitempty" json:"cancelationReason,omitempty"`
	ServiceCategory       []CodeableConcept        `bson:"serviceCategory,omitempty" json:"serviceCategory,omitempty"`
	ServiceType           []CodeableConcept        `bson:"serviceType,omitempty" json:"serviceType,omitempty"`
	Specialty             []CodeableConcept        `bson:"specialty,omitempty" json:"specialty,omitempty"`
	AppointmentType       *CodeableConcept         `bson:"appointmentType,omitempty" json:"appointmentType,omitempty"`
	ReasonCode            []CodeableConcept        `bson:"reasonCode,omitempty" json:"reasonCode,omitempty"`
	ReasonReference       []Reference              `bson:"reasonReference,omitempty" json:"reasonReference,omitempty"`
	Priority              *int                     `bson:"priority,omitempty" json:"priority,omitempty"`
	Description           *string                  `bson:"description,omitempty" json:"description,omitempty"`
	SupportingInformation []Reference              `bson:"supportingInformation,omitempty" json:"supportingInformation,omitempty"`
	Start                 *string                  `bson:"start,omitempty" json:"start,omitempty"`
	End                   *string                  `bson:"end,omitempty" json:"end,omitempty"`
	MinutesDuration       *int                     `bson:"minutesDuration,omitempty" json:"minutesDuration,omitempty"`
	Slot                  []Reference              `bson:"slot,omitempty" json:"slot,omitempty"`
	Created               *string                  `bson:"created,omitempty" json:"created,omitempty"`
	Comment               *string                  `bson:"comment,omitempty" json:"comment,omitempty"`
	PatientInstruction    *string                  `bson:"patientInstruction,omitempty" json:"patientInstruction,omitempty"`
	BasedOn               []Reference              `bson:"basedOn,omitempty" json:"basedOn,omitempty"`
	Participant           []AppointmentParticipant `bson:"participant" json:"participant"`
	RequestedPeriod       []Period                 `bson:"requestedPeriod,omitempty" json:"requestedPeriod,omitempty"`
}
type AppointmentParticipant struct {
	Id                *string              `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension          `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension          `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Type              []CodeableConcept    `bson:"type,omitempty" json:"type,omitempty"`
	Actor             *Reference           `bson:"actor,omitempty" json:"actor,omitempty"`
	Required          *ParticipantRequired `bson:"required,omitempty" json:"required,omitempty"`
	Status            ParticipationStatus  `bson:"status" json:"status"`
	Period            *Period              `bson:"period,omitempty" json:"period,omitempty"`
}
type OtherAppointment Appointment

// MarshalJSON marshals the given Appointment as JSON into a byte slice
func (r Appointment) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherAppointment
		ResourceType string `json:"resourceType"`
	}{
		OtherAppointment: OtherAppointment(r),
		ResourceType:     "Appointment",
	})
}

// UnmarshalAppointment unmarshals a Appointment.
func UnmarshalAppointment(b []byte) (Appointment, error) {
	var appointment Appointment
	if err := json.Unmarshal(b, &appointment); err != nil {
		return appointment, err
	}
	return appointment, nil
}
