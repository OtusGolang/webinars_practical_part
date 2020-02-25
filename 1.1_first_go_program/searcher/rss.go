package searcher

import "strings" // Работа со строками

/* rss - программное представление XML вида:

<?xml version="1.0" encoding="utf-8" ?>
<rss version="2.0">
<channel>
<title>Яндекс.Новости: Здоровье</title>
<link>https://news.yandex.ru/health.html?from=rss</link>
...
<item>
...
</item>
</channel>
</rss>
*/
type rss struct {
	Channel channel `xml:"channel"` // Структурные теги, которые использует "encoding/xml"
}

type channel struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	Items       []item `xml:"item"`
}

type item struct {
	Title       string    `xml:"title"`
	Link        string    `xml:"link"`
	Description string    `xml:"description"`
	GUID        string    `xml:"guid"`
	Enclosure   enclosure `xml:"enclosure"`
	PubDate     string    `xml:"pubDate"`
}

// HasKeyword - метод структуры item;
// говорит, содержит ли элемент в своём заголовке или описании строку word
func (i item) HasKeyword(word string) bool {
	return strings.Contains(strings.ToLower(i.Title), word) ||
		strings.Contains(strings.ToLower(i.Description), word)
}

type enclosure struct {
	URL    string `xml:"url,attr"`
	Length int64  `xml:"length,attr"`
	Type   string `xml:"type,attr"`
}
