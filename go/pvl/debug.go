// Copyright 2016 Keybase, Inc. All rights reserved. Use of
// this source code is governed by the included BSD license.

package pvl

import (
	"fmt"

	libkb "github.com/keybase/client/go/libkb"
	keybase1 "github.com/keybase/client/go/protocol"
)

func debugWithState(g libkb.ProofContext, state PvlScriptState, format string, arg ...interface{}) {
	s := fmt.Sprintf(format, arg...)
	g.GetLog().Debug("PVL @(service:%v script:%v pc:%v) %v",
		debugServiceToString(state.Service), state.WhichScript, state.PC, s)
}

func debugWithStateError(g libkb.ProofContext, state PvlScriptState, err libkb.ProofError) {
	g.GetLog().Debug("PVL @(service:%v script:%v pc:%v) Error code=%v: %v",
		debugServiceToString(state.Service), state.WhichScript, state.PC, err.GetProofStatus(), err.GetDesc())
}

func debugWithPosition(g libkb.ProofContext, service keybase1.ProofType, whichscript int, pc int, format string, arg ...interface{}) {
	s := fmt.Sprintf(format, arg...)
	g.GetLog().Debug("PVL @(service:%v script:%v pc:%v) %v",
		debugServiceToString(service), whichscript, pc, s)
}

func debug(g libkb.ProofContext, format string, arg ...interface{}) {
	s := fmt.Sprintf(format, arg...)
	g.GetLog().Debug("PVL %v", s)
}

// debugServiceToString returns the name of a service or number string if it is invalid.
func debugServiceToString(service keybase1.ProofType) string {
	s, err := pvlServiceToString(service)
	if err != nil {
		return string(service)
	}
	return s
}
