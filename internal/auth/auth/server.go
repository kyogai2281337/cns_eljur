package apiserver

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/kyogai2281337/cns_eljur/internal/auth/model"
	"github.com/kyogai2281337/cns_eljur/internal/auth/store"

	"github.com/google/uuid"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

var (
	errEmailOrPassInvalid = errors.New("incorrect email or pass")
	errNotAuthed          = errors.New("user not authorized")
)

const (
	sessionName        = "journal_auth"
	ctxKeyUser  ctxKey = iota
	ctxKeyReqID
)

type ctxKey int8

type server struct {
	router       *mux.Router
	sessionStore sessions.Store
	store        store.Store
}

func NewServer(store store.Store, sessStore sessions.Store) *server {
	s := &server{
		router:       mux.NewRouter(),
		sessionStore: sessStore,
		store:        store,
	}
	s.configureRouter()
	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) configureRouter() {
	s.router.Use(s.setReqID)
	s.router.Use(s.logReq)
	s.router.Use(handlers.CORS(handlers.AllowedOrigins([]string{"*"})))
	s.router.HandleFunc("/signup", s.handleUserCreate()).Methods("POST")
	s.router.HandleFunc("/signin", s.handleSessionCreate()).Methods("POST")

	priv := s.router.PathPrefix("/private").Subrouter()
	priv.Use(s.authUser)
	priv.HandleFunc("/profile", s.HandleWhoami()).Methods("GET")
	priv.HandleFunc("/delete", s.HandleDelete()).Methods("GET")
	priv.HandleFunc("/logout", s.HandleLogout()).Methods("GET")
}

func (s *server) HandleLogout() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, err := s.sessionStore.Get(r, sessionName)
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}
		delete(session.Values, "user_id")
		if err := session.Save(r, w); err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

	}
}

func (s *server) HandleDelete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, err := s.sessionStore.Get(r, sessionName)
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}
		id, ok := session.Values["user_id"]
		if !ok || id == nil {
			s.error(w, r, http.StatusUnauthorized, errNotAuthed)
			return
		}
		var userID int64
		switch v := id.(type) {
		case int64:
			userID = v
		default:
			s.error(w, r, http.StatusUnauthorized, errNotAuthed)
			return
		}
		err = s.store.User().Delete(userID)
		if err != nil {
			s.error(w, r, http.StatusUnauthorized, errNotAuthed)
			return
		}
	}
}

func (s *server) setReqID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := uuid.New().String()
		w.Header().Set("X-Request-ID", id)
		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), ctxKeyReqID, id)))
	})
}

func (s *server) logReq(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[INFO]Started %s %s => %s || %s", r.Method, r.RequestURI, r.Context().Value(ctxKeyReqID), r.RemoteAddr)

		start := time.Now()
		rw := &ResWriter{w, http.StatusOK}
		next.ServeHTTP(rw, r)

		log.Printf("[%v]Completed with %s in %v => %s || %s", rw.Code, http.StatusText(rw.Code), time.Since(start), r.Context().Value(ctxKeyReqID), r.RemoteAddr)
	})

}

func (s *server) handleUserCreate() http.HandlerFunc {
	type request struct {
		Email string `json:"email"`
		Pass  string `json:"password"`
		First string `json:"first"`
		Last  string `json:"last"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}
		//
		// Здесь дописать маханику для конфига!!!
		//
		uRole, err := s.store.Role().Find(1)
		if err != nil {
			s.error(w, r, http.StatusNotFound, err)
			return
		}
		u := &model.User{
			Email:     req.Email,
			Pass:      req.Pass,
			FirstName: req.First,
			LastName:  req.Last,
			Role:      uRole,
		}
		if err := s.store.User().Create(u); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}
		u.Sanitize()
		s.respond(w, r, http.StatusCreated, u)
	}
}

func (s *server) handleSessionCreate() http.HandlerFunc {
	type request struct {
		Email string `json:"email"`
		Pass  string `json:"password"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		u, err := s.store.User().FindByEmail(req.Email)
		if err != nil || !u.ComparePass(req.Pass) {
			s.error(w, r, http.StatusUnauthorized, errEmailOrPassInvalid)
			return
		}

		session, err := s.sessionStore.Get(r, sessionName)
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}
		session.Values["auth"] = true

		// Генерация JWT-Access и JWT-Refresh
		session.Values["access"], session.Values["refresh"], err = GenAuthJWT(u)
		//pr := session.Values["access"]
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
		}
		if err := s.sessionStore.Save(r, w, session); err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		s.respond(w, r, http.StatusOK, map[string]interface{}{
			"status":  "OK",
			"auth":    true,
			"access":  session.Values["access"], // pr
			"refresh": session.Values["refresh"],
		})
	}
}

func (s *server) authUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, err := s.sessionStore.Get(r, sessionName)
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}
		auth, ok := session.Values["auth"]
		if !ok || !auth.(bool) {
			s.error(w, r, http.StatusUnauthorized, errNotAuthed)
			return
		}
		access, ok := session.Values["access"]
		if !ok {
			s.error(w, r, http.StatusUnauthorized, errNotAuthed)
			return
		}
		refresh, ok := session.Values["refresh"]
		if !ok {
			s.error(w, r, http.StatusUnauthorized, errNotAuthed)
			return
		}

		uJWT, err := ValidateViaJWT(access.(string), refresh.(string))
		if err != nil {
			s.error(w, r, http.StatusUnauthorized, err)
			return
		}

		session.Values["access"], err = uJWT.GenJWT("a")
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		if err := s.sessionStore.Save(r, w, session); err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		if uJWT.ID <= 0 {
			s.error(w, r, http.StatusUnauthorized, errNotAuthed)
			return
		}
		u, err := s.store.User().Find(uJWT.ID)
		if err != nil {
			s.error(w, r, http.StatusUnauthorized, errNotAuthed)
			return
		}
		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), ctxKeyUser, u)))
	})
}

func (s *server) HandleWhoami() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.respond(w, r, http.StatusOK, r.Context().Value(ctxKeyUser).(*model.User))
	}
}

func (s *server) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	s.respond(w, r, code, map[string]string{"error": err.Error()})
}

func (s *server) respond(w http.ResponseWriter, _ *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}