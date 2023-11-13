package folders

import (
	"github.com/gofrs/uuid"
)

// Return all folders with the specified OrgID as a struct
func GetFilteredFolders(req *FetchFolderRequest) (*FetchFolderResponse, error) {

	// Fetch all folders with the requested OrgID
	foldersOfOrg, err := FetchAllFoldersByParams(req.OrgID, req.Status) //, req.DataSet)

	// Return immediately on error
	if err != nil {
		return nil, err
	}

	// Find length of foldersOfOrg
	count := len(foldersOfOrg)

	// Return resulting folders as struct
	res := &FetchFolderResponse{Folders: foldersOfOrg, Count: count}

	return res, nil
}

// Fetch all folders with the specified OrgID
func FetchAllFoldersByParams(orgID uuid.UUID, status string) ([]*Folder, error) {
	var folders []*Folder
	var err error

	// Parse JSON data from sample.json
	folders, err = GetJSONData("sample.json")
	if err != nil {
		return nil, err
	}

	// Iterate over all folders and append folders with specified OrgID to resFolders
	resFolders := []*Folder{}

	count := 1

	// Different conditionals for status being searched for
	// Switch-case outside of for loop to conserve processing
	switch status {
	case "existing":
		for _, folder := range folders {
			// Get folders with any or specified OrgID that has not been deleted
			if (orgID == uuid.Nil || folder.OrgId == orgID) && !folder.Deleted {
				folder.Number = count
				resFolders = append(resFolders, folder)
				count++
			}
		}
	case "deleted":
		for _, folder := range folders {
			// Get folders with any or specified OrgID that has been deleted
			if (orgID == uuid.Nil || folder.OrgId == orgID) && folder.Deleted {
				folder.Number = count
				resFolders = append(resFolders, folder)
				count++
			}
		}
	default:
		for _, folder := range folders {
			// Get folders with any or specified OrgID (any status)
			if orgID == uuid.Nil || folder.OrgId == orgID {
				folder.Number = count
				resFolders = append(resFolders, folder)
				count++
			}
		}
	}

	return resFolders, nil
}
