package task01

import "gorm.io/gorm"

/*
假设有两个表： accounts 表（包含字段 id 主键， balance 账户余额）和
transactions 表（包含字段 id 主键， from_account_id 转出账户ID，
to_account_id 转入账户ID， amount 转账金额）。
要求 ：
编写一个事务，实现从账户 A 向账户 B 转账 100 元的操作。在事务中，
需要先检查账户 A 的余额是否足够，如果足够则从账户 A 扣除 100 元，
向账户 B 增加 100 元，并在 transactions 表中记录该笔转账信息。如果余额不足，则回滚事务。
*/
type Account struct {
	ID      int
	Balance int
}

type Transaction struct {
	ID            int
	FromAccountId int
	ToAccountId   int
	Amount        int
}

/*func (a *Account) BeforeUpdate(tx *gorm.DB) (err error) {

	return
}*/

/*func (a *Account) AfterUpdate(tx *gorm.DB) (err error) {


	return
}*/

func Run(db *gorm.DB) {
	err := db.AutoMigrate(&Account{})
	if err != nil {
		return
	}

	err1 := db.AutoMigrate(&Transaction{})
	if err1 != nil {
		return
	}
	/*db.Create(&Account{
		ID:      1,
		Balance: 50,
	})
	db.Create(&Account{
		ID:      2,
		Balance: 150,
	})
	db.Create(&Account{
		ID:      3,
		Balance: 100,
	})*/

	tx := db.Begin()
	account := Account{}
	tx.Model(&Account{}).Where("id = ?", 2).Find(&account)
	if tx.Error != nil {
		tx.Rollback()
	}
	val := account.Balance - 100
	if val < 0 {
		tx.Rollback()
	}
	tx.Model(&Account{}).Where("id=?", 2).Update("balance", val)
	if tx.Error != nil {
		tx.Rollback()
	}
	account3 := Account{}
	tx.Where("id = ?", 3).Find(&account3)
	if tx.Error != nil {
		tx.Rollback()
	}
	addA := account3.Balance + 100
	tx.Model(&Account{}).Where("id=?", 3).Update("balance", addA)
	if tx.Error != nil {
		tx.Rollback()
	}
	tran := Transaction{FromAccountId: 2, ToAccountId: 3, Amount: 100}
	tx.Create(&tran)

	if tx.Error != nil {
		tx.Rollback()
	}
	tx.Commit()
}
