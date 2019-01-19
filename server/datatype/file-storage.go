package datatype

//FileStorage interface for FileStorage
type FileStorage interface {
	Put(path string, contentType string, content []byte, acl string) error
}
