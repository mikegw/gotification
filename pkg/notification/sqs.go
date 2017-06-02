package notification

import (
    "github.com/aws/aws-sdk-go/service/sqs"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/aws"
)

/*--- Public Data Types ---*/


/* Implements Persistor interface */
type SQSPersistorImpl struct {
    MessageSender sqsMessageSender
}

func (persistor SQSPersistorImpl) Persist(model Model) (string, error) {
    message := sqsMessageInput{
        Body: model.Payload,
        QueueURL: "https://sqs.us-east-1.amazonaws.com/654041455006/mike_test",
    }
    result, err := persistor.MessageSender.SendMessage(message)
    if err != nil {
        return "", err
    }
    return result.MessageId, nil
}



/*--- Public Functions ---*/

func NewSQSPersistor(config *aws.Config) (persistor SQSPersistorImpl) {
    messageSender := newSQSMessageSender(config)
    return SQSPersistorImpl{messageSender}
}



/*--- Mocks ---*/

type MockMessageSender struct {
    sqsInput sqsMessageInput
}

func (m *MockMessageSender) SendMessage(input sqsMessageInput) (sqsMessageOutput, error) {
    m.sqsInput = input
    return sqsMessageOutput{}, nil
}

func (m MockMessageSender) InputBody() string {
    return m.sqsInput.Body
}



/*--- Private ---*/

type sqsMessageInput struct {
    Body string
    QueueURL string
}

type sqsMessageOutput struct {
    MessageId string
}

type sqsMessageSender interface {
    SendMessage(input sqsMessageInput) (output sqsMessageOutput, err error)
}

type sqsMessageSenderImpl struct {
    sqsSession *sqs.SQS
}

func (s *sqsMessageSenderImpl) SendMessage(input sqsMessageInput) (output sqsMessageOutput, err error) {
    sendMessageInput := newSendMessageInput(input)
    sendMessageOutput, err := s.sqsSession.SendMessage(sendMessageInput)
    if err != nil {
        return output, err
    }
    output = newSQSMessageOutput(sendMessageOutput)
    return output, nil
}



func newSendMessageInput(input sqsMessageInput) (*sqs.SendMessageInput) {
    return &sqs.SendMessageInput{
        MessageBody:  aws.String(input.Body),
        QueueUrl:     aws.String(input.QueueURL),
    }
}

func newSQSMessageOutput(output *sqs.SendMessageOutput) (sqsMessageOutput) {
    return sqsMessageOutput{
        MessageId: *output.MessageId,
    }
}

func newSQSMessageSender(config *aws.Config) sqsMessageSender {
    sess := session.Must(session.NewSession(config))
    return &sqsMessageSenderImpl{sqs.New(sess)}
}
