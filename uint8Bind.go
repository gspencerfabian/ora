// Copyright 2014 Rana Ian. All rights reserved.
// Use of this source code is governed by The MIT License
// found in the accompanying LICENSE file.

package ora

/*
#include <oci.h>
#include <stdlib.h>
#include <string.h>
*/
import "C"
import (
	"unsafe"
)

type uint8Bind struct {
	environment *Environment
	ocibnd      *C.OCIBind
	ociNumber   C.OCINumber
}

func (uint8Bind *uint8Bind) bind(value uint8, position int, ocistmt *C.OCIStmt) error {
	r := C.OCINumberFromInt(
		uint8Bind.environment.ocierr, //OCIError            *err,
		unsafe.Pointer(&value),       //const void          *inum,
		1, //uword               inum_length,
		C.OCI_NUMBER_UNSIGNED, //uword               inum_s_flag,
		&uint8Bind.ociNumber)  //OCINumber           *number );
	if r == C.OCI_ERROR {
		return uint8Bind.environment.ociError()
	}
	r = C.OCIBindByPos2(
		ocistmt, //OCIStmt      *stmtp,
		(**C.OCIBind)(&uint8Bind.ocibnd),     //OCIBind      **bindpp,
		uint8Bind.environment.ocierr,         //OCIError     *errhp,
		C.ub4(position),                      //ub4          position,
		unsafe.Pointer(&uint8Bind.ociNumber), //void         *valuep,
		C.sb8(C.sizeof_OCINumber),            //sb8          value_sz,
		C.SQLT_VNU,                           //ub2          dty,
		nil,                                  //void         *indp,
		nil,                                  //ub2          *alenp,
		nil,                                  //ub2          *rcodep,
		0,                                    //ub4          maxarr_len,
		nil,                                  //ub4          *curelep,
		C.OCI_DEFAULT)                        //ub4          mode );
	if r == C.OCI_ERROR {
		return uint8Bind.environment.ociError()
	}
	return nil
}

func (uint8Bind *uint8Bind) setPtr() error {
	return nil
}

func (uint8Bind *uint8Bind) close() {
	defer func() {
		recover()
	}()
	uint8Bind.ocibnd = nil
	uint8Bind.environment.uint8BindPool.Put(uint8Bind)
}
