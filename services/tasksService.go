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

func (ts *TaskService) GetUsersTasks(userId primitive.ObjectID) (tasks []models.TaskOutput, err error) {
	filter := bson.M{"user_id": userId}
	cursor, err := ts.tasks.Find(*ts.ctx, filter)
	if err != nil {
		return nil, err
	}
	tasks = make([]models.TaskOutput, 0)
	for cursor.Next(*ts.ctx) {
		var taskOut models.TaskOutput
		err = cursor.Decode(&taskOut)
		if err != nil {
			return tasks, err
		}
		tasks = append(tasks, taskOut)
	}
	return tasks, nil
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
