package workflow

import (
	"log"

	"temporal-go-demo/app/activity"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

type Worker struct {
	w worker.Worker
	c client.Client
}

func NewWorker() *Worker {
	// Create the client object just once per process
	c, err := client.NewLazyClient(client.Options{})
	if err != nil {
		log.Fatalln("unable to create Temporal client", err)
		panic(err)
	}
	// This worker hosts both Workflow and Activity functions
	w := worker.New(c, TransferMoneyTaskQueue, worker.Options{})
	w.RegisterWorkflow(TransferMoneyWorkflow)
	w.RegisterActivity(activity.TransferMoneyActivity)
	w.RegisterActivity(activity.UpdateCachedActivity)
	return &Worker{
		w: w,
		c: c,
	}
}

func (w *Worker) Start() {
	err := w.w.Run(worker.InterruptCh())
	if err != nil {
		panic(err)
	}
}

func (w *Worker) Close() {
	w.c.Close()
}
