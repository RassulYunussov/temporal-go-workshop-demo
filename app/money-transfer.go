package app

import (
	"context"
	"fmt"
	"time"

	"go.temporal.io/sdk/workflow"
)

func TransferMoneyWorkflow(ctx workflow.Context, amount int, fromAccount string, toAccount string) error {
	fmt.Println("TransferMoneyWorkflow")
	options := workflow.ActivityOptions{
		StartToCloseTimeout: time.Second * 5,
		//RetryPolicy:         &temporal.RetryPolicy{MaximumAttempts: 1},
	}

	ctx = workflow.WithActivityOptions(ctx, options)
	fmt.Println("Withdraw")
	err := workflow.ExecuteActivity(ctx, WithdrawAmount, amount, fromAccount, "initial withdrawal").Get(ctx, nil)
	if err != nil {
		return err
	}
	fmt.Println("Deposit")
	err = workflow.ExecuteActivity(ctx, DepositAmount, amount, toAccount, "initial deposit").Get(ctx, nil)
	if err != nil {
		fmt.Println("Deposit errorered")
		return err
	}
	fmt.Println("Finish")
	// if err != nil {
	// 	// return money back to account
	// 	err = workflow.ExecuteActivity(ctx, RollbackAmount, amount, fromAccount).Get(ctx, nil)
	// 	return err
	// }
	return nil
}

func WithdrawAmount(ctx context.Context, amount int, fromAccount string, description string) error {
	fmt.Printf("Withdraw %d from %s, %s\n", amount, fromAccount, description)
	return nil
}

func RollbackAmount(ctx context.Context, amount int, toAccount string) error {
	fmt.Printf("Return %d to %s\n", amount, toAccount)
	return nil
}

func DepositAmount(ctx context.Context, amount int, toAccount string, description string) error {
	// if rand.Intn(100) < 50 {
	// 	fmt.Printf("Deposit %d to %s, %s\n", amount, toAccount, description)
	// 	return nil
	// }
	// fmt.Printf("Failed to deposit %d to %s\n", amount, toAccount)
	// return fmt.Errorf("failed to deposit %d to %s", amount, toAccount)
	fmt.Printf("Deposit %d to %s, %s\n", amount, toAccount, description)
	return nil
}
