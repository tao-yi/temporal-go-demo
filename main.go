package main

import (
	"temporal-go-demo/app/activity"
	"temporal-go-demo/app/workflow"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"go.temporal.io/sdk/client"
)

type Transaction struct {
	FromAccount int     `json:"from_account"`
	ToAccount   int     `json:"to_account"`
	Amount      float32 `json:"amount"`
}

func main() {
	app := fiber.New()
	worker := workflow.NewWorker()
	go func() {
		worker.Start()
	}()
	defer worker.Close()

	cli := workflow.NewClient()
	defer cli.Close()

	app.Post("/api/transaction", func(c *fiber.Ctx) error {
		var t Transaction
		if err := c.BodyParser(&t); err != nil {
			return err
		}

		options := client.StartWorkflowOptions{
			ID:        "transfer-money-workflow",
			TaskQueue: workflow.TransferMoneyTaskQueue,
		}
		transferDetails := activity.TransferDetails{
			Amount:      t.Amount,
			FromAccount: t.FromAccount,
			ToAccount:   t.ToAccount,
			ReferenceID: uuid.New().String(),
		}
		res, err := cli.ExecuteWorkflow(c.Context(), options, workflow.TransferMoneyWorkflow, transferDetails)
		if err != nil {
			return err
		}
		if err = res.Get(c.Context(), nil); err != nil {
			return err
		}

		return c.SendString("success")
	})

	err := app.Listen(":3000")
	if err != nil {
		panic(err)
	}
}
