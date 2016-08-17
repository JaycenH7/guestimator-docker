package main

import (
	"log"

	"github.com/mrap/guestimator/db"
	"github.com/mrap/guestimator/models"
	"github.com/mrap/guestimator/parser"
	"github.com/mrap/guestimator/server"
)

// DB - development database
var DB = db.DevDB

func main() {
	// if _, err := DB.Exec("DELETE FROM questions;DELETE FROM wikipages;"); err != nil {
	// 	log.Fatalln(err)
	// }
	//
	// FetchAndSaveTopWikis()

	questions, err := GetLastestQuestions(5)
	if err != nil {
		log.Fatalln(err)
	}

	engine := server.NewMatchHandler()
	server.AddMatch("test", 1, questions)
	engine.Run(":3000")
}

// FetchAndSaveTopWikis - function
func FetchAndSaveTopWikis() {
	wikis := parser.FetchTopWikis()
	for _, wp := range wikis {
		var err error

		if err = models.CreateWikipage(DB, &wp); err != nil {
			log.Fatalln(err)
		}
		log.Println("Created wikipage:", wp.Title)

		questions := wp.ExtractQuestions()
		log.Printf("Creating %d questions", len(questions))

		for _, q := range questions {
			q.WikipageID = wp.ID
			if err = models.CreateQuestion(DB, &q); err != nil {
				log.Fatalln(err)
			}
		}
	}

}

// GetLastestQuestions - function
func GetLastestQuestions(limit int) ([]models.Question, error) {
	questions := []models.Question{}
	return questions, DB.Model(&questions).Column("question.*", "Wikipage").Limit(limit).Select()
}
