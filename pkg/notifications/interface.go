package notifications

// Notifier defines the interface which every type of notification provider
// needs to implement to claim support in openterrascan
type Notifier interface {
	Init(interface{}) error
	SendNotification(interface{}) error
}
