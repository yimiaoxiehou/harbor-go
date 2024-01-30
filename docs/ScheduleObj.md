# ScheduleObj

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cron** | **string** | A cron expression, a time-based job scheduler. | [optional] [default to null]
**NextScheduledTime** | [**time.Time**](time.Time.md) | The next time to schedule to run the job. | [optional] [default to null]
**Type_** | **string** | The schedule type. The valid values are &#39;Hourly&#39;, &#39;Daily&#39;, &#39;Weekly&#39;, &#39;Custom&#39;, &#39;Manual&#39; and &#39;None&#39;. &#39;Manual&#39; means to trigger it right away and &#39;None&#39; means to cancel the schedule.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


