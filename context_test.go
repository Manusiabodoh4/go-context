package gocontext

import (
	"context"
	"fmt"
	"testing"
)

func TestCreateContext(t *testing.T) {

	background := context.Background()
	fmt.Println(background)

	todo := context.TODO()
	fmt.Println(todo)

}
