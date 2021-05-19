package binding

import (
	"encoding/json"
	"net/http"
	"testing"
)

type model4 struct {
	ID    int    `binding:"name(model2.model3.model4.id)"`
	Name  string `binding:"name(model2.model3.model4.name)"`
	ID2   int    `binding:"name(xxx)"`
	Name2 string `binding:"name(yyy)"`
}

type model3 struct {
	ID     int    `binding:"name(model2.model3.id)"`
	Name   string `binding:"name(model2.model3.name)"`
	Model4 model4 `binding:"name(model4)"`
}

type model2 struct {
	ID     int    `binding:"name(model2.id)"`
	Name   string `binding:"name(model2.name)"`
	Model3 model3 `binding:"name(model3)"`
}

type model struct {
	ID     int    `binding:"name(id)"`
	Name   string `binding:"name(name)"`
	Model2 model2 `binding:"name(model2)"`
}

func buildParamReq() *http.Request {
	req, _ := http.NewRequest(http.MethodGet, "", nil)
	req.Form = map[string][]string{}
	req.Form.Set("id", "100")
	req.Form.Set("name", "zhangsan")
	req.Form.Set("model2.id", "200")
	req.Form.Set("model2.name", "lisi")
	req.Form.Set("model2.model3.id", "300")
	req.Form.Set("model2.model3.name", "wangwu")
	req.Form.Set("model2.model3.model4.id", "400")
	req.Form.Set("model2.model3.model4.name", "zhaoliu")
	req.Form.Set("xxx", "400")
	req.Form.Set("yyy", "zha111111oliu")
	return req
}

func buildHeaderReq() *http.Request {
	req, _ := http.NewRequest(http.MethodGet, "", nil)
	req.Header = map[string][]string{}
	req.Header.Set("id", "100")
	req.Header.Set("name", "zhangsan")
	req.Header.Set("model2.id", "200")
	req.Header.Set("model2.name", "lisi")
	req.Header.Set("model2.model3.id", "300")
	req.Header.Set("model2.model3.name", "wangwu")
	req.Header.Set("model2.model3.model4.id", "400")
	req.Header.Set("model2.model3.model4.name", "zhaoliu")
	return req
}

func TestBindingReqParam(t *testing.T) {
	m := &model{}
	binding := NewWithType(m, Param)
	binding.BindReq(buildParamReq())
	bs, _ := json.Marshal(m)
	t.Log(string(bs))
}

func TestBindingReqHeader(t *testing.T) {
	m := &model{}
	binding := NewWithType(m, Header)
	binding.BindReq(buildHeaderReq())
	bs, _ := json.Marshal(m)
	t.Log(string(bs))
}
