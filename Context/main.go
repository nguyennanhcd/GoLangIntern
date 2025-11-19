// bài toán ở đây là gì, bài toán chính là làm sao mà handle được 1 đơn hàng phải thành công trong vòng 2s

package main

import (
	"context"
	"log"
	"net/http"
	"time"
)

// without using context
func placeOrderWithoutContext(orderID string) error {
	log.Printf("Bắt đầu xử lý đơn hàng: %s\n", orderID)
	// Giả sử: thời gian xử lý mất 3 giây (inventory, payment..)
	time.Sleep(3 * time.Second)

	log.Printf("Xử lý đơn hàng %s thành công (sau 3 giây)\n", orderID)
	return nil // Thành công
}

func OrderHandler(w http.ResponseWriter, r *http.Request) {
	orderID := "GO-12345"
	//
	err := placeOrderWithoutContext(orderID)

	if err != nil {
		http.Error(w, "Lỗi xử lý đơn hàng", http.StatusInternalServerError) // 500 Internal Server Error
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Đặt hàng thành công!"))
}

func OrderHandlerSelect(w http.ResponseWriter, r *http.Request) {
	orderID := "GO-12345"
	resultChan := make(chan error, 1)

	go func() {
		err := placeOrderWithoutContext(orderID)
		resultChan <- err
	}()

	select {
	case err := <-resultChan:
		if err != nil {
			log.Printf("Xử lý đơn hàng %s thất bại \n", orderID)
			http.Error(w, "Lỗi xử lý đơn hàng", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Đặt hàng thành công!"))

	case <-time.After(2 * time.Second):
		log.Printf("Xử lý đơn hàng %s quá 2 giây, trả lỗi về client \n", orderID)
		http.Error(w, "Yêu cầu quá thời gian chờ, vui lòng thử lại sau", http.StatusGatewayTimeout) // 504 Gateway Timeout
	}
}

func OrderHandlerWithContext(w http.ResponseWriter, r *http.Request) {
	orderID := "GO-12345"

	// Tạo context có timeout 2 giây
	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
	defer cancel()

	err := placeOrderWithContext(ctx, orderID)

	if err != nil {
		log.Printf("Xử lý đơn hàng %s thất bại: %v\n", orderID, err)
		http.Error(w, "Lỗi xử lý đơn hàng quá thời gian", http.StatusGatewayTimeout)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Đặt hàng thành công!"))
}

// PlaceOrderContext
func placeOrderWithContext(ctx context.Context, orderID string) error {
	log.Printf("Bắt đầu xử lý đơn hàng: %s\n", orderID)

	select {
	case <-time.After(3 * time.Second): // giả lập xử lý mất 3 giây
		log.Printf("Xử lý đơn hàng %s thành công!\n", orderID)
		return nil
	case <-ctx.Done(): // Nếu context bị huỷ
		log.Printf("Huỷ xử lý đơn hàng %s: %v\n", orderID, ctx.Err())
		return ctx.Err()
	}
}

func main() {
	http.HandleFunc("/order", OrderHandlerSelect)
	log.Println("Server đang chạy tại http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
