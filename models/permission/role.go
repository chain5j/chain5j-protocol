// Package permission
//
// @author: xwc1125
package permission

// RoleType 角色类型
type RoleType uint64

// 同事及同事级别以上的角色能够同步私有数据
const (
	ADMIN      RoleType = 0 // 管理员[共识中添加][全部权限]
	SUPERVISOR RoleType = 1 // 监管员[可同步全数据，不能产块、广播]
	COLLEAGUE  RoleType = 2 // 同事[可同步全数据，不能产块、广播]
	PEER       RoleType = 3 // 成员[共识中添加]，[同步无Extra数据，能产块、广播]
	OBSERVER   RoleType = 4 // 观察者[同步无Extra数据，不能产块、广播]
	OTHER      RoleType = 5 // 其他[STRICT模式，无法进入链]
)

// ChainAccessType 链准入类型
type ChainAccessType uint64

const (
	STRICT ChainAccessType = iota // 强制限制，跟权限有关，OTHER无权限进入
	LIGHT                         // 所有节点可以进入，other 只能同步txHash，
	NONE                          // 所有节点可以进入，都可以同步，节点会给其他任意节点广播
)

// DataSyncOp 数据同步类型
type DataSyncOp uint64

// TODO 如果只是同步Hash，那么state数据无法同步，因此当前不支持只同步hash
const (
	FULL DataSyncOp = iota // 全数据
	MUCH                   // 无extra数据
	LESS                   // 只有Hash
	NULL                   // 无数据
)
