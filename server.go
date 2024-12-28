package main

import (
	"html/template"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/gomarkdown/markdown"
	"github.com/labstack/echo/v4"
)

type BlogPost struct {
	Title     string
	Slug      string
	Content   template.HTML
	CreatedAt time.Time
}

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func loadBlogPosts() ([]BlogPost, error) {
	var posts []BlogPost
	files, err := os.ReadDir("blog-posts")
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".md") {
			content, err := os.ReadFile(filepath.Join("blog-posts", file.Name()))
			if err != nil {
				continue
			}

			lines := strings.Split(string(content), "\n")
			var title string
			for _, line := range lines {
				if strings.HasPrefix(line, "# ") {
					title = strings.TrimPrefix(line, "# ")
					break
				}
			}

			slug := strings.TrimSuffix(file.Name(), ".md")
			html := markdown.ToHTML(content, nil, nil)

			posts = append(posts, BlogPost{
				Title:     title,
				Slug:      slug,
				Content:   template.HTML(html),
				CreatedAt: time.Now(),
			})
		}
	}

	sort.Slice(posts, func(i, j int) bool {
		return posts[i].CreatedAt.After(posts[j].CreatedAt)
	})

	return posts, nil
}

func main() {
	e := echo.New()

	t := &Template{
		templates: template.Must(template.ParseGlob("templates/*.html")),
	}

	e.Renderer = t

	e.GET("/", func(c echo.Context) error {
		posts, err := loadBlogPosts()
		if err != nil {
			return err
		}
		return c.Render(http.StatusOK, "index.html", posts)
	})

	e.GET("/blog/:slug", func(c echo.Context) error {
		slug := c.Param("slug")
		posts, err := loadBlogPosts()
		if err != nil {
			return err
		}

		for _, post := range posts {
			if post.Slug == slug {
				return c.Render(http.StatusOK, "blog-post.html", post)
			}
		}

		return echo.NewHTTPError(http.StatusNotFound, "Blog post not found")
	})

	e.Static("/", ".")

	e.Logger.Fatal(e.Start(":1323"))
}
