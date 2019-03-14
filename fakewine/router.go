package fakewine

/**
*
* @author Liu Weiyi
* @date 2019-03-14 16:00
 */

type Router interface {
	Routes
	Group(string, ...Handle) *GRouter
}

type Routes interface {
	ANY(string, ...Handle) Routes
	GET(string, ...Handle) Routes
	POST(string, ...Handle) Routes
	DELETE(string, ...Handle) Routes
	PUT(string, ...Handle) Routes
	PATCH(string, ...Handle) Routes
	OPTIONS(string, ...Handle) Routes
	CONNECT(string, ...Handle) Routes
	HEAD(string, ...Handle) Routes

	Handle(string, string, ...Handle) Routes
	Static(string, string) Routes
	StaticFile(string, string) Routes
}

type GRouter struct {
	Router
	root   bool
	path   string
	engine *Fengine
}

func (group *GRouter) handle(relativePath, method string, handles ...Handle) Routes {
	// TODO 计算绝对路径且添加路由
	return group.returnObj()
}

// It can use every HTTP method , but you need enter the method by yourself
func (group *GRouter) Handle(relativePath, method string, handles ...Handle) Routes {
	//TODO 校验HTTP method 是否合法
	group.handle(relativePath, method, handles...)
}

// Handle any HTTP method I can handle
func (group *GRouter) ANY(relativePath string, handles ...Handle) Routes {
	group.Handle(relativePath, "GET", handles...)
	group.Handle(relativePath, "POST", handles...)
	group.Handle(relativePath, "DELETE", handles...)
	group.Handle(relativePath, "PUT", handles...)
	group.Handle(relativePath, "PATCH", handles...)
	group.Handle(relativePath, "OPTIONS", handles...)
	group.Handle(relativePath, "HEAD", handles...)
	return group.returnObj()
}

// Handle specific HTTP method
func (group *GRouter) GET(relativePath string, handles ...Handle) Routes {
	return group.Handle(relativePath, "GET", handles...)
}

func (group *GRouter) POST(relativePath string, handles ...Handle) Routes {
	return group.Handle(relativePath, "POST", handles...)
}

func (group *GRouter) DELETE(relativePath string, handles ...Handle) Routes {
	return group.Handle(relativePath, "DELETE", handles...)

}

func (group *GRouter) PUT(relativePath string, handles ...Handle) Routes {
	return group.Handle(relativePath, "PUT", handles...)
}

func (group *GRouter) PATCH(relativePath string, handles ...Handle) Routes {
	return group.Handle(relativePath, "PATCH", handles...)
}

func (group *GRouter) OPTIONS(relativePath string, handles ...Handle) Routes {
	return group.Handle(relativePath, "OPTIONS", handles...)
}

func (group *GRouter) HEAD(relativePath string, handles ...Handle) Routes {
	return group.Handle(relativePath, "HEAD", handles...)
}

// Static Func
// Static serves files from the given file system root.
func (group *GRouter) Static(relativePath, fileRoot string) Routes {
	// TODO 打开文件夹作为 Static serves
	return group.returnObj()
}

// Static serves one file from given file name
func (group *GRouter) StaticFIle(relativePath, fileName string) Routes {
	// TODO 返回单个文件
	return group.returnObj()
}

func (group *GRouter) returnObj() Routes {
	if group.root {
		return group.engine
	}
	return group
}
