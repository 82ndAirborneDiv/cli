// +build windows

// Copyright 2016 Alan Shreve github@inconshreveable.
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

package updater

import (
	"syscall"
	"unsafe"

	"github.com/ActiveState/cli/internal/logging"
)

func hideFile(path string) error {
	kernel32 := syscall.NewLazyDLL("kernel32.dll")
	setFileAttributes := kernel32.NewProc("SetFileAttributesW")

	r1, _, err := setFileAttributes.Call(uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(path))), 2)
	if err != nil {
		logging.Error("Encountered error setting file attributes: %v", err)
	}

	if r1 == 0 {
		return err
	} else {
		return nil
	}
}