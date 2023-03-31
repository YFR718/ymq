package topic

import (
	"errors"
	"github.com/YFR718/ymq/pkg/common"
)

type Topic struct {
	Name        string
	Partitions  int
	Replication int
	MessageSize int
}

type TopicManager struct {
	Topics map[string]chan []byte
}

func (t *TopicManager) Create(topic Topic) error {
	if _, ok := t.Topics[topic.Name]; ok {
		return errors.New(common.TOPIC_EXIST)
	}
	t.Topics[topic.Name] = make(chan []byte, 10000)
	return nil
}

func (t *TopicManager) Delete(topic Topic) error {
	if _, ok := t.Topics[topic.Name]; !ok {
		return errors.New(common.TOPIC_NOT_EXIST)
	}
	delete(t.Topics, topic.Name)
	return nil
}

func (t *TopicManager) Send(topic Topic, data []byte) error {
	if _, ok := t.Topics[topic.Name]; !ok {
		return errors.New(common.TOPIC_NOT_EXIST)
	}
	t.Topics[topic.Name] <- data
	return nil
}

var TopicManagerInstance *TopicManager

func InitManager() {
	TopicManagerInstance = &TopicManager{
		Topics: make(map[string]chan []byte),
	}
}
