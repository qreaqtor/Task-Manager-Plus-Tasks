package services

import (
	"context"
	"errors"
	"task-manager-plus-tasks/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TaskService struct {
	ctx   *context.Context
	tasks *mongo.Collection
}

func NewTaskService() TaskService {
	return TaskService{
		tasks: tasksCollection,
		ctx:   &ctx,
	}
}

func (ts *TaskService) CreateTask(taskIn models.TaskInput) error {
	_, err := ts.tasks.InsertOne(*ts.ctx, taskIn)
	return err
}

func (ts *TaskService) GetWorkspaceTasks(workspaceId primitive.ObjectID) (tasks []models.TaskOutput, err error) {
	filter := bson.M{"workspace_id": workspaceId}
	cursor, err := ts.tasks.Find(*ts.ctx, filter)
	if err != nil {
		return
	}
	for cursor.Next(*ts.ctx) {
		var taskOut models.TaskOutput
		err = cursor.Decode(&taskOut)
		if err != nil {
			return
		}
		tasks = append(tasks, taskOut)
	}
	return
}

func (ts *TaskService) DeleteTask(taskId primitive.ObjectID) error {
	filter := bson.M{"_id": taskId}
	result, err := ts.tasks.DeleteOne(*ts.ctx, filter)
	if err != err {
		return err
	}
	if result.DeletedCount != 1 {
		return errors.New("no matched document found for delete")
	}
	return nil
}

func (ts *TaskService) DeleteWorkspaceTasks(workspaceId primitive.ObjectID) error {
	filter := bson.M{"workspace_id": workspaceId}
	_, err := ts.tasks.DeleteMany(*ts.ctx, filter)
	if err != err {
		return err
	}
	return nil
}
