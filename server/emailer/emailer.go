package emailer

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/FuturICT2/fin4-core/server/datatype"
)

// EmailerConfig specifies configuration options and credentials for accessing Amazon SES.
type EmailerConfig struct {
	// Endpoint is the AWS endpoint to use for requests.
	Endpoint string

	// AccessKeyID is your Amazon AWS access key ID.
	AccessKeyID string

	// SecretAccessKey is your Amazon AWS secret key.
	SecretAccessKey string
	// From is the default from email
	From string
}

// DefaultEmailer default emailer
var DefaultEmailer EmailerConfig

func Setup(cfg datatype.Config) {
	DefaultEmailer = EmailerConfig{
		Endpoint:        "https://email.us-east-1.amazonaws.com",
		AccessKeyID:     cfg.AwsSesKey,
		SecretAccessKey: cfg.AwsSesSecret,
		From:            cfg.AwsEmailFrom,
	}
}

// SendEmail sends a plain text email.
func (c *EmailerConfig) SendEmail(to, subject, body string) (string, error) {
	data := make(url.Values)
	data.Add("Action", "SendEmail")
	data.Add("Source", c.From)
	data.Add("Destination.ToAddresses.member.1", to)
	data.Add("Message.Subject.Data", subject)
	data.Add("Message.Body.Text.Data", body)
	data.Add("AWSAccessKeyId", c.AccessKeyID)
	return sesPost(data, c.Endpoint, c.AccessKeyID, c.SecretAccessKey)
}

// SendEmailHTML sends a HTML email.
func (c *EmailerConfig) SendEmailHTML(to, subject, bodyHTML, bodyText string) (string, error) {
	data := make(url.Values)
	data.Add("Action", "SendEmail")
	data.Add("Source", c.From)
	data.Add("Destination.ToAddresses.member.1", to)
	data.Add("Message.Subject.Data", subject)
	data.Add("Message.Body.Text.Data", bodyText)
	data.Add("Message.Body.Html.Data", bodyHTML)
	data.Add("AWSAccessKeyId", c.AccessKeyID)

	return sesPost(data, c.Endpoint, c.AccessKeyID, c.SecretAccessKey)
}

// SendRawEmail sends a raw email.
func (c *EmailerConfig) SendRawEmail(raw []byte) (string, error) {
	data := make(url.Values)
	data.Add("Action", "SendRawEmail")
	data.Add("RawMessage.Data", base64.StdEncoding.EncodeToString(raw))
	data.Add("AWSAccessKeyId", c.AccessKeyID)

	return sesPost(data, c.Endpoint, c.AccessKeyID, c.SecretAccessKey)
}

func sesPost(data url.Values, endpoint, accessKeyID, secretAccessKey string) (string, error) {
	body := strings.NewReader(data.Encode())
	req, err := http.NewRequest("POST", endpoint, body)
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	now := time.Now().UTC()
	// date format: "Tue, 25 May 2010 21:20:27 +0000"
	date := now.Format("Mon, 02 Jan 2006 15:04:05 -0700")
	req.Header.Set("Date", date)

	h := hmac.New(sha256.New, []uint8(secretAccessKey))
	h.Write([]uint8(date))
	signature := base64.StdEncoding.EncodeToString(h.Sum(nil))
	auth := fmt.Sprintf("AWS3-HTTPS AWSAccessKeyId=%s, Algorithm=HmacSHA256, Signature=%s", accessKeyID, signature)
	req.Header.Set("X-Amzn-Authorization", auth)

	r, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("http error: %s", err)
		return "", err
	}

	resultbody, _ := ioutil.ReadAll(r.Body)
	r.Body.Close()

	if r.StatusCode != 200 {
		log.Printf("error, status = %d", r.StatusCode)

		log.Printf("error response: %s", resultbody)
		return "", fmt.Errorf("error code %d. response: %s", r.StatusCode, resultbody)
	}

	return string(resultbody), nil
}
