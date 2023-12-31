// Code generated by goctl. DO NOT EDIT.
package types

type LoginRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type LoginReply struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

type UserDetailRequest struct {
	Identity string `json:"identity"`
}

type UserDetailReply struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type MailCodeSendRequest struct {
	Email string `json:"email"`
}

type MailCodeSendReply struct {
}

type UserRegisterRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Code     string `json:"code"`
}

type UserRegisterReply struct {
}

type FileUploadRequest struct {
	Hash string `json:"hash,optional"`
	Name string `json:"name,optional"`
	Ext  string `json:"ext,optional"`
	Size int64  `json:"size,optional"`
	Path string `json:"path,optional"`
}

type FileUploadReply struct {
	Identity string `json:"identity"`
	Ext      string `json:"ext"`
	Name     string `json:"name"`
}

type UserRepositorySaveRequest struct {
	ParentId           int64  `json:"parent_id"`
	RepositoryIdentity string `json:"repositoryIdentity"`
	Ext                string `json:"ext"`
	Name               string `json:"name"`
}

type UserRepositorySaveReply struct {
	Identity string `json:"repository"`
}

type UserFileListRequest struct {
	Id   int `json:"id,optional"`
	Page int `json:"page,optional"`
	Size int `json:"size,optional"`
}

type UserFileListReply struct {
	List  []*UserFile `json:"list"`
	Count int64       `json:"count"`
}

type UserFile struct {
	Id                 int    `json:"id"`
	Identity           string `json:"identity"`
	RepositoryIdentity string `json:"repository_identity"`
	Name               string `json:"name"`
	Ext                string `json:"ext"`
	Path               string `json:"path"`
	Size               int64  `json:"size"`
}

type UserFileNameUpdateRequest struct {
	Identity string `json:"identity"`
	Name     string `json:"name"`
}

type UserFileNameUpdateReply struct {
}

type UserFloderCreateRequest struct {
	ParentId int64  `json:"parent_id"`
	Name     string `json:"name"`
}

type UserFloderCreateReply struct {
	Identity string `json:"identity"`
}

type UserFileDeleteRequest struct {
	Identity string `json:"identity"`
}

type UserFileDeleteReply struct {
}

type UserFileMoveRequest struct {
	Identity       string `json:"identity"`
	ParentIdentity string `json:"parent_identity"`
}

type UserFileMoveReply struct {
}

type ShareBasicCreateRequest struct {
	UserRepositoryIdentity string `json:"user_repository_identity"`
	ExpiredTime            int    `json:"expired_time"`
}

type ShareBasicCreateReply struct {
	Identity string `json:"identity"`
}

type ShareBasicDetailRequest struct {
	Identity string `json:"identity"`
}

type ShareBasicDetailReply struct {
	RepositoryIdentity string `json:"repository_identity"`
	Name               string `json:"name"`
	Ext                string `json:"ext"`
	Size               int64  `json:"size"`
	Path               string `json:"path"`
}

type ShareBasicSaveRequest struct {
	RepositoryIdentity string `json:"repository_identity"`
	ParentId           int64  `json:"parent_id"`
}

type ShareBasicSaveReply struct {
	Identity string `json:"identity"`
}

type RefreshAuthorizationRequest struct {
}

type RefreshAuthorizationReply struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

type FileUploadPrepareRequest struct {
	Md5  string `json:"md5"`
	Name string `json:"name"`
	Ext  string `json:"ext"`
}

type FileUploadPrepareReply struct {
	Identity string `json:"identity"`
	UploadId string `json:"upload_id"`
	Key      string `json:"key"`
}

type FileUploadChunkRequest struct {
}

type FileUploadChunkReply struct {
	Etag string `json:"etag"` // MD5
}

type FileUploadChunkCompleteRequest struct {
	Md5        string      `json:"md5"`
	Name       string      `json:"name"`
	Ext        string      `json:"ext"`
	Size       int64       `json:"size"`
	Key        string      `json:"key"`
	UploadId   string      `json:"upload_id"`
	CosObjects []CosObject `json:"cos_objects"`
}

type CosObject struct {
	PartNumber int    `json:"part_number"`
	Etag       string `json:"etag"`
}

type FileUploadChunkCompleteReply struct {
	Identity string `json:"identity"` // 存储池identity
}
