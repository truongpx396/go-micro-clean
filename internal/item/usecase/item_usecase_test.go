package usecase

// import (
// 	"errors"
// 	"go-micro-clean/modules/item/domain/enums"
// 	"go-micro-clean/modules/item/domain/models"
// 	"testing"
// 	"time"

// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/mock"
// 	"gorm.io/gorm"
// )

// // MockItemRepository is a mock implementation of the ItemRepository interface.
// type MockItemRepository struct {
// 	mock.Mock
// }

// func (m *MockItemRepository) Create(item *models.Item) error {
// 	args := m.Called(item)
// 	return args.Error(0)
// }

// func (m *MockItemRepository) GetByID(id uint) (*models.Item, error) {
// 	args := m.Called(id)
// 	return args.Get(0).(*models.Item), args.Error(1)
// }

// func (m *MockItemRepository) GetByName(name string) (*models.Item, error) {
// 	args := m.Called(name)
// 	return args.Get(0).(*models.Item), args.Error(1)
// }

// func (m *MockItemRepository) Update(item *models.Item) error {
// 	args := m.Called(item)
// 	return args.Error(0)
// }

// func (m *MockItemRepository) Delete(id uint) error {
// 	args := m.Called(id)
// 	return args.Error(0)
// }

// func (m *MockItemRepository) List(pagination *models.Pagination, filters ...func(*gorm.DB) *gorm.DB) ([]models.Item, error) {
// 	args := m.Called(pagination, filters)
// 	return args.Get(0).([]models.Item), args.Error(1)
// }

// func TestGetItemByID(t *testing.T) {
// 	mockRepo := new(MockItemRepository)
// 	itemUC := NewItemUsecase(mockRepo)

// 	// Define the item to be returned by the mock repository
// 	expectedItem := &models.Item{
// 		ID:        1,
// 		Name:      "Test Item",
// 		Type:      enums.Physical,
// 		CreatedAt: time.Now(),
// 		UpdatedAt: time.Now(),
// 	}

// 	// Set up the mock repository to return the expected item
// 	mockRepo.On("GetByID", uint(1)).Return(expectedItem, nil)

// 	// Call the use case method
// 	item, err := itemUC.GetItemByID(1)

// 	// Assert the results
// 	assert.NoError(t, err)
// 	assert.Equal(t, expectedItem, item)

// 	// Verify that the mock repository method was called
// 	mockRepo.AssertExpectations(t)
// }

// func TestGetItemByID_NotFound(t *testing.T) {
// 	mockRepo := new(MockItemRepository)
// 	itemUC := NewItemUsecase(mockRepo)

// 	// Set up the mock repository to return an error
// 	mockRepo.On("GetByID", uint(1)).Return(nil, errors.New("item not found"))

// 	// Call the use case method
// 	item, err := itemUC.GetItemByID(1)

// 	// Assert the results
// 	assert.Error(t, err)
// 	assert.Nil(t, item)
// 	assert.Equal(t, "item not found", err.Error())

// 	// Verify that the mock repository method was called
// 	mockRepo.AssertExpectations(t)
// }
