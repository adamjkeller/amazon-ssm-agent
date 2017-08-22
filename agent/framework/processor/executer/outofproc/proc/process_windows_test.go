// Copyright 2016 Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may not
// use this file except in compliance with the License. A copy of the
// License is located at
//
// http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
// either express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// +build windows

// Package process wraps up the os.Process interface and also provides os-specific process lookup functions
package proc

import (
	"os/exec"
	"testing"

	"time"

	"github.com/aws/amazon-ssm-agent/agent/log"
	"src/github.com/stretchr/testify/assert"
)

//TODO add process start time
func TestIsProcessExists(t *testing.T) {
	cmd := exec.Command("cmd", "timeout", "100")
	err := cmd.Start()
	assert.NoError(t, err)
	pid := cmd.Process.Pid
	//do not call wait in case process is recycled
	assert.True(t, IsProcessExists(log.NewMockLog(), pid, time.Now()))
}
