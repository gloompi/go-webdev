package main

import "fmt"

func main() {
	name := "Esenzhanov Kubanychbek"

	tpl := `
	<!DOCTYPE html>
	<html lang="en">
	<head>
	<meta charset="UTF-8">
	<title>Yo! Wassap everybody</title>
	</head>
	<body>
	<h1>` + name + `</h1>
	</body>
	</html>
	`

	fmt.Println(tpl)
}
