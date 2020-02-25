package searcher

import (
	"encoding/xml" // Работа с XML
	"fmt"          // Работа с форматированным вводом-выводом
	"io"           // Работа с вводом-выводом
	"net/http"     // Работа с HTTP
	"sync"         // Работа с примитивами синхронизации
)

// Search выполняет конкурентный поиск новостей со словом keyword среди рассылок feeds.
// Результат записывается в out
func Search(feeds []string, keyword string, out io.Writer) {
	var outMu sync.Mutex

	// Механизм синхронизации группы горутин
	var wg sync.WaitGroup
	wg.Add(len(feeds))

	// Обрабатываем каждый URL в отдельной горутине
	for _, f := range feeds {
		go func(url string) {
			defer wg.Done()

			// Ошибки игнорируются, но в реальной разработке их стоит логировать
			// или выводить в отдельный канал
			_, _ = fmt.Fprintf(out, "Process %s...\n", url)
			_ = getFeedAndSearch(url, keyword, out, &outMu)
		}(f)
	}

	// Ждём, пока не будут обработаны все RSS-ленты
	wg.Wait()
}

// getFeedAndSearch получает XML по url, парсит RSS ленту и ищет в ней новости с keyword
func getFeedAndSearch(url, keyword string, out io.Writer, outMut sync.Locker) error {
	// Выполняем HTTP GET-запрос
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Декодируем XML в структуру
	var feed rss
	if err := xml.NewDecoder(resp.Body).Decode(&feed); err != nil {
		return err
	}

	// Проходимся по новостям
	for _, item := range feed.Channel.Items {
		if item.HasKeyword(keyword) {
			if err := safeWrite(fmt.Sprintf("\n%s (%s)\n", item.Title, item.Link), out, outMut); err != nil {
				return err
			}
		}
	}
	return nil
}

// safeWrite синхронизирует запись в out между горутинами
func safeWrite(str string, out io.Writer, outMut sync.Locker) (err error) {
	outMut.Lock()
	defer outMut.Unlock()
	_, err = fmt.Fprint(out, str)
	return
}
