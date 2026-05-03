package main

import (
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"blog.ariellarin.com/internal/models"
)


func (app *application) home(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go") // do this to add to the response header map
	// You must call this before either w.WriteHeader or w.Write

	// Initialize a slice containing the paths to the two files. It's important
    // to note that the file containing our base template must be the first
    // file in the slice.
    
	files := []string{
        "./ui/html/base.html",
		"./ui/html/partials/nav.html",
        "./ui/html/pages/home.html",
    }
    // Use the template.ParseFiles() function to read the files and store the
    // templates in a template set. Notice that we use ... to pass the contents 
    // of the files slice as variadic arguments.
    ts, err := template.ParseFiles(files...)
    if err != nil {
		app.serverError(w, r, err)
        return
    }
	
    // Use the ExecuteTemplate() method to write the content of the "base" 
    // template as the response body.
    err = ts.ExecuteTemplate(w, "base", nil)
    if err != nil {
		app.serverError(w, r, err)
    }
}

func (app *application) blogView(w http.ResponseWriter, r *http.Request) {
    id, err := strconv.Atoi(r.PathValue("id"))
    if err != nil || id < 1 {
        http.NotFound(w, r)
        return
    }

    // Use the SnippetModel's Get() method to retrieve the data for a
    // specific record based on its ID. If no matching record is found,
    // return a 404 Not Found response.
    post, err := app.posts.Get(id)
    if err != nil {
        if errors.Is(err, models.ErrNoRecord) {
            http.NotFound(w, r)
        } else {
            app.serverError(w, r, err)
        }
        return
    }

    // Write the snippet data as a plain-text HTTP response body.
    fmt.Fprintf(w, "%+v", post)
}

func (app *application) blogPostComposer(w http.ResponseWriter, r *http.Request) {
	message := "This is where we can create a new blog post"
	w.Write([]byte(message))
}

func (app *application) createBlogPost(w http.ResponseWriter, r *http.Request) {
	// Create some variables holding dummy data. We'll remove these later on
    // during development.
    title := "My life as an engineer"
    content := "I graduated college in 2025\nNow, I work in Chicago in financial technolgy\nI miss Minnesota"
    expires := 7

    // Pass the data to the BlogPostModel.Insert() method, receiving the
    // ID of the new record back.
    id, err := app.posts.Insert(title, content, expires)
    if err != nil {
        app.serverError(w, r, err)
        return
    }

    // Redirect the user to the relevant page for the snippet.
    http.Redirect(w, r, fmt.Sprintf("/blog/view/%d", id), http.StatusSeeOther)
}



