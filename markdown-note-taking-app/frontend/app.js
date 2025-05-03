async function listNotes() {
    console.log("Listing notes...")
    const response = await fetch("http://localhost:8080/api/listNotes")
    const notes = await response.json()
    const notesList = document.getElementById("notes-table-body");
    notesList.innerHTML = "";
    notes.forEach(note => {
                const noteElement = document.createElement("tr");
                noteElement.innerHTML = `<td>${note.id}</td><td>${note.title}</td><td><button onclick="redirectToReadNote(${note.id})">Read</button></td>`;
                notesList.appendChild(noteElement);
            });
        }

async function readNote() {
    const id = new URLSearchParams(window.location.search).get("id")
    if (id === null || id === undefined || id === "") {
        console.log("No ID provided")
        return
    }
    console.log("Reading note..." + id)
    const response = await fetch("http://localhost:8080/api/readNote?id=" + id)
    const note = await response.json()
    console.table(note)
    document.getElementById("note").innerHTML = note.content
}

function redirectToReadNote(id) {
    console.log("Redirecting to read note..." + id)
    window.location.href = "note.html?id=" + id
}

function redirectToListNotes() {
    console.log("Redirecting to list notes...")
    window.location.href = "index.html"
}