package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)


func (app *application) createBookHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create new book..")
}
func (app *application) showBookHandler(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	id, err :=strconv.ParseInt(params.ByName("id"), 10, 64)
	 if err != nil || id < 1 {
        http.NotFound(w, r)
        return
    }
   
    fmt.Fprintf(w, "show the details of book %d\n", id)
}


