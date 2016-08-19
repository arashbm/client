// Copyright 2015 Keybase, Inc. All rights reserved. Use of
// this source code is governed by the included BSD license.

package libkb

import (
	"errors"
	"fmt"
	"strings"
	"sync"

	"github.com/PuerkitoBio/goquery"
	keybase1 "github.com/keybase/client/go/protocol"
	jsonw "github.com/keybase/go-jsonw"
)

func pvlServiceToString(service keybase1.ProofType) (string, ProofError) {
	for name, stat := range keybase1.ProofTypeMap {
		if service == stat {
			return strings.ToLower(name), nil
		}
	}

	return "", NewProofError(keybase1.ProofStatus_INVALID_PVL,
		"Unsupported service %v", service)
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

func pvlJSONStringOrMarshal(object *jsonw.Wrapper) (string, error) {
	s, err := object.GetString()
	if err == nil {
		return s, nil
	}
	b, err := object.Marshal()
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// Get the HTML contents of all elements in a selection, concatenated by a space.
func pvlSelectionContents(selection *goquery.Selection, useAttr bool, attr string) (string, error) {
	len := selection.Length()
	results := make([]string, len)
	errs := make([]error, len)
	var wg sync.WaitGroup
	wg.Add(len)
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
		wg.Done()
	})
	wg.Wait()
	for _, err := range errs {
		if err != nil {
			return "", err
		}
	}
	return strings.Join(results, " "), nil
}
