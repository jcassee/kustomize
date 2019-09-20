// Copyright 2019 The Kubernetes Authors.
// SPDX-License-Identifier: Apache-2.0

package main_test

import (
	"testing"

	"sigs.k8s.io/kustomize/v3/pkg/kusttest"
	plugins_test "sigs.k8s.io/kustomize/v3/pkg/plugins/test"
)

func TestTestGeneratorPlugin(t *testing.T) {
	tc := plugins_test.NewEnvForTest(t).Set()
	defer tc.Reset()

	tc.BuildExecPlugin(
		"someteam.example.com", "v1", "TestGenerator")

	th := kusttest_test.NewKustTestPluginHarness(t, "/app")

	m := th.LoadAndRunGenerator(`
apiVersion: someteam.example.com/v1
kind: TestGenerator
metadata:
  name: whatever
`)
	th.AssertActualEqualsExpected(m, `
apiVersion: someteam.example.com/v1
kind: TestGenerator
metadata:
  name: whatever
`)
}
