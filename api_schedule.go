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

type ScheduleApiService service

/*
ScheduleApiService
Get scheduler paused status
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param jobType The type of the job. &#39;all&#39; stands for all job types, current only support query with all
 * @param optional nil or *ScheduleApiGetSchedulePausedOpts - Optional Parameters:
     * @param "XRequestId" (optional.String) -  An unique ID for the request

@return SchedulerStatus
*/

type ScheduleApiGetSchedulePausedOpts struct {
	XRequestId optional.String
}

func (a *ScheduleApiService) GetSchedulePaused(ctx context.Context, jobType string, localVarOptionals *ScheduleApiGetSchedulePausedOpts) (SchedulerStatus, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Get")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue SchedulerStatus
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/schedules/{job_type}/paused"
	localVarPath = strings.Replace(localVarPath, "{"+"job_type"+"}", fmt.Sprintf("%v", jobType), -1)

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
			var v SchedulerStatus
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

		if localVarHttpResponse.StatusCode == 404 {
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
ScheduleApiService
List schedules
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *ScheduleApiListSchedulesOpts - Optional Parameters:
     * @param "XRequestId" (optional.String) -  An unique ID for the request
     * @param "Page" (optional.Int64) -  The page number
     * @param "PageSize" (optional.Int64) -  The size of per page

@return []ScheduleTask
*/

type ScheduleApiListSchedulesOpts struct {
	XRequestId optional.String
	Page       optional.Int64
	PageSize   optional.Int64
}

func (a *ScheduleApiService) ListSchedules(ctx context.Context, localVarOptionals *ScheduleApiListSchedulesOpts) ([]ScheduleTask, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Get")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue []ScheduleTask
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/schedules"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Page.IsSet() {
		localVarQueryParams.Add("page", parameterToString(localVarOptionals.Page.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.PageSize.IsSet() {
		localVarQueryParams.Add("page_size", parameterToString(localVarOptionals.PageSize.Value(), ""))
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
			var v []ScheduleTask
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

		if localVarHttpResponse.StatusCode == 404 {
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
