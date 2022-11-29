// Copyright 2022 eatmoreapple.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package redis

// Scanner is the interface that wraps the Scan method.
type Scanner interface {
	Scan(v any) error
}
