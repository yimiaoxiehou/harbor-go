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
	"os"
	"strings"

	"github.com/pkg/errors"

	"github.com/antihax/optional"
)

// Linger please
var (
	_ context.Context
)

type ScanDataExportApiService service

/*
ScanDataExportApiService Download the scan data export file
Download the scan data report. Default format is CSV
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param executionId Execution ID
 * @param optional nil or *ScanDataExportApiDownloadScanDataOpts - Optional Parameters:
     * @param "XRequestId" (optional.String) -  An unique ID for the request
     * @param "Format" (optional.String) -  The format of the data to be exported. e.g. CSV or PDF

@return *os.File
*/

type ScanDataExportApiDownloadScanDataOpts struct {
	XRequestId optional.String
	Format     optional.String
}

func (a *ScanDataExportApiService) DownloadScanData(ctx context.Context, executionId int32, localVarOptionals *ScanDataExportApiDownloadScanDataOpts) (*os.File, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Get")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue *os.File
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/export/cve/download/{execution_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"execution_id"+"}", fmt.Sprintf("%v", executionId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Format.IsSet() {
		localVarQueryParams.Add("format", parameterToString(localVarOptionals.Format.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"text/csv"}

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
			var v *os.File
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
ScanDataExportApiService Export scan data for selected projects
Export scan data for selected projects
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param xScanDataType The type of scan data to export
 * @param criteria The criteria for the export
 * @param optional nil or *ScanDataExportApiExportScanDataOpts - Optional Parameters:
     * @param "XRequestId" (optional.String) -  An unique ID for the request

@return ScanDataExportJob
*/

type ScanDataExportApiExportScanDataOpts struct {
	XRequestId optional.String
}

func (a *ScanDataExportApiService) ExportScanData(ctx context.Context, xScanDataType string, criteria ScanDataExportRequest, localVarOptionals *ScanDataExportApiExportScanDataOpts) (ScanDataExportJob, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Post")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue ScanDataExportJob
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/export/cve"

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
	localVarHeaderParams["X-Scan-Data-Type"] = parameterToString(xScanDataType, "")
	// body params
	localVarPostBody = &criteria
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
			var v ScanDataExportJob
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

		if localVarHttpResponse.StatusCode == 405 {
			var v Errors
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, errors.WithStack(newErr)
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, errors.WithStack(newErr)
		}

		if localVarHttpResponse.StatusCode == 409 {
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
ScanDataExportApiService Get the specific scan data export execution
Get the scan data export execution specified by ID
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param executionId Execution ID
 * @param optional nil or *ScanDataExportApiGetScanDataExportExecutionOpts - Optional Parameters:
     * @param "XRequestId" (optional.String) -  An unique ID for the request

@return ScanDataExportExecution
*/

type ScanDataExportApiGetScanDataExportExecutionOpts struct {
	XRequestId optional.String
}

func (a *ScanDataExportApiService) GetScanDataExportExecution(ctx context.Context, executionId int32, localVarOptionals *ScanDataExportApiGetScanDataExportExecutionOpts) (ScanDataExportExecution, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Get")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue ScanDataExportExecution
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/export/cve/execution/{execution_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"execution_id"+"}", fmt.Sprintf("%v", executionId), -1)

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
			var v ScanDataExportExecution
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
ScanDataExportApiService Get a list of specific scan data export execution jobs for a specified user
Get a list of specific scan data export execution jobs for a specified user
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *ScanDataExportApiGetScanDataExportExecutionListOpts - Optional Parameters:
     * @param "XRequestId" (optional.String) -  An unique ID for the request

@return ScanDataExportExecutionList
*/

type ScanDataExportApiGetScanDataExportExecutionListOpts struct {
	XRequestId optional.String
}

func (a *ScanDataExportApiService) GetScanDataExportExecutionList(ctx context.Context, localVarOptionals *ScanDataExportApiGetScanDataExportExecutionListOpts) (ScanDataExportExecutionList, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Get")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue ScanDataExportExecutionList
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/export/cve/executions"

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
			var v ScanDataExportExecutionList
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
