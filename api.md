# Shared Params Types

- <a href="https://pkg.go.dev/github.com/satstream/satstream-go-sdk/shared">shared</a>.<a href="https://pkg.go.dev/github.com/satstream/satstream-go-sdk/shared#OrderParam">OrderParam</a>

# Shared Response Types

- <a href="https://pkg.go.dev/github.com/satstream/satstream-go-sdk/shared">shared</a>.<a href="https://pkg.go.dev/github.com/satstream/satstream-go-sdk/shared#Order">Order</a>

# Pets

Params Types:

- <a href="https://pkg.go.dev/github.com/satstream/satstream-go-sdk">satstream</a>.<a href="https://pkg.go.dev/github.com/satstream/satstream-go-sdk#PetParam">PetParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/satstream/satstream-go-sdk">satstream</a>.<a href="https://pkg.go.dev/github.com/satstream/satstream-go-sdk#APIResponse">APIResponse</a>
- <a href="https://pkg.go.dev/github.com/satstream/satstream-go-sdk">satstream</a>.<a href="https://pkg.go.dev/github.com/satstream/satstream-go-sdk#Pet">Pet</a>

Methods:

- <code title="post /pet">client.Pets.<a href="https://pkg.go.dev/github.com/satstream/satstream-go-sdk#PetService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/satstream/satstream-go-sdk">satstream</a>.<a href="https://pkg.go.dev/github.com/satstream/satstream-go-sdk#PetNewParams">PetNewParams</a>) (<a href="https://pkg.go.dev/github.com/satstream/satstream-go-sdk">satstream</a>.<a href="https://pkg.go.dev/github.com/satstream/satstream-go-sdk#Pet">Pet</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /pet/{petId}">client.Pets.<a href="https://pkg.go.dev/github.com/satstream/satstream-go-sdk#PetService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, petID <a href="https://pkg.go.dev/builtin#int64">int64</a>) (<a href="https://pkg.go.dev/github.com/satstream/satstream-go-sdk">satstream</a>.<a href="https://pkg.go.dev/github.com/satstream/satstream-go-sdk#Pet">Pet</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /pet">client.Pets.<a href="https://pkg.go.dev/github.com/satstream/satstream-go-sdk#PetService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/satstream/satstream-go-sdk">satstream</a>.<a href="https://pkg.go.dev/github.com/satstream/satstream-go-sdk#PetUpdateParams">PetUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/satstream/satstream-go-sdk">satstream</a>.<a href="https://pkg.go.dev/github.com/satstream/satstream-go-sdk#Pet">Pet</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /pet/{petId}">client.Pets.<a href="https://pkg.go.dev/github.com/satstream/satstream-go-sdk#PetService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, petID <a href="https://pkg.go.dev/builtin#int64">int64</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="get /pet/findByStatus">client.Pets.<a href="https://pkg.go.dev/github.com/satstream/satstream-go-sdk#PetService.FindByStatus">FindByStatus</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/satstream/satstream-go-sdk">satstream</a>.<a href="https://pkg.go.dev/github.com/satstream/satstream-go-sdk#PetFindByStatusParams">PetFindByStatusParams</a>) ([]<a href="https://pkg.go.dev/github.com/satstream/satstream-go-sdk">satstream</a>.<a href="https://pkg.go.dev/github.com/satstream/satstream-go-sdk#Pet">Pet</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /pet/findByTags">client.Pets.<a href="https://pkg.go.dev/github.com/satstream/satstream-go-sdk#PetService.FindByTags">FindByTags</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/satstream/satstream-go-sdk">satstream</a>.<a href="https://pkg.go.dev/github.com/satstream/satstream-go-sdk#PetFindByTagsParams">PetFindByTagsParams</a>) ([]<a href="https://pkg.go.dev/github.com/satstream/satstream-go-sdk">satstream</a>.<a href="https://pkg.go.dev/github.com/satstream/satstream-go-sdk#Pet">Pet</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /pet/{petId}">client.Pets.<a href="https://pkg.go.dev/github.com/satstream/satstream-go-sdk#PetService.UpdateByID">UpdateByID</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, petID <a href="https://pkg.go.dev/builtin#int64">int64</a>, body <a href="https://pkg.go.dev/github.com/satstream/satstream-go-sdk">satstream</a>.<a href="https://pkg.go.dev/github.com/satstream/satstream-go-sdk#PetUpdateByIDParams">PetUpdateByIDParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="post /pet/{petId}/uploadImage">client.Pets.<a href="https://pkg.go.dev/github.com/satstream/satstream-go-sdk#PetService.UploadImage">UploadImage</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, petID <a href="https://pkg.go.dev/builtin#int64">int64</a>, params <a href="https://pkg.go.dev/github.com/satstream/satstream-go-sdk">satstream</a>.<a href="https://pkg.go.dev/github.com/satstream/satstream-go-sdk#PetUploadImageParams">PetUploadImageParams</a>) (<a href="https://pkg.go.dev/github.com/satstream/satstream-go-sdk">satstream</a>.<a href="https://pkg.go.dev/github.com/satstream/satstream-go-sdk#APIResponse">APIResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Store

Response Types:

- <a href="https://pkg.go.dev/github.com/satstream/satstream-go-sdk">satstream</a>.<a href="https://pkg.go.dev/github.com/satstream/satstream-go-sdk#StoreInventoryResponse">StoreInventoryResponse</a>

Methods:

- <code title="post /store/order">client.Store.<a href="https://pkg.go.dev/github.com/satstream/satstream-go-sdk#StoreService.NewOrder">NewOrder</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/satstream/satstream-go-sdk">satstream</a>.<a href="https://pkg.go.dev/github.com/satstream/satstream-go-sdk#StoreNewOrderParams">StoreNewOrderParams</a>) (<a href="https://pkg.go.dev/github.com/satstream/satstream-go-sdk/shared">shared</a>.<a href="https://pkg.go.dev/github.com/satstream/satstream-go-sdk/shared#Order">Order</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /store/inventory">client.Store.<a href="https://pkg.go.dev/github.com/satstream/satstream-go-sdk#StoreService.Inventory">Inventory</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/satstream/satstream-go-sdk">satstream</a>.<a href="https://pkg.go.dev/github.com/satstream/satstream-go-sdk#StoreInventoryResponse">StoreInventoryResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Order

Methods:

- <code title="get /store/order/{orderId}">client.Store.Order.<a href="https://pkg.go.dev/github.com/satstream/satstream-go-sdk#StoreOrderService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orderID <a href="https://pkg.go.dev/builtin#int64">int64</a>) (<a href="https://pkg.go.dev/github.com/satstream/satstream-go-sdk/shared">shared</a>.<a href="https://pkg.go.dev/github.com/satstream/satstream-go-sdk/shared#Order">Order</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /store/order/{orderId}">client.Store.Order.<a href="https://pkg.go.dev/github.com/satstream/satstream-go-sdk#StoreOrderService.DeleteOrder">DeleteOrder</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orderID <a href="https://pkg.go.dev/builtin#int64">int64</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

# User

Params Types:

- <a href="https://pkg.go.dev/github.com/satstream/satstream-go-sdk">satstream</a>.<a href="https://pkg.go.dev/github.com/satstream/satstream-go-sdk#UserParam">UserParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/satstream/satstream-go-sdk">satstream</a>.<a href="https://pkg.go.dev/github.com/satstream/satstream-go-sdk#User">User</a>

Methods:

- <code title="post /user">client.User.<a href="https://pkg.go.dev/github.com/satstream/satstream-go-sdk#UserService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/satstream/satstream-go-sdk">satstream</a>.<a href="https://pkg.go.dev/github.com/satstream/satstream-go-sdk#UserNewParams">UserNewParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="get /user/{username}">client.User.<a href="https://pkg.go.dev/github.com/satstream/satstream-go-sdk#UserService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, username <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/satstream/satstream-go-sdk">satstream</a>.<a href="https://pkg.go.dev/github.com/satstream/satstream-go-sdk#User">User</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /user/{username}">client.User.<a href="https://pkg.go.dev/github.com/satstream/satstream-go-sdk#UserService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, existingUsername <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/satstream/satstream-go-sdk">satstream</a>.<a href="https://pkg.go.dev/github.com/satstream/satstream-go-sdk#UserUpdateParams">UserUpdateParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="delete /user/{username}">client.User.<a href="https://pkg.go.dev/github.com/satstream/satstream-go-sdk#UserService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, username <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="post /user/createWithList">client.User.<a href="https://pkg.go.dev/github.com/satstream/satstream-go-sdk#UserService.NewWithList">NewWithList</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/satstream/satstream-go-sdk">satstream</a>.<a href="https://pkg.go.dev/github.com/satstream/satstream-go-sdk#UserNewWithListParams">UserNewWithListParams</a>) (<a href="https://pkg.go.dev/github.com/satstream/satstream-go-sdk">satstream</a>.<a href="https://pkg.go.dev/github.com/satstream/satstream-go-sdk#User">User</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /user/login">client.User.<a href="https://pkg.go.dev/github.com/satstream/satstream-go-sdk#UserService.Login">Login</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/satstream/satstream-go-sdk">satstream</a>.<a href="https://pkg.go.dev/github.com/satstream/satstream-go-sdk#UserLoginParams">UserLoginParams</a>) (<a href="https://pkg.go.dev/builtin#string">string</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /user/logout">client.User.<a href="https://pkg.go.dev/github.com/satstream/satstream-go-sdk#UserService.Logout">Logout</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
