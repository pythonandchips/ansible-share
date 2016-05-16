package storage

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type S3Storage struct {
	bucket string
}

func NewS3Storage(bucket string) S3Storage {
	return S3Storage{bucket: bucket}
}

func (s3Storage S3Storage) session() *s3.S3 {
	config := &aws.Config{Region: aws.String("us-east-1")}
	session := session.New()
	return s3.New(session, config)
}

func (s3Storage S3Storage) Put(name, version string, data []byte) error {
	params := &s3.PutObjectInput{
		Bucket: aws.String(s3Storage.bucket),
		Key:    aws.String(filepath.Join(name, version)),
		Body:   bytes.NewReader(data),
	}
	session := s3Storage.session()
	_, err := session.PutObject(params)
	if err != nil {
		fmt.Println("Error uploading role, ", err)
	}
	return err
}

func (s3Storage S3Storage) LatestVersion(name string) (string, error) {
	result, err := s3Storage.listBucket(name)
	if err != nil {
		return "", err
	}
	if len(result.Contents) == 0 {
		return "", errors.New("Role not found in store")
	}
	var fullPath string
	currentTime := time.Date(1970, 1, 1, 1, 1, 1, 1, time.UTC)
	for _, c := range result.Contents {
		if c.LastModified.After(currentTime) {
			fullPath = *c.Key
			currentTime = *c.LastModified
		}
	}
	versionIndex := strings.Index(fullPath, "/") + 1
	version := fullPath[versionIndex:len(fullPath)]
	return version, nil
}

func (s3Storage S3Storage) List() ([]string, error) {
	params := &s3.ListObjectsV2Input{
		Bucket: aws.String(s3Storage.bucket),
	}
	session := s3Storage.session()
	result, err := session.ListObjectsV2(params)
	if err != nil {
		return []string{}, err
	}
	if len(result.Contents) == 0 {
		return []string{}, errors.New("Role not found in store")
	}
	roles := []string{}
	for _, c := range result.Contents {
		key := *c.Key
		folderIndex := strings.Index(key, "/")
		role := key[0:folderIndex]
		isIncluded := false
		for _, r := range roles {
			if role == r {
				isIncluded = true
				break
			}
		}
		if !isIncluded {
			roles = append(roles, role)
		}
	}
	return roles, nil
}

func (s3Storage S3Storage) Get(path, version string) ([]byte, error) {
	var fullPath string
	fullPath = filepath.Join(path, version)
	return s3Storage.getObject(fullPath)
}

func (s3Storage S3Storage) listBucket(prefix string) (*s3.ListObjectsV2Output, error) {
	params := &s3.ListObjectsV2Input{
		Bucket: aws.String(s3Storage.bucket),
		Prefix: aws.String(prefix),
	}
	session := s3Storage.session()
	return session.ListObjectsV2(params)
}

func (s3Storage S3Storage) getObject(path string) ([]byte, error) {
	params := &s3.GetObjectInput{
		Bucket: aws.String(s3Storage.bucket),
		Key:    aws.String(path),
	}
	session := s3Storage.session()
	object, err := session.GetObject(params)
	if err != nil {
		return []byte{}, err
	}
	return ioutil.ReadAll(object.Body)
}
