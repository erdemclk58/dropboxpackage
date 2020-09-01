package nodes

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"

	"bitbucket.org/mosteknoloji/robomotion-go-lib/message"
	"bitbucket.org/mosteknoloji/robomotion-go-lib/runtime"
	"github.com/atotto/clipboard"
	"github.com/tj/go-dropbox"
	"github.com/tj/go-dropy"
)

type DownloadFile struct {
	runtime.Node `spec:"id=Robomotion.Dropbox.DownloadFile,name=Download File,icon=mdiDropbox,color=#007ee5"`

	// Input1
	InClientID runtime.InVariable `spec:"title=Client ID,scope=Message,name=client_id,messageOnly"`

	// Input2
	InDropboxPath runtime.InVariable `spec:"title=Dropbox Path,scope=Custom"`

	// Input3
	InLocalPath runtime.InVariable `spec:"title=Local Path,scope=Custom"`
}

func (n *DownloadFile) OnCreate() error {
	return nil
}

func (n *DownloadFile) OnMessage(ctx message.Context) (err error) {
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

	path1, err := n.InLocalPath.GetString(ctx)
	if err != nil {
		return err
	}
	reader, errdown := client.Download(path)
	if errdown != nil {
		log.Fatal(errdown)
	}
	defer reader.Close()

	data := []byte{}

	for {
		buffer := make([]byte, 1024)
		n, errdown := reader.Read(buffer)
		if errdown != nil && errdown != io.EOF {
			log.Fatal(errdown)
		}

		data = append(data, buffer[:n]...)
		if errdown == io.EOF {
			break
		}
	}

	errdown = ioutil.WriteFile(path1, data, 0644)
	if errdown != nil {
		log.Fatal(errdown)
	}

	return nil
}

func (n *DownloadFile) OnClose() error {
	return nil
}

func setClipboard(text string) (err error) {
	clipboard.WriteAll(fmt.Sprintf("%v", text))
	return nil
}
