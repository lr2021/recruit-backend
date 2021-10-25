package mail

import (
	"github.com/go-gomail/gomail"
	"github.com/lr2021/recruit-backend/general/config"
	"log"
	"time"
)

type sendMail struct {
	Msg *gomail.Message
	Goal chan int
}

var ch = make(chan sendMail, 20)

func Init() {
	go func() {
		d := gomail.NewDialer(config.SMTP_HOST, 587, config.SMTP_USER, config.SMTP_PASS)

		var s gomail.SendCloser
		var err error
		open := false
		for {
			select {
			case m, ok := <-ch:
				if !ok {
					m.Goal <- -1
					return
				}
				if !open {
					s, err = d.Dial()
					times := 0
					for err != nil && times < 5 {
						log.Print(err)
						repeatSec := 5 + 5 * times
						time.Sleep(time.Duration(repeatSec) * time.Second)
						s, err = d.Dial()
						if err == nil {
							open = true
							break
						}
						times++
					}
					if err != nil {
						panic(err)
					}
				}
				if err := gomail.Send(s, m.Msg); err != nil {
					m.Goal <- -1
					log.Print(err)
				}
				m.Goal <- 0
			// Close the connection to the SMTP server if no email was sent in
			// the last 30 seconds.
			case <-time.After(30 * time.Second):
				if open {
					if err := s.Close(); err != nil {
						panic(err)
					}
					open = false
				}
			}
		}
	}()

	close(ch)
}

func Send(m *gomail.Message) (bool, error) {
	statusChan := make(chan int)
	ch <- sendMail{
		Msg:  m,
		Goal: statusChan,
	}
	status := <-statusChan
	if status == -1 {
		return false, nil
	}
	return true, nil
}

