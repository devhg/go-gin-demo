package notify

// OnPanicNotify 发生 panic 时进行通知
// func OnPanicNotify(ctx *gin.Context, err interface{}, stackInfo string) {
// 	cfg := config.AppSetting.Mail
// 	if cfg.Host == "" || cfg.Port == 0 || cfg.UserName == "" || cfg.PassWord == "" || len(cfg.To) == 0 {
// 		ctx.Logger().Error("Mail config error")
// 		return
// 	}

// 	subject, body, htmlErr := NewPanicHTMLEmail(
// 		ctx.Method(),
// 		ctx.Host(),
// 		ctx.URI(),
// 		ctx.Trace().ID(),
// 		err,
// 		stackInfo,
// 	)
// 	if htmlErr != nil {
// 		ctx.Logger().Error("NewPanicHTMLEmail error", zap.Error(htmlErr))
// 		return
// 	}

// 	options := &mail.Options{
// 		MailHost: cfg.Host,
// 		MailPort: cfg.Port,
// 		MailUser: cfg.User,
// 		MailPass: cfg.Pass,
// 		MailTo:   cfg.To,
// 		Subject:  subject,
// 		Body:     body,
// 	}
// 	sendErr := mail.Send(options)
// 	if sendErr != nil {
// 		ctx.Logger().Error("Mail Send error", zap.Error(sendErr))
// 	}

// 	return
// }
