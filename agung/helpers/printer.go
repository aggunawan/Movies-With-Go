package helpers

import (
	"anggafirdaus10/agung/entities"
	"fmt"
	"github.com/rodaine/table"
	"strings"
	"time"
)

func PrintMoviesTable(movies []entities.Movie) {
	output := table.New("ID", "Title", "Director", "Genre", "Year", "Country")

	if len(movies) == 0 {
		output.AddRow("Empty")
	}

	for _, m := range movies {
		var genres []string
		date, _ := time.Parse(time.RFC3339, m.ReleasedAt)

		for _, genre := range m.Genres {
			genres = append(genres, genre.Name)
		}

		output.AddRow(
			m.Id,
			m.Title,
			m.Director.Name,
			strings.Join(genres, ", "),
			fmt.Sprintf("%s %s %s", date.Format("02"), date.Format("Jan"), date.Format("2006")),
			m.Country.Name,
		)
	}

	output.Print()
}

func PrintGenresTable(genres []entities.Genre) {
	output := table.New("ID", "Name")

	if len(genres) == 0 {
		output.AddRow("Empty")
	}

	for _, g := range genres {
		output.AddRow(g.Id, g.Name)
	}

	output.Print()
}