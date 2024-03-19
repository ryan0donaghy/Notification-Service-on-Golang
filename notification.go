// notification.go

package notification

// NotificationService defines the interface for sending notifications.
type NotificationService interface {
    SendNotification(message string, recipient string) error
}
