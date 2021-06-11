// Copyright (c) Roman Atachiants and contributors. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for details.

package column

import (
	"sync"

	"github.com/kelindar/bitmap"
)

// Bitmaps represents a pool of bitmaps
var bitmaps = &sync.Pool{
	New: func() interface{} {
		return &bitmap.Bitmap{}
	},
}

// aquireBitmap acquires a bitmap for a transaction
func aquireBitmap(source *bitmap.Bitmap) *bitmap.Bitmap {
	b := bitmaps.Get().(*bitmap.Bitmap)
	source.Clone(b)
	return b
}

// releaseBitmap releases a bitmap back to the pool
func releaseBitmap(b *bitmap.Bitmap) {
	bitmaps.Put(b)
}

// --------------------------- Transaction ----------------------------

// Txn represents a transaction which supports filtering and projection.
type Txn struct {
	owner *Collection
	index *bitmap.Bitmap
}

// With applies a logical AND operation to the current query and the specified index.
func (txn Txn) With(index string, extra ...string) Txn {
	if idx, ok := txn.owner.cols[index]; ok {
		txn.index.And(idx.Bitmap())
	} else {
		txn.index.Clear()
	}

	// go through extra indexes
	for _, e := range extra {
		if idx, ok := txn.owner.cols[e]; ok {
			txn.index.And(idx.Bitmap())
		} else {
			txn.index.Clear()
		}
	}
	return txn
}

// Without applies a logical AND NOT operation to the current query and the specified index.
func (txn Txn) Without(index string, extra ...string) Txn {
	if idx, ok := txn.owner.cols[index]; ok {
		txn.index.AndNot(idx.Bitmap())
	}

	// go through extra indexes
	for _, e := range extra {
		if idx, ok := txn.owner.cols[e]; ok {
			txn.index.AndNot(idx.Bitmap())
		}
	}
	return txn
}

// Union computes a union between the current query and the specified index.
func (txn Txn) Union(index string, extra ...string) Txn {
	if idx, ok := txn.owner.cols[index]; ok {
		txn.index.Or(idx.Bitmap())
	}

	// go through extra indexes
	for _, e := range extra {
		if idx, ok := txn.owner.cols[e]; ok {
			txn.index.Or(idx.Bitmap())
		}
	}
	return txn
}

// WithValue applies a filter predicate over values for a specific properties. It filters
// down the items in the query.
func (txn Txn) WithValue(property string, predicate func(v interface{}) bool) Txn {
	if p, ok := txn.owner.cols[property]; ok {
		txn.index.Filter(func(x uint32) bool {
			if v, ok := p.Value(x); ok {
				return predicate(v)
			}
			return false
		})
	}
	return txn
}

// WithFloat64 filters down the values based on the specified predicate. The column for
// this filter must be numerical and convertible to float64.
func (txn Txn) WithFloat64(property string, predicate func(v float64) bool) Txn {
	if p, ok := txn.owner.cols[property]; ok {
		if n, ok := p.(numerical); ok {
			txn.index.Filter(func(x uint32) bool {
				if v, ok := n.Float64(x); ok {
					return predicate(v)
				}
				return false
			})
		}
	}
	return txn
}

// WithInt64 filters down the values based on the specified predicate. The column for
// this filter must be numerical and convertible to int64.
func (txn Txn) WithInt64(property string, predicate func(v int64) bool) Txn {
	if p, ok := txn.owner.cols[property]; ok {
		if n, ok := p.(numerical); ok {
			txn.index.Filter(func(x uint32) bool {
				if v, ok := n.Int64(x); ok {
					return predicate(v)
				}
				return false
			})
		}
	}
	return txn
}

// WithUint64 filters down the values based on the specified predicate. The column for
// this filter must be numerical and convertible to uint64.
func (txn Txn) WithUint64(property string, predicate func(v uint64) bool) Txn {
	if p, ok := txn.owner.cols[property]; ok {
		if n, ok := p.(numerical); ok {
			txn.index.Filter(func(x uint32) bool {
				if v, ok := n.Uint64(x); ok {
					return predicate(v)
				}
				return false
			})
		}
	}
	return txn
}

// WithString filters down the values based on the specified predicate. The column for
// this filter must be a string.
func (txn Txn) WithString(property string, predicate func(v string) bool) Txn {
	return txn.WithValue(property, func(v interface{}) bool {
		return predicate(v.(string))
	})
}

// Count returns the number of objects matching the query
func (txn Txn) Count() int {
	return int(txn.index.Count())
}

// Range iterates over the objects with the given properties, but does not perform any
// locking.
func (txn Txn) Range(fn func(Selector) bool) {
	txn.index.Range(func(x uint32) bool {
		return fn(Selector{
			index: x,
			owner: txn.owner,
		})
	})
}

// --------------------------- Selector ----------------------------

// Selector represents a column selector which can retrieve a value by a column name.
type Selector struct {
	index uint32
	owner *Collection
}

// Value reads a value for a current row at a given column.
func (s *Selector) Value(column string) interface{} {
	if c, ok := s.owner.cols[column]; ok {
		v, _ := c.Value(s.index)
		return v
	}
	return nil
}

// String reads a string value for a current row at a given column.
func (s *Selector) String(column string) string {
	if c, ok := s.owner.cols[column]; ok {
		if v, ok := c.Value(s.index); ok {
			return v.(string)
		}
	}
	return ""
}

// Float64 reads a float64 value for a current row at a given column.
func (s *Selector) Float64(column string) float64 {
	if c, ok := s.owner.cols[column]; ok {
		if n, ok := c.(numerical); ok {
			v, _ := n.Float64(s.index)
			return v
		}
	}
	return 0
}

// Int64 reads an int64 value for a current row at a given column.
func (s *Selector) Int64(column string) int64 {
	if c, ok := s.owner.cols[column]; ok {
		if n, ok := c.(numerical); ok {
			v, _ := n.Int64(s.index)
			return v
		}
	}
	return 0
}

// Uint64 reads a uint64 value for a current row at a given column.
func (s *Selector) Uint64(column string) uint64 {
	if c, ok := s.owner.cols[column]; ok {
		if n, ok := c.(numerical); ok {
			v, _ := n.Uint64(s.index)
			return v
		}
	}
	return 0
}

// Bool reads a boolean value for a current row at a given column.
func (s *Selector) Bool(column string) bool {
	if c, ok := s.owner.cols[column]; ok {
		return c.Contains(s.index)
	}
	return false
}