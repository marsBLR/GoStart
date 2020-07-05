package routes

import (
	"blog-Go_SR/db/documents"
	"blog-Go_SR/models"
	"blog-Go_SR/session"
	"blog-Go_SR/utils"
	"fmt"
	"net/http"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"gopkg.in/mgo.v2"
)

func EditHandler(rnd render.Render, r *http.Request, params martini.Params, db *mgo.Database, s *session.Session) {
	if !s.IsAuthorized {
		rnd.Redirect("/")
	}
	postsCollection := db.C("posts")

	id := params["id"]
	postDocument := documents.PostDocument{}
	err := postsCollection.FindId(id).One(&postDocument)
	if err != nil {
		rnd.Redirect("/")
		return
	}
	post := models.Post{postDocument.Id, postDocument.Title, postDocument.ContentHtml, postDocument.ContentMarkdown}

	model := models.EditPostModel{}
	model.IsAuthorized = s.IsAuthorized
	model.Post = post
	rnd.HTML(200, "write", model)
}

func ViewHandler(rnd render.Render, r *http.Request, params martini.Params, db *mgo.Database, s *session.Session) {
	postsCollection := db.C("posts")

	id := params["id"]
	postDocument := documents.PostDocument{}
	err := postsCollection.FindId(id).One(&postDocument)
	if err != nil {
		rnd.Redirect("/")
		return
	}
	post := models.Post{postDocument.Id, postDocument.Title, postDocument.ContentHtml, postDocument.ContentMarkdown}

	model := models.ViewPostModel{}
	model.IsAuthorized = s.IsAuthorized
	model.Post = post
	rnd.HTML(200, "view", model)
}

func WriteHandler(rnd render.Render, s *session.Session) {
	if !s.IsAuthorized {
		rnd.Redirect("/")
	}
	model := models.EditPostModel{}
	model.IsAuthorized = s.IsAuthorized
	model.Post = models.Post{}
	rnd.HTML(200, "write", model)
}

func DeleteHandler(rnd render.Render, r *http.Request, params martini.Params, db *mgo.Database, s *session.Session) {
	if !s.IsAuthorized {
		rnd.Redirect("/")
	}
	id := params["id"]
	if id == "" {
		rnd.Redirect("/")
	}

	postsCollection := db.C("posts")
	postsCollection.RemoveId(id)

	rnd.Redirect("/")
}

func SavePostHandler(rnd render.Render, r *http.Request, db *mgo.Database, s *session.Session) {
	if !s.IsAuthorized {
		rnd.Redirect("/")
	}
	id := r.FormValue("id")
	title := r.FormValue("title")
	contentMarkdown := r.FormValue("content")
	contentHtml := utils.ConvertMaarkDownToHtml(contentMarkdown)

	postDocument := documents.PostDocument{id, title, contentHtml, contentMarkdown}

	postsCollection := db.C("posts")

	if id != "" {
		postsCollection.UpdateId(id, postDocument)

	} else {
		fmt.Println("next func Generate")
		id = utils.GenerateId()
		postDocument.Id = id
		postsCollection.Insert(postDocument)
	}

	rnd.Redirect("/")
}

func GetHtmlHandler(rnd render.Render, r *http.Request) {
	md := r.FormValue("md")
	html := utils.ConvertMaarkDownToHtml(md)

	rnd.JSON(200, map[string]interface{}{"html": html})
}
