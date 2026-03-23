// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/kataras/golog"
	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/errors"
	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/util/httputils"
	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/util/jsonutils"
)

const (
	authorizationPrefix = "ZStack"

	headerKeyAuthorization = "Authorization"
	headerKeyDate          = "Date"
	headerSessionId        = "x-session-id"

	responseKeyInventories = "inventories"
	responseKeyInventory   = "inventory"
	responseKeyLocation    = "location"
	responseKeyData        = "data"
	responseKeyTotal       = "total"

	accountLoginPath = "/v1/accounts/login"
	userLoginPath    = "/v1/accounts/users/login"
	logoutPathPrefix = "/v1/accounts/sessions/"
)

type ZSHttpClient struct {
	*ZSConfig

	sessionId string //  Valid when using account or user login authentication

	httpClient *http.Client
}

func NewZSHttpClient(config *ZSConfig) *ZSHttpClient {
	// Initialize httpClient
	httpClient := httputils.GetAdaptiveTimeoutClient()
	httputils.SetClientProxyFunc(httpClient, config.proxyFunc)
	ts, _ := httpClient.Transport.(*http.Transport)
	httpClient.Transport = httputils.GetCheckTransport(ts,
		func(req *http.Request) (func(resp *http.Response), error) {
			if !config.readOnly {
				return nil, nil
			}
			if req.Method == http.MethodGet || req.Method == http.MethodHead {
				return nil, nil
			}

			url := req.URL.Path
			uri := url[strings.Index(url, config.contextPath)+len(config.contextPath):]
			// Authentication login
			if req.Method == http.MethodPut && (uri == accountLoginPath || uri == userLoginPath) {
				return nil, nil
			}
			// Authentication logout
			if req.Method == http.MethodDelete && strings.HasPrefix(uri, logoutPathPrefix) {
				return nil, nil
			}

			return nil, errors.Wrapf(errors.ErrAccountReadOnly, "%s %s", req.Method, req.URL.Path)
		})
	return &ZSHttpClient{
		ZSConfig:   config,
		httpClient: httpClient,
	}
}

// Valid when using account or user login authentication
func (cli *ZSHttpClient) LoadSession(sessionId string) {
	cli.sessionId = sessionId
}
func (cli *ZSHttpClient) unloadSession() {
	cli.sessionId = ""
}

////////////////////////////// List(Query) ///////////////////////

func (cli *ZSHttpClient) ListAll(ctx context.Context, resource string, params *param.QueryParam, retVal interface{}) error {
	result := []jsonutils.JSONObject{}
	start, limit := 0, 50
	for {
		params = params.Start(start).Limit(limit).ReplyWithCount(true)

		urlStr := cli.getListURL(resource, params.Values)
		_, resp, err := cli.httpList(ctx, urlStr)
		if err != nil {
			return err
		}

		objs, err := resp.GetArray(responseKeyInventories)
		if err != nil {
			return err
		}

		result = append(result, objs...)
		if start+limit > len(result) {
			inventories := jsonutils.Marshal(map[string][]jsonutils.JSONObject{responseKeyInventories: result})
			return inventories.Unmarshal(retVal, responseKeyInventories)
		}

		start += limit
	}
}

func (cli *ZSHttpClient) Page(ctx context.Context, resource string, params *param.QueryParam, retVal interface{}) (int, error) {
	return cli.PageWithKey(ctx, resource, responseKeyInventories, params, retVal)
}

func (cli *ZSHttpClient) PageWithKey(ctx context.Context, resource, responseKey string, params *param.QueryParam, retVal interface{}) (int, error) {
	params.ReplyWithCount(true)
	err := cli.ListWithRespKey(ctx, resource, responseKey, params, retVal)
	if err != nil {
		return 0, err
	}

	return strconv.Atoi(params.Get(responseKeyTotal))
}

func (cli *ZSHttpClient) List(ctx context.Context, resource string, params *param.QueryParam, retVal interface{}) error {
	return cli.ListWithRespKey(ctx, resource, responseKeyInventories, params, retVal)
}

func (cli *ZSHttpClient) ListWithRespKey(ctx context.Context, resource, responseKey string, params *param.QueryParam, retVal interface{}) error {
	urlStr := cli.getListURL(resource, params.Values)
	_, resp, err := cli.httpList(ctx, urlStr)
	if err != nil {
		return err
	}

	if len(responseKey) == 0 {
		return resp.Unmarshal(retVal)
	}

	if len(params.Get("replyWithCount")) > 0 {
		total, err := resp.GetString(responseKeyTotal)
		if err != nil {
			return err
		}

		params.Set(responseKeyTotal, total)
	}

	return resp.Unmarshal(retVal, responseKey)
}

func (cli *ZSHttpClient) httpList(ctx context.Context, urlStr string) (http.Header, jsonutils.JSONObject, error) {
	header, err := cli.getHeader(urlStr, http.MethodGet)
	if err != nil {
		return nil, nil, err
	}

	respHeader, resp, err := httputils.JSONRequest(cli.httpClient, ctx, httputils.GET, urlStr, header, nil, cli.debug)
	if err != nil {
		if e, ok := err.(*httputils.JSONClientError); ok {
			if strings.Contains(e.Details, "wrong accessKey signature") || strings.Contains(e.Details, "access key id") {
				return nil, nil, errors.Wrapf(errors.ErrInvalidAccessKey, "%s", err.Error())
			}
		}
		return nil, nil, errors.Wrapf(err, "%s %s", http.MethodGet, urlStr)
	}

	return respHeader, resp, nil
}

func (cli *ZSHttpClient) getListURL(resource string, params url.Values) string {
	url := cli.getURL(resource, "", "")
	if len(params) > 0 {
		url = fmt.Sprintf("%s?%s", url, params.Encode())
	}
	return url
}

////////////////////////////// Get(Retrieve Details) ///////////////////////

func (cli *ZSHttpClient) Get(ctx context.Context, resource, resourceId string, params interface{}, retVal interface{}) error {
	urlStr := cli.getGetURL(resource, resourceId, "")
	if params != nil {
		urlValues, err := param.ConvertStruct2UrlValues(params)
		if err != nil {
			return err
		}
		urlStr = fmt.Sprintf("%s?%s", urlStr, urlValues.Encode())
	}

	_, _, resp, err := cli.httpGet(ctx, urlStr, false)
	if err != nil {
		return err
	}
	inventories, err := resp.GetArray(responseKeyInventories)
	if err != nil {
		return err
	}

	if len(inventories) == 0 {
		return errors.ErrNotFound
	}

	if len(inventories) > 1 {
		return errors.ErrDuplicateId
	}

	return inventories[0].Unmarshal(retVal)
}

func (cli *ZSHttpClient) GetWithRespKey(ctx context.Context, resource, resourceId, responseKey string, params interface{}, retVal interface{}) error {
	return cli.GetWithSpec(ctx, resource, resourceId, "", responseKey, params, retVal)
}

func (cli *ZSHttpClient) GetWithSpec(ctx context.Context, resource, resourceId, spec, responseKey string, params interface{}, retVal interface{}) error {
	urlStr := cli.getGetURL(resource, resourceId, spec)
	if params != nil {
		urlValues, err := param.ConvertStruct2UrlValues(params)
		if err != nil {
			return err
		}
		urlStr = fmt.Sprintf("%s?%s", urlStr, urlValues.Encode())
	}
	_, _, resp, err := cli.httpGet(ctx, urlStr, false)
	if err != nil {
		return err
	}

	if retVal == nil {
		return nil
	}

	if len(responseKey) == 0 {
		return resp.Unmarshal(retVal)
	}

	return resp.Unmarshal(retVal, responseKey)
}

func (cli *ZSHttpClient) httpGet(ctx context.Context, urlStr string, async bool) (string, http.Header, jsonutils.JSONObject, error) {
	header := http.Header{}

	var respHeader http.Header
	var resp jsonutils.JSONObject
	startTime := time.Now()
	for time.Since(startTime) < 5*time.Minute {
		header, err := cli.getHeader(urlStr, http.MethodGet)
		if err != nil {
			return "", nil, nil, err
		}

		httpRespHeader, httpResp, err := httputils.JSONRequest(cli.httpClient, ctx, httputils.GET, urlStr, header, nil, cli.debug)
		if err != nil {
			if strings.Contains(err.Error(), "exceeded while awaiting headers") {
				time.Sleep(time.Second * 5)
				continue
			}
			return "", nil, nil, errors.Wrapf(err, "%s %s", http.MethodGet, urlStr)
		}

		respHeader = httpRespHeader
		resp = httpResp
		break
	}

	var location string
	if resp.Contains(responseKeyLocation) {
		location, _ = resp.GetString(responseKeyLocation)
		if !async {
			resultHeader, result, err := cli.httpWait(ctx, header, http.MethodGet, urlStr, jsonutils.NewDict(), location)
			return location, resultHeader, result, err
		}
	}

	return location, respHeader, resp, nil
}

func (cli *ZSHttpClient) getGetURL(resource, resourceId, spec string) string {
	return cli.getURL(resource, resourceId, spec)
}

////////////////////////////// Post(Create) ///////////////////////

func (cli *ZSHttpClient) Post(ctx context.Context, resource string, params interface{}, retVal interface{}) error {
	return cli.PostWithRespKey(ctx, resource, responseKeyInventory, params, retVal)
}

func (cli *ZSHttpClient) PostWithRespKey(ctx context.Context, resource, responseKey string, params interface{}, retVal interface{}) error {
	_, err := cli.PostWithAsync(ctx, resource, responseKey, params, retVal, false)
	return err
}

func (cli *ZSHttpClient) PostWithAsync(ctx context.Context, resource, responseKey string, params interface{}, retVal interface{}, async bool) (string, error) {
	urlStr := cli.getPostURL(resource)
	location, _, resp, err := cli.httpPost(ctx, urlStr, jsonMarshal(params), async)
	if err != nil {
		return location, err
	}

	if async || retVal == nil {
		return location, nil
	}

	if len(responseKey) == 0 {
		return location, resp.Unmarshal(retVal)
	}

	return location, resp.Unmarshal(retVal, responseKey)
}

func (cli *ZSHttpClient) httpPost(ctx context.Context, urlStr string, params jsonutils.JSONObject, async bool) (string, http.Header, jsonutils.JSONObject, error) {
	header, err := cli.getHeader(urlStr, http.MethodPost)
	if err != nil {
		return "", nil, nil, err
	}

	respHeader, resp, err := httputils.JSONRequest(cli.httpClient, ctx, httputils.POST, urlStr, header, params, cli.debug)
	if err != nil {
		return "", nil, nil, errors.Wrapf(err, "%s %s %s", http.MethodPost, urlStr, params.String())
	}

	var location string
	if resp.Contains(responseKeyLocation) {
		location, _ = resp.GetString(responseKeyLocation)
		if !async {
			resultHeader, result, err := cli.httpWait(ctx, header, http.MethodPost, urlStr, params, location)
			return location, resultHeader, result, err
		}
	}

	return location, respHeader, resp, nil
}

func (cli *ZSHttpClient) getPostURL(resource string) string {
	return cli.getURL(resource, "", "")
}

////////////////////////////// Put(Update) ///////////////////////

func (cli *ZSHttpClient) Put(ctx context.Context, resource, resourceId string, params interface{}, retVal interface{}) error {
	return cli.PutWithRespKey(ctx, resource, resourceId, responseKeyInventory, params, retVal)
}

func (cli *ZSHttpClient) PutWithRespKey(ctx context.Context, resource, resourceId, responseKey string, params interface{}, retVal interface{}) error {
	return cli.PutWithSpec(ctx, resource, resourceId, "actions", responseKey, params, retVal)
}

func (cli *ZSHttpClient) PutWithSpec(ctx context.Context, resource, resourceId, spec, responseKey string, params interface{}, retVal interface{}) error {
	_, err := cli.PutWithAsync(ctx, resource, resourceId, spec, responseKey, params, retVal, false)
	return err
}

func (cli *ZSHttpClient) PutWithAsync(ctx context.Context, resource, resourceId, spec, responseKey string, params interface{}, retVal interface{}, async bool) (string, error) {
	urlStr := cli.getPutURL(resource, resourceId, spec)
	location, _, resp, err := cli.httpPut(ctx, urlStr, jsonMarshal(params), async)
	if err != nil {
		return location, err
	}

	if async || retVal == nil {
		return location, nil
	}

	if len(responseKey) == 0 {
		fmt.Println(resp.String())
		return location, resp.Unmarshal(retVal)
	}

	return location, resp.Unmarshal(retVal, responseKey)
}

func (cli *ZSHttpClient) httpPut(ctx context.Context, urlStr string, params jsonutils.JSONObject, async bool) (string, http.Header, jsonutils.JSONObject, error) {
	header, err := cli.getHeader(urlStr, http.MethodPut)
	if err != nil {
		return "", nil, nil, err
	}

	respHeader, resp, err := httputils.JSONRequest(cli.httpClient, ctx, httputils.PUT, urlStr, header, params, cli.debug)
	if err != nil {
		return "", nil, nil, errors.Wrapf(err, "%s %s %s", http.MethodPut, urlStr, params.String())
	}

	var location string
	if resp.Contains(responseKeyLocation) {
		location, _ = resp.GetString(responseKeyLocation)
		if !async {
			resultHeader, result, err := cli.httpWait(ctx, header, http.MethodPut, urlStr, params, location)
			return location, resultHeader, result, err
		}
	}

	return location, respHeader, resp, nil
}

func (cli *ZSHttpClient) getPutURL(resource, resourceId, spec string) string {
	return cli.getURL(resource, resourceId, spec)
}

////////////////////////////// Delete(Delete) ///////////////////////

func (cli *ZSHttpClient) Delete(ctx context.Context, resource, resourceId, deleteMode string) error {
	return cli.DeleteWithSpec(ctx, resource, resourceId, "", fmt.Sprintf("deleteMode=%s", deleteMode), nil)
}

func (cli *ZSHttpClient) DeleteWithSpec(ctx context.Context, resource, resourceId, spec, paramsStr string, retVal interface{}) error {
	_, err := cli.DeleteWithAsync(ctx, resource, resourceId, spec, paramsStr, retVal, false)
	return err
}

func (cli *ZSHttpClient) DeleteWithAsync(ctx context.Context, resource, resourceId, spec, paramsStr string, retVal interface{}, async bool) (string, error) {
	urlStr := cli.getDeleteURL(resource, resourceId, spec, paramsStr)
	location, _, resp, err := cli.httpDelete(ctx, urlStr, async)
	if err != nil {
		return location, err
	}

	if async || retVal == nil {
		return location, nil
	}

	return location, resp.Unmarshal(retVal, responseKeyInventory)
}

func (cli *ZSHttpClient) httpDelete(ctx context.Context, urlStr string, async bool) (string, http.Header, jsonutils.JSONObject, error) {
	header, err := cli.getHeader(urlStr, http.MethodDelete)
	if err != nil {
		return "", nil, nil, err
	}

	respHeader, resp, err := httputils.JSONRequest(cli.httpClient, ctx, httputils.DELETE, urlStr, header, nil, cli.debug)
	if err != nil {
		return "", nil, nil, errors.Wrapf(err, "%s %s", http.MethodDelete, urlStr)
	}

	var location string
	if resp.Contains(responseKeyLocation) {
		location, _ = resp.GetString(responseKeyLocation)
		if !async {
			resultHeader, result, err := cli.httpWait(ctx, header, http.MethodDelete, urlStr, jsonutils.NewDict(), location)
			return location, resultHeader, result, err
		}
	}

	return location, respHeader, resp, nil
}

func (cli *ZSHttpClient) getDeleteURL(resource, resourceId, spec, paramsStr string) string {
	url := cli.getURL(resource, resourceId, spec)
	if len(paramsStr) > 0 {
		if strings.Contains(url, "?") {
			url = fmt.Sprintf("%s&%s", url, paramsStr)
		} else {
			url = fmt.Sprintf("%s?%s", url, paramsStr)
		}
	}
	return url
}

////////////////////////////// Common Func ///////////////////////

func (cli *ZSHttpClient) getURL(resource, resourceId, spec string) string {
	url := cli.getRequestURL(resource)
	if len(resourceId) > 0 {
		url = fmt.Sprintf("%s/%s", url, resourceId)
		if len(spec) > 0 {
			url = fmt.Sprintf("%s/%s", url, spec)
		}
	}
	return url
}

func (cli *ZSHttpClient) getRequestURL(resource string) string {
	if cli.port == 80 {
		if cli.contextPath == "" {
			return fmt.Sprintf("http://%s/%s", cli.hostname, resource)
		} else {
			return fmt.Sprintf("http://%s/%s/%s", cli.hostname, cli.contextPath, resource)
		}
	} else {
		if cli.contextPath == "" {
			return fmt.Sprintf("http://%s:%d/%s", cli.hostname, cli.port, resource)
		} else {
			return fmt.Sprintf("http://%s:%d/%s/%s", cli.hostname, cli.port, cli.contextPath, resource)
		}
	}
}

func (cli *ZSHttpClient) getHeader(url, method string) (http.Header, error) {
	if len(cli.authType) == 0 {
		return nil, errors.Errorf("no auth info is configured")
	}

	switch cli.authType {
	case AuthTypeAccessKey:
		return cli.getAccessKeyHeader(url, method)
	case AuthTypeAccount:
		return cli.getLoginHeader(url, method)
	case AuthTypeAccountUser:
		return cli.getLoginHeader(url, method)
	default:
		return nil, errors.ErrNotSupported
	}
}

func (cli *ZSHttpClient) getLoginHeader(url, method string) (http.Header, error) {
	if cli.authType != AuthTypeAccount && cli.authType != AuthTypeAccountUser {
		return nil, errors.ErrNotSupported
	}

	contextPath := fmt.Sprintf("/%s", cli.contextPath)
	uri := url[strings.Index(url, contextPath)+len(contextPath):]
	if strings.Contains(url, "?") {
		uri = uri[:strings.Index(uri, "?")]
	}

	if uri == accountLoginPath || uri == userLoginPath {
		return nil, nil
	}

	if len(cli.sessionId) == 0 {
		return nil, errors.Errorf("need to login first")
	}

	authorization := fmt.Sprintf("%s %s", "OAuth", cli.sessionId)

	header := http.Header{}
	header.Add(headerKeyAuthorization, authorization)
	header.Add(headerSessionId, cli.sessionId)
	return header, nil
}

func (cli *ZSHttpClient) getAccessKeyHeader(url, method string) (http.Header, error) {
	if cli.authType != AuthTypeAccessKey {
		return nil, errors.ErrNotSupported
	}

	if len(cli.accessKeyId) == 0 || len(cli.accessKeySecret) == 0 {
		return nil, errors.ErrInvalidAccessKey
	}

	header := http.Header{}

	//loc, _ := time.LoadLocation("Asia/Shanghai")
	//date := time.Now().In(loc).Format("Mon, 02 Jan 2006 15:04:05 MST")
	//date := time.Now().In(time.UTC).Format("Mon, 02 Jan 2006 15:04:05 MST")
	//date := time.Now().Local().Format("Mon, 02 Jan 2006 15:04:05 MST")

	date := formatDateForZStack(time.Now().Local())

	contextPath := fmt.Sprintf("/%s", cli.contextPath)
	uri := url[strings.Index(url, contextPath)+len(contextPath):]
	if strings.Contains(url, "?") {
		uri = uri[:strings.Index(uri, "?")]
	}

	hmac := hmac.New(sha1.New, []byte(cli.accessKeySecret))
	hmac.Write([]byte(fmt.Sprintf("%s\n%s\n%s", method, date, uri)))
	signature := base64.StdEncoding.EncodeToString(hmac.Sum(nil))

	authorization := fmt.Sprintf("%s %s:%s", authorizationPrefix, cli.accessKeyId, signature)

	header.Add(headerKeyAuthorization, authorization)
	header.Add(headerKeyDate, date)
	return header, nil
}

func (cli *ZSHttpClient) Wait(ctx context.Context, location, responseKey string, retVal interface{}) error {
	header, err := cli.getHeader(location, http.MethodGet)
	if err != nil {
		return err
	}
	_, result, err := cli.httpWait(ctx, header, "", "", jsonutils.NewDict(), location)
	if err != nil {
		return err
	}

	if result == nil {
		return nil
	}

	if len(responseKey) == 0 {
		return result.Unmarshal(retVal)
	}

	return result.Unmarshal(retVal, responseKey)
}

func (cli *ZSHttpClient) httpWait(ctx context.Context, header http.Header, action string, requestURL string, params jsonutils.JSONObject, location string) (http.Header, jsonutils.JSONObject, error) {

	configHost := fmt.Sprintf("%s:%d", cli.hostname, cli.port)
	if !strings.Contains(location, configHost) {
		splitRegex := "/zstack/v1/api-jobs"
		parts := strings.Split(location, splitRegex)
		if len(parts) > 1 {
			location = fmt.Sprintf("http://%s%s%s", configHost, splitRegex, parts[1])
			golog.Debugf("Replaced callback URL from %s to %s", location, location)
		}
	}

	return retryCallback(ctx, func() (http.Header, jsonutils.JSONObject, error) {

		resp, err := httputils.Request(cli.httpClient, ctx, httputils.GET, location, header, nil, cli.debug)
		if err != nil {
			return nil, nil, errors.Wrap(err, fmt.Sprintf("wait location %s", location))
		}

		if resp.StatusCode != 200 {
			if resp.StatusCode == 202 {
				httputils.CloseResponse(resp)
				return nil, nil, errors.NewJobRunningError(fmt.Sprintf("StatusCode: %d, Job Still Running", resp.StatusCode))
			}
			_, result, err := parseJSONResponseForHttpWait(resp, cli.debug)
			return nil, nil, fmt.Errorf("StatusCode: %d, Reponse: %v, Error: %v", resp.StatusCode, result, err)
		}

		return parseJSONResponseForHttpWait(resp, cli.debug)
	}, action, requestURL, params.String(), cli.retryInterval, cli.retryTimes)
}

func parseJSONResponseForHttpWait(resp *http.Response, debug bool) (http.Header, jsonutils.JSONObject, error) {
	resultHeader, result, err := httputils.ParseJSONResponse("", resp, nil, debug)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			golog.Errorf("httpWait not found %s", err.Error())
			return nil, nil, err //errors.ErrNotFound
		}

		return nil, nil, err
	}
	return resultHeader, result, nil
}

func retryCallback(ctx context.Context, fn func() (http.Header, jsonutils.JSONObject, error), action, requestURL, params string, interval, retryTimes int) (http.Header, jsonutils.JSONObject, error) {
	for {
		// Check if context is cancelled before each retry
		select {
		case <-ctx.Done():
			return nil, nil, ctx.Err()
		default:
			header, body, err := fn()
			if err == nil {
				return header, body, nil
			}

			if retryTimes == 0 {
				return header, body, err
			}
			if !errors.IsJobRunningError(err) {
				return header, body, err
			}

			golog.Debugf("Wait for job %s %s %s complete , lastest result ： %s", action, requestURL, params, err.Error())
			time.Sleep(time.Duration(interval) * time.Second)
			retryTimes--
		}
	}
}

func jsonMarshal(params interface{}) jsonutils.JSONObject {
	return jsonutils.Marshal(params)
}

var offsetToAbbr = map[int]string{
	6*3600 + 30*60: "MMT", // +0630 Myanmar Time
	7 * 3600:       "ICT", // +0700 Indochina Time
}

func formatDateForZStack(t time.Time) string {
	loc := t.Location()
	zoneName, zoneOffset := t.Zone()

	if (strings.HasPrefix(zoneName, "+") || strings.HasPrefix(zoneName, "-")) && len(zoneName) == 5 {
		if abbr, ok := offsetToAbbr[zoneOffset]; ok {
			loc = time.FixedZone(abbr, zoneOffset)
		}
	}

	t = t.In(loc)
	return t.Format("Mon, 02 Jan 2006 15:04:05 MST")
}
