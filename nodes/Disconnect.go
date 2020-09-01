package nodes

import (
	"bitbucket.org/mosteknoloji/robomotion-go-lib/message"
	"bitbucket.org/mosteknoloji/robomotion-go-lib/runtime"
)

type Disconnect struct {
	runtime.Node `spec:"id=Robomotion.Dropbox.Disconnect,name=Disconnect,icon=mdiDropbox,color=#007ee5"`

	//Input
	InClientID runtime.InVariable `spec:"title=Client ID,scope=Message,name=client_id,messageOnly"`
}

func (n *Disconnect) OnCreate() error {
	return nil
}

func (n *Disconnect) OnMessage(ctx message.Context) (err error) {
	client_id, err := n.InClientID.GetString(ctx)
	if err != nil {
		return err
	}
	removeToken(client_id)
	return nil
}

func (n *Disconnect) OnClose() error {
	return nil
}
