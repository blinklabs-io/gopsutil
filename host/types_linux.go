// SPDX-License-Identifier: BSD-3-Clause
//go:build ignore

/*
Input to cgo -godefs.
*/

package host

/*
#include <sys/types.h>
#include <utmp.h>

enum {
       sizeofPtr = sizeof(void*),
};

*/
import "C"

// Machine characteristics; for internal use.

const (
	sizeofPtr      = C.sizeofPtr
	sizeofShort    = C.sizeof_short
	sizeofInt      = C.sizeof_int
	sizeofLong     = C.sizeof_long
	sizeofLongLong = C.sizeof_longlong
	sizeOfUtmp     = C.sizeof_struct_utmp
)

// Basic types

type (
	_C_short     C.short
	_C_int       C.int
	_C_long      C.long
	_C_long_long C.longlong
)

type (
	utmp        C.struct_utmp
	exit_status C.struct_exit_status
	timeval     C.struct_timeval
)
