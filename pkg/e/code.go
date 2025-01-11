package e

// 存放状态码信息

const (
	Success       = 200
	Error         = 500
	InvalidParams = 400

	//uer相关错误码

	ErrorExistUser             = 30001
	ErrorFailEncryption        = 30002
	ErrorExistUserNotFound     = 30003
	ErrorNotCompare            = 30004
	ErrorAuthToken             = 30005
	ErrorAuthCheckTokenTimeout = 30006
	ErrorUploadFail            = 30007
	ErrorSendEmailFail         = 30008

	//product  模块错误

	ErrorProductImgUpload = 40000

	// Favorite 收藏夹错误

	ErrorFavoriteExist = 50000
)
