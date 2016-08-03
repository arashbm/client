// Auto-generated by avdl-compiler v1.3.4 (https://github.com/keybase/node-avdl-compiler)
//   Input file: avdl/passphrase_common.avdl

package keybase1

import (
	rpc "github.com/keybase/go-framed-msgpack-rpc"
)

type Feature struct {
	Allow        bool   `codec:"allow" json:"allow"`
	DefaultValue bool   `codec:"defaultValue" json:"defaultValue"`
	Readonly     bool   `codec:"readonly" json:"readonly"`
	Label        string `codec:"label" json:"label"`
}

type GUIEntryFeatures struct {
	ShowTyping Feature `codec:"showTyping" json:"showTyping"`
}

type PassphraseType int

const (
	PassphraseType_NONE               PassphraseType = 0
	PassphraseType_PAPER_KEY          PassphraseType = 1
	PassphraseType_PASS_PHRASE        PassphraseType = 2
	PassphraseType_VERIFY_PASS_PHRASE PassphraseType = 3
)

var PassphraseTypeMap = map[string]PassphraseType{
	"NONE":               0,
	"PAPER_KEY":          1,
	"PASS_PHRASE":        2,
	"VERIFY_PASS_PHRASE": 3,
}

type GUIEntryArg struct {
	WindowTitle string           `codec:"windowTitle" json:"windowTitle"`
	Prompt      string           `codec:"prompt" json:"prompt"`
	Username    string           `codec:"username" json:"username"`
	SubmitLabel string           `codec:"submitLabel" json:"submitLabel"`
	CancelLabel string           `codec:"cancelLabel" json:"cancelLabel"`
	RetryLabel  string           `codec:"retryLabel" json:"retryLabel"`
	Type        PassphraseType   `codec:"type" json:"type"`
	Features    GUIEntryFeatures `codec:"features" json:"features"`
}

type GetPassphraseRes struct {
	Passphrase  string `codec:"passphrase" json:"passphrase"`
	StoreSecret bool   `codec:"storeSecret" json:"storeSecret"`
}

type PassphraseCommonInterface interface {
}

func PassphraseCommonProtocol(i PassphraseCommonInterface) rpc.Protocol {
	return rpc.Protocol{
		Name:    "keybase.1.passphraseCommon",
		Methods: map[string]rpc.ServeHandlerDescription{},
	}
}

type PassphraseCommonClient struct {
	Cli rpc.GenericClient
}
