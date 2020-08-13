package main

import (
	"html/template"
)

var tmplHtml = `
<!DOCTYPE html>
<html>
	<head>
		<meta name="viewport" content="width=device-width, initial-scale=1">
		<title>Bybit Funding</title>
		<link rel="stylesheet" href="/static/css/bootstrap.css" />
		<link rel="stylesheet" href="/static/css/style.css" />
	</head>

	<body>
		<div class="container">
			<div class="row my-3">
				<div class="col col-md-4 offset-md-4">
					<div class="row">
						<div class="col">
							Current:
						</div>
						<div class="col">
							{{ .Current }}%
						</div>
					</div>
					<div class="row">
						<div class="col">
							Predicted:
						</div>
						<div class="col">
							{{ .Predicted }}%
						</div>
					</div>
				</div>
			</div>
		</div>
	</body>
</html>
`

var tmpl = template.Must(template.New("index").Parse(tmplHtml))
