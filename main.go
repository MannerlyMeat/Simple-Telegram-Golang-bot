package main

import (
	// "io"

	"log"
	"os"
	"time"

	tele "gopkg.in/telebot.v4"
)

func main() {

	pref := tele.Settings{
		Token:  string(os.Getenv("TOKEN")),
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}

	faqBtn := tele.InlineButton{
		Text:   "FAQ",
		Unique: "faq",
	}

	supportBtn := tele.InlineButton{
		Text:   "Поддержка",
		Unique: "support",
	}

	question1Btn := tele.InlineButton{
		Text:   "Вопрос 1",
		Unique: "question1",
	}

	question2Btn := tele.InlineButton{
		Text:   "Вопрос 2",
		Unique: "question2",
	}

	question3Btn := tele.InlineButton{
		Text:   "Вопрос 3",
		Unique: "question3",
	}

	question4Btn := tele.InlineButton{
		Text:   "Вопрос 4",
		Unique: "question4",
	}

	question5Btn := tele.InlineButton{
		Text:   "Вопрос 5",
		Unique: "question5",
	}

	question6Btn := tele.InlineButton{
		Text:   "Вопрос 6",
		Unique: "question6",
	}

	b.Handle("/start", func(c tele.Context) error {
		//Открываем файл
		data, err := os.ReadFile("messages/greetings.txt")
		if err != nil {
			log.Fatal(err)
		}

		//Формируем кнопки
		catalogBtn := tele.ReplyButton{
			Text: "Каталог бренда",
		}
		helpBtn := tele.ReplyButton{
			Text: "Служба заботы",
		}
		appBtn := tele.ReplyButton{
			Text: "Приложение",
		}

		replyKeyboard := tele.ReplyMarkup{
			ReplyKeyboard: [][]tele.ReplyButton{
				{catalogBtn}, {helpBtn}, {appBtn},
			},
		}

		//Преобразуем текст из файла в строку
		text := string(data)

		//Отправка сообщение пользователю
		return c.Send(&tele.Photo{File: tele.FromDisk("images/greetings.jpg"), Caption: text}, &replyKeyboard)
	})

	b.Handle("Приложение", func(c tele.Context) error {
		//Открываем файл
		data, err := os.ReadFile("messages/app.txt")
		if err != nil {
			log.Fatal(err)
		}

		//Преобразуем текст из файла в строку
		text := string(data)
		//Отправка сообщение пользователю
		return c.Send(text)
	})

	b.Handle("Каталог бренда", func(c tele.Context) error {
		//Открываем файл
		data, err := os.ReadFile("messages/catalog.txt")
		if err != nil {
			log.Fatal(err)
		}

		ozonBtn := tele.InlineButton{
			Text: "Ozon",
			URL:  "https://www.ozon.ru/",
		}
		wbBtn := tele.InlineButton{
			Text: "Wildberries",
			URL:  "https://www.wildberries.ru/",
		}

		inlineKeyboard := tele.ReplyMarkup{
			InlineKeyboard: [][]tele.InlineButton{
				{ozonBtn}, {wbBtn},
			},
		}

		//Преобразуем текст из файла в строку
		text := string(data)
		//Отправка сообщение пользователю
		return c.Send(&tele.Photo{File: tele.FromDisk("images/catalog.jpg"), Caption: text}, &inlineKeyboard)
	})

	b.Handle("Служба заботы", func(c tele.Context) error {
		//Открываем файл
		data, err := os.ReadFile("messages/faq_main.txt")
		if err != nil {
			log.Fatal(err)
		}

		inlineKeyboard := tele.ReplyMarkup{
			InlineKeyboard: [][]tele.InlineButton{
				{faqBtn}, {supportBtn},
			},
		}

		//Преобразуем текст из файла в строку
		text := string(data)
		//Отправка сообщение пользователю
		return c.Send(&tele.Photo{File: tele.FromDisk("images/faq_main.jpg"), Caption: text}, &inlineKeyboard)
	})

	b.Handle(&faqBtn, func(c tele.Context) error {
		inlineKeyboard := tele.ReplyMarkup{
			InlineKeyboard: [][]tele.InlineButton{
				{question1Btn}, {question2Btn}, {question3Btn}, {question4Btn}, {question5Btn}, {question6Btn},
			},
		}
		return c.Send("Выбери вопрос из списка:", &inlineKeyboard)
	})

	b.Handle(&supportBtn, func(c tele.Context) error {
		data, err := os.ReadFile("messages/support.txt")
		if err != nil {
			log.Fatal(err)
		}
		text := string(data)
		return c.Send(&tele.Photo{File: tele.FromDisk("images/support.jpg"), Caption: text})
	})

	b.Handle(&question1Btn, func(c tele.Context) error {
		data, err := os.ReadFile("messages/faq_answers/1.txt")
		if err != nil {
			log.Fatal(err)
		}
		text := string(data)
		return c.Send(text)
	})

	b.Handle(&question2Btn, func(c tele.Context) error {
		data, err := os.ReadFile("messages/faq_answers/2.txt")
		if err != nil {
			log.Fatal(err)
		}
		text := string(data)
		return c.Send(text)
	})

	b.Handle(&question3Btn, func(c tele.Context) error {
		data, err := os.ReadFile("messages/faq_answers/3.txt")
		if err != nil {
			log.Fatal(err)
		}
		text := string(data)
		return c.Send(text)
	})

	b.Handle(&question4Btn, func(c tele.Context) error {
		data, err := os.ReadFile("messages/faq_answers/4.txt")
		if err != nil {
			log.Fatal(err)
		}
		text := string(data)
		return c.Send(text)
	})

	b.Handle(&question5Btn, func(c tele.Context) error {
		data, err := os.ReadFile("messages/faq_answers/5.txt")
		if err != nil {
			log.Fatal(err)
		}
		text := string(data)
		return c.Send(text)
	})

	b.Handle(&question6Btn, func(c tele.Context) error {
		data, err := os.ReadFile("messages/faq_answers/6.txt")
		if err != nil {
			log.Fatal(err)
		}
		text := string(data)
		return c.Send(text)
	})

	b.Start()
}
