package cdn

import (
    "fmt"
    "library/util"
    "net/url"
    "time"
)

type Client struct {
    config *Config
}

type Config struct {
    AccessKeyId     string `json:"accessKeyId"`
    AccessKeySecret string `json:"accessKeySecret"`
    Key             string `json:"key"`
    Endpoint        string `json:"endpoint"`
}

func New(c *Config) *Client {
    return &Client{
        config: c,
    }
}

func (c *Client) Signature(path string) (string, error) {
    u, err := url.Parse(c.config.Endpoint)
    if err != nil {
        return "", err
    }

    timestamp := time.Now().Add(time.Second * 1800).Format("200601021504")
    str := fmt.Sprintf("%v%v/%v", c.config.Key, timestamp, path)
    md5 := util.SumMd5([]byte(str))
    u.Path = fmt.Sprintf("%v/%v/%v", timestamp, md5, path)

    return url.QueryUnescape(u.String())
}
