package models

import (
	"database/sql"
	"errors"
	"time"
)

// Define a Post type to hold the data for an individual blog post. Notice how
// the fields of the struct correspond to the fields in our MySQL blogPosts
// table?
type Post struct {
    ID      int
    Title   string
    Content string
    Created time.Time
    Expires time.Time
}

// Define a BlogPostModel type which wraps a sql.DB connection pool.
type BlogPostModel struct {
    DB *sql.DB
    /*A sql.DB is not a single connection, but a connection pool manager.
    Basically, we have one long-lived sql.DB object that manages many real TCP connctions
    to MySQL behind the scenes. Each query borrows one connection from the pool, uses it, then returns it. 

    Without connection pooling, every queue opens/closes a fresh TCP connection; slow and expensive.
    But, with pooling, we can resuse existing connections leading to better performance, concurrency, and less DB churn.
    */
}

// This will insert a new blog post into the database.
func (m *BlogPostModel) Insert(title string, content string, expires int) (int, error) {
    // Write the SQL statement we want to execute.
    stmt := "INSERT INTO blogPosts (title, content, created, expires) VALUES(?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))"

    // Use the Exec() method on the embedded connection pool to execute the
    // statement. The first parameter is the SQL statement, followed by the
    // values for the placeholder parameters: title, content and expiry in
    // that order. This method returns a sql.Result type, which contains some
    // basic information about what happened when the statement was executed.
    result, err := m.DB.Exec(stmt, title, content, expires)
    if err != nil {
        return 0, err
    }

    // Use the LastInsertId() method on the result to get the ID of our
    // newly inserted record in the posts table.
    id, err := result.LastInsertId()
    if err != nil {
        return 0, err
    }

    // The ID returned has the type int64, so we convert it to an int type
    // before returning.
    return int(id), nil
}

// This will return a specific blog post based on its id.
func (m *BlogPostModel) Get(id int) (Post, error) {
    // Write the SQL statement we want to execute.
    stmt := "SELECT id, title, content, created, expires FROM blogPosts WHERE expires > UTC_TIMESTAMP() AND id = ?"

    // Use the QueryRow() method on the connection pool to execute our
    // SQL statement, passing in the untrusted id variable as the value for the
    // placeholder parameter. This returns a pointer to a sql.Row value which
    // holds the result from the database.
    row := m.DB.QueryRow(stmt, id)

    // Initialize a new zeroed Post struct.
    var s Post

    // Use row.Scan() to copy the values from each field in sql.Row to the
    // corresponding field in the Post struct. Notice that the arguments
    // to row.Scan are *pointers* to the place you want to copy the data into,
    // and the number of arguments must be exactly the same as the number of
    // columns returned by your statement.
    err := row.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
    if err != nil {
        // If the query returns no rows, then row.Scan() will return a
        // sql.ErrNoRows error. We use the errors.Is() function check for that
        // error specifically, and return our own ErrNoRecord error
        // instead (we'll create this in a moment).
        if errors.Is(err, sql.ErrNoRows) {
            return Post{}, ErrNoRecord
        } else {
            return Post{}, err
        }
    }

    // If everything went OK, then return the filled Post struct.
    return s, nil
}

// This will return the 10 most recently created blog posts.
func (m *BlogPostModel) Latest() ([]Post, error) {
        // Grab ten most recent blog posts
		stmt := "SELECT id, title, content, created, expires FROM blogPosts WHERE expires > UTC_TIMESTAMP() ORDER BY id DESC LIMIT 10"
	
		// Use the Query() method on the connection pool to execute our
		// SQL statement. This returns a sql.Rows resultset containing the result of
		// our query.
		rows, err := m.DB.Query(stmt)
		if err != nil {
			return nil, err
		}
	
		// We defer rows.Close() to ensure the sql.Rows resultset is
		// always properly closed before the Latest() method returns. This defer
		// statement should come *after* you check for an error from the Query()
		// method. Otherwise, if Query() returns an error, you'll get a panic
		// trying to close a nil resultset.
		defer rows.Close()
	
		// Initialize an empty slice to hold the Post structs.
		var posts []Post
	
		// Use rows.Next to iterate through the rows in the resultset. This
		// prepares the first (and then each subsequent) row to be acted on by the
		// rows.Scan() method. If iteration over all the rows completes then the
		// resultset automatically closes itself and frees up the underlying
		// database connection.
		for rows.Next() {
			// Create a new zero value Post struct.
			var s Post
			// Use rows.Scan() to copy the values from each field in the row to the
			// new Post struct that we created. Again, the arguments to row.Scan()
			// must be pointers to the place you want to copy the data into, and the
			// number of arguments must be exactly the same as the number of
			// columns returned by your statement.
			err = rows.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
			if err != nil {
				return nil, err
			}
			// Append it to the slice of posts.
			posts = append(posts, s)
		}
	
		// When the rows.Next() loop has finished we call rows.Err() to retrieve any
		// error that was encountered during the iteration. It's important to
		// call this - don't assume the iteration completed successfully over the 
		// entire result set.
		if err = rows.Err(); err != nil {
			return nil, err
		}
	
		// If everything went OK then return the posts slice.
		return posts, nil
}

func (m *BlogPostModel) NonExpired() (int64, error) {
    stmt := "SELECT COUNT(*) FROM blogPosts WHERE expires > UTC_TIMESTAMP()"
    row := m.DB.QueryRow(stmt)

    var numNonExpiredRows int64
    /*The database driver (for MySQL, Postgres, etc.) always returns aggregate results
    like COUNT(*) as a 64-bit integer, because the database itself can store and return
    very large numbers (billions or trillions of rows)
    
    If you use int in Go, it might be 32 bits on some systems, which could overflow if 
    the database ever returns a value larger than 2,147,483,647. By using int64, you
    guarantee that your Go variable can always hold the value the database returns,
    no matter the machine or the size of the table.*/
    
    err := row.Scan(&numNonExpiredRows) //pointer to where we want to copy data to

    if err != nil {
        return 0, err
    }

    return numNonExpiredRows, nil
	
}


func (m *BlogPostModel) ExpiredCount() (int64, error) {
    stmt := "SELECT COUNT(*) FROM blogPosts WHERE expires <= UTC_TIMESTAMP()"
    row := m.DB.QueryRow(stmt)

    var numExpiredRows int64
    
    err := row.Scan(&numExpiredRows) //pointer to where we want to copy data to

    if err != nil {
        return 0, err
    }

    return numExpiredRows, nil
}

func (m *BlogPostModel) ExistsByID(ID int) (bool, error) {
    stmt := "SELECT EXISTS (SELECT 1 FROM blogPosts WHERE id = ? AND expires > UTC_TIMESTAMP())"

    row := m.DB.QueryRow(stmt, ID)

    var idExists int
    
    err := row.Scan(&idExists)

    if err != nil {
        return false, err
    }

    if idExists == 0 {
        return false, nil
    }
    return true, nil
}

func (m *BlogPostModel) LatestByLimit(numNonExpiredLimit int) ([]Post, error) {
		stmt := "SELECT id, title, content, created, expires FROM blogPosts WHERE expires > UTC_TIMESTAMP() ORDER BY id DESC LIMIT ?"

		rows, err := m.DB.Query(stmt, numNonExpiredLimit)
		if err != nil {
			return nil, err
		}
	
		defer rows.Close()
	
		var posts []Post
	
		for rows.Next() {
			var s Post

			err = rows.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
			if err != nil {
				return nil, err
			}

			posts = append(posts, s)
		}
	
		if err = rows.Err(); err != nil {
			return nil, err
		}
	
		return posts, nil
}