package usertest

import (
	"fmt"
	"net/http"
	"html/template"
	"appengine"
	"appengine/user"
	
)

type TEMPSUB_OPENID struct {
	
	Email	string
	AuthDomain	string
	ID	string
	User	string
	Url		string
	
}


func init() {
    http.HandleFunc("/userOpenIdTest", rootOpenId)
    http.HandleFunc("/userOpenIdTest/loginOpenId", loginOpenId)
    http.HandleFunc("/userOpenIdTest/loginedOpenId", logined)
}

func rootOpenId(w http.ResponseWriter, r *http.Request){

	c := appengine.NewContext(r)
	c.Debugf("rootOpenId start")

	fmt.Fprintf(w, root_html)

	c.Debugf("rootOpenId end")
}

var root_html = `
<html>
  <body>
    <form action="/userOpenIdTest/loginOpenId" method="post">
    	<input type="text" name="identity" value="input identity" />
    	<input type="submit" value="post" />
    </form>
  </body>
<html>
`

func loginOpenId(w http.ResponseWriter, r *http.Request){

	c := appengine.NewContext(r)
	c.Debugf("loginOpenId start")
	
	identity := r.FormValue("identity")
	
	login_url, err := user.LoginURLFederated(c, "/userOpenIdTest/loginedOpenID",identity)
	
	if err != nil {
		fmt.Fprintf(w, "error.",err)
		return
	}

	
	http.Redirect(w, r, login_url, http.StatusFound)
	c.Debugf("loginOpenId end")
}



func loginedOpenId(w http.ResponseWriter, r *http.Request){

	c := appengine.NewContext(r)
	
	u := user.Current(c)
	//
	if u == nil {
		
		http.Redirect(w, r, "/userOpenIdTest", http.StatusFound)
	}
	
	//
	logout_url, err := user.LogoutURL(c, "/userOpenIdTest")
	
	if err != nil {
		fmt.Fprintf(w, "error.")
		return
	
	}
	
	tmpl_data := &TEMPSUB_OPENID {
		User : u.FederatedIdentity,
		Url : logout_url,
	}
	
	tmpl := template.Must(template.New("Logined").Parse(loginedOpenId_template))
	tmpl.Execute(w, tmpl_data)
	
}

var loginedOpenId_template = `
<html>
  <body>
    <p> your identity : {{.User}}</p>
    <a href="{{.Url}}">logout</a>
  </body>
<html>
`


