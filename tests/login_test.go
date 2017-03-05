package test

import (
	"testing"
	"runtime"
	"strings"
	"net/url"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	
	"github.com/astaxie/beego"
	. "github.com/smartystreets/goconvey/convey"

	_ "dreamfish/dfsso/routers"
	"dreamfish/dfsso/models/loginModel"
	"dreamfish/dfsso/models/initDBModel"
)

func init() {
	_, file, _, _ := runtime.Caller(1)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".." + string(filepath.Separator))))
	beego.TestBeegoInit(apppath)

	initDBModel.RegisterDB()
}

func TestLoginModel(t *testing.T) {
	username := "admin"
	password := "12"
	_, err := loginModel.Login(username, password)	

	Convey("Subject: Test Station Endpoint\n", t, func() {
	        Convey("err should be nil", func() {
	                So(err, ShouldEqual, nil)
	        })
	})
}

func TestLoginGet(t *testing.T) {
	r, _ := http.NewRequest("GET", "/login", nil)
	r.RequestURI = "/login"
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	Convey("Subject: Test Station Endpoint\n", t, func() {
	        Convey("Status Code Should Be 200", func() {
	                So(w.Code, ShouldEqual, 200)
	        })
	        Convey("The Result Should Not Be Empty", func() {
	                So(w.Body.Len(), ShouldBeGreaterThan, 0)
	        })
	})
}

func TestLoginPost(t *testing.T) {
	form := url.Values{}
    form.Set("username", "123")
    form.Set("password", "456")
    b := strings.NewReader(form.Encode())
	r, _ := http.NewRequest("POST", "/login", b)
	r.RequestURI = "/login"
	w := httptest.NewRecorder()
	
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	Convey("Subject: Test Station Endpoint\n", t, func() {
	        Convey("Status Code Should Be 200", func() {
	                So(w.Code, ShouldEqual, 200)
	        })
	        Convey("The Result Should Not Be Empty", func() {
	                So(w.Body.Len(), ShouldBeGreaterThan, 0)
	        })
	})
}



