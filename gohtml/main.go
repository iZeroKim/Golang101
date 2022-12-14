package main

import ("net/http"
		"fmt"
		"text/template"
		"github.com/gorilla/mux"
)

// type Info struct{
// 	Title string
// 	Name string
// 	Fruits [3]string
// }
const(
	CONN_HOST ="localhost"
	CONN_PORT = "8080"
)

var GetRequestHandler = http.HandlerFunc(
	func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, "<h1>Get Request</h1>")
	}
)

var PostRequestHandler = http.HandlerFunc( 
	func(w http.ResponseWriter, req *http.Request){
		fmt.Fprintf(w, "<h1>Post Request</h1>")
	}
)

var PassedVariableHandler = http.HandlerFunc(
	func(w http.ResponseWriter, req *http.Request){
		variables := mux.Vars(r)
		name := vars[name]
		fmt.Fprintf(w, "<h1>Passed Variable: " + name + "</h1>") 
	}
)

func main(){

	// const doc =`
	// <!DOCTYPE html>
	// <html lang="en">
	// <head>
	// 	<meta charset="UTF-8">
	// 	<meta http-equiv="X-UA-Compatible" content="IE=edge">
	// 	<meta name="viewport" content="width=device-width, initial-scale=1.0">
	// 	<title>{{.Title}}</title>
	// </head>
	// <body>
	// 	<h3>Hi, {{.Name}}. The fruits are:</h3>
	// 	<ul>
	// 		{{range .Fruits}}
	// 			<li>
	// 				{{.}}
	// 			</li>
	// 		{{end}}
	// 	</ul>
	// </body>
	// </html>
	// `
	// http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request){
	// 	w.Header().Add("Content Type", "text/html")
	// 	templates:= template.New("template")
	// 	templates.New("doc").Parse(doc)
	// 	info := Info{
	// 		"My Fruits",
	// 		"Kim",
	// 		[3]string{
	// 			"Sweet melon",
	// 			"Water melon",
	// 			"Pineapple",
	// 		},
	// 	}
	// 	templates.Lookup("doc").Execute(w, info)
	// })
	// fmt.Println("Starting new server at: http://localhost:8000")
	// http.ListenAndServe(":8000", nil)

	router := mux.NewRouter()
	router.Handle("/", GetRequestHandler).Methods("GET")
	router.Handle("/post", PostRequestHandler).Methods("POST")
	router.Handle("/hello/{name}", PassedVariableHandler).Methods("GET", "PUT")

	
	http.ListenAndServe(CONN_HOST+":"+CONN_PORT, router)
}