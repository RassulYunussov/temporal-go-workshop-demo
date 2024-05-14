package main

import (
	"context"
	"log"
	"playground/app"
	"time"

	"github.com/pborman/uuid"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

const MONEY_TRANSFER_QUEUE = "TRANSFER_QUEUE"

func main() {
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("unable to create Temporal client", err)
	}
	defer c.Close()

	// This worker hosts both Workflow and Activity functions
	w := worker.New(c, MONEY_TRANSFER_QUEUE, worker.Options{})
	w.RegisterWorkflow(app.TransferMoneyWorkflow)
	w.RegisterActivity(app.WithdrawAmount)
	w.RegisterActivity(app.DepositAmount)
	w.RegisterActivity(app.RollbackAmount)

	go func() {
		for {
			time.Sleep(time.Second * 5)
			options := client.StartWorkflowOptions{
				ID:        uuid.New(),
				TaskQueue: MONEY_TRANSFER_QUEUE,
			}
			// Start the Workflow
			run, err := c.ExecuteWorkflow(context.TODO(), options, app.TransferMoneyWorkflow, 100, "SOURCE", "DESTINATION")
			if err != nil {
				log.Fatalln("unable to submit Workflow", err)
			}
			log.Println("Started workflow: " + run.GetRunID())
		}
	}()

	// Start listening to the Task Queue
	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("unable to start Worker", err)
	}
}
