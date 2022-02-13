package Sql

/* NodeInfo接口 */
type NodeInfo interface {
	Insert(string) error                                        // 添加新节点
	Update(string) error                                        // 根据json更新节点信息
	Delete(string) error                                        // 根据id删除节点
	Get(string, string, string, string) (string, error)         // 模糊聚合查询
	GetAbsolute(string, string, string, string) (string, error) // 绝对匹配聚合查询
	GetToJson(string) (string, error)                           // 查询指定id的结果到并将结果转换为json格式
	ReadByLocate(string) (string, error)                        // 读取符合地理位置的节点信息并返回对应json格式信息
	ReadByID(string) (string, error)                            // 读取ID为id的某个节点信息并返回对应json格式信息
	GetIP(string) (string, error)                               // 读取ID为id的某个节点的IP地址
	GetRSAPublic(string) (string, error)                        // 根据ID获取节点的RSA公钥
	GetRSAPrivate(string) (string, error)                       // 根据ID获取节点的RSA私钥
	GetByID(string) (*NodeInformation, error)                   // 根据ID获取节点信息对象
	GetByType(int) (*[]NodeInformation, error)                  // 根据类型获取节点信息对象
	GetCheck(string) (string, error)                            // 根据ID获取节点的校验信息
	GetFeatures(string) (string, error)                         // 根据ID获取节点的特征信息
	GetIDsByBelong(string) ([]string, error)                    // 根据所属信息获取所有匹配的节点的ID
}

/* AccountInfo接口 */
type AccountInfo interface {
	Get(string) (string, error)                    // 根据ID获取账户信息
	GetType(string) (int, error)                   // 根据ID获取节点的类型
	GetByType(int) (*[]AccountInformations, error) // 根据账户类型获取符合的账户并返回
	Add() error                                    // 添加账户信息
	Delete(string) error                           // 根据ID删除账户信息
	Update() error                                 // 根据ID更新账户信息
	GetPasswdFragment(string) (string, error)      // 根据ID获取账户密码特征信息
	GetFragment(string) (string, error)            // 查询指定ID的用户的密码残片
	GetOrganization(string) (string, error)        // 查询指定ID的用户的所属机构信息
	Exist(string) (bool, error)                    // 判断账户是否存在
}

/* 操作日志表接口 */
type OperationRecordsInterface interface {
	Insert() error                            // 记录操作日志
	GetByID(string) (string, error)           // 根据操作账户ID获取操作日志
	GetByTime(string, string) (string, error) // 根据日期半径获取操作日志
	GetResult(string) (string, error)         // 只获取指定id的操作结果日志
	Delete(string) error                      // 删除操作日志
}

// Organization接口
type OrganizationInfo interface {
	Insert() error                                 // 新增企业信息
	Update() error                                 // 更新企业信息
	Delete(string) error                           // 删除企业信息
	GetByID(string) (*Organization, error)         // 根据企业ID获取企业信息
	GetByName(string) (*[]Organization, error)     // 根据企业名称获取企业信息
	GetByType(string) (*[]Organization, error)     // 根据企业类型获取企业信息
	GetByStandard(string) (*[]Organization, error) //    // 根据服从标准获取企业信息
}
