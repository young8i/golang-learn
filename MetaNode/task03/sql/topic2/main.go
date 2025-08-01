package main

import (
	"database/sql"
	"fmt"
	"log"
	"metanode/task03"
)

/**
 * accounts 表（包含字段 id 主键， balance 账户余额）和
 * transactions 表（包含字段 id 主键， from_account_id 转出账户ID， to_account_id 转入账户ID， amount 转账金额）。
 */
type Account struct {
	Id      int     `gorm:"primaryKey"`
	Balance float64 //账户余额
}
type Transaction struct {
	Id            int     `gorm:"primaryKey"`
	FromAccountId int     //转出账户ID
	ToAccountId   int     //转入账户ID
	Amount        float64 //转账金额
}

// 编写一个事务，实现从账户 A 向账户 B 转账 100 元的操作。
// 在事务中，需要先检查账户 A 的余额是否足够，如果足够则从账户 A 扣除 100 元，向账户 B 增加 100 元，并在 transactions 表中记录该笔转账信息。
// 如果余额不足，则回滚事务。
func transferFunds(db *sql.DB, fromID, toID int, amount float64) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	// 查询账户A余额
	var balance float64
	err = tx.QueryRow("SELECT balance FROM accounts WHERE id = ?", fromID).Scan(&balance)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("查询账户A余额失败: %v", err)
	}

	// 检查余额是否充足
	if balance < amount {
		tx.Rollback()
		return fmt.Errorf("账户A余额不足: %.2f < %.2f", balance, amount)
	}

	// 扣除账户A余额
	_, err = tx.Exec("UPDATE accounts SET balance = balance - ? WHERE id = ?", amount, fromID)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("扣除账户A余额失败: %v", err)
	}

	// 增加账户B余额
	_, err = tx.Exec("UPDATE accounts SET balance = balance + ? WHERE id = ?", amount, toID)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("增加账户B余额失败: %v", err)
	}

	// 写入转账记录
	_, err = tx.Exec(`INSERT INTO transactions (from_account_id, to_account_id, amount) VALUES (?, ?, ?)`,
		fromID, toID, amount)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("记录转账失败: %v", err)
	}

	// 提交事务
	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("事务提交失败: %v", err)
	}

	return nil
}

func main() {
	//初始化数据库
	db := task03.InitDb()

	err := transferFunds(db, 1, 2, 100.00)
	if err != nil {
		log.Printf("转账失败: %v", err)
	} else {
		fmt.Println("转账成功！")
	}
}
