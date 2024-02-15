# Toolkit

This project is about preparing modules, but also includes small projects related to using these modules. If you want to examine / run it locally, move the folder names `usage-of-toolkit` to a another folder. Afterwards, add the folder with the remaining files into the `usage-of-toolkit`. This way, you will be able to operate it without any problems.

A simple example how to create a reusable Go module with commonly used tools.

The included Tools are:

- Get a random string of length n
- Upload a file to a specified directory
- Create a directory, including all parent directories, if it does not already exist
- Create a URL safe slug from a string
- Download a static file
- Read JSON
- Write JSON
- Produce a JSON encoded error response
- Post JSON to a remote service

## Installation

```
go get -u github.com/erenyusufduran/toolkit
```

## Usage

### Get a Random String

```go
import (
	"fmt"
	"github.com/erenyusufduran/toolkit"
)

func main() {
	var tools toolkit.Tools
	s := tools.RandomString(10)
	fmt.Println(s)
}
```

### Upload a File to Specified Directory

Here is html code:

```html
<form action="http://localhost:8080/upload-one" method="POST" enctype="multipart/form-data">
    <div class="mb-3">
        <label for="fileUpload" class="form-label">Choose a file...</label>
        <input type="file" class="form-control" id="fileUpload" name="uploaded">
    </div>
    <input class="btn btn-primary" type="submit" value="Upload file">
</form>
```

and here is go code for upload one file and upload more than one f'les:

```go
import (
	"fmt"
	"log"
	"net/http"
	"github.com/erenyusufduran/toolkit"
)

func main() {
	mux := routes()

	log.Println("Starting server on port 8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}

func routes() http.Handler {
	mux := http.NewServeMux()

	mux.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("."))))
	mux.HandleFunc("/upload", uploadFiles)
	mux.HandleFunc("/upload-one", uploadFile)

	return mux
}

func uploadFiles(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	t := toolkit.Tools{
		MaxFileSize:      1024 * 1024 * 1024,
		AllowedFileTypes: []string{"image/jpeg", "image/png", "image/gif"},
	}

	files, err := t.UploadFiles(r, "./uploads")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	out := ""
	for _, item := range files {
		out += fmt.Sprintf("Uploaded %s to the uploads folder, renamed to %s\n", item.OriginalFileName, item.NewFileName)
	}

	_, _ = w.Write([]byte(out))
}

func uploadFile(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	t := toolkit.Tools{
		MaxFileSize:      1024 * 1024 * 1024,
		AllowedFileTypes: []string{"image/jpeg", "image/png", "image/gif"},
	}

	f, err := t.UploadOneFile(r, "./uploads")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, _ = w.Write([]byte(fmt.Sprintf("Uploaded 1 file, %s. to the uploads folder", f.OriginalFileName)))

}
```


### Create a Directory

```go
import "github.com/erenyusufduran/toolkit"

func main() {
	var tools toolkit.Tools

	tools.CreateDirIfNotExists("./test-dir")
}
```

### Create a Slug From a String

```go
import (
	"log"
	"github.com/erenyusufduran/toolkit"
)

func main() {
	toSlug := "NOW!!_ is the time 123"
	var tools toolkit.Tools

	slugified, err := tools.Slugify(toSlug)
	if err != nil {
		log.Println(err)
	}
	log.Println(slugified) // now-is-the-time-123
}
```

### Download a Static File

There is html code:

```html
<p>
    <a href="/download">Click here to download the file</a>
</p>
```

and here is go code for download a static file:

```go
import (
	"log"
	"net/http"
	"github.com/erenyusufduran/toolkit"
)

func main() {
	mux := routes()

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}

func routes() http.Handler {
	mux := http.NewServeMux()

	mux.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("."))))
	mux.HandleFunc("/download", downloadFile)

	return mux
}

func downloadFile(w http.ResponseWriter, r *http.Request) {
	t := toolkit.Tools{}
	t.DownloadStaticFile(w, r, "./files", "pic.jpeg", "puppy.jpeg")
}
```

### Working With JSON

HTML side:

```html
<body>
    <div class="container">
        <div class="row">
            <div class="col">
                <h1 class="mt-2">JSON Functionality</h1>
                <hr>
                <form>
                    <div class="mb-3">
                        <label for="json" class="form-label">JSON to Send:</label>
                        <textarea name="json" id="json" rows="5" style="font-family: 'Courier New', Courier, monospace;" class="form-control">
                        {
                            "action": "some action",
                            "message": "some message"
                        }
                        </textarea>
                    </div>
                    <a id="pushBtn" class="btn btn-primary">Push JSON</a>
                </form>
                <hr>
                <p><strong>Response from server:</strong></p>
                <div style="outline: 1px solid silver; padding: 2em;">
                    <pre id="response">No response from server yet...</pre>
                </div>
            </div>
        </div>
    </div>

    <script>
        let pushBtn = document.getElementById("pushBtn")
        let jsonPayload = document.getElementById("json")
        let serverResponse = document.getElementById("response")

        pushBtn.addEventListener("click", function () {
            console.log("clicked json is", jsonPayload.value)

            const payload = jsonPayload.value
            const headers = new Headers()

            const body = {
                method: "POST",
                body: payload,
                headers: headers
            }

            headers.append("Content-Type", "application/json")

			// you can change the remote-service whatever you use
            fetch("http://localhost:8081/remote-service", body).then((response) => response.json()).then((data) => {
                serverResponse.innerHTML = JSON.stringify(data, undefined, 4)
            }).catch((e) => {
                serverResponse.innerHTML = "<br><br>Error: " + error
            })
        }) 
    </script>
</body>
```

```go
import (
	"log"
	"net/http"

	"github.com/erenyusufduran/toolkit"
)

type RequestPayload struct {
	Action  string `json:"action"`
	Message string `json:"message"`
}

type ResponsePayload struct {
	StatusCode int    `json:"status_code,omitempty"`
	Message    string `json:"message"`
}

func main() {
	mux := http.NewServeMux()

	mux.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("."))))
	mux.HandleFunc("/receive-post", receivePost)
	mux.HandleFunc("/remote-service", remoteService)
	mux.HandleFunc("/simulated-service", simulatedService)

	log.Println("Starting service...")

	err := http.ListenAndServe(":8081", mux)
	if err != nil {
		log.Fatal(err)
	}
}

func receivePost(w http.ResponseWriter, r *http.Request) {
	var requestPayload RequestPayload
	var t toolkit.Tools

	err := t.ReadJSON(w, r, &requestPayload)
	if err != nil {
		t.ErrorJSON(w, err)
		return
	}

	responsePayload := ResponsePayload{
		Message: "hit the handler okay, and sending response",
	}

	err = t.WriteJSON(w, http.StatusAccepted, responsePayload)
	if err != nil {
		log.Println(err)
	}
}

func remoteService(w http.ResponseWriter, r *http.Request) {
	var requestPayload RequestPayload
	var t toolkit.Tools

	err := t.ReadJSON(w, r, &requestPayload)
	if err != nil {
		t.ErrorJSON(w, err)
		return
	}

	// call a remote service
	_, statusCode, err := t.PushJSONToRemote("http://localhost:8081/simulated-service", requestPayload)
	if err != nil {
		t.ErrorJSON(w, err)
		return
	}

	responsePayload := ResponsePayload{
		Message:    "hit the handler okay, and sending response",
		StatusCode: statusCode,
	}

	err = t.WriteJSON(w, http.StatusAccepted, responsePayload)
	if err != nil {
		log.Println(err)
	}
}

func simulatedService(w http.ResponseWriter, r *http.Request) {
	payload := ResponsePayload{
		Message: "ok",
	}

	var t toolkit.Tools
	_ = t.WriteJSON(w, http.StatusOK, payload)
}
```