package service

import (
	"context"
	"strconv"

	"github.com/Limerc/E_commerce/gomall/app/checkout/infra/rpc"
	"github.com/Limerc/E_commerce/gomall/rpc_gen/kitex_gen/cart"
	checkout "github.com/Limerc/E_commerce/gomall/rpc_gen/kitex_gen/checkout"
	"github.com/Limerc/E_commerce/gomall/rpc_gen/kitex_gen/order"
	"github.com/Limerc/E_commerce/gomall/rpc_gen/kitex_gen/payment"
	"github.com/Limerc/E_commerce/gomall/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
)

type CheckoutService struct {
	ctx context.Context
} // NewCheckoutService new CheckoutService
func NewCheckoutService(ctx context.Context) *CheckoutService {
	return &CheckoutService{ctx: ctx}
}

// Run create note info
func (s *CheckoutService) Run(req *checkout.CheckoutReq) (resp *checkout.CheckoutResp, err error) {
	// Finish your business logic.
	// 获取购物车商品
	cartResult, err := rpc.CartClient.GetCart(s.ctx, &cart.GetCartReq{UserId: req.UserId})
	if err != nil {
		return nil, kerrors.NewGRPCBizStatusError(5005001, err.Error())
	}
	if cartResult == nil || cartResult.Items == nil {
		return nil, kerrors.NewGRPCBizStatusError(5004001, "cart is empty")
	}

	var (
		total float32
		oi    []*order.OrderItem
	)

	for _, cartItem := range cartResult.Items {
		productResp, resultErr := rpc.ProductClient.GetProduct(s.ctx, &product.GetProductReq{
			Id: cartItem.ProductId, // 实际开发要把rpc放在循环外，否则对性能有很大的影响
		})

		if resultErr != nil {
			return nil, resultErr
		}

		if productResp.Product == nil {
			continue
		}

		p := productResp.Product.Price

		cost := p * float32(cartItem.Quantity)
		total += cost

		oi = append(oi, &order.OrderItem{
			Item: &cart.CartItem{
				ProductId: cartItem.ProductId,
				Quantity:  cartItem.Quantity,
			},
			Cost: cost,
		})
	}

	// 下面创建订单
	var orderId string
	// // 订单服务未开发，先创建个虚拟的订单号
	// u, _ := uuid.NewRandom()
	// orderId = u.String()

	// 引入下单接口
	ZipCodeInt, err := strconv.Atoi(req.Address.ZipCode)
	if err != nil {
		return nil, kerrors.NewBizStatusError(5004002, "zipcode is invalid")
	}
	orderResp, err := rpc.OrderClient.PlaceOrder(s.ctx, &order.PlaceOrderReq{
		UserId:       req.UserId,
		Email:        req.Email,
		Address:      &order.Address{
			StreetAddress:  req.Address.StreetAddress,
			City:           req.Address.City,
			State:          req.Address.State,
			Country:  req.Address.Country,
			ZipCode:        int32(ZipCodeInt),
		},
		Items: oi,
	})
	if err != nil {
		return nil, kerrors.NewGRPCBizStatusError(5004002, "cart is empty")
	}

	if orderResp.Order != nil && orderResp.Order != nil {
		orderId = orderResp.Order.OrderId
	}

	// 创建支付请求订单信息
	payReq := &payment.ChargeReq{
		UserId:  req.UserId,
		OrderId: orderId,
		Amount:  total,
		CreditCard: &payment.CreditCardInfo{
			CreditCardNumber:          req.CreditCard.CreditCardNumber,
			CreditCardCvv:             req.CreditCard.CreditCardCvv,
			CreditCardExpirationYear:  req.CreditCard.CreditCardExpirationYear,
			CreditCardExpirationMonth: req.CreditCard.CreditCardExpirationMonth,
		},
	}
	// 清空购物车
	_, err = rpc.CartClient.EmptyCart(s.ctx, &cart.EmptyCartReq{UserId: req.UserId})
	if err != nil {
		klog.Error(err.Error())
	}
	// 调用支付服务，扣款
	paymentResult, err := rpc.PaymentClient.Charge(s.ctx, payReq)
	if err != nil {
		return nil, err
	}
	// 记录订单信息
	klog.Info(paymentResult)

	resp = &checkout.CheckoutResp{
		OrderId:       orderId,
		TransactionId: paymentResult.TransactionId,
	}
	return resp, nil
}
