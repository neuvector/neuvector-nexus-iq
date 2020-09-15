package integration

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/neuvector/neuvector-nexus-iq/neuvector"
	"net/http"
	"time"
)

// Endpoint paths
const (
	webhookPath         = "/webhook"
	healthLivenessPath  = "/health/liveness"
	healthReadinessPath = "/health/readiness"
)

type Server struct {
	Config      RootConfig
	Integration neuVectorNexusIq

	httpServer *http.Server
	router     *http.ServeMux
}

var _ http.Handler = &Server{}

func NewServer(config *RootConfig) *Server {
	return &Server{
		Config: *config,
		Integration: &NeuVectorNexusIq{
			Config: *config,
		},
	}
}

func (s *Server) Setup() {
	r := http.NewServeMux()

	r.HandleFunc(webhookPath, s.handleWebhook)
	r.HandleFunc(healthLivenessPath, s.handleHealthLiveness)
	r.HandleFunc(healthReadinessPath, s.handleHealthReadiness)

	s.router = r
}

func (s *Server) Start() {
	if s.router == nil {
		s.Setup()
	}

	if s.Integration == nil {
		log.Fatalf("integration is not configured for server")
	}

	addr := fmt.Sprintf("%s:%d", s.Config.Address, s.Config.Port)

	log.Infof("Start NeuVector Nexus IQ integration server")
	s.httpServer = &http.Server{Addr: addr, Handler: s.router}
	err := s.httpServer.ListenAndServe()
	if err != nil {
		log.Errorf(err.Error())
	}

	log.Infof("NeuVector Nexus IQ integration server stopped")
}

func (s *Server) ShutdownWithTimeout(timeout time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	err := s.httpServer.Shutdown(ctx)

	if err != nil {
		return err
	}
	return nil
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

// handleWebhook is the http handler which receives the webhook request from the NeuVector controller
func (s *Server) handleWebhook(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "invalid method", http.StatusMethodNotAllowed)
	}

	// Decode JSON body
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()

	wr := new(neuvector.WebhookRequest)

	err := d.Decode(&wr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Process webhook request
	err = s.Integration.ProcessWebhookRequest(wr)
	if err != nil {
		log.Errorf(err.Error())

		// Error response
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Success response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
}

func (s *Server) handleHealthLiveness(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}

func (s *Server) handleHealthReadiness(w http.ResponseWriter, r *http.Request) {
	// FOLLOWUP respond readiness only when NeuVector and Nexus IQ are reachable, i.e. authenticate succeeds with both
	// parallel query nv and nx for availability implement isAvailable in the respective client libraries

	w.WriteHeader(200)
}
