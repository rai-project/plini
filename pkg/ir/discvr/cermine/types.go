package cermine

type StringName struct {
	Surname    string `xml:"surname" json:"surname"`
	GivenNames string `xml:"given-names" json:"given_names"`
}

type Ref struct {
	Id           string       `xml:"id,attr" json:"id"`
	StringNames  []StringName `xml:"mixed-citation>string-name" json:"string_names"`
	ArticleTitle string       `xml:"mixed-citation>article-title" json:"article_title"`
	Source       string       `xml:"mixed-citation>source" json:"source"`
	Year         string       `xml:"mixed-citation>year" json:"year"`
	Volume       string       `xml:"mixed-citation>volume" json:"volume"`
}

type Subsection struct {
	Title      string   `xml:"title" json:"title"`
	Id         string   `xml:"id,attr" json:"id"`
	Paragraphs []string `xml:"p" json:"paragraphs"`
}

type Section struct {
	Title       string       `xml:"title" json:"title"`
	Id          string       `xml:"id,attr" json:"id"`
	Paragraphs  []string     `xml:"p" json:"paragraphs"`
	Subsections []Subsection `xml:"sec" json:"subsections"`
}

type Article struct {
	Title      string    `xml:"front>article-meta>title-group>article-title" json:"title"`
	Abstract   string    `xml:"front>article-meta>abstract>p" json:"abstract"`
	Year       int       `xml:"front>article-meta>pub-date>year" json:"year"`
	Authors    []string  `xml:"front>article-meta>contrib-group>contrib>string-name" json:"authors"`
	Sections   []Section `xml:"body>sec" json:"sections"`
	References []Ref     `xml:"back>ref-list>ref" json:"references"`
}
