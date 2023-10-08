// Copyright (c) 2023. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package leaderelection

type HeartBeatRequest struct {
	Lead uint64
	Term uint64
}

type HeartBeatResponse struct {
	Success bool
}
