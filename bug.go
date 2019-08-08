package zeitbug

import (
  "bytes"
  "fmt"
  "io"
  "net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
  buf := &bytes.Buffer{}
  io.Copy(buf, r.Body)
  w.Write([]byte(fmt.Sprintf("Got %d bytes of request body:\n", buf.Len())))
  io.Copy(w, buf)
}
