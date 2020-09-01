package nodes

import (
	"time"

	"bitbucket.org/mosteknoloji/robomotion-go-lib/message"
	"bitbucket.org/mosteknoloji/robomotion-go-lib/runtime"
	"github.com/tj/go-dropbox"
	"github.com/tj/go-dropy"
)

type Stat struct {
	runtime.Node `spec:"id=Robomotion.Dropbox.Stat,name=Stat,icon=mdiDropbox,color=#007ee5"`

	// Input1
	InClientID runtime.InVariable `spec:"title=Client ID,scope=Message,name=client_id,messageOnly"`

	// Input2
	InDropboxPath runtime.InVariable `spec:"title=Dropbox Path,scope=Custom"`

	//Output

	OutStats runtime.OutVariable `spec:"title=Stats,scope=Message,name=stats,messageOnly"`
}

type Stats struct {
	IsDir   bool      `json:"is_dir"`
	ModTime time.Time `json:"mod_time"`
	Name    string    `json:"name"`
	Size    int64     `json:"size"`
}

func (n *Stat) OnCreate() error {
	return nil
}

func (n *Stat) OnMessage(ctx message.Context) (err error) {

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

	info, err := client.Stat(path)
	if err != nil {
		return err
	}

	v := &Stats{
		Size:    info.Size(),
		Name:    info.Name(),
		IsDir:   info.IsDir(),
		ModTime: info.ModTime(),
	}

	err = n.OutStats.Set(ctx, v)
	if err != nil {
		return err
	}

	return nil

}

func (n *Stat) OnClose() error {
	return nil
}
