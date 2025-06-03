package task

import (
	"database/sql"
	"fmt"
)

func main_tran() {
	db := initDB()

	flatAb := checkAccountBalance(db, 100, 2)
	fmt.Println("是否可以转账", flatAb)
	if flatAb { //余额充足可以转账
		tx, err := db.Begin()
		if err != nil {
			fmt.Println("开启事务失败")
			return
		}
		// 使用defer处理事务回滚
		var transactionErr error // 新增变量来跟踪事务错误
		// 使用defer处理事务回滚
		defer func() {
			if transactionErr != nil {
				tx.Rollback()
				fmt.Println("事务回滚:", err)
			} else {
				err = tx.Commit()
				if err != nil {
					fmt.Println("事务提交失败:", err)
					tx.Rollback()
				} else {
					fmt.Println("事务提交成功")
				}
			}
		}()

		fmt.Println("进行转账....")
		if transactionErr = transfer(tx, 2, 1, 100); transactionErr != nil {
			fmt.Println("转账失败")
			return
		}
		fmt.Println("转账信息记录...")
		if transactionErr = insertTransactions(tx, 2, 1, 100); transactionErr != nil {

			fmt.Println("记录失败")
		}
	} else {
		fmt.Println("余额不足...")
	}
}

func checkAccountBalance(db *sql.DB, amount float32, id int) bool {
	var account Account
	err := db.QueryRow("SELECT balance FROM ajxd_account WHERE id= ?", id).
		Scan(&account.Balance)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Errorf("account %d 查找不到", id)
			return false
		}
		return false
	}
	if amount > account.Balance {
		fmt.Errorf("account %d 金额不足", id)
		return false
	}
	return true
}

func transfer(tx *sql.Tx, from int, to int, amount float32) error {
	_, errF := tx.Exec("update ajxd_account set balance= balance-? where id= ?", amount, from)
	if errF != nil {
		//需要处理错误，并记录日志
		return fmt.Errorf("扣款失败： %v", errF)
	}
	_, errTo := tx.Exec("update ajxd_account set balance= balance+? where id= ?", amount, to)
	if errTo != nil {
		//需要处理错误，并记录日志
		return fmt.Errorf("收款失败： %v", errTo)
	}
	return nil
}

func insertTransactions(tx *sql.Tx, from int, to int, amount float32) error {
	_, err := tx.Exec("INSERT INTO  transactions (from_account_id,to_account_id,amount)VALUES (?,?,?)", from, to, amount)
	if err != nil {
		return fmt.Errorf("记录失败： %v", err)
	}
	return nil
}

type Account struct {
	Id      int `db:"id,omitempty"`
	Balance float32
}

type Transactions struct {
	Id            int `db:"id,omitempty"`
	FromAccountId int
	ToAccountId   int
	Amount        float32
}
