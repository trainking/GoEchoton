package delayq

type DelayQ struct {
	DeadletterExchange string // 死信交换机名称
	DeadletterQueue    string // 死信队列名称

}
