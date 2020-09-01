package nodes

import (
	"bitbucket.org/mosteknoloji/robomotion-go-lib/message"
	"bitbucket.org/mosteknoloji/robomotion-go-lib/runtime"
	"github.com/tj/go-dropbox"
	"github.com/tj/go-dropy"
)

type Move struct {
	runtime.Node `spec:"id=Robomotion.Dropbox.Move,name=Move,icon=mdiDropbox,color=#007ee5"`

	// Input1
	InClientID runtime.InVariable `spec:"title=Client ID,scope=Message,name=client_id,messageOnly"`

	// Input2
	InSourcePath runtime.InVariable `spec:"title=Source Path,scope=Custom"`

	// Input3
	InDestinationPath runtime.InVariable `spec:"title=Destination Path,scope=Custom"`
}

func (n *Move) OnCreate() error {
	return nil
}

func (n *Move) OnMessage(ctx message.Context) (err error) {
	client_id, err := n.InClientID.GetString(ctx)
	if err != nil {
		return err
	}
	token := getToken(client_id)
	client := dropy.New(dropbox.New(dropbox.NewConfig(token)))

	path, err := n.InSourcePath.GetString(ctx)
	if err != nil {
		return err
	}

	path1, err := n.InDestinationPath.GetString(ctx)
	if err != nil {
		return err
	}
	client.Move(path, path1)

	return nil
}

func (n *Move) OnClose() error {
	return nil
}
