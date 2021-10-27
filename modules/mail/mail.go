package mail

import (
	"bytes"
	"go-hexagonal-auth/config"
	"go-hexagonal-auth/constants"
	"html/template"
	"strings"

	log "github.com/sirupsen/logrus"

	gomail "gopkg.in/gomail.v2"
)

type mailer struct {
	cfg *config.Mailer
}

type MailDetail struct {
	To       []string
	Cc       []string
	Subject  string
	Template string
	Data     interface{}
}

type MailInterface interface {
	Send(detail MailDetail) error
}

func NewMailer(cfg *config.Mailer) MailInterface {
	return &mailer{
		cfg: cfg,
	}
}

func SetMail(to []string, cc []string, subject string, data interface{}, template string) *MailDetail {
	return &MailDetail{
		To:       to,
		Cc:       cc,
		Subject:  subject,
		Data:     data,
		Template: template,
	}
}

func (m *mailer) Send(detail MailDetail) error {
	mainTemplate := constants.MailMainTemplate
	logoTemplate := constants.MailLogoTemplate
	footerTemplate := constants.MailFooterTemplate

	t := template.Must(template.ParseFiles(mainTemplate, logoTemplate, footerTemplate, detail.Template))
	var tpl bytes.Buffer
	if err := t.ExecuteTemplate(&tpl, "layout", detail.Data); err != nil {
		log.Error(err)
		return err
	}

	result := tpl.String()
	result = strings.ReplaceAll(result, "{LOGO_DESI}", constants.DESILogoURL)

	mailer := gomail.NewMessage()
	mailer.SetHeader("From", m.cfg.Sender)
	mailer.SetHeader("To", detail.To...)
	for _, cc := range detail.Cc {
		mailer.SetAddressHeader("Cc", cc, cc)
	}
	mailer.SetHeader("Subject", detail.Subject)
	mailer.SetBody("text/html", result)

	dialer := gomail.NewPlainDialer(
		m.cfg.Server,
		int(m.cfg.Port),
		m.cfg.Username,
		m.cfg.Password,
	)

	if err := dialer.DialAndSend(mailer); err != nil {
		log.Error(err)
		return err
	}

	return nil
}
