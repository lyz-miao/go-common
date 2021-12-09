package oss

import (
    "bytes"
    "crypto/md5"
    "encoding/base64"
    "github.com/aliyun/aliyun-oss-go-sdk/oss"
    "io/ioutil"
)

type Config struct {
    AccessKeyId     string `json:"accessKeyId"`
    AccessKeySecret string `json:"accessKeySecret"`
    Endpoint        string `json:"endpoint"`
    Bucket          string `json:"bucket"`
    Custom          bool   `json:"custom"`
}

type Client struct {
    config *Config
    Client *oss.Client
}

func New(c *Config) *Client {
    ossClient, err := oss.New(c.Endpoint, c.AccessKeyId, c.AccessKeySecret, oss.UseCname(c.Custom), oss.EnableCRC(true))
    if err != nil {
        panic(err)
    }

    return &Client{
        config: c,
        Client: ossClient,
    }
}

func (o *Client) PutObject(key string, mime string, data []byte) error {
    bucketName := o.config.Bucket
    bucket, err := o.Client.Bucket(bucketName)
    if err != nil {
        return err
    }

    h := md5.New()
    h.Write(data)
    b := base64.StdEncoding.EncodeToString(h.Sum(nil))
    option := []oss.Option{
        oss.ContentMD5(b),
        oss.ContentType(mime),
    }

    err = bucket.PutObject(key, bytes.NewReader(data), option...)
    if err != nil {
        return err
    } else {
        return nil
    }
}

func (o *Client) GetObject(key string) ([]byte, error) {
    bucketName := o.config.Bucket
    bucket, err := o.Client.Bucket(bucketName)
    if err != nil {
        return nil, err
    }
    body, err := bucket.GetObject(key)
    if err != nil {
        return nil, err
    }
    defer body.Close()

    return ioutil.ReadAll(body)
}

func (o *Client) SignURL(key string, method oss.HTTPMethod, expired int64, option ...oss.Option) (string, error) {
    bucketName := o.config.Bucket
    bucket, err := o.Client.Bucket(bucketName)
    if err != nil {
        return "", err
    }

    return bucket.SignURL(key, method, expired, option...)
}
