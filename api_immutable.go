/*
 * Harbor API
 *
 * These APIs provide services for manipulating Harbor project.
 *
 * API version: 2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/antihax/optional"
	"github.com/pkg/errors"
)

// Linger please
var (
	_ context.Context
)

type ImmutableApiService service

/*
ImmutableApiService Add an immutable tag rule to current project
This endpoint add an immutable tag rule to the project
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param projectNameOrId The name or id of the project
 * @param immutableRule
 * @param optional nil or *ImmutableApiCreateImmuRuleOpts - Optional Parameters:
     * @param "XRequestId" (optional.String) -  An unique ID for the request
     * @param "XIsResourceName" (optional.Bool) -  The flag to indicate whether the parameter which supports both name and id in the path is the name of the resource. When the X-Is-Resource-Name is false and the parameter can be converted to an integer, the parameter will be as an id, otherwise, it will be as a name.


*/

type ImmutableApiCreateImmuRuleOpts struct {
	XRequestId      optional.String
	XIsResourceName optional.Bool
}

func (a *ImmutableApiService) CreateImmuRule(ctx context.Context, projectNameOrId string, immutableRule ImmutableRule, localVarOptionals *ImmutableApiCreateImmuRuleOpts) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/projects/{project_name_or_id}/immutabletagrules"
	localVarPath = strings.Replace(localVarPath, "{"+"project_name_or_id"+"}", fmt.Sprintf("%v", projectNameOrId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if localVarOptionals != nil && localVarOptionals.XRequestId.IsSet() {
		localVarHeaderParams["X-Request-Id"] = parameterToString(localVarOptionals.XRequestId.Value(), "")
	}
	if localVarOptionals != nil && localVarOptionals.XIsResourceName.IsSet() {
		localVarHeaderParams["X-Is-Resource-Name"] = parameterToString(localVarOptionals.XIsResourceName.Value(), "")
	}
	// body params
	localVarPostBody = &immutableRule
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, errors.WithStack(err)
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarHttpResponse, errors.WithStack(err)
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}

		if localVarHttpResponse.StatusCode == 400 {
			var v Errors
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHttpResponse, errors.WithStack(newErr)
			}
			newErr.model = v
			return localVarHttpResponse, errors.WithStack(newErr)
		}

		if localVarHttpResponse.StatusCode == 401 {
			var v Errors
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHttpResponse, errors.WithStack(newErr)
			}
			newErr.model = v
			return localVarHttpResponse, errors.WithStack(newErr)
		}

		if localVarHttpResponse.StatusCode == 403 {
			var v Errors
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHttpResponse, errors.WithStack(newErr)
			}
			newErr.model = v
			return localVarHttpResponse, errors.WithStack(newErr)
		}

		if localVarHttpResponse.StatusCode == 404 {
			var v Errors
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHttpResponse, errors.WithStack(newErr)
			}
			newErr.model = v
			return localVarHttpResponse, errors.WithStack(newErr)
		}

		if localVarHttpResponse.StatusCode == 500 {
			var v Errors
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHttpResponse, errors.WithStack(newErr)
			}
			newErr.model = v
			return localVarHttpResponse, errors.WithStack(newErr)
		}

		return localVarHttpResponse, errors.WithStack(newErr)
	}

	return localVarHttpResponse, nil
}

/*
ImmutableApiService Delete the immutable tag rule.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param projectNameOrId The name or id of the project
 * @param immutableRuleId The ID of the immutable rule
 * @param optional nil or *ImmutableApiDeleteImmuRuleOpts - Optional Parameters:
     * @param "XRequestId" (optional.String) -  An unique ID for the request
     * @param "XIsResourceName" (optional.Bool) -  The flag to indicate whether the parameter which supports both name and id in the path is the name of the resource. When the X-Is-Resource-Name is false and the parameter can be converted to an integer, the parameter will be as an id, otherwise, it will be as a name.


*/

type ImmutableApiDeleteImmuRuleOpts struct {
	XRequestId      optional.String
	XIsResourceName optional.Bool
}

func (a *ImmutableApiService) DeleteImmuRule(ctx context.Context, projectNameOrId string, immutableRuleId int64, localVarOptionals *ImmutableApiDeleteImmuRuleOpts) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/projects/{project_name_or_id}/immutabletagrules/{immutable_rule_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"project_name_or_id"+"}", fmt.Sprintf("%v", projectNameOrId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"immutable_rule_id"+"}", fmt.Sprintf("%v", immutableRuleId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if localVarOptionals != nil && localVarOptionals.XRequestId.IsSet() {
		localVarHeaderParams["X-Request-Id"] = parameterToString(localVarOptionals.XRequestId.Value(), "")
	}
	if localVarOptionals != nil && localVarOptionals.XIsResourceName.IsSet() {
		localVarHeaderParams["X-Is-Resource-Name"] = parameterToString(localVarOptionals.XIsResourceName.Value(), "")
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, errors.WithStack(err)
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarHttpResponse, errors.WithStack(err)
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}

		if localVarHttpResponse.StatusCode == 400 {
			var v Errors
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHttpResponse, errors.WithStack(newErr)
			}
			newErr.model = v
			return localVarHttpResponse, errors.WithStack(newErr)
		}

		if localVarHttpResponse.StatusCode == 401 {
			var v Errors
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHttpResponse, errors.WithStack(newErr)
			}
			newErr.model = v
			return localVarHttpResponse, errors.WithStack(newErr)
		}

		if localVarHttpResponse.StatusCode == 403 {
			var v Errors
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHttpResponse, errors.WithStack(newErr)
			}
			newErr.model = v
			return localVarHttpResponse, errors.WithStack(newErr)
		}

		if localVarHttpResponse.StatusCode == 500 {
			var v Errors
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHttpResponse, errors.WithStack(newErr)
			}
			newErr.model = v
			return localVarHttpResponse, errors.WithStack(newErr)
		}

		return localVarHttpResponse, errors.WithStack(newErr)
	}

	return localVarHttpResponse, nil
}

/*
ImmutableApiService List all immutable tag rules of current project
This endpoint returns the immutable tag rules of a project
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param projectNameOrId The name or id of the project
 * @param optional nil or *ImmutableApiListImmuRulesOpts - Optional Parameters:
     * @param "XRequestId" (optional.String) -  An unique ID for the request
     * @param "XIsResourceName" (optional.Bool) -  The flag to indicate whether the parameter which supports both name and id in the path is the name of the resource. When the X-Is-Resource-Name is false and the parameter can be converted to an integer, the parameter will be as an id, otherwise, it will be as a name.
     * @param "Page" (optional.Int64) -  The page number
     * @param "PageSize" (optional.Int64) -  The size of per page
     * @param "Q" (optional.String) -  Query string to query resources. Supported query patterns are \&quot;exact match(k&#x3D;v)\&quot;, \&quot;fuzzy match(k&#x3D;~v)\&quot;, \&quot;range(k&#x3D;[min~max])\&quot;, \&quot;list with union releationship(k&#x3D;{v1 v2 v3})\&quot; and \&quot;list with intersetion relationship(k&#x3D;(v1 v2 v3))\&quot;. The value of range and list can be string(enclosed by \&quot; or &#39;), integer or time(in format \&quot;2020-04-09 02:36:00\&quot;). All of these query patterns should be put in the query string \&quot;q&#x3D;xxx\&quot; and splitted by \&quot;,\&quot;. e.g. q&#x3D;k1&#x3D;v1,k2&#x3D;~v2,k3&#x3D;[min~max]
     * @param "Sort" (optional.String) -  Sort the resource list in ascending or descending order. e.g. sort by field1 in ascending orderr and field2 in descending order with \&quot;sort&#x3D;field1,-field2\&quot;

@return []ImmutableRule
*/

type ImmutableApiListImmuRulesOpts struct {
	XRequestId      optional.String
	XIsResourceName optional.Bool
	Page            optional.Int64
	PageSize        optional.Int64
	Q               optional.String
	Sort            optional.String
}

func (a *ImmutableApiService) ListImmuRules(ctx context.Context, projectNameOrId string, localVarOptionals *ImmutableApiListImmuRulesOpts) ([]ImmutableRule, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Get")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue []ImmutableRule
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/projects/{project_name_or_id}/immutabletagrules"
	localVarPath = strings.Replace(localVarPath, "{"+"project_name_or_id"+"}", fmt.Sprintf("%v", projectNameOrId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Page.IsSet() {
		localVarQueryParams.Add("page", parameterToString(localVarOptionals.Page.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.PageSize.IsSet() {
		localVarQueryParams.Add("page_size", parameterToString(localVarOptionals.PageSize.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Q.IsSet() {
		localVarQueryParams.Add("q", parameterToString(localVarOptionals.Q.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Sort.IsSet() {
		localVarQueryParams.Add("sort", parameterToString(localVarOptionals.Sort.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if localVarOptionals != nil && localVarOptionals.XRequestId.IsSet() {
		localVarHeaderParams["X-Request-Id"] = parameterToString(localVarOptionals.XRequestId.Value(), "")
	}
	if localVarOptionals != nil && localVarOptionals.XIsResourceName.IsSet() {
		localVarHeaderParams["X-Is-Resource-Name"] = parameterToString(localVarOptionals.XIsResourceName.Value(), "")
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, errors.WithStack(err)
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, errors.WithStack(err)
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, errors.WithStack(err)
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		return localVarReturnValue, localVarHttpResponse, errors.WithStack(err)
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}

		if localVarHttpResponse.StatusCode == 200 {
			var v []ImmutableRule
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, errors.WithStack(newErr)
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, errors.WithStack(newErr)
		}

		if localVarHttpResponse.StatusCode == 400 {
			var v Errors
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, errors.WithStack(newErr)
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, errors.WithStack(newErr)
		}

		if localVarHttpResponse.StatusCode == 401 {
			var v Errors
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, errors.WithStack(newErr)
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, errors.WithStack(newErr)
		}

		if localVarHttpResponse.StatusCode == 403 {
			var v Errors
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, errors.WithStack(newErr)
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, errors.WithStack(newErr)
		}

		if localVarHttpResponse.StatusCode == 500 {
			var v Errors
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, errors.WithStack(newErr)
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, errors.WithStack(newErr)
		}

		return localVarReturnValue, localVarHttpResponse, errors.WithStack(newErr)
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
ImmutableApiService Update the immutable tag rule or enable or disable the rule
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param projectNameOrId The name or id of the project
 * @param immutableRuleId The ID of the immutable rule
 * @param immutableRule
 * @param optional nil or *ImmutableApiUpdateImmuRuleOpts - Optional Parameters:
     * @param "XRequestId" (optional.String) -  An unique ID for the request
     * @param "XIsResourceName" (optional.Bool) -  The flag to indicate whether the parameter which supports both name and id in the path is the name of the resource. When the X-Is-Resource-Name is false and the parameter can be converted to an integer, the parameter will be as an id, otherwise, it will be as a name.


*/

type ImmutableApiUpdateImmuRuleOpts struct {
	XRequestId      optional.String
	XIsResourceName optional.Bool
}

func (a *ImmutableApiService) UpdateImmuRule(ctx context.Context, projectNameOrId string, immutableRuleId int64, immutableRule ImmutableRule, localVarOptionals *ImmutableApiUpdateImmuRuleOpts) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/projects/{project_name_or_id}/immutabletagrules/{immutable_rule_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"project_name_or_id"+"}", fmt.Sprintf("%v", projectNameOrId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"immutable_rule_id"+"}", fmt.Sprintf("%v", immutableRuleId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if localVarOptionals != nil && localVarOptionals.XRequestId.IsSet() {
		localVarHeaderParams["X-Request-Id"] = parameterToString(localVarOptionals.XRequestId.Value(), "")
	}
	if localVarOptionals != nil && localVarOptionals.XIsResourceName.IsSet() {
		localVarHeaderParams["X-Is-Resource-Name"] = parameterToString(localVarOptionals.XIsResourceName.Value(), "")
	}
	// body params
	localVarPostBody = &immutableRule
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, errors.WithStack(err)
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarHttpResponse, errors.WithStack(err)
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}

		if localVarHttpResponse.StatusCode == 400 {
			var v Errors
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHttpResponse, errors.WithStack(newErr)
			}
			newErr.model = v
			return localVarHttpResponse, errors.WithStack(newErr)
		}

		if localVarHttpResponse.StatusCode == 401 {
			var v Errors
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHttpResponse, errors.WithStack(newErr)
			}
			newErr.model = v
			return localVarHttpResponse, errors.WithStack(newErr)
		}

		if localVarHttpResponse.StatusCode == 403 {
			var v Errors
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHttpResponse, errors.WithStack(newErr)
			}
			newErr.model = v
			return localVarHttpResponse, errors.WithStack(newErr)
		}

		if localVarHttpResponse.StatusCode == 500 {
			var v Errors
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHttpResponse, errors.WithStack(newErr)
			}
			newErr.model = v
			return localVarHttpResponse, errors.WithStack(newErr)
		}

		return localVarHttpResponse, errors.WithStack(newErr)
	}

	return localVarHttpResponse, nil
}
