package controllers

import "github.com/robfig/revel"
import "net/http"

type Html string

type App struct {
	*revel.Controller
}

func (r Html) Apply(req *revel.Request, resp *revel.Response) {
	resp.WriteHeader(http.StatusOK, "text/html")
	resp.Out.Write([]byte(r))
}

func (c App) Index(msg string) revel.Result {
	return Html(msg)
}
