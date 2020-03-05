package xsql

import (
	"database/sql"
	"log"
)

type Cluster struct {
	name     string
	master   *sql.DB
	slaves   []*sql.DB
	selector Selector
}

func newDB(options ConnectionOptions) (*sql.DB, error) {
	db, err := sql.Open(options.DriverName, options.Dsn)
	if err != nil {
		return nil, err
	}
	db.SetMaxIdleConns(options.MaxIdleConn)
	db.SetMaxOpenConns(options.MaxOpenConn)
	db.SetConnMaxLifetime(options.ConnMaxLifeTime)
	return db, nil
}

// OpenCluster init new gorm client cluster with mysql database.
// name is instance name of current mysql instance.
// selector to select which slave node to use.
// masterOption is master node options
// slaveOptions are slave nodes options
func OpenCluster(name string, selector Selector, masterOption ConnectionOptions, slaveOptions []ConnectionOptions) (*Cluster, error) {
	if selector == nil {
		selector = RoundRobinSelector()
	}
	db, err := newDB(masterOption)
	if err != nil {
		return nil, err
	}
	slaves := make([]*sql.DB, 0, len(slaveOptions))
	for _, slaveCfg := range slaveOptions {
		db, err := newDB(slaveCfg)
		if err != nil {
			// open connection to slave node failed, not affect the whole cluster state
			log.Printf("[xdatabase/xsql] open slave node failed, driver_name(%s), dsn(%s), details(%+v\n)", slaveCfg.DriverName, slaveCfg.Dsn, err)
			continue
		}
		slaves = append(slaves, db)
	}
	return &Cluster{
		name:     name,
		master:   db,
		slaves:   slaves,
		selector: selector,
	}, nil
}

// Master return master node to do more operation.
// 返回主节点。
func (c *Cluster) Master() *sql.DB {
	return c.master
}

// SlaveBySelector get slave db use user `Selector`.
// 使用用户指定的选择器选择从节点。
func (c *Cluster) SlaveBySelector(selector Selector) *sql.DB {
	if len(c.slaves) == 0 {
		return c.master
	}
	if len(c.slaves) == 1 {
		return c.slaves[0]
	}
	return c.slaves[selector.SelectDB(len(c.slaves))]
}

// SlaveByKey get slave db use shard `Selector` with shard key.
// 使用key来选择从节点。
func (c *Cluster) SlaveByKey(key string) *sql.DB {
	return c.SlaveBySelector(ShardSelector(key))
}

// Slave get slave db use default `Selector`.
// 使用默认的选择器选择从节点。
func (c *Cluster) Slave() *sql.DB {
	return c.SlaveBySelector(c.selector)
}

// // Query will retrieve data slice from database, prefer to use slave node than master node.
// // 从database 获取多条数据，优先从对应的从节点读取。
// func (c *Cluster) Query(dest interface{}, sql string, args ...interface{}) error {
// 	return c.Slave().Raw(sql, args...).Find(dest).Error
// }

// // QueryOne will retrieve one record from database, prefer to use slave node than master node.
// // 从database 获取单条数据，优先从对应的从节点读取。
// func (c *Cluster) QueryOne(dest interface{}, sql string, args ...interface{}) error {
// 	return c.Slave().Raw(sql, args...).First(dest).Error
// }

// Exec will do update/insert/delete operation with master node.
// 在主节点上执行 更新/插入/删除 操作。
func (c *Cluster) Exec(sql string, args ...interface{}) error {
	_, err := c.Master().Exec(sql, args...)
	return err
}

// Run will do operation on master or slave node with user choose.
// 用户可选的在何种节点上执行操作。
func (c *Cluster) Run(fn func(master *sql.DB, slave *sql.DB) error) error {
	return fn(c.Master(), c.Slave())
}

// Run will do operation on master or slave node with user choose, with key to select which slave node.
// 用户可选的在何种节点上执行操作，key 用来选择从节点。
func (c *Cluster) RunShard(key string, fn func(master *sql.DB, slave *sql.DB) error) error {
	return fn(c.Master(), c.SlaveByKey(key))
}

// Transaction will do transaction on master node, with commit/rollback automatically.
// 在主节点上执行事务，可以自动执行commit/rollback。
func (c *Cluster) Transaction(fn func(tx *sql.Tx) error) error {
	master := c.master
	db, err := master.Begin()
	if err != nil {
		return err
	}
	if err := fn(db); err != nil {
		_ = db.Rollback()
		return err
	}
	return db.Commit()
}
