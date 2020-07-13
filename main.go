package main

import (
	"io"
	"net/http"
)

var (
	data = `<html>
<body>
<p id="demo2">some</p>
  <div id="demo"></div>
  <button onclick="onClick()">aaaaa</button>
  <script>
	function onClick() {
	  let child = document.createElement("p");
	  child.innerHTML = "aaa";
	  let el = document.getElementById("demo");
	  el.appendChild(child);
	  let el2 = document.getElementById("demo2");
	  el2.innerHTML = "bbb";
	}
  </script>
</body>
</html>
	`
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, data)
	})
	http.ListenAndServe(":9091", nil)
}
