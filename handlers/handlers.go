package handlers

import (
	"html/template"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/gofrs/uuid"
	"github.com/pentest-duck/sc-interns-extension/folders"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	var webPageData folders.PaginatedFetchFolderResponse

	// Parse URL
	u, err := url.Parse(r.URL.String())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Extract search queries
	params := u.Query()
	orgID := strings.TrimSpace(params.Get("org-id"))
	status := params.Get("status")
	if status != "existing" && status != "deleted" { // set to all for invalid status
		status = "all"
	}
	pageSize, _ := strconv.Atoi(params.Get("size"))
	if pageSize < folders.MinFoldersPerPage || pageSize > folders.MaxFoldersPerPage { // set to default for invalid size
		pageSize = folders.DefaultPagesize
	}
	page, _ := strconv.Atoi(params.Get("page"))
	if page < 1 { // set to 1 for invalid page
		page = 1
	}

	// Prepare request to fetch folders
	req := &folders.FetchFolderRequest{
		OrgID:    uuid.FromStringOrNil(orgID),
		Status:   status,
		PageSize: pageSize,
	}

	// Get all filtered folders for specified OrgID and status
	res, err := folders.GetFilteredFolders(req)
	checkError(err)

	// Get folders for specific page
	webPageData = *folders.GetPage(res.Folders, page, pageSize)
	webPageData.Params = req

	// Run index.html template
	tmpl, err := template.ParseFiles("index.html")
	checkError(err)

	tmpl.Execute(w, webPageData)
}

// Log errors
func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
