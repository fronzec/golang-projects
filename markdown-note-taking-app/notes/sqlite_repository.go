package notes

import (
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
)

type SqliteNoteRepository struct {
    db *sql.DB
}

func NewSqliteNoteRepository(db *sql.DB) *SqliteNoteRepository {
    return &SqliteNoteRepository{db: db}
}

func (r *SqliteNoteRepository) Create(note *Note) (int64, error) {
    res, err := r.db.Exec("INSERT INTO notes (title, content) VALUES (?, ?)", note.Title, note.Content)
    if err != nil {
        return 0, err
    }
    id, err := res.LastInsertId()
    if err != nil {
        return 0, err
    }
    return id, nil
}

func (r *SqliteNoteRepository) Get(id int64) (*Note, error) {
    note := &Note{}
    err := r.db.QueryRow("SELECT id, title, content FROM notes WHERE id = ?", id).Scan(&note.ID, &note.Title, &note.Content)
    if err != nil {
        return nil, err
    }
    return note, nil
}

func (r *SqliteNoteRepository) Update(note *Note) error {
    _, err := r.db.Exec("UPDATE notes SET title = ?, content = ? WHERE id = ?", note.Title, note.Content, note.ID)
    return err
}

func (r *SqliteNoteRepository) Delete(id int64) error {
    _, err := r.db.Exec("DELETE FROM notes WHERE id = ?", id)
    return err
}

func (r *SqliteNoteRepository) List() ([]*Note, error) {
    rows, err := r.db.Query("SELECT id, title, content FROM notes")
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    var notes []*Note
    for rows.Next() {
        note := &Note{}
        err := rows.Scan(&note.ID, &note.Title, &note.Content)
        if err != nil {
            return nil, err
        }
        notes = append(notes, note)
    }
    return notes, nil
}
