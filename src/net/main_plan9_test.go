// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package net

import "os/exec"

func installTestHooks() {}

func uninstallTestHooks() {}

// forceCloseSockets must be called only from TestMain.
func forceCloseSockets() {}

func enableSocketConnect() {}

func disableSocketConnect(network string) {}

func addCmdInheritedHandle(cmd *exec.Cmd, fd uintptr) {}
