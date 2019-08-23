/*
Copyright 2019 Google LLC

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

package v1alpha1

import (
	"knative.dev/pkg/apis"
)

// GetCondition returns the condition currently associated with the given type, or nil.
func (s *SchedulerStatus) GetCondition(t apis.ConditionType) *apis.Condition {
	return schedulerCondSet.Manage(s).GetCondition(t)
}

// IsReady returns true if the resource is ready overall.
func (s *SchedulerStatus) IsReady() bool {
	return schedulerCondSet.Manage(s).IsHappy()
}

// InitializeConditions sets relevant unset conditions to Unknown state.
func (s *SchedulerStatus) InitializeConditions() {
	schedulerCondSet.Manage(s).InitializeConditions()
}

// MarkPullSubscriptionNotReady sets the condition that the underlying PullSubscription
// is not ready and why
func (s *SchedulerStatus) MarkPullSubscriptionNotReady(reason, messageFormat string, messageA ...interface{}) {
	schedulerCondSet.Manage(s).MarkFalse(PullSubscriptionReady, reason, messageFormat, messageA...)
}

// MarkPullSubscriptionReady sets the condition that the underlying PullSubscription is ready
func (s *SchedulerStatus) MarkPullSubscriptionReady() {
	schedulerCondSet.Manage(s).MarkTrue(PullSubscriptionReady)
}

// MarkTopicNotReady sets the condition that the Topic was not created and why
func (s *SchedulerStatus) MarkTopicNotReady(reason, messageFormat string, messageA ...interface{}) {
	schedulerCondSet.Manage(s).MarkFalse(TopicReady, reason, messageFormat, messageA...)
}

// MarkTopicReady sets the condition that the underlying Topic was created successfully
func (s *SchedulerStatus) MarkTopicReady() {
	schedulerCondSet.Manage(s).MarkTrue(TopicReady)
}

// MarkSchedulerNotReady sets the condition that the GCS has been configured to send Notifications
func (s *SchedulerStatus) MarkSchedulerJobNotReady(reason, messageFormat string, messageA ...interface{}) {
	schedulerCondSet.Manage(s).MarkFalse(SchedulerJobReady, reason, messageFormat, messageA...)
}

func (s *SchedulerStatus) MarkSchedulerJobReady() {
	schedulerCondSet.Manage(s).MarkTrue(SchedulerJobReady)
}
