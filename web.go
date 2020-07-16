package main

import (
	"log"
	"fmt"
	"net/http"
	"io"
	"os"
)

func main() {
	fileDir := "/tmp/"
	fileUrl := "https://golangcode.com/logo.svg"
	err := DownloadFile(fileDir +"logo.svg", fileUrl)
	if err != nil {
		panic(err)
	}
	fmt.Println("Downloaded: " + fileUrl)

	fs := http.FileServer( http.Dir( "/tmp/" ) )

	http.HandleFunc( "/", func( res http.ResponseWriter, req *http.Request ) {
		res.Header().Set( "Content-Type", "text/html" );
		fmt.Fprint( res, "<h1>This is my go test</h1>" )
	} )

	http.Handle( "/static/", http.StripPrefix( "/static", fs ) )

	log.Fatal(http.ListenAndServe( ":9000", nil ))
}

func DownloadFile(filepath string, url string) error {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}
