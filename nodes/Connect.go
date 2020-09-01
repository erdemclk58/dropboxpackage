package nodes

import (
	"bitbucket.org/mosteknoloji/robomotion-go-lib/message"
	"bitbucket.org/mosteknoloji/robomotion-go-lib/runtime"
)

type Connect struct {
	runtime.Node `spec:"id=Robomotion.Dropbox.Connect,name=Connect,icon=mdiDropbox,color=#007ee5"`

	// Output
	OutClientID runtime.OutVariable `spec:"title=Client ID,scope=Message,name=client_id,messageOnly"`

	// Option
	OptToken runtime.Credential `spec:"title=Dropbox Token,category=4"`
}

func (n *Connect) OnCreate() error {
	return nil
}

func (n *Connect) OnMessage(ctx message.Context) (err error) {

	item, err := n.OptToken.Get()
	if err != nil {

		return err
	}

	token := item["value"].(string)

	clientID := addToken(token)
	//clients[clientID] = token
	n.OutClientID.Set(ctx, clientID)

	return nil
}

func (n *Connect) OnClose() error {
	return nil
}

/*func setClipboard(text string) (err error) {
	clipboard.WriteAll(fmt.Sprintf("%v", text))
	return nil
}*/
