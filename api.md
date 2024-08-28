# Shared Params Types

- <a href="https://pkg.go.dev/github.com/stainless-sdks/satstream-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/satstream-go/shared#OrderParam">OrderParam</a>

# Shared Response Types

- <a href="https://pkg.go.dev/github.com/stainless-sdks/satstream-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/satstream-go/shared#Order">Order</a>

# Pets

Params Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/satstream-go">satstream</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/satstream-go#PetParam">PetParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/satstream-go">satstream</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/satstream-go#APIResponse">APIResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/satstream-go">satstream</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/satstream-go#Pet">Pet</a>

Methods:

- <code title="post /pet">client.Pets.<a href="https://pkg.go.dev/github.com/stainless-sdks/satstream-go#PetService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/satstream-go">satstream</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/satstream-go#PetNewParams">PetNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/satstream-go">satstream</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/satstream-go#Pet">Pet</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /pet/{petId}">client.Pets.<a href="https://pkg.go.dev/github.com/stainless-sdks/satstream-go#PetService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, petID <a href="https://pkg.go.dev/builtin#int64">int64</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/satstream-go">satstream</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/satstream-go#Pet">Pet</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /pet">client.Pets.<a href="https://pkg.go.dev/github.com/stainless-sdks/satstream-go#PetService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/satstream-go">satstream</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/satstream-go#PetUpdateParams">PetUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/satstream-go">satstream</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/satstream-go#Pet">Pet</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /pet/{petId}">client.Pets.<a href="https://pkg.go.dev/github.com/stainless-sdks/satstream-go#PetService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, petID <a href="https://pkg.go.dev/builtin#int64">int64</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="get /pet/findByStatus">client.Pets.<a href="https://pkg.go.dev/github.com/stainless-sdks/satstream-go#PetService.FindByStatus">FindByStatus</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/satstream-go">satstream</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/satstream-go#PetFindByStatusParams">PetFindByStatusParams</a>) ([]<a href="https://pkg.go.dev/github.com/stainless-sdks/satstream-go">satstream</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/satstream-go#Pet">Pet</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /pet/findByTags">client.Pets.<a href="https://pkg.go.dev/github.com/stainless-sdks/satstream-go#PetService.FindByTags">FindByTags</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/satstream-go">satstream</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/satstream-go#PetFindByTagsParams">PetFindByTagsParams</a>) ([]<a href="https://pkg.go.dev/github.com/stainless-sdks/satstream-go">satstream</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/satstream-go#Pet">Pet</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /pet/{petId}">client.Pets.<a href="https://pkg.go.dev/github.com/stainless-sdks/satstream-go#PetService.UpdateByID">UpdateByID</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, petID <a href="https://pkg.go.dev/builtin#int64">int64</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/satstream-go">satstream</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/satstream-go#PetUpdateByIDParams">PetUpdateByIDParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="post /pet/{petId}/uploadImage">client.Pets.<a href="https://pkg.go.dev/github.com/stainless-sdks/satstream-go#PetService.UploadImage">UploadImage</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, petID <a href="https://pkg.go.dev/builtin#int64">int64</a>, params <a href="https://pkg.go.dev/github.com/stainless-sdks/satstream-go">satstream</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/satstream-go#PetUploadImageParams">PetUploadImageParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/satstream-go">satstream</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/satstream-go#APIResponse">APIResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Store

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/satstream-go">satstream</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/satstream-go#StoreInventoryResponse">StoreInventoryResponse</a>

Methods:

- <code title="post /store/order">client.Store.<a href="https://pkg.go.dev/github.com/stainless-sdks/satstream-go#StoreService.NewOrder">NewOrder</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/satstream-go">satstream</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/satstream-go#StoreNewOrderParams">StoreNewOrderParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/satstream-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/satstream-go/shared#Order">Order</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /store/inventory">client.Store.<a href="https://pkg.go.dev/github.com/stainless-sdks/satstream-go#StoreService.Inventory">Inventory</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/satstream-go">satstream</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/satstream-go#StoreInventoryResponse">StoreInventoryResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Order

Methods:

- <code title="get /store/order/{orderId}">client.Store.Order.<a href="https://pkg.go.dev/github.com/stainless-sdks/satstream-go#StoreOrderService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orderID <a href="https://pkg.go.dev/builtin#int64">int64</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/satstream-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/satstream-go/shared#Order">Order</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /store/order/{orderId}">client.Store.Order.<a href="https://pkg.go.dev/github.com/stainless-sdks/satstream-go#StoreOrderService.DeleteOrder">DeleteOrder</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orderID <a href="https://pkg.go.dev/builtin#int64">int64</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

# User

Params Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/satstream-go">satstream</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/satstream-go#UserParam">UserParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/satstream-go">satstream</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/satstream-go#User">User</a>

Methods:

- <code title="post /user">client.User.<a href="https://pkg.go.dev/github.com/stainless-sdks/satstream-go#UserService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/satstream-go">satstream</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/satstream-go#UserNewParams">UserNewParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="get /user/{username}">client.User.<a href="https://pkg.go.dev/github.com/stainless-sdks/satstream-go#UserService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, username <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/satstream-go">satstream</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/satstream-go#User">User</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /user/{username}">client.User.<a href="https://pkg.go.dev/github.com/stainless-sdks/satstream-go#UserService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, existingUsername <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/satstream-go">satstream</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/satstream-go#UserUpdateParams">UserUpdateParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="delete /user/{username}">client.User.<a href="https://pkg.go.dev/github.com/stainless-sdks/satstream-go#UserService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, username <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="post /user/createWithList">client.User.<a href="https://pkg.go.dev/github.com/stainless-sdks/satstream-go#UserService.NewWithList">NewWithList</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/satstream-go">satstream</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/satstream-go#UserNewWithListParams">UserNewWithListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/satstream-go">satstream</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/satstream-go#User">User</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /user/login">client.User.<a href="https://pkg.go.dev/github.com/stainless-sdks/satstream-go#UserService.Login">Login</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/satstream-go">satstream</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/satstream-go#UserLoginParams">UserLoginParams</a>) (<a href="https://pkg.go.dev/builtin#string">string</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /user/logout">client.User.<a href="https://pkg.go.dev/github.com/stainless-sdks/satstream-go#UserService.Logout">Logout</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
