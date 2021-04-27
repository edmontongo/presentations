package google_cloud_functions

const exampleHTML = `{{define "header" -}}
<html>
	<head>
		<title>Live, from google cloud!</title>
		<link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/4.1.0/css/bootstrap.min.css">
		<script src="https://code.jquery.com/jquery-3.3.1.slim.min.js" integrity="sha384-q8i/X+965DzO0rT7abK41JStQIAqVgRVzpbzo5smXKp4YfRvH+8abtTE1Pi6jizo" crossorigin="anonymous"></script>
		<script src="https://stackpath.bootstrapcdn.com/bootstrap/4.1.0/js/bootstrap.min.js" integrity="sha384-uefMccjFJAIv6A+rW+L4AHf99KvxDjWSu1z9VI8SKNVmz4sk7buKt/6v9KI65qnm" crossorigin="anonymous"></script>
		<meta name="viewport" content="width=device-width, initial-scale=1.0">

		<style type="text/css">
			h3 {
				margin-top: 1em;
			}
		</style>
	</head>
	<body>
		<nav class="navbar navbar-expand-lg navbar-dark bg-dark">
			<div class="container">
				<a class="navbar-brand" href="#">I'm a cloud function!</a>
			</div>
		</nav><div class="container">
{{- end}}

{{define "footer" -}}
</div>
	</body>
</html>
{{- end}}

{{define "menu"}}
		<ul class="list-group row">
			<div class="list-group-item form-group container">
				<a href="/HTML/">Home</a>
			</div>
			<div class="list-group-item form-group container">
				<a href="/HTML/A">Page A</a>
			</div>
			<div class="list-group-item form-group container">
				<a href="/HTML/B">Page B</a>
			</div>
		</ul>
{{end}}

{{define "home"}}
{{template "header"}}
<h3>This is the home page</h3>
{{template "menu"}}
{{template "footer"}}
{{end}}

{{define "pageA"}}
{{template "header"}}
<h3>This is page A</h3>
{{template "menu"}}
{{template "footer"}}
{{end}}

{{define "pageB"}}
{{template "header"}}
<h3>This is page B!</h3>
{{template "menu"}}
{{template "footer"}}
{{end}}
`
