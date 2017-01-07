package usertest

import (
	"fmt"
	"net/http"
	"html/template"
	"appengine"
	"appengine/user"
	
)

type TEMPSUB struct {
	
	Email	string
	AuthDomain	string
	ID	string
	Url		string
	
}


func init() {
    http.HandleFunc("/", root)
    http.HandleFunc("/logined", logined)
}

func root(w http.ResponseWriter, r *http.Request){

	c := appengine.NewContext(r)
	
	login_url, err := user.LoginURL(c, "/logined")
	
	if err != nil {
		fmt.Fprintf(w, "error.")
		return
	}

	t_data := &TEMPSUB {
		Url : login_url,
	}
	
	tmpl := template.Must(template.New("ROOT").Parse(root_template))
	tmpl.Execute(w, t_data)
}

var root_template = `
<html>
  <body>
    <a href="{{.Url}}">login</a>
  </body>
<html>
`

func logined(w http.ResponseWriter, r *http.Request){

	c := appengine.NewContext(r)
	
	u := user.Current(c)
	
	logout_url, err := user.LogoutURL(c, "/")
	
	if err != nil {
		fmt.Fprintf(w, "error.")
		return
	
	}
	
	tmpl_data := &TEMPSUB {
		Email : u.Email,
		AuthDomain : u.AuthDomain,
		ID : u.ID,
		Url : logout_url,
	}
	
	tmpl := template.Must(template.New("Logined").Parse(logined_template))
	tmpl.Execute(w, tmpl_data)
	
}

var logined_template = `
<html>
  <body>
    <p> your email : {{.Email}}</p>
    <p> your domain : {{.AuthDomain}}</p>
    <p> your ID : {{.ID}}</p>
    <a href="{{.Url}}">logout</a>
  </body>
<html>
`

