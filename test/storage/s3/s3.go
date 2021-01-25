package s3

import (
	"fmt"
	"os"
	"sync"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var opt struct {
	bucket string
	key    string
}

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "s3",
		Short: "run s3 examples",
		RunE:  s3Run,
	}

	flags := cmd.Flags()
	flags.StringVarP(&opt.bucket, "bucket", "b", "", "bucket")
	flags.StringVarP(&opt.key, "key", "k", "", "key for object")

	return cmd
}

func s3Run(_ *cobra.Command, args []string) error {
	if opt.bucket == "" || opt.key == "" {
		return errors.Errorf("bucket or key is empty")
	}
	if len(args) < 1 {
		return errors.Errorf("local file is empty")
	}
	filename := args[0]

	s3Config := &aws.Config{
		Credentials:      credentials.NewStaticCredentials("minio_access_key", "minio_secret_key", ""),
		Endpoint:         aws.String("http://10.68.192.160:9000"),
		Region:           aws.String("us-east-1"),
		DisableSSL:       aws.Bool(true),
		S3ForcePathStyle: aws.Bool(true),
	}
	sess, err := session.NewSession(s3Config)
	if err != nil {
		return errors.WithMessage(err, "session.NewSession")
	}
	file, err := os.Open(filename)
	if err != nil {
		return errors.WithMessage(err, "open local file")
	}
	fileInfo, err := file.Stat()
	if err != nil {
		return errors.WithMessage(err, "stat file")
	}

	reader := &CustomReader{
		fp:      file,
		size:    fileInfo.Size(),
		signMap: map[int64]struct{}{},
	}

	uploader := s3manager.NewUploader(sess, func(u *s3manager.Uploader) {
		u.PartSize = 5 * 1024 * 1024
		u.LeavePartsOnError = true
	})

	output, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(opt.bucket),
		Key:    aws.String(opt.key),
		Body:   reader,
	})
	if err != nil {
		return errors.WithMessage(err, "s3manager Upload")
	}

	fmt.Println()
	fmt.Println(output.Location)
	return nil
}

type CustomReader struct {
	fp      *os.File
	size    int64
	read    int64
	signMap map[int64]struct{}
	mux     sync.Mutex
}

func (r *CustomReader) Read(p []byte) (int, error) {
	return r.fp.Read(p)
}

func (r *CustomReader) ReadAt(p []byte, off int64) (int, error) {
	n, err := r.fp.ReadAt(p, off)
	if err != nil {
		return n, err
	}

	r.mux.Lock()
	// Ignore the first signature call
	if _, ok := r.signMap[off]; ok {
		// Got the length have read( or means has uploaded), and you can construct your message
		r.read += int64(n)
		fmt.Printf("\rtotal read:%d    progress:%d%%", r.read, int(float32(r.read*100)/float32(r.size)))
	} else {
		r.signMap[off] = struct{}{}
	}
	r.mux.Unlock()
	return n, err
}

func (r *CustomReader) Seek(offset int64, whence int) (int64, error) {
	return r.fp.Seek(offset, whence)
}
