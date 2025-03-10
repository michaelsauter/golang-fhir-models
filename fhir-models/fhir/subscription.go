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

// Subscription is documented here http://hl7.org/fhir/StructureDefinition/Subscription
type Subscription struct {
	Id                *string             `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta               `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string             `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string             `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative          `bson:"text,omitempty" json:"text,omitempty"`
	Extension         []Extension         `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension         `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Status            SubscriptionStatus  `bson:"status" json:"status"`
	Contact           []ContactPoint      `bson:"contact,omitempty" json:"contact,omitempty"`
	End               *string             `bson:"end,omitempty" json:"end,omitempty"`
	Reason            string              `bson:"reason" json:"reason"`
	Criteria          string              `bson:"criteria" json:"criteria"`
	Error             *string             `bson:"error,omitempty" json:"error,omitempty"`
	Channel           SubscriptionChannel `bson:"channel" json:"channel"`
}
type SubscriptionChannel struct {
	Id                *string                 `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension             `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension             `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Type              SubscriptionChannelType `bson:"type" json:"type"`
	Endpoint          *string                 `bson:"endpoint,omitempty" json:"endpoint,omitempty"`
	Payload           *string                 `bson:"payload,omitempty" json:"payload,omitempty"`
	Header            []string                `bson:"header,omitempty" json:"header,omitempty"`
}
type OtherSubscription Subscription

// MarshalJSON marshals the given Subscription as JSON into a byte slice
func (r Subscription) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherSubscription
		ResourceType string `json:"resourceType"`
	}{
		OtherSubscription: OtherSubscription(r),
		ResourceType:      "Subscription",
	})
}

// UnmarshalSubscription unmarshals a Subscription.
func UnmarshalSubscription(b []byte) (Subscription, error) {
	var subscription Subscription
	if err := json.Unmarshal(b, &subscription); err != nil {
		return subscription, err
	}
	return subscription, nil
}
