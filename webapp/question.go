package webapp

import (
	"fmt"
	"html/template"
	"net/http"
	"time"

	"appengine"
	"appengine/datastore"
	"appengine/user"
)

type Quiz struct {
	time_created time.Time
	author string
	questions []Question
}

type Question struct {
	time_created time.Time
	category string
	text string
	sub_questions []SubQuestion
}

type SubQuestion struct {
	text string
	answer string
	points int32
}

func QuizKey(c appengine.Context) *datastore.Key {
	return datastore.NewKey(c, "Quiz", "quiz1", 0, nil)
}

func CreateQuestion(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)	
	sub_question = SubQuestion{
		text: r.FormValue("text"),
		answer: r.FormValue("answer"),
		points: r.FormValue("points"),
	}
	parts []SubQuestion{sub_question}
	question = Question{
		text: r.FormValue("text"),
		category: r.FormValue("category"),
		time_created: time.Now(),
		sub_questions: parts,
	}
	question_key := datastore.NewIncompleteKey(c, "Question", QuizKey(c))
	_, err := datastore.Put(c, question_key, &question)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/ask", http.StatusFound)
}

func AskQuestion(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	q := datastore.NewQuery("Question").Ancestor(QuizKey(c))
	.Order("-time_created").Limit(10)
	if u == nil {
		url, err := user.LoginURL(c, r.URL.String())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Location", url)
		w.WriteHeader(http.StatusFound)
		return
	}
	fmt.Fprintf(w, "Hello, %v", u)
}

var questionTemplate = template.Must(template.New("question").Parse(`
<html>
  <head>
    <title>Add Question</title>
  </head>
  <body>
    {{range .}}
      <pre>{{.text}}</pre>
    {{end}}
    <form action="/createquestion" method="post">
      <div><textarea name="text" rows="3" cols="60"></textarea></div>
      <div><textarea name="answer" rows="3" cols="60"></textarea></div>
      <div><textarea name="points" rows="3" cols="60"></textarea></div>
      <div><textarea name="category" rows="3" cols="60"></textarea></div>
      <div><input type="submit" value="Sign Question"></div>
    </form>
  </body>
</html>
`))
