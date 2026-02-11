package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/marisasha/kinolog/pkg/models"
)

func GetMovieInfoFromAI(title *string, year *int) (*models.Movie, error) {
	apiKey := os.Getenv("GROQ_API_KEY")

	prompt := fmt.Sprintf(`Найди информацию о фильме или сериале "%s" %d.
Верни СТРОГО JSON без пояснений и без обратных кавычек :

{
  "description": "string ( 200-250 символов)",
  "poster_url": "string (ссылка на постер к фильму или сериалу )",
  "title": "string (название на русском)",
  "type": "film или serial",
  "year": number (год выпуска),
  "actors": [
    {
      "bio_url": "string (URL на страницу биографии)",
      "first_name": "string",
      "last_name": "string",
      "role": "actor или director"
    }
  ](минимум 3 актера и обязательно режисера )
}`, *title, *year)

	type ActorResponse struct {
		BioURL    string `json:"bio_url"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Role      string `json:"role"`
	}

	type MovieResponse struct {
		Description string          `json:"description"`
		PosterURL   string          `json:"poster_url"`
		Title       string          `json:"title"`
		Type        string          `json:"type"`
		Year        int             `json:"year"`
		Actors      []ActorResponse `json:"actors"`
	}

	url := os.Getenv("AI_URL") // Groq api url

	requestBody := map[string]interface{}{
		"model": os.Getenv("AI_MODEL"), //Groq api model
		"messages": []map[string]string{
			{
				"role":    "user",
				"content": prompt,
			},
		},
		"temperature": 0.7,
		"max_tokens":  500,
		"response_format": map[string]string{
			"type": "json_object",
		},
	}

	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		return nil, fmt.Errorf("ошибка маршалинга JSON: %w", err)
	}

	req, err := http.NewRequest("POST", url, strings.NewReader(string(jsonData)))
	if err != nil {
		return nil, fmt.Errorf("ошибка создания запроса: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("ошибка отправки запроса: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("ошибка чтения ответа: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("ошибка API: %d - %s", resp.StatusCode, string(body))
	}

	type GroqResponse struct {
		Choices []struct {
			Message struct {
				Content string `json:"content"`
			} `json:"message"`
		} `json:"choices"`
	}

	var groqResp GroqResponse
	if err := json.Unmarshal(body, &groqResp); err != nil {
		return nil, err
	}

	if len(groqResp.Choices) == 0 {
		return nil, errors.New("Пустой ответ API")
	}

	content := groqResp.Choices[0].Message.Content

	content = strings.TrimSpace(content)
	content = strings.TrimPrefix(content, "```json")
	content = strings.TrimPrefix(content, "```")
	content = strings.TrimSuffix(content, "```")
	content = strings.TrimSpace(content)

	var movieResp MovieResponse
	if err := json.Unmarshal([]byte(content), &movieResp); err != nil {
		return nil, err
	}

	movie := &models.Movie{
		Title:       movieResp.Title,
		Type:        movieResp.Type,
		Year:        movieResp.Year,
		Description: movieResp.Description,
		PosterURL:   movieResp.PosterURL,
		Status:      nil,
		Mark:        nil,
		Review:      nil,
	}

	for _, actorResp := range movieResp.Actors {
		actor := models.MovieActor{
			FirstName: actorResp.FirstName,
			LastName:  actorResp.LastName,
			Role:      actorResp.Role,
			BioUrl:    actorResp.BioURL,
		}
		movie.Actors = append(movie.Actors, actor)
	}

	return movie, nil
}
