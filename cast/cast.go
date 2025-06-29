package cast

import "time"

const errorMsg = "unable to cast %#v of type %T to %T"
const errorMsgWith = "unable to cast %#v of type %T to %T: %w"

// Basic is a type parameter constraint for functions accepting basic types.
//
// It represents the supported basic types this package can cast to.
type Basic interface {
	string | bool | Number | time.Time | time.Duration
}

// ToE casts any value to a [Basic] type.
func ToE[T Basic](i any) (T, error) {
	var t T

	var v any
	var err error

	switch any(t).(type) {
	case string:
		v, err = ToStringE(i)
	case bool:
		v, err = ToBoolE(i)
	case int:
		v, err = toNumberE[int](i, parseInt[int])
	case int8:
		v, err = toNumberE[int8](i, parseInt[int8])
	case int16:
		v, err = toNumberE[int16](i, parseInt[int16])
	case int32:
		v, err = toNumberE[int32](i, parseInt[int32])
	case int64:
		v, err = toNumberE[int64](i, parseInt[int64])
	case uint:
		v, err = toUnsignedNumberE[uint](i, parseUint[uint])
	case uint8:
		v, err = toUnsignedNumberE[uint8](i, parseUint[uint8])
	case uint16:
		v, err = toUnsignedNumberE[uint16](i, parseUint[uint16])
	case uint32:
		v, err = toUnsignedNumberE[uint32](i, parseUint[uint32])
	case uint64:
		v, err = toUnsignedNumberE[uint64](i, parseUint[uint64])
	case float32:
		v, err = toNumberE[float32](i, parseFloat[float32])
	case float64:
		v, err = toNumberE[float64](i, parseFloat[float64])
	case time.Time:
		v, err = ToTimeE(i)
	case time.Duration:
		v, err = ToDurationE(i)
	}

	if err != nil {
		return t, err
	}

	return v.(T), nil
}

// Must is a helper that wraps a call to a cast function and panics if the error is non-nil.
func Must[T any](i any, err error) T {
	if err != nil {
		panic(err)
	}

	return i.(T)
}

// To casts any value to a [Basic] type.
func To[T Basic](i any) T {
	v, _ := ToE[T](i)

	return v
}

// ToBool casts any value to a(n) bool type.
func ToBool(i any) bool {
	v, _ := ToBoolE(i)
	return v
}

// ToString casts any value to a(n) string type.
func ToString(i any) string {
	v, _ := ToStringE(i)
	return v
}

// ToTime casts any value to a(n) time.Time type.
func ToTime(i any) time.Time {
	v, _ := ToTimeE(i)
	return v
}

// ToTimeInDefaultLocation casts any value to a(n) time.Time type.
func ToTimeInDefaultLocation(i any, location *time.Location) time.Time {
	v, _ := ToTimeInDefaultLocationE(i, location)
	return v
}

// ToDuration casts any value to a(n) time.Duration type.
func ToDuration(i any) time.Duration {
	v, _ := ToDurationE(i)
	return v
}

// ToInt casts any value to a(n) int type.
func ToInt(i any) int {
	v, _ := ToIntE(i)
	return v
}

// ToInt8 casts any value to a(n) int8 type.
func ToInt8(i any) int8 {
	v, _ := ToInt8E(i)
	return v
}

// ToInt16 casts any value to a(n) int16 type.
func ToInt16(i any) int16 {
	v, _ := ToInt16E(i)
	return v
}

// ToInt32 casts any value to a(n) int32 type.
func ToInt32(i any) int32 {
	v, _ := ToInt32E(i)
	return v
}

// ToInt64 casts any value to a(n) int64 type.
func ToInt64(i any) int64 {
	v, _ := ToInt64E(i)
	return v
}

// ToUint casts any value to a(n) uint type.
func ToUint(i any) uint {
	v, _ := ToUintE(i)
	return v
}

// ToUint8 casts any value to a(n) uint8 type.
func ToUint8(i any) uint8 {
	v, _ := ToUint8E(i)
	return v
}

// ToUint16 casts any value to a(n) uint16 type.
func ToUint16(i any) uint16 {
	v, _ := ToUint16E(i)
	return v
}

// ToUint32 casts any value to a(n) uint32 type.
func ToUint32(i any) uint32 {
	v, _ := ToUint32E(i)
	return v
}

// ToUint64 casts any value to a(n) uint64 type.
func ToUint64(i any) uint64 {
	v, _ := ToUint64E(i)
	return v
}

// ToFloat32 casts any value to a(n) float32 type.
func ToFloat32(i any) float32 {
	v, _ := ToFloat32E(i)
	return v
}

// ToFloat64 casts any value to a(n) float64 type.
func ToFloat64(i any) float64 {
	v, _ := ToFloat64E(i)
	return v
}

// ToStringMapString casts any value to a(n) map[string]string type.
func ToStringMapString(i any) map[string]string {
	v, _ := ToStringMapStringE(i)
	return v
}

// ToStringMapStringSlice casts any value to a(n) map[string][]string type.
func ToStringMapStringSlice(i any) map[string][]string {
	v, _ := ToStringMapStringSliceE(i)
	return v
}

// ToStringMapBool casts any value to a(n) map[string]bool type.
func ToStringMapBool(i any) map[string]bool {
	v, _ := ToStringMapBoolE(i)
	return v
}

// ToStringMapInt casts any value to a(n) map[string]int type.
func ToStringMapInt(i any) map[string]int {
	v, _ := ToStringMapIntE(i)
	return v
}

// ToStringMapInt64 casts any value to a(n) map[string]int64 type.
func ToStringMapInt64(i any) map[string]int64 {
	v, _ := ToStringMapInt64E(i)
	return v
}

// ToStringMap casts any value to a(n) map[string]any type.
func ToStringMap(i any) map[string]any {
	v, _ := ToStringMapE(i)
	return v
}

// ToSlice casts any value to a(n) []any type.
func ToSlice(i any) []any {
	v, _ := ToSliceE(i)
	return v
}

// ToBoolSlice casts any value to a(n) []bool type.
func ToBoolSlice(i any) []bool {
	v, _ := ToBoolSliceE(i)
	return v
}

// ToStringSlice casts any value to a(n) []string type.
func ToStringSlice(i any) []string {
	v, _ := ToStringSliceE(i)
	return v
}

// ToIntSlice casts any value to a(n) []int type.
func ToIntSlice(i any) []int {
	v, _ := ToIntSliceE(i)
	return v
}

// ToInt64Slice casts any value to a(n) []int64 type.
func ToInt64Slice(i any) []int64 {
	v, _ := ToInt64SliceE(i)
	return v
}

// ToUintSlice casts any value to a(n) []uint type.
func ToUintSlice(i any) []uint {
	v, _ := ToUintSliceE(i)
	return v
}

// ToFloat64Slice casts any value to a(n) []float64 type.
func ToFloat64Slice(i any) []float64 {
	v, _ := ToFloat64SliceE(i)
	return v
}

// ToDurationSlice casts any value to a(n) []time.Duration type.
func ToDurationSlice(i any) []time.Duration {
	v, _ := ToDurationSliceE(i)
	return v
}

// ToBoolSliceE casts any value to a(n) []bool type.
func ToBoolSliceE(i any) ([]bool, error) {
	return toSliceE[bool](i)
}

// ToDurationSliceE casts any value to a(n) []time.Duration type.
func ToDurationSliceE(i any) ([]time.Duration, error) {
	return toSliceE[time.Duration](i)
}

// ToIntSliceE casts any value to a(n) []int type.
func ToIntSliceE(i any) ([]int, error) {
	return toSliceE[int](i)
}

// ToInt8SliceE casts any value to a(n) []int8 type.
func ToInt8SliceE(i any) ([]int8, error) {
	return toSliceE[int8](i)
}

// ToInt16SliceE casts any value to a(n) []int16 type.
func ToInt16SliceE(i any) ([]int16, error) {
	return toSliceE[int16](i)
}

// ToInt32SliceE casts any value to a(n) []int32 type.
func ToInt32SliceE(i any) ([]int32, error) {
	return toSliceE[int32](i)
}

// ToInt64SliceE casts any value to a(n) []int64 type.
func ToInt64SliceE(i any) ([]int64, error) {
	return toSliceE[int64](i)
}

// ToUintSliceE casts any value to a(n) []uint type.
func ToUintSliceE(i any) ([]uint, error) {
	return toSliceE[uint](i)
}

// ToUint8SliceE casts any value to a(n) []uint8 type.
func ToUint8SliceE(i any) ([]uint8, error) {
	return toSliceE[uint8](i)
}

// ToUint16SliceE casts any value to a(n) []uint16 type.
func ToUint16SliceE(i any) ([]uint16, error) {
	return toSliceE[uint16](i)
}

// ToUint32SliceE casts any value to a(n) []uint32 type.
func ToUint32SliceE(i any) ([]uint32, error) {
	return toSliceE[uint32](i)
}

// ToUint64SliceE casts any value to a(n) []uint64 type.
func ToUint64SliceE(i any) ([]uint64, error) {
	return toSliceE[uint64](i)
}

// ToFloat32SliceE casts any value to a(n) []float32 type.
func ToFloat32SliceE(i any) ([]float32, error) {
	return toSliceE[float32](i)
}

// ToFloat64SliceE casts any value to a(n) []float64 type.
func ToFloat64SliceE(i any) ([]float64, error) {
	return toSliceE[float64](i)
}
