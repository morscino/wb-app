/*
Package upload handles file uploads.
*/
package upload

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"sync"

	"github.com/MastoCred-Inc/web-app/utility/environment"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
)

// Uploader represent a collection of methods for file upload operation
//go:generate mockgen -source upload.go -destination ./mock/mock_upload.go -package mock Uploader
type Uploader interface {
	UploadFiles(cleanFiles []FileInput, folder string) ([]FileAttachment, []error)
	ValidateFile(fileInput FileInput, maxFileSize int64, attachmentKinds []AttachmentKind) error
	UploadFile(fileInput FileInput, folder string) (*FileAttachment, error)
}

//Upload object
type Upload struct {
	logger zerolog.Logger
	env    *environment.Env
}

//NewUpload constructor
func NewUpload(logger zerolog.Logger, ev *environment.Env) *Uploader {
	upload := Uploader(&Upload{logger, ev})
	return &upload
}

// s3FileSession initializes S3 session
func (u *Upload) s3FileSession() (*session.Session, error) {
	accessKeyID := u.env.Get("AWS_S3_ACCESS_KEY_ID")
	secretAccessKey := u.env.Get("AWS_S3_SECRET_ACCESS_KEY")
	region := u.env.Get("AWS_S3_REGION")
	return session.NewSession(
		&aws.Config{
			Region: aws.String(region),
			Credentials: credentials.NewStaticCredentials(
				accessKeyID,
				secretAccessKey,
				"", // a token will be created when the session it's used.
			),
		})
}

// UploadFiles upload files concurrently to the S3 buckets
func (u *Upload) UploadFiles(cleanFiles []FileInput, folder string) ([]FileAttachment, []error) {
	var savedFiles = make([]FileAttachment, 0)
	var errValues = make([]error, 0)
	var wg sync.WaitGroup

	for _, f := range cleanFiles {
		wg.Add(1)
		file := f
		t := file.Type
		if file.Type == nil {
			r := ReportFileType
			t = &r
		}
		go func(fi FileInput, wg *sync.WaitGroup) {
			defer wg.Done()
			result, err := u.UploadFile(fi, folder)
			if err != nil {
				errValues = append(errValues, err)
			} else {
				savedFiles = append(savedFiles, FileAttachment{
					Kind:      result.Kind,
					URL:       result.URL,
					Size:      result.Size,
					Extension: result.Extension,
					Name:      file.Name,
					Type:      *t,
				})
			}
		}(f, &wg)
	}
	wg.Wait()

	return savedFiles, errValues
}

// UploadFile is the method that handle all file uploads,
func (u *Upload) UploadFile(fileInput FileInput, folder string) (*FileAttachment, error) {
	extension := filepath.Ext(fileInput.Name)

	// create a unique url for the file
	url := fmt.Sprintf("%s/%s%s", folder, uuid.New().String(), extension)

	sess, err := u.s3FileSession()
	if err != nil {
		u.logger.Err(err).Msgf("uploadFile:s3FileSession (%v)", err)
		return nil, err
	}

	output, err := s3manager.NewUploader(sess).Upload(&s3manager.UploadInput{
		Key:                  aws.String(url),
		Bucket:               aws.String(os.Getenv("AWS_BUCKET_NAME")),
		ACL:                  aws.String("public-read"),
		ContentType:          aws.String(http.DetectContentType(fileInput.Content)),
		ContentDisposition:   aws.String("attachment"),
		ServerSideEncryption: aws.String("AES256"),
		StorageClass:         aws.String("INTELLIGENT_TIERING"),
		Body:                 bytes.NewReader(fileInput.Content),
	})
	if err != nil {
		u.logger.Err(err).Msgf("uploadFile:Upload (%v)", err)
		return nil, err
	}

	return &FileAttachment{
		Kind:      fileInput.Kind,
		URL:       output.Location,
		Size:      fileInput.Size,
		Extension: extension,
	}, nil
}

// ValidateFile validates the file provided
func (u *Upload) ValidateFile(fileInput FileInput, maxFileSize int64, attachmentKinds []AttachmentKind) error {
	if fileInput.Size > maxFileSize {
		return fmt.Errorf("%s file should of size of %d, MB or less", fileInput.Name, maxFileSize)
	}
	allowed := false
	for _, kind := range attachmentKinds {
		if kind == AttachmentKindMap[fileInput.Kind] {
			allowed = true
			break
		}
	}
	if !allowed {
		return fmt.Errorf("file uploaded should be of the format(s): %s", aggregatedAttachmentKind(attachmentKinds))
	}

	return nil
}

// aggregatedAttachmentKind return joined, well formatted attachments
func aggregatedAttachmentKind(attachments []AttachmentKind) string {
	result := ""
	for _, attachment := range attachments {
		result += fmt.Sprintf("%s, ", attachment)
	}
	return result[:len(result)-2] // remove the comma and space for the last item
}
