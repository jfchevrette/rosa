/*
Copyright (c) 2020 Red Hat, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// IMPORTANT: This file has been generated automatically, refrain from modifying it manually as all
// your changes will be lost when the file is generated again.

package v1 // github.com/openshift-online/ocm-sdk-go/authorizations/v1

// SelfTermsReviewRequestBuilder contains the data and logic needed to build 'self_terms_review_request' objects.
//
// Representation of Red Hat's Terms and Conditions for using OpenShift Dedicated and Amazon Red Hat OpenShift [Terms]
// review requests.
type SelfTermsReviewRequestBuilder struct {
	bitmap_   uint32
	eventCode string
	siteCode  string
}

// NewSelfTermsReviewRequest creates a new builder of 'self_terms_review_request' objects.
func NewSelfTermsReviewRequest() *SelfTermsReviewRequestBuilder {
	return &SelfTermsReviewRequestBuilder{}
}

// Empty returns true if the builder is empty, i.e. no attribute has a value.
func (b *SelfTermsReviewRequestBuilder) Empty() bool {
	return b == nil || b.bitmap_ == 0
}

// EventCode sets the value of the 'event_code' attribute to the given value.
//
//
func (b *SelfTermsReviewRequestBuilder) EventCode(value string) *SelfTermsReviewRequestBuilder {
	b.eventCode = value
	b.bitmap_ |= 1
	return b
}

// SiteCode sets the value of the 'site_code' attribute to the given value.
//
//
func (b *SelfTermsReviewRequestBuilder) SiteCode(value string) *SelfTermsReviewRequestBuilder {
	b.siteCode = value
	b.bitmap_ |= 2
	return b
}

// Copy copies the attributes of the given object into this builder, discarding any previous values.
func (b *SelfTermsReviewRequestBuilder) Copy(object *SelfTermsReviewRequest) *SelfTermsReviewRequestBuilder {
	if object == nil {
		return b
	}
	b.bitmap_ = object.bitmap_
	b.eventCode = object.eventCode
	b.siteCode = object.siteCode
	return b
}

// Build creates a 'self_terms_review_request' object using the configuration stored in the builder.
func (b *SelfTermsReviewRequestBuilder) Build() (object *SelfTermsReviewRequest, err error) {
	object = new(SelfTermsReviewRequest)
	object.bitmap_ = b.bitmap_
	object.eventCode = b.eventCode
	object.siteCode = b.siteCode
	return
}
