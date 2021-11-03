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

// THIS FILE IS GENERATED BY https://github.com/lifebox-healthcare/golang-fhir-models
// PLEASE DO NOT EDIT BY HAND

// Timing is documented here http://hl7.org/fhir/StructureDefinition/Timing
type Timing struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Event             []string         `bson:"event,omitempty" json:"event,omitempty"`
	Repeat            *TimingRepeat    `bson:"repeat,omitempty" json:"repeat,omitempty"`
	Code              *CodeableConcept `bson:"code,omitempty" json:"code,omitempty"`
}
type TimingRepeat struct {
	Id             *string      `bson:"id,omitempty" json:"id,omitempty"`
	Extension      []Extension  `bson:"extension,omitempty" json:"extension,omitempty"`
	BoundsDuration *Duration    `bson:"boundsDuration,omitempty" json:"boundsDuration,omitempty"`
	BoundsRange    *Range       `bson:"boundsRange,omitempty" json:"boundsRange,omitempty"`
	BoundsPeriod   *Period      `bson:"boundsPeriod,omitempty" json:"boundsPeriod,omitempty"`
	Count          *int         `bson:"count,omitempty" json:"count,omitempty"`
	CountMax       *int         `bson:"countMax,omitempty" json:"countMax,omitempty"`
	Duration       *string      `bson:"duration,omitempty" json:"duration,omitempty"`
	DurationMax    *string      `bson:"durationMax,omitempty" json:"durationMax,omitempty"`
	DurationUnit   *string      `bson:"durationUnit,omitempty" json:"durationUnit,omitempty"`
	Frequency      *int         `bson:"frequency,omitempty" json:"frequency,omitempty"`
	FrequencyMax   *int         `bson:"frequencyMax,omitempty" json:"frequencyMax,omitempty"`
	Period         *string      `bson:"period,omitempty" json:"period,omitempty"`
	PeriodMax      *string      `bson:"periodMax,omitempty" json:"periodMax,omitempty"`
	PeriodUnit     *string      `bson:"periodUnit,omitempty" json:"periodUnit,omitempty"`
	DayOfWeek      []DaysOfWeek `bson:"dayOfWeek,omitempty" json:"dayOfWeek,omitempty"`
	TimeOfDay      []string     `bson:"timeOfDay,omitempty" json:"timeOfDay,omitempty"`
	When           []string     `bson:"when,omitempty" json:"when,omitempty"`
	Offset         *int         `bson:"offset,omitempty" json:"offset,omitempty"`
}
