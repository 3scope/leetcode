package main

func rand7() int {
	return 7
}

func rand10() int {
	// Rand method to get 0 ~ 6.
	a := rand7() - 1
	b := rand7()
	// Rand method to get 1 ~ 49.
	c := a*7 + b
	// Rand method to get 1 ~ 40.
	// 对于 41 ~ 49，比较简单的处理方式是直接抛弃，直到获取的数字是 1 到 40 为止（拒绝采样）。每次运行程序会生成 1 到 40 的概率 p 为 40/49，根据独立事件的期望公式 Ex = n * p, 程序运行的期望运行次数为 n 为1.225(49 / 40)，每次运行会调用 2 次 rand7，所以 rand7 的调用期望为 1.225 * 2。
	if c <= 40 {
		// Set value to 0 ~ 39, and then divided by 4.
		return (c-1)/4 + 1
	}
	// Rand method to get 1 ~ 9.
	c = c - 40
	a = rand7() - 1
	// Rand method to get 1 ~ 63.
	c = a*9 + c
	if c <= 60 {
		return (c-1)/6 + 1
	}
	// Rand method to get 1 ~ 3.
	c = c - 60
	a = rand7() - 1
	// Rand method to get 1 ~ 21.
	c = a*3 + c
	if c <= 20 {
		return (c-1)/2 + 1
	}

	return rand10()
}
