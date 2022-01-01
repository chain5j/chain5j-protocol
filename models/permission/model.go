// Package permission
//
// @author: xwc1125
package permission

// DataPermissionOp 交易数据权限
type DataPermissionOp uint64

const (
	AddOp    DataPermissionOp = iota // 增
	DelOp                            // 删
	UpdateOp                         // 改
	QueryOp                          // 查
)

// DataPermissionOpData 交易数据权限
type DataPermissionOpData struct {
	Addr     string           // 地址
	Name     string           // 名称
	Height   uint64           // 有效期截止高度，0无限制
	RoleType RoleType         // 角色
	Opt      DataPermissionOp // 操作
}
