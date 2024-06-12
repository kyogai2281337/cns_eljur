package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kyogai2281337/cns_eljur/pkg/sql/store"
)

type Server struct {
	App   *fiber.App
	Store store.Store
}

func NewServer(store store.Store) *Server {
	s := &Server{
		App:   fiber.New(),
		Store: store,
	}
	s.App.Use("/api", func(c *fiber.Ctx) error {
		return c.Next()
	})
	return s
}

func (s *Server) ServeHTTP(addr string) error {
	return s.App.Listen(addr)
}

// func (s *Server) AuthMWare(next fiber.Handler) fiber.Handler {
// 	// return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 	// 	session, err := s.sessionStore.Get(r, sessionName)
// 	// 	if err != nil {
// 	// 		s.error(w, r, http.StatusInternalServerError, err)
// 	// 		return
// 	// 	}
// 	// 	auth, ok := session.Values["auth"]
// 	// 	if !ok || !auth.(bool) {
// 	// 		s.error(w, r, http.StatusUnauthorized, errNotAuthed)
// 	// 		return
// 	// 	}
// 	// 	refresh, ok := session.Values["refresh"]
// 	// 	if !ok {
// 	// 		s.error(w, r, http.StatusUnauthorized, errNotAuthed)
// 	// 		return
// 	// 	}

// 	// 	uJWT, err := GetUserDataJWT(refresh.(string))
// 	// 	if err != nil {
// 	// 		s.error(w, r, http.StatusUnauthorized, err)
// 	// 		return
// 	// 	}
// 	// 	log.Printf("User JWT: %+v", uJWT)

// 	// 	if err := s.sessionStore.Save(r, w, session); err != nil {
// 	// 		s.error(w, r, http.StatusInternalServerError, err)
// 	// 		return
// 	// 	}

// 	// 	if uJWT.ID <= 0 {
// 	// 		s.error(w, r, http.StatusUnauthorized, errNotAuthed)
// 	// 		return
// 	// 	}
// 	// 	u, err := s.store.User().Find(uJWT.ID)
// 	// 	if err != nil {
// 	// 		s.error(w, r, http.StatusUnauthorized, errNotAuthed)
// 	// 		return
// 	// 	}
// 	// 	next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), ctxKeyUser, u)))
// 	// })
// }
