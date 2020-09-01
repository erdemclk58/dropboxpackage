package nodes

import (
	"strings"

	"bitbucket.org/mosteknoloji/robomotion-go-lib/message"
	"bitbucket.org/mosteknoloji/robomotion-go-lib/runtime"
	"github.com/tj/go-dropbox"
	"github.com/tj/go-dropy"
)

type WriteFile struct {
	runtime.Node `spec:"id=Robomotion.Dropbox.WriteFile,name=Write File,icon=mdiDropbox,color=#007ee5"`

	// Input1
	InClientID runtime.InVariable `spec:"title=Client ID,scope=Message,name=client_id,messageOnly"`

	// Input2
	InFilename runtime.InVariable `spec:"title=File Name,scope=Custom"`

	// Input3
	InContext runtime.InVariable `spec:"title=Context,scope=Custom"`
}

func (n *WriteFile) OnCreate() error {
	return nil
}

func (n *WriteFile) OnMessage(ctx message.Context) (err error) {
	client_id, err := n.InClientID.GetString(ctx)
	if err != nil {
		return err
	}
	token := getToken(client_id)
	client := dropy.New(dropbox.New(dropbox.NewConfig(token)))

	path, err := n.InFilename.GetString(ctx)
	if err != nil {
		return err
	}
	path1, err := n.InContext.GetString(ctx)
	if err != nil {
		return err
	}
	r := strings.NewReader(path1)
	client.Upload(path, r)
	return nil
}

func (n *WriteFile) OnClose() error {
	return nil
}
