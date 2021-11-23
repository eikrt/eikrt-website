package main

import (
    "log"
    "net/http"
    "html/template"
)

func frontpageHandler (w http.ResponseWriter, r *http.Request) {
        files := []string {


                "templates/frontpage.html",
                "templates/base.html",
        }
        t, err := template.ParseFiles(files...)
        if err != nil {
                log.Println(err.Error())
                http.Error(w, "Internal Server Error", 500)
                return
        }
        
        err = t.Execute(w,nil)
        if err != nil {
                log.Println(err.Error())
                http.Error(w, "Internal Server Error", 500)
        }
}

func reviewsHandler (w http.ResponseWriter, r *http.Request) {
        files := []string {


                "templates/reviews.html",
                "templates/base.html",
        }
        t, err := template.ParseFiles(files...)
        if err != nil {
                log.Println(err.Error())
                http.Error(w, "Internal Server Error", 500)
                return
        }
        
        err = t.Execute(w,nil)
        if err != nil {
                log.Println(err.Error())
                http.Error(w, "Internal Server Error", 500)
        }
}
func projectsHandler (w http.ResponseWriter, r *http.Request) {
        files := []string {


                "templates/projects.html",
                "templates/base.html",
        }
        t, err := template.ParseFiles(files...)
        if err != nil {
                log.Println(err.Error())
                http.Error(w, "Internal Server Error", 500)
                return
        }
        
        err = t.Execute(w,nil)
        if err != nil {
                log.Println(err.Error())
                http.Error(w, "Internal Server Error", 500)
        }
}


func tutorialsHandler (w http.ResponseWriter, r *http.Request) {
        files := []string {


                "templates/tutorials.html",
                "templates/base.html",
        }
        t, err := template.ParseFiles(files...)
        if err != nil {
                log.Println(err.Error())
                http.Error(w, "Internal Server Error", 500)
                return
        }
        
        err = t.Execute(w,nil)
        if err != nil {
                log.Println(err.Error())
                http.Error(w, "Internal Server Error", 500)
        }
}

func aboutHandler (w http.ResponseWriter, r *http.Request) {
        files := []string {


                "templates/about.html",
                "templates/base.html",
        }
        t, err := template.ParseFiles(files...)
        if err != nil {
                log.Println(err.Error())
                http.Error(w, "Internal Server Error", 500)
                return
        }
        
        err = t.Execute(w,nil)
        if err != nil {
                log.Println(err.Error())
                http.Error(w, "Internal Server Error", 500)
        }
}

func main() {

    http.HandleFunc("/", frontpageHandler)
    http.HandleFunc("/projects", projectsHandler)
    http.HandleFunc("/reviews", reviewsHandler)
    http.HandleFunc("/tutorials", tutorialsHandler)
    http.HandleFunc("/about", aboutHandler)
    http.Handle("/style/", http.StripPrefix("/style/", http.FileServer(http.Dir("./style/")))) 
    http.Handle("/res/", http.StripPrefix("/res/", http.FileServer(http.Dir("./res/")))) 
    log.Fatal(http.ListenAndServe(":8080", nil))
}
