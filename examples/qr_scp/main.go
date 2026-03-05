// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

// Package main demonstrates a minimal Query/Retrieve SCP.
// It supports C-ECHO, C-FIND, C-MOVE, and C-GET with an in-memory index.
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"os/signal"
	"path/filepath"
	"regexp"
	"strings"
	"syscall"
	"time"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/parser"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/uid"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
	"github.com/cocosip/go-dicom/pkg/network/dimse"
	"github.com/cocosip/go-dicom/pkg/network/server"
)

var (
	port    = flag.Int("port", 11113, "Port to listen on")
	dataDir = flag.String("data-dir", "", "Optional DICOM folder used to build query index")
	verbose = flag.Bool("verbose", false, "Enable verbose logging")
)

type qrRecord struct {
	PatientName       string
	PatientID         string
	StudyInstanceUID  string
	StudyDate         string
	AccessionNumber   string
	Modality          string
	SeriesInstanceUID string
	SOPInstanceUID    string
	SOPClassUID       string
	SourcePath        string
}

type qrRepository struct {
	records []qrRecord
}

func main() {
	flag.Parse()

	records, err := loadRecords(*dataDir, *verbose)
	if err != nil {
		log.Fatalf("failed to load query index: %v", err)
	}
	if len(records) == 0 {
		records = sampleRecords()
	}

	repo := &qrRepository{records: records}

	fmt.Printf("=== DICOM Query/Retrieve SCP Example ===\n")
	fmt.Printf("Port:       %d\n", *port)
	fmt.Printf("Records:    %d\n", len(records))
	if *dataDir == "" {
		fmt.Printf("Data dir:   (built-in sample data)\n")
	} else {
		fmt.Printf("Data dir:   %s\n", *dataDir)
	}
	fmt.Println("Supports:   C-ECHO, C-FIND, C-MOVE, C-GET")
	fmt.Println()

	srv := server.New(
		server.WithPort(*port),
		server.WithAssociationTimeout(30*time.Second),
		server.WithRequestTimeout(30*time.Second),
	)

	srv.SetCEchoHandler(func(_ context.Context, req *dimse.CEchoRequest) (*dimse.CEchoResponse, error) {
		if *verbose {
			log.Printf("C-ECHO request messageID=%d", req.MessageID())
		}
		return dimse.NewCEchoResponseFromRequest(req, 0x0000), nil
	})

	srv.SetCFindHandler(repo.handleCFind)
	srv.SetCMoveHandler(repo.handleCMove)
	srv.SetCGetHandler(repo.handleCGet)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	errCh := make(chan error, 1)
	go func() {
		fmt.Printf("Listening on :%d\n", *port)
		errCh <- srv.ListenAndServe(ctx)
	}()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)

	select {
	case sig := <-sigCh:
		fmt.Printf("\nReceived signal: %v\n", sig)
		shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer shutdownCancel()
		if err := srv.Shutdown(shutdownCtx); err != nil {
			log.Printf("shutdown error: %v", err)
		}
	case err := <-errCh:
		if err != nil && err != context.Canceled {
			log.Fatalf("server stopped with error: %v", err)
		}
	}

	fmt.Println("Server stopped.")
}

func (r *qrRepository) handleCFind(_ context.Context, req *dimse.CFindRequest) ([]*dimse.CFindResponse, error) {
	matched := r.filterRecords(req.QueryLevel(), req.DataDataset())
	results := reduceResults(req.QueryLevel(), matched)

	if *verbose {
		log.Printf("C-FIND level=%s matched=%d", req.QueryLevel(), len(results))
	}

	responses := make([]*dimse.CFindResponse, 0, len(results)+1)
	for _, rec := range results {
		identifier := buildFindIdentifier(req.QueryLevel(), rec)
		responses = append(responses, dimse.NewCFindResponseFromRequest(req, 0xFF00, identifier))
	}

	responses = append(responses, dimse.NewCFindResponseFromRequest(req, 0x0000, nil))
	return responses, nil
}

func (r *qrRepository) handleCMove(_ context.Context, req *dimse.CMoveRequest) ([]*dimse.CMoveResponse, error) {
	matched := r.filterRecords(req.QueryLevel(), req.DataDataset())
	total := safeUint16(len(matched))

	if *verbose {
		log.Printf("C-MOVE level=%s destination=%s subops=%d", req.QueryLevel(), req.MoveDestination(), total)
	}

	if total == 0 {
		return []*dimse.CMoveResponse{dimse.NewCMoveResponseSuccess(req.MessageID(), req.AffectedSOPClassUID())}, nil
	}

	return []*dimse.CMoveResponse{
		dimse.NewCMoveResponsePending(req.MessageID(), req.AffectedSOPClassUID(), total, 0, 0, 0),
		dimse.NewCMoveResponsePending(req.MessageID(), req.AffectedSOPClassUID(), 0, total, 0, 0),
		dimse.NewCMoveResponseSuccess(req.MessageID(), req.AffectedSOPClassUID()),
	}, nil
}

func (r *qrRepository) handleCGet(_ context.Context, req *dimse.CGetRequest) ([]*dimse.CGetResponse, error) {
	matched := r.filterRecords(req.QueryLevel(), req.DataDataset())
	total := safeUint16(len(matched))

	if *verbose {
		log.Printf("C-GET level=%s subops=%d", req.QueryLevel(), total)
	}

	if total == 0 {
		return []*dimse.CGetResponse{dimse.NewCGetResponseSuccess(req.MessageID(), req.AffectedSOPClassUID())}, nil
	}

	return []*dimse.CGetResponse{
		dimse.NewCGetResponsePending(req.MessageID(), req.AffectedSOPClassUID(), total, 0, 0, 0),
		dimse.NewCGetResponsePending(req.MessageID(), req.AffectedSOPClassUID(), 0, total, 0, 0),
		dimse.NewCGetResponseSuccess(req.MessageID(), req.AffectedSOPClassUID()),
	}, nil
}

func (r *qrRepository) filterRecords(level dimse.QueryRetrieveLevel, query *dataset.Dataset) []qrRecord {
	if len(r.records) == 0 {
		return nil
	}

	results := make([]qrRecord, 0, len(r.records))
	for _, rec := range r.records {
		if !matchByLevel(level, rec) {
			continue
		}
		if !matchQuery(query, rec) {
			continue
		}
		results = append(results, rec)
	}
	return results
}

func matchByLevel(level dimse.QueryRetrieveLevel, rec qrRecord) bool {
	switch level {
	case dimse.QueryRetrieveLevelPatient:
		return rec.PatientID != "" || rec.PatientName != ""
	case dimse.QueryRetrieveLevelStudy:
		return rec.StudyInstanceUID != ""
	case dimse.QueryRetrieveLevelSeries:
		return rec.StudyInstanceUID != "" && rec.SeriesInstanceUID != ""
	case dimse.QueryRetrieveLevelImage:
		return rec.StudyInstanceUID != "" && rec.SeriesInstanceUID != "" && rec.SOPInstanceUID != ""
	default:
		return rec.StudyInstanceUID != ""
	}
}

func matchQuery(query *dataset.Dataset, rec qrRecord) bool {
	if query == nil {
		return true
	}

	if !matchString(queryValue(query, tag.PatientName), rec.PatientName) {
		return false
	}
	if !matchString(queryValue(query, tag.PatientID), rec.PatientID) {
		return false
	}
	if !matchString(queryValue(query, tag.StudyInstanceUID), rec.StudyInstanceUID) {
		return false
	}
	if !matchString(queryValue(query, tag.SeriesInstanceUID), rec.SeriesInstanceUID) {
		return false
	}
	if !matchString(queryValue(query, tag.SOPInstanceUID), rec.SOPInstanceUID) {
		return false
	}
	if !matchString(queryValue(query, tag.AccessionNumber), rec.AccessionNumber) {
		return false
	}
	if !matchString(queryValue(query, tag.Modality), rec.Modality) {
		return false
	}
	if !matchDate(queryValue(query, tag.StudyDate), rec.StudyDate) {
		return false
	}

	return true
}

func reduceResults(level dimse.QueryRetrieveLevel, records []qrRecord) []qrRecord {
	if len(records) == 0 {
		return nil
	}

	results := make([]qrRecord, 0, len(records))
	seen := make(map[string]struct{}, len(records))

	for _, rec := range records {
		key := dedupeKey(level, rec)
		if key == "" {
			continue
		}
		if _, exists := seen[key]; exists {
			continue
		}
		seen[key] = struct{}{}
		results = append(results, rec)
	}

	return results
}

func dedupeKey(level dimse.QueryRetrieveLevel, rec qrRecord) string {
	switch level {
	case dimse.QueryRetrieveLevelPatient:
		return rec.PatientID + "|" + rec.PatientName
	case dimse.QueryRetrieveLevelStudy:
		return rec.StudyInstanceUID
	case dimse.QueryRetrieveLevelSeries:
		return rec.StudyInstanceUID + "|" + rec.SeriesInstanceUID
	case dimse.QueryRetrieveLevelImage:
		return rec.SOPInstanceUID
	default:
		return rec.StudyInstanceUID
	}
}

func buildFindIdentifier(level dimse.QueryRetrieveLevel, rec qrRecord) *dataset.Dataset {
	ds := dataset.New()
	_ = ds.Add(element.NewString(tag.QueryRetrieveLevel, vr.CS, []string{string(level)}))

	addString(ds, tag.PatientName, vr.PN, rec.PatientName)
	addString(ds, tag.PatientID, vr.LO, rec.PatientID)
	addString(ds, tag.StudyInstanceUID, vr.UI, rec.StudyInstanceUID)
	addString(ds, tag.StudyDate, vr.DA, rec.StudyDate)
	addString(ds, tag.AccessionNumber, vr.SH, rec.AccessionNumber)
	addString(ds, tag.Modality, vr.CS, rec.Modality)
	addString(ds, tag.ModalitiesInStudy, vr.CS, rec.Modality)
	addString(ds, tag.SeriesInstanceUID, vr.UI, rec.SeriesInstanceUID)
	addString(ds, tag.SOPInstanceUID, vr.UI, rec.SOPInstanceUID)
	addString(ds, tag.SOPClassUID, vr.UI, rec.SOPClassUID)
	addString(ds, tag.RetrieveAETitle, vr.AE, "QRSCP")

	return ds
}

func loadRecords(dir string, logSkipped bool) ([]qrRecord, error) {
	dir = strings.TrimSpace(dir)
	if dir == "" {
		return nil, nil
	}

	info, err := os.Stat(dir)
	if err != nil {
		return nil, err
	}
	if !info.IsDir() {
		return nil, fmt.Errorf("%s is not a directory", dir)
	}

	records := make([]qrRecord, 0)
	walkErr := filepath.Walk(dir, func(path string, fileInfo os.FileInfo, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if fileInfo.IsDir() {
			return nil
		}

		result, err := parser.ParseFile(path)
		if err != nil {
			if logSkipped {
				log.Printf("skip non-DICOM file %s: %v", path, err)
			}
			return nil
		}

		rec := qrRecord{
			PatientName:       strings.TrimSpace(result.Dataset.TryGetString(tag.PatientName)),
			PatientID:         strings.TrimSpace(result.Dataset.TryGetString(tag.PatientID)),
			StudyInstanceUID:  strings.TrimSpace(result.Dataset.TryGetString(tag.StudyInstanceUID)),
			StudyDate:         strings.TrimSpace(result.Dataset.TryGetString(tag.StudyDate)),
			AccessionNumber:   strings.TrimSpace(result.Dataset.TryGetString(tag.AccessionNumber)),
			Modality:          strings.TrimSpace(result.Dataset.TryGetString(tag.Modality)),
			SeriesInstanceUID: strings.TrimSpace(result.Dataset.TryGetString(tag.SeriesInstanceUID)),
			SOPInstanceUID:    strings.TrimSpace(result.Dataset.TryGetString(tag.SOPInstanceUID)),
			SOPClassUID:       strings.TrimSpace(result.Dataset.TryGetString(tag.SOPClassUID)),
			SourcePath:        path,
		}

		if rec.StudyInstanceUID == "" {
			if logSkipped {
				log.Printf("skip file without StudyInstanceUID: %s", path)
			}
			return nil
		}
		if rec.SeriesInstanceUID == "" {
			rec.SeriesInstanceUID = rec.StudyInstanceUID + ".1"
		}
		if rec.SOPInstanceUID == "" {
			rec.SOPInstanceUID = rec.SeriesInstanceUID + ".1"
		}
		if rec.SOPClassUID == "" {
			rec.SOPClassUID = uid.SecondaryCaptureImageStorage.UID()
		}

		records = append(records, rec)
		return nil
	})

	if walkErr != nil {
		return nil, walkErr
	}

	return records, nil
}

func sampleRecords() []qrRecord {
	return []qrRecord{
		{
			PatientName:       "DOE^JOHN",
			PatientID:         "P001",
			StudyInstanceUID:  "1.2.840.113619.2.55.3.2831164357.781.170000001.1",
			StudyDate:         "20260220",
			AccessionNumber:   "ACC-1001",
			Modality:          "CT",
			SeriesInstanceUID: "1.2.840.113619.2.55.3.2831164357.781.170000001.2",
			SOPInstanceUID:    "1.2.840.113619.2.55.3.2831164357.781.170000001.3",
			SOPClassUID:       uid.CTImageStorage.UID(),
		},
		{
			PatientName:       "DOE^JOHN",
			PatientID:         "P001",
			StudyInstanceUID:  "1.2.840.113619.2.55.3.2831164357.781.170000001.1",
			StudyDate:         "20260220",
			AccessionNumber:   "ACC-1001",
			Modality:          "CT",
			SeriesInstanceUID: "1.2.840.113619.2.55.3.2831164357.781.170000001.2",
			SOPInstanceUID:    "1.2.840.113619.2.55.3.2831164357.781.170000001.4",
			SOPClassUID:       uid.CTImageStorage.UID(),
		},
		{
			PatientName:       "SMITH^JANE",
			PatientID:         "P002",
			StudyInstanceUID:  "1.2.840.113619.2.55.3.2831164357.781.170000002.1",
			StudyDate:         "20260224",
			AccessionNumber:   "ACC-1002",
			Modality:          "MR",
			SeriesInstanceUID: "1.2.840.113619.2.55.3.2831164357.781.170000002.2",
			SOPInstanceUID:    "1.2.840.113619.2.55.3.2831164357.781.170000002.3",
			SOPClassUID:       uid.MRImageStorage.UID(),
		},
	}
}

func queryValue(ds *dataset.Dataset, t *tag.Tag) string {
	if ds == nil {
		return ""
	}
	value, _ := ds.GetString(t)
	return strings.TrimSpace(value)
}

func matchString(pattern, value string) bool {
	pattern = strings.TrimSpace(pattern)
	if pattern == "" {
		return true
	}

	value = strings.TrimSpace(value)
	if value == "" {
		return false
	}

	patternUpper := strings.ToUpper(pattern)
	valueUpper := strings.ToUpper(value)

	if !strings.ContainsAny(patternUpper, "*?") {
		return patternUpper == valueUpper
	}

	reExpr := regexp.QuoteMeta(patternUpper)
	reExpr = strings.ReplaceAll(reExpr, "\\*", ".*")
	reExpr = strings.ReplaceAll(reExpr, "\\?", ".")
	re, err := regexp.Compile("^" + reExpr + "$")
	if err != nil {
		return false
	}
	return re.MatchString(valueUpper)
}

func matchDate(pattern, date string) bool {
	pattern = strings.TrimSpace(pattern)
	if pattern == "" {
		return true
	}

	date = strings.TrimSpace(date)
	if date == "" {
		return false
	}

	if strings.Contains(pattern, "-") {
		parts := strings.SplitN(pattern, "-", 2)
		start := strings.TrimSpace(parts[0])
		end := strings.TrimSpace(parts[1])
		if start != "" && date < start {
			return false
		}
		if end != "" && date > end {
			return false
		}
		return true
	}

	return matchString(pattern, date)
}

func addString(ds *dataset.Dataset, t *tag.Tag, valueVR *vr.VR, value string) {
	value = strings.TrimSpace(value)
	if value == "" {
		return
	}
	_ = ds.AddOrUpdate(element.NewString(t, valueVR, []string{value}))
}

func safeUint16(n int) uint16 {
	if n < 0 {
		return 0
	}
	if n > math.MaxUint16 {
		return math.MaxUint16
	}
	return uint16(n)
}
