### 1. N/A

1. route definition

- Url: /user/login
- Method: POST
- Request: `LoginRequest`
- Response: `LoginReply`

2. request definition



```golang
type LoginRequest struct {
	Name string `json:"name"`
	Password string `json:"password"`
}
```


3. response definition



```golang
type LoginReply struct {
	Token string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}
```

### 2. N/A

1. route definition

- Url: /user/detail
- Method: GET
- Request: `UserDetailRequest`
- Response: `UserDetailReply`

2. request definition



```golang
type UserDetailRequest struct {
	Identity string `json:"identity"`
}
```


3. response definition



```golang
type UserDetailReply struct {
	Name string `json:"name"`
	Email string `json:"email"`
}
```

### 3. N/A

1. route definition

- Url: /mail/code/send/register
- Method: POST
- Request: `MailCodeSendRequest`
- Response: `MailCodeSendReply`

2. request definition



```golang
type MailCodeSendRequest struct {
	Email string `json:"email"`
}
```


3. response definition



```golang
type MailCodeSendReply struct {
}
```

### 4. N/A

1. route definition

- Url: /user/register
- Method: POST
- Request: `UserRegisterRequest`
- Response: `UserRegisterReply`

2. request definition



```golang
type UserRegisterRequest struct {
	Name string `json:"name"`
	Password string `json:"password"`
	Email string `json:"email"`
	Code string `json:"code"`
}
```


3. response definition



```golang
type UserRegisterReply struct {
}
```

### 5. N/A

1. route definition

- Url: /share/basic/detail
- Method: GET
- Request: `ShareBasicDetailRequest`
- Response: `ShareBasicDetailReply`

2. request definition



```golang
type ShareBasicDetailRequest struct {
	Identity string `json:"identity"`
}
```


3. response definition



```golang
type ShareBasicDetailReply struct {
	RepositoryIdentity string `json:"repository_identity"`
	Name string `json:"name"`
	Ext string `json:"ext"`
	Size int64 `json:"size"`
	Path string `json:"path"`
}
```

### 6. N/A

1. route definition

- Url: /file/upload
- Method: POST
- Request: `FileUploadRequest`
- Response: `FileUploadReply`

2. request definition



```golang
type FileUploadRequest struct {
	Hash string `json:"hash,optional"`
	Name string `json:"name,optional"`
	Ext string `json:"ext,optional"`
	Size int64 `json:"size,optional"`
	Path string `json:"path,optional"`
}
```


3. response definition



```golang
type FileUploadReply struct {
	Identity string `json:"identity"`
	Ext string `json:"ext"`
	Name string `json:"name"`
}
```

### 7. N/A

1. route definition

- Url: /user/repository/save
- Method: POST
- Request: `UserRepositorySaveRequest`
- Response: `UserRepositorySaveReply`

2. request definition



```golang
type UserRepositorySaveRequest struct {
	ParentId int64 `json:"parent_id"`
	RepositoryIdentity string `json:"repositoryIdentity"`
	Ext string `json:"ext"`
	Name string `json:"name"`
}
```


3. response definition



```golang
type UserRepositorySaveReply struct {
	Identity string `json:"repository"`
}
```

### 8. N/A

1. route definition

- Url: /user/file/list
- Method: GET
- Request: `UserFileListRequest`
- Response: `UserFileListReply`

2. request definition



```golang
type UserFileListRequest struct {
	Id int `json:"id,optional"`
	Page int `json:"page,optional"`
	Size int `json:"size,optional"`
}
```


3. response definition



```golang
type UserFileListReply struct {
	List []*UserFile `json:"list"`
	Count int64 `json:"count"`
}
```

### 9. N/A

1. route definition

- Url: /user/file/name/update
- Method: POST
- Request: `UserFileNameUpdateRequest`
- Response: `UserFileNameUpdateReply`

2. request definition



```golang
type UserFileNameUpdateRequest struct {
	Identity string `json:"identity"`
	Name string `json:"name"`
}
```


3. response definition



```golang
type UserFileNameUpdateReply struct {
}
```

### 10. N/A

1. route definition

- Url: /user/folder/create
- Method: POST
- Request: `UserFloderCreateRequest`
- Response: `UserFloderCreateReply`

2. request definition



```golang
type UserFloderCreateRequest struct {
	ParentId int64 `json:"parent_id"`
	Name string `json:"name"`
}
```


3. response definition



```golang
type UserFloderCreateReply struct {
	Identity string `json:"identity"`
}
```

### 11. N/A

1. route definition

- Url: /user/file/delete
- Method: DELETE
- Request: `UserFileDeleteRequest`
- Response: `UserFileDeleteReply`

2. request definition



```golang
type UserFileDeleteRequest struct {
	Identity string `json:"identity"`
}
```


3. response definition



```golang
type UserFileDeleteReply struct {
}
```

### 12. N/A

1. route definition

- Url: /user/file/move
- Method: PUT
- Request: `UserFileMoveRequest`
- Response: `UserFileMoveReply`

2. request definition



```golang
type UserFileMoveRequest struct {
	Identity string `json:"identity"`
	ParentIdentity string `json:"parent_identity"`
}
```


3. response definition



```golang
type UserFileMoveReply struct {
}
```

### 13. N/A

1. route definition

- Url: /share/basic/create
- Method: POST
- Request: `ShareBasicCreateRequest`
- Response: `ShareBasicCreateReply`

2. request definition



```golang
type ShareBasicCreateRequest struct {
	UserRepositoryIdentity string `json:"user_repository_identity"`
	ExpiredTime int `json:"expired_time"`
}
```


3. response definition



```golang
type ShareBasicCreateReply struct {
	Identity string `json:"identity"`
}
```

### 14. N/A

1. route definition

- Url: /share/basic/save
- Method: POST
- Request: `ShareBasicSaveRequest`
- Response: `ShareBasicSaveReply`

2. request definition



```golang
type ShareBasicSaveRequest struct {
	RepositoryIdentity string `json:"repository_identity"`
	ParentId int64 `json:"parent_id"`
}
```


3. response definition



```golang
type ShareBasicSaveReply struct {
	Identity string `json:"identity"`
}
```

### 15. N/A

1. route definition

- Url: /refresh/authorization
- Method: POST
- Request: `RefreshAuthorizationRequest`
- Response: `RefreshAuthorizationReply`

2. request definition



```golang
type RefreshAuthorizationRequest struct {
}
```


3. response definition



```golang
type RefreshAuthorizationReply struct {
	Token string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}
```

### 16. N/A

1. route definition

- Url: /file/upload/prepare
- Method: POST
- Request: `FileUploadPrepareRequest`
- Response: `FileUploadPrepareReply`

2. request definition



```golang
type FileUploadPrepareRequest struct {
	Md5 string `json:"md5"`
	Name string `json:"name"`
	Ext string `json:"ext"`
}
```


3. response definition



```golang
type FileUploadPrepareReply struct {
	Identity string `json:"identity"`
	UploadId string `json:"upload_id"`
	Key string `json:"key"`
}
```

### 17. N/A

1. route definition

- Url: /file/upload/chunk
- Method: POST
- Request: `FileUploadChunkRequest`
- Response: `FileUploadChunkReply`

2. request definition



```golang
type FileUploadChunkRequest struct {
}
```


3. response definition



```golang
type FileUploadChunkReply struct {
	Etag string `json:"etag"` // MD5
}
```

### 18. N/A

1. route definition

- Url: /file/upload/chunk/complete
- Method: POST
- Request: `FileUploadChunkCompleteRequest`
- Response: `FileUploadChunkCompleteReply`

2. request definition



```golang
type FileUploadChunkCompleteRequest struct {
	Md5 string `json:"md5"`
	Name string `json:"name"`
	Ext string `json:"ext"`
	Size int64 `json:"size"`
	Key string `json:"key"`
	UploadId string `json:"upload_id"`
	CosObjects []CosObject `json:"cos_objects"`
}
```


3. response definition



```golang
type FileUploadChunkCompleteReply struct {
	Identity string `json:"identity"` // 存储池identity
}
```

