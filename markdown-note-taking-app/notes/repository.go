package notes

type NoteRepository interface {
    Create(note *Note) (int64, error)
    Get(id int64) (*Note, error)
    Update(note *Note) error
    Delete(id int64) error
    List() ([]*Note, error)
}
