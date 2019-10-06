package gopmdk

// #cgo LDFLAGS: -lpmem
// #include <libpmem.h>
// #include <stdlib.h>
import "C"

import (
	"reflect"
	"unsafe"
)

// void *pmem_map_file(const char *path, size_t len, int flags, mode_t mode, size_t *mapped_lenp, int *is_pmemp);
func MapFile(path string, len int, flags int, mode int) (addr unsafe.Pointer, mappedLen int, isPmem bool) {
	pathc := C.CString(path)
	mappedLenc := C.size_t(mappedLen)
	isPmemc := C.int(0)
	addr = C.pmem_map_file(pathc, C.size_t(len), C.int(flags), C.mode_t(mode), &mappedLenc, &isPmemc)
	C.free(unsafe.Pointer(pathc))
	isPmem = isPmemc != 0
	mappedLen = int(mappedLenc)
	return
}

// int pmem_unmap(void *addr, size_t len);
func Unmap(addr interface{}, size int) {
	C.pmem_unmap(unsafe.Pointer(reflect.ValueOf(addr).Elem().UnsafeAddr()), C.size_t(size))
}

// int pmem_is_pmem(const void *addr, size_t len);
func IsPmem(addr interface{}, size int) {
	C.pmem_is_pmem(unsafe.Pointer(reflect.ValueOf(addr).Elem().UnsafeAddr()), C.size_t(size))
}

// void pmem_persist(const void *addr, size_t len);
func Persist(addr interface{}, size int) {
	ptr, ok := addr.(unsafe.Pointer)
	if !ok {
		ptr = unsafe.Pointer(reflect.ValueOf(addr).Elem().UnsafeAddr())
	}
	C.pmem_persist(ptr, C.size_t(size))
}

// int pmem_msync(const void *addr, size_t len);
func Msync(addr interface{}, size int) {
	C.pmem_msync(unsafe.Pointer(reflect.ValueOf(addr).Elem().UnsafeAddr()), C.size_t(size))
}

// int pmem_has_auto_flush(void);
func HasAutoFlush() int {
	return int(C.pmem_has_auto_flush())
}

// void pmem_flush(const void *addr, size_t len);
func Flush(addr interface{}, len int) {
	C.pmem_flush(unsafe.Pointer(reflect.ValueOf(addr).Elem().UnsafeAddr()), C.size_t(len))
}

// void pmem_deep_flush(const void *addr, size_t len);
func DeepFlush(addr interface{}, size int) {
	C.pmem_deep_flush(unsafe.Pointer(reflect.ValueOf(addr).Elem().UnsafeAddr()), C.size_t(size))
}

// int pmem_deep_drain(const void *addr, size_t len);
func DeepDrain(addr interface{}, size int) int {
	return int(C.pmem_deep_drain(unsafe.Pointer(reflect.ValueOf(addr).Elem().UnsafeAddr()), C.size_t(size)))
}

// int pmem_deep_persist(const void *addr, size_t len);
func DeepPersist(addr interface{}, size int) int {
	return int(C.pmem_deep_persist(unsafe.Pointer(reflect.ValueOf(addr).Elem().UnsafeAddr()), C.size_t(size)))
}

// void pmem_drain(void);
func Drain() { C.pmem_drain() }

// int pmem_has_hw_drain(void);
func HasHwDrain() int { return int(C.pmem_has_hw_drain()) }

// void *pmem_memmove_persist(void *pmemdest, const void *src, size_t len);
func MemmovePersist(dst, src interface{}, size int) {
	dstv := unsafe.Pointer(reflect.ValueOf(dst).Elem().UnsafeAddr())
	srcv := unsafe.Pointer(reflect.ValueOf(src).Elem().UnsafeAddr())
	C.pmem_memmove_persist(dstv, srcv, C.size_t(size))
}

// void *pmem_memcpy_persist(void *pmemdest, const void *src, size_t len);
func MemcpyPersist(dst, src interface{}, size int) {
	dstv := unsafe.Pointer(reflect.ValueOf(dst).Elem().UnsafeAddr())
	srcv := unsafe.Pointer(reflect.ValueOf(src).Elem().UnsafeAddr())
	C.pmem_memcpy_persist(dstv, srcv, C.size_t(size))
}

// void *pmem_memset_persist(void *pmemdest, int c, size_t len);
func MemsetPersist(dst interface{}, c int, size int) {
	dstv := unsafe.Pointer(reflect.ValueOf(dst).Elem().UnsafeAddr())
	C.pmem_memset_persist(dstv, C.int(c), C.size_t(size))
}

// void *pmem_memmove_nodrain(void *pmemdest, const void *src, size_t len);
func MemmoveNodrain(dst, src interface{}, len int) {
	dstv := unsafe.Pointer(reflect.ValueOf(dst).Elem().UnsafeAddr())
	srcv := unsafe.Pointer(reflect.ValueOf(src).Elem().UnsafeAddr())
	C.pmem_memmove_nodrain(dstv, srcv, C.size_t(len))
}

// void *pmem_memcpy_nodrain(void *pmemdest, const void *src, size_t len);
func MemcpyNodrain(dst, src interface{}, len int) {
	dstv := unsafe.Pointer(reflect.ValueOf(dst).Elem().UnsafeAddr())
	srcv := unsafe.Pointer(reflect.ValueOf(src).Elem().UnsafeAddr())
	C.pmem_memcpy_nodrain(dstv, srcv, C.size_t(len))
}

// void *pmem_memset_nodrain(void *pmemdest, int c, size_t len);
func MemsetNodrain(dst interface{}, c int, len int) {
	dstv := unsafe.Pointer(reflect.ValueOf(dst).Elem().UnsafeAddr())
	C.pmem_memset_persist(dstv, C.int(c), C.size_t(len))
}

// void *pmem_memmove(void *pmemdest, const void *src, size_t len, unsigned flags);
func Memmove(dst, src interface{}, len int, flags uint) {
	dstv := unsafe.Pointer(reflect.ValueOf(dst).Elem().UnsafeAddr())
	srcv := unsafe.Pointer(reflect.ValueOf(src).Elem().UnsafeAddr())
	C.pmem_memmove(dstv, srcv, C.size_t(len), C.unsigned(flags))
}

// void *pmem_memcpy(void *pmemdest, const void *src, size_t len, unsigned flags);
func Memcpy(dst, src interface{}, len int, flags uint) {
	dstv := unsafe.Pointer(reflect.ValueOf(dst).Elem().UnsafeAddr())
	srcv := unsafe.Pointer(reflect.ValueOf(src).Elem().UnsafeAddr())
	C.pmem_memcpy(dstv, srcv, C.size_t(len), C.unsigned(flags))
}

// void *pmem_memset(void *pmemdest, int c, size_t len, unsigned flags);
func Memset(dst interface{}, c int, len int, flags uint) {
	dstv := unsafe.Pointer(reflect.ValueOf(dst).Elem().UnsafeAddr())
	C.pmem_memset(dstv, C.int(c), C.size_t(len), C.unsigned(flags))
}

// const char *pmem_check_version(unsigned major_required, unsigned minor_required);
func CheckVersion(major, minor uint) string {
	return C.GoString(C.pmem_check_version(C.unsigned(major), C.unsigned(minor)))
}

// const char *pmem_errormsg(void);
func Errormsg() string { return C.GoString(C.pmem_errormsg()) }
