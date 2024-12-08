package synodb

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"runtime"

	"github.com/khulnasoft/synodb/internal/flags"
)

// Collection of all synodb clients
type Client struct {
	baseUrl    *url.URL
	token      string
	cliVersion string
	Org        string

	// Single instance to be reused by all clients
	base *client

	Instances     *InstancesClient
	Databases     *DatabasesClient
	Feedback      *FeedbackClient
	Organizations *OrganizationsClient
	ApiTokens     *ApiTokensClient
	Locations     *LocationsClient
	Tokens        *TokensClient
	Users         *UsersClient
	Plans         *PlansClient
	Subscriptions *SubscriptionClient
	Billing       *BillingClient
	Groups        *GroupsClient
}

// Client struct that will be aliases by all other clients
type client struct {
	client *Client
}

func New(base *url.URL, token string, cliVersion string, org string) *Client {
	c := &Client{baseUrl: base, token: token, cliVersion: cliVersion, Org: org}

	c.base = &client{c}
	c.Instances = (*InstancesClient)(c.base)
	c.Databases = (*DatabasesClient)(c.base)
	c.Feedback = (*FeedbackClient)(c.base)
	c.Organizations = (*OrganizationsClient)(c.base)
	c.ApiTokens = (*ApiTokensClient)(c.base)
	c.Locations = (*LocationsClient)(c.base)
	c.Tokens = (*TokensClient)(c.base)
	c.Users = (*UsersClient)(c.base)
	c.Plans = (*PlansClient)(c.base)
	c.Subscriptions = (*SubscriptionClient)(c.base)
	c.Billing = (*BillingClient)(c.base)
	c.Groups = (*GroupsClient)(c.base)
	return c
}

func (t *Client) newRequest(method, urlPath string, body io.Reader) (*http.Request, error) {
	url, err := url.Parse(t.baseUrl.String())
	if err != nil {
		return nil, err
	}
	url, err = url.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(method, url.String(), body)
	if err != nil {
		return nil, err
	}
	if t.token != "" {
		req.Header.Add("Authorization", fmt.Sprint("Bearer ", t.token))
	}
	req.Header.Add("SynodbCliVersion", t.cliVersion)
	parsedCliVersion := t.cliVersion
	if parsedCliVersion != "dev" {
		parsedCliVersion = t.cliVersion[1:]
	}
	req.Header.Add("User-Agent", fmt.Sprintf("synodb/%s (%s/%s)", parsedCliVersion, runtime.GOOS, runtime.GOARCH))
	req.Header.Add("Content-Type", "application/json")
	return req, nil
}

func (t *Client) do(method, path string, body io.Reader) (*http.Response, error) {
	req, err := t.newRequest(method, path, body)
	var reqDump string
	if flags.Debug() {
		reqDump = dumpRequest(req)
	}
	if err != nil {
		return nil, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	if flags.Debug() {
		printDumps(reqDump, dumpResponse(resp))
	}
	return resp, nil
}

func printDumps(req, resp string) {
	if req != "" {
		fmt.Println(req)
	}
	if resp != "" {
		fmt.Println(resp)
	}
}

func dumpRequest(req *http.Request) string {
	dump, err := httputil.DumpRequestOut(req, true)
	if err != nil {
		return ""
	}
	return string(dump)
}

func dumpResponse(req *http.Response) string {
	dump, err := httputil.DumpResponse(req, true)
	if err != nil {
		return ""
	}
	return string(dump)
}

func (t *Client) Get(path string, body io.Reader) (*http.Response, error) {
	return t.do("GET", path, body)
}

func (t *Client) Post(path string, body io.Reader) (*http.Response, error) {
	return t.do("POST", path, body)
}

func (t *Client) Patch(path string, body io.Reader) (*http.Response, error) {
	return t.do("PATCH", path, body)
}

func (t *Client) Upload(path string, fileData *os.File) (*http.Response, error) {
	var b bytes.Buffer
	w := multipart.NewWriter(&b)
	formFile, err := w.CreateFormFile("file", fileData.Name())
	if err != nil {
		w.Close()
		return nil, err
	}

	if _, err := io.Copy(formFile, fileData); err != nil {
		w.Close()
		return nil, err
	}
	w.Close()

	req, err := t.newRequest("POST", path, &b)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", w.FormDataContentType())
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (t *Client) Delete(path string, body io.Reader) (*http.Response, error) {
	return t.do("DELETE", path, body)
}
