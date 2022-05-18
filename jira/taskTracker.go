package jira

import (
	"context"
	"errors"
	"fmt"
	"github.com/andygrunwald/go-jira"
	"github.com/vlamai/papir_jira/config"
	"github.com/vlamai/papir_jira/proto/pb"
	"log"
	"net/http"
)

type TaskTracker struct {
	pb.UnimplementedTaskTrackerServer

	Client *jira.Client
}

func New(config config.Jira) (TaskTracker, error) {
	tp := jira.BasicAuthTransport{
		Username: config.Name,
		Password: config.Token,
	}

	client, err := jira.NewClient(tp.Client(), config.BaseUrl)
	if err != nil {
		return TaskTracker{}, err
	}

	tracker := TaskTracker{
		Client: client,
	}

	return tracker, nil
}

func (t TaskTracker) GetTaskName(_ context.Context, req *pb.GetTaskNameRequest) (*pb.GetTaskNameResponse, error) {
	taskId := req.TaskId
	issue, responce, err := t.Client.Issue.Get(taskId, nil)
	if err != nil {
		log.Printf("JIRA TASK ID: %s | %v\n", taskId, err)
		return nil, err
	}

	if responce.StatusCode != http.StatusOK {
		err = errors.New(fmt.Sprintf("JIRA TASK ID: %s | %v", taskId, responce.Body))
		log.Printf("%v", err)
		return nil, err
	}

	taskName := fmt.Sprintf(
		"%s: %s\n%s",
		issue.Key,
		issue.Fields.Summary,
		issue.Fields.Type.Name,
	)
	response := pb.GetTaskNameResponse{
		TaskName: taskName,
	}
	return &response, nil
}
