package folders

import "github.com/gofrs/uuid"

type FetchFolderRequest struct {
	OrgID    uuid.UUID
	Status   string
	PageSize int
}

type FetchFolderResponse struct {
	Folders []*Folder
	Count   int
}

// ADDED: returns folders within the page and tokens to other pages
// Token is 0 if there are no more pages in the direction (previous/next)
type PaginatedFetchFolderResponse struct {
	Params       *FetchFolderRequest
	TotalCount   int
	PageSize     int
	CurrentPage  int
	PreviousPage int
	NextPage     int
	FirstPage    int
	LastPage     int
	Folders      []*Folder
	NilUUID      uuid.UUID // Pass to template to check for empty OrgID search
}
