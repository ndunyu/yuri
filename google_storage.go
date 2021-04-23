package yuri

///disclaimer::
///Do not use the this storage class
///I only created it with my use case in mind
import (
	"context"
	"fmt"
	"io"
	"os"
	"time"

	"cloud.google.com/go/storage"
	"google.golang.org/api/option"
)

func InitGoogleStorage(path string) (*storage.Client, error) {
	opt := option.WithCredentialsFile(path)
	ctx := context.Background()
	client, err := storage.NewClient(ctx, opt)
	return client, err

}

type GoogleStorage struct {
	Client    *storage.Client
	ProjectId string
}

func CreateBucket(Client *storage.Client, projectId, bucketName string) error {
	ctx := context.Background()
	bucket := Client.Bucket(bucketName)
	ctx, cancel := context.WithTimeout(ctx, time.Second*20)
	defer cancel()
	if err := bucket.Create(ctx, projectId, nil); err != nil {
		return err
	}
	return nil

}

func UploadFile(Client *storage.Client, w io.Writer, bucket, objectName, filename string) error {
	ctx := context.Background()
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	ctx, cancel := context.WithTimeout(ctx, time.Second*50)
	defer cancel()
	wc := Client.Bucket(bucket).Object(objectName).NewWriter(ctx)
	if _, err = io.Copy(wc, f); err != nil {
		return fmt.Errorf("io.Copy: %v", err)
	}
	if err := wc.Close(); err != nil {
		return fmt.Errorf("Writer.Close: %v", err)
	}
	fmt.Fprintf(w, "Blob %v uploaded.\n", objectName)
	return nil
}
