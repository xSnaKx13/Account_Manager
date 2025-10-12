package cloud

type CloudDb struct {
	URL string
}

func NewCloudDb(url string) *CloudDb {
	return &CloudDb{
		URL: url,
	}
}

func (db *CloudDb) Read() ([]byte, error) {
	return []byte{}, nil
}

func (db *CloudDb) Write(content []byte) {

}
