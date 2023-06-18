// package proto defines api request and response proto
package proto

type ParamToListMyArchive struct {
	ArchiveKey  string `json:"archive_key"`
	Description string `json:"description"`
}

type ParamToCreateMyNewAchive struct {
	ArchiveKey   *string `json:"archive_key"`
	ArchiveValue *string `json:"archive_value"`
	Description  *string `json:"description"`
}
