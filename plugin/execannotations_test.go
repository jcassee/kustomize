// Copyright 2019 The Kubernetes Authors.
// SPDX-License-Identifier: Apache-2.0

package plugin

import (
	"testing"

	"sigs.k8s.io/kustomize/v3/pkg/kusttest"
	plugins_test "sigs.k8s.io/kustomize/v3/pkg/plugins/test"
	"sigs.k8s.io/kustomize/v3/pkg/types"
)

func TestNeedsHashAnnotation(t *testing.T) {
	tc := plugins_test.NewEnvForTest(t).Set()
	defer tc.Reset()

	tc.BuildExecPlugin(
		"someteam.example.com", "v1", "TestGenerator")

	th := kusttest_test.NewKustTestPluginHarness(t, "/app")

	rm := th.LoadAndRunGenerator(`
apiVersion: someteam.example.com/v1
kind: TestGenerator
metadata:
  name: whatever
  annotations:
    kustomize.config.k8s.io/needs-hash: "true"
`)

	for _, r := range rm.Resources() {
		if !r.NeedHashSuffix() {
			t.Fatalf("resources should need hash suffix: %v", r)
		}
	}
}

func TestBehaviorAnnotation(t *testing.T) {
	tc := plugins_test.NewEnvForTest(t).Set()
	defer tc.Reset()

	tc.BuildExecPlugin(
		"someteam.example.com", "v1", "TestGenerator")

	th := kusttest_test.NewKustTestPluginHarness(t, "/app")

	rm := th.LoadAndRunGenerator(`
apiVersion: someteam.example.com/v1
kind: TestGenerator
metadata:
  name: whatever
  annotations:
    kustomize.config.k8s.io/behavior: "merge"
`)

	for _, r := range rm.Resources() {
		if r.Behavior() != types.BehaviorMerge {
			t.Fatalf("resources should have behavior merge: %v", r)
		}
	}
}
