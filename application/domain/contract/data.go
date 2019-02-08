package contract

// DataManager holds the methods that manipulates the main data.
type DataManager interface {
	repoManager
	Close() error
}

// TransactionManager holds the methods that manipulates the main
// data, from within a transaction.
type TransactionManager interface {
	repoManager
	Rollback() error
	Commit() error
}
