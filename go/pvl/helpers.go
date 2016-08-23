// Copyright 2015 Keybase, Inc. All rights reserved. Use of
// this source code is governed by the included BSD license.

package pvl

import (
	b64 "encoding/base64"
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	libkb "github.com/keybase/client/go/libkb"
	keybase1 "github.com/keybase/client/go/protocol"
	jsonw "github.com/keybase/go-jsonw"
)

// Substitute vars for %{name} in the string.
// Only substitutes whitelisted variables.
// It is an error to refer to an unknown variable or undefined numbered group.
// Match is an optional slice which is a regex match.
func pvlSubstitute(template string, state PvlScriptState, match []string) (string, libkb.ProofError) {
	vars := state.Vars
	webish := (state.Service == keybase1.ProofType_DNS || state.Service == keybase1.ProofType_GENERIC_WEB_SITE)

	var outerr libkb.ProofError
	// Regex to find %{name} occurrences.
	re := regexp.MustCompile("%{[\\w]+}")
	pvlSubstituteOne := func(vartag string) string {
		// Strip off the %, {, and }
		varname := vartag[2 : len(vartag)-1]
		var value string
		switch varname {
		case "username_service":
			if !webish {
				value = vars.UsernameService
			} else {
				outerr = libkb.NewProofError(keybase1.ProofStatus_INVALID_PVL,
					"Cannot use username_service in proof type %v", state.Service)
			}
		case "username_keybase":
			value = vars.UsernameKeybase
		case "sig":
			value = b64.StdEncoding.EncodeToString(vars.Sig)
		case "sig_id_medium":
			value = vars.SigIDMedium
		case "sig_id_short":
			value = vars.SigIDShort
		case "hostname":
			if webish {
				value = vars.Hostname
			} else {
				outerr = libkb.NewProofError(keybase1.ProofStatus_INVALID_PVL,
					"Cannot use username_service in proof type %v", state.Service)
			}
		default:
			var i int
			i, err := strconv.Atoi(varname)
			if err == nil {
				if i >= 0 && i < len(match) {
					value = match[i]
				} else {
					outerr = libkb.NewProofError(keybase1.ProofStatus_BAD_API_URL,
						"Substitution argument %v out of range of match", i)
				}
			} else {
				outerr = libkb.NewProofError(keybase1.ProofStatus_INVALID_PVL,
					"Unrecognized variable: %v", varname)
			}
		}
		return regexp.QuoteMeta(value)
	}
	res := re.ReplaceAllStringFunc(template, pvlSubstituteOne)
	if outerr != nil {
		return template, outerr
	}
	return res, nil
}

func pvlServiceToString(service keybase1.ProofType) (string, libkb.ProofError) {
	for name, stat := range keybase1.ProofTypeMap {
		if service == stat {
			return strings.ToLower(name), nil
		}
	}

	return "", libkb.NewProofError(keybase1.ProofStatus_INVALID_PVL, "Unsupported service %v", service)
}

func pvlJSONHasKey(w *jsonw.Wrapper, key string) bool {
	return !w.AtKey(key).IsNil()
}

// Return the elements of an array.
func pvlJSONUnpackArray(w *jsonw.Wrapper) ([]*jsonw.Wrapper, error) {
	w, err := w.ToArray()
	if err != nil {
		return nil, err
	}
	length, err := w.Len()
	if err != nil {
		return nil, err
	}
	res := make([]*jsonw.Wrapper, length)
	for i := 0; i < length; i++ {
		res[i] = w.AtIndex(i)
	}
	return res, nil
}

// Return the elements of an array or values of a map.
func pvlJSONGetChildren(w *jsonw.Wrapper) ([]*jsonw.Wrapper, error) {
	dict, err := w.ToDictionary()
	isDict := err == nil
	array, err := w.ToArray()
	isArray := err == nil

	switch {
	case isDict:
		keys, err := dict.Keys()
		if err != nil {
			return nil, err
		}
		var res = make([]*jsonw.Wrapper, len(keys))
		for i, key := range keys {
			res[i] = dict.AtKey(key)
		}
		return res, nil
	case isArray:
		return pvlJSONUnpackArray(array)
	default:
		return nil, errors.New("got children of non-container")
	}
}

// pvlJSONStringSimple converts a simple json object into a string.
// Simple objects are those that are not arrays or objects.
// Non-simple objects result in an error.
func pvlJSONStringSimple(object *jsonw.Wrapper) (string, error) {
	x, err := object.GetString()
	if err == nil {
		return x, nil
	}
	y, err := object.GetInt()
	if err == nil {
		return string(y), nil
	}
	z, err := object.GetBool()
	if err == nil {
		if z {
			return "true", nil
		}
		return "false", nil
	}
	isnil := object.IsNil()
	if isnil {
		return "null", nil
	}

	return "", fmt.Errorf("Non-simple object: %v", object)
}

// pvlSelectionContents gets the HTML contents of all elements in a selection, concatenated by a space.
// If getting the contents/attr value of any elements fails, that does not cause an error.
// The result can be an empty string.
func pvlSelectionContents(selection *goquery.Selection, useAttr bool, attr string) string {
	length := selection.Length()
	results := make([]string, length)
	errs := make([]error, length)
	selection.Each(func(i int, element *goquery.Selection) {
		if useAttr {
			res, ok := element.Attr(attr)
			results[i] = res
			if !ok {
				errs[i] = fmt.Errorf("Could not get attr %v of element", attr)
			}
		} else {
			results[i] = element.Text()
			errs[i] = nil
		}
	})

	var goodresults []string
	for i, err := range errs {
		if err == nil {
			goodresults = append(goodresults, results[i])
		}
	}
	return strings.Join(goodresults, " ")
}
