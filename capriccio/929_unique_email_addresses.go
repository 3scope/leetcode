package main

import "strings"

// 电子邮件地址通过“@”分开，前面是本地名，后面是域名。
func numUniqueEmails(emails []string) int {
	// 用于保存结果。
	count := 0
	// 用于判断域名是否存在。
	isExist := make(map[string]bool)
	for _, email := range emails {
		temp := strings.Split(email, "@")
		// 处理前缀，即处理本地名。
		local := temp[0]
		local = strings.Split(local, "+")[0]
		localTemp := strings.Split(local, ".")
		local = ""
		for _, value := range localTemp {
			local += value
		}
		// 处理后缀，即处理域名。
		domain := temp[1]
		email = local + "@" + domain
		if !isExist[email] {
			count++
			isExist[email] = true
		}
	}

	return count
}
