# \UserApi

All URIs are relative to *https://api.satstream.io/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UserApiKeysGet**](UserApi.md#UserApiKeysGet) | **Get** /user/api-keys | Get API keys
[**UserApiKeysIdDelete**](UserApi.md#UserApiKeysIdDelete) | **Delete** /user/api-keys/{id} | Delete API key
[**UserAppsAppIdGet**](UserApi.md#UserAppsAppIdGet) | **Get** /user/apps/{app_id} | Get user app
[**UserAppsGet**](UserApi.md#UserAppsGet) | **Get** /user/apps | Get user apps
[**UserAppsPost**](UserApi.md#UserAppsPost) | **Post** /user/apps | Create user app
[**UserPaymentsCreateIntentPost**](UserApi.md#UserPaymentsCreateIntentPost) | **Post** /user/payments/create-intent | Create payment intent
[**UserRefreshTokenPost**](UserApi.md#UserRefreshTokenPost) | **Post** /user/refresh-token | Refresh token
[**UserSubscriptionsGet**](UserApi.md#UserSubscriptionsGet) | **Get** /user/subscriptions | Get active subscription


# **UserApiKeysGet**
> []TypesApiKey UserApiKeysGet(ctx, )
Get API keys

Get all API keys for the authenticated user

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]TypesApiKey**](types.ApiKey.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UserApiKeysIdDelete**
> map[string]Object UserApiKeysIdDelete(ctx, id)
Delete API key

Delete a specific API key for the authenticated user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| API key ID | 

### Return type

**map[string]Object**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UserAppsAppIdGet**
> TypesUserApp UserAppsAppIdGet(ctx, appId)
Get user app

Get a specific app for the authenticated user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appId** | **string**| App ID | 

### Return type

[**TypesUserApp**](types.UserApp.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UserAppsGet**
> []TypesUserApp UserAppsGet(ctx, )
Get user apps

Get all apps for the authenticated user

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]TypesUserApp**](types.UserApp.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UserAppsPost**
> TypesUserApp UserAppsPost(ctx, app)
Create user app

Create a new app for the authenticated user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **app** | [**RequestsCreateUserAppRequest**](RequestsCreateUserAppRequest.md)| App details | 

### Return type

[**TypesUserApp**](types.UserApp.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UserPaymentsCreateIntentPost**
> ResponsesCreatePaymentIntentResponse UserPaymentsCreateIntentPost(ctx, intent)
Create payment intent

Create a payment intent for the authenticated user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **intent** | [**RequestsCreatePaymentIntentRequest**](RequestsCreatePaymentIntentRequest.md)| Payment intent details | 

### Return type

[**ResponsesCreatePaymentIntentResponse**](responses.CreatePaymentIntentResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UserRefreshTokenPost**
> ResponsesRefreshTokenResponse UserRefreshTokenPost(ctx, )
Refresh token

Refresh the authentication token for the current user

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ResponsesRefreshTokenResponse**](responses.RefreshTokenResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UserSubscriptionsGet**
> []TypesSubscription UserSubscriptionsGet(ctx, )
Get active subscription

Get the active subscription for the authenticated user

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]TypesSubscription**](types.Subscription.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

