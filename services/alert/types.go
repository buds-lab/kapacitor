package alert

import "github.com/influxdata/kapacitor/alert"

type HandlerSpecRegistrar interface {
	// RegisterHandlerSpec saves the handler spec and registers a handler defined by the spec
	RegisterHandlerSpec(HandlerSpec)
	// DeregisterHandlerSpec deletes the handler spec and deregisters the defined handler.
	DeregisterHandlerSpec(id string)
	// UpdateHandlerSpec updates the old spec with the new spec and takes care of registering new handlers based on the new spec.
	UpdateHandlerSpec(oldSpec, newSpec HandlerSpec)
	// Handlers returns a list of handler specs that match the pattern.
	HandlerSpecs(pattern string) []HandlerSpec
}

type TopicStatuser interface {
	// TopicStatus returns the status of all topics that match the pattern and have at least minLevel.
	TopicStatus(patter string, minLevel alert.Level) []alert.TopicStatus
	// TopicStatusEvents returns the specific events for each topic that matches the pattern.
	// Only events greater or equal to minLevel will be returned
	TopicStatusEvents(patter string, minLevel alert.Level) map[string]map[string]alert.EventState
}

type HandlerRegistrar interface {
	// RegisterHandler registers the handler instance for the listed topics.
	RegisterHandler(topics []string, h alert.Handler)
	// DeregisterHandler removes the handler from the listed topics.
	DeregisterHandler(topics []string, h alert.Handler)
}

type Eventer interface {
	// Collect accepts a new event for processing.
	Collect(event alert.Event) error
	// UpdateEvent updates an existing event with a previously known state.
	UpdateEvent(topic string, event alert.EventState) error
	// EventState returns the current events state.
	EventState(topic, event string) (alert.EventState, bool)
}

type TopicPersister interface {
	// CloseTopic closes a topic but does not delete its state.
	CloseTopic(topic string) error
	// DeleteTopic closes a topic and deletes all state associated with the topic.
	DeleteTopic(topic string) error
	// RestoreTopic signals that a topic should be restored from persisted state.
	RestoreTopic(topic string) error
}
