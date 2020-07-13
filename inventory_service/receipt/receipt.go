package receipt

import (
	"io/ioutil"
	"path/filepath"
	"time"
)

//ReceiptDirectory stores a path to the directory where all the files are stored
var ReceiptDirectory string = filepath.Join("uploads")

//Receipt ...
type Receipt struct {
	ReceiptName string    `json:"name"`
	UploadDate  time.Time `json:"uploadDate"`
}

//GetReceipts is a function that goes thrue the list of files
//in our uploads folder
func GetReceipts() ([]Receipt, error) {
	print()
	receipts := make([]Receipt, 0)
	files, err := ioutil.ReadDir(ReceiptDirectory)
	if err != nil {
		return nil, err
	}
	for _, f := range files {
		receipts = append(receipts, Receipt{ReceiptName: f.Name(), UploadDate: f.ModTime()})
	}
	return receipts, nil
}
