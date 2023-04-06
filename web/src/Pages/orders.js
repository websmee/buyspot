import { useDispatch, useSelector } from "react-redux";
import { getOrders, updateOrdersData } from "Store/reducer";
import { useEffect } from "react";
import { Navigate } from 'react-router-dom'

import Order from "Components/order"
import OrderSellModal from "Components/orderSellModal"
import OrdersHeader from 'Layouts/ordersHeader';
import ErrorMessage from "Components/errorMessage";
import Footer from "Layouts/footer";

function Orders() {
    const dispatch = useDispatch();
    const orders = useSelector((state) => state.orders);
    const unauthorized = useSelector((state) => state.unauthorized);

    useEffect(() => {
        const intervalId = setInterval(() => {
            dispatch(updateOrdersData());
        }, 1000);
        dispatch(getOrders());

        return () => {
          clearInterval(intervalId);
        };
    }, [dispatch]);

    if (unauthorized) {
        return <Navigate to='/login' />
    }

    return (
        <>
            <OrdersHeader />

            <div className="page-content header-clear-medium">
                <ErrorMessage />

                {orders.map((order, i) =>
                    <Order key={i} modalId={"sell-modal-" + i} order={order} />
                )}
            </div>

            {orders.map((order, i) =>
                <OrderSellModal key={i} id={"sell-modal-" + i} order={order} />
            )}

            <Footer />
        </>
    )
}

export default Orders