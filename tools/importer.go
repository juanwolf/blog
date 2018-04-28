package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
	"strings"
)

// Post ...
type Post struct {
	ID         int         `json:"id"`
	PubDate    string      `json:"pub_date"`
	Image      string      `json:"image"`
	Keywords   interface{} `json:"keywords"`
	KeywordsEn interface{} `json:"keywords_en"`
	KeywordsFr interface{} `json:"keywords_fr"`
	Title      string      `json:"title"`
	TitleEn    string      `json:"title_en"`
	TitleFr    string      `json:"title_fr"`
	Text       string      `json:"text"`
	TextEn     string      `json:"text_en"`
	TextFr     string      `json:"text_fr"`
	Slug       string      `json:"slug"`
	SlugEn     string      `json:"slug_en"`
	SlugFr     string      `json:"slug_fr"`
	Category   struct {
		ID            int    `json:"id"`
		Name          string `json:"name"`
		NameEn        string `json:"name_en"`
		NameFr        string `json:"name_fr"`
		Description   string `json:"description"`
		DescriptionEn string `json:"description_en"`
		DescriptionFr string `json:"description_fr"`
		Slug          string `json:"slug"`
		SlugEn        string `json:"slug_en"`
		SlugFr        string `json:"slug_fr"`
	} `json:"category"`
	Tags []struct {
		ID            int    `json:"id"`
		Name          string `json:"name"`
		NameEn        string `json:"name_en"`
		NameFr        string `json:"name_fr"`
		Description   string `json:"description"`
		DescriptionEn string `json:"description_en"`
		DescriptionFr string `json:"description_fr"`
		Slug          string `json:"slug"`
		SlugEn        string `json:"slug_en"`
		SlugFr        string `json:"slug_fr"`
	} `json:"tags"`
}

func writePost(post *Post) {
	date := strings.Split(post.PubDate, "T")[0]
	filenameEn := fmt.Sprintf("../content/post/%s-%s.en.md", date, post.SlugEn)
	filenameFr := fmt.Sprintf("../content/post/%s-%s.fr.md", date, post.SlugEn)
	pandocCmdEn := fmt.Sprintf("echo \"%s\" | pandoc -f html -t gfm", post.TextEn)
	pandocCmdFr := fmt.Sprintf("echo \"%s\" | pandoc -f html -t gfm", post.TextFr)
	markdownEn, _ := exec.Command("bash", "-c", pandocCmdEn).Output()

	markdownFr, _ := exec.Command("bash", "-c", pandocCmdFr).Output()

	metadataTemplate := `
---
title: %s
date: %s
tags: [%s]
categories: ["%s"]

draft: false
author: "Jean-Loup Adde"
---

%s
`
	tagNamesEn := []string{}
	tagNamesFr := []string{}

	for _, tag := range post.Tags {
		tagNamesEn = append(tagNamesEn, "\""+tag.NameEn+"\"")
		tagNamesFr = append(tagNamesFr, "\""+tag.NameFr+"\"")
	}

	postEn := fmt.Sprintf(metadataTemplate, post.TitleEn, date,
		strings.Join(tagNamesEn, ", "), post.Category.NameEn, markdownEn)

	postFr := fmt.Sprintf(metadataTemplate, post.TitleFr, date,
		strings.Join(tagNamesFr, ", "), post.Category.NameFr, markdownFr)

	ioutil.WriteFile(filenameFr, []byte(postFr), 0664)
	ioutil.WriteFile(filenameEn, []byte(postEn), 0664)
}

func main() {
	url := "https://blog.juanwolf.fr/api/posts/?format=json"
	response, _ := http.Get(url)
	postsJSON, _ := ioutil.ReadAll(response.Body)
	var postsArray []Post
	if err := json.Unmarshal(postsJSON, &postsArray); err != nil {
		log.Fatal(err)
	}
	for _, post := range postsArray {
		writePost(&post)
	}
}
