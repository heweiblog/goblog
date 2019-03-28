package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/sevlyar/go-daemon"
	"goblog/models"
	_ "goblog/routers"
	"log"
	"os"
)

func init() {
	models.RegisterDB()
}

func main() {
	if 2 == len(os.Args) && os.Args[1] == "-b" {
		cntxt := &daemon.Context{
			PidFileName: "goblog_pid",
			PidFilePerm: 0644,
			LogFileName: "goblog.log",
			LogFilePerm: 0640,
			WorkDir:     "./",
			Umask:       027,
			Args:        []string{"[goblog]"},
		}

		d, err := cntxt.Reborn()
		if err != nil {
			log.Fatal("Unable to run: ", err)
		}
		if d != nil {
			return
		}
		defer cntxt.Release()
	}

	//orm.Debug = true
	orm.Debug = false
	orm.RunSyncdb("default", false, true)
	beego.Run()
}
