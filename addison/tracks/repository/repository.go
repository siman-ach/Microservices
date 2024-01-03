package repository

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

// Repository struct holds a reference to the *sql.DB object
type Repository struct {
	DB *sql.DB
}

var repo Repository

// Init function initializes the repository by connecting to the SQLite database at /tmp/test.db.
func Init() {
	if db, err := sql.Open("sqlite3", "/tmp/test.db"); err == nil {
		repo = Repository{DB: db}
	} else {
		// Log error and message when failing to create table
		log.Printf("Database initialisation: %v", err)
	}
}

// Create function creates a new Tracks table in the database if it doesn't already exist.
func Create() int {
	const sql = "CREATE TABLE IF NOT EXISTS Tracks(Id TEXT PRIMARY KEY, Audio TEXT)"
	if _, err := repo.DB.Exec(sql); err == nil {
		return 0
	} else {
		// Log error and message when failing to delete tracks
		log.Printf("Failed to create table: %v", err)
		return -1
	}
}

// Clear function deletes all records from the Tracks table.
func Clear() int {
	const sql = "DELETE FROM Tracks"
	if _, err := repo.DB.Exec(sql); err == nil {
		return 0
	} else {
		// Log error and message when failing to read track
		log.Printf("Failed to delete tracks: %v", err)
		return -1
	}
}

// Read function retrieves a music track with the specified ID from the Tracks table.
func Read(id string) (Track, int64) {
	const sql = "SELECT * FROM Tracks WHERE Id = ?"
	if stmt, err := repo.DB.Prepare(sql); err == nil {
		defer stmt.Close()
		var t Track
		row := stmt.QueryRow(id)
		if err := row.Scan(&t.Id, &t.Audio); err == nil {
			return t, 1
		} else {
			// Log error and message when failing to read track
			log.Printf("Failed to read track: %v", err)
			return Track{}, 0
		}
	} else {
		// Log error and message when failing to prepare SQL read statement
		log.Printf("Failed to prepare SQL read statement: %v", err)
		return Track{}, -1
	}
}

// Update function updates the audio data of a music track with the specified ID in the Tracks table.
func Update(t Track) int64 {
	const sql = "UPDATE Tracks SET Audio = ? WHERE Id = ?"
	if stmt, err := repo.DB.Prepare(sql); err == nil {
		defer stmt.Close()
		if res, err := stmt.Exec(t.Audio, t.Id); err == nil {
			if n, err := res.RowsAffected(); err == nil {
				return n
			} else {
				log.Printf("Failed to read rows affected: %v", err)
			}
		} else {
			log.Printf("Failed to execute SQL update statement: %v", err)
		}
	} else {
		log.Printf("Failed to prepare SQL update statement: %v", err)
	}
	return -1
}

// Insert function inserts a new music track into the Tracks table.
func Insert(t Track) int64 {
	const sql = "INSERT INTO Tracks(Id, Audio) VALUES (?, ?)"
	if stmt, err := repo.DB.Prepare(sql); err == nil {
		defer stmt.Close()

		if res, err := stmt.Exec(t.Id, t.Audio); err == nil {
			if n, err := res.RowsAffected(); err == nil {
				return n
			} else {
				log.Printf("Failed to read rows affected: %v", err)
			}
		} else {
			log.Printf("Failed to execute SQL insert statement: %v", err)
		}
	} else {
		log.Printf("Failed to prepare SQL insert statement: %v", err)
	}
	return -1
}

// ListTracks function lists all music tracks in the Tracks table.
func ListTracks() []string {
	const sql = "SELECT * FROM Tracks"
	rows, err := repo.DB.Query(sql)
	if err != nil {
		// Log error and message when failing to query tracks
		log.Printf("Failed to query tracks: %v", err)
		return nil
	}
	defer rows.Close()
	var tracks []string
	for rows.Next() {
		var id string
		var audio string
		if err := rows.Scan(&id, &audio); err == nil {
			tracks = append(tracks, id)
		} else {
			// Log error and message when failing to scan track
			log.Printf("Failed to scan track: %v", err)
		}
	}
	if err := rows.Err(); err != nil {
		// Log error and message when there is an error iterating through rows
		log.Printf("Failed to iterate through rows: %v", err)
	}
	return tracks
}

// Delete function deletes a music track with the specified ID from the Tracks table.
func Delete(id string) int64 {
	const sql = "DELETE FROM Tracks WHERE Id = ?"
	if stmt, err := repo.DB.Prepare(sql); err == nil {
		defer stmt.Close()
		if res, err := stmt.Exec(id); err == nil {
			if n, err := res.RowsAffected(); err == nil {
				return n
			} else {
				// In case of failure to read rows within SQL delete statement, log output
				log.Printf("Failed to read the given rows in statement: %v", err)
				return -1
			}
		} else {
			// In case of failure to execute SQL delete statement, log output
			log.Printf("Failed to execute SQL delete statement: %v", err)
			return -1
		}
	} else {
		// Failed to prepare SQL delete statement, log output
		log.Printf("Failed to prepare SQL delete statement: %v", err)
		return -1
	}
}
