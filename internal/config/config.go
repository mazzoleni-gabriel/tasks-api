package config

type (
	Configuration struct {
		Addr               string
		TopicTaskFeed      string
		TopicTaskFeedScope string
	}
)

func NewConfig() (Configuration, error) {
	// @todo load configs from properties
	conf := Configuration{
		Addr:               ":8080",
		TopicTaskFeed:      "fake_topic_task_feed",
		TopicTaskFeedScope: "fake",
	}
	return conf, nil
}
