package notes

import (
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

var rootPath = "./"

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	 //1. Param input for multipart file upload
	 r.ParseMultipartForm(200 << 20) // Maximum of 200MB file allowed

	file, handler, err := r.FormFile("file")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Error parsing form file")
		return
	}

	err = saveFile(file, handler)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "Error saving file")
		return
	}

	defer file.Close()

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "File uploaded successfully")
}

func saveFile(file multipart.File, handler *multipart.FileHeader) error {
    //3. Create a temporary file to our directory
	tempFolderPath := fmt.Sprintf("%s%s", rootPath, "/tempFiles")
	tempFileName := fmt.Sprintf("upload-%s-*.%s", fileNameWithoutExtension(handler.Filename), filepath.Ext(handler.Filename))

	tempFile, err := os.CreateTemp(tempFolderPath, tempFileName)
	if err != nil {
		return err
	}
	defer tempFile.Close()

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		return err
	}
	_, err = tempFile.Write(fileBytes)
	if err != nil {
		return err
	}

	return nil

}



func fileNameWithoutExtension(filename string) string {
	return filename[:len(filename)-len(filepath.Ext(filename))]
   }