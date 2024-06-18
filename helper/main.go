package helper

import (
	"mime/multipart"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

func ResData(status int, message string, errorText string, data interface{}) (jsonObj any) {
	return gin.H{"code": status, "message": message, "error": errorText, "payload": data}
}

func IsImage(file *multipart.FileHeader) bool {
	switch strings.ToLower(filepath.Ext(file.Filename)) {
	case ".jpg", ".jpeg", ".jfif", ".png", ".tiff", ".tif", ".raw", ".svg", ".webp":
		return true
	default:
		return false
	}
}

// func UploadFile(fileName string, fileData []byte, path string, username string) (string, error) {
// 	session, err := session.NewSession(&aws.Config{
// 		Region:      aws.String(AWS_S3_REGION),
// 		Credentials: credentials.NewStaticCredentials(AWS_ACCESS_KEY_ID, AWS_SECRET_ACCESS_KEY, ""),
// 	})

// 	if err != nil {
// 		logrus.Fatal(err)
// 	}
// 	if username == "" {
// 		username = "IMAGE"
// 	}
// 	//fileCompress := CompressImageResource(fileData)
// 	_, err = s3.New(session).PutObject(&s3.PutObjectInput{
// 		Bucket: aws.String(AWS_S3_BUCKET),
// 		Key:    aws.String("CHATALLINONE/" + path + "/" + username + "/" + fileName),
// 		ACL:    aws.String("private"),
// 		Body:   bytes.NewReader(fileData),
// 		// ContentLength:        aws.Int64(fileSize),
// 		ContentType:          aws.String(http.DetectContentType(fileData)),
// 		ContentDisposition:   aws.String(""),
// 		ServerSideEncryption: aws.String("AES256"),
// 	})
// 	if err != nil {
// 		return "", err
// 	}

// 	urlpath := DOMAIN_URL_S3 + "/CHATALLINONE/" + path + "/" + username + "/" + fileName
// 	return urlpath, err
// }
