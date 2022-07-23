// Copyright 2020 The gVisor Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at //
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by "stringer -type headerType ."; DO NOT EDIT.

package stack

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[linkHeader-0]
	_ = x[networkHeader-1]
	_ = x[transportHeader-2]
	_ = x[numHeaderType-3]
}

const _headerType_name = "linkHeadernetworkHeadertransportHeadernumHeaderType"

var _headerType_index = [...]uint8{0, 10, 23, 38, 51}

func (i headerType) String() string {
	if i < 0 || i >= headerType(len(_headerType_index)-1) {
		return "headerType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _headerType_name[_headerType_index[i]:_headerType_index[i+1]]
}
