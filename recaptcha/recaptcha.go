package recaptcha

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

type (
	recaptcha struct {
		config *Config
	}

	Config struct {
		Key string `json:"key"`
	}

	recaptchaRequest struct {
		Secret   string  `json:"secret"`
		Response string  `json:"response"`
		RemoteIP *string `json:"remoteip"`
	}

	recaptchaResponse struct {
		Success     bool      `json:"success"`
		ChallengeTs *string   `json:"challenge_ts,omitempty"`
		Hostname    *string   `json:"hostname,omitempty"`
		Score       *float64  `json:"score,omitempty"`
		Action      *string   `json:"action,omitempty"`
		ErrorCodes  *[]string `json:"error-codes,omitempty"`
	}
)

func NewClient(config *Config) (*recaptcha, error) {
	if config == nil {
		return nil, nil
	}

	return &recaptcha{
		config: config,
	}, nil
}

func (r *recaptcha) Verify(token string, ip string) (*recaptchaResponse, error) {
	apiUrl := "https://www.recaptcha.net/recaptcha/api/siteverify"
	key := r.config.Key
	postParam := url.Values{
		"secret":   {key},
		"response": {token},
		"remoteip": {ip},
	}
	resp, err := http.PostForm(apiUrl, postParam)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	response := &recaptchaResponse{}
	if err = json.Unmarshal(buf, &response); err != nil {
		return nil, err
	}

	return response, nil
}
