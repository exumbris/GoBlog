package main

import (
	"database/sql"
	"flag"
	"log/slog"
	"net/http"
	"os"

	"blog.ariellarin.com/internal/models"

	_ "github.com/go-sql-driver/mysql"
)

// Define an application struct to hold the application-wide dependencies for the
// web application. This is for dependency injection.
type application struct {
    logger *slog.Logger
	posts *models.BlogPostModel
}


/*
Purpose of our main:
1. Parisng the runtime configuration settings for the application
2. Establishing the dependencies for the handlers
3. Running the HTTP server
*/

func main() {
    addr := flag.String("addr", ":4000", "HTTP network address")
	dsn := flag.String("dsn", "web:pass@/blog?parseTime=true", "MySQL data source name")
    flag.Parse()

    logger := slog.New(slog.NewTextHandler(os.Stdout, nil))



    db, err := openDB(*dsn)
    if err != nil {
        logger.Error(err.Error())
        os.Exit(1)
    }


	defer db.Close() // we wait until the main function ends; then close DB


	// This means that app is of type *application (a pointer to an application struct)
    app := &application{ // creates a pointer to the struct
        logger: logger,
		posts: &models.BlogPostModel{DB: db},
    }

    logger.Info("starting server", "addr", *addr)



    
    // Call the new app.routes() method to get the servemux containing our routes,
    // and pass that to http.ListenAndServe().
    err = http.ListenAndServe(*addr, app.routes())
    logger.Error(err.Error())
    os.Exit(1)
}

// The openDB() function wraps sql.Open() and returns a sql.DB connection pool
// for a given DSN.
func openDB(dsn string) (*sql.DB, error) {
    db, err := sql.Open("mysql", dsn) // initializes database connection pool for future use
    if err != nil {
        return nil, err
    }

    err = db.Ping()
    if err != nil {
        db.Close()
        return nil, err
    }

    return db, nil
}