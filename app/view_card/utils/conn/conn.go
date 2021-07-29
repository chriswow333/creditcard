package conn

type Service interface {
	GetConn() (*Connection, error)
	Commit(conn *Connection) error
	RollBack(conn *Connection) error
	Exec(conn *Connection, sql string, updater ...interface{}) error
}

type Connection struct {
	Connection interface{}
}
