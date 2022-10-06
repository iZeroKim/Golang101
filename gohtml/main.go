package main

import ("net/http"
		"fmt"
)

type Info struct{
	Title string
	Name string
	Fruits [3]string
}

func main(){

	const doc =`
	<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta http-equiv="X-UA-Compatible" content="IE=edge">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>{{.Title}}</title>
	</head>
	<body>
		<h3>Hi, {{.Name}}. The fruits are:</h3>
		<ul>
			{{range .Fruit}}
				<li>
					{{.}}
				</li>
			{{end}}
		</ul>
	</body>
	</html>
	`
	fmt.Println("Starting new server at: http://localhost:8000")
	http.ListenAndServe(":8000", nil)
}