package workflow

import (
	"time"

	"temporal-go-demo/app/activity"

	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
)

const TransferMoneyTaskQueue = "TRANSFER_MONEY_TASK_QUEUE"

func TransferMoneyWorkflow(ctx workflow.Context, transferDetails activity.TransferDetails) error {
	// RetryPolicy specifies how to automatically handle retries if an Activity fails.
	retrypolicy := &temporal.RetryPolicy{
		InitialInterval:        time.Second,
		BackoffCoefficient:     2.0, // 1.0,
		MaximumInterval:        time.Minute,
		MaximumAttempts:        5,
		NonRetryableErrorTypes: []string{},
	}
	options := workflow.ActivityOptions{
		// Timeout options specify when to automatically timeout Activity functions.
		StartToCloseTimeout: time.Minute,
		// Optionally provide a customized RetryPolicy.
		// Temporal retries failures by default, this is just an example.
		RetryPolicy: retrypolicy,
	}

	ctx = workflow.WithActivityOptions(ctx, options)

	err := workflow.ExecuteActivity(ctx, activity.TransferMoneyActivity, transferDetails).Get(ctx, nil)
	if err != nil {
		return err
	}

	err = workflow.ExecuteActivity(ctx, activity.UpdateCachedActivity, activity.UpdateCacheArgs{
		Amount:    -transferDetails.Amount,
		AccountID: transferDetails.FromAccount,
	}).Get(ctx, nil)
	if err != nil {
		return err
	}

	err = workflow.ExecuteActivity(ctx, activity.UpdateCachedActivity, activity.UpdateCacheArgs{
		Amount:    transferDetails.Amount,
		AccountID: transferDetails.ToAccount,
	}).Get(ctx, nil)
	if err != nil {
		return err
	}

	return nil
}
