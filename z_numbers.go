// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package column

import "github.com/kelindar/bitmap"

// --------------------------- float32s ----------------------------

// columnFloat32 represents a generic column
type columnFloat32 struct {
	fill bitmap.Bitmap // The fill-list
	data []float32     // The actual values
}

// makeFloat32s creates a new vector or float32s
func makeFloat32s() Column {
	return &columnFloat32{
		fill: make(bitmap.Bitmap, 0, 4),
		data: make([]float32, 0, 64),
	}
}

// Set sets a value at a specified index
func (p *columnFloat32) Set(idx uint32, value interface{}) {
	size := uint32(len(p.data))
	for i := size; i <= idx; i++ {
		p.data = append(p.data, 0)
	}

	// Set the data at index
	p.fill.Set(idx)
	p.data[idx] = value.(float32)
}

// Value retrieves a value at a specified index
func (p *columnFloat32) Value(idx uint32) (interface{}, bool) {
	if idx >= uint32(len(p.data)) || !p.fill.Contains(idx) {
		return float32(0), false
	}

	return p.data[idx], true
}

// Float64 retrieves a float64 value at a specified index
func (p *columnFloat32) Float64(idx uint32) (float64, bool) {
	if idx >= uint32(len(p.data)) || !p.fill.Contains(idx) {
		return 0, false
	}

	return float64(p.data[idx]), true
}

// Int64 retrieves an int64 value at a specified index
func (p *columnFloat32) Int64(idx uint32) (int64, bool) {
	if idx >= uint32(len(p.data)) || !p.fill.Contains(idx) {
		return 0, false
	}

	return int64(p.data[idx]), true
}

// Uint64 retrieves an uint64 value at a specified index
func (p *columnFloat32) Uint64(idx uint32) (uint64, bool) {
	if idx >= uint32(len(p.data)) || !p.fill.Contains(idx) {
		return 0, false
	}

	return uint64(p.data[idx]), true
}

// Del removes a value at a specified index
func (p *columnFloat32) Del(idx uint32) {
	p.fill.Remove(idx)
	p.data[idx] = 0
}

// Bitmap returns the associated index bitmap.
func (p *columnFloat32) Bitmap() bitmap.Bitmap {
	return p.fill
}

// Contains checks whether the column has a value at a specified index.
func (p *columnFloat32) Contains(idx uint32) bool {
	return p.fill.Contains(idx)
}

// --------------------------- float64s ----------------------------

// columnFloat64 represents a generic column
type columnFloat64 struct {
	fill bitmap.Bitmap // The fill-list
	data []float64     // The actual values
}

// makeFloat64s creates a new vector or float64s
func makeFloat64s() Column {
	return &columnFloat64{
		fill: make(bitmap.Bitmap, 0, 4),
		data: make([]float64, 0, 64),
	}
}

// Set sets a value at a specified index
func (p *columnFloat64) Set(idx uint32, value interface{}) {
	size := uint32(len(p.data))
	for i := size; i <= idx; i++ {
		p.data = append(p.data, 0)
	}

	// Set the data at index
	p.fill.Set(idx)
	p.data[idx] = value.(float64)
}

// Value retrieves a value at a specified index
func (p *columnFloat64) Value(idx uint32) (interface{}, bool) {
	if idx >= uint32(len(p.data)) || !p.fill.Contains(idx) {
		return float64(0), false
	}

	return p.data[idx], true
}

// Float64 retrieves a float64 value at a specified index
func (p *columnFloat64) Float64(idx uint32) (float64, bool) {
	if idx >= uint32(len(p.data)) || !p.fill.Contains(idx) {
		return 0, false
	}

	return float64(p.data[idx]), true
}

// Int64 retrieves an int64 value at a specified index
func (p *columnFloat64) Int64(idx uint32) (int64, bool) {
	if idx >= uint32(len(p.data)) || !p.fill.Contains(idx) {
		return 0, false
	}

	return int64(p.data[idx]), true
}

// Uint64 retrieves an uint64 value at a specified index
func (p *columnFloat64) Uint64(idx uint32) (uint64, bool) {
	if idx >= uint32(len(p.data)) || !p.fill.Contains(idx) {
		return 0, false
	}

	return uint64(p.data[idx]), true
}

// Del removes a value at a specified index
func (p *columnFloat64) Del(idx uint32) {
	p.fill.Remove(idx)
	p.data[idx] = 0
}

// Bitmap returns the associated index bitmap.
func (p *columnFloat64) Bitmap() bitmap.Bitmap {
	return p.fill
}

// Contains checks whether the column has a value at a specified index.
func (p *columnFloat64) Contains(idx uint32) bool {
	return p.fill.Contains(idx)
}

// --------------------------- ints ----------------------------

// columnInt represents a generic column
type columnInt struct {
	fill bitmap.Bitmap // The fill-list
	data []int         // The actual values
}

// makeInts creates a new vector or ints
func makeInts() Column {
	return &columnInt{
		fill: make(bitmap.Bitmap, 0, 4),
		data: make([]int, 0, 64),
	}
}

// Set sets a value at a specified index
func (p *columnInt) Set(idx uint32, value interface{}) {
	size := uint32(len(p.data))
	for i := size; i <= idx; i++ {
		p.data = append(p.data, 0)
	}

	// Set the data at index
	p.fill.Set(idx)
	p.data[idx] = value.(int)
}

// Value retrieves a value at a specified index
func (p *columnInt) Value(idx uint32) (interface{}, bool) {
	if idx >= uint32(len(p.data)) || !p.fill.Contains(idx) {
		return int(0), false
	}

	return p.data[idx], true
}

// Float64 retrieves a float64 value at a specified index
func (p *columnInt) Float64(idx uint32) (float64, bool) {
	if idx >= uint32(len(p.data)) || !p.fill.Contains(idx) {
		return 0, false
	}

	return float64(p.data[idx]), true
}

// Int64 retrieves an int64 value at a specified index
func (p *columnInt) Int64(idx uint32) (int64, bool) {
	if idx >= uint32(len(p.data)) || !p.fill.Contains(idx) {
		return 0, false
	}

	return int64(p.data[idx]), true
}

// Uint64 retrieves an uint64 value at a specified index
func (p *columnInt) Uint64(idx uint32) (uint64, bool) {
	if idx >= uint32(len(p.data)) || !p.fill.Contains(idx) {
		return 0, false
	}

	return uint64(p.data[idx]), true
}

// Del removes a value at a specified index
func (p *columnInt) Del(idx uint32) {
	p.fill.Remove(idx)
	p.data[idx] = 0
}

// Bitmap returns the associated index bitmap.
func (p *columnInt) Bitmap() bitmap.Bitmap {
	return p.fill
}

// Contains checks whether the column has a value at a specified index.
func (p *columnInt) Contains(idx uint32) bool {
	return p.fill.Contains(idx)
}

// --------------------------- int16s ----------------------------

// columnInt16 represents a generic column
type columnInt16 struct {
	fill bitmap.Bitmap // The fill-list
	data []int16       // The actual values
}

// makeInt16s creates a new vector or int16s
func makeInt16s() Column {
	return &columnInt16{
		fill: make(bitmap.Bitmap, 0, 4),
		data: make([]int16, 0, 64),
	}
}

// Set sets a value at a specified index
func (p *columnInt16) Set(idx uint32, value interface{}) {
	size := uint32(len(p.data))
	for i := size; i <= idx; i++ {
		p.data = append(p.data, 0)
	}

	// Set the data at index
	p.fill.Set(idx)
	p.data[idx] = value.(int16)
}

// Value retrieves a value at a specified index
func (p *columnInt16) Value(idx uint32) (interface{}, bool) {
	if idx >= uint32(len(p.data)) || !p.fill.Contains(idx) {
		return int16(0), false
	}

	return p.data[idx], true
}

// Float64 retrieves a float64 value at a specified index
func (p *columnInt16) Float64(idx uint32) (float64, bool) {
	if idx >= uint32(len(p.data)) || !p.fill.Contains(idx) {
		return 0, false
	}

	return float64(p.data[idx]), true
}

// Int64 retrieves an int64 value at a specified index
func (p *columnInt16) Int64(idx uint32) (int64, bool) {
	if idx >= uint32(len(p.data)) || !p.fill.Contains(idx) {
		return 0, false
	}

	return int64(p.data[idx]), true
}

// Uint64 retrieves an uint64 value at a specified index
func (p *columnInt16) Uint64(idx uint32) (uint64, bool) {
	if idx >= uint32(len(p.data)) || !p.fill.Contains(idx) {
		return 0, false
	}

	return uint64(p.data[idx]), true
}

// Del removes a value at a specified index
func (p *columnInt16) Del(idx uint32) {
	p.fill.Remove(idx)
	p.data[idx] = 0
}

// Bitmap returns the associated index bitmap.
func (p *columnInt16) Bitmap() bitmap.Bitmap {
	return p.fill
}

// Contains checks whether the column has a value at a specified index.
func (p *columnInt16) Contains(idx uint32) bool {
	return p.fill.Contains(idx)
}

// --------------------------- int32s ----------------------------

// columnInt32 represents a generic column
type columnInt32 struct {
	fill bitmap.Bitmap // The fill-list
	data []int32       // The actual values
}

// makeInt32s creates a new vector or int32s
func makeInt32s() Column {
	return &columnInt32{
		fill: make(bitmap.Bitmap, 0, 4),
		data: make([]int32, 0, 64),
	}
}

// Set sets a value at a specified index
func (p *columnInt32) Set(idx uint32, value interface{}) {
	size := uint32(len(p.data))
	for i := size; i <= idx; i++ {
		p.data = append(p.data, 0)
	}

	// Set the data at index
	p.fill.Set(idx)
	p.data[idx] = value.(int32)
}

// Value retrieves a value at a specified index
func (p *columnInt32) Value(idx uint32) (interface{}, bool) {
	if idx >= uint32(len(p.data)) || !p.fill.Contains(idx) {
		return int32(0), false
	}

	return p.data[idx], true
}

// Float64 retrieves a float64 value at a specified index
func (p *columnInt32) Float64(idx uint32) (float64, bool) {
	if idx >= uint32(len(p.data)) || !p.fill.Contains(idx) {
		return 0, false
	}

	return float64(p.data[idx]), true
}

// Int64 retrieves an int64 value at a specified index
func (p *columnInt32) Int64(idx uint32) (int64, bool) {
	if idx >= uint32(len(p.data)) || !p.fill.Contains(idx) {
		return 0, false
	}

	return int64(p.data[idx]), true
}

// Uint64 retrieves an uint64 value at a specified index
func (p *columnInt32) Uint64(idx uint32) (uint64, bool) {
	if idx >= uint32(len(p.data)) || !p.fill.Contains(idx) {
		return 0, false
	}

	return uint64(p.data[idx]), true
}

// Del removes a value at a specified index
func (p *columnInt32) Del(idx uint32) {
	p.fill.Remove(idx)
	p.data[idx] = 0
}

// Bitmap returns the associated index bitmap.
func (p *columnInt32) Bitmap() bitmap.Bitmap {
	return p.fill
}

// Contains checks whether the column has a value at a specified index.
func (p *columnInt32) Contains(idx uint32) bool {
	return p.fill.Contains(idx)
}

// --------------------------- int64s ----------------------------

// columnInt64 represents a generic column
type columnInt64 struct {
	fill bitmap.Bitmap // The fill-list
	data []int64       // The actual values
}

// makeInt64s creates a new vector or int64s
func makeInt64s() Column {
	return &columnInt64{
		fill: make(bitmap.Bitmap, 0, 4),
		data: make([]int64, 0, 64),
	}
}

// Set sets a value at a specified index
func (p *columnInt64) Set(idx uint32, value interface{}) {
	size := uint32(len(p.data))
	for i := size; i <= idx; i++ {
		p.data = append(p.data, 0)
	}

	// Set the data at index
	p.fill.Set(idx)
	p.data[idx] = value.(int64)
}

// Value retrieves a value at a specified index
func (p *columnInt64) Value(idx uint32) (interface{}, bool) {
	if idx >= uint32(len(p.data)) || !p.fill.Contains(idx) {
		return int64(0), false
	}

	return p.data[idx], true
}

// Float64 retrieves a float64 value at a specified index
func (p *columnInt64) Float64(idx uint32) (float64, bool) {
	if idx >= uint32(len(p.data)) || !p.fill.Contains(idx) {
		return 0, false
	}

	return float64(p.data[idx]), true
}

// Int64 retrieves an int64 value at a specified index
func (p *columnInt64) Int64(idx uint32) (int64, bool) {
	if idx >= uint32(len(p.data)) || !p.fill.Contains(idx) {
		return 0, false
	}

	return int64(p.data[idx]), true
}

// Uint64 retrieves an uint64 value at a specified index
func (p *columnInt64) Uint64(idx uint32) (uint64, bool) {
	if idx >= uint32(len(p.data)) || !p.fill.Contains(idx) {
		return 0, false
	}

	return uint64(p.data[idx]), true
}

// Del removes a value at a specified index
func (p *columnInt64) Del(idx uint32) {
	p.fill.Remove(idx)
	p.data[idx] = 0
}

// Bitmap returns the associated index bitmap.
func (p *columnInt64) Bitmap() bitmap.Bitmap {
	return p.fill
}

// Contains checks whether the column has a value at a specified index.
func (p *columnInt64) Contains(idx uint32) bool {
	return p.fill.Contains(idx)
}

// --------------------------- uints ----------------------------

// columnUint represents a generic column
type columnUint struct {
	fill bitmap.Bitmap // The fill-list
	data []uint        // The actual values
}

// makeUints creates a new vector or uints
func makeUints() Column {
	return &columnUint{
		fill: make(bitmap.Bitmap, 0, 4),
		data: make([]uint, 0, 64),
	}
}

// Set sets a value at a specified index
func (p *columnUint) Set(idx uint32, value interface{}) {
	size := uint32(len(p.data))
	for i := size; i <= idx; i++ {
		p.data = append(p.data, 0)
	}

	// Set the data at index
	p.fill.Set(idx)
	p.data[idx] = value.(uint)
}

// Value retrieves a value at a specified index
func (p *columnUint) Value(idx uint32) (interface{}, bool) {
	if idx >= uint32(len(p.data)) || !p.fill.Contains(idx) {
		return uint(0), false
	}

	return p.data[idx], true
}

// Float64 retrieves a float64 value at a specified index
func (p *columnUint) Float64(idx uint32) (float64, bool) {
	if idx >= uint32(len(p.data)) || !p.fill.Contains(idx) {
		return 0, false
	}

	return float64(p.data[idx]), true
}

// Int64 retrieves an int64 value at a specified index
func (p *columnUint) Int64(idx uint32) (int64, bool) {
	if idx >= uint32(len(p.data)) || !p.fill.Contains(idx) {
		return 0, false
	}

	return int64(p.data[idx]), true
}

// Uint64 retrieves an uint64 value at a specified index
func (p *columnUint) Uint64(idx uint32) (uint64, bool) {
	if idx >= uint32(len(p.data)) || !p.fill.Contains(idx) {
		return 0, false
	}

	return uint64(p.data[idx]), true
}

// Del removes a value at a specified index
func (p *columnUint) Del(idx uint32) {
	p.fill.Remove(idx)
	p.data[idx] = 0
}

// Bitmap returns the associated index bitmap.
func (p *columnUint) Bitmap() bitmap.Bitmap {
	return p.fill
}

// Contains checks whether the column has a value at a specified index.
func (p *columnUint) Contains(idx uint32) bool {
	return p.fill.Contains(idx)
}

// --------------------------- uint16s ----------------------------

// columnUint16 represents a generic column
type columnUint16 struct {
	fill bitmap.Bitmap // The fill-list
	data []uint16      // The actual values
}

// makeUint16s creates a new vector or uint16s
func makeUint16s() Column {
	return &columnUint16{
		fill: make(bitmap.Bitmap, 0, 4),
		data: make([]uint16, 0, 64),
	}
}

// Set sets a value at a specified index
func (p *columnUint16) Set(idx uint32, value interface{}) {
	size := uint32(len(p.data))
	for i := size; i <= idx; i++ {
		p.data = append(p.data, 0)
	}

	// Set the data at index
	p.fill.Set(idx)
	p.data[idx] = value.(uint16)
}

// Value retrieves a value at a specified index
func (p *columnUint16) Value(idx uint32) (interface{}, bool) {
	if idx >= uint32(len(p.data)) || !p.fill.Contains(idx) {
		return uint16(0), false
	}

	return p.data[idx], true
}

// Float64 retrieves a float64 value at a specified index
func (p *columnUint16) Float64(idx uint32) (float64, bool) {
	if idx >= uint32(len(p.data)) || !p.fill.Contains(idx) {
		return 0, false
	}

	return float64(p.data[idx]), true
}

// Int64 retrieves an int64 value at a specified index
func (p *columnUint16) Int64(idx uint32) (int64, bool) {
	if idx >= uint32(len(p.data)) || !p.fill.Contains(idx) {
		return 0, false
	}

	return int64(p.data[idx]), true
}

// Uint64 retrieves an uint64 value at a specified index
func (p *columnUint16) Uint64(idx uint32) (uint64, bool) {
	if idx >= uint32(len(p.data)) || !p.fill.Contains(idx) {
		return 0, false
	}

	return uint64(p.data[idx]), true
}

// Del removes a value at a specified index
func (p *columnUint16) Del(idx uint32) {
	p.fill.Remove(idx)
	p.data[idx] = 0
}

// Bitmap returns the associated index bitmap.
func (p *columnUint16) Bitmap() bitmap.Bitmap {
	return p.fill
}

// Contains checks whether the column has a value at a specified index.
func (p *columnUint16) Contains(idx uint32) bool {
	return p.fill.Contains(idx)
}

// --------------------------- uint32s ----------------------------

// columnUint32 represents a generic column
type columnUint32 struct {
	fill bitmap.Bitmap // The fill-list
	data []uint32      // The actual values
}

// makeUint32s creates a new vector or uint32s
func makeUint32s() Column {
	return &columnUint32{
		fill: make(bitmap.Bitmap, 0, 4),
		data: make([]uint32, 0, 64),
	}
}

// Set sets a value at a specified index
func (p *columnUint32) Set(idx uint32, value interface{}) {
	size := uint32(len(p.data))
	for i := size; i <= idx; i++ {
		p.data = append(p.data, 0)
	}

	// Set the data at index
	p.fill.Set(idx)
	p.data[idx] = value.(uint32)
}

// Value retrieves a value at a specified index
func (p *columnUint32) Value(idx uint32) (interface{}, bool) {
	if idx >= uint32(len(p.data)) || !p.fill.Contains(idx) {
		return uint32(0), false
	}

	return p.data[idx], true
}

// Float64 retrieves a float64 value at a specified index
func (p *columnUint32) Float64(idx uint32) (float64, bool) {
	if idx >= uint32(len(p.data)) || !p.fill.Contains(idx) {
		return 0, false
	}

	return float64(p.data[idx]), true
}

// Int64 retrieves an int64 value at a specified index
func (p *columnUint32) Int64(idx uint32) (int64, bool) {
	if idx >= uint32(len(p.data)) || !p.fill.Contains(idx) {
		return 0, false
	}

	return int64(p.data[idx]), true
}

// Uint64 retrieves an uint64 value at a specified index
func (p *columnUint32) Uint64(idx uint32) (uint64, bool) {
	if idx >= uint32(len(p.data)) || !p.fill.Contains(idx) {
		return 0, false
	}

	return uint64(p.data[idx]), true
}

// Del removes a value at a specified index
func (p *columnUint32) Del(idx uint32) {
	p.fill.Remove(idx)
	p.data[idx] = 0
}

// Bitmap returns the associated index bitmap.
func (p *columnUint32) Bitmap() bitmap.Bitmap {
	return p.fill
}

// Contains checks whether the column has a value at a specified index.
func (p *columnUint32) Contains(idx uint32) bool {
	return p.fill.Contains(idx)
}

// --------------------------- uint64s ----------------------------

// columnUint64 represents a generic column
type columnUint64 struct {
	fill bitmap.Bitmap // The fill-list
	data []uint64      // The actual values
}

// makeUint64s creates a new vector or uint64s
func makeUint64s() Column {
	return &columnUint64{
		fill: make(bitmap.Bitmap, 0, 4),
		data: make([]uint64, 0, 64),
	}
}

// Set sets a value at a specified index
func (p *columnUint64) Set(idx uint32, value interface{}) {
	size := uint32(len(p.data))
	for i := size; i <= idx; i++ {
		p.data = append(p.data, 0)
	}

	// Set the data at index
	p.fill.Set(idx)
	p.data[idx] = value.(uint64)
}

// Value retrieves a value at a specified index
func (p *columnUint64) Value(idx uint32) (interface{}, bool) {
	if idx >= uint32(len(p.data)) || !p.fill.Contains(idx) {
		return uint64(0), false
	}

	return p.data[idx], true
}

// Float64 retrieves a float64 value at a specified index
func (p *columnUint64) Float64(idx uint32) (float64, bool) {
	if idx >= uint32(len(p.data)) || !p.fill.Contains(idx) {
		return 0, false
	}

	return float64(p.data[idx]), true
}

// Int64 retrieves an int64 value at a specified index
func (p *columnUint64) Int64(idx uint32) (int64, bool) {
	if idx >= uint32(len(p.data)) || !p.fill.Contains(idx) {
		return 0, false
	}

	return int64(p.data[idx]), true
}

// Uint64 retrieves an uint64 value at a specified index
func (p *columnUint64) Uint64(idx uint32) (uint64, bool) {
	if idx >= uint32(len(p.data)) || !p.fill.Contains(idx) {
		return 0, false
	}

	return uint64(p.data[idx]), true
}

// Del removes a value at a specified index
func (p *columnUint64) Del(idx uint32) {
	p.fill.Remove(idx)
	p.data[idx] = 0
}

// Bitmap returns the associated index bitmap.
func (p *columnUint64) Bitmap() bitmap.Bitmap {
	return p.fill
}

// Contains checks whether the column has a value at a specified index.
func (p *columnUint64) Contains(idx uint32) bool {
	return p.fill.Contains(idx)
}