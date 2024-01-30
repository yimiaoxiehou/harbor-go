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

type LdapApiService service

/*
LdapApiService Import selected available ldap users.
This endpoint adds the selected available ldap users to harbor based on related configuration parameters from the system. System will try to guess the user email address and realname, add to harbor user information. If have errors when import user, will return the list of importing failed uid and the failed reason.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param uidList The uid listed for importing. This list will check users validity of ldap service based on configuration from the system.
 * @param optional nil or *LdapApiImportLdapUserOpts - Optional Parameters:
     * @param "XRequestId" (optional.String) -  An unique ID for the request


*/

type LdapApiImportLdapUserOpts struct {
	XRequestId optional.String
}

func (a *LdapApiService) ImportLdapUser(ctx context.Context, uidList LdapImportUsers, localVarOptionals *LdapApiImportLdapUserOpts) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/ldap/users/import"

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
	// body params
	localVarPostBody = &uidList
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
			var v []LdapFailedImportUser
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
LdapApiService Ping available ldap service.
This endpoint ping the available ldap service for test related configuration parameters.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *LdapApiPingLdapOpts - Optional Parameters:
     * @param "XRequestId" (optional.String) -  An unique ID for the request
     * @param "Ldapconf" (optional.Interface of LdapConf) -  ldap configuration. support input ldap service configuration. If it is a empty request, will load current configuration from the system.

@return LdapPingResult
*/

type LdapApiPingLdapOpts struct {
	XRequestId optional.String
	Ldapconf   optional.Interface
}

func (a *LdapApiService) PingLdap(ctx context.Context, localVarOptionals *LdapApiPingLdapOpts) (LdapPingResult, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Post")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue LdapPingResult
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/ldap/ping"

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
	// body params
	if localVarOptionals != nil && localVarOptionals.Ldapconf.IsSet() {

		localVarOptionalLdapconf, localVarOptionalLdapconfok := localVarOptionals.Ldapconf.Value().(LdapConf)
		if !localVarOptionalLdapconfok {
			return localVarReturnValue, nil, reportError("ldapconf should be LdapConf")
		}
		localVarPostBody = &localVarOptionalLdapconf
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
			var v LdapPingResult
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
LdapApiService Search available ldap groups.
This endpoint searches the available ldap groups based on related configuration parameters. support to search by groupname or groupdn.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *LdapApiSearchLdapGroupOpts - Optional Parameters:
     * @param "XRequestId" (optional.String) -  An unique ID for the request
     * @param "Groupname" (optional.String) -  Ldap group name
     * @param "Groupdn" (optional.String) -  The LDAP group DN

@return []UserGroup
*/

type LdapApiSearchLdapGroupOpts struct {
	XRequestId optional.String
	Groupname  optional.String
	Groupdn    optional.String
}

func (a *LdapApiService) SearchLdapGroup(ctx context.Context, localVarOptionals *LdapApiSearchLdapGroupOpts) ([]UserGroup, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Get")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue []UserGroup
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/ldap/groups/search"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Groupname.IsSet() {
		localVarQueryParams.Add("groupname", parameterToString(localVarOptionals.Groupname.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Groupdn.IsSet() {
		localVarQueryParams.Add("groupdn", parameterToString(localVarOptionals.Groupdn.Value(), ""))
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
			var v []UserGroup
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
LdapApiService Search available ldap users.
This endpoint searches the available ldap users based on related configuration parameters. Support searched by input ladp configuration, load configuration from the system and specific filter.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *LdapApiSearchLdapUserOpts - Optional Parameters:
     * @param "XRequestId" (optional.String) -  An unique ID for the request
     * @param "Username" (optional.String) -  Registered user ID

@return []LdapUser
*/

type LdapApiSearchLdapUserOpts struct {
	XRequestId optional.String
	Username   optional.String
}

func (a *LdapApiService) SearchLdapUser(ctx context.Context, localVarOptionals *LdapApiSearchLdapUserOpts) ([]LdapUser, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Get")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue []LdapUser
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/ldap/users/search"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Username.IsSet() {
		localVarQueryParams.Add("username", parameterToString(localVarOptionals.Username.Value(), ""))
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
			var v []LdapUser
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
