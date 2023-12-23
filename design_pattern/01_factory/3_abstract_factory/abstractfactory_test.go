package abstractfactory

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAbstractFactory(t *testing.T) {
	var (
		factory DAOFactory
		ret     string

		orderMainDAO   OrderMainDAO
		orderDetailDao OrderDetailDAO
	)

	factory = &RDBDAOFactory{}
	orderMainDAO = factory.CreateOrderMainDAO()
	orderDetailDao = factory.CreateOrderDetailDAO()
	ret = orderMainDAO.SaveOrderMain()
	assert.Equal(t, "rdb main save", ret)
	ret = orderDetailDao.SaveOrderDetail()
	assert.Equal(t, "rdb detail save", ret)

	factory = &XMLDAOFactory{}
	orderMainDAO = factory.CreateOrderMainDAO()
	orderDetailDao = factory.CreateOrderDetailDAO()
	ret = orderMainDAO.SaveOrderMain()
	assert.Equal(t, "xml main save", ret)
	ret = orderDetailDao.SaveOrderDetail()
	assert.Equal(t, "xml detail save", ret)
}
