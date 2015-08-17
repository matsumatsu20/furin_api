package controllers

import "github.com/robfig/revel"

type Hoge struct {
    *revel.Controller
}

func (c Hoge) Index() revel.Result {
    var hogejson = `json:"sss", sss:"json"`;
    return c.RenderJson(hogejson);
}
