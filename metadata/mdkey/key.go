package mdkey

const (
	// RPCCallerNamespace indicate caller namespace
	//
	// 调用方业务的命名空间
	RPCCallerNamespace = "x-rpc-caller-namespace"
	// RPCCallerName indicate caller name
	//
	// 调用方业务的名称
	RPCCallerName = "x-rpc-caller-name"
	// RPCCallerVersion indicate caller version
	//
	// 调用方业务的版本
	RPCCallerVersion = "x-rpc-caller-version"
	// RPCCalleeNamespace indicate caller namespace
	//
	// 被调用方业务的命名空间
	RPCCalleeNamespace = "x-rpc-callee-namespace"
	// RPCCalleeName indicate caller name
	//
	// 被调用方业务的名称
	RPCCalleeName = "x-rpc-callee-name"
	// RPCCalleeVersion indicate caller version
	//
	// 被调用方业务的版本
	RPCCalleeVersion = "x-rpc-callee-version"
	// RPCRequestId indicate a unique call
	//
	// 每次请求的唯一ID
	RPCRequestId = "x-rpc-request-id"
	// RPCRequestTime is begin time of a rpc call
	//
	// 每次请求开始的时间，是机器时间，不可过度依赖这个属性
	RPCRequestTime = "x-rpc-request-time"
	// RPCMethod indicate that method called
	//
	// 每次请求的方法名
	RPCMethod = "x-rpc-method"
	// RPCAuth indicate rpc auth method & token
	//
	// 每次请求时携带的认证信息
	RPCAuthorization = "x-rpc-auth"
)

const (
	// DBType indicate database type
	//
	// 数据库操作中的数据库类型
	DBType = "x-db-type"
	// DBAddr indicate database address
	//
	// 数据库操作中的数据库地址
	DBAddr = "x-db-addr"
	// DBQuery indicate the query of operation
	//
	// 数据库操作中的数据库操作语句
	DBQuery = "x-db-query"
)
