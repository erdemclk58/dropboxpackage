package nodes

import (
	"bitbucket.org/mosteknoloji/robomotion-go-lib/message"
	"bitbucket.org/mosteknoloji/robomotion-go-lib/runtime"
	"github.com/tj/go-dropbox"
	"github.com/tj/go-dropy"
)

type Delete struct {
	runtime.Node `spec:"id=Robomotion.Dropbox.Delete,name=Delete,icon=mdiDropbox,color=#007ee5"`

	// Input1
	InClientID runtime.InVariable `spec:"title=Client ID,scope=Message,name=client_id,messageOnly"`

	// Input2
	InDropboxPath runtime.InVariable `spec:"title=Dropbox Path,scope=Custom"`
}

func (n *Delete) OnCreate() error {
	return nil
}

func (n *Delete) OnMessage(ctx message.Context) (err error) {

	client_id, err := n.InClientID.GetString(ctx)
	if err != nil {
		return err
	}
	token := getToken(client_id)
	client := dropy.New(dropbox.New(dropbox.NewConfig(token)))

	path, err := n.InDropboxPath.GetString(ctx)
	if err != nil {
		return err
	}
	client.Delete(path)
	return nil
}

func (n *Delete) OnClose() error {
	return nil
}
