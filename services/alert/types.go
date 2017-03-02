package alert

import "github.com/influxdata/kapacitor/alert"

// HandlerSpecRegistrar is responsible for registering and persisting handler spec definitions.
type HandlerSpecRegistrar interface {
	// RegisterHandlerSpec saves the handler spec and registers a handler defined by the spec
	RegisterHandlerSpec(HandlerSpec) error
	// DeregisterHandlerSpec deletes the handler spec and deregisters the defined handler.
	DeregisterHandlerSpec(id string) error
	// UpdateHandlerSpec updates the old spec with the new spec and takes care of registering new handlers based on the new spec.
	UpdateHandlerSpec(oldSpec, newSpec HandlerSpec) error
	// HandlerSpec returns a handler spec
	HandlerSpec(id string) (HandlerSpec, error)
	// Handlers returns a list of handler specs that match the pattern.
	HandlerSpecs(pattern string) ([]HandlerSpec, error)
}

// Topics is responsible for querying  the status of topics and their events.
type Topics interface {
	// TopicStatus returns the status of the specified topic,
	TopicStatus(topic string) (alert.TopicStatus, error)

	// TopicStatusEvents returns the current state of events for the specified topic.
	// Only events greater or equal to minLevel will be returned
	TopicStatusEvents(topic string, minLevel alert.Level) (map[string]alert.EventState, error)

	// TopicEventState returns the current state of the event.
	TopicEventState(topic, event string) (alert.EventState, bool)

	// ListTopicStatus returns the status of all topics that match the pattern and have at least minLevel.
	ListTopicStatus(pattern string, minLevel alert.Level) (map[string]alert.TopicStatus, error)
}

// AnonHandlerRegistrar is responsible for directly registering handlers for anonymous topics.
// This is to be used only when the origin of the handler is not defined by a handler spec.
type AnonHandlerRegistrar interface {
	// RegisterHandler registers the handler instance for the listed topics.
	RegisterAnonHandler(topics []string, h alert.Handler)
	// DeregisterHandler removes the handler from the listed topics.
	DeregisterAnonHandler(topics []string, h alert.Handler)
}

// Events is responsible for accepting events for processing and reporting on the state of events.
type Events interface {
	// Collect accepts a new event for processing.
	Collect(event alert.Event) error
	// UpdateEvent updates an existing event with a previously known state.
	UpdateEvent(topic string, event alert.EventState) error
	// TopicEventState returns the current events state.
	TopicEventState(topic, event string) (alert.EventState, bool)
}

// TopicPersister is responsible for controlling the persistence of topic state.
type TopicPersister interface {
	// CloseTopic closes a topic but does not delete its state.
	CloseTopic(topic string) error
	// DeleteTopic closes a topic and deletes all state associated with the topic.
	DeleteTopic(topic string) error
	// RestoreTopic signals that a topic should be restored from persisted state.
	RestoreTopic(topic string) error
}
