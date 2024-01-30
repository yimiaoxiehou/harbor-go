# \JobserviceApi

All URIs are relative to *http://localhost/api/v2.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActionGetJobLog**](JobserviceApi.md#ActionGetJobLog) | **Get** /jobservice/jobs/{job_id}/log | Get job log by job id
[**ActionPendingJobs**](JobserviceApi.md#ActionPendingJobs) | **Put** /jobservice/queues/{job_type} | stop and clean, pause, resume pending jobs in the queue
[**GetWorkerPools**](JobserviceApi.md#GetWorkerPools) | **Get** /jobservice/pools | Get worker pools
[**GetWorkers**](JobserviceApi.md#GetWorkers) | **Get** /jobservice/pools/{pool_id}/workers | Get workers
[**ListJobQueues**](JobserviceApi.md#ListJobQueues) | **Get** /jobservice/queues | list job queues
[**StopRunningJob**](JobserviceApi.md#StopRunningJob) | **Put** /jobservice/jobs/{job_id} | Stop running job


# **ActionGetJobLog**
> string ActionGetJobLog(ctx, jobId, optional)
Get job log by job id

Get job log by job id, it is only used by administrator

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **jobId** | **string**| The id of the job. | 
 **optional** | ***JobserviceApiActionGetJobLogOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a JobserviceApiActionGetJobLogOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRequestId** | **optional.String**| An unique ID for the request | 

### Return type

**string**

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ActionPendingJobs**
> ActionPendingJobs(ctx, jobType, actionRequest, optional)
stop and clean, pause, resume pending jobs in the queue

stop and clean, pause, resume pending jobs in the queue

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **jobType** | **string**| The type of the job. &#39;all&#39; stands for all job types | 
  **actionRequest** | [**ActionRequest**](ActionRequest.md)|  | 
 **optional** | ***JobserviceApiActionPendingJobsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a JobserviceApiActionPendingJobsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xRequestId** | **optional.String**| An unique ID for the request | 

### Return type

 (empty response body)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetWorkerPools**
> []WorkerPool GetWorkerPools(ctx, optional)
Get worker pools

Get worker pools

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***JobserviceApiGetWorkerPoolsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a JobserviceApiGetWorkerPoolsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **optional.String**| An unique ID for the request | 

### Return type

[**[]WorkerPool**](WorkerPool.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetWorkers**
> []Worker GetWorkers(ctx, poolId, optional)
Get workers

Get workers in current pool

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **poolId** | **string**| The name of the pool. &#39;all&#39; stands for all pools | 
 **optional** | ***JobserviceApiGetWorkersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a JobserviceApiGetWorkersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRequestId** | **optional.String**| An unique ID for the request | 

### Return type

[**[]Worker**](Worker.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListJobQueues**
> []JobQueue ListJobQueues(ctx, optional)
list job queues

list job queue

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***JobserviceApiListJobQueuesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a JobserviceApiListJobQueuesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **optional.String**| An unique ID for the request | 

### Return type

[**[]JobQueue**](JobQueue.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StopRunningJob**
> StopRunningJob(ctx, jobId, optional)
Stop running job

Stop running job

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **jobId** | **string**| The id of the job. | 
 **optional** | ***JobserviceApiStopRunningJobOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a JobserviceApiStopRunningJobOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRequestId** | **optional.String**| An unique ID for the request | 

### Return type

 (empty response body)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

