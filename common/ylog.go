package common

import (
	"encoding/json"
	"fmt"
	"github.com/kataras/golog"
	"github.com/lhlyu/iyu/util"
	"os"
	"strconv"
	"strings"
	"time"
)

const _json = "json"
const _timeFormat = "2006-01-02 15:04:05"

type ylog struct {
	g          *golog.Logger
	outWay     string
	timeFormat string
	level      string
}

type logJson struct {
	Id uint64 `json:"gid"`
	T  string `json:"time"`
	P  string `json:"position"`
	C  string `json:"content"`
}

func NewYlog(level, timeFormat, outFile, outWay string) *ylog {
	g := golog.New()
	yg := &ylog{
		timeFormat: _timeFormat,
	}
	if level != "" {
		g.SetLevel(level)
		yg.level = level
	}
	if timeFormat != "" {
		g.SetTimeFormat(timeFormat)
		yg.timeFormat = timeFormat
	}
	if outFile != "" {
		fl, e := os.Create(outFile)
		if e != nil {
			panic(e)
		}
		g.SetOutput(fl)
	}
	if outWay == _json {
		yg.outWay = _json
	}
	return yg
}

func (y *ylog) Debug(v ...interface{}) {
	gid := util.GetGID()
	funcName, fileName, line := util.CurrentInfo(2)
	lgJson := logJson{
		Id: gid,
		T:  time.Now().Format(y.timeFormat),
		P:  strings.Join([]string{funcName, fileName, strconv.Itoa(line)}, " "),
		C:  fmt.Sprint(v),
	}
	if y.outWay == _json {
		bytes, _ := json.Marshal(lgJson)
		y.g.Print(string(bytes))
	} else {
		y.g.Debugf("| %v | %v | %v\n", lgJson.Id, lgJson.P, lgJson.C)
	}
}
