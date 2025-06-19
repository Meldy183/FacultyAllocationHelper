package cookies

import (
	"errors"
	"log"
	"net/http"

	"gitlab.pg.innopolis.university/f.markin/fah/auth/internal/config"
)

type CookiesService struct {
	Cookie config.Cookies
}

func NewCookiesService(cookieConfig config.Cookies) *CookiesService {
	return &CookiesService{
		Cookie: cookieConfig,
	}
}

func (c *CookiesService) SetAccessTokenCookie(w http.ResponseWriter, token string) error {
	if token == "" {
		log.Printf("err: empty token")
		return errors.New("empty JWT token provided")
	}
	var a http.SameSite
	if c.Cookie.AccessToken.SameSite == "lax" {
		a = http.SameSiteLaxMode
	}

	cookie := &http.Cookie{
		Name:     c.Cookie.AccessToken.Name,
		Value:    token,
		HttpOnly: c.Cookie.AccessToken.HTTPOnly,
		Secure:   c.Cookie.AccessToken.Secure,
		//  подумать
		SameSite: http.SameSite(a),
		Path:     c.Cookie.AccessToken.Path,
		MaxAge:   c.Cookie.AccessToken.MaxAge,
	}
	http.SetCookie(w, cookie)
	return nil
}
func (c *CookiesService) SetRefreshTokenCookie(w http.ResponseWriter, token string) error {
	if token == "" {
		return errors.New("empty JWT token provided")
	}
	var a http.SameSite
	if c.Cookie.RefreshToken.SameSite == "strict" {
		a = http.SameSiteStrictMode
	}
	cookie := &http.Cookie{
		Name:     c.Cookie.RefreshToken.Name,
		Value:    token,
		HttpOnly: c.Cookie.RefreshToken.HTTPOnly,
		Secure:   c.Cookie.RefreshToken.Secure,
		// подумать
		SameSite: http.SameSite(a),
		Path:     c.Cookie.RefreshToken.Path,
		MaxAge:   c.Cookie.RefreshToken.MaxAge}
	http.SetCookie(w, cookie)
	return nil
}
func (c *CookiesService) ClearAuthCookies(w http.ResponseWriter) error {
	// есл  мы будем храниь refresh токены в бд  (как сессии) то тут можно будет добавить изменения
	if c.Cookie.AccessToken.Name == "" || c.Cookie.RefreshToken.Name == "" {
		return http.ErrNoCookie
	}

	http.SetCookie(w, &http.Cookie{
		Name:     c.Cookie.AccessToken.Name,
		Value:    "",
		HttpOnly: c.Cookie.AccessToken.HTTPOnly,
		Secure:   c.Cookie.AccessToken.Secure,
		SameSite: http.SameSiteLaxMode,
		Path:     c.Cookie.AccessToken.Path,
		MaxAge:   -1,
	})
	http.SetCookie(w, &http.Cookie{
		Name:     c.Cookie.RefreshToken.Name,
		Value:    "",
		HttpOnly: c.Cookie.RefreshToken.HTTPOnly,
		Secure:   c.Cookie.RefreshToken.Secure,
		SameSite: http.SameSiteStrictMode,
		Path:     c.Cookie.RefreshToken.Path,
		MaxAge:   -1})
	return nil
}
func (c *CookiesService) ExtractTokenFromCookie(r *http.Request, cookieName string) (string, error) {
	cookie, err := r.Cookie(cookieName)
	if err != nil {
		return "", err
	}
	if cookie == nil || cookie.Value == "" {
		return "", http.ErrNoCookie
	}
	return cookie.Value, nil
}
