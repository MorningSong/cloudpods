// Copyright 2019 Yunion
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

package multicloud

import (
	api "yunion.io/x/cloudmux/pkg/apis/compute"
)

type SDisk struct {
	SVirtualResourceBase
	SBillingBase
}

func (self *SDisk) GetExtSnapshotPolicyIds() ([]string, error) {
	return []string{""}, nil
}

func (self *SDisk) GetIStorageId() string {
	return ""
}

func (self *SDisk) GetIops() int {
	return 0
}

func (self *SDisk) GetPreallocation() string {
	return api.DISK_PREALLOCATION_OFF
}
