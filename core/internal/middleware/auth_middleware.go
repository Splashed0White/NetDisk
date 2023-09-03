// 声明一个名为middleware的包
package middleware

// 导入需要的包
import (
	"NetDisk/core/utils" // 一个自定义的包，用于处理token相关的操作
	"net/http"           // 一个标准库包，用于处理HTTP请求和响应
)

// 定义一个名为AuthMiddleware的结构体，没有任何字段
type AuthMiddleware struct {
}

// 定义一个名为NewAuthMiddleware的函数，返回一个指向AuthMiddleware结构体的指针
func NewAuthMiddleware() *AuthMiddleware {
	return &AuthMiddleware{} // 创建一个空的AuthMiddleware结构体，并返回其地址
}

// 定义一个名为Handle的方法，属于AuthMiddleware结构体，接受一个http.HandlerFunc类型的参数，返回一个http.HandlerFunc类型的值
func (m *AuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) { // 返回一个匿名函数，该函数符合http.HandlerFunc类型，接受两个参数：w是用于写入响应的对象，r是用于读取请求的对象
		// 这是一个待办事项，表示需要实现中间件的功能，完成后删除这一行

		// Passthrough to next handler if need
		// 这是一个注释，表示如果需要的话，可以直接调用next函数，将请求传递给下一个处理器

		auth := r.Header.Get("Authorization") // 从请求头中获取Authorization字段的值，赋给auth变量
		if auth == "" {                       // 如果auth变量为空字符串，表示没有提供认证信息
			w.WriteHeader(http.StatusUnauthorized) // 向响应头中写入状态码为401（未授权）的信息
			w.Write([]byte("UnAuthorized"))        // 向响应体中写入字节切片（字符串）"UnAuthorized"
			return                                 // 结束函数执行，不再调用next函数
		}
		uc, err := utils.AnalyzeToken(auth) // 调用utils包中的AnalyzeToken函数，传入auth变量作为参数，返回两个值：uc是用户信息的结构体，err是错误信息
		if err != nil {                     // 如果err变量不为空，表示分析token过程中出现了错误
			w.WriteHeader(http.StatusUnauthorized) // 同样向响应头中写入状态码为401（未授权）的信息
			w.Write([]byte(err.Error()))           // 向响应体中写入字节切片（字符串），内容是err变量的错误信息
			return                                 // 结束函数执行，不再调用next函数
		}
		r.Header.Set("UserId", uc.ID)             // 将用户信息结构体中的ID字段的值设置到请求头中的UserId字段中
		r.Header.Set("UserIdentity", uc.Identity) // 将用户信息结构体中的Identity字段的值设置到请求头中的UserIdentity字段中
		r.Header.Set("UserName", uc.Name)         // 将用户信息结构体中的Name字段的值设置到请求头中的UserName字段中

		next(w, r) // 调用next函数，将经过修改后的请求和响应对象传递给下一个处理器
	}
}
