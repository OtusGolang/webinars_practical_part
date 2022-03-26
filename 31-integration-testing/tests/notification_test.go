package scripts

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/cucumber/godog"
	"github.com/cucumber/messages-go/v16"
	"github.com/streadway/amqp"
)

var amqpDSN = os.Getenv("TESTS_AMQP_DSN")

func init() {
	if amqpDSN == "" {
		amqpDSN = "amqp://guest:guest@localhost:5672/"
	}
}

const (
	queueName                 = "ToNotificationTest"
	notificationsExchangeName = "UserNotifications"
)

type notifyTest struct {
	conn          *amqp.Connection
	ch            *amqp.Channel
	messages      [][]byte
	messagesMutex *sync.RWMutex
	stopSignal    chan struct{}

	responseStatusCode int
	responseBody       []byte
}

func panicOnErr(err error) {
	if err != nil {
		panic(err)
	}
}

func (test *notifyTest) startConsuming(ctx context.Context, _ *messages.Pickle) (context.Context, error) {
	test.messages = make([][]byte, 0)
	test.messagesMutex = new(sync.RWMutex)
	test.stopSignal = make(chan struct{})

	var err error

	test.conn, err = amqp.Dial(amqpDSN)
	panicOnErr(err)

	test.ch, err = test.conn.Channel()
	panicOnErr(err)

	// Consume
	_, err = test.ch.QueueDeclare(queueName, true, true, true, false, nil)
	panicOnErr(err)

	err = test.ch.QueueBind(queueName, "", notificationsExchangeName, false, nil)
	panicOnErr(err)

	events, err := test.ch.Consume(queueName, "", true, true, false, false, nil)
	panicOnErr(err)

	go func(stop <-chan struct{}) {
		for {
			select {
			case <-stop:
				return
			case event := <-events:
				test.messagesMutex.Lock()
				test.messages = append(test.messages, event.Body)
				test.messagesMutex.Unlock()
			}
		}
	}(test.stopSignal)

	return ctx, nil
}

func (test *notifyTest) stopConsuming(ctx context.Context, _ *messages.Pickle, _ error) (context.Context, error) {
	test.stopSignal <- struct{}{}

	panicOnErr(test.ch.Close())
	panicOnErr(test.conn.Close())
	test.messages = nil

	return ctx, nil
}

func (test *notifyTest) iSendRequestTo(httpMethod, addr string) (err error) {
	var r *http.Response

	switch httpMethod {
	case http.MethodGet:
		r, err = http.Get(addr)
	default:
		err = fmt.Errorf("unknown method: %s", httpMethod)
	}

	if err != nil {
		return
	}
	test.responseStatusCode = r.StatusCode
	test.responseBody, err = ioutil.ReadAll(r.Body)
	return
}

func (test *notifyTest) theResponseCodeShouldBe(code int) error {
	if test.responseStatusCode != code {
		return fmt.Errorf("unexpected status code: %d != %d", test.responseStatusCode, code)
	}
	return nil
}

func (test *notifyTest) theResponseShouldMatchText(text string) error {
	if string(test.responseBody) != text {
		return fmt.Errorf("unexpected text: %s != %s", test.responseBody, text)
	}
	return nil
}

func (test *notifyTest) iSendRequestToWithData(httpMethod, addr, contentType string, data *messages.PickleDocString) (err error) {
	var r *http.Response

	switch httpMethod {
	case http.MethodPost:
		replacer := strings.NewReplacer("\n", "", "\t", "")
		cleanJson := replacer.Replace(data.Content)
		r, err = http.Post(addr, contentType, bytes.NewReader([]byte(cleanJson)))
	default:
		err = fmt.Errorf("unknown method: %s", httpMethod)
	}

	if err != nil {
		return
	}
	test.responseStatusCode = r.StatusCode
	test.responseBody, err = ioutil.ReadAll(r.Body)
	return
}

func (test *notifyTest) iReceiveEventWithText(text string) error {
	time.Sleep(3 * time.Second) // На всякий случай ждём обработки евента

	test.messagesMutex.RLock()
	defer test.messagesMutex.RUnlock()

	for _, msg := range test.messages {
		if string(msg) == text {
			return nil
		}
	}
	return fmt.Errorf("event with text '%s' was not found in %s", text, test.messages)
}

func InitializeScenario(s *godog.ScenarioContext) {
	test := new(notifyTest)

	s.Before(test.startConsuming)

	s.Step(`^I send "([^"]*)" request to "([^"]*)"$`, test.iSendRequestTo)
	s.Step(`^The response code should be (\d+)$`, test.theResponseCodeShouldBe)
	s.Step(`^The response should match text "([^"]*)"$`, test.theResponseShouldMatchText)

	s.Step(`^I send "([^"]*)" request to "([^"]*)" with "([^"]*)" data:$`, test.iSendRequestToWithData)
	s.Step(`^I receive event with text "([^"]*)"$`, test.iReceiveEventWithText)

	s.After(test.stopConsuming)
}
