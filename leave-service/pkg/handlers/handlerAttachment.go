package handlers

import (
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

const attachmentPath = "./static/attachments/"
const userName = "tester"

// Insert Claim attachment to file
// TODO: Rename this if function will be reused
func (rep *Repository) UploadLeaveAttachment(w http.ResponseWriter, r *http.Request) {
	// Check if the request method is POST
	var renameAttachments []string
	log.Println("Check Upload")
	if r.Method == "POST" {
		// Parse the form data and retrieve the uploaded file
		r.ParseMultipartForm(10 << 20) // 10 MB maximum file size
		files := r.MultipartForm.File["attachments[]"]
		for index, handler := range files {
			file, err := handler.Open()
			if err != nil {
				log.Println("Error opening file:", err)
				return
			}
			defer file.Close()

			//Rename the file
			reFName := renameAttachment(r, handler.Filename, strconv.Itoa(index+1))
			renameAttachments = append(renameAttachments, reFName)

			//Get the service name from the URL (TODO: Check whether to Add into helper function or not ?)
			sURL := strings.Split(r.URL.Path, "/")
			service := sURL[1]
			//Create the directory if not exist

			uploadPath := attachmentPath + service + "/"
			// uploadPath := filepath.Join(attachmentPath, service)
			uploadPath, err = rep.Mkdir(uploadPath)

			if err != nil {
				log.Println("Error creating directory:", err)
				return
			}

			// Create a new file on the server to save the uploaded file
			f, err := os.OpenFile(uploadPath+"/"+reFName, os.O_WRONLY|os.O_CREATE, 0666)
			if err != nil {
				log.Println("Error creating file:", err)
				return
			}
			defer f.Close()

			// Copy the contents of the uploaded file to the new file on the server
			io.Copy(f, file)

			//TODO Send a response to the client
			log.Println("File %s uploaded successfully!", handler.Filename)
		}

		fileAttachSepComma := strings.Join(renameAttachments, ",")
		var resp = jsonResponse{
			Error:   false,
			Message: "Attachment uploaded",
			Data:    fileAttachSepComma,
		}

		//response to be send back
		rep.writeJSON(w, http.StatusAccepted, resp)
	}
}

// Rename the file attachment based on the service name, username and current time stamp and original file name and  extension
func renameAttachment(r *http.Request, fName string, iteration string) string {
	sURL := strings.Split(r.URL.Path, "/")
	service := sURL[1]
	//TODO: Need to figure out how to get username or user
	//Hardcoded
	cTime := time.Now().Format("20060102150405")
	//Should ID be added to the file name?
	newFileName := service + "_" + userName + "_" + cTime + "_" + iteration + filepath.Ext(fName)
	return newFileName
}
