// https://go.dev/doc/tutorial/add-a-test
package helper

import (
	"booking-app/09_package_unittest/helper"
	"os"
	"testing"
)

func TestGetUserInput(t *testing.T) {
	// 保存原始的标准输入
	oldStdin := os.Stdin
	defer func() { os.Stdin = oldStdin }() // 测试结束后恢复标准输入
	// 我们首先保存原始的标准输入，并在测试结束后恢复它，以确保不会影响其他测试或程序的正常运行。

	// 创建一个管道，用于模拟标准输入
	// 并将标准输入重定向到管道的读取端。这样我们就可以通过管道的写入端模拟用户输入。
	r, w, err := os.Pipe()
	if err != nil { // nil 是一个预定义的标识符，表示指针、接口、映射、切片、通道和函数类型的零值。它类似于其他语言中的 null 或 None。
		t.Fatalf("os.Pipe() failed: %v", err)
	}

	// 将标准输入重定向到管道的读取端
	os.Stdin = r

	// 准备模拟的用户输入
	userInput := "John\nDoe\njohn.doe@example.com\n2\n"

	// 将模拟的用户输入写入管道的写入端
	_, err = w.Write([]byte(userInput))
	if err != nil {
		t.Fatalf("Write to pipe failed: %v", err)
	}

	// 关闭管道的写入端，以便读取端知道输入已经结束
	w.Close()

	// 调用被测试的函数
	firstName, lastName, email, userTickets := helper.GetUserInput()

	// 验证返回值是否符合预期
	expectedFirstName := "John"
	expectedLastName := "Doe"
	expectedEmail := "john.doe@example.com"
	expectedUserTickets := uint(2)

	if firstName != expectedFirstName {
		t.Errorf("Expected first name %q, but got %q", expectedFirstName, firstName)
	}

	if lastName != expectedLastName {
		t.Errorf("Expected last name %q, but got %q", expectedLastName, lastName)
	}

	if email != expectedEmail {
		t.Errorf("Expected email %q, but got %q", expectedEmail, email)
	}

	if userTickets != expectedUserTickets {
		t.Errorf("Expected user tickets %d, but got %d", expectedUserTickets, userTickets)
	}
}

// go test -v tests/helper_test.go
// 将会看到:
// === RUN   TestGetUserInput
// Enter your first name:
// Enter your lastr name:
// Enter your email:
// Enter your number of tickets:
// --- PASS: TestGetUserInput (0.00s)
// PASS
// ok      command-line-arguments  0.283s
