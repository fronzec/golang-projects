// An example of how to use fetch to create a note
document.addEventListener('DOMContentLoaded', () => {
    const createNoteForm = document.getElementById('create-note');
    createNoteForm.addEventListener('submit', (e) => {
        e.preventDefault();
        const title = document.getElementById('title').value;
        const content = document.getElementById('content').value;
        fetch('http://localhost:8080/notes', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ title, content }),
        })
            .then((response) => response.json())
            .then((data) => {
                console.log(data);
                createNoteForm.reset();
            })
            .catch((error) => console.error('Error creating note:', error));
    });
});