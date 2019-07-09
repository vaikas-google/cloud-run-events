/*
Copyright 2019 The Knative Authors

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
	"context"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestPullSubscriptionDefaults(t *testing.T) {
	tests := []struct {
		name  string
		start *PullSubscription
		want  *PullSubscription
	}{{
		name: "non-nil structured",
		start: &PullSubscription{
			Spec: PullSubscriptionSpec{
				Mode: ModeCloudEventsStructured,
			},
		},
		want: &PullSubscription{
			Spec: PullSubscriptionSpec{
				Mode:              ModeCloudEventsStructured,
				RetentionDuration: &defaultRetentionDuration,
				AckDeadline:       &defaultAckDeadline,
			},
		},
	}, {
		name: "non-nil push",
		start: &PullSubscription{
			Spec: PullSubscriptionSpec{
				Mode: ModePushCompatible,
			},
		},
		want: &PullSubscription{
			Spec: PullSubscriptionSpec{
				Mode:              ModePushCompatible,
				RetentionDuration: &defaultRetentionDuration,
				AckDeadline:       &defaultAckDeadline,
			},
		},
	}, {
		name: "non-nil invalid",
		start: &PullSubscription{
			Spec: PullSubscriptionSpec{
				Mode: "invalid",
			},
		},
		want: &PullSubscription{
			Spec: PullSubscriptionSpec{
				Mode:              ModeCloudEventsBinary,
				RetentionDuration: &defaultRetentionDuration,
				AckDeadline:       &defaultAckDeadline,
			},
		},
	}, {
		name: "nil",
		start: &PullSubscription{
			ObjectMeta: metav1.ObjectMeta{},
			Spec:       PullSubscriptionSpec{},
		},
		want: &PullSubscription{
			Spec: PullSubscriptionSpec{
				Mode:              ModeCloudEventsBinary,
				RetentionDuration: &defaultRetentionDuration,
				AckDeadline:       &defaultAckDeadline,
			},
		},
	}}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := test.start
			got.SetDefaults(context.Background())

			if diff := cmp.Diff(test.want, got); diff != "" {
				t.Errorf("failed to get expected (-want, +got) = %v", diff)
			}
		})
	}
}

func TestPullSubscriptionDefaults_NoChange(t *testing.T) {
	days2 := 2 * 24 * time.Hour
	secs60 := 60 * time.Second
	want := &PullSubscription{
		Spec: PullSubscriptionSpec{
			Mode:              ModeCloudEventsBinary,
			AckDeadline:       &secs60,
			RetentionDuration: &days2,
		},
	}

	got := want.DeepCopy()
	got.SetDefaults(context.Background())
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("failed to get expected (-want, +got) = %v", diff)
	}
}
