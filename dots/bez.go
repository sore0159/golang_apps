package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

const MAXSIZE int64 = 1048576

var BEZPTS [][2]float64

func bezAPI(w http.ResponseWriter, r *http.Request) {
	var bytes []byte
	var err error
	bytes, err = ioutil.ReadAll(io.LimitReader(r.Body, MAXSIZE))
	if err != nil {
		http.Error(w, "read body failure", 500)
		return
	}
	err = r.Body.Close()
	if err != nil {
		http.Error(w, "read body close failure", 500)
		return
	}
	bez := [][2]float64{}
	err = json.Unmarshal(bytes, &bez)
	if err != nil {
		http.Error(w, "json unmarshal failure", 500)
		return
	}
	if len(bez) != 3 && len(bez) != 4 {
		http.Error(w, "invalid bez points", 500)
		return
	}
	BEZPTS = bez
	fmt.Fprintf(w, "OK")
}

func bezPNG(w http.ResponseWriter, r *http.Request) {
	if len(BEZPTS) == 3 || len(BEZPTS) == 4 {
		DrawBez(w, BEZPTS)
		return
	}
	http.Error(w, "invalid bez points", 500)
}

func bezHTML(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, `<!doctype html>
<html>
<head>
<meta charset="utf-8">
<title>Bezier Curves</title>
<link rel="shortcut icon" href="yd32.ico">
</head>
<body>
<h2 style="margin:0px;">Bezier Curves: draw me like one of your french mathematical equations</h2>
        <center>
        <canvas id="can" width=500 height=500 style="margin-top:0px;border: 2px solid black"></canvas>
        <img id="pic" width=500 height=500 style="margin-top:10px;border: 2px solid black">
		<br>
		<button style="font-size:x-large" id="changeButton">Quadratic</button>
		<button style="font-size:x-large" id="subButton">Submit</button>

        </center>
</body>
<script src="bez.js" type="text/javascript"></script>

</html>`)
}

func bezJS(w http.ResponseWriter, r *http.Request) {
	f, err := os.Open("bez.js")
	if err != nil {
		http.Error(w, "can't open bez.js", 500)
		return
	}
	w.Header().Set("Content-Type", "application/javascript")
	if _, err = io.Copy(w, f); err != nil {
		fmt.Println("Error copying js to http: ", err)
	}
}
