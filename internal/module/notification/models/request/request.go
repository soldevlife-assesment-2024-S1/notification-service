package request

type Pagination struct {
	Page int `json:"page" form:"page" required:"true" validate:"required,numeric"`
	Size int `json:"size" form:"size" required:"true" validate:"required,numeric"`
}

type InquiryTicketAmount struct {
	TicketID    int64 `json:"ticket_id" form:"ticket_id" required:"true" validate:"required"`
	TotalTicket int   `json:"total_ticket" form:"total_ticket" required:"true" validate:"required,numeric"`
}

type CheckStockTicket struct {
	TicketDetailID string `form:"ticket_detail_id"`
}

type DecrementTicketStock struct {
	TicketDetailID int64 `json:"ticket_detail_id" form:"ticket_detail_id" validate:"required"`
	TotalTickets   int64 `json:"total_tickets" form:"total_tickets" validate:"required"`
}

type IncrementTicketStock struct {
	TicketDetailID int64 `json:"ticket_detail_id" form:"ticket_detail_id" validate:"required"`
	TotalTickets   int64 `json:"total_tickets" form:"total_tickets" validate:"required"`
}

type PoisonedQueue struct {
	TopicTarget string      `json:"topic_target" validate:"required"`
	ErrorMsg    string      `json:"error_msg" validate:"required"`
	Payload     interface{} `json:"payload" validate:"required"`
}

type NotificationMessage struct {
	Message        string `json:"message" validate:"required"`
	EmailRecipient string `json:"email_recipient" validate:"required"`
}

type NotificationInvoice struct {
	BookingID         string  `json:"booking_id" validate:"required"`
	PaymentAmount     float64 `json:"payment_amount" validate:"required"`
	PaymentExpiration string  `json:"payment_expiration" validate:"required"`
	EmailRecipient    string  `json:"email_recipient" validate:"required"`
}

type NotificationPayment struct {
	BookingID      string `json:"booking_id" validate:"required"`
	Message        string `json:"message" validate:"required"`
	PaymentMethod  string `json:"payment_method" validate:"required"`
	EmailRecipient string `json:"email_recipient" validate:"required"`
}

type SendEmail struct {
	EmailAddress string   `json:"email_address" validate:"required"`
	To           string   `json:"to" validate:"required"`
	Cc           string   `json:"cc"`
	Bcc          string   `json:"bcc"`
	Subject      string   `json:"subject" validate:"required"`
	Body         string   `json:"body" validate:"required"`
	Attachments  []string `json:"attachments"`
}
