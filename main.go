// main.go

package main

import (
    "notification_service/pkg/notification"
)

func main() {
    // Initialize the notification service
    notificationService := notification.NewPushNotificationService()

    // Example usage
    message := "Hello, world!"
    recipient := "example@example.com"
    err := notificationService.SendNotification(message, recipient)
    if err != nil {
        // Handle error
        panic(err)
    }
}
