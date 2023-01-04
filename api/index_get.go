package api

import (
	"fmt"
	"strings"

	"github.com/1nv8rzim/otn/config"
	"github.com/gin-gonic/gin"
)

func index_get(ctx *gin.Context) {
	index := fmt.Sprintf(`
	<!DOCTYPE html>
	<html>
		<head>
			<title>OTN</title>
		</head>
		<style>
			body {
				width: 500px;
				margin: auto;
			}
		</style>
		<body>
			<pre>
<center>OTN - One Time Note</center>

Description:
	OTN lets you create notes from the commandline
	These notes can only be read once and will be deleted afterwards
	Likewise, the notes are only stored for a limited time

Usage:
	&lt;command&gt; | curl -F 'content=<-' %v

Source Code:
	<a href="https://github.com/1nv8rZim/otn">https://github.com/1nv8rZim/otn</a>
			</pre>
		</body>
	</html>
	`, strings.Split(config.URL, "//")[1])

	ctx.Data(200, "text/html; charset=utf-8", []byte(index))
}
