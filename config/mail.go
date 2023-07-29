// Package config 站点配置信息
package config

import "gohub/pkg/config"

func init() {
	config.Add("mail", func() map[string]any {
		return map[string]any{

			// 默认是 Mailhog 的配置
			"smtp": map[string]any{
				"host":     config.Env("MAIL_HOST", "localhost"),
				"port":     config.Env("MAIL_PORT", 1025),
				"username": config.Env("MAIL_USERNAME", ""),
				"password": config.Env("MAIL_PASSWORD", ""),
			},

			"from": map[string]any{
				"address": config.Env("MAIL_FROM_ADDRESS", "gohub@example.com"),
				"name":    config.Env("MAIL_FROM_NAME", "Gohub"),
			},
		}
	})
}
