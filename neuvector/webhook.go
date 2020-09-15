package neuvector

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"
)

// Webhook event types
const (
	EventRegistryScanReport = "Registry.Scan.Report"

	EventContainerScanReport = "Container.Scan.Report"
)

// WebhookRequest represents an inbound webhook request from the NeuVector controller
type WebhookRequest struct {
	Text     string `json:"text"`
	Username string `json:"username"`
}

type WebhookData struct {
	Event  string
	Level  string
	Values map[string]string
}

// Webhook keys

const (
	webhookKeyName            = "name"
	webhookKeyLevel           = "level"
	webhookKeyRegistry        = "registry"
	webhookKeyRegistryName    = "registry_name"
	webhookKeyImageId         = "image_id"
	webhookKeyWorkloadId      = "workload_id"
	webhookKeyWorkloadName    = "workload_name"
	webhookKeyWorkloadImage   = "workload_image"
	webhookKeyWorkloadService = "workload_service"
	webhookKeyRepository      = "repository"
	webhookKeyTag             = "tag"
)

type WebhookRegistryScanReportData struct {
	Registry     string
	RegistryName string
	ImageId      string
	Repository   string
	Tag          string
}

type WebhookContainerScanReportData struct {
	WorkloadId      string
	WorkloadName    string
	WorkloadImage   string
	WorkloadService string
}

type WebhookHandler struct {
}

var webhookTextRe = regexp.MustCompile(`^\*(?P<audit>[^\*\n]*)\*\n(?P<event>[^\n]*)\n>>>\s+(?P<values>[^\n]*)$`)

const (
	webhookTextReAudit  = "audit"
	webhookTextReEvent  = "event"
	webhookTextReValues = "values"
)

func (w *WebhookRequest) ToWebhookData() (*WebhookData, error) {
	groupNames := webhookTextRe.SubexpNames()

	matches := webhookTextRe.FindAllStringSubmatch(w.Text, -1)
	if len(matches) != 1 {
		return nil, fmt.Errorf("failed to parse webhook")
	}

	if groupNames[3] != webhookTextReValues {
		return nil, fmt.Errorf("failed to parse webhook")
	}

	values := matches[0][3]

	valuesMap := map[string]string{}

	for _, kvPair := range strings.Split(values, ",") {
		splitKvPair := strings.SplitN(kvPair, "=", 2)
		valuesMap[splitKvPair[0]] = splitKvPair[1]
	}

	event, err := requireKeyInWebhook(valuesMap, webhookKeyName)
	if err != nil {
		return nil, err
	}

	level, err := requireKeyInWebhook(valuesMap, webhookKeyLevel)
	if err != nil {
		return nil, err
	}

	wd := &WebhookData{
		Event:  event,
		Level:  level,
		Values: valuesMap,
	}

	return wd, nil
}

func requireKeyInWebhook(m map[string]string, key string) (string, error) {
	value, hasKey := m[key]
	if !hasKey {
		return "", fmt.Errorf("key %s not available in webhook", key)
	}
	return value, nil
}

func (w *WebhookData) ToWebhookRegistryScanReportData() (*WebhookRegistryScanReportData, error) {
	if w.Event != EventRegistryScanReport {
		return nil, fmt.Errorf("webhook data is not a registry scan report event")
	}

	registry, err := requireKeyInWebhook(w.Values, webhookKeyRegistry)
	if err != nil {
		return nil, err
	}

	registryName, err := requireKeyInWebhook(w.Values, webhookKeyRegistryName)
	if err != nil {
		return nil, err
	}

	imageId, err := requireKeyInWebhook(w.Values, webhookKeyImageId)
	if err != nil {
		return nil, err
	}

	repository, err := requireKeyInWebhook(w.Values, webhookKeyRepository)
	if err != nil {
		return nil, err
	}

	tag, err := requireKeyInWebhook(w.Values, webhookKeyTag)
	if err != nil {
		return nil, err
	}

	wd := &WebhookRegistryScanReportData{
		Registry:     registry,
		RegistryName: registryName,
		ImageId:      imageId,
		Repository:   repository,
		Tag:          tag,
	}

	return wd, nil
}

func (w *WebhookData) ToWebhookContainerScanReportData() (*WebhookContainerScanReportData, error) {
	if w.Event != EventContainerScanReport {
		return nil, fmt.Errorf("webhook data is not a container scan report event")
	}

	workloadId, err := requireKeyInWebhook(w.Values, webhookKeyWorkloadId)
	if err != nil {
		return nil, err
	}

	workloadName, err := requireKeyInWebhook(w.Values, webhookKeyWorkloadName)
	if err != nil {
		return nil, err
	}

	workloadImage, err := requireKeyInWebhook(w.Values, webhookKeyWorkloadImage)
	if err != nil {
		return nil, err
	}

	workloadService, err := requireKeyInWebhook(w.Values, webhookKeyWorkloadService)
	if err != nil {
		return nil, err
	}

	wd := &WebhookContainerScanReportData{
		WorkloadId:      workloadId,
		WorkloadName:    workloadName,
		WorkloadImage:   workloadImage,
		WorkloadService: workloadService,
	}

	return wd, nil
}

var _ http.Handler = &WebhookHandler{}

func (h *WebhookHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()

	content := new(WebhookRequest)

	err := d.Decode(&content)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	log.Println(*content)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
}
