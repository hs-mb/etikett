//go:build js && wasm

package label

import (
	"context"

	"github.com/coder/websocket"
	"honnef.co/go/js/dom/v2"
)

func SendPrintServer(source string) error {
	addr := dom.GetWindow().Document().QuerySelector("meta[name='print-addr']").(*dom.HTMLMetaElement).Content()
	ctx := context.Background()
	c, _, err := websocket.Dial(ctx, addr, nil)
	if err != nil {
		return err
	}
	err = c.Write(ctx, websocket.MessageBinary, []byte(source))
	if err != nil {
		return err
	}
	return c.Close(websocket.StatusNormalClosure, "")
}
