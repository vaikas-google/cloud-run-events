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

package resources

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	duckv1beta1 "knative.dev/pkg/apis/duck/v1beta1"
	apisv1alpha1 "knative.dev/pkg/apis/v1alpha1"

	"github.com/google/knative-gcp/pkg/apis/events/v1alpha1"
	pubsubv1alpha1 "github.com/google/knative-gcp/pkg/apis/pubsub/v1alpha1"
)

func TestMakePullSubscription(t *testing.T) {
	source := &v1alpha1.Scheduler{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "bucket-name",
			Namespace: "bucket-namespace",
			UID:       "bucket-uid",
		},
		Spec: v1alpha1.SchedulerSpec{
			Bucket:  "this-bucket",
			Project: "project-123",
			GCSSecret: corev1.SecretKeySelector{
				LocalObjectReference: corev1.LocalObjectReference{
					Name: "eventing-secret-name",
				},
				Key: "eventing-secret-key",
			},
			SourceSpec: duckv1beta1.SourceSpec{
				Sink: apisv1alpha1.Destination{
					ObjectReference: &corev1.ObjectReference{
						APIVersion: "v1",
						Kind:       "Kitchen",
						Name:       "sink",
					},
				},
				CloudEventOverrides: &duckv1beta1.CloudEventOverrides{
					Extensions: map[string]string{
						"foo": "bar",
					},
				},
			},
		},
	}

	got := MakePullSubscription(source, "topic-abc")

	yes := true
	want := &pubsubv1alpha1.PullSubscription{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: "bucket-namespace",
			Name:      "bucket-name",
			Labels: map[string]string{
				"receive-adapter": "scheduler.events.cloud.run",
			},
			OwnerReferences: []metav1.OwnerReference{{
				APIVersion:         "events.cloud.run/v1alpha1",
				Kind:               "Scheduler",
				Name:               "bucket-name",
				UID:                "bucket-uid",
				Controller:         &yes,
				BlockOwnerDeletion: &yes,
			}},
		},
		Spec: pubsubv1alpha1.PullSubscriptionSpec{
			Secret: &corev1.SecretKeySelector{
				LocalObjectReference: corev1.LocalObjectReference{
					Name: "eventing-secret-name",
				},
				Key: "eventing-secret-key",
			},
			Project: "project-123",
			Topic:   "topic-abc",
			Sink: apisv1alpha1.Destination{
				ObjectReference: &corev1.ObjectReference{
					APIVersion: "v1",
					Kind:       "Kitchen",
					Name:       "sink",
				},
			},
			CloudEventOverrides: &pubsubv1alpha1.CloudEventOverrides{
				Extensions: map[string]string{
					"foo": "bar",
				},
			},
		},
	}

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("unexpected (-want, +got) = %v", diff)
	}
}
