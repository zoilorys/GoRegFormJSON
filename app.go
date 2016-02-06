package main

import (
  "net/http"
  "html/template"
  "user"
)

var templates = template.Must(template.ParseFiles("src/templates/index.html",
  "src/templates/list.html",
  "src/templates/error.html",
  "src/templates/success.html",
  "src/templates/failure.html"))

func main() {
  http.HandleFunc("/", indexHandler)
  http.HandleFunc("/view/", viewHandler)
  http.HandleFunc("/register", registerHandler)
  http.HandleFunc("/assets/", staticHandler)
  http.ListenAndServe(":8080", nil)
}

func indexHandler(res http.ResponseWriter, req *http.Request) {
  err := templates.ExecuteTemplate(res, "index.html", nil)
  if err != nil {
    ierr := templates.ExecuteTemplate(res, "error.html", struct{Text string}{Text: err.Error()})
    if ierr != nil {
      http.Error(res, ierr.Error(), http.StatusInternalServerError)
    }
  }
}

func viewHandler(res http.ResponseWriter, req *http.Request) {
  users, err := user.List()
  if err != nil {
    sendError(res, err.Error())
  }

  perr := templates.ExecuteTemplate(res, "list.html", struct{Users []user.User}{Users: users})
  if perr != nil {
    sendError(res, perr.Error())
  }
}

func registerHandler(res http.ResponseWriter, req *http.Request) {
  user := user.User{
    Name: req.FormValue("name"),
    Email: req.FormValue("email"),
    Password: req.FormValue("password")}

  rerr := user.Save()
  if rerr != nil {
    ferr := templates.ExecuteTemplate(res, "failure.html", nil)
    if ferr != nil {
      sendError(res, ferr.Error())
    }
  }

  serr := templates.ExecuteTemplate(res, "success.html", nil)
  if serr != nil {
    sendError(res, serr.Error())
  }
}

func sendError(res http.ResponseWriter, text string) {
  err := templates.ExecuteTemplate(res, "error.html", struct{Text string}{Text: text})
  if err != nil {
    http.Error(res, err.Error(), http.StatusInternalServerError)
  }
}

func staticHandler(res http.ResponseWriter, req *http.Request) {
  path := req.URL.Path[1:]
  http.ServeFile(res, req, path)
}
