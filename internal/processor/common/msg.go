package common

import (
	"fmt"
	"github.com/LampardNguyen234/whale-alert/internal/clients/common"
	"github.com/iancoleman/strcase"
	"strings"
)

type TxMsg struct {
	From      string
	To        string
	Amount    string
	Token     string
	TokenName string
	TxHash    string
}

func FormatTxURL(txHash string) string {
	return common.Explorer.FormatTxURL(txHash)
}

type MsgFormatter struct {
	title string
	resp  string
}

func (f *MsgFormatter) FormatTitle(title string) *MsgFormatter {
	f.title = title
	return f
}

func (f *MsgFormatter) FormatKeyValueMsg(key string, value interface{}) *MsgFormatter {
	f.resp = f.resp + fmt.Sprintf("%v: %v\n", strcase.ToCamel(key), value)
	return f
}

func (f *MsgFormatter) FormatMsg(msg string) *MsgFormatter {
	f.resp = f.resp + fmt.Sprintf("%v\n", msg)
	return f
}

func (f *MsgFormatter) String() string {
	resp := fmt.Sprintf("========== [%v] ==========\n", strings.ToUpper(f.title))
	resp = resp + f.resp
	return resp
}
