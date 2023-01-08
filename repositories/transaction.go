package repositories

import (
	"dewe/models"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	FindTransactions() ([]models.Transaction, error)
	GetTransaction(ID int) (models.Transaction, error)
	CreateTransaction(transaction models.Transaction) (models.Transaction, error)
	UpdateTransaction(transaction models.Transaction, ID int) (models.Transaction, error)
	DeleteTransaction(transaction models.Transaction) (models.Transaction, error)
	UpdateTransactions(status string, ID int) error
	GetOneTransaction(ID string) (models.Transaction, error)
}

func RepositoryTransaction(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindTransactions() ([]models.Transaction, error) {
	var transactions []models.Transaction
	err := r.db.Preload("User").Preload("Trip").Preload("Trip.Country").Find(&transactions).Error

	return transactions, err
}

func (r *repository) GetTransaction(ID int) (models.Transaction, error) {
	var transaction models.Transaction
	err := r.db.Preload("User").Preload("Trip").Preload("Trip.Country").First(&transaction, ID).Error

	return transaction, err
}

func (r *repository) CreateTransaction(transaction models.Transaction) (models.Transaction, error) {
	err := r.db.Preload("Trip").Preload("Trip.Country").Create(&transaction).Error

	// if transaction.Status == "success" {
	// 	transaction.Trip.Quota = transaction.Trip.Quota - transaction.CounterQty
	// }

	return transaction, err
}

func (r *repository) UpdateTransaction(transaction models.Transaction, ID int) (models.Transaction, error) {

	err := r.db.Model(&transaction).Updates(transaction).Error

	return transaction, err
}

func (r *repository) DeleteTransaction(transaction models.Transaction) (models.Transaction, error) {
	err := r.db.Preload("Trip.Country").Delete(&transaction).Error

	return transaction, err
}

func (r *repository) UpdateTransactions(status string, ID int) error {
	var transaction models.Transaction
	r.db.Preload("Trip").First(&transaction, ID)

	// If is different & Status is "success" decrement product quantity
	if status != transaction.Status && status == "success" {
		var trip models.Trip
		r.db.First(&trip, transaction.Trip.ID)
		trip.Quota = trip.Quota - transaction.CounterQty
		r.db.Save(&trip)
	}

	transaction.Status = status

	err := r.db.Save(&transaction).Error

	return err
}

func (r *repository) GetOneTransaction(ID string) (models.Transaction, error) {
	var transaction models.Transaction
	err := r.db.Preload("Trip").Preload("Trip.User").Preload("User").First(&transaction, "id = ?", ID).Error

	return transaction, err
}
